#!/bin/bash

protoc \
	--ts_opt=esModuleInterop=true \
	--ts_out="frontend/src/generated" \
	--go_out="generated/protos" \
	protos/*.proto

