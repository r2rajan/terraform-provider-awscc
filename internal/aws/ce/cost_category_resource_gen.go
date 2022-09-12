// Code generated by generators/resource/main.go; DO NOT EDIT.

package ce

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
	registry.AddResourceTypeFactory("awscc_ce_cost_category", costCategoryResourceType)
}

// costCategoryResourceType returns the Terraform awscc_ce_cost_category resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CE::CostCategory resource type.
func costCategoryResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Cost category ARN",
			//   "pattern": "^arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$",
			//   "type": "string"
			// }
			Description: "Cost category ARN",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"default_value": {
			// Property: DefaultValue
			// CloudFormation resource type schema:
			// {
			//   "description": "The default value for the cost category",
			//   "maxLength": 50,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The default value for the cost category",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"effective_start": {
			// Property: EffectiveStart
			// CloudFormation resource type schema:
			// {
			//   "description": "ISO 8601 date time with offset format",
			//   "maxLength": 25,
			//   "minLength": 20,
			//   "pattern": "^\\d{4}-\\d\\d-\\d\\dT\\d\\d:\\d\\d:\\d\\d(([+-]\\d\\d:\\d\\d)|Z)$",
			//   "type": "string"
			// }
			Description: "ISO 8601 date time with offset format",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 50,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"rule_version": {
			// Property: RuleVersion
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "CostCategoryExpression.v1"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"CostCategoryExpression.v1",
				}),
			},
		},
		"rules": {
			// Property: Rules
			// CloudFormation resource type schema:
			// {
			//   "description": "JSON array format of Expression in Billing and Cost Management API",
			//   "type": "string"
			// }
			Description: "JSON array format of Expression in Billing and Cost Management API",
			Type:        types.StringType,
			Required:    true,
		},
		"split_charge_rules": {
			// Property: SplitChargeRules
			// CloudFormation resource type schema:
			// {
			//   "description": "Json array format of CostCategorySplitChargeRule in Billing and Cost Management API",
			//   "type": "string"
			// }
			Description: "Json array format of CostCategorySplitChargeRule in Billing and Cost Management API",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Cost Category enables you to map your cost and usage into meaningful categories. You can use Cost Category to organize your costs using a rule-based engine.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CE::CostCategory").WithTerraformTypeName("awscc_ce_cost_category")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"default_value":      "DefaultValue",
		"effective_start":    "EffectiveStart",
		"name":               "Name",
		"rule_version":       "RuleVersion",
		"rules":              "Rules",
		"split_charge_rules": "SplitChargeRules",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
