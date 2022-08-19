package test

import (
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the simple Terraform module in examples/terraform-basic-example using Terratest.
func TestTerraformBasicExample(t *testing.T) {
	t.Parallel()

	expectedText := "Data Engineering-Shared Cluster"
	exampleFolder := test_structure.CopyTerraformFolderToTemp(t, "../", "../..")
	planFilePath := filepath.Join(exampleFolder, "plan.out")

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1::Set the path to the Terraform code that will be tested.
		// The path to where our Terraform code is located
		TerraformDir: "../../",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"cluster_name": expectedText,
		},

		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"terraform.tfvars"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor:      true,
		PlanFilePath: planFilePath,
	})

	plan := terraform.InitAndPlanAndShowWithStruct(t, terraformOptions)
	terraform.RequirePlannedValuesMapKeyExists(t, plan, "module.azure_databricks_demo.databricks_cluster.team_cluster")

	databricksCluster := plan.ResourcePlannedValuesMap["module.azure_databricks_demo.databricks_cluster.team_cluster"]
	clusterName := databricksCluster.AttributeValues["cluster_name"]

	assert.Equal(t, expectedText, clusterName)

}
