package com.cheng.exchange.aeron;

import io.aeron.ExclusivePublication;
import io.aeron.Image;
import io.aeron.cluster.codecs.CloseReason;
import io.aeron.cluster.service.ClientSession;
import io.aeron.cluster.service.Cluster;
import io.aeron.cluster.service.ClusteredService;
import io.aeron.logbuffer.Header;
import org.agrona.DirectBuffer;

import java.util.Objects;

public final class ExchangeClusteredService implements ClusteredService {
    public interface CommandHandler {
        void onCommand(
            ClientSession session,
            long timestamp,
            DirectBuffer buffer,
            int offset,
            int length,
            Header header);
    }

    private final CommandHandler commandHandler;
    private Cluster cluster;
    private Cluster.Role role;
    private int committedMessages;

    public ExchangeClusteredService() {
        this((session, timestamp, buffer, offset, length, header) -> {
        });
    }

    public ExchangeClusteredService(final CommandHandler commandHandler) {
        this.commandHandler = Objects.requireNonNull(commandHandler, "commandHandler");
    }

    @Override
    public void onStart(final Cluster cluster, final Image snapshotImage) {
        this.cluster = Objects.requireNonNull(cluster, "cluster");
        // TODO: load snapshotImage into exchange-core once snapshot format exists.
    }

    @Override
    public void onSessionOpen(final ClientSession session, final long timestamp) {
        // TODO: track client sessions if the exchange needs per-session egress.
    }

    @Override
    public void onSessionClose(final ClientSession session, final long timestamp, final CloseReason closeReason) {
        // TODO: release any session-local resources once they exist.
    }

    @Override
    public void onSessionMessage(
        final ClientSession session,
        final long timestamp,
        final DirectBuffer buffer,
        final int offset,
        final int length,
        final Header header) {
        committedMessages++;
        commandHandler.onCommand(session, timestamp, buffer, offset, length, header);
    }

    @Override
    public void onTimerEvent(final long correlationId, final long timestamp) {
        // TODO: route cluster timers to the deterministic core if needed.
    }

    @Override
    public void onTakeSnapshot(final ExclusivePublication snapshotPublication) {
        // TODO: publish exchange-core snapshot bytes.
    }

    @Override
    public void onRoleChange(final Cluster.Role newRole) {
        this.role = Objects.requireNonNull(newRole, "newRole");
    }

    @Override
    public void onTerminate(final Cluster cluster) {
        // TODO: close Rust resources after the FFI boundary exists.
    }

    public Cluster cluster() {
        return cluster;
    }

    public Cluster.Role role() {
        return role;
    }

    public int committedMessages() {
        return committedMessages;
    }
}
