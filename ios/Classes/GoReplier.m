#import "GoReplier.h"
#import <Flutter/Flutter.h>

@implementation GoReplier {
  FlutterReply _callback;
}
+ (instancetype)replierWithCallback:(FlutterReply)callback {
  return [[GoReplier alloc] initWithCallback:callback];
}

- initWithCallback:(FlutterReply)callback {
  self = [super init];
  NSAssert(self, @"Super init cannot be nil");
  _callback = callback;
  return self;
}

- (BOOL)reply:(NSData*)p0 error:(NSError**)error {
  if (_callback)
    _callback(p0);
  return YES;
}

@end
