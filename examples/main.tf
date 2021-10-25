terraform {
  required_providers {
    fortiwebcloud = {
      source  = "fortinet/terraform/fortiwebcloud"
      version = "1.0.2"
    }
  }
}

provider "fortiwebcloud" {
  hostname  = "api.fortiweb-cloud.com"
  api_token = "specify your token"
}

resource "fortiwebcloud_app" "app_example" {
  app_name    = "from_terraform"
  domain_name = "www.example.com"
  app_service = {
    http  = 80
    https = 443
  }
  origin_server_ip      = "your server ip"
  origin_server_service = "HTTPS"
  cdn                   = false
  continent_cdn         = false
}


resource "fortiwebcloud_openapi_validation" "openapi_validation_example" {
  app_name = "from_terraform"
  enable   = true
  action   = "alert_deny"
  validation_files = [
    "/path/openapi_validation_file.yaml",
    "/path/openapi_validation_file2.yaml"
  ]
  depends_on = [
    fortiwebcloud_app.app_example
  ]
}
