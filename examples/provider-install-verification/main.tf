

terraform {
  required_providers {
    unicode = {
      //source = "hashicorp.com/edu/hashicups"
      source = "hashicorp.com/edu/unicode"
    }
  }
}

provider "unicode" {
  //user = "tray"
  user = var.user
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
  id          = "example14"
  name        = "example app234 for bob the builder the 2nd"
  description = "example"
  updated_at  = "2021-07-01T00:00:00Z"
  created_at  = "2021-07-01T00:00:00Z"
}

/*resource "unicode_unicode_string" "my_string" {
  app_id = unicode_app.example_app.id
  value  = "SEAN"
}*/

/*output "name" {
  value = resource.unicode_app.example_app
}*/

/*resource "unicode_unicode_string" "my_string" {
  app_id = unicode_app.example_app.id
  value  = "Horse"
}
*/
/*resource "unicode_unicode_string" "my_string" {
  app_id = unicode_app.example_app.id
  value  = "H!"
}*/

/*resource "unicode_unicode_string" "my_string2" {
  app_id = unicode_app.example_app.id
  value  = "Hello, World!"
}

output "my_string2" {
  value = resource.unicode_unicode_string.my_string2
}

*/

/*resource "unicode_unicode_string" "my_string" {
  app_id = unicode_app.example_app.id
  value  = "!"
}*/

/*resource "unicode_unicode_string" "my_string2" {
  app_id = unicode_unicode_string.my_string.app_id
  value  = "Hello, BOBS WORLD!"
}

resource "unicode_unicode_string" "my_string3" {
  app_id = unicode_unicode_string.my_string2.app_id
  value  = "Hello, TRUCE OF THE MATTER!"
}


output "my_string2" {
  value = resource.unicode_unicode_string.my_string2
}
*/

/*output "my_string" {
  value = resource.unicode_unicode_string.my_string
}*/

/*resource "unicode_unicode_string" "my_string" {
  app_id = unicode_app.example_app.id
  name   = "Hello, World!"
  index  = 0
  id     = "example"
}

*/

/*output "my_gamer" {
  value = resource.unicode_unicode_string.my_string
}
*/

// Get Resources and Print Them
