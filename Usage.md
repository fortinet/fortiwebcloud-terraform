# Terraform: FortiWeb Cloud as a provider

FortiWeb Cloud's Terraform support provides customers with more ways to efficiently deploy, manage, and automate application security across multiple cloud environments.
By using Terraform, various IT infrastructure needs can be automated, thereby diminishing mistakes from repetitive manual configurations.

The Terraform FortiWeb Cloud provider can be used to automatically onboard or delete applications.

The following example demonstrates how to use the Terraform FortiWeb Cloud provider to perform simple configuration changes on FortiWeb Cloud. It requires the following:

1. FortiWeb Cloud: 20.2.d or later.
2. FortiWeb Cloud Provider: This example uses terraform-provider-fortiwebcloud 1.0.0.
3. Terraform: This example uses Terraform 0.12.26.

<a name="572d3f5e"></a>
###### To configure FortiWeb Cloud with Terraform Provider module support:

1. Download `terraform-provider-fortiwebcloud` to your local directory `~/.terraform.d/plugins`.

1. Create a new file with the .tf extension for configuring your FortiWeb Cloud:

    ```
    $touch main.tf
    $ ls
    main.tf
    ```

1. Edit the `main.tf` Terraform configuration file. In the example below, you will connect to the FortiWeb Cloud API gateway. The api_token entered must have write privileges on FortiWeb Cloud.

   ```
   # Configure the FortiwebCloud Provider
   provider "fortiwebcloud" {
     hostname   = "api.fortiweb-cloud.com"
     api_token = "specify your api key secret"
   }
   ```

1. Create the resources for onboarding your application. Specify your application name, domain name, server service and CDN preference. Enable CDN to increase access speed across multiple regions.

    ```
    resource "fortiwebcloud_app" "app_example" {
    app_name = "from_terraform"
    domain_name = "www.example.com"
    app_service = {
        http= 80
        https= 443
    }
    origin_server_ip = "93.184.216.34"
    origin_server_service = "HTTPS"
    cdn = false
    }
    ```

1. Save your Terraform configuration file.

1. In the terminal, enter _terraform init_ to initialize the working directory:

    ```
    $ terraform init
    Initializing the backend...
    Initializing provider plugins...
    Terraform has been successfully initialized!
    You may now begin working with Terraform. Try running "terraform plan" to see any changes that are required for your infrastructure. All Terraform commands should now work.
    If you ever set or change modules or backend configuration for Terraform, rerun this command to reinitialize your working directory. If you forget, other commands will detect it and remind you to do so if necessary.
    ```

1. Run _terraform -v_ to verify the version of the loaded provider module:

    ```
    $ terraform -v
    Terraform v0.12.26
    + provider.fortiwebcloud v1.0.0
    ```

1. Enter _terraform plan_ to parse the configuration file and read from the FortiWeb Cloud configurations to see what Terraform changes. The example below onboards an application to FortiWeb Cloud.

    ```
    $ terraform plan
    Refreshing Terraform state in-memory prior to plan...
    The refreshed state will be used to calculate this plan, but will not be
    persisted to local or remote state storage.
    \------------------------------------------------------------------------
    An execution plan has been generated and is shown below.
    Resource actions are indicated with the following symbols:
        \+ create
    Terraform will perform the following actions:
        \# fortiwebcloud_app.app_example will be created
        \+ resource "fortiwebcloud_app" "app_example" {
            \+ app_name              = "from_terraform"
            \+ app_service           = {
                \+ "http"  = 80
                \+ "https" = 443
            }
            \+ block                 = false
            \+ cdn                   = false
            \+ cname                 = (known after apply)
            \+ domain_name           = "www.example.com"
            \+ ep_id                 = (known after apply)
            \+ id                    = (known after apply)
            \+ origin_server_ip      = "93.184.216.34"
            \+ origin_server_port    = 443
            \+ origin_server_service = "HTTPS"
        }
    Plan: 1 to add, 0 to change, 0 to destroy.
    \------------------------------------------------------------------------
    Note: You didn't specify an "-out" parameter to save this plan, so Terraform
    can't guarantee that exactly these actions will be performed if
    "terraform apply" is subsequently run.
    ```

1. Enter _terraform apply_ to continue the configuration:

    ```
    $ terraform apply
    An execution plan has been generated and is shown below.
        Resource actions are indicated with the following symbols:
        \+ create
        Terraform will perform the following actions:
    \# fortiwebcloud_app.app_example will be created
    \+ resource "fortiwebcloud_app" "app_example" {
        \+ app_name              = "from_terraform"
        \+ app_service           = {
            \+ "http"  = 80
            \+ "https" = 443
            }
        \+ block                 = false
        \+ cdn                   = false
        \+ cname                 = (known after apply)
        \+ domain_name           = "www.example.com"
        \+ ep_id                 = (known after apply)
        \+ id                    = (known after apply)
        \+ origin_server_ip      = "93.184.216.34"
        \+ origin_server_port    = 443
        \+ origin_server_service = "HTTPS"
        }
    Plan: 1 to add, 0 to change, 0 to destroy.
    Do you want to perform these actions?
    Terraform will perform the actions described above.
    Only 'yes' will be accepted to approve.
    Enter a value: yes
    fortiwebcloud_app.app_example: Creating...
    fortiwebcloud_app.app_example: Creation complete after 4s [id=from_terraform]
    Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
    ```

The application is now onboarded on FortiWeb Cloud.

1. To delete the application,  enter _terraform destroy_ to delete the configuration from FortiWeb Cloud.```

    ```
    $ terraform destroy
    fortiwebcloud_app.app_example: Refreshing state... [id=from_terraform]
    An execution plan has been generated and is shown below.
    Resource actions are indicated with the following symbols:
        \- destroy
    Terraform will perform the following actions:
        \# fortiwebcloud_app.app_example will be destroyed
        \- resource "fortiwebcloud_app" "app_example" {
            \- app_name              = "from_terraform" -> null
            \- app_service           = {
                \- "http"  = 80
                \- "https" = 443
            } -> null
            \- block                 = false -> null
            \- cdn                   = false -> null
            \- domain_name           = "www.example.com" -> null
            \- id                    = "from_terraform" -> null
            \- origin_server_ip      = "93.184.216.34" -> null
            \- origin_server_port    = 443 -> null
            \- origin_server_service = "HTTPS" -> null
        }
    Plan: 0 to add, 0 to change, 1 to destroy.
    Do you really want to destroy all resources?
        Terraform will destroy all your managed infrastructure, as shown above.
        There is no undo. Only 'yes' will be accepted to confirm.
        Enter a value: yes
    fortiwebcloud_app.app_example: Destroying... [id=from_terraform]
    fortiwebcloud_app.app_example: Destruction complete after 3s
    Destroy complete! Resources: 1 destroyed.
    ```

1. At this time, the modify operation is not supported. The application can be modified using the GUI/API.
