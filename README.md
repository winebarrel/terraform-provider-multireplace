# terraform-provider-multireplace

[![CI](https://github.com/winebarrel/terraform-provider-multireplace/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/terraform-provider-multireplace/actions/workflows/ci.yml)
[![terraform docs](https://img.shields.io/badge/terraform-docs-%35835CC?logo=terraform)](https://registry.terraform.io/providers/winebarrel/multireplace/latest/docs)

Terraform provider for replacing multiple strings.

## Usage

```tf
terraform {
  required_providers {
    multireplace = {
      source = "winebarrel/multireplace"
    }
  }
}

#########################
# multireplace function
#########################

output "london_bridge" {
  value = provider::multireplace::multireplace(
    "London Bridge Is Falling Down, Falling down, falling down",
    { Falling = "Winding", falling = "jumping" }
  )
  #=> london_bridge = "London Bridge Is Winding Down, Winding down, jumping down"

  # value = provider::multireplace::multireplace(
  #   "London Bridge Is Falling Down, Falling down, falling down",
  # 	{ Falling = "Winding" },
  # 	{ falling = "jumping" }
  # )
}

output "birmingham_bridge" {
  value = provider::multireplace::multireplace(
    "Birmingham Bridge Is Falling Down, Falling down, falling down",
    { "/(?i)falling/" = "rising", "/(?i)down/" = "up" }
  )
  #=> birmingham_bridge = "Birmingham Bridge Is rising up, rising up, rising up"
}

#########################
# jsonunescape function
#########################

locals {
  # see https://developer.hashicorp.com/terraform/language/functions/jsonencode
  jsonencode_html = jsonencode({
    link = "<a href=\"https://example.com?foo=bar&zoo=baz\">Open</a>"
  })

  jsonunescape_html = provider::multireplace::jsonunescape(local.jsonencode_html)
}

output "html" {
  value = <<-EOT
    ${local.jsonencode_html}
    ---
    ${local.jsonunescape_html}
  EOT
  #=> html = <<-EOT
  #       {"link":"\u003ca href=\"https://example.com?foo=bar\u0026zoo=baz\"\u003eOpen\u003c/a\u003e"}
  #       ---
  #       {"link":"<a href=\"https://example.com?foo=bar&zoo=baz\">Open</a>"}
  #   EOT
}
```

## Run locally for development

```sh
cp multireplace.tf.sample multireplace.tf
make tf-plan
make tf-apply
# make tf-console
```
