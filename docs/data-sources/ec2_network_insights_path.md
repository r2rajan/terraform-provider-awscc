---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_network_insights_path Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::NetworkInsightsPath
---

# awscc_ec2_network_insights_path (Data Source)

Data Source schema for AWS::EC2::NetworkInsightsPath



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `created_date` (String)
- `destination` (String)
- `destination_arn` (String)
- `destination_ip` (String)
- `destination_port` (Number)
- `filter_at_destination` (Attributes) (see [below for nested schema](#nestedatt--filter_at_destination))
- `filter_at_source` (Attributes) (see [below for nested schema](#nestedatt--filter_at_source))
- `network_insights_path_arn` (String)
- `network_insights_path_id` (String)
- `protocol` (String)
- `source` (String)
- `source_arn` (String)
- `source_ip` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--filter_at_destination"></a>
### Nested Schema for `filter_at_destination`

Read-Only:

- `destination_address` (String)
- `destination_port_range` (Attributes) (see [below for nested schema](#nestedatt--filter_at_destination--destination_port_range))
- `source_address` (String)
- `source_port_range` (Attributes) (see [below for nested schema](#nestedatt--filter_at_destination--source_port_range))

<a id="nestedatt--filter_at_destination--destination_port_range"></a>
### Nested Schema for `filter_at_destination.destination_port_range`

Read-Only:

- `from_port` (Number)
- `to_port` (Number)


<a id="nestedatt--filter_at_destination--source_port_range"></a>
### Nested Schema for `filter_at_destination.source_port_range`

Read-Only:

- `from_port` (Number)
- `to_port` (Number)



<a id="nestedatt--filter_at_source"></a>
### Nested Schema for `filter_at_source`

Read-Only:

- `destination_address` (String)
- `destination_port_range` (Attributes) (see [below for nested schema](#nestedatt--filter_at_source--destination_port_range))
- `source_address` (String)
- `source_port_range` (Attributes) (see [below for nested schema](#nestedatt--filter_at_source--source_port_range))

<a id="nestedatt--filter_at_source--destination_port_range"></a>
### Nested Schema for `filter_at_source.destination_port_range`

Read-Only:

- `from_port` (Number)
- `to_port` (Number)


<a id="nestedatt--filter_at_source--source_port_range"></a>
### Nested Schema for `filter_at_source.source_port_range`

Read-Only:

- `from_port` (Number)
- `to_port` (Number)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)


