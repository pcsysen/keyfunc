package keyfunc

import (
	"fmt"
)

const (

	// hs256 represents a cryptographic key generated by an SHA-256 hash and an HMAC algorithm.
	hs256 = "HS256"

	// hs384 represents a cryptographic key generated by an SHA-384 hash and an HMAC algorithm.
	hs384 = "HS384"

	// hs512 represents a cryptographic key generated by an SHA-512 hash and an HMAC algorithm.
	hs512 = "HS512"
)

// HMAC returns the HMAC key used for parsing and verifying the JWT.
//
// The jsonKey data structure does not
func (j *jsonKey) HMAC() (key []byte, err error) {

	// Confirm the key is already present as expected.
	j.precomputedMux.RLock()
	defer j.precomputedMux.RUnlock()
	if j.precomputed != nil {
		var ok bool
		if key, ok = j.precomputed.([]byte); ok {
			return key, nil
		}
	}

	// The HMAC key should have been pre-attached.
	return nil, fmt.Errorf("%w: hmac", ErrMissingAssets)
}
