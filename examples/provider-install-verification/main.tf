

terraform {
  required_providers {
    unicode = {
      //source = "hashicorp.com/edu/hashicups"
      source = "hashicorp.com/edu/unicode"
    }
  }
}

provider "unicode" {
  user = "tray"
}

data "unicode_unicode_chars" "example" {
  unicode_char = "🦔"
}

output "unicode_char" {
  value = data.unicode_unicode_chars.example
}

/*data "hashicups_unicode_character" "char" {
  unicode_char = "🦔"
}
*/

resource "unicode_app" "example_app" {
  id          = "examplrrrre"
  name        = "example app234"
  description = "example"
  updated_at  = "2021-07-01T00:00:00Z"
  created_at  = "2021-07-01T00:00:00Z"
}

output "name" {
  value = resource.unicode_app.example_app
}

resource "unicode_unicode_string" "my_string" {
  app_id = unicode_app.example_app.id
  name   = "Hello, World!"
  index  = 0
  id     = "example"
}

output "my_gamer" {
  value = resource.unicode_unicode_string.my_string
}
