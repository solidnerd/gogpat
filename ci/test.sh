#!/bin/bash 
set -eu
PKG_LIST=$1
pushd ci
docker-compose up -d

until curl -sSf http://localhost:10080/users/sign_in  > /dev/null
do
  sleep 10
done
popd 
go test -v ${PKG_LIST}
pushd ci
docker-compose down -v
popd
