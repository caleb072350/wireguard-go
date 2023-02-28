package device

import "sync"

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
}
