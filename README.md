![](https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg#align=left&display=inline&height=144&margin=%5Bobject%20Object%5D&originHeight=72&originWidth=300&status=done&style=none&width=600)

# Terraform Provider

- Website: [https://www.terraform.io](https://www.terraform.io)

- [![](https://badges.gitter.im/hashicorp-terraform/Lobby.png#align=left&display=inline&height=20&margin=%5Bobject%20Object%5D&originHeight=20&originWidth=92&status=done&style=none&width=92)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13.x or higher
- [Go](https://golang.org/doc/install) 1.16.x (to build the provider plugin)

## Compatibility
This integration has been tested against Terraform version 1.0.6. Versions above this are expected to work but have not been tested.


## Setup the Provider

1. Obtain the specific version tag of source code and extract it.

    https://github.com/fortinet/fortiwebcloud-terraform/tags

    ex, for the version tag 1.0.2
    ```sh
    $ wget https://github.com/fortinet/fortiwebcloud-terraform/archive/refs/tags/1.0.2.tar.gz
    $ tar xvzf 1.0.2.tar.gz
    $ cd fortiwebcloud-terraform-1.0.2
    ```

2. Build the plugin
    ```sh
    go build
    ```
    the result is terraform-provider-fortiwebcloud

3. Create the plugin directory and move the provider to it.

    The plugin path is ~/.terraform.d/plugins/fortinet/terraform/fortiwebcloud/[VERSION]/[TARGET]/
    - VERSION is a string as the version of the tag
    - TARGET specifies a particular target platform
        - darwin_amd64 for Mac
        - linux_arm for linux
        - windows_amd64 for Windows


    Example in linux environment and the version is 1.0.2
    ```sh
    $ mkdir -p ~/.terraform.d/plugins/fortinet/terraform/fortiwebcloud/1.0.2/linux_amd64/
    $ mv terraform-provider-fortiwebcloud ~/.terraform.d/plugins/fortinet/terraform/fortiwebcloud/1.0.2/linux_amd64/
    ```
## Using the Provider

If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it into your plugins directory,  run `terraform init` to initialize it.

```sh
$ terraform init
```
## License
[MIT](https://github.com/fortinet/fortiwebcloud-terraform/blob/master/LICENSE)

## Support

Please contact your Fortinet representation for any comments, questions, considerations, and/or concerns.
