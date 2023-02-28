package device

import (
	"sync"
	"sync/atomic"

	"github.com/caleb072350/wireguard-go/tun"
)

const (
	DeviceRoutineNumberPerCPU     = 3
	DeviceRoutineNumberAdditional = 2
)

type Device struct {
	isUp     AtomicBool // device is (going) up
	isClosed AtomicBool // device is closed? (acting as guard)
	log      *Logger

	// synchronized resources (locks acquired in order)

	state struct {
		starting sync.WaitGroup
		stopping sync.WaitGroup
		sync.Mutex
		changing AtomicBool
		current  bool
	}

	net struct {
		starting sync.WaitGroup
		stopping sync.WaitGroup
		sync.RWMutex
		bind   Bind   // bind interface
		port   uint16 // listening port
		fwmark uint32 // mark value (0 = disabled)
	}

	staticIdentity struct {
		sync.RWMutex
		privateKey NoisePrivateKey
		publicKey  NoisePublicKey
	}

	peers struct {
		sync.RWMutex
		keyMap map[NoisePublicKey]*Peer
	}

	// unprotected / "self-synchronising resources"

	allowedips    allowedips
	indexTable    indexTable
	cookieChecker cookieChecker

	rate struct {
		underLoadUntil atomic.Value
		limiter        ratelimiter.ratelimiter
	}

	pool struct {
		messageBufferPool        *sync.Pool
		messageBufferReuseChan   chan *[MaxMessageSize]byte
		inboundElementPool       *sync.Pool
		inboundElementReuseChan  chan *QueueInboundElement
		outboundElementPool      *sync.Pool
		outboundElementReuseChan chan *QueueOutboundElement
	}

	queue struct {
		encryption chan *QueueOutboundElement
		decryption chan *QueueInboundElement
		handshake  chan QueueHandshakeElement
	}

	signals struct {
		stop chan struct{}
	}

	tun struct {
		device tun.Device
		mtu    int32
	}
}
