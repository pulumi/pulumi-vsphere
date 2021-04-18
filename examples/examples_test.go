// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{}
}

func skipIfNoVsphereConfig(t *testing.T) {
	_, server := os.LookupEnv("VSPHERE_SERVER")
	if !server {
		t.Skipf("Skipping: VSPHERE_SERVER is not set")
	}
	_, user := os.LookupEnv("VSPHERE_USER")
	if !user {
		t.Skipf("Skipping: VSPHERE_USER is not set")
	}
	_, pass := os.LookupEnv("VSPHERE_PASSWORD")
	if !pass {
		t.Skipf("Skipping: VSPHERE_PASSWORD is not set")
	}
}
