package discord

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"os"
)

type RequestSignature struct {
	Signature string
	Timestamp string
}

func VerifySignature(r *http.Request) error {
	signature := getRequestSignature(r)

	_rawBody, err := io.ReadAll(r.Body)

	r.Body = io.NopCloser(bytes.NewBuffer(_rawBody))

	if err != nil {
		println("Fail to read body")
		os.Exit(1)
	}

	message := []byte(signature.Timestamp + string(_rawBody))
	publicKey := os.Getenv("DISCORD_PUBLIC_KEY")

	hexPublicKey, _ := hex.DecodeString(publicKey)
	hexSignature, _ := hex.DecodeString(signature.Signature)

	verified := ed25519.Verify(hexPublicKey, message, hexSignature)

	println("Verified:", verified)

	if !verified {
		return errors.New("invalid signature")
	}

	return nil
}

func getRequestSignature(r *http.Request) RequestSignature {
	signature := r.Header.Get("X-Signature-Ed25519")
	timestamp := r.Header.Get("X-Signature-Timestamp")

	return RequestSignature{
		Signature: signature,
		Timestamp: timestamp,
	}
}
