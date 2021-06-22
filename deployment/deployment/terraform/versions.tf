terraform {
  required_providers {
    heroku = {
      source = "heroku/heroku"
    }
    mysql = {
      source  = "winebarrel/mysql"
      version = "~> 1.10.2"
    }
  }
  required_version = ">= 0.13"
}
