#!/usr/bin/env bash
RUN_NAME="ApiService"

mkdir -p output/bin
cp script/* output/
chmod +x output/run.sh

go build -o output/bin/${RUN_NAME}
