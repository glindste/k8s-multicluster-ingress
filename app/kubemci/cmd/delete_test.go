// Copyright 2017 Google Inc.
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

package cmd

import (
	"testing"
)

func TestValidateDeleteArgs(t *testing.T) {
	// validateDeleteArgs should return an error with empty options.
	options := DeleteOptions{}
	if err := validateDeleteArgs(&options, []string{}); err == nil {
		t.Errorf("Expected error for emtpy options")
	}

	// validateDeleteArgs should return an error with missing load balancer name.
	options = DeleteOptions{
		IngressFilename:    "ingress.yaml",
		GCPProject:         "gcp-project",
		KubeconfigFilename: "kubeconfig",
	}
	if err := validateDeleteArgs(&options, []string{}); err == nil {
		t.Errorf("Expected error for missing load balancer name")
	}

	// validateDeleteArgs should return an error with missing ingress.
	options = DeleteOptions{
		GCPProject:         "gcp-project",
		KubeconfigFilename: "kubeconfig",
	}
	if err := validateDeleteArgs(&options, []string{"lbname"}); err == nil {
		t.Errorf("Expected error for missing ingress")
	}

	// validateDeleteArgs should return an error with missing gcp project.
	options = DeleteOptions{
		IngressFilename:    "ingress.yaml",
		KubeconfigFilename: "kubeconfig",
	}
	if err := validateDeleteArgs(&options, []string{"lbname"}); err == nil {
		t.Errorf("Expected error for missing gcp project")
	}

	// validateDeleteArgs should return an error with missing kubeconfig.
	options = DeleteOptions{
		IngressFilename: "ingress.yaml",
		GCPProject:      "gcp-project",
	}
	if err := validateDeleteArgs(&options, []string{"lbname"}); err == nil {
		t.Errorf("Expected error for missing kubeconfig")
	}

	// validateDeleteArgs should succeed when all arguments are passed as expected.
	options = DeleteOptions{
		IngressFilename:    "ingress.yaml",
		GCPProject:         "gcp-project",
		KubeconfigFilename: "kubeconfig",
	}
	if err := validateDeleteArgs(&options, []string{"lbname"}); err != nil {
		t.Errorf("unexpected error from validateDeleteArgs: %s", err)
	}
}
