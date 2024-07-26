// Copyright 2024 Coralogix Ltd.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     https://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package coralogix

var (
	archiveLogsResourceName = "coralogix_archive_logs.test"
)

//func TestAccCoralogixResourceResourceArchiveLogs(t *testing.T) {
//	resource.Test(t, resource.TestCase{
//		PreCheck:                 func() { testAccPreCheck(t) },
//		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccCoralogixResourceArchiveLogs(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					resource.TestCheckResourceAttr(archiveLogsResourceName, "bucket", "coralogix-c4c-eu2-prometheus-data"),
//					resource.TestCheckResourceAttr(archiveLogsResourceName, "active", "true"),
//				),
//			},
//			{
//				ResourceName:      archiveLogsResourceName,
//				ImportState:       true,
//				ImportStateVerify: true,
//			},
//			{
//				Config: testAccCoralogixResourceArchiveLogsUpdate(),
//				Check: resource.ComposeAggregateTestCheckFunc(
//					resource.TestCheckResourceAttr(archiveLogsResourceName, "bucket", "coralogix-c4c-eu2-prometheus-data"),
//					resource.TestCheckResourceAttr(archiveLogsResourceName, "active", "false"),
//				),
//			},
//		},
//	})
//}

func testAccCoralogixResourceArchiveLogs() string {
	return `resource "coralogix_archive_logs" "test" {
 	bucket = "coralogix-c4c-eu2-prometheus-data"
}
`
}

func testAccCoralogixResourceArchiveLogsUpdate() string {
	return `resource "coralogix_archive_logs" "test" {
  		bucket = coralogix-c4c-eu2-prometheus-data
 		active = false
}
`
}
