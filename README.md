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

output "birmingham_bridge" {
  value = provider::multireplace::multireplace(
    "Birmingham Bridge Is Falling Down, Falling down, falling down",
    { "/(?i)falling/" = "rising", "/(?i)down/" = "up" }
  )
}

output "html" {
  value = provider::multireplace::jsonunescape(
    jsonencode({
      link = "<a href=\"https://example.com?foo=bar&zoo=baz\">Open</a>"
    })
  )
}
```

## Run locally for development

```sh
cp multireplace.tf.sample multireplace.tf
make tf-plan
make tf-apply
# make tf-console
```
