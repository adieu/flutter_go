package com.sebible.fluttergo;

import android.util.Log;

import io.flutter.plugin.common.BinaryMessenger;

import types.Manager;
import types.Receiver;
import types.Sender;


public class GoManager implements Manager {
    private BinaryMessenger messenger;

    public GoManager(BinaryMessenger messenger) {
        this.messenger = messenger;
    }

    @Override
    public Receiver newReceiver(String p0) throws Exception {
        return new GoReceiver(p0, this.messenger);
    }

    @Override
    public Sender newSender(String p0) throws Exception {
        return new GoSender(p0, this.messenger);
    }
}
