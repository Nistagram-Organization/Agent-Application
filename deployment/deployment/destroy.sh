#!/bin/sh

# DESTROY TERRAFORM BACKEND ON HEROKU POSTGRES
# heroku apps:destroy --app $TERRAFORM_PG_BACKEND

echo "Running destory.sh"

ALL_HEROKU_APPS=$(heroku apps) && export ALL_HEROKU_APPS

echo $ALL_HEROKU_APPS
echo $TERRAFORM_PG_BACKEND

if [[ $ALL_HEROKU_APPS =~ "$TERRAFORM_PG_BACKEND" ]]; 
then
  echo "It's there!"
else 
  echo "It's no there!"
  heroku create $TERRAFORM_PG_BACKEND
  heroku addons:create heroku-postgresql:hobby-dev --app $TERRAFORM_PG_BACKEND
fi


cd terraform || exit
DATABASE_URL=$(heroku config:get DATABASE_URL --app "$TERRAFORM_PG_BACKEND") && export DATABASE_URL
terraform init -backend-config="conn_str=$DATABASE_URL"

terraform destroy -auto-approve -var agent-products-name="$APP_NAME_AGENT_PRODUCTS" \
                                -var agent-invoices-name="$APP_NAME_AGENT_INVOICES"   \
                                -var agent-reports-name="$APP_NAME_AGENT_REPORTS"