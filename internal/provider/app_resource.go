package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	_ resource.Resource = &UnicodeAppResource{}
)

func NewUnicodeAppResource() resource.Resource {
	return &UnicodeAppResource{}
}

type UnicodeAppResource struct {
}

func (r *UnicodeAppResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app"
}

// What Types and Annotations Will I need if I want to create a resource that will create a unicode character
type UnicodeAppModel struct {
	Id          string `json:"id" tfsdk:"id"`
	Name        string `json:"name" tfsdk:"name"`
	Description string `json:"description" tfsdk:"description"`
	Created_at  string `json:"created_at" tfsdk:"created_at"`
	Updated_at  string `json:"updated_at" tfsdk:"updated_at"`
}

func (r *UnicodeAppResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	//resp.Schema = schema.Schema{}
	//resp.Schema = map[string]resource.Attribute{
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Required: true,
			},
			"updated_at": schema.StringAttribute{
				Required: true,
			},
		},
	}

}

// Create a New Resource
func (r *UnicodeAppResource) Create(_ context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	//resp.State = req.Plan
	var plan UnicodeAppModel
	diags := req.Config.Get(context.Background(), &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	rt := &UnicodeAppModel{
		Id:          plan.Id,
		Name:        plan.Name,
		Description: plan.Description,
		Created_at:  plan.Created_at,
		Updated_at:  plan.Updated_at,
	}

	diags = resp.State.Set(context.Background(), rt)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *UnicodeAppResource) Read(_ context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.State = req.State
}

func (r *UnicodeAppResource) Update(_ context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//resp.State = req.NewState
	resp.State = req.State
}

func (r *UnicodeAppResource) Delete(_ context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State = req.State
}
