// https://github.com/apache/thrift/blob/master/tutorial/shared.thrift

/**
 * This Thrift file can be included by other Thrift files that want to share
 * these definitions.
 */

namespace go tutorial

struct SharedStruct {
  1: i32 key
  2: string value
}

service SharedService {
  SharedStruct getStruct(1: i32 key)
}