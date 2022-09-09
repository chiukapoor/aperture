/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"net/http"

	"github.com/clarketm/json"
	"github.com/fluxninja/aperture/operator/api/v1alpha1"
	"github.com/fluxninja/aperture/pkg/config"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// AgentHooks injects the default spec of Aperture Agent in CR.
type AgentHooks struct {
	decoder *admission.Decoder
}

// Handle receives incomming requests from MutatingWebhook for newly created Pods and injects Agent container in them.
func (agentHooks *AgentHooks) Handle(ctx context.Context, req admission.Request) admission.Response {
	agent := &v1alpha1.Agent{}

	unmarshaller, err := config.KoanfUnmarshallerConstructor{}.NewKoanfUnmarshaller(req.Object.Raw)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}

	err = unmarshaller.Unmarshal(agent)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	updatedAgent, err := json.Marshal(agent)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}

	return admission.PatchResponseFromRaw(req.Object.Raw, updatedAgent)
}

// InjectDecoder injects the decoder.
func (agentHooks *AgentHooks) InjectDecoder(d *admission.Decoder) error {
	agentHooks.decoder = d
	return nil
}