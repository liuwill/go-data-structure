#!/usr/bin/env bash
testPath=$1

if [ ! -n "$testPath" ]; then
  echo "选择的目录不存在"
elif [ ! -d "$testPath" ];then
  echo "选择的目录不存在"
else
  cd $testPath && go test --cover
fi
