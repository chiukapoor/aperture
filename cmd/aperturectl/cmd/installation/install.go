package installation

import (
	"fmt"

	"github.com/fluxninja/aperture/cmd/aperturectl/cmd/utils"
	"github.com/fluxninja/aperture/operator/api"
	_ "github.com/fluxninja/aperture/operator/api/agent/v1alpha1"
	_ "github.com/fluxninja/aperture/operator/api/controller/v1alpha1"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	InstallCmd.PersistentFlags().StringVar(&kubeConfig, "kube-config", "", "Path to the Kubernetes cluster config. Defaults to '~/.kube/config'")
	InstallCmd.PersistentFlags().StringVar(&valuesFile, "values-file", "", "Values YAML file containing parameters to customize the installation")
	InstallCmd.PersistentFlags().StringVar(&version, "version", apertureLatestVersion, "Version of the Aperture")
	InstallCmd.PersistentFlags().StringVar(&namespace, "namespace", "", "Namespace in which the component will be installed. Defaults to component name")

	InstallCmd.AddCommand(controllerInstallCmd)
	InstallCmd.AddCommand(agentInstallCmd)
	InstallCmd.AddCommand(istioConfigInstallCmd)
}

// InstallCmd is the command to install Aperture Controller and Aperture Agent on Kubernetes.
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Aperture components",
	Long: `
Use this command to install Aperture Controller and Agent on your Kubernetes cluster.`,
	SilenceErrors: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error

		kubeRestConfig, err = utils.GetKubeConfig(kubeConfig)
		if err != nil {
			return err
		}

		err = api.SchemeBuilder.AddToScheme(scheme.Scheme)
		if err != nil {
			return fmt.Errorf("failed to create Kubernetes client: %w", err)
		}
		kubeClient, err = client.New(kubeRestConfig, client.Options{
			Scheme: scheme.Scheme,
		})
		if err != nil {
			return fmt.Errorf("failed to create Kubernetes client: %w", err)
		}

		if namespace == "" {
			namespace = defaultNS
		}

		if err = manageNamespace(); err != nil {
			return err
		}

		latestVersion, err = utils.ResolveLatestVersion()
		if err != nil {
			return err
		}

		if version == "" || version == apertureLatestVersion {
			version = latestVersion
		}
		return nil
	},
}
