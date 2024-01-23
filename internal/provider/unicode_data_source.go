package provider

import (
	"context"
	//unicode_client "terraform-unicode-pranoSA/internal/unicode"
	unicode_client "terraform-provider-unicode/internal/unicode"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &UnicodeDataSource{}
var _ datasource.DataSourceWithConfigure = &UnicodeDataSource{}

type UnicodeDataSource struct {
	unicodeClient *unicode_client.UnicodeProviderClient
}

func MakeUnicodeDataSource() func() datasource.DataSource {
	return func() datasource.DataSource {
		return &UnicodeDataSource{}
	}
}

/*func NewUnicodeDataSource(unicodeClient *unicode_client.UnicodeProviderClient) *UnicodeDataSource {
	return &UnicodeDataSource{unicodeClient: unicodeClient}
}*/

/**
 *
 * Now I need to implement the Metadata, Schema and Read methods.
 *
 * Metadata : Configure the Resorce Identifier Essentially
 *
 * Schema: Identify the attributes that will be returned by the data source
 * Also Attributes that are required ?
 *
 * Required : true,
 * Computed True
 * Type of Attribute
 *
 */

func (d *UnicodeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_unicode_chars"
}

func (d *UnicodeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {

	resp.Diagnostics.AddWarning("configure not implemented", "")

	unicodeClient, ok := req.ProviderData.(*unicode_client.UnicodeProviderClient)

	if unicodeClient == nil {
		resp.Diagnostics.AddWarning("invalid provider configuration", "IS NULL")
	}

	if !ok {
		resp.Diagnostics.AddWarning("invalid provider configuration", "")
		unicodeClient = unicode_client.NewUnicodeProviderClient("bob")
		//return
	}

	d.unicodeClient = unicodeClient

	return
}

func (d *UnicodeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Unicode data source",
		Attributes: map[string]schema.Attribute{
			"unicode_char": schema.StringAttribute{
				MarkdownDescription: "Unicode character",
				Computed:            false,
				Required:            true,
			},
			"unicode_name": schema.StringAttribute{
				MarkdownDescription: "Unicode name",
				Computed:            true,
			},
			"unicode_block": schema.StringAttribute{
				MarkdownDescription: "Unicode block",
				Computed:            true,
			},
			"unicode_category": schema.StringAttribute{
				MarkdownDescription: "Unicode category",
				Computed:            true,
			},
		},
	}
}

type UnicodeRequest struct {
	UnicodeChar string `tfsdk:"unicode_char"`
}

func (d *UnicodeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var unicodeChar UnicodeRequest

	var res unicode_client.UnicodeData

	req.Config.Get(ctx, &unicodeChar)

	unicodeCharData, err := d.unicodeClient.GetUnicodeCharData(unicodeChar.UnicodeChar)
	if err != nil {
		return
	}

	res.Char = unicodeCharData.Char
	res.Block = unicodeCharData.Block
	res.Category = unicodeCharData.Category
	res.Name = unicodeCharData.Name

	diags := resp.State.Set(context.Background(), &res)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (d *UnicodeDataSource) DataSourceWithConfigure() {}
