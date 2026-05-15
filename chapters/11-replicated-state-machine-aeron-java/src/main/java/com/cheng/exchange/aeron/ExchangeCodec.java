package com.cheng.exchange.aeron;

public final class ExchangeCodec {
    private ExchangeCodec() {
    }

    public static byte[] passThrough(byte[] bytes) {
        return bytes;
    }
}

