terraform {
  backend "pg" {
  }
}

provider "heroku" {
}

# Configure the MySQL provider
provider "mysql" {
  endpoint = "127.0.0.1:3306"
  username = "root"
  password = "root"
}

# Create a Database
resource "mysql_database" "mysql_db" {
  name = "mysql-database"
}

variable "agent-products-name" {
  description = "Unique name of the agent products app"
}
variable "agent-invoices-name" {
  description = "Unique name of the agent invoices app"
}
variable "agent-reports-name" {
  description = "Unique name of the agent reports app"
}


## ---------- AGENT-PRODUCTS ----------- ##
resource "heroku_app" "agent-products" {
  name   = var.agent-products-name
  region = "eu"
  stack  = "container"
}

resource "heroku_build" "agent-products" {
  app = heroku_app.agent-products.id

  source {
    path = "agent-products"
  }
}


## ---------- AGENT-INVOICES ----------- ##
resource "heroku_app" "agent-invoices" {
  name   = var.agent-invoices-name
  region = "eu"
  stack  = "container"
}

resource "heroku_build" "agent-invoices" {
  app = heroku_app.agent-invoices.id

  source {
    path = "agent-invoices"
  }
}


## ---------- AGENT-REPORTS ----------- ##
resource "heroku_app" "agent-reports" {
  name   = var.agent-reports-name
  region = "eu"
  stack  = "container"
}

resource "heroku_build" "agent-reports" {
  app = heroku_app.agent-reports.id

  source {
    path = "agent-reports"
  }
}


## ---------- OUTPUTS ----------- ##

output "agent-products-url" {
  value = "https://${heroku_app.agent-products.name}.herokuapp.com"
}
output "agent-invoices-url" {
  value = "https://${heroku_app.agent-invoices.name}.herokuapp.com"
}
output "agent-reports-url" {
  value = "https://${heroku_app.agent-reports.name}.herokuapp.com"
}