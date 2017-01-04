#!/bin/bash

releasePath="build/"

if [ ! -d "$releasePath" ]; then
  echo "create" + $releasePath
  mkdir -p $releasePath
fi

go build -o build/cron main.go
GOOS=windows GOARCH=386 go build -o build/cron.exe main.go
echo "buid done"
cp -R templates build/
echo "copy config"
cp config.yaml build/
echo "copy templates complete"
cp -R public build/
echo "copy static complete"