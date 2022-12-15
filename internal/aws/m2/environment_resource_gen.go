// Code generated by generators/resource/main.go; DO NOT EDIT.

package m2

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_m2_environment", environmentResource)
}

// environmentResource returns the Terraform awscc_m2_environment resource.
// This Terraform resource corresponds to the CloudFormation AWS::M2::Environment resource.
func environmentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The description of the environment.",
			//	  "maxLength": 500,
			//	  "minLength": 0,
			//	  "type": "string"
			//	}
			Description: "The description of the environment.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 500),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"engine_type": {
			// Property: EngineType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The target platform for the environment.",
			//	  "enum": [
			//	    "microfocus",
			//	    "bluage"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "The target platform for the environment.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"microfocus",
					"bluage",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"engine_version": {
			// Property: EngineVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The version of the runtime engine for the environment.",
			//	  "pattern": "^\\S{1,10}$",
			//	  "type": "string"
			//	}
			Description: "The version of the runtime engine for the environment.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^\\S{1,10}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"environment_arn": {
			// Property: EnvironmentArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the runtime environment.",
			//	  "pattern": "",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the runtime environment.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"environment_id": {
			// Property: EnvironmentId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The unique identifier of the environment.",
			//	  "pattern": "^\\S{1,80}$",
			//	  "type": "string"
			//	}
			Description: "The unique identifier of the environment.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"high_availability_config": {
			// Property: HighAvailabilityConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Defines the details of a high availability configuration.",
			//	  "properties": {
			//	    "DesiredCapacity": {
			//	      "maximum": 100,
			//	      "minimum": 1,
			//	      "type": "integer"
			//	    }
			//	  },
			//	  "required": [
			//	    "DesiredCapacity"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "Defines the details of a high availability configuration.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"desired_capacity": {
						// Property: DesiredCapacity
						Type:     types.Int64Type,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(1, 100),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"instance_type": {
			// Property: InstanceType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The type of instance underlying the environment.",
			//	  "pattern": "^\\S{1,20}$",
			//	  "type": "string"
			//	}
			Description: "The type of instance underlying the environment.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^\\S{1,20}$"), ""),
			},
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.",
			//	  "maxLength": 2048,
			//	  "type": "string"
			//	}
			Description: "The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the environment.",
			//	  "pattern": "^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$",
			//	  "type": "string"
			//	}
			Description: "The name of the environment.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"preferred_maintenance_window": {
			// Property: PreferredMaintenanceWindow
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.",
			//	  "pattern": "^\\S{1,50}$",
			//	  "type": "string"
			//	}
			Description: "Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^\\S{1,50}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"publicly_accessible": {
			// Property: PubliclyAccessible
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Specifies whether the environment is publicly accessible.",
			//	  "type": "boolean"
			//	}
			Description: "Specifies whether the environment is publicly accessible.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"security_group_ids": {
			// Property: SecurityGroupIds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The list of security groups for the VPC associated with this environment.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "pattern": "^\\S{1,50}$",
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The list of security groups for the VPC associated with this environment.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^\\S{1,50}$"), "")),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"storage_configurations": {
			// Property: StorageConfigurations
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The storage configurations defined for the runtime environment.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "description": "Defines the storage configuration for an environment.",
			//	    "properties": {
			//	      "Efs": {
			//	        "additionalProperties": false,
			//	        "description": "Defines the storage configuration for an Amazon EFS file system.",
			//	        "properties": {
			//	          "FileSystemId": {
			//	            "description": "The file system identifier.",
			//	            "pattern": "^\\S{1,200}$",
			//	            "type": "string"
			//	          },
			//	          "MountPoint": {
			//	            "description": "The mount point for the file system.",
			//	            "pattern": "^\\S{1,200}$",
			//	            "type": "string"
			//	          }
			//	        },
			//	        "required": [
			//	          "FileSystemId",
			//	          "MountPoint"
			//	        ],
			//	        "type": "object"
			//	      },
			//	      "Fsx": {
			//	        "additionalProperties": false,
			//	        "description": "Defines the storage configuration for an Amazon FSx file system.",
			//	        "properties": {
			//	          "FileSystemId": {
			//	            "description": "The file system identifier.",
			//	            "pattern": "^\\S{1,200}$",
			//	            "type": "string"
			//	          },
			//	          "MountPoint": {
			//	            "description": "The mount point for the file system.",
			//	            "pattern": "^\\S{1,200}$",
			//	            "type": "string"
			//	          }
			//	        },
			//	        "required": [
			//	          "FileSystemId",
			//	          "MountPoint"
			//	        ],
			//	        "type": "object"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The storage configurations defined for the runtime environment.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"efs": {
						// Property: Efs
						Description: "Defines the storage configuration for an Amazon EFS file system.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"file_system_id": {
									// Property: FileSystemId
									Description: "The file system identifier.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("^\\S{1,200}$"), ""),
									},
								},
								"mount_point": {
									// Property: MountPoint
									Description: "The mount point for the file system.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("^\\S{1,200}$"), ""),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"fsx": {
						// Property: Fsx
						Description: "Defines the storage configuration for an Amazon FSx file system.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"file_system_id": {
									// Property: FileSystemId
									Description: "The file system identifier.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("^\\S{1,200}$"), ""),
									},
								},
								"mount_point": {
									// Property: MountPoint
									Description: "The mount point for the file system.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile("^\\S{1,200}$"), ""),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"subnet_ids": {
			// Property: SubnetIds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The unique identifiers of the subnets assigned to this runtime environment.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "pattern": "^\\S{1,50}$",
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The unique identifiers of the subnets assigned to this runtime environment.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^\\S{1,50}$"), "")),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Tags associated to this environment.",
			//	  "patternProperties": {
			//	    "": {
			//	      "maxLength": 256,
			//	      "minLength": 0,
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Tags associated to this environment.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
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
		Description: "Represents a runtime environment that can run migrated mainframe applications.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::M2::Environment").WithTerraformTypeName("awscc_m2_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":                  "Description",
		"desired_capacity":             "DesiredCapacity",
		"efs":                          "Efs",
		"engine_type":                  "EngineType",
		"engine_version":               "EngineVersion",
		"environment_arn":              "EnvironmentArn",
		"environment_id":               "EnvironmentId",
		"file_system_id":               "FileSystemId",
		"fsx":                          "Fsx",
		"high_availability_config":     "HighAvailabilityConfig",
		"instance_type":                "InstanceType",
		"kms_key_id":                   "KmsKeyId",
		"mount_point":                  "MountPoint",
		"name":                         "Name",
		"preferred_maintenance_window": "PreferredMaintenanceWindow",
		"publicly_accessible":          "PubliclyAccessible",
		"security_group_ids":           "SecurityGroupIds",
		"storage_configurations":       "StorageConfigurations",
		"subnet_ids":                   "SubnetIds",
		"tags":                         "Tags",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
