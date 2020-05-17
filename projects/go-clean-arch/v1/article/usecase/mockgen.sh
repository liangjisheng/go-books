#!/bin/bash

mkdir mocks
mockgen -destination mocks/ArticleUsercase.go v1/article/usecase ArticleUsecase
