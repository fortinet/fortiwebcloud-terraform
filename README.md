![](https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg#align=left&display=inline&height=144&margin=%5Bobject%20Object%5D&originHeight=72&originWidth=300&status=done&style=none&width=600)

# Terraform Provider

- Website: [https://www.terraform.io](https://www.terraform.io)

- [![](https://badges.gitter.im/hashicorp-terraform/Lobby.png#align=left&display=inline&height=20&margin=%5Bobject%20Object%5D&originHeight=20&originWidth=92&status=done&style=none&width=92)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x or higher
- [Go](https://golang.org/doc/install) 1.13.x (to build the provider plugin)

## Setup the Provider

1. Obtain the source code and extract it.
    ```sh
    $ tar xvzf terraform-provider-fortiwebcloud.tar.gz
    ```

2. Create the plugin directory and move the provider to it.
    ```sh
    $ mkdir ~/.terraform.d/plugins
    $ mv terraform-provider-fortiwebcloud/terraform-provider-fortiwebcloud ~/.terraform.d/plugins
    ```
## Using the Provider

If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it into your plugins directory,  run `terraform init` to initialize it.

```sh
$ terraform init
```
## License
[MIT](https://github.com/fortinet/terraform-secure-remote-access/blob/master/LICENSE)

## Support

Please contact your Fortinet representation for any comments, questions, considerations, and/or concerns.
