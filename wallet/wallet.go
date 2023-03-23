package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/claerhead/go_blockchain/utils"
)

const (
	signature     string = "d2d42c0a9a707223719f7441cd2c047da7485d7194ad6cc518e960cc95c351d139996b99e4aea9a4778032833efea7b3a0dd25bb77cce404575069b90873d3a4"
	hashedMessage string = "c33084feaa65adbbbebd0c9bf292a26ffc6dea97b170d88e501ab4865591aafd"
	privateKey    string = "307702010104204227fadabcafad367b71d92ad8755dada1c4319a8d236fc69b6d1d3edeadf930a00a06082a8648ce3d030107a14403420004195356dda3befe57c09ab934e71b8ca5f6f79c48694d42098d4e86e46e89c143548a6742e11cc17ab6b95fa970b58f2908e5fca2affc8e49217d2311020a0ac3"
)

func Start() {
	privByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	_, err = x509.ParseECPrivateKey(privByte)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	utils.HandleErr(err)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)
}
