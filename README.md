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

#=> london_bridge = "London Bridge Is Winding Down, Winding down, jumping down"
output "london_bridge" {
  value = provider::multireplace::multireplace(
    "London Bridge Is Falling Down, Falling down, falling down",
    { Falling = "Winding", falling = "jumping" }
  )

  # value = provider::multireplace::multireplace(
  #   "London Bridge Is Falling Down, Falling down, falling down",
  # 	{ Falling = "Winding" },
  # 	{ falling = "jumping" }
  # )
}

#=> birmingham_bridge = "Birmingham Bridge Is rising up, rising up, rising up"
output "birmingham_bridge" {
  value = provider::multireplace::multireplace(
    "Birmingham Bridge Is Falling Down, Falling down, falling down",
    { "/(?i)falling/" = "rising", "/(?i)down/" = "up" }
  )
}

#=> html = <<-EOT
#       {"link":"<a href=\"https://example.com?foo=bar&zoo=baz\">Open</a>"}
#   EOT
output "html" {
  value = format(" %s\n", provider::multireplace::jsonunescape(
    # see https://developer.hashicorp.com/terraform/language/functions/jsonencode
    jsonencode({
      link = "<a href=\"https://example.com?foo=bar&zoo=baz\">Open</a>"
    })
  ))
}
```

## Run locally for development

```sh
cp multireplace.tf.sample multireplace.tf
make tf-plan
make tf-apply
# make tf-console
```
