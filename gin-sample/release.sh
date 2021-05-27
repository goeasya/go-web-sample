#!/bin/bash

VERSION=0.1

docker build --network host -t gin-sample:${VERSION} .

