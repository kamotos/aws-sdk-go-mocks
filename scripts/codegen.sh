#!/bin/bash
set -euxo pipefail

sdk_version=1.36.23

IFS=$'
'

pushd "$(go env GOPATH)/pkg/mod/github.com/aws/aws-sdk-go@v${sdk_version}/service"
interfaces=("$(grep '^type .* interface {$' */*/interface.go)")
popd

for t in ${interfaces[@]}; do
  interface_name=$(echo "$t" | awk '{print $2}')
  service_name=$(echo "$t" | awk -F '\\/' '{print $1}')
  destination_service=${service_name}mock
  mkdir -p service/${destination_service}
  mockgen \
    -package ${destination_service} \
    -destination service/${destination_service}/mock.go \
    github.com/aws/aws-sdk-go/service/$service_name/${service_name}iface \
    $interface_name &
done
