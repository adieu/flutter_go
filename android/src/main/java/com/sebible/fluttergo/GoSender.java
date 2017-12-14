package com.sebible.fluttergo;

import android.util.Log;
import java.nio.ByteBuffer;

import io.flutter.plugin.common.BasicMessageChannel;
import io.flutter.plugin.common.BinaryMessenger;
import io.flutter.plugin.common.BinaryCodec;

import types.Sender;
import types.Replier;
import types.Channel;


public class GoSender implements Sender, BasicMessageChannel.MessageHandler<ByteBuffer> {
    private static final String TAG = "GoSender";

    private String name;
    private Channel channel;

    public GoSender(String name, BinaryMessenger messenger) {
        name = name;
        BasicMessageChannel channel = new BasicMessageChannel<ByteBuffer>(messenger, name, BinaryCodec.INSTANCE);
        channel.setMessageHandler(this);
    }

    @Override
    public void sendMessage(byte[] p0, Replier p1) throws Exception {
       this.channel.sendMessage(p0, p1);
    }

    @Override
    public void setChannel(Channel p0) throws Exception {
        this.channel = p0;
    }

    @Override
    public void onMessage(final ByteBuffer message, BasicMessageChannel.Reply<ByteBuffer> reply) {
	try {
	    this.sendMessage(message.array(), new GoReplier(reply));
	}
	catch (Exception e) {
            Log.e(TAG, "Send message exception: ", e);
	}
	return;
    }
}
