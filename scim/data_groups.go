package scim

import (
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/databricks/terraform-provider-databricks/common"
)

// DataSourceGroups searches for groups based on display_name
func DataSourceGroups() common.Resource {
	type groupData struct {
		ID             string `json:"id,omitempty" tf:"computed"`
		DisplayName    string `json:"display_name,omitempty" tf:"computed"`
		ExternalID     string `json:"external_id,omitempty" tf:"computed"`
		AclPrincipalID string `json:"acl_principal_id,omitempty" tf:"computed"`
	}
	type groupsData struct {
		DisplayNameContains string      `json:"display_name_contains,omitempty" tf:"computed"`
		DisplayNames        []string    `json:"display_names,omitempty" tf:"computed,slice_set"`
		Groups              []groupData `json:"groups,omitempty" tf:"computed"`
	}
	return common.DataResource(groupsData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		response := e.(*groupsData)
		groupsAPI := NewGroupsAPI(ctx, c)

		var filter string

		if response.DisplayNameContains != "" {
			filter = fmt.Sprintf(`displayName co "%s"`, response.DisplayNameContains)
		}
		groupList, err := groupsAPI.Filter(filter, "id,displayName,externalId")
		if err != nil {
			return err
		}
		if len(groupList.Resources) == 0 {
			log.Printf("[INFO] cannot find groups with display name containing %s", response.DisplayNameContains)
		}
		for _, group := range groupList.Resources {
			response.DisplayNames = append(response.DisplayNames, group.DisplayName)
			response.Groups = append(response.Groups, groupData{
				ID:             group.ID,
				DisplayName:    group.DisplayName,
				ExternalID:     group.ExternalID,
				AclPrincipalID: fmt.Sprintf("groups/%s", group.DisplayName),
			})
		}
		sort.Strings(response.DisplayNames)
		sort.Slice(response.Groups, func(i, j int) bool {
			return response.Groups[i].DisplayName < response.Groups[j].DisplayName
		})
		return nil
	})
}
