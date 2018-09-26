#!/bin/bash

DEMO_PATH=`pwd`
GOLANG_DOCKER_VERSION=golang:alpine
DOCKER_CONTAINER=go-data-structure
APP_NAME=go-data-structure
DOCKER_RUN_PATH=/go/src/$APP_NAME

usage()
{
  echo "USAGE: $CALLER [-h] COMMANDS"
  echo "       $CALLER [ --help ]"
  echo "COMMANDS:"
  echo "    install       鍒涘缓docker瀹瑰櫒"
  echo "    start         鍚姩docker瀹瑰櫒"
  echo "    test          杩愯鍗曞厓娴嬭瘯"
  echo "    enter         杩涘叆docker瀹瑰櫒"
  echo "Run '$CALLER COMMAND --help' for more information on a command."
  exit 1
}

installDocker () {
  echo "start install docker container"

  docker run -itd \
    -v $DEMO_PATH:$DOCKER_RUN_PATH \
    --name $DOCKER_CONTAINER \
    $GOLANG_DOCKER_VERSION &> /dev/null

  # docker run -itd \
  #   -v $DEMO_PATH:/go/src/go-kata-practice \
  #   --name golang-kata \
  #   golang:1.10 &> /dev/null
}

startDocker () {
  docker start $DOCKER_CONTAINER
}

enterDocker () {
  docker exec -it $DOCKER_CONTAINER /bin/sh
}

execTestKata () {
  docker exec -it $DOCKER_CONTAINER sh $DOCKER_RUN_PATH/scripts/init.sh
  # docker exec -it $DOCKER_CONTAINER sh -c 'cp /etc/apk/repositories /etc/apk/repositories.bak;echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories;apk add --no-cache make'
  docker exec -it $DOCKER_CONTAINER sh -c "cd $DOCKER_RUN_PATH;make coverhtml"
  # docker exec -it $DOCKER_CONTAINER sh -c 'apk del git'
}

checkDockerContainerStatus () {
  check_result=`docker ps -a | grep $DOCKER_CONTAINER`

  if [ ! "$check_result" ]; then
    return 0
  fi

  check_exit=`echo $check_result | grep Exited`
  if [ ! "$check_exit" ]; then
    return 2
  fi
  return 1
}

checkDockerInstalled () {
  check_result=`which docker`

  if [ ! "$check_result" ]; then
    echo "娌℃湁妫€娴嬪埌docker锛岃鍏堝畨瑁卍ocker涔嬪悗鍐嶈瘯"
    exit 1
  fi
}

if [ $# -ne 1 ] ; then
    usage
fi

checkDockerInstalled

checkDockerContainerStatus
DOCKER_START_STATUS=$?

if [ "$1" = "install" ]; then
  if [ "$DOCKER_START_STATUS" != "0" ]; then
    echo "瀹瑰櫒宸茬粡瀛樺湪"
    exit 0
  fi
  installDocker
  exit 0
fi


if [ "$1" = "start" ]; then
  if [ "$DOCKER_START_STATUS" = "2" ]; then
    echo "container is working"
    exit 0
  elif [ "$DOCKER_START_STATUS" = "0" ]; then
    echo "container not exist"
    installDocker
  fi
  startDocker
elif [ "$1" = "enter" ]; then
  if [ "$DOCKER_START_STATUS" = "0" ]; then
    echo "container not exist"
    installDocker
  elif [ "$DOCKER_START_STATUS" != "2" ]; then
    startDocker
  fi

  enterDocker
elif [ "$1" = "test" ]; then
  if [ "$DOCKER_START_STATUS" = "0" ]; then
    echo "container not exist"
    installDocker
  elif [ "$DOCKER_START_STATUS" != "2" ]; then
    startDocker
  fi

  execTestKata
fi

