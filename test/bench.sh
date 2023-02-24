#!/bin/bash

# args.length > 1
if [ $# -gt 1 ] || [ $1 == "help" ]; then
    echo "Usage: bash ./bench.sh <module_name> - to benchmark single module"
    echo "Or:    bash ./bench.sh all           - to benchmark all modules"
    exit
fi

if [ $1 == "all" ]; then
  go test -bench=. -benchtime=10x -benchmem comment_test.go favorite_test.go feed_test.go message_test.go publish_test.go relation_test.go user_test.go main_test.go
else
  go test -bench=. -benchtime=10x -benchmem $1_test.go main_test.go
fi


