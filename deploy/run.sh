#!/bin/sh

# Exit immediately if a command exits with a non-zero status.
set -e

cleaner sbu-prepare

sap-btp-service-operator-migration -namespace ${POD_NAMESPACE} run

cleaner final-clean
