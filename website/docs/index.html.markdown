---
layout: "fortiwebcloud"
page_title: "FortiWeb Cloud: fortiwebcloud"
sidebar_current: "docs-fortiwebcloud-index"
description: |-
  The fortiwebcloud provider interacts with Fortiweb Cloud.
---

# fortiwebcloud

fortiwebcloud is used to interact with the resources supported by FortiWeb Cloud. Prior to use, the provider must be configured with the proper credentials.

## Configuration for FortiWeb Cloud

### Example Usage

```hcl
# Configure the FortiWeb Cloud Provider.
provider "fortiwebcloud" {
  hostname     = "api.fortiweb-cloud.com"
  api_token = "specify your token"
}
# Create Application.
resource "fortiwebcloud_app" "app_example" {
  app_name = "app-name"
  domain_name = "www.example.com"
  app_service = {
      http= 80
      https= 443
  }
  origin_server_ip = "1.1.1.1"
  origin_server_service = "HTTPS"
  cdn = false
}

```

### Argument Reference

The following arguments are supported:

* `hostname` - (Required) The API for FortiWeb Cloud(api.fortiweb-cloud.com).
* `username` - (Optional) The account of FortiWeb Cloud.
* `password` - (Optional) The password of your FortiWeb Cloud account.
* `api_token` - (Optional) The api key secret of your FortiWeb Cloud account.
