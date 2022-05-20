#!/bin/bash -x

RELATIVE_PROJECT_DIR_PATH=..
PROJECT_DIR=$(cd $(dirname $0);cd $RELATIVE_PROJECT_DIR_PATH; pwd)

# init environment variables
cp -v ${PROJECT_DIR}/backend/.env.example ${PROJECT_DIR}/backend/.env
cp -v ${PROJECT_DIR}/frontend/.env.example ${PROJECT_DIR}/frontend/.env

# init mysql log files
mkdir -p ${PROJECT_DIR}/log/mysql
touch ${PROJECT_DIR}/log/mysql/mysqld.log

# init nginx log files
mkdir -p ${PROJECT_DIR}/log/nginx
touch ${PROJECT_DIR}/log/nginx/access.log
touch ${PROJECT_DIR}/log/nginx/error.log

# init log file permission
find ${PROJECT_DIR}/log -type f -print | xargs chmod 666
