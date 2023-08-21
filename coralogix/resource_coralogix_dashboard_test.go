package coralogix

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"google.golang.org/protobuf/types/known/wrapperspb"
	"terraform-provider-coralogix/coralogix/clientset"
	dashboard "terraform-provider-coralogix/coralogix/clientset/grpc/coralogix-dashboards/v1"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var dashboardResourceName = "coralogix_dashboard.test"

func TestAccCoralogixResourceDashboard(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDashboardDestroy,
		Steps: []resource.TestStep{
			{

				Config: testAccCoralogixResourceDashboard(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(dashboardResourceName, "id"),
					resource.TestCheckResourceAttr(dashboardResourceName, "name", "test"),
					resource.TestCheckResourceAttr(dashboardResourceName, "description", "dashboards team is messing with this 🗿"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.height", "19"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.0.title", "status 4XX"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.0.definition.line_chart.query_definitions.0.query.metrics.promql_query", "http_requests_total{status!~\"4..\"}"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.0.definition.line_chart.legend.is_visible", "true"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.0.definition.line_chart.legend.columns.0", "max"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.0.definition.line_chart.legend.columns.1", "last"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.0.width", "0"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.1.title", "count"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.1.definition.line_chart.query_definitions.0.query.logs.aggregations.0.type", "count"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.1.definition.line_chart.legend.is_visible", "true"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.1.definition.line_chart.legend.columns.0", "min"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.1.definition.line_chart.legend.columns.1", "max"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.1.definition.line_chart.legend.columns.2", "sum"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.1.definition.line_chart.legend.columns.3", "avg"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.1.definition.line_chart.legend.columns.4", "last"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.1.width", "10"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.2.title", "error throwing pods"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.2.definition.line_chart.query_definitions.0.query.logs.lucene_query", "coralogix.metadata.severity=5 OR coralogix.metadata.severity=\"6\" OR coralogix.metadata.severity=\"4\""),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.2.definition.line_chart.query_definitions.0.query.logs.group_by.0", "coralogix.metadata.subsystemName"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.2.definition.line_chart.query_definitions.0.query.logs.aggregations.0.type", "count"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.2.definition.line_chart.legend.is_visible", "true"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.2.definition.line_chart.legend.columns.0", "max"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.2.definition.line_chart.legend.columns.1", "last"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.0.widgets.2.width", "0"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.height", "28"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.title", "dashboards-api logz"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.description", "warnings, errors, criticals"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.definition.data_table.query.logs.filters.0.field", "coralogix.metadata.applicationName"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.definition.data_table.query.logs.filters.0.operator.type", "equals"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.definition.data_table.query.logs.filters.0.operator.selected_values.0", "staging"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.definition.data_table.results_per_page", "20"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.definition.data_table.row_style", "one_line"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.definition.data_table.columns.0.field", "coralogix.timestamp"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.definition.data_table.columns.1.field", "textObject.textObject.textObject.kubernetes.pod_id"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.definition.data_table.columns.2.field", "coralogix.text"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.definition.data_table.columns.3.field", "coralogix.metadata.applicationName"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.definition.data_table.columns.4.field", "coralogix.metadata.subsystemName"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.definition.data_table.columns.5.field", "coralogix.metadata.sdkId"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.definition.data_table.columns.6.field", "textObject.log_obj.e2e_test.config"),
					resource.TestCheckResourceAttr(dashboardResourceName, "layout.sections.0.rows.1.widgets.0.width", "0"),
					resource.TestCheckResourceAttr(dashboardResourceName, "variables.0.name", "test_variable"),
					resource.TestCheckResourceAttr(dashboardResourceName, "variables.0.definition.multi_select.selected_values.0", "1"),
					resource.TestCheckResourceAttr(dashboardResourceName, "variables.0.definition.multi_select.selected_values.1", "2"),
					resource.TestCheckResourceAttr(dashboardResourceName, "variables.0.definition.multi_select.selected_values.2", "3"),
					resource.TestCheckResourceAttr(dashboardResourceName, "variables.0.definition.multi_select.source.constant_list.0", "1"),
					resource.TestCheckResourceAttr(dashboardResourceName, "variables.0.definition.multi_select.source.constant_list.1", "2"),
					resource.TestCheckResourceAttr(dashboardResourceName, "variables.0.definition.multi_select.source.constant_list.2", "3"),
				),
			},
			{
				ResourceName:      dashboardResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCoralogixResourceDashboardFromJson(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	filePath := parent + "/examples/dashboard/dashboard.json"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDashboardDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceDashboardFromJson(filePath),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(dashboardResourceName, "id"),
				),
			},
		},
	})
}

func testAccCheckDashboardDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*clientset.ClientSet).Dashboards()

	ctx := context.TODO()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "coralogix_dashboard" {
			continue
		}

		dashboardId := wrapperspb.String(rs.Primary.ID)
		resp, err := client.GetDashboard(ctx, &dashboard.GetDashboardRequest{DashboardId: dashboardId})
		if err == nil {
			if resp.GetDashboard().GetId().GetValue() == rs.Primary.ID {
				return fmt.Errorf("dashboard still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCoralogixResourceDashboard() string {
	return `resource "coralogix_dashboard" test {
  name        = "test"
  description = "dashboards team is messing with this 🗿"
  layout      = {
    sections = [
      {
        rows = [
          {
            height  = 19
            widgets = [
              {
                title      = "status 4XX"
                definition = {
                  line_chart = {
                    query_definitions = [
                      {
                        query = {
                          metrics = {
                            promql_query = "http_requests_total{status!~\"4..\"}"
                          }
                        }
                      },
                    ]
                    legend = {
                      is_visible = true
                      columns     = ["max", "last"]
                    }
                  }
                }
                width = 0
              },
              {
                title      = "count"
                definition = {
                  line_chart = {
                    query_definitions = [
                      {
                        query = {
                          logs = {
                            aggregations = [
                              {
                                type = "count"
                              },
                            ]
                          }
                        }
                      },
                    ]
			      legend = {
                   		is_visible = true
                   		 columns     = ["min", "max", "sum", "avg", "last"]
                  	}
                  } 
                }
                width = 10
              },
              {
                title      = "error throwing pods"
                definition = {
                  line_chart = {
                    query_definitions = [
                      {
                        query = {
                          logs = {
                            lucene_query = "coralogix.metadata.severity=5 OR coralogix.metadata.severity=\"6\" OR coralogix.metadata.severity=\"4\""
                            group_by     = ["coralogix.metadata.subsystemName"]
                            aggregations = [
                              {
                                type = "count"
                              },
                            ]
                          }
                        }
                      },
                    ]
                    legend = {
                      is_visible = true
                      columns     = ["max", "last"]
                    }
                  }
                }
                width = 0
              }
            ]
          },
          {
            height  = 28
            widgets = [
              {
                title       = "dashboards-api logz"
                description = "warnings, errors, criticals"
                definition  = {
                  data_table = {
                    query = {
                      logs = {
                        filters = [
                          {
                            field    = "coralogix.metadata.applicationName"
                            operator = {
                              type            = "equals"
                              selected_values = ["staging"]
                            }
                          }
                        ]
                      }
                    }
                    results_per_page = 20
                    row_style        = "one_line"
                    columns          = [
                      {
                        field = "coralogix.timestamp"
                      },
                      {
                        field = "textObject.textObject.textObject.kubernetes.pod_id"
                      },
                      {
                        field = "coralogix.text"
                      },
                      {
                        field = "coralogix.metadata.applicationName"
                      },
                      {
                        field = "coralogix.metadata.subsystemName"
                      },
                      {
                        field = "coralogix.metadata.sdkId"
                      },
                      {
                        field = "textObject.log_obj.e2e_test.config"
                      },
                    ]
                  }
                }
                width = 0
              }
            ],
          },
        ]
      },
    ]
  }
  variables = [
    {
      name       = "test_variable"
      definition = {
        multi_select = {
          selected_values = ["1", "2", "3"]
          source          = {
            constant_list = ["1", "2", "3"]
          }
        }
      }
    },
  ]
}
`
}

func testAccCoralogixResourceDashboardFromJson(jsonFilePath string) string {
	return fmt.Sprintf(`resource "coralogix_dashboard" test {
   		content_json = file("%s")
	}
`, jsonFilePath)
}
