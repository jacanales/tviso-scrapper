#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
ROOT="${SCRIPT_DIR}/../.."
OUTPUT_FILE="wiremock-body-transformer.jar"
EXTENSION_URL="https://repo1.maven.org/maven2/com/opentable/wiremock-body-transformer/1.1.3/wiremock-body-transformer-1.1.3.jar"
EXTENSION_FOLDER="${ROOT}/var/wiremock/"

if ! [[ -d ${EXTENSION_FOLDER} ]]; then
    mkdir -p ${EXTENSION_FOLDER}
fi

if ! [[ -f ${EXTENSION_FOLDER}${OUTPUT_FILE} ]]; then
    (cd ${EXTENSION_FOLDER} && curl -o ${OUTPUT_FILE} ${EXTENSION_URL} --silent)
fi
