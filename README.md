[![Build Status](https://travis-ci.com/Badmuts/hsleiden-ipsenh-api.svg?token=F8VcDxDboKvhDwqC3zq8&branch=master)](https://travis-ci.com/Badmuts/hsleiden-ipsenh-api)
[![Docker Pulls](https://img.shields.io/docker/pulls/badmuts/hsleiden-ipsenh-api.svg?maxAge=3600)](https://hub.docker.com/r/badmuts/hsleiden-ipsenh-api/)
[![Docker Stars](https://img.shields.io/docker/stars/badmuts/hsleiden-ipsenh-api.svg?maxAge=3600)](https://hub.docker.com/r/badmuts/hsleiden-ipsenh-api/)
[![Download Insomia workspace](https://img.shields.io/badge/Download%20REST%20workspace-latest-6e60cc.svg)](https://drive.google.com/open?id=0B9S6iWoU5nj4cU11eW1CbklRLVk)

# hsleiden-ipsenh-api
Api for IoT platform created during IPSENH - groep 2

## Getting started
```sh
# Start via makefile
$ make start 
# Or start via docker-compose
$ docker-compose up -d
# Retrieve logs for api container
$ make logs
```

**Testing**
```sh
$ make suite
```
