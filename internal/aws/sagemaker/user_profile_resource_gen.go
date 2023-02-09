// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_sagemaker_user_profile", userProfileResource)
}

// userProfileResource returns the Terraform awscc_sagemaker_user_profile resource.
// This Terraform resource corresponds to the CloudFormation AWS::SageMaker::UserProfile resource.
func userProfileResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DomainId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the associated Domain.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"domain_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the associated Domain.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SingleSignOnUserIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A specifier for the type of value specified in SingleSignOnUserValue. Currently, the only supported value is \"UserName\". If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.",
		//	  "pattern": "UserName",
		//	  "type": "string"
		//	}
		"single_sign_on_user_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A specifier for the type of value specified in SingleSignOnUserValue. Currently, the only supported value is \"UserName\". If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("UserName"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SingleSignOnUserValue
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The username of the associated AWS Single Sign-On User for this UserProfile. If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"single_sign_on_user_value": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The username of the associated AWS Single Sign-On User for this UserProfile. If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags to apply to the user profile.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 128,
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
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of tags to apply to the user profile.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// Tags is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: UserProfileArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user profile Amazon Resource Name (ARN).",
		//	  "maxLength": 256,
		//	  "pattern": "arn:aws[a-z\\-]*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:user-profile/.*",
		//	  "type": "string"
		//	}
		"user_profile_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The user profile Amazon Resource Name (ARN).",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UserProfileName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the UserProfile.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"user_profile_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the UserProfile.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UserSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A collection of settings.",
		//	  "properties": {
		//	    "ExecutionRole": {
		//	      "description": "The user profile Amazon Resource Name (ARN).",
		//	      "maxLength": 2048,
		//	      "minLength": 20,
		//	      "pattern": "^arn:aws[a-z\\-]*:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$",
		//	      "type": "string"
		//	    },
		//	    "JupyterServerAppSettings": {
		//	      "additionalProperties": false,
		//	      "description": "The Jupyter server's app settings.",
		//	      "properties": {
		//	        "DefaultResourceSpec": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "InstanceType": {
		//	              "description": "The instance type that the image version runs on.",
		//	              "enum": [
		//	                "system",
		//	                "ml.t3.micro",
		//	                "ml.t3.small",
		//	                "ml.t3.medium",
		//	                "ml.t3.large",
		//	                "ml.t3.xlarge",
		//	                "ml.t3.2xlarge",
		//	                "ml.m5.large",
		//	                "ml.m5.xlarge",
		//	                "ml.m5.2xlarge",
		//	                "ml.m5.4xlarge",
		//	                "ml.m5.8xlarge",
		//	                "ml.m5.12xlarge",
		//	                "ml.m5.16xlarge",
		//	                "ml.m5.24xlarge",
		//	                "ml.c5.large",
		//	                "ml.c5.xlarge",
		//	                "ml.c5.2xlarge",
		//	                "ml.c5.4xlarge",
		//	                "ml.c5.9xlarge",
		//	                "ml.c5.12xlarge",
		//	                "ml.c5.18xlarge",
		//	                "ml.c5.24xlarge",
		//	                "ml.p3.2xlarge",
		//	                "ml.p3.8xlarge",
		//	                "ml.p3.16xlarge",
		//	                "ml.g4dn.xlarge",
		//	                "ml.g4dn.2xlarge",
		//	                "ml.g4dn.4xlarge",
		//	                "ml.g4dn.8xlarge",
		//	                "ml.g4dn.12xlarge",
		//	                "ml.g4dn.16xlarge",
		//	                "ml.r5.large",
		//	                "ml.r5.xlarge",
		//	                "ml.r5.2xlarge",
		//	                "ml.r5.4xlarge",
		//	                "ml.r5.8xlarge",
		//	                "ml.r5.12xlarge",
		//	                "ml.r5.16xlarge",
		//	                "ml.r5.24xlarge",
		//	                "ml.p3dn.24xlarge",
		//	                "ml.m5d.large",
		//	                "ml.m5d.xlarge",
		//	                "ml.m5d.2xlarge",
		//	                "ml.m5d.4xlarge",
		//	                "ml.m5d.8xlarge",
		//	                "ml.m5d.12xlarge",
		//	                "ml.m5d.16xlarge",
		//	                "ml.m5d.24xlarge",
		//	                "ml.g5.xlarge",
		//	                "ml.g5.2xlarge",
		//	                "ml.g5.4xlarge",
		//	                "ml.g5.8xlarge",
		//	                "ml.g5.12xlarge",
		//	                "ml.g5.16xlarge",
		//	                "ml.g5.24xlarge",
		//	                "ml.g5.48xlarge"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "SageMakerImageArn": {
		//	              "description": "The ARN of the SageMaker image that the image version belongs to.",
		//	              "maxLength": 256,
		//	              "pattern": "^arn:aws(-[\\w]+)*:sagemaker:.+:[0-9]{12}:image/[a-z0-9]([-.]?[a-z0-9])*$",
		//	              "type": "string"
		//	            },
		//	            "SageMakerImageVersionArn": {
		//	              "description": "The ARN of the image version created on the instance.",
		//	              "maxLength": 256,
		//	              "pattern": "^arn:aws(-[\\w]+)*:sagemaker:.+:[0-9]{12}:image-version/[a-z0-9]([-.]?[a-z0-9])*/[0-9]+$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "KernelGatewayAppSettings": {
		//	      "additionalProperties": false,
		//	      "description": "The kernel gateway app settings.",
		//	      "properties": {
		//	        "CustomImages": {
		//	          "description": "A list of custom SageMaker images that are configured to run as a KernelGateway app.",
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "description": "A custom SageMaker image.",
		//	            "properties": {
		//	              "AppImageConfigName": {
		//	                "description": "The Name of the AppImageConfig.",
		//	                "maxLength": 63,
		//	                "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}",
		//	                "type": "string"
		//	              },
		//	              "ImageName": {
		//	                "description": "The name of the CustomImage. Must be unique to your account.",
		//	                "maxLength": 63,
		//	                "pattern": "^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$",
		//	                "type": "string"
		//	              },
		//	              "ImageVersionNumber": {
		//	                "description": "The version number of the CustomImage.",
		//	                "minimum": 0,
		//	                "type": "integer"
		//	              }
		//	            },
		//	            "required": [
		//	              "AppImageConfigName",
		//	              "ImageName"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "maxItems": 30,
		//	          "minItems": 0,
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        },
		//	        "DefaultResourceSpec": {
		//	          "additionalProperties": false,
		//	          "description": "The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app.",
		//	          "properties": {
		//	            "InstanceType": {
		//	              "description": "The instance type that the image version runs on.",
		//	              "enum": [
		//	                "system",
		//	                "ml.t3.micro",
		//	                "ml.t3.small",
		//	                "ml.t3.medium",
		//	                "ml.t3.large",
		//	                "ml.t3.xlarge",
		//	                "ml.t3.2xlarge",
		//	                "ml.m5.large",
		//	                "ml.m5.xlarge",
		//	                "ml.m5.2xlarge",
		//	                "ml.m5.4xlarge",
		//	                "ml.m5.8xlarge",
		//	                "ml.m5.12xlarge",
		//	                "ml.m5.16xlarge",
		//	                "ml.m5.24xlarge",
		//	                "ml.c5.large",
		//	                "ml.c5.xlarge",
		//	                "ml.c5.2xlarge",
		//	                "ml.c5.4xlarge",
		//	                "ml.c5.9xlarge",
		//	                "ml.c5.12xlarge",
		//	                "ml.c5.18xlarge",
		//	                "ml.c5.24xlarge",
		//	                "ml.p3.2xlarge",
		//	                "ml.p3.8xlarge",
		//	                "ml.p3.16xlarge",
		//	                "ml.g4dn.xlarge",
		//	                "ml.g4dn.2xlarge",
		//	                "ml.g4dn.4xlarge",
		//	                "ml.g4dn.8xlarge",
		//	                "ml.g4dn.12xlarge",
		//	                "ml.g4dn.16xlarge",
		//	                "ml.r5.large",
		//	                "ml.r5.xlarge",
		//	                "ml.r5.2xlarge",
		//	                "ml.r5.4xlarge",
		//	                "ml.r5.8xlarge",
		//	                "ml.r5.12xlarge",
		//	                "ml.r5.16xlarge",
		//	                "ml.r5.24xlarge",
		//	                "ml.p3dn.24xlarge",
		//	                "ml.m5d.large",
		//	                "ml.m5d.xlarge",
		//	                "ml.m5d.2xlarge",
		//	                "ml.m5d.4xlarge",
		//	                "ml.m5d.8xlarge",
		//	                "ml.m5d.12xlarge",
		//	                "ml.m5d.16xlarge",
		//	                "ml.m5d.24xlarge",
		//	                "ml.g5.xlarge",
		//	                "ml.g5.2xlarge",
		//	                "ml.g5.4xlarge",
		//	                "ml.g5.8xlarge",
		//	                "ml.g5.12xlarge",
		//	                "ml.g5.16xlarge",
		//	                "ml.g5.24xlarge",
		//	                "ml.g5.48xlarge"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "SageMakerImageArn": {
		//	              "description": "The ARN of the SageMaker image that the image version belongs to.",
		//	              "maxLength": 256,
		//	              "pattern": "^arn:aws(-[\\w]+)*:sagemaker:.+:[0-9]{12}:image/[a-z0-9]([-.]?[a-z0-9])*$",
		//	              "type": "string"
		//	            },
		//	            "SageMakerImageVersionArn": {
		//	              "description": "The ARN of the image version created on the instance.",
		//	              "maxLength": 256,
		//	              "pattern": "^arn:aws(-[\\w]+)*:sagemaker:.+:[0-9]{12}:image-version/[a-z0-9]([-.]?[a-z0-9])*/[0-9]+$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "RStudioServerProAppSettings": {
		//	      "additionalProperties": false,
		//	      "description": "A collection of settings that configure user interaction with the RStudioServerPro app.",
		//	      "properties": {
		//	        "AccessStatus": {
		//	          "description": "Indicates whether the current user has access to the RStudioServerPro app.",
		//	          "enum": [
		//	            "ENABLED",
		//	            "DISABLED"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "UserGroup": {
		//	          "description": "The level of permissions that the user has within the RStudioServerPro app. This value defaults to User. The Admin value allows the user access to the RStudio Administrative Dashboard.",
		//	          "enum": [
		//	            "R_STUDIO_ADMIN",
		//	            "R_STUDIO_USER"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "SecurityGroups": {
		//	      "description": "The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
		//	      "items": {
		//	        "maxLength": 32,
		//	        "pattern": "[-0-9a-zA-Z]+",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 5,
		//	      "minItems": 0,
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "SharingSettings": {
		//	      "additionalProperties": false,
		//	      "description": "The sharing settings.",
		//	      "properties": {
		//	        "NotebookOutputOption": {
		//	          "description": "Whether to include the notebook cell output when sharing the notebook. The default is Disabled.",
		//	          "enum": [
		//	            "Allowed",
		//	            "Disabled"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "S3KmsKeyId": {
		//	          "description": "When NotebookOutputOption is Allowed, the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.",
		//	          "maxLength": 2048,
		//	          "pattern": ".*",
		//	          "type": "string"
		//	        },
		//	        "S3OutputPath": {
		//	          "description": "When NotebookOutputOption is Allowed, the Amazon S3 bucket used to store the shared notebook snapshots.",
		//	          "maxLength": 1024,
		//	          "pattern": "^(https|s3)://([^/]+)/?(.*)$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"user_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ExecutionRole
				"execution_role": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The user profile Amazon Resource Name (ARN).",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(20, 2048),
						stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[a-z\\-]*:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: JupyterServerAppSettings
				"jupyter_server_app_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DefaultResourceSpec
						"default_resource_spec": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: InstanceType
								"instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The instance type that the image version runs on.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.OneOf(
											"system",
											"ml.t3.micro",
											"ml.t3.small",
											"ml.t3.medium",
											"ml.t3.large",
											"ml.t3.xlarge",
											"ml.t3.2xlarge",
											"ml.m5.large",
											"ml.m5.xlarge",
											"ml.m5.2xlarge",
											"ml.m5.4xlarge",
											"ml.m5.8xlarge",
											"ml.m5.12xlarge",
											"ml.m5.16xlarge",
											"ml.m5.24xlarge",
											"ml.c5.large",
											"ml.c5.xlarge",
											"ml.c5.2xlarge",
											"ml.c5.4xlarge",
											"ml.c5.9xlarge",
											"ml.c5.12xlarge",
											"ml.c5.18xlarge",
											"ml.c5.24xlarge",
											"ml.p3.2xlarge",
											"ml.p3.8xlarge",
											"ml.p3.16xlarge",
											"ml.g4dn.xlarge",
											"ml.g4dn.2xlarge",
											"ml.g4dn.4xlarge",
											"ml.g4dn.8xlarge",
											"ml.g4dn.12xlarge",
											"ml.g4dn.16xlarge",
											"ml.r5.large",
											"ml.r5.xlarge",
											"ml.r5.2xlarge",
											"ml.r5.4xlarge",
											"ml.r5.8xlarge",
											"ml.r5.12xlarge",
											"ml.r5.16xlarge",
											"ml.r5.24xlarge",
											"ml.p3dn.24xlarge",
											"ml.m5d.large",
											"ml.m5d.xlarge",
											"ml.m5d.2xlarge",
											"ml.m5d.4xlarge",
											"ml.m5d.8xlarge",
											"ml.m5d.12xlarge",
											"ml.m5d.16xlarge",
											"ml.m5d.24xlarge",
											"ml.g5.xlarge",
											"ml.g5.2xlarge",
											"ml.g5.4xlarge",
											"ml.g5.8xlarge",
											"ml.g5.12xlarge",
											"ml.g5.16xlarge",
											"ml.g5.24xlarge",
											"ml.g5.48xlarge",
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: SageMakerImageArn
								"sage_maker_image_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The ARN of the SageMaker image that the image version belongs to.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthAtMost(256),
										stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws(-[\\w]+)*:sagemaker:.+:[0-9]{12}:image/[a-z0-9]([-.]?[a-z0-9])*$"), ""),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: SageMakerImageVersionArn
								"sage_maker_image_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The ARN of the image version created on the instance.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthAtMost(256),
										stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws(-[\\w]+)*:sagemaker:.+:[0-9]{12}:image-version/[a-z0-9]([-.]?[a-z0-9])*/[0-9]+$"), ""),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The Jupyter server's app settings.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: KernelGatewayAppSettings
				"kernel_gateway_app_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CustomImages
						"custom_images": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: AppImageConfigName
									"app_image_config_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The Name of the AppImageConfig.",
										Required:    true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthAtMost(63),
											stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}"), ""),
										}, /*END VALIDATORS*/
									}, /*END ATTRIBUTE*/
									// Property: ImageName
									"image_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The name of the CustomImage. Must be unique to your account.",
										Required:    true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthAtMost(63),
											stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]([-.]?[a-zA-Z0-9]){0,62}$"), ""),
										}, /*END VALIDATORS*/
									}, /*END ATTRIBUTE*/
									// Property: ImageVersionNumber
									"image_version_number": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Description: "The version number of the CustomImage.",
										Optional:    true,
										Computed:    true,
										Validators: []validator.Int64{ /*START VALIDATORS*/
											int64validator.AtLeast(0),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
											int64planmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Description: "A list of custom SageMaker images that are configured to run as a KernelGateway app.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(0, 30),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: DefaultResourceSpec
						"default_resource_spec": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: InstanceType
								"instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The instance type that the image version runs on.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.OneOf(
											"system",
											"ml.t3.micro",
											"ml.t3.small",
											"ml.t3.medium",
											"ml.t3.large",
											"ml.t3.xlarge",
											"ml.t3.2xlarge",
											"ml.m5.large",
											"ml.m5.xlarge",
											"ml.m5.2xlarge",
											"ml.m5.4xlarge",
											"ml.m5.8xlarge",
											"ml.m5.12xlarge",
											"ml.m5.16xlarge",
											"ml.m5.24xlarge",
											"ml.c5.large",
											"ml.c5.xlarge",
											"ml.c5.2xlarge",
											"ml.c5.4xlarge",
											"ml.c5.9xlarge",
											"ml.c5.12xlarge",
											"ml.c5.18xlarge",
											"ml.c5.24xlarge",
											"ml.p3.2xlarge",
											"ml.p3.8xlarge",
											"ml.p3.16xlarge",
											"ml.g4dn.xlarge",
											"ml.g4dn.2xlarge",
											"ml.g4dn.4xlarge",
											"ml.g4dn.8xlarge",
											"ml.g4dn.12xlarge",
											"ml.g4dn.16xlarge",
											"ml.r5.large",
											"ml.r5.xlarge",
											"ml.r5.2xlarge",
											"ml.r5.4xlarge",
											"ml.r5.8xlarge",
											"ml.r5.12xlarge",
											"ml.r5.16xlarge",
											"ml.r5.24xlarge",
											"ml.p3dn.24xlarge",
											"ml.m5d.large",
											"ml.m5d.xlarge",
											"ml.m5d.2xlarge",
											"ml.m5d.4xlarge",
											"ml.m5d.8xlarge",
											"ml.m5d.12xlarge",
											"ml.m5d.16xlarge",
											"ml.m5d.24xlarge",
											"ml.g5.xlarge",
											"ml.g5.2xlarge",
											"ml.g5.4xlarge",
											"ml.g5.8xlarge",
											"ml.g5.12xlarge",
											"ml.g5.16xlarge",
											"ml.g5.24xlarge",
											"ml.g5.48xlarge",
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: SageMakerImageArn
								"sage_maker_image_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The ARN of the SageMaker image that the image version belongs to.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthAtMost(256),
										stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws(-[\\w]+)*:sagemaker:.+:[0-9]{12}:image/[a-z0-9]([-.]?[a-z0-9])*$"), ""),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: SageMakerImageVersionArn
								"sage_maker_image_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The ARN of the image version created on the instance.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthAtMost(256),
										stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws(-[\\w]+)*:sagemaker:.+:[0-9]{12}:image-version/[a-z0-9]([-.]?[a-z0-9])*/[0-9]+$"), ""),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The kernel gateway app settings.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RStudioServerProAppSettings
				"r_studio_server_pro_app_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AccessStatus
						"access_status": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Indicates whether the current user has access to the RStudioServerPro app.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"ENABLED",
									"DISABLED",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
								stringplanmodifier.RequiresReplace(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: UserGroup
						"user_group": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The level of permissions that the user has within the RStudioServerPro app. This value defaults to User. The Admin value allows the user access to the RStudio Administrative Dashboard.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"R_STUDIO_ADMIN",
									"R_STUDIO_USER",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
								stringplanmodifier.RequiresReplace(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "A collection of settings that configure user interaction with the RStudioServerPro app.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SecurityGroups
				"security_groups": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The security groups for the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(0, 5),
						listvalidator.ValueStringsAre(
							stringvalidator.LengthAtMost(32),
							stringvalidator.RegexMatches(regexp.MustCompile("[-0-9a-zA-Z]+"), ""),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SharingSettings
				"sharing_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: NotebookOutputOption
						"notebook_output_option": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Whether to include the notebook cell output when sharing the notebook. The default is Disabled.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"Allowed",
									"Disabled",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: S3KmsKeyId
						"s3_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "When NotebookOutputOption is Allowed, the AWS Key Management Service (KMS) encryption key ID used to encrypt the notebook cell output in the Amazon S3 bucket.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(2048),
								stringvalidator.RegexMatches(regexp.MustCompile(".*"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: S3OutputPath
						"s3_output_path": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "When NotebookOutputOption is Allowed, the Amazon S3 bucket used to store the shared notebook snapshots.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(1024),
								stringvalidator.RegexMatches(regexp.MustCompile("^(https|s3)://([^/]+)/?(.*)$"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The sharing settings.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A collection of settings.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::SageMaker::UserProfile",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::UserProfile").WithTerraformTypeName("awscc_sagemaker_user_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_status":                    "AccessStatus",
		"app_image_config_name":            "AppImageConfigName",
		"custom_images":                    "CustomImages",
		"default_resource_spec":            "DefaultResourceSpec",
		"domain_id":                        "DomainId",
		"execution_role":                   "ExecutionRole",
		"image_name":                       "ImageName",
		"image_version_number":             "ImageVersionNumber",
		"instance_type":                    "InstanceType",
		"jupyter_server_app_settings":      "JupyterServerAppSettings",
		"kernel_gateway_app_settings":      "KernelGatewayAppSettings",
		"key":                              "Key",
		"notebook_output_option":           "NotebookOutputOption",
		"r_studio_server_pro_app_settings": "RStudioServerProAppSettings",
		"s3_kms_key_id":                    "S3KmsKeyId",
		"s3_output_path":                   "S3OutputPath",
		"sage_maker_image_arn":             "SageMakerImageArn",
		"sage_maker_image_version_arn":     "SageMakerImageVersionArn",
		"security_groups":                  "SecurityGroups",
		"sharing_settings":                 "SharingSettings",
		"single_sign_on_user_identifier":   "SingleSignOnUserIdentifier",
		"single_sign_on_user_value":        "SingleSignOnUserValue",
		"tags":                             "Tags",
		"user_group":                       "UserGroup",
		"user_profile_arn":                 "UserProfileArn",
		"user_profile_name":                "UserProfileName",
		"user_settings":                    "UserSettings",
		"value":                            "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
