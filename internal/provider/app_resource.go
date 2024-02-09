package provider

import (
	"context"
	"strings"

	unicode_client "terraform-provider-unicode/internal/unicode"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var (
	_ resource.Resource              = &UnicodeAppResource{}
	_ resource.ResourceWithConfigure = &UnicodeAppResource{}
)

func NewUnicodeAppResource() resource.Resource {
	return &UnicodeAppResource{}
}

type UnicodeAppResource struct {
	unicodeClient *unicode_client.UnicodeProviderClient
}

func (r *UnicodeAppResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app"
}

// What Types and Annotations Will I need if I want to create a resource that will create a unicode character

func (r *UnicodeAppResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	//resp.Schema = schema.Schema{}
	//resp.Schema = map[string]resource.Attribute{

	/*resp.Schema = tfsdk.Schema {
		Attributes: map[string]tfsdk.Attribute{
			"app_id": {
				Type: tfsdk.StringType,
				compute: true,
			},
			"name": {
				Type: tfsdk.StringType,
				Required: true,
			},
			"description": {
				Type: tfsdk.StringType,
				Required: true,
			},
			"created_at": {
				Type: tfsdk.StringType,
				Required: true,
			},
			"updated_at": {
				Type: tfsdk.StringType,
				Required: true,
			},
		},

	}*/

	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
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
func (r *UnicodeAppResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan unicode_client.UnicodeAppModel
	//resp.State = req.NewState
	//req.Config.Get(ctx, &plan)
	req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	s := req.Plan.Raw
	resp.Diagnostics.AddWarning("Client Resource Requested s", s.String())
	strings.Split(s.String(), "")

	//req.Plan.Set()

	resp.Diagnostics.AddWarning("Client Resource Requested Nem", plan.Name)

	//Get the Resource Client
	res, err := r.unicodeClient.CreateApplication(unicode_client.UnicodeAppModel{
		Name:        plan.Name,
		Description: plan.Description,
		Created_at:  plan.Created_at,
		Updated_at:  plan.Updated_at,
	})
	if err != nil {
		resp.Diagnostics.AddError("Unable to get Unicode Applications", err.Error())
		return
	}

	var fake_response unicode_client.UnicodeAppModel

	fake_response.Created_at = plan.Created_at
	fake_response.Description = plan.Description
	fake_response.Id = res.Id
	fake_response.Name = plan.Name
	fake_response.Updated_at = plan.Updated_at

	resp.Diagnostics.AddWarning("Client Resource ID", res.Id.String())
	resp.Diagnostics.AddWarning("Client Resource Name", fake_response.Name)

	diags := resp.State.Set(context.Background(), fake_response)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *UnicodeAppResource) Read(_ context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var request *unicode_client.UnicodeAppModel

	req.State.Get(context.Background(), &request)

	//
	//
	response, err := r.unicodeClient.GetApplication(request.Id.ValueString())

	if err != nil {
		//resp.Diagnostics.AddError("Unable to get Unicode Applications", err.Error())
		//return
		resp.Diagnostics.AddWarning("May Not Be Created Yet", err.Error())
		req.State.Get(context.Background(), response)
	}

	diags := resp.State.Set(context.Background(), response)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *UnicodeAppResource) Update(_ context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan unicode_client.UnicodeAppModel
	//resp.State = req.NewState
	req.Config.Get(context.Background(), &plan)
	//req.State.Get(context.Background(), &plan)

	var old_plan unicode_client.UnicodeAppModel

	req.State.Get(context.Background(), &old_plan)

	resp.Diagnostics.AddWarning("Client Resource", old_plan.Id.ValueString())

	//Delete Old Resource
	err := r.unicodeClient.DeleteApplication(old_plan.Id.ValueString()) //Assuming It keeps Same ID -> Stop Being a Silly Goose ...
	if err != nil {
		resp.Diagnostics.AddError("Unable to get Unicode Applications", err.Error())
		return
	}

	//Get the Resource Client
	res, err := r.unicodeClient.CreateApplication(unicode_client.UnicodeAppModel{
		Id:          plan.Id,
		Name:        plan.Name,
		Description: plan.Description,
		Created_at:  plan.Created_at,
		Updated_at:  plan.Updated_at,
	})
	if err != nil {
		resp.Diagnostics.AddError("Unable to get Unicode Applications", err.Error())
		return
	}

	diags := resp.State.Set(context.Background(), res)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *UnicodeAppResource) Delete(_ context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//Delete Application

	var plan unicode_client.UnicodeAppModel

	req.State.Get(context.Background(), &plan)

	err := r.unicodeClient.DeleteApplication(plan.Id.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Unable to get Unicode Applications", err.Error())
		return
	}

	// Return
	resp.State.Set(context.Background(), nil)

}

func (r *UnicodeAppResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	unicodeClient, ok := req.ProviderData.(*unicode_client.UnicodeProviderClient)

	if !ok {
		//resp.Diagnostics.AddError("Unable to create client in resource", "Client is NULL After NewUnicodeProviderClient")
		//return
		resp.Diagnostics.AddWarning("Unable to create client in resource", "Client is NULL After NewUnicodeProviderClient")
		unicodeClient = unicode_client.NewUnicodeProviderClient("bob")
	}

	r.unicodeClient = unicodeClient

	return
}
