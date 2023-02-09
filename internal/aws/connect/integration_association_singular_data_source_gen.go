// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_connect_integration_association", integrationAssociationDataSource)
}

// integrationAssociationDataSource returns the Terraform awscc_connect_integration_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::Connect::IntegrationAssociation resource.
func integrationAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: InstanceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Connect instance identifier",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"instance_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Connect instance identifier",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IntegrationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of Integration being associated with the instance",
		//	  "maxLength": 140,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"integration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of Integration being associated with the instance",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IntegrationAssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Identifier of the association with Connect Instance",
		//	  "pattern": "^[a-zA-Z]{1}(?:-?[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"integration_association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Identifier of the association with Connect Instance",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IntegrationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the integration type to be associated with the instance",
		//	  "enum": [
		//	    "LEX_BOT",
		//	    "LAMBDA_FUNCTION"
		//	  ],
		//	  "type": "string"
		//	}
		"integration_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the integration type to be associated with the instance",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Connect::IntegrationAssociation",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::IntegrationAssociation").WithTerraformTypeName("awscc_connect_integration_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"instance_id":                "InstanceId",
		"integration_arn":            "IntegrationArn",
		"integration_association_id": "IntegrationAssociationId",
		"integration_type":           "IntegrationType",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
