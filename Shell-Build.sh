#!/bin/bash
# this script can be run
# under MSYS/Cygwin/WSL(with go and upx installed)

export GOARCH=386
export GOOS=windows
go.exe build -ldflags "-s -w" Shell.go
upx.exe Shell.exe
echo "# Build Successfully!"
ls -lh Shell.exe
echo "# Build under:"
uname -a
