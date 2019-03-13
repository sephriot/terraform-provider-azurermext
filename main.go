package main

import (
"github.com/hashicorp/terraform/plugin"
"github.com/hashicorp/terraform/terraform"
"github.com/sephriot/terraform-provider-azurermext/azurermext"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return azurermext.Provider()
		},
	})
}