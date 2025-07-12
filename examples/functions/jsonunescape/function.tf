
output "html" {
  value = provider::multireplace::jsonunescape(
    jsonencode({
      link = "<a href=\"https://example.com?foo=bar&zoo=baz\">Open</a>"
    })
  )
}
