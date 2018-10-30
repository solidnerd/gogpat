#!/bin/bash 
set -eu
PKG_LIST=$1

versions=( \
  11.1.4 \
  11.2.3 \
  11.3.4 \
)

for version in "${versions[@]}"; do
  # pushd ci
  # printf "Testing with GitLab $version\n\n"
  # GITLAB_VERSION=${version} docker-compose up -d
  # until curl -sSf http://localhost:10080/explore  > /dev/null
  # do
  #   sleep 10
  # done
  # popd
  echo "$PKG_LIST"
  printf "Testing...\n\n"
  go test -v -covermode=count -coverprofile=profile.cov -timeout=1200s ${PKG_LIST}
  pushd ci
  printf "Teardown...\n\n"
  GITLAB_VERSION=${version} docker-compose down -v --rmi 'local'
  popd
done

