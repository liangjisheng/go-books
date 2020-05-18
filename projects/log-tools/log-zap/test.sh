#!/bin/bash

rm -rf logs
go test -run TestInitZapV2Logger -v
