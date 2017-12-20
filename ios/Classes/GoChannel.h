@import Plugin;
#import <Flutter/Flutter.h>

@interface GoChannel : NSObject<GoTypesChannel>
+ (instancetype)channelWithName:(NSString*)name
               binaryMessenger:(NSObject<FlutterBinaryMessenger>*)messenger;
- initWithName:(NSString*)name channel:(FlutterBasicMessageChannel*)channel;
@end
