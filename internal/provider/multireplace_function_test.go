package provider_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestMultiReplacerFunction_OK(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::multireplace::multireplace(
							"London Bridge Is Falling Down, Falling down, falling down",
							{ Falling = "Winding", falling = "jumping" }
						)
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "London Bridge Is Winding Down, Winding down, jumping down"),
				),
			},
		},
	})
}

func TestMultiReplacerFunction_OK_NoReplace(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::multireplace::multireplace(
							"London Bridge Is Falling Down, Falling down, falling down",
							{ xFalling = "Winding", xfalling = "jumping" }
						)
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "London Bridge Is Falling Down, Falling down, falling down"),
				),
			},
		},
	})
}

func TestMultiReplacerFunction_OK_VarArgs(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::multireplace::multireplace(
							"London Bridge Is Falling Down, Falling down, falling down",
							{ Falling = "Winding" },
							{ falling = "jumping" }
						)
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "London Bridge Is Winding Down, Winding down, jumping down"),
				),
			},
		},
	})
}

func TestMultiReplacerFunction_OK_Regexp(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::multireplace::multireplace(
							"London Bridge Is Falling Down, Falling down, falling down",
							{ "/(?i)falling/" = "raising", "/(?i)down/" = "up" }
						)
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "London Bridge Is raising up, raising up, raising up"),
				),
			},
		},
	})
}

func TestMultiReplacerFunction_OK_VarArgs_Regexp(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::multireplace::multireplace(
							"London Bridge Is Falling Down, Falling down, falling down",
							{ "/(?i)falling/" = "raising" },
							{ "/(?i)down/" = "up" }
						)
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "London Bridge Is raising up, raising up, raising up"),
				),
			},
		},
	})
}

func TestMultiReplacerFunction_Err_InvalidRegexp(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::multireplace::multireplace(
							"blah blah blah",
							{ "/)/" = "N/A" }
						)
					}
				`,
				ExpectError: regexp.MustCompile(`(?m)error\s+parsing\s+regexp`),
			},
		},
	})
}

func TestMultiReplacerFunction_Err_NoVarArgs(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::multireplace::multireplace(
							"blah blah blah"
						)
					}
				`,
				ExpectError: regexp.MustCompile(`(?m)wrong\s+number\s+of\s+arguments`),
			},
		},
	})
}
