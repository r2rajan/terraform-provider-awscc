// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package billingconductor

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_billingconductor_billing_group", billingGroupResource)
}

// billingGroupResource returns the Terraform awscc_billingconductor_billing_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::BillingConductor::BillingGroup resource.
func billingGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountGrouping
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "LinkedAccountIds": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "pattern": "[0-9]{12}",
		//	        "type": "string"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "required": [
		//	    "LinkedAccountIds"
		//	  ],
		//	  "type": "object"
		//	}
		"account_grouping": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LinkedAccountIds
				"linked_account_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Required:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeAtLeast(1),
						setvalidator.ValueStringsAre(
							stringvalidator.RegexMatches(regexp.MustCompile("[0-9]{12}"), ""),
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Billing Group ARN",
		//	  "pattern": "arn:aws(-cn)?:billingconductor::[0-9]{12}:billinggroup/?[0-9]{12}",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Billing Group ARN",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ComputationPreference
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "PricingPlanArn": {
		//	      "description": "ARN of the attached pricing plan",
		//	      "pattern": "arn:aws(-cn)?:billingconductor::[0-9]{12}:pricingplan/[a-zA-Z0-9]{10}",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "PricingPlanArn"
		//	  ],
		//	  "type": "object"
		//	}
		"computation_preference": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PricingPlanArn
				"pricing_plan_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "ARN of the attached pricing plan",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("arn:aws(-cn)?:billingconductor::[0-9]{12}:pricingplan/[a-zA-Z0-9]{10}"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Creation timestamp in UNIX epoch time format",
		//	  "type": "integer"
		//	}
		"creation_time": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Creation timestamp in UNIX epoch time format",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1024),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Latest modified timestamp in UNIX epoch time format",
		//	  "type": "integer"
		//	}
		"last_modified_time": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Latest modified timestamp in UNIX epoch time format",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_\\+=\\.\\-@]+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9_\\+=\\.\\-@]+"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: PrimaryAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This account will act as a virtual payer account of the billing group",
		//	  "pattern": "[0-9]{12}",
		//	  "type": "string"
		//	}
		"primary_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This account will act as a virtual payer account of the billing group",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[0-9]{12}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Size
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Number of accounts in the billing group",
		//	  "type": "integer"
		//	}
		"size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Number of accounts in the billing group",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "ACTIVE",
		//	    "PRIMARY_ACCOUNT_MISSING"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StatusReason
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"status_reason": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "A billing group is a set of linked account which belong to the same end customer. It can be seen as a virtual consolidated billing family.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::BillingConductor::BillingGroup").WithTerraformTypeName("awscc_billingconductor_billing_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_grouping":       "AccountGrouping",
		"arn":                    "Arn",
		"computation_preference": "ComputationPreference",
		"creation_time":          "CreationTime",
		"description":            "Description",
		"key":                    "Key",
		"last_modified_time":     "LastModifiedTime",
		"linked_account_ids":     "LinkedAccountIds",
		"name":                   "Name",
		"pricing_plan_arn":       "PricingPlanArn",
		"primary_account_id":     "PrimaryAccountId",
		"size":                   "Size",
		"status":                 "Status",
		"status_reason":          "StatusReason",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
