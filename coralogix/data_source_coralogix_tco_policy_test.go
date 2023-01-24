package coralogix

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var tcoPolicyDataSourceName = "data." + tcoPolicyResourceName

func TestAccCoralogixDataSourceTCOPolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceTCOPolicy() +
					testAccCoralogixResourceTCOPolicy_read(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(tcoPolicyDataSourceName, "id"),
				),
			},
		},
	})
}

func testAccCoralogixResourceTCOPolicy_read() string {
	return `data "coralogix_tco_policy" "test" {
		id = coralogix_tco_policy.test.id
}
`
}