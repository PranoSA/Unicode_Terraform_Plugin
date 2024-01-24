// Copyright (c) HashiCorp, Inc.l
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"
	unicode_client "terraform-provider-unicode/internal/unicode"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var (
	_ provider.Provider = &unicodeProvider{}
)

func NewUnicodeDataSource() datasource.DataSource {
	return &UnicodeDataSource{}
}

type unicodeProvider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &unicodeProvider{
			version: version,
		}
	}
}

// hashicupsProvider is the provider implementation.
type hashicupsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// hashicupsProviderModel maps provider schema data to a Go type.

func (up *unicodeProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "unicode"
	resp.Version = up.version

	//Metadata missing
}

// Schema defines the provider-level schema for configuration data.
// Schema defines the provider-level schema for configuration data.

func (up *unicodeProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"user": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (up *unicodeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config unicodeProviderModel
	diags := req.Config.Get(ctx, &config)

	resp.Diagnostics.AddWarning("configure not implemented", config.User.String())

	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	//client := unicode_client.NewUnicodeProviderClient(config.User.String())
	client1 := unicode_client.NewUnicodeProviderClient("bob")

	resp.DataSourceData = client1

	if resp.DataSourceData == nil {
		resp.Diagnostics.AddWarning("Unable to create client", "Client is NULL After NewUnicodeProviderClient")
		//return
	}

	resp.Diagnostics.AddWarning("configure not implemented 2", client1.Username)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.User.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("user"),
			"Unknown User",
			"User is unknown",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	username := os.Getenv("USER")

	if !config.User.IsNull() {
		username = config.User.ValueString()
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("user"),
			"Missing User",
			"User is missing",
		)
	}

	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError("Unable to create clien 2t", "Client is NULL After NewUnicodeProviderClient")
		return
	}

	client := unicode_client.NewUnicodeProviderClient(username)

	if client == nil {
		resp.Diagnostics.AddError("Unable to create client provider", "Client is NULL After NewUnicodeProviderClient")
		return
	}

	//resp.DataSourceData = client
	resp.Diagnostics.AddWarning("PLEASE MAKE IT !!", "PLEASE MAKE IT !!")

	resp.ResourceData = client

}

// DataSources defines the data sources implemented in the provider.
func (up *unicodeProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewUnicodeDataSource,
	}
}

func (up *unicodeProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewUnicodeAppResource,
		NewUnicodeStringResource,
	}
}

func (up *unicodeProviderModel) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "unicode"
	resp.Version = up.User.String()
}

func (up *unicodeProviderModel) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"user": schema.StringAttribute{
				Optional: false,
			},
		},
	}
}

type unicodeProviderModel struct {
	User types.String `tfsdk:"user"`
}
