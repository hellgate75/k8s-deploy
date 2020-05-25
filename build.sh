#!/bin/sh
if [ "" = "$GOPATH" ]; then
	GOPATH="$(realpath ~/go)"
fi
go build -buildmode=exe -o $GOPATH/bin/ github.com/hellgate75/k8s-deploy/cmd/k8srepo/...
