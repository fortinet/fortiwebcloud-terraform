---
layout: "fortiwebcloud"
page_title: "FortiWeb Cloud: fortiwebcloud_app"
sidebar_current: "docs-fortiwebcloud-app"
description: |-
  This resource supports application creation.
---

# fortiwebcloud_app
This resource supports application creation.

## Example Usage
```hcl
resource "fortiwebcloud_app" "example" {
  app_name = "test-app"
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
## Argument Reference
The following arguments are supported:

* `app_name` - (Required) Application name.
* `domain_name` - (Required) Your domain (for example , www.example.com).
* `extra_domains` - (Optional) A list of extra domains (for example , www.example1.com).
* `app_service` - (Optional) Application service .
* `origin_server_ip` - (Required) Origin server IP or domain.
* `origin_server_service` - (Optional) Origin server service  or domain. Defaults to HTTPS.
* `origin_server_port` - (Optional) Origin server port. Defaults to 443.
* `cdn` - (Optional) Enable CDN or not. Defaults to false.
* `block` - (Optional) Enable block_mode or not. Defaults to false.
* `template` - (Optional) The template name.

## Attributes Reference
The following attributes are exported:
* `ep_id` - Application id.
* `cname` - CNAME of the application.
