package descriptors

import (
	"github.com/caicloud/nirvana/definition"

	handler "github.com/caicloud/cyclone/pkg/server/handler/v1alpha1"
	httputil "github.com/caicloud/cyclone/pkg/util/http"
)

func init() {
	register(tenant...)
}

var tenant = []definition.Descriptor{
	{
		Path:        "/tenants",
		Description: "Tenant APIs",
		Definitions: []definition.Definition{
			{
				Method:      definition.Create,
				Function:    handler.CreateTenant,
				Description: "Create tenant",
				Parameters: []definition.Parameter{
					{
						Source:      definition.Body,
						Description: "tenant",
					},
				},
				Results: definition.DataErrorResults("tenant"),
			},
			{
				Method:      definition.Get,
				Function:    handler.ListTenants,
				Description: "List tenants",
				Results:     definition.DataErrorResults("tenants"),
			},
		},
	},
	{
		Path: "/tenants/{tenant}",
		Definitions: []definition.Definition{
			{
				Method:      definition.Get,
				Function:    handler.GetTenant,
				Description: "Get tenant",
				Parameters: []definition.Parameter{
					{
						Source: definition.Path,
						Name:   httputil.TenantNamePathParameterName,
					},
				},
				Results: definition.DataErrorResults("tenant"),
			},
			{
				Method:      definition.Update,
				Function:    handler.UpdateTenant,
				Description: "Update tenant",
				Parameters: []definition.Parameter{
					{
						Source: definition.Path,
						Name:   httputil.TenantNamePathParameterName,
					},
					{
						Source:      definition.Body,
						Description: "tenant",
					},
				},
				Results: definition.DataErrorResults("tenant"),
			},
			{
				Method:      definition.Delete,
				Function:    handler.DeleteTenant,
				Description: "Delete tenant",
				Parameters: []definition.Parameter{
					{
						Source: definition.Path,
						Name:   httputil.TenantNamePathParameterName,
					},
				},
				Results: []definition.Result{definition.ErrorResult()},
			},
		},
	},
}
