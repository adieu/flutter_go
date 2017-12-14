#import "GoManager.h"
#import "GoSender.h"
#import "GoReceiver.h"

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

- (id<GoTypesReceiver>)newReceiver:(NSString*)p0 error:(NSError**)error {
  return [GoReceiver receiverWithName:p0 binaryMessenger:_messenger];
}

- (id<GoTypesSender>)newSender:(NSString*)p0 error:(NSError**)error {
  return [GoSender senderWithName:p0 binaryMessenger:_messenger];
}

@end
