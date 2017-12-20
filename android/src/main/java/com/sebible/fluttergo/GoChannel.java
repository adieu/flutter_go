package com.sebible.fluttergo;

import android.util.Log;
import java.nio.ByteBuffer;

import io.flutter.plugin.common.BasicMessageChannel;
import io.flutter.plugin.common.BinaryMessenger;
import io.flutter.plugin.common.BinaryCodec;

import types.Channel;
import types.Replier;


public class GoChannel implements Channel, BasicMessageChannel.MessageHandler<ByteBuffer> {
    private static final String TAG = "GoChannel";

    private String name;
    private BasicMessageChannel _channel;

    public GoChannel(String name, BinaryMessenger messenger) {
        this.name = name;
        this._channel = new BasicMessageChannel<ByteBuffer>(messenger, name, BinaryCodec.INSTANCE);
        this._channel.setMessageHandler(this);
    }

    @Override
    public void send(byte[] message, final Replier p1) throws Exception {
       this._channel.send(message, new BasicMessageChannel.Reply<byte[]>() {
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

    @Override
    public void onMessage(final ByteBuffer message, BasicMessageChannel.Reply<ByteBuffer> reply) {
	try {
 	    Channel c = channel.Channel.connect(name);
	    c.send(message.array(), new GoReplier(reply));
	}
	catch (Exception e) {
            Log.e(TAG, "Channel message exception: ", e);
	}
	return;
    }
}
