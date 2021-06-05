<<<<<<< HEAD
package main

import (
	"math"
	"math/big"
	"crypto/sha256"
	"bytes"
	"fmt"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 10

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}


// Run performs a proof-of-work
// implement

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashData(),//pow.block.Hash
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	nonce := 0
	var hashInt big.Int
	var hash [32]byte
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			pow.block.Hash = hash[:]
			break
		} else {
			if nonce < maxNonce{
				nonce++
			} else {
				nonce = -1
				return nonce,hash[:]
			}
		}
	}

	return nonce, pow.block.Hash
}

// Validate validates block's PoW
// implement
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashData(),//pow.block.Hash
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	if hashInt.Cmp(pow.target) == -1{
		return true
	} else {
		return false
	}
}
=======
package main

import (
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 10

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}


// Run performs a proof-of-work
// implement
func (pow *ProofOfWork) Run() (int, []byte) {
	nonce := 0

	return nonce, pow.block.Hash
}

// Validate validates block's PoW
// implement
func (pow *ProofOfWork) Validate() bool {
	return true
}
>>>>>>> a8f7e8aacf32b00972f74cc3119f9dd72652936f
