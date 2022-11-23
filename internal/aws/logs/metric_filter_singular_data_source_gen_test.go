// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package logs_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSLogsMetricFilterDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Logs::MetricFilter", "awscc_logs_metric_filter", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSLogsMetricFilterDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Logs::MetricFilter", "awscc_logs_metric_filter", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
