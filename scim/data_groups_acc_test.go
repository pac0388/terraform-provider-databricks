package scim_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

const groups = `
resource "databricks_group" "this" {
	display_name = "TF Group {var.RANDOM}"
}

data databricks_groups "this" {
	display_name_contains = ""
	depends_on = [databricks_group.this]
}`

func TestAccDataSourceGroupsOnAWS(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "TEST_EC2_INSTANCE_PROFILE")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: groups,
	})
}

func TestAccDataSourceGroupsOnGCP(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "GOOGLE_CREDENTIALS")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: groups,
	})
}

func TestAccDataSourceGroupsOnAzure(t *testing.T) {
	acceptance.GetEnvOrSkipTest(t, "ARM_CLIENT_ID")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: groups,
	})
}
