/*
Copyright 2020 The Tekton Authors

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

package common

import (
	"errors"
	"strings"

	"github.com/tektoncd/operator/pkg/apis/operator/v1alpha1"
	informer "github.com/tektoncd/operator/pkg/client/informers/externalversions/operator/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"knative.dev/pkg/apis"
)

const (
	PipelineNotReady       = "tekton-pipelines not ready"
	PipelineNotFound       = "tekton-pipelines not installed"
	TriggerNotReady        = "tekton-triggers not ready"
	TriggerNotFound        = "tekton-triggers not installed"
	NamespaceIgnorePattern = "^(openshift|kube)-|^open-cluster-management-agent-addon$|^open-cluster-management-agent$|^dedicated-admin$|^kube-node-lease$|^kube-public$|^kube-system$"
)

func PipelineReady(informer informer.TektonPipelineInformer) (*v1alpha1.TektonPipeline, error) {
	ppln, err := getPipelineRes(informer)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, errors.New(PipelineNotFound)
		}
		return nil, err
	}
	if isUpgradePending(ppln.GetStatus()) {
		return nil, v1alpha1.DEPENDENCY_UPGRADE_PENDING_ERR
	}
	if !ppln.Status.IsReady() {
		return nil, errors.New(PipelineNotReady)
	}
	return ppln, nil
}

// isUpgradePending checks if the component status indicates an upgrade is pending
func isUpgradePending(status v1alpha1.TektonComponentStatus) bool {
	if status == nil {
		return false
	}
	readyCondition := status.GetCondition(apis.ConditionReady)
	if readyCondition == nil {
		return false
	}
	return strings.Contains(readyCondition.Message, v1alpha1.UpgradePending)
}

func PipelineTargetNamspace(informer informer.TektonPipelineInformer) (string, error) {
	ppln, err := getPipelineRes(informer)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}
	return ppln.Spec.TargetNamespace, nil
}

func getPipelineRes(informer informer.TektonPipelineInformer) (*v1alpha1.TektonPipeline, error) {
	res, err := informer.Lister().Get(v1alpha1.PipelineResourceName)
	return res, err
}

func TriggerReady(informer informer.TektonTriggerInformer) (*v1alpha1.TektonTrigger, error) {
	trigger, err := getTriggerRes(informer)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, errors.New(TriggerNotFound)
		}
		return nil, err
	}
	if trigger.GetStatus() != nil && strings.Contains(trigger.GetStatus().GetCondition(apis.ConditionReady).Message, v1alpha1.UpgradePending) {
		return nil, v1alpha1.DEPENDENCY_UPGRADE_PENDING_ERR
	}
	if !trigger.Status.IsReady() {
		return nil, errors.New(TriggerNotReady)
	}
	return trigger, nil
}

func getTriggerRes(informer informer.TektonTriggerInformer) (*v1alpha1.TektonTrigger, error) {
	res, err := informer.Lister().Get(v1alpha1.TriggerResourceName)
	return res, err
}
