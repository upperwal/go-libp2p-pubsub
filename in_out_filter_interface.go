package pubsub

import "github.com/libp2p/go-libp2p-core/peer"

type PeerFilter interface {
	FilterSubscriber(peer.ID, string) bool
	FilterPublisher(peer.ID, []string) []bool
}
