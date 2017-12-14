package com.sebible.fluttergo;

import android.util.Log;
import java.nio.ByteBuffer;

import io.flutter.plugin.common.BasicMessageChannel;

import types.Replier;

public class GoReplier implements Replier {
    private BasicMessageChannel.Reply<ByteBuffer> r;

    public GoReplier(BasicMessageChannel.Reply<ByteBuffer> reply) {
        this.r = reply;
    }

    @Override
    public void reply(final byte[] p0) throws Exception {
        final ByteBuffer buffer = ByteBuffer.allocateDirect(p0.length);
        buffer.put(p0);
        this.r.reply(buffer);
    }
}
