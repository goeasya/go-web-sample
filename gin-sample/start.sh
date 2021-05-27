#!/bin/bash

APP_IMAGE="gin-sample:0.1"
# Set the config type. One of: zookeeper|env
CONFIG_TYPE="env"

# zookeeper
# If CONFIG_TYPE is zookeeper, must set:
CONFIG_ZK_NODE=""
CONFIG_ZK_SERVER=""

# env
# If CONFIG_TYPE is env, must set
ENV_DB_TYPE=""
# format: USER:PASSWORD@tcp(IP:POET)/DATABASE
ENV_DB_COON_INFO=""
# jwt secret key
ENV_SECRET_KEY=""
# http listen address
ENV_API_ADDR=""
# Set the logging level. One of: debug|info|warn|error (default "info")
ENV_LOG_LEVEL=""


if [[ ${CONFIG_TYPE} == "env" ]]; then
    docker run --restart always --network host -d --name artiman \
    -e CONFIG_TYPE=${CONFIG_TYPE} \
    -e DB_TYPE=${ENV_DB_TYPE} \
    -e DB_CONN_INFO=${ENV_DB_COON_INFO} \
    -e SECRET_KEY=${ENV_SECRET_KEY} \
    -e API_ADDR=${ENV_API_ADDR} \
    -e LOGLEVEL=${ENV_LOG_LEVEL} \
    ${APP_IMAGE}
fi

if [[ ${CONFIG_TYPE} == "zookeeper" ]]; then
    docker run --restart always --network host -d --name artiman \
    -e CONFIG_TYPE=${CONFIG_TYPE} \
    -e CONFIG_ZK_NODE=${CONFIG_ZK_NODE} \
    -e CONFIG_ZK_SERVER=${CONFIG_ZK_SERVER} \
    ${APP_IMAGE}
fi

