package com.sebible.fluttergo;

import android.util.Log;

import io.flutter.plugin.common.BasicMessageChannel;
import io.flutter.plugin.common.BinaryMessenger;
import io.flutter.plugin.common.BinaryCodec;

import types.Receiver;
import types.Replier;
import types.Channel;


public class GoReceiver implements Receiver {
    private static final String TAG = "GoReceiver";

    BasicMessageChannel flutterChannel;
    private String name;
    private Channel channel;

    public GoReceiver(String name, BinaryMessenger messenger) {
        name = name;
        flutterChannel = new BasicMessageChannel<>(messenger, name, BinaryCodec.INSTANCE);
    }

    public void onMessage(byte[] p0, final Replier p1) throws Exception {
       this.flutterChannel.send(p0, new BasicMessageChannel.Reply<byte[]>() {
            public void reply(byte[] r) {
                try {
                    p1.reply(r);
                }
                catch (Exception e) {
                    Log.e(TAG, "Reply message exception: ", e);
                }
                return;
            }
       });
    }

    public void setChannel(Channel p0) throws Exception {
        this.channel = p0;
    }
}
