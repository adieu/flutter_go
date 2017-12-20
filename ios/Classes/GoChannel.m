#import "GoChannel.h"
#import "GoReplier.h"

@implementation GoChannel {
  NSString* _name;
  FlutterBasicMessageChannel* _channel;
}
+ (instancetype)channelWithName:(NSString*)name
               binaryMessenger:(NSObject<FlutterBinaryMessenger>*)messenger {
  FlutterBasicMessageChannel* channel = [FlutterBasicMessageChannel
      messageChannelWithName:name
             binaryMessenger:messenger
                       codec:[FlutterBinaryCodec sharedInstance]];
  GoChannel* instance = [[GoChannel alloc] initWithName:name channel:channel];
  [channel setMessageHandler:^(id message, FlutterReply callback) {
    [instance handleOnMessage:message callback:callback];
  }];
  return instance;
}
- initWithName:(NSString*)name channel:(FlutterBasicMessageChannel*)channel {
  self = [super init];
  NSAssert(self, @"Super init cannot be nil");
  _name = name;
  _channel = channel;
  return self;
}

- (BOOL)send:(NSData*)p0 p1:(id<GoTypesReplier>)p1 error:(NSError**)error {
  [_channel sendMessage:p0 reply:^(NSData* message) {
    NSError *error;
    [p1 reply:message error:&error];
  }];
  return YES;
}

- (void)handleOnMessage:(id)message callback:(FlutterReply)callback {
  NSError *error;
  id<GoTypesChannel> channel = GoChannelConnect(_name, &error);
  [channel send:message
                 p1:[GoReplier replierWithCallback:callback]
              error:&error];
}

@end
