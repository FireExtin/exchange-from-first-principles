package com.cheng.exchange.aeron;

import io.aeron.ExclusivePublication;

public final class ExchangeSnapshot {
    public long offer(final ExclusivePublication publication, final byte[] snapshotBytes) {
        throw new UnsupportedOperationException("Encode and publish exchange-core snapshot later.");
    }

    public void load(byte[] snapshotBytes) {
        throw new UnsupportedOperationException("Load exchange-core state later.");
    }
}
