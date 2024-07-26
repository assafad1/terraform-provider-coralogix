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

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"terraform-provider-coralogix/coralogix/clientset"
)

var _ datasource.DataSourceWithConfigure = &ArchiveMetricsDataSource{}

func NewArchiveMetricsDataSource() datasource.DataSource {
	return &ArchiveMetricsDataSource{}
}

type ArchiveMetricsDataSource struct {
	client *clientset.ArchiveMetricsClient
}

func (d *ArchiveMetricsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_archive_metrics"
}

func (d *ArchiveMetricsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	clientSet, ok := req.ProviderData.(*clientset.ClientSet)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *clientset.ClientSet, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = clientSet.ArchiveMetrics()
}

func (d *ArchiveMetricsDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	var r ArchiveMetricsResource
	var resourceResp resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &resourceResp)

	resp.Schema = convertSchemaWithoutID(resourceResp.Schema)
}

func (d *ArchiveMetricsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ArchiveMetricsResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get refreshed archive-metrics value from Coralogix
	id := data.ID.ValueString()
	log.Print("[INFO] Reading archive-metrics")
	getResp, err := d.client.GetArchiveMetrics(ctx)
	if err != nil {
		log.Printf("[ERROR] Received error: %s", err.Error())
		if status.Code(err) == codes.NotFound {
			resp.Diagnostics.AddWarning(
				err.Error(),
				fmt.Sprintf("archive-metrics %q is in state, but no longer exists in Coralogix backend", id),
			)
		} else {
			resp.Diagnostics.AddError(
				"Error reading archive-metrics",
				formatRpcErrors(err, getArchiveMetricsURL, ""),
			)
		}
		return
	}
	log.Printf("[INFO] Received archive-metrics: %s", protojson.Format(getResp))

	data, diags := flattenArchiveMetrics(ctx, getResp.GetTenantConfig())
	if diags.HasError() {
		resp.Diagnostics = diags
		return
	}
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
