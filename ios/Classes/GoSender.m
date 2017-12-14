#import "GoSender.h"
#import "GoReplier.h"

@implementation GoSender {
  id<GoTypesChannel> _channel;
}
+ (instancetype)senderWithName:(NSString*)name
               binaryMessenger:(NSObject<FlutterBinaryMessenger>*)messenger {
  FlutterBasicMessageChannel* channel = [FlutterBasicMessageChannel
      messageChannelWithName:name
             binaryMessenger:messenger
                       codec:[FlutterBinaryCodec sharedInstance]];
  GoSender* instance = [[GoSender alloc] init];
  [channel setMessageHandler:^(id message, FlutterReply callback) {
    [instance handleOnMessage:message callback:callback];
  }];
  return instance;
}

- (BOOL)sendMessage:(NSData*)p0 p1:(id<GoTypesReplier>)p1 error:(NSError**)error {
  [_channel sendMessage:p0 p1:p1 error:error];
  return YES;
}

- (BOOL)setChannel:(id<GoTypesChannel>)p0 error:(NSError**)error {
  _channel = p0;
  return YES;
}

- (void)handleOnMessage:(id)message callback:(FlutterReply)callback {
  NSError *error;
  [self sendMessage:message
                 p1:[GoReplier replierWithCallback:callback]
              error:&error];
}

@end
