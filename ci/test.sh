#!/bin/bash 
set -eu
PKG_LIST=$1

versions=( \
  10.2.5 \
  10.8.4 \
  11.0.2 \
)

for version in "${versions[@]}"; do
  pushd ci
  printf "Testing with GitLab $version\n\n"
  GITLAB_VERSION=${version} docker-compose up -d
  until curl -sSf http://localhost:10080/explore  > /dev/null
  do
    sleep 10
  done
  popd
  printf "Testing...\n\n"
  go test -v -covermode=count -coverprofile=profile.cov -timeout=1200s ${PKG_LIST}
  pushd ci
  printf "Teardown...\n\n"
  GITLAB_VERSION=${version} docker-compose down -v --rmi 'local'
  popd
done

