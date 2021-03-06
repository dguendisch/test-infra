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

package testrun

import (
	argov1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	tmv1beta1 "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1"
	"github.com/gardener/test-infra/pkg/testmachinery/argo"
	"github.com/gardener/test-infra/pkg/testmachinery/config"
	"github.com/gardener/test-infra/pkg/testmachinery/testdefinition"
	"github.com/gardener/test-infra/pkg/testmachinery/testflow"
)

// New takes a testrun crd and creates a new Testrun representation.
// It fetches testruns from specified testdeflocations and generates a testflow object.
func New(tr *tmv1beta1.Testrun) (*Testrun, error) {

	testDefinitions, err := testdefinition.NewTestDefinitions(tr.Spec.TestLocations)
	if err != nil {
		return nil, err
	}

	globalConfig := config.New(tr.Spec.Config)

	// create initial prepare step
	prepare := testdefinition.NewPrepare("Prepare", tr.Spec.Kubeconfigs)
	tf, err := testflow.New(testflow.FlowIDTest, &tr.Spec.TestFlow, testDefinitions, globalConfig, prepare)
	if err != nil {
		return nil, err
	}

	onExitPrepare := testdefinition.NewPrepare("PostPrepare", tr.Spec.Kubeconfigs)
	onExitFlow, err := testflow.New(testflow.FlowIDExit, &tr.Spec.OnExit, testDefinitions, globalConfig, onExitPrepare)
	if err != nil {
		return nil, err
	}

	return &Testrun{Info: tr, Testflow: tf, OnExitTestflow: onExitFlow}, nil
}

// GetWorkflow returns the argo workflow object of this testrun.
func (tr *Testrun) GetWorkflow(name, namespace string, pullImageSecretNames []string) (*argov1.Workflow, error) {
	testrunName := "testrun"
	onExitName := "exit-handler"

	templates, err := tr.Testflow.GetTemplates(testrunName)
	if err != nil {
		return nil, err
	}

	onExitTemplates, err := tr.OnExitTestflow.GetTemplates(onExitName)
	if err != nil {
		return nil, err
	}
	return argo.CreateWorkflow(name, namespace, testrunName, onExitName, append(templates, onExitTemplates...), tr.Testflow.GetLocalVolumes(), tr.Info.Spec.TTLSecondsAfterFinished, pullImageSecretNames)
}
