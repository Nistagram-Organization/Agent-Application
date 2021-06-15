#!/bin/bash
docker-compose -f docker-compose.test.yml up -d

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

AGENT_PRODUCTS_TEST_EXIT_CODE=return_exit_code agent-products
AGENT_REPORTS_TEST_EXIT_CODE=return_exit_code agent-reports
AGENT_INVOICES_TEST_EXIT_CODE=return_exit_code agent-invoices

echo "agent-products tests returned $AGENT_PRODUCTS_TEST_EXIT_CODE"
echo "agent-reports tests returned $AGENT_REPORTS_TEST_EXIT_CODE"
echo "agent-invoices tests returned $AGENT_INVOICES_TEST_EXIT_CODE"

if ["$AGENT_PRODUCTS_TEST_EXIT_CODE" = 1]; then
	exit 1
fi

if ["$AGENT_INVOICES_TEST_EXIT_CODE" = 1]; then
	exit 1
fi

if ["$AGENT_REPORTS_TEST_EXIT_CODE" = 1]; then
	exit 1
fi