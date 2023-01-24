---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "coralogix_tco_policy Data Source - terraform-provider-coralogix"
subcategory: ""
description: "Coralogix TCO-Policy. Api-key is required for this resource. For more information

- https://coralogix.com/docs/tco-optimizer-api ."

---

# coralogix_tco_policy (Data Source)

Coralogix TCO-Policy. Api-key is required for this resource.
For more information - https://coralogix.com/docs/tco-optimizer-api .

## Example Usage

```hcl
data "coralogix_tco_policy" "imported_coralogix_tco_policy" {
  id = coralogix_tco_policy.tco_policy_example.id
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Read-Only

- `application_name` (List of Object) The applications to apply the policy on. Applies the policy on all the
  applications by default. (see [below for nested schema](#nestedatt--application_name))
- `enabled` (Boolean) Determines weather the policy will be enabled. True by default.
- `id` (String) The ID of this resource.
- `name` (String) The policy name.
- `order` (Number) Determines the policy's order between the other policies. By default, will be added last.
- `priority` (String) The policy description. Can be one of ["high" "medium" "low" "block"].
- `severities` (Set of String) The severities to apply the policy on. Can be few
  of ["warning" "error" "critical" "debug" "verbose" "info"].
- `subsystem_name` (List of Object) The subsystems to apply the policy on. Applies the policy on all the subsystems by
  default. (see [below for nested schema](#nestedatt--subsystem_name))

<a id="nestedatt--application_name"></a>

### Nested Schema for `application_name`

Read-Only:

- `includes` (Boolean)
- `is` (Boolean)
- `is_not` (Boolean)
- `rule` (String)
- `rules` (Set of String)
- `starts_with` (Boolean)

<a id="nestedatt--subsystem_name"></a>

### Nested Schema for `subsystem_name`

Read-Only:

- `includes` (Boolean)
- `is` (Boolean)
- `is_not` (Boolean)
- `rule` (String)
- `rules` (Set of String)
- `starts_with` (Boolean)

