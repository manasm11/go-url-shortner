package inmemory

import (
	"encoding/json"
	"log"

	"github.com/manasm11/go-url-shortner/shortner"
)

type InMemorySerializer struct{}

func (s *InMemorySerializer) Decode(bs []byte) (r *shortner.Redirect, err error) {
	r = &shortner.Redirect{}
	log.Println("InMemorySerializer.Decode", string(bs))
	err = json.Unmarshal(bs, r)
	return
}

func (s *InMemorySerializer) Encode(r *shortner.Redirect) ([]byte, error) {
	return json.Marshal(r)
}
