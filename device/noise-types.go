package device

import (
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/chacha20poly1305"
)

const (
	NoisePublicKeySize  = 32
	NoisePrivateKeySize = 32
)

type (
	NoisePublicKey    [NoisePublicKeySize]byte
	NoisePrivateKey   [NoisePrivateKeySize]byte
	NoiseSymmetricKey [chacha20poly1305.KeySize]byte
	NoiseNonce        uint64 // padded to 12-bytes
)

func loadExactHex(dst []byte, src string) error {
	slice, err := hex.DecodeString(src)
	if err != nil {
		return err
	}
	if len(slice) != len(dst) {
		return errors.New("hex string does not fit the slice")
	}
	copy(dst, slice)
	return nil
}
