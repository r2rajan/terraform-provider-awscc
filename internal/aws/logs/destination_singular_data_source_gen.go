// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_logs_destination", destinationDataSource)
}

// destinationDataSource returns the Terraform awscc_logs_destination data source.
// This Terraform data source corresponds to the CloudFormation AWS::Logs::Destination resource.
func destinationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
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
		"destination_name": {
			// Property: DestinationName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the destination resource",
			//	  "maxLength": 512,
			//	  "minLength": 1,
			//	  "pattern": "^[^:*]{1,512}$",
			//	  "type": "string"
			//	}
			Description: "The name of the destination resource",
			Type:        types.StringType,
			Computed:    true,
		},
		"destination_policy": {
			// Property: DestinationPolicy
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An IAM policy document that governs which AWS accounts can create subscription filters against this destination.",
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "An IAM policy document that governs which AWS accounts can create subscription filters against this destination.",
			Type:        types.StringType,
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource",
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource",
			Type:        types.StringType,
			Computed:    true,
		},
		"target_arn": {
			// Property: TargetArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)",
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Logs::Destination",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::Destination").WithTerraformTypeName("awscc_logs_destination")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"destination_name":   "DestinationName",
		"destination_policy": "DestinationPolicy",
		"role_arn":           "RoleArn",
		"target_arn":         "TargetArn",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
