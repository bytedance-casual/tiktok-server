#!/bin/bash

# args.length > 1
if [ $# -gt 1 ] || [ $1 == "help" ]; then
    echo "Usage: bash ./test.sh <module_name> - to test single module"
    echo "Or:    bash ./test.sh all           - to test all modules"
    exit
fi

if [ $1 == "all" ]; then
  go test comment_test.go favorite_test.go feed_test.go message_test.go publish_test.go relation_test.go user_test.go main_test.go -v
else
  go test $1_test.go main_test.go -v
fi


