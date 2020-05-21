#!/bin/bash

go test -run=TestLog -v

# 因为默认情况下 go test 会运行单元测试，为了防止单元测试的输出影响我们
# 查看基准测试的结果，可以使用-run=匹配一个从来没有的单元测试方法，过滤
# 掉单元测试的输出，我们这里使用none，因为我们基本上不会创建这个名字的
# 单元测试方法
go test -bench=BenchmarkLog -benchtime=5s -run=none
