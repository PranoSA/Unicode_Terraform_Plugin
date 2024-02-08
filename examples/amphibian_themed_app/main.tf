terraform {
  required_providers {
    unicode = {
      //source = "hashicorp.com/edu/hashicups"
      source = "hashicorp.com/edu/unicode"
    }
  }
}

provider "unicode" {
  // Configuration options
  user = var.user
}


resource "unicode_app" "amphibian_application" {
  name        = "Amphibian Application"
  description = "Amphibian App v2"
  id          = "1215902212-123091SDSA"
  updated_at  = "2024-02-07T00:00:00Z"
  created_at  = "2024-02-07T10:00:00Z"
}

resource "unicode_unicode_string" "frog_string" {
  app_id = unicode_app.amphibian_application.id
  value  = "🐸"
}

resource "unicode_unicode_string" "water_string" {
  app_id = unicode_app.amphibian_application.id
  value  = "💧"
}

resource "unicode_unicode_string" "lily_pad_string" {
  app_id = unicode_app.amphibian_application.id
  value  = "🪴"
}

resource "unicode_unicode_string" "tadpole_string" {
  app_id = unicode_app.amphibian_application.id
  value  = "🐸🥚"
}

resource "unicode_unicode_string" "amphibian_string" {
  app_id = unicode_app.amphibian_application.id
  value  = "🐸🦎"
}

