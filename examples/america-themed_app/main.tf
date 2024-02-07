

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


resource "unicode_app" "america3" {
  name        = "USA Application"
  description = "America App v2"
  id          = "1215902212-123091SDSA"
  updated_at  = "2021-07-01T00:00:00Z"
  created_at  = "2021-07-01T00:00:00Z"
}

resource "unicode_unicode_string" "america_string" {
  app_id = unicode_app.america3.id
  value  = "🇺🇸"
}


resource "unicode_unicode_string" "football_string" {
  app_id = unicode_app.america3.id
  value  = "🏈"
}

resource "unicode_unicode_string" "baseball_string" {
  app_id = unicode_app.america3.id
  value  = "⚾"
}

resource "unicode_unicode_string" "fastfood_string" {
  app_id = unicode_app.america3.id
  value  = "🍔🥤🍟🌭"
}

