#!/bin/bash

HEROKU_EMAIL=${1}
HEROKU_API_KEY=${2}
VERSION=${3}
TERRAFORM_PG_BACKEND=${4}
NAME=${5:-staging}
DOCKERHUB_USERNAME=${6:-goveed20}
CONTAINER_NAME=${7:-terraform-deploy}

APP_NAME_AGENT_PRODUCTS=agent-products
APP_NAME_AGENT_INVOICES=agent-invoices
APP_NAME_AGENT_REPORTS=agent-reports

APP_IMAGE_NAME_AGENT_PRODUCTS=${DOCKERHUB_USERNAME}/${APP_NAME_AGENT_PRODUCTS}:${VERSION}
APP_IMAGE_NAME_AGENT_INVOICES=${DOCKERHUB_USERNAME}/${APP_NAME_AGENT_INVOICES}:${VERSION}
APP_IMAGE_NAME_AGENT_REPORTS=${DOCKERHUB_USERNAME}/${APP_NAME_AGENT_REPORTS}:${VERSION}


APP_NAME_AGENT_PRODUCTS_HEROKU=agentproducts${NAME}
APP_NAME_AGENT_INVOICES_HEROKU=agentinvoices${NAME}
APP_NAME_AGENT_REPORTS_HEROKU=agentreports${NAME}
# create can be used to create container from base image
# add new files, env variables
# and then in separete comand to start container

# here we use it for running terraform scripts to deploy on Heroku staging and prod env
# we override CMD but ENTRYPOINT probably does heroku login
DOCKER_BUILDKIT=1 docker create \
  --workdir /deployment \
  --entrypoint sh \
  --env APP_IMAGE_NAME_AGENT_PRODUCTS="${APP_IMAGE_NAME_AGENT_PRODUCTS}" \
  --env APP_IMAGE_NAME_AGENT_INVOICES="${APP_IMAGE_NAME_AGENT_INVOICES}" \
  --env APP_IMAGE_NAME_AGENT_REPORTS="${APP_IMAGE_NAME_AGENT_REPORTS}" \
  --env APP_NAME_AGENT_PRODUCTS="${APP_NAME_AGENT_PRODUCTS_HEROKU}" \
  --env APP_NAME_AGENT_INVOICES="${APP_NAME_AGENT_INVOICES_HEROKU}" \
  --env APP_NAME_AGENT_REPORTS="${APP_NAME_AGENT_REPORTS_HEROKU}" \
  --env HEROKU_API_KEY="${HEROKU_API_KEY}" \
  --env HEROKU_EMAIL="${HEROKU_EMAIL}" \
  --env TERRAFORM_PG_BACKEND="${TERRAFORM_PG_BACKEND}" \
  --name "$CONTAINER_NAME" \
  danijelradakovic/heroku-terraform \
  installBash.sh

docker cp deployment/. "${CONTAINER_NAME}":/deployment/
docker start -i "${CONTAINER_NAME}"
docker rm "${CONTAINER_NAME}"