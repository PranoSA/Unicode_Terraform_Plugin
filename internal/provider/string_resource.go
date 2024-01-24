package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type UnicodeStringModel struct {
	Id    string `json:"id" tfsdk:"id"`
	Name  string `json:"name" tfsdk:"name"`
	Index int    `json:"index" tfsdk:"index"`
	AppId string `json:"app_id" tfsdk:"app_id"`
}

func NewUnicodeStringResource() resource.Resource {
	return &unicodeStringResource{}
}

type unicodeStringResource struct{}

func (r *unicodeStringResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_unicode_string"
}

func (r *unicodeStringResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {

	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"app_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
			"index": schema.NumberAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Required: false,
				Optional: true,
			},
		},
	}
}

func (r *unicodeStringResource) Create(_ context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan UnicodeStringModel

	diags := req.Config.Get(context.Background(), &plan)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	rr := &UnicodeStringModel{
		Id:    plan.Id,
		Name:  plan.Name,
		Index: plan.Index,
		AppId: plan.AppId,
	}

	diags = resp.State.Set(context.Background(), rr)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *unicodeStringResource) Read(_ context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var plan UnicodeStringModel

	diags := req.State.Get(context.Background(), &plan)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	rr := &UnicodeStringModel{
		Id:    plan.Id,
		Name:  plan.Name,
		Index: plan.Index,
		AppId: plan.AppId,
	}

	diags = resp.State.Set(context.Background(), rr)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *unicodeStringResource) Update(_ context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Didn't Change
}

func (r *unicodeStringResource) Delete(_ context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//
}
