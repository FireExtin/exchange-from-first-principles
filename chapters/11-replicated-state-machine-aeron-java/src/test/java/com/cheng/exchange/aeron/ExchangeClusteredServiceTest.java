package com.cheng.exchange.aeron;

import org.agrona.concurrent.UnsafeBuffer;
import org.junit.jupiter.api.Test;

import java.nio.ByteBuffer;
import java.util.concurrent.atomic.AtomicInteger;

import static org.junit.jupiter.api.Assertions.assertEquals;

final class ExchangeClusteredServiceTest {
    @Test
    void sessionMessageIsTheCommittedCommandBoundary() {
        final AtomicInteger calls = new AtomicInteger();
        final ExchangeClusteredService service = new ExchangeClusteredService(
            (session, timestamp, buffer, offset, length, header) -> {
                assertEquals(7, timestamp);
                assertEquals(1, offset);
                assertEquals(3, length);
                assertEquals(2, buffer.getByte(offset));
                calls.incrementAndGet();
            });

        final UnsafeBuffer buffer = new UnsafeBuffer(ByteBuffer.wrap(new byte[]{1, 2, 3, 4, 5}));
        service.onSessionMessage(null, 7, buffer, 1, 3, null);

        assertEquals(1, calls.get());
        assertEquals(1, service.committedMessages());
    }
}
