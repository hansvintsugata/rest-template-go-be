#!/bin/bash
for file in ./api/http/*.yaml; do
        f=$(basename $file .yaml)
        mkdir -p ./internal/http/$f/port/genhttp

        oapi-codegen -generate types \
        -o ./internal/http/$f/port/genhttp/openapi_types.gen.go \
        -templates ./pkg/codegen/templates/ \
        -package genhttp \
        api/http/$f.yaml

        oapi-codegen -generate chi-server \
        -o ./internal/http/$f/port/genhttp/openapi_server.gen.go \
        -templates ./pkg/codegen/templates/ \
        -package genhttp \
        api/http/$f.yaml
done