// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_elasticbeanstalk_application", applicationDataSource)
}

// applicationDataSource returns the Terraform awscc_elasticbeanstalk_application data source.
// This Terraform data source corresponds to the CloudFormation AWS::ElasticBeanstalk::Application resource.
func applicationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"application_name": {
			// Property: ApplicationName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.",
			//	  "type": "string"
			//	}
			Description: "A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Your description of the application.",
			//	  "type": "string"
			//	}
			Description: "Your description of the application.",
			Type:        types.StringType,
			Computed:    true,
		},
		"resource_lifecycle_config": {
			// Property: ResourceLifecycleConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.",
			//	  "properties": {
			//	    "ServiceRole": {
			//	      "description": "The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a ResourceLifecycleConfig for the application. After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again. You can, however, specify it in subsequent updates to change the Service Role to another value.",
			//	      "type": "string"
			//	    },
			//	    "VersionLifecycleConfig": {
			//	      "additionalProperties": false,
			//	      "description": "Defines lifecycle settings for application versions.",
			//	      "properties": {
			//	        "MaxAgeRule": {
			//	          "additionalProperties": false,
			//	          "description": "Specify a max age rule to restrict the length of time that application versions are retained for an application.",
			//	          "properties": {
			//	            "DeleteSourceFromS3": {
			//	              "description": "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
			//	              "type": "boolean"
			//	            },
			//	            "Enabled": {
			//	              "description": "Specify true to apply the rule, or false to disable it.",
			//	              "type": "boolean"
			//	            },
			//	            "MaxAgeInDays": {
			//	              "description": "Specify the number of days to retain an application versions.",
			//	              "type": "integer"
			//	            }
			//	          },
			//	          "type": "object"
			//	        },
			//	        "MaxCountRule": {
			//	          "additionalProperties": false,
			//	          "description": "Specify a max count rule to restrict the number of application versions that are retained for an application.",
			//	          "properties": {
			//	            "DeleteSourceFromS3": {
			//	              "description": "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
			//	              "type": "boolean"
			//	            },
			//	            "Enabled": {
			//	              "description": "Specify true to apply the rule, or false to disable it.",
			//	              "type": "boolean"
			//	            },
			//	            "MaxCount": {
			//	              "description": "Specify the maximum number of application versions to retain.",
			//	              "type": "integer"
			//	            }
			//	          },
			//	          "type": "object"
			//	        }
			//	      },
			//	      "type": "object"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"service_role": {
						// Property: ServiceRole
						Description: "The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a ResourceLifecycleConfig for the application. After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again. You can, however, specify it in subsequent updates to change the Service Role to another value.",
						Type:        types.StringType,
						Computed:    true,
					},
					"version_lifecycle_config": {
						// Property: VersionLifecycleConfig
						Description: "Defines lifecycle settings for application versions.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"max_age_rule": {
									// Property: MaxAgeRule
									Description: "Specify a max age rule to restrict the length of time that application versions are retained for an application.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"delete_source_from_s3": {
												// Property: DeleteSourceFromS3
												Description: "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
												Type:        types.BoolType,
												Computed:    true,
											},
											"enabled": {
												// Property: Enabled
												Description: "Specify true to apply the rule, or false to disable it.",
												Type:        types.BoolType,
												Computed:    true,
											},
											"max_age_in_days": {
												// Property: MaxAgeInDays
												Description: "Specify the number of days to retain an application versions.",
												Type:        types.Int64Type,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"max_count_rule": {
									// Property: MaxCountRule
									Description: "Specify a max count rule to restrict the number of application versions that are retained for an application.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"delete_source_from_s3": {
												// Property: DeleteSourceFromS3
												Description: "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
												Type:        types.BoolType,
												Computed:    true,
											},
											"enabled": {
												// Property: Enabled
												Description: "Specify true to apply the rule, or false to disable it.",
												Type:        types.BoolType,
												Computed:    true,
											},
											"max_count": {
												// Property: MaxCount
												Description: "Specify the maximum number of application versions to retain.",
												Type:        types.Int64Type,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ElasticBeanstalk::Application",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticBeanstalk::Application").WithTerraformTypeName("awscc_elasticbeanstalk_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_name":          "ApplicationName",
		"delete_source_from_s3":     "DeleteSourceFromS3",
		"description":               "Description",
		"enabled":                   "Enabled",
		"max_age_in_days":           "MaxAgeInDays",
		"max_age_rule":              "MaxAgeRule",
		"max_count":                 "MaxCount",
		"max_count_rule":            "MaxCountRule",
		"resource_lifecycle_config": "ResourceLifecycleConfig",
		"service_role":              "ServiceRole",
		"version_lifecycle_config":  "VersionLifecycleConfig",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
