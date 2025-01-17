package installation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fluxninja/aperture/pkg/log"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/imdario/mergo"
	"golang.org/x/exp/slices"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/engine"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/releaseutil"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

var (
	valuesFile     string
	kubeConfig     string
	kubeRestConfig *rest.Config
	version        string
	latestVersion  string
	namespace      string
	kubeClient     client.Client
	timeout        int
)

const (
	apertureLatestVersion  = "latest"
	defaultNS              = "default"
	controller             = "controller"
	apertureController     = "aperture-controller"
	agent                  = "agent"
	apertureAgent          = "aperture-agent"
	istioConfig            = "istioconfig"
	istioConfigReleaseName = "aperture-envoy-filter"
)

// getTemplets loads CRDs, hooks and manifests from the Helm chart.
func getTemplets(chartName, releaseName string, order releaseutil.KindSortOrder) ([]chart.CRD, []*release.Hook, []releaseutil.Manifest, error) {
	chartURL := fmt.Sprintf("https://fluxninja.github.io/aperture/%s-%s.tgz", chartName, version)

	resp, err := http.Get(chartURL) //nolint
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to download chart: %s", err)
	}
	defer resp.Body.Close()

	ch, err := loader.LoadArchive(resp.Body)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to load chart: %s", err)
	}

	values := ch.Values
	if valuesFile != "" {
		values, err = chartutil.ReadValuesFile(valuesFile)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to read values: %s", err)
		}
	} else if releaseName == agent && slices.Compare(order, releaseutil.UninstallOrder) == 0 {
		values = map[string]interface{}{
			"agent": map[string]interface{}{
				"config": map[string]interface{}{
					"etcd": map[string]interface{}{
						"endpoints": []string{"dummy"},
					},
					"prometheus": map[string]interface{}{
						"address": "dummy",
					},
				},
			},
		}
	}

	renderedValues, err := chartutil.ToRenderValues(ch, values, chartutil.ReleaseOptions{Name: releaseName, Namespace: namespace}, chartutil.DefaultCapabilities)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to read values: %s", err)
	}

	files, err := engine.RenderWithClient(ch, renderedValues, kubeRestConfig)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to render chart: %s", err)
	}

	for k := range files {
		if strings.HasSuffix(k, "NOTES.txt") {
			delete(files, k)
		}
	}

	hooks, manifests, err := releaseutil.SortManifests(files, chartutil.DefaultVersionSet, order)

	return ch.CRDObjects(), hooks, manifests, err
}

// applyManifest creates/updates the generated manifest to Kubernetes.
func applyManifest(manifest string) error {
	unstructuredObject, err := prepareUnstructuredObject(manifest)
	if err != nil {
		return err
	}

	log.Info().Msgf("Applying - %s/%s", unstructuredObject.GetKind(), unstructuredObject.GetName())
	attempt := 0
	for attempt < 5 {
		attempt++
		if err = applyObjectToKubernetes(unstructuredObject); err != nil && strings.Contains(err.Error(), "no matches for kind") {
			time.Sleep(time.Second * time.Duration(attempt))
			continue
		}
		break
	}

	if err != nil {
		return fmt.Errorf("failed to apply - %s/%s, Error - '%s'", unstructuredObject.GetKind(), unstructuredObject.GetName(), err)
	}
	return nil
}

// applyObjectToKubernetes applies the given object to Kubernetes.
func applyObjectToKubernetes(unstructuredObject *unstructured.Unstructured) error {
	key := types.NamespacedName{
		Name:      unstructuredObject.GetName(),
		Namespace: unstructuredObject.GetNamespace(),
	}
	existing := unstructuredObject.DeepCopy()
	err := kubeClient.Get(context.Background(), key, existing)
	if err != nil && !apierrors.IsNotFound(err) {
		return err
	}

	if !apierrors.IsNotFound(err) {
		opts := cmpopts.IgnoreFields(
			metav1.ObjectMeta{},
			"Generation", "ResourceVersion", "SelfLink", "UID",
			"CreationTimestamp", "DeletionTimestamp", "DeletionGracePeriodSeconds",
			"OwnerReferences", "Finalizers",
		)

		// Check if there are any differences between the existing and the new object
		if !cmp.Equal(unstructuredObject, existing, opts) {
			err = mergo.Map(&unstructuredObject.Object, &existing.Object)
			if err != nil {
				return err
			}
			err = kubeClient.Patch(context.Background(), unstructuredObject, client.MergeFrom(existing))
		}
	} else {
		err = kubeClient.Create(context.Background(), unstructuredObject)
	}

	return err
}

// deleteManifest deletes the generated manifest from Kubernetes.
func deleteManifest(manifest string) error {
	unstructuredObject, err := prepareUnstructuredObject(manifest)
	if err != nil {
		return err
	}

	log.Info().Msgf("Deleting - %s/%s", unstructuredObject.GetKind(), unstructuredObject.GetName())

	err = kubeClient.Delete(context.Background(), unstructuredObject)
	if apierrors.IsNotFound(err) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("failed to delete - %s/%s, Error - '%s'", unstructuredObject.GetKind(), unstructuredObject.GetName(), err)
	}
	return nil
}

// manageNamespace creates namespace if not present.
func manageNamespace() error {
	ns := &corev1.Namespace{}
	err := kubeClient.Get(context.Background(), types.NamespacedName{Name: namespace}, ns)
	if apierrors.IsNotFound(err) {
		ns.Name = namespace
		if err = kubeClient.Create(context.Background(), ns); err != nil {
			return err
		}
	}

	return nil
}

// waitForHook waits for 1 successful execution of the hook.
func waitForHook(name string, ctx context.Context) error {
	job := &batchv1.Job{}
	for {
		err := kubeClient.Get(ctx, types.NamespacedName{Namespace: namespace, Name: name}, job)
		if err != nil || job.Status.Succeeded != 1 {
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}
	return nil
}

// prepareUnstructuredObject prepares unstructured.Unstructured from given YAML string.
func prepareUnstructuredObject(manifest string) (*unstructured.Unstructured, error) {
	content := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(manifest), &content)
	if err != nil {
		return nil, err
	}

	unstructuredObject := &unstructured.Unstructured{
		Object: content,
	}

	if unstructuredObject.GetNamespace() == "" {
		unstructuredObject.SetNamespace(namespace)
	}

	return unstructuredObject, nil
}

// handleInstall handles installation for given chart using given release name.
func handleInstall(chartName, releaseName string) error {
	crds, _, manifests, err := getTemplets(chartName, releaseName, releaseutil.InstallOrder)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, crd := range crds {
		if err = applyManifest(string(crd.File.Data)); err != nil {
			errs = append(errs, err)
		}
	}

	for _, manifest := range manifests {
		if err = applyManifest(manifest.Content); err != nil {
			errs = append(errs, err)
		}
	}

	for _, err := range errs {
		log.Error().Msg(err.Error())
	}

	if len(errs) != 0 {
		return fmt.Errorf("failed to complete install successfully")
	}
	return nil
}

// handleUnInstall handles uninstallation for given chart using given release name.
func handleUnInstall(chartName, releaseName string) error {
	crds, hooks, manifests, err := getTemplets(chartName, releaseName, releaseutil.UninstallOrder)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, hook := range hooks {
		log.Info().Msgf("Executing hook - %s", hook.Name)
		if err = applyManifest(hook.Manifest); err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		if err = waitForHook(hook.Name, ctx); err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				return fmt.Errorf("timed out waiting for pre-delete hook completion")
			}
			return err
		}

		if err = deleteManifest(hook.Manifest); err != nil {
			errs = append(errs, err)
		}

		if err = kubeClient.DeleteAllOf(
			context.Background(), &corev1.Pod{}, client.InNamespace(namespace), client.MatchingLabels{"job-name": hook.Name}); err != nil {
			errs = append(errs, err)
		}
	}

	for _, manifest := range manifests {
		if err = deleteManifest(manifest.Content); err != nil {
			errs = append(errs, err)
		}
	}

	for _, crd := range crds {
		if err = deleteManifest(string(crd.File.Data)); err != nil {
			errs = append(errs, err)
		}
	}

	for _, err := range errs {
		log.Error().Msg(err.Error())
	}

	if len(errs) != 0 {
		return fmt.Errorf("failed to complete uninstall successfully")
	}
	return nil
}
