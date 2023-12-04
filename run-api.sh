#!/bin/bash
cd /home/atorres/Dev/Go/polaris/polaris-chatbot
export LIBRARY_PATH=$PWD 
export C_INCLUDE_PATH=$PWD
rm -fr cache
/usr/local/go1.21.0.linux-amd64/bin/go run cmd/ecatrom2000/main.go
#
