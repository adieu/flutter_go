#import "GoManager.h"
#import "GoChannel.h"

@implementation GoManager {
  NSObject<FlutterBinaryMessenger>* _messenger;
}
+ (instancetype)managerWithMessenger:(NSObject<FlutterBinaryMessenger>*)messenger {
  return [[GoManager alloc] initWithMessenger:messenger];
}

- initWithMessenger:(NSObject<FlutterBinaryMessenger>*)messenger {
  self = [super init];
  NSAssert(self, @"Super init cannot be nil");
  _messenger = messenger;
  return self;
}

//- (void)dealloc {
//  [super dealloc];
//}
//
- (id<GoTypesChannel>)newChannel:(NSString*)p0 error:(NSError**)error {
  return [GoChannel channelWithName:p0 binaryMessenger:_messenger];
}

@end
