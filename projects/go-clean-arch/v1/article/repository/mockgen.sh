#!/bin/bash

mkdir mocks
mockgen -destination mocks/ArticleRepository.go -source=repository.go
