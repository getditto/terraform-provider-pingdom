package main

import (
	"github.com/getditto/terraform-provider-pingdom/pingdom"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pingdom.Provider,
	})
}
