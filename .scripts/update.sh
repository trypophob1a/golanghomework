#!/bin/zsh

for d in $(ls)
do
  if [[ $d == homework* ]]; then
    cd $d
    echo "Update deps in ${d}..."
    go mod tidy
    go get -t -u ./...
    cd ..
  fi
done
