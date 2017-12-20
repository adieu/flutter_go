package com.sebible.fluttergo;

import android.util.Log;
import java.nio.ByteBuffer;

import io.flutter.plugin.common.BasicMessageChannel;
import io.flutter.plugin.common.BinaryMessenger;
import io.flutter.plugin.common.BinaryCodec;

import types.Channel;
import channel.Channel;


public class GoChannel implements Channel, BasicMessageChannel.MessageHandler<ByteBuffer> {
    private static final String TAG = "GoChannel";

    private String name;
    private BasicMessageChannel channel;

    public GoChannel(String name, BinaryMessenger messenger) {
        name = name;
        channel = new BasicMessageChannel<ByteBuffer>(messenger, name, BinaryCodec.INSTANCE);
        channel.setMessageHandler(this);
    }

    @Override
    public void sendMessage(byte[] p0, Replier p1) throws Exception {
       this.channel.sendMessage(p0, p1);
    }

    @Override
    public void onMessage(final ByteBuffer message, BasicMessageChannel.Reply<ByteBuffer> reply) {
	try {
        Channel channel = channel.Channel.connect(name);
	    channel.sendMessage(message.array(), new GoReplier(reply));
	}
	catch (Exception e) {
            Log.e(TAG, "Channel message exception: ", e);
	}
	return;
    }
}
