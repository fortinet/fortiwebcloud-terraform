
provider "fortiwebcloud" {
  hostname     = "api.fortiweb-cloud.com"
  api_token = "specify your token"
}

resource "fortiwebcloud_app" "app_example" {
  app_name = "from_terraform"
  domain_name = "www.example.com"
  app_service = {
      http= 80
      https= 443
  }
  origin_server_ip = "your server ip"
  origin_server_service = "HTTPS"
  cdn = false
}
