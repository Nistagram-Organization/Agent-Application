#!/bin/sh

echo "Running installBash.sh"

apk update
apk add bash
/bin/bash ./deploy.sh