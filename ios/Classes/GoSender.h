@import Plugin;
#import <Flutter/Flutter.h>

@interface GoSender : NSObject<GoTypesSender>
+ (instancetype)senderWithName:(NSString*)name
               binaryMessenger:(NSObject<FlutterBinaryMessenger>*)messenger;
@end
