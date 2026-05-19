package com.cheng.exchange.aeron;

import io.aeron.cluster.client.AeronCluster;
import org.agrona.DirectBuffer;

import java.util.Objects;

public final class ExchangeClient {
    private final AeronCluster cluster;

    public ExchangeClient(final AeronCluster cluster) {
        this.cluster = Objects.requireNonNull(cluster, "cluster");
    }

    public long offer(final DirectBuffer commandBuffer, final int offset, final int length) {
        return cluster.offer(commandBuffer, offset, length);
    }
}
