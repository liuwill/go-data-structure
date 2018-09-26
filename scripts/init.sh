#!/bin/bash

check () {
  check_result=`which make`

  if [ ! "$check_result" ]; then
    return 0
  fi

  return 1
}

install () {
  cp /etc/apk/repositories /etc/apk/repositories.bak
  echo "http://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories
  apk add --no-cache make
}

check
CHECK_STATUS=$?

if [ "$CHECK_STATUS" = "0" ]; then
  echo "make is not installed, will install automatically"
  install
elif [ "$CHECK_STATUS" = "1" ]; then
  echo "detect make"
fi
