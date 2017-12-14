package com.sebible.fluttergo;

import android.util.Log;
import io.flutter.plugin.common.PluginRegistry.Registrar;
import io.flutter.plugin.common.BinaryMessenger;


import plugin.Plugin;
import channel.Channel;

/**
 * GoPlugin
 */
public class GoPlugin {
  private static final String TAG = "GoPlugin";
  /**
   * Plugin registration.
   */
  public static void registerWith(Registrar registrar) {
    Plugin.bootstrap();

    final GoManager manager = new GoManager(registrar.messenger());
    try {
        Channel.init(manager);
    }
    catch (Exception e) {
      Log.e(TAG, "Channel init exception: ", e);
    }
  }
}
