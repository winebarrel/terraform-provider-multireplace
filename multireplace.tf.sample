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
