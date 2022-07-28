#!/bin/sh
start(){
  echo "编译 $1 server"
  go build -o ../bin/$1 cmd/$1/main.go

  echo "启动 $1 server"
#  nohup ../bin/$1 $2 >> /dev/null 2>&1 &
  nohup ../bin/$1 ../bin/conf/$2 >> ../bin/logs/$1/$3 2>&1 &
}

start gateway gateway.json gateway.log
start account account.json account.log
start game game1.json game1.log
start chat chat1.json chat1.log
#start game game2.json game2.log
#start chat chat2.json chat2.log
