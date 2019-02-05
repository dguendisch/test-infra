// Copyright 2019 Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"context"
	"fmt"
	"strings"

	argov1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	tmv1beta1 "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1"
	"github.com/gardener/test-infra/pkg/testmachinery"
	"github.com/gardener/test-infra/pkg/testmachinery/testrun"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func getWorkflowName(tr *tmv1beta1.Testrun) string {
	return fmt.Sprintf("%s-wf", tr.Name)
}

func (r *TestrunReconciler) createWorkflow(ctx context.Context, testrunDef *tmv1beta1.Testrun) (*argov1.Workflow, error) {
	tr, err := testrun.New(testrunDef)
	if err != nil {
		return nil, fmt.Errorf("Error parsing testrun: %s", err.Error())
	}

	wf, err := tr.GetWorkflow(getWorkflowName(testrunDef), testrunDef.Namespace, r.getImagePullSecrets(ctx))
	if err != nil {
		return nil, err
	}

	if err := controllerutil.SetControllerReference(testrunDef, wf, r.scheme); err != nil {
		return nil, err
	}

	wfFinalizers := sets.NewString(wf.Finalizers...)
	if !wfFinalizers.Has(tmv1beta1.SchemeGroupVersion.Group) {
		wfFinalizers.Insert(tmv1beta1.SchemeGroupVersion.Group)
		wf.Finalizers = wfFinalizers.UnsortedList()
	}

	return wf, nil
}

func (r *TestrunReconciler) getImagePullSecrets(ctx context.Context) []string {
	configMap := &corev1.ConfigMap{}
	err := r.Get(ctx, types.NamespacedName{Name: testmachinery.ConfigMapName, Namespace: testmachinery.GetConfig().Namespace}, configMap)
	if err != nil {
		log.Warnf("Cannot fetch Testmachinery config: %s", err.Error())
		return nil
	}

	pullSecretNames := configMap.Data["secrets.PullSecrets"]

	if pullSecretNames == "" {
		return nil
	}

	return strings.Split(pullSecretNames, ",")
}