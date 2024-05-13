#!/usr/bin/env bash

MERCURE_PUBLISHER_JWT_KEY='!ChangeMe!' MERCURE_SUBSCRIBER_JWT_KEY='!ChangeMe!' SERVER_NAME=:9000 ./bin/mercure-bin/mercure run &

./bin/openceptor-bin/openceptor.eu -c ./bin/openceptor-bin/config.yml &


