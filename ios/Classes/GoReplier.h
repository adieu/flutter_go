@import Plugin;
#import <Flutter/Flutter.h>

@interface GoReplier : NSObject<GoTypesReplier>
+ (instancetype)replierWithCallback:(FlutterReply)callback;
@end
