package auth

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"log"
)

const saltSize = 64

func makePassword(rawPassword string) string {
	buffer := make([]byte, saltSize, saltSize+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buffer)
	if err != nil {
		log.Fatal("random read failed ->", err)
	}

	hash := sha1.New()
	hash.Write(buffer)
	hash.Write([]byte(rawPassword))

	return base64.URLEncoding.EncodeToString(hash.Sum(buffer))
}

func checkPassword(rawPossiblePassword string, user User) bool {
	data, _ := base64.URLEncoding.DecodeString(user.PasswordHash)
	if len(data) != saltSize+sha1.Size {
		log.Fatal("wrong length of data")
		return false
	}
	hash := sha1.New()
	hash.Write(data[:saltSize])
	hash.Write([]byte(rawPossiblePassword))
	return bytes.Equal(hash.Sum(nil), data[saltSize:])
}
