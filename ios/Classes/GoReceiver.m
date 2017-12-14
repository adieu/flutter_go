#import "GoReceiver.h"
#import "GoReplier.h"

@implementation GoReceiver {
  FlutterBasicMessageChannel* _flutterChannel;
  id<GoTypesChannel> _channel;
}
+ (instancetype)receiverWithName:(NSString*)name
               binaryMessenger:(NSObject<FlutterBinaryMessenger>*)messenger {
  FlutterBasicMessageChannel* channel = [FlutterBasicMessageChannel
      messageChannelWithName:name
             binaryMessenger:messenger
                       codec:[FlutterBinaryCodec sharedInstance]];
  GoReceiver* instance = [[GoReceiver alloc] initWithFlutterChannel:channel];
  return instance;
}

- initWithFlutterChannel:(FlutterBasicMessageChannel*)channel {
  self = [super init];
  NSAssert(self, @"Super init cannot be nil");
  _flutterChannel = channel;
  return self;
}

- (BOOL)onMessage:(NSData*)p0 p1:(id<GoTypesReplier>)p1 error:(NSError**)error {
  [_flutterChannel
      sendMessage:[[FlutterBinaryCodec sharedInstance] decode:p0]
      reply:^(id reply) {
        if (p1) {
  	  NSError* _error;
          [p1 reply:[[FlutterBinaryCodec sharedInstance] encode:reply] error:&_error];
	  if (_error != nil && error != nil) {
	    *error = _error;
	  }
        }
      }];
  return (*error == nil);
}

- (BOOL)setChannel:(id<GoTypesChannel>)p0 error:(NSError**)error {
  _channel = p0;
  return YES;
}

@end
