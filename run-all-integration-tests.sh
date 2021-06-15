#!/bin/bash
docker compose -f docker-compose.test.yml up -d

is_finished() {
	service_name="$1"
	status="$(docker inspect "$service_name" --format='{{.State.Status}}')"

	if [ "$status" = "exited" ]; then
		return 0
	else
		return 1
	fi
}

return_exit_code() {
	service_name="$1"
	status="$(docker inspect "$service_name" --format='{{.State.ExitCode}}')"
	return status
}

while ! is_finished agent-products; do sleep 20; done
while ! is_finished agent-reports; do sleep 20; done
while ! is_finished agent-invoices; do sleep 20; done

return [ [ return_exit_code agent-products ] || [ return_exit_code agent-invoices ] || [ return_exit_code agent-reports ] ]