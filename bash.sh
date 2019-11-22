#!/usr/bin/env bash
name=$1
type=$2

#测试单个函数 ./bash.sh 函数名
if [[ -n $name ]];then
    go test -v -test.run $name
#基准测试单个函数 ./bash 函数名 b
elif [[ $type = "b" ]];then
    go test -bench $name -run=^$ -benchmem
else
    go test -v
fi