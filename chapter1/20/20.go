package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type secret struct {
	ID         string
	CreateTime time.Time

	token string
}

func (s *secret) Read(p []byte) (int, error) {
	return bytes.NewBuffer(p).WriteString(s.token)
}

func NewSecret() io.Reader {
	return &secret{
		ID:         "dummy_id",
		CreateTime: time.Now(),
		token:      "dummy_token",
	}
}

func main() {
	s := NewSecret()

	err := json.NewEncoder(os.Stdout).Encode(s)
	if err != nil {
		fmt.Println("failed to json encode, error =", err)
	}
}
