package scim

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/require"
)

func TestDataGroupsReadByDisplayName(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups?attributes=id%2CdisplayName%2CexternalId&filter=displayName%20co%20%22test%22",
				Response: GroupList{
					Resources: []Group{
						{
							ID:          "abc1",
							DisplayName: "test-group-1",
							ExternalID:  "ext1",
						},
						{
							ID:          "abc2",
							DisplayName: "test-group-2",
							ExternalID:  "ext2",
						},
					},
				},
			},
		},
		Resource:    DataSourceGroups(),
		HCL:         `display_name_contains = "test"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"display_names": []string{"test-group-1", "test-group-2"},
	})
}

func TestDataGroupsReadNotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups?attributes=id%2CdisplayName%2CexternalId&filter=displayName%20co%20%22notfound%22",
				Response: GroupList{},
			},
		},
		Resource:    DataSourceGroups(),
		HCL:         `display_name_contains = "notfound"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"display_names":          []string{},
		"groups":                 []any{},
		"display_name_contains":  "notfound",
	})
}

func TestDataGroupsReadNoFilter(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups?attributes=id%2CdisplayName%2CexternalId",
				Response: GroupList{
					Resources: []Group{
						{
							ID:          "abc2",
							DisplayName: "group-b",
						},
						{
							ID:          "abc1",
							DisplayName: "group-a",
						},
					},
				},
			},
		},
		Resource:    DataSourceGroups(),
		HCL:         ``,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"display_names": []string{"group-a", "group-b"},
	})
}

func TestDataGroupsReadError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/preview/scim/v2/Groups?attributes=id%2CdisplayName%2CexternalId&filter=displayName%20co%20%22test%22",
				Status:   500,
			},
		},
		Resource:    DataSourceGroups(),
		HCL:         `display_name_contains = "test"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.Apply(t)
	require.Error(t, err)
}
