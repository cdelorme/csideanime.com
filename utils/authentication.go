package authentication

import (
	"crypto/rand"
	"io"
)

// uuid version 4: populate 16 bytes with random bits, then set version number
// @link: http://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_.28random.29
func GenerateTokenSignature() []byte {
	sig := make([]byte, 16)

	// populate with random bits
	rander := rand.Reader
	if _, err := io.ReadFull(rander, sig); err != nil {
		panic(err.Error()) // supposedly, failure is not possible
	}

	// set version
	sig[6] = (sig[6] & 0x0f) | 0x40
	sig[8] = (sig[8] & 0x3f) | 0x80

	return sig
}

// Generate an Authorize Token
func GenerateAuthorizeToken() (string, error) {
	return strings.TrimRight(tokenbase64.URLEncoding.EncodeToString(GenerateTokenSignature()), "="), nil
}

// Generate an Access Token & Refresh (if applicable)
func GenerateAccessToken(fresh bool) (token string, refresh string) {
	token = strings.TrimRight(tokenbase64.URLEncoding.EncodeToString(GenerateTokenSignature()), "=")
	if fresh {
		refresh = strings.TrimRight(tokenbase64.URLEncoding.EncodeToString(GenerateTokenSignature()), "=")
	}
}

// we need a struct with structure support for persistence of token requests?
