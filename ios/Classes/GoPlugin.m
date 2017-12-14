#import "GoPlugin.h"
#import "GoManager.h"

@implementation GoPlugin
+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
  GoPluginBootstrap();
  GoManager* manager = [GoManager
      managerWithMessenger:[registrar messenger]];
  NSError* error;
  GoChannelInit(manager, &error);
}

@end
