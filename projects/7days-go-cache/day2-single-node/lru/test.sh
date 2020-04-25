#!/bin/bash

go test -run TestGet -v
go test -run TestRemoveOldest -v
go test -run TestOnEvicted -v
