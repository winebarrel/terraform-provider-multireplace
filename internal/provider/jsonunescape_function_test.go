package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestJsonUnescapeFunction_OK(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::multireplace::jsonunescape(
							<<-EOT
								{ "link" : "\u003ca href=\"https://example.com?foo=bar\u0026zoo=baz\"\u003eOpen\u003c/a\u003e" }
							EOT
						)
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", `{ "link" : "<a href=\"https://example.com?foo=bar&zoo=baz\">Open</a>" }`+"\n"),
				),
			},
		},
	})
}

func TestJsonUnescapeFunction_OK_WithJsonencode(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::multireplace::jsonunescape(
							jsonencode({
								link = "<a href=\"https://example.com?foo=bar&zoo=baz\">Open</a>"
							})
						)
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", `{"link":"<a href=\"https://example.com?foo=bar&zoo=baz\">Open</a>"}`),
				),
			},
		},
	})
}
