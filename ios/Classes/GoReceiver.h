@import Plugin;
#import <Flutter/Flutter.h>

@interface GoReceiver : NSObject<GoTypesReceiver>
+ (instancetype)receiverWithName:(NSString*)name
               binaryMessenger:(NSObject<FlutterBinaryMessenger>*)messenger;
@end
