package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
)

// GenerateInviteCode ...
func GenerateInviteCode(n int) string {
	letterRunes := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var bb bytes.Buffer
	bb.Grow(n)
	l := uint32(len(letterRunes))
	for i := 0; i < n; i++ {
		bb.WriteRune(letterRunes[binary.BigEndian.Uint32(readBytes(4))%l])
	}
	return bb.String()
}

func readBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
