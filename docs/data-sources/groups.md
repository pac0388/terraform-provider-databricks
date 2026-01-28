---
subcategory: "Security"
---

# databricks_groups Data Source

Retrieves `display_names` of all [databricks_group](../resources/group.md) based on their `display_name`

-> This data source can be used with an account or workspace-level provider.

## Example Usage

Adding all groups of which display name contains `analytics` to a parent group

```hcl
data "databricks_group" "parent" {
  display_name = "parent-group"
}

data "databricks_groups" "analytics" {
  display_name_contains = "analytics"
}

resource "databricks_group_member" "my_member_group" {
  for_each  = toset(data.databricks_groups.analytics.display_names)
  group_id  = data.databricks_group.parent.id
  member_id = data.databricks_group.analytics_by_name[each.value].id
}

data "databricks_group" "analytics_by_name" {
  for_each     = toset(data.databricks_groups.analytics.display_names)
  display_name = each.value
}
```

## Argument Reference

Data source allows you to pick groups by the following attributes

- `display_name_contains` - (Optional) Only return [databricks_group](group.md) display name that match the given name string

## Attribute Reference

Data source exposes the following attributes:

- `display_names` - List of `display_names` of groups. Individual group can be retrieved using [databricks_group](group.md) data source or from `groups` attribute.
- `groups` - List of objects describing individual groups. Each object has the following attributes:
  - `id` - The id of the group (SCIM ID).
  - `display_name` - Display name of the [group](../resources/group.md), e.g. `Analytics Team`.
  - `external_id` - ID of the group in an external identity provider.
  - `acl_principal_id` - identifier for use in [databricks_access_control_rule_set](../resources/access_control_rule_set.md), e.g. `groups/analytics-team`.

## Related Resources

The following resources are used in the same context:

- [End to end workspace management](../guides/workspace-management.md) guide.
- [databricks_current_user](current_user.md) data to retrieve information about [databricks_user](../resources/user.md) or [databricks_service_principal](../resources/service_principal.md), that is calling Databricks REST API.
- [databricks_group](../resources/group.md) to manage [Account-level](https://docs.databricks.com/aws/en/admin/users-groups/groups) or [Workspace-level](https://docs.databricks.com/aws/en/admin/users-groups/workspace-local-groups) groups.
- [databricks_group](group.md) data to retrieve information about [databricks_group](../resources/group.md) members, entitlements and instance profiles.
- [databricks_group_instance_profile](../resources/group_instance_profile.md) to attach [databricks_instance_profile](../resources/instance_profile.md) (AWS) to [databricks_group](../resources/group.md).
- [databricks_group_member](../resources/group_member.md) to attach [users](../resources/user.md) and [groups](../resources/group.md) as group members.
- [databricks_permissions](../resources/permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
- [databricks_user](../resources/user.md) to manage [users](https://docs.databricks.com/administration-guide/users-groups/users.html) in Databricks workspace.
