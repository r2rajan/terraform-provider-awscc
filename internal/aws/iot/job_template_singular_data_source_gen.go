// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iot_job_template", jobTemplateDataSource)
}

// jobTemplateDataSource returns the Terraform awscc_iot_job_template data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoT::JobTemplate resource.
func jobTemplateDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"abort_config": {
			// Property: AbortConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The criteria that determine when and how a job abort takes place.",
			//	  "properties": {
			//	    "CriteriaList": {
			//	      "insertionOrder": false,
			//	      "items": {
			//	        "additionalProperties": false,
			//	        "description": "The criteria that determine when and how a job abort takes place.",
			//	        "properties": {
			//	          "Action": {
			//	            "description": "The type of job action to take to initiate the job abort.",
			//	            "enum": [
			//	              "CANCEL"
			//	            ],
			//	            "type": "string"
			//	          },
			//	          "FailureType": {
			//	            "description": "The type of job execution failures that can initiate a job abort.",
			//	            "enum": [
			//	              "FAILED",
			//	              "REJECTED",
			//	              "TIMED_OUT",
			//	              "ALL"
			//	            ],
			//	            "type": "string"
			//	          },
			//	          "MinNumberOfExecutedThings": {
			//	            "description": "The minimum number of things which must receive job execution notifications before the job can be aborted.",
			//	            "minimum": 1,
			//	            "type": "integer"
			//	          },
			//	          "ThresholdPercentage": {
			//	            "description": "The minimum percentage of job execution failures that must occur to initiate the job abort.",
			//	            "maximum": 100,
			//	            "type": "number"
			//	          }
			//	        },
			//	        "required": [
			//	          "Action",
			//	          "FailureType",
			//	          "MinNumberOfExecutedThings",
			//	          "ThresholdPercentage"
			//	        ],
			//	        "type": "object"
			//	      },
			//	      "minItems": 1,
			//	      "type": "array"
			//	    }
			//	  },
			//	  "required": [
			//	    "CriteriaList"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "The criteria that determine when and how a job abort takes place.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"criteria_list": {
						// Property: CriteriaList
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"action": {
									// Property: Action
									Description: "The type of job action to take to initiate the job abort.",
									Type:        types.StringType,
									Computed:    true,
								},
								"failure_type": {
									// Property: FailureType
									Description: "The type of job execution failures that can initiate a job abort.",
									Type:        types.StringType,
									Computed:    true,
								},
								"min_number_of_executed_things": {
									// Property: MinNumberOfExecutedThings
									Description: "The minimum number of things which must receive job execution notifications before the job can be aborted.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"threshold_percentage": {
									// Property: ThresholdPercentage
									Description: "The minimum percentage of job execution failures that must occur to initiate the job abort.",
									Type:        types.Float64Type,
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
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A description of the Job Template.",
			//	  "maxLength": 2028,
			//	  "pattern": "[^\\p{C}]+",
			//	  "type": "string"
			//	}
			Description: "A description of the Job Template.",
			Type:        types.StringType,
			Computed:    true,
		},
		"document": {
			// Property: Document
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The job document. Required if you don't specify a value for documentSource.",
			//	  "maxLength": 32768,
			//	  "type": "string"
			//	}
			Description: "The job document. Required if you don't specify a value for documentSource.",
			Type:        types.StringType,
			Computed:    true,
		},
		"document_source": {
			// Property: DocumentSource
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An S3 link to the job document to use in the template. Required if you don't specify a value for document.",
			//	  "maxLength": 1350,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "An S3 link to the job document to use in the template. Required if you don't specify a value for document.",
			Type:        types.StringType,
			Computed:    true,
		},
		"job_arn": {
			// Property: JobArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Optional for copying a JobTemplate from a pre-existing Job configuration.",
			//	  "type": "string"
			//	}
			Description: "Optional for copying a JobTemplate from a pre-existing Job configuration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"job_executions_rollout_config": {
			// Property: JobExecutionsRolloutConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Allows you to create a staged rollout of a job.",
			//	  "properties": {
			//	    "ExponentialRolloutRate": {
			//	      "additionalProperties": false,
			//	      "description": "The rate of increase for a job rollout. This parameter allows you to define an exponential rate for a job rollout.",
			//	      "properties": {
			//	        "BaseRatePerMinute": {
			//	          "description": "The minimum number of things that will be notified of a pending job, per minute at the start of job rollout. This parameter allows you to define the initial rate of rollout.",
			//	          "minimum": 1,
			//	          "type": "integer"
			//	        },
			//	        "IncrementFactor": {
			//	          "description": "The exponential factor to increase the rate of rollout for a job.",
			//	          "maximum": 5,
			//	          "minimum": 1,
			//	          "type": "number"
			//	        },
			//	        "RateIncreaseCriteria": {
			//	          "additionalProperties": false,
			//	          "description": "The criteria to initiate the increase in rate of rollout for a job.",
			//	          "properties": {
			//	            "NumberOfNotifiedThings": {
			//	              "minimum": 1,
			//	              "type": "integer"
			//	            },
			//	            "NumberOfSucceededThings": {
			//	              "minimum": 1,
			//	              "type": "integer"
			//	            }
			//	          },
			//	          "type": "object"
			//	        }
			//	      },
			//	      "required": [
			//	        "BaseRatePerMinute",
			//	        "IncrementFactor",
			//	        "RateIncreaseCriteria"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "MaximumPerMinute": {
			//	      "description": "The maximum number of things that will be notified of a pending job, per minute. This parameter allows you to create a staged rollout.",
			//	      "minimum": 1,
			//	      "type": "integer"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Allows you to create a staged rollout of a job.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"exponential_rollout_rate": {
						// Property: ExponentialRolloutRate
						Description: "The rate of increase for a job rollout. This parameter allows you to define an exponential rate for a job rollout.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"base_rate_per_minute": {
									// Property: BaseRatePerMinute
									Description: "The minimum number of things that will be notified of a pending job, per minute at the start of job rollout. This parameter allows you to define the initial rate of rollout.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"increment_factor": {
									// Property: IncrementFactor
									Description: "The exponential factor to increase the rate of rollout for a job.",
									Type:        types.Float64Type,
									Computed:    true,
								},
								"rate_increase_criteria": {
									// Property: RateIncreaseCriteria
									Description: "The criteria to initiate the increase in rate of rollout for a job.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"number_of_notified_things": {
												// Property: NumberOfNotifiedThings
												Type:     types.Int64Type,
												Computed: true,
											},
											"number_of_succeeded_things": {
												// Property: NumberOfSucceededThings
												Type:     types.Int64Type,
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
					"maximum_per_minute": {
						// Property: MaximumPerMinute
						Description: "The maximum number of things that will be notified of a pending job, per minute. This parameter allows you to create a staged rollout.",
						Type:        types.Int64Type,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"job_template_id": {
			// Property: JobTemplateId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 64,
			//	  "minLength": 1,
			//	  "pattern": "[a-zA-Z0-9_-]+",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"presigned_url_config": {
			// Property: PresignedUrlConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Configuration for pre-signed S3 URLs.",
			//	  "properties": {
			//	    "ExpiresInSec": {
			//	      "description": "How number (in seconds) pre-signed URLs are valid.",
			//	      "maximum": 3600,
			//	      "minimum": 60,
			//	      "type": "integer"
			//	    },
			//	    "RoleArn": {
			//	      "description": "The ARN of an IAM role that grants grants permission to download files from the S3 bucket where the job data/updates are stored. The role must also grant permission for IoT to download the files.",
			//	      "maxLength": 2048,
			//	      "minLength": 20,
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "RoleArn"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "Configuration for pre-signed S3 URLs.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"expires_in_sec": {
						// Property: ExpiresInSec
						Description: "How number (in seconds) pre-signed URLs are valid.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"role_arn": {
						// Property: RoleArn
						Description: "The ARN of an IAM role that grants grants permission to download files from the S3 bucket where the job data/updates are stored. The role must also grant permission for IoT to download the files.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Metadata that can be used to manage the JobTemplate.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair to associate with a resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The tag's key.",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The tag's value.",
			//	        "maxLength": 256,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Value",
			//	      "Key"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 50,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "Metadata that can be used to manage the JobTemplate.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The tag's key.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The tag's value.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"timeout_config": {
			// Property: TimeoutConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Specifies the amount of time each device has to finish its execution of the job.",
			//	  "properties": {
			//	    "InProgressTimeoutInMinutes": {
			//	      "description": "Specifies the amount of time, in minutes, this device has to finish execution of this job.",
			//	      "maximum": 10080,
			//	      "minimum": 1,
			//	      "type": "integer"
			//	    }
			//	  },
			//	  "required": [
			//	    "InProgressTimeoutInMinutes"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "Specifies the amount of time each device has to finish its execution of the job.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"in_progress_timeout_in_minutes": {
						// Property: InProgressTimeoutInMinutes
						Description: "Specifies the amount of time, in minutes, this device has to finish execution of this job.",
						Type:        types.Int64Type,
						Computed:    true,
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
		Description: "Data Source schema for AWS::IoT::JobTemplate",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::JobTemplate").WithTerraformTypeName("awscc_iot_job_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"abort_config":                   "AbortConfig",
		"action":                         "Action",
		"arn":                            "Arn",
		"base_rate_per_minute":           "BaseRatePerMinute",
		"criteria_list":                  "CriteriaList",
		"description":                    "Description",
		"document":                       "Document",
		"document_source":                "DocumentSource",
		"expires_in_sec":                 "ExpiresInSec",
		"exponential_rollout_rate":       "ExponentialRolloutRate",
		"failure_type":                   "FailureType",
		"in_progress_timeout_in_minutes": "InProgressTimeoutInMinutes",
		"increment_factor":               "IncrementFactor",
		"job_arn":                        "JobArn",
		"job_executions_rollout_config":  "JobExecutionsRolloutConfig",
		"job_template_id":                "JobTemplateId",
		"key":                            "Key",
		"maximum_per_minute":             "MaximumPerMinute",
		"min_number_of_executed_things":  "MinNumberOfExecutedThings",
		"number_of_notified_things":      "NumberOfNotifiedThings",
		"number_of_succeeded_things":     "NumberOfSucceededThings",
		"presigned_url_config":           "PresignedUrlConfig",
		"rate_increase_criteria":         "RateIncreaseCriteria",
		"role_arn":                       "RoleArn",
		"tags":                           "Tags",
		"threshold_percentage":           "ThresholdPercentage",
		"timeout_config":                 "TimeoutConfig",
		"value":                          "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
