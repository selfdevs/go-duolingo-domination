package discord

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

const invalidSignature = "signature"
const invalidTimestamp = "timestamp"

func createRequestWithInvalidSignature() *http.Request {
	req, err := http.NewRequest("GET", "/interactions", nil)
	if err != nil {
		panic(err)
	}
	req.Body = io.NopCloser(bytes.NewBuffer([]byte("body")))
	req.Header.Set("X-Signature-Ed25519", invalidSignature)
	req.Header.Set("X-Signature-Timestamp", invalidTimestamp)
	return req
}

func TestGetRequestSignature(t *testing.T) {
	req := createRequestWithInvalidSignature()
	signature := GetRequestSignature(req)

	if signature.Signature != "signature" {
		t.Fatalf("Expected signature to be \"signature\", got \"%s\"", signature.Signature)
	}
	if signature.Timestamp != "timestamp" {
		t.Fatalf("Expected timestamp to be \"timestamp\", got \"%s\"", signature.Timestamp)
	}
}

func TestVerifySignature(t *testing.T) {
	req := createRequestWithInvalidSignature()
	err := VerifySignature(req)
	if err != nil {
		t.Skip()
	}
	t.Fail()
}
