package provider

import (
	"context"
	unicode_client "terraform-provider-unicode/internal/unicode"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	_ resource.Resource              = &unicodeStringResource{}
	_ resource.ResourceWithConfigure = &unicodeStringResource{}
)

func NewUnicodeStringResource() resource.Resource {
	return &unicodeStringResource{}
}

type unicodeStringResource struct {
	unicodeClient *unicode_client.UnicodeProviderClient
}

func (r *unicodeStringResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_unicode_string"
}

func (r *unicodeStringResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	//resp.State = req.NewState
	//req.Config.Get(context.Background(), &plan)

	//r.unicodeClient = unicode_client.NewUnicodeProviderClient("test")
	unicodeClient, ok := req.ProviderData.(*unicode_client.UnicodeProviderClient)

	if !ok {
		//resp.Diagnostics.AddError("Unable to create client in resource", "Client is NULL After NewUnicodeProviderClient")
		//return
		resp.Diagnostics.AddWarning("Unable to create client in resource", "Client is NULL After NewUnicodeProviderClient")
		unicodeClient = unicode_client.NewUnicodeProviderClient("bob")
	}

	r.unicodeClient = unicodeClient
}

func (r *unicodeStringResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {

	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"app_id": schema.StringAttribute{
				Required: true,
			},
			"value": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
		},
	}
}

func (r *unicodeStringResource) Create(_ context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan unicode_client.UnicodeStringModel

	diags := req.Config.Get(context.Background(), &plan)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Now Create the Conversionf
	res, c, err := r.unicodeClient.AddConversionToApplication(plan)

	if err != nil {
		resp.Diagnostics.AddError("Unable to Create Conversion", err.Error())
		return
	}

	diags = resp.State.Set(context.Background(), res)

	resp.Diagnostics.Append(diags...)

	var conv string
	for _, v := range c {
		conv += "" + v.Value
	}

	resp.Diagnostics.AddWarning("Conversion Created", conv)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *unicodeStringResource) Read(_ context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var plan unicode_client.UnicodeStringModel

	diags := req.State.Get(context.Background(), &plan)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Now Create the Conversion
	rr, err := r.unicodeClient.GetConversionFromApplication(plan)
	if err != nil {
		resp.Diagnostics.AddWarning("Conversion Not Found", err.Error())

		rr = &plan
		//return
	}

	diags = resp.State.Set(context.Background(), rr)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *unicodeStringResource) Update(_ context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Didn't Change

	// Not Implemented For Now ..
	var plan unicode_client.UnicodeStringModel

	diags := req.Config.Get(context.Background(), &plan)

	resp.Diagnostics.Append(diags...)

	req.State.Set(context.Background(), plan)
}

func (r *unicodeStringResource) Delete(_ context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//
	var plan unicode_client.UnicodeStringModel

	diags := req.State.Get(context.Background(), &plan)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Now Create the Conversion
	_, err := r.unicodeClient.RemoveConversionFromApplication(plan)

	if err != nil {
		resp.Diagnostics.AddError("Unable to Delete Conversion", err.Error())
		return
	}

	diags = resp.State.Set(context.Background(), plan)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
}
