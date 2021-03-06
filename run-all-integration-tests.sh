#!/bin/bash
docker-compose -f docker-compose.test.yml up -d

is_finished() {
	local service_name="$1"
	local status="$(docker inspect -f "{{.State.Status}}" "$service_name")"

	echo "STATUS: $status, CONTAINER: $service_name"

	if [ "$status" = "exited" ]; then
		return 0
	else
		return 1
	fi
}

return_exit_code() {
	local service_name="$1"
	local exit_code="$(docker inspect -f "{{.State.ExitCode}}" "$service_name")"
	
	echo "EXIT_CODE: $exit_code, CONTAINER: $service_name"

	if [ "$exit_code" = 0 ]; then
		return 0
	else
		return 1
	fi
}

while ! is_finished agent-products; do sleep 20; done
while ! is_finished agent-reports; do sleep 20; done
while ! is_finished agent-invoices; do sleep 20; done

AGENT_PRODUCTS_TEST_EXIT_CODE="$(return_exit_code agent-products)"
AGENT_REPORTS_TEST_EXIT_CODE="$(return_exit_code agent-reports)"
AGENT_INVOICES_TEST_EXIT_CODE="$(return_exit_code agent-invoices)"

echo "agent-products tests returned $AGENT_PRODUCTS_TEST_EXIT_CODE"
echo "agent-reports tests returned $AGENT_REPORTS_TEST_EXIT_CODE"
echo "agent-invoices tests returned $AGENT_INVOICES_TEST_EXIT_CODE"

if [ "$AGENT_PRODUCTS_TEST_EXIT_CODE" -eq 1 ]; then
	echo "::set-output name=tests_exit_code::$($AGENT_PRODUCTS_TEST_EXIT_CODE)"
fi

if [ "$AGENT_INVOICES_TEST_EXIT_CODE" -eq 1 ]; then
	echo "::set-output name=tests_exit_code::$($AGENT_INVOICES_TEST_EXIT_CODE)"
fi

if [ "$AGENT_REPORTS_TEST_EXIT_CODE" -eq 1 ]; then
	echo "::set-output name=tests_exit_code::$($AGENT_REPORTS_TEST_EXIT_CODE)"
fi

echo "::set-output name=tests_exit_code::0"