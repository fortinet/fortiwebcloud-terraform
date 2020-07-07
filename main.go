package main

import (
	"terraform-provider-fortiwebcloud/fortiwebcloud"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fortiwebcloud.Provider})
}
