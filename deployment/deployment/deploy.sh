#!/bin/sh

# CREATE TERRAFORM BACKEND ON HEROKU POSTGRES
# heroku create $TERRAFORM_PG_BACKEND
# heroku addons:create heroku-postgresql:hobby-dev --app $TERRAFORM_PG_BACKEND

ALL_HEROKU_APPS=$(heroku apps) && export ALL_HEROKU_APPS

echo $ALL_HEROKU_APPS
echo $TERRAFORM_PG_BACKEND

if [[ $ALL_HEROKU_APPS =~ "$TERRAFORM_PG_BACKEND" ]]; 
then
  echo "It's there!"
else 
  echo "It's no there!"
  heroku create $TERRAFORM_PG_BACKEND
  heroku addons:create jawsdb:kitefin --app $TERRAFORM_PG_BACKEND
fi


cd terraform || exit
JAWSDB_URL=$(heroku config:get JAWSDB_URL --app "$TERRAFORM_PG_BACKEND") && export JAWSDB_URL


# prepare Dockerfile-s

rm -rf ./agent-products/Dockerfile
rm -rf ./agent-invoices/Dockerfile
rm -rf ./agent-reports/Dockerfile


echo "FROM $APP_IMAGE_NAME_AGENT_PRODUCTS" >> ./agent-products/Dockerfile
cat ./agent-products/Dockerfile

echo "FROM $APP_IMAGE_NAME_AGENT_INVOICES" >> ./agent-invoices/Dockerfile
cat ./agent-invoices/Dockerfile

echo "FROM $APP_IMAGE_NAME_AGENT_REPORTS" >> ./agent-reports/Dockerfile
cat ./agent-reports/Dockerfile

terraform init -backend-config="conn_str=$JAWSDB_URL"

echo "Connected to database"

terraform apply -auto-approve -var agent-products-name="$APP_NAME_AGENT_PRODUCTS" \
                              -var agent-invoices-name="$APP_NAME_AGENT_INVOICES"   \
                              -var agent-reports-name="$APP_NAME_AGENT_REPORTS"