package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.ProviderWithFunctions = &MultiReplaceProvider{}

type MultiReplaceProvider struct {
	version string
}

type MultiReplaceProviderModel struct {
}

func (p *MultiReplaceProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "multireplace"
	resp.Version = p.version
}

func (p *MultiReplaceProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *MultiReplaceProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data MultiReplaceProvider

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.DataSourceData = data
	resp.ResourceData = data
}

func (p *MultiReplaceProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		// No Resources
	}
}

func (p *MultiReplaceProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// No DataSources
	}
}

func (p *MultiReplaceProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewMultiRepaceFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &MultiReplaceProvider{
			version: version,
		}
	}
}
