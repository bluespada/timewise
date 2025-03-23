// code source : https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go
package crypt

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"runtime"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type Argon2ID struct {
	config Argon2Params
}

type Argon2Params struct {
	Memory      uint32
	Parallelism uint8
	Iterations  uint32
	SaltLength  uint32
	KeyLength   uint32
	Papper      string
}

var DefaultArgon2Params = Argon2Params{
	Memory:      64 * 1024,
	Parallelism: uint8(runtime.NumCPU()),
	Iterations:  3,
	SaltLength:  16,
	KeyLength:   32,
	Papper:      "default",
}

func NewArgon(config Argon2Params) *Argon2ID {
	return &Argon2ID{
		config: config,
	}
}

func (a *Argon2ID) Hash(password string) (string, error) {
	salt, err := a.generateRandomBytes(a.config.SaltLength)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password+a.config.Papper), salt, a.config.Iterations, a.config.Memory, a.config.Parallelism, a.config.KeyLength)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, a.config.Memory, a.config.Iterations, a.config.Parallelism, b64Salt, b64Hash)
	return encodedHash, nil
}

func (a *Argon2ID) Compare(passwordHash string, password string) (bool, error) {

	p, salt, hash, err := a.decodeHash(passwordHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password+a.config.Papper), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func (a *Argon2ID) generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (a *Argon2ID) decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
