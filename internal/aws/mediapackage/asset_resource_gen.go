// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_mediapackage_asset", assetResourceType)
}

// assetResourceType returns the Terraform awscc_mediapackage_asset resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::MediaPackage::Asset resource type.
func assetResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the Asset.",
			//   "type": "string"
			// }
			Description: "The ARN of the Asset.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The time the Asset was initially submitted for Ingest.",
			//   "type": "string"
			// }
			Description: "The time the Asset was initially submitted for Ingest.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"egress_endpoints": {
			// Property: EgressEndpoints
			// CloudFormation resource type schema:
			// {
			//   "description": "The list of egress endpoints available for the Asset.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The endpoint URL used to access an Asset using one PackagingConfiguration.",
			//     "properties": {
			//       "PackagingConfigurationId": {
			//         "description": "The ID of the PackagingConfiguration being applied to the Asset.",
			//         "type": "string"
			//       },
			//       "Url": {
			//         "description": "The URL of the parent manifest for the repackaged Asset.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "PackagingConfigurationId",
			//       "Url"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The list of egress endpoints available for the Asset.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"packaging_configuration_id": {
						// Property: PackagingConfigurationId
						Description: "The ID of the PackagingConfiguration being applied to the Asset.",
						Type:        types.StringType,
						Required:    true,
					},
					"url": {
						// Property: Url
						Description: "The URL of the parent manifest for the repackaged Asset.",
						Type:        types.StringType,
						Required:    true,
					},
				},
			),
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique identifier for the Asset.",
			//   "type": "string"
			// }
			Description: "The unique identifier for the Asset.",
			Type:        types.StringType,
			Required:    true,
		},
		"packaging_group_id": {
			// Property: PackagingGroupId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the PackagingGroup for the Asset.",
			//   "type": "string"
			// }
			Description: "The ID of the PackagingGroup for the Asset.",
			Type:        types.StringType,
			Required:    true,
		},
		"resource_id": {
			// Property: ResourceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The resource ID to include in SPEKE key requests.",
			//   "type": "string"
			// }
			Description: "The resource ID to include in SPEKE key requests.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"source_arn": {
			// Property: SourceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "ARN of the source object in S3.",
			//   "type": "string"
			// }
			Description: "ARN of the source object in S3.",
			Type:        types.StringType,
			Required:    true,
		},
		"source_role_arn": {
			// Property: SourceRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The IAM role_arn used to access the source S3 bucket.",
			//   "type": "string"
			// }
			Description: "The IAM role_arn used to access the source S3 bucket.",
			Type:        types.StringType,
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A collection of tags associated with a resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::MediaPackage::Asset",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaPackage::Asset").WithTerraformTypeName("awscc_mediapackage_asset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                        "Arn",
		"created_at":                 "CreatedAt",
		"egress_endpoints":           "EgressEndpoints",
		"id":                         "Id",
		"key":                        "Key",
		"packaging_configuration_id": "PackagingConfigurationId",
		"packaging_group_id":         "PackagingGroupId",
		"resource_id":                "ResourceId",
		"source_arn":                 "SourceArn",
		"source_role_arn":            "SourceRoleArn",
		"tags":                       "Tags",
		"url":                        "Url",
		"value":                      "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
