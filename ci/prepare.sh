#!/bin/bash
GODIR=$GOPATH/src/${CIRCLE_REPOSITORY_URL#https://}
mkdir -p "$(dirname "$GODIR")"
ln -sfv "$(pwd -P)" "$GODIR"
cd $GODIR
