package simulator

import (
	"context"
	"crypto/ecdsa"
	"flag-example/blocks/interfaces"
	"fmt"

	"github.com/sirupsen/logrus"
)

type BlockReceiver interface {
	ReceiveOGBlock(ctx context.Context, block interfaces.SignedCurieBlock, pubKey *ecdsa.PublicKey) error
	ReceiveNGBlock(ctx context.Context, block interfaces.SignedCurieBlock) error
}

func ReceiveOGBlock(ctx context.Context, block interfaces.SignedCurieBlock, pubKey *ecdsa.PublicKey) error {
	/* Check Received Data for Validation */

	logrus.Info("@@ STEP_2_1 @@")

	// Hashing Received Data
	hash := block.Hash()

	logrus.Info("@@ STEP_2_2 @@")

	// Decryption Signature && Compare Hashing and Decryption Signature
	// 1. 수신 데이터로부터 서명 데이터 추출하기
	sig := block.Signature()

	logrus.Info("@@ STEP_2_3 @@")

	// 2. 공개키 이용하여 Verify() 함수 호출하기
	if sig.Verify(pubKey, hash) {
		// If it is valid, Send Normal ACK to Check Node
		fmt.Println("Received Data is Valid")
	} else {
		// If it is invalid, Send NACK to Check Node
		fmt.Println("Received Data is Non-Valid")
	}

	logrus.Info("@@ STEP_2_4 @@")

	return nil
}

func ReceiveNGBlock(ctx context.Context, block interfaces.SignedCurieBlock) error {
	/* Check Received Data for Validation */

	// Hashing Received Data
	// hash := block.Hash()

	// Decryption Signature && Compare Hashing and Decryption Signature
	// 1. 수신 데이터로부터 서명 데이터 추출하기
	// sig := block.Signature()

	// 2. 공개키 이용하여 Verify() 함수 호출하기

	// If it is valid, Send Normal ACK to Check Node

	// If it is invalid, Send NACK to Check Node

	return nil
}
