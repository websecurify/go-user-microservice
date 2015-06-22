#!/usr/bin/env bash

# ---
# ---
# ---

CSD=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

# ---
# ---
# ---

docker run \
	--rm \
	-v "${CSD}/src/v1:/go/src/v1"\
	-v "${CSD}/bin:/src" \
	-v "/var/run/docker.sock:/var/run/docker.sock" \
	"centurylink/golang-builder" \
	"websecurify/go-user-microservice"
	
# ---
# ---
# ---

rm "${CSD}/bin/go-user-microservice"

# ---
