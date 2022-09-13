#!/usr/bin/env bash

# thrift 0.15.0
if [ ! -d thrift-gen ]; then
  mkdir thrift-gen # if necessary
fi
thrift -r -gen go -out thrift-gen IDL/calculator.thrift
