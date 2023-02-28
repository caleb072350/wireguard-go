package device

import "sync"

const (
	PeerRoutineNumber = 3
)

type Peer struct {
	isRunning    AtomicBool
	sync.RWMutex // Mostly protects endpoint, but is generally taken whenever we modify peer
	keypairs     keypairs
}
