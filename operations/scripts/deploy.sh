#!/usr/bin/env bash
#
# usage: ./deploy.sh BRANCH IMAGE
#
# example: ./deploy.sh master master-1231ad
#
INSTANCE="ec2-35-158-19-66.eu-central-1.compute.amazonaws.com"
TRAVIS_BRANCH=$1
CURRENT=$2

conn()
{
    ssh -o "IdentitiesOnly yes" \
        -o "StrictHostKeyChecking no" \
        -o "User ec2-user" \
        -i "$(pwd)/operations/secrets/travis-aws" \
        $INSTANCE $1
}

cat docker-compose.testing.yml | conn "mkdir .deploy/; \
  cat > .deploy/docker-compose.yml; \
  TRAVIS_BRANCH=$TRAVIS_BRANCH CURRENT=$CURRENT docker-compose -p api-$TRAVIS_BRANCH -f .deploy/docker-compose.yml up -d; \
  rm -rf .deploy/"
