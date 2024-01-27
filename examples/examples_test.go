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

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	checkCreds(t)
	return integration.ProgramTestOptions{
		RunUpdateTest:        false,
		ExpectRefreshChanges: true,
	}
}

func checkCreds(t *testing.T) {
	_, exists := os.LookupEnv("DBT_CLOUD_HOST_URL")
	if !exists {
		t.Skipf("Skipping test due to missing DBT_CLOUD_HOST_URL environment variable")
	}

	_, exists = os.LookupEnv("DBT_CLOUD_ACCOUNT_ID")
	if !exists {
		t.Skipf("Skipping test due to missing DBT_CLOUD_ACCOUNT_ID environment variable")
	}

	_, exists = os.LookupEnv("DBT_CLOUD_TOKEN")
	if !exists {
		t.Skipf("Skipping test due to missing DBT_CLOUD_TOKEN environment variable")
	}
}
