package ecdsad

import (
	"crypto/ecdsa"
	"flag-example/crypto/ecdsa/common"
	curiepb "flag-example/proto"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type Signature struct {
	sig []byte
}

func InitSignFromProto(i interface{}) (common.Signature, error) {
	switch pb := i.(type) {
	case *curiepb.Signature:
		return &Signature{sig: pb.Sig}, nil
	default:
		return nil, errors.Wrapf(errors.New("unsupported signed curie block"), "unable to create block from type %T", i)
	}
}

func (s *Signature) Verify(pubKey *ecdsa.PublicKey, msg []byte) bool {
	// if crypto.VerifySignature(crypto.CompressPubkey(pubKey), msg, s.sig) {
	// 	return true
	// } else {
	// 	logrus.Info("Original Public Key: ", crypto.CompressPubkey(pubKey))
	// 	logrus.Info("Signature: ", s.sig)
	// 	return false
	// }

	sigPublicKey, err := crypto.Ecrecover(msg, s.sig)
	if err != nil {
		logrus.WithError(err).Error("Failed to recover public key")
		return false
	}

	recoveredPubKey, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		log.Fatalf("Failed to unmarshal public key: %v", err)
	}

	// 실제 주소 비교
	recoveredAddr := crypto.PubkeyToAddress(*recoveredPubKey)
	originalAddr := crypto.PubkeyToAddress(*pubKey)
	if recoveredAddr.Hex() == originalAddr.Hex() {
		return true
	} else {
		fmt.Println("Recovered Address: ", recoveredAddr.Hex())
		fmt.Println("Original Address: ", originalAddr.Hex())
		return false
	}

	// logrus.Info("Original Public Key: ", len(sigPublicKey))
	// logrus.Info("Recovered Public Key: ", len(crypto.CompressPubkey(pubKey)))
	// k, err := crypto.DecompressPubkey(sigPublicKey)
	// if err != nil {
	// 	logrus.WithError(err).Error("Failed to decompress public key")
	// 	return false
	// }

	// if bytes.Equal(crypto.CompressPubkey(k), crypto.CompressPubkey(pubKey)) {
	// 	return true
	// } else {
	// 	logrus.Info("Original Public Key: ", crypto.CompressPubkey(pubKey))
	// 	logrus.Info("Recovered Public Key: ", crypto.CompressPubkey(k))
	// 	logrus.Info(bytes.Equal(crypto.CompressPubkey(k), crypto.CompressPubkey(pubKey)))
	// 	logrus.Info("Signature: ", s.sig)
	// 	return false
	// }

	// if (sigPublicKey != nil) && bytes.Equal(crypto.CompressPubkey(pubKey), sigPublicKey) {
	// 	return true
	// } else {
	// 	logrus.Info("Original Public Key: ", crypto.CompressPubkey(pubKey))
	// 	logrus.Info("Recovered Public Key: ", sigPublicKey)
	// 	logrus.Info(bytes.Equal(crypto.CompressPubkey(pubKey), sigPublicKey))
	// 	logrus.Info("Signature: ", s.sig)
	// 	return false
	// }
}

func (s *Signature) Marshal() []byte {
	return s.sig
}

func (s *Signature) Proto() (proto.Message, error) {
	return &curiepb.Signature{
		Sig: s.sig,
	}, nil
}
