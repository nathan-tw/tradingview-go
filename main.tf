terraform {
  required_providers {
    heroku = {
      source  = "heroku/heroku"
      version = "~> 4.0"
    }
  }
}

provider "heroku" {
  email   = var.heroku_email
  api_key = var.heroku_api_key
}

resource "heroku_app" "app" {
  name   = "terraform-test-app-1"
  region = "us"

  config_vars = {
    FOOBAR = "baz"
  }

  buildpacks = [
    "heroku/go"
  ]
}

resource "heroku_build" "app" {
  app        = heroku_app.app.id
  buildpacks = ["https://github.com/heroku/heroku-buildpack-go"]


  source {
    # This app uses a community buildpack, set it in `buildpacks` above.
    path     = "."
  }
}

resource "heroku_formation" "app" {
  app        = heroku_app.app.id
  type       = "web"
  quantity   = 1
  size       = "Free"
  depends_on = [heroku_build.app]
}