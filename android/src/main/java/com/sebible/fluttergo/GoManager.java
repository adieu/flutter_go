package com.sebible.fluttergo;

import android.util.Log;

import io.flutter.plugin.common.BinaryMessenger;

import types.Manager;
import types.Channel;


public class GoManager implements Manager {
    private BinaryMessenger messenger;

    public GoManager(BinaryMessenger messenger) {
        this.messenger = messenger;
    }

    @Override
    public Channel newChannel(String p0) throws Exception {
        return new GoChannel(p0, this.messenger);
    }
}
