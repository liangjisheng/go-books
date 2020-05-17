#!/bin/bash

mkdir mocks
mockgen -destination mocks/AuthorRepository.go -source=repository.go
