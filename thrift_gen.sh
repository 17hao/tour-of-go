#!/usr/bin/env bash

# thrift 0.13.0
if [ ! -d thrift-gen ]; then
  mkdir thrift-gen
fi

thrift -r -gen go:package_prefix=github.com/17hao/tour-of-go/thrift-gen/ -out thrift-gen IDL/calculator.thrift
