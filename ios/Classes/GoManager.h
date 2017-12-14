@import Plugin;
#import <Flutter/Flutter.h>

@interface GoManager : NSObject<GoTypesManager>
+ (instancetype)managerWithMessenger:(NSObject<FlutterBinaryMessenger>*)messenger;
@end
