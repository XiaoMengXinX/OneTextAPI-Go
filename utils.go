package onetext

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// New onetext
func New() *onetext {
	return &onetext{nil}
}

// ReadBytes read json from bytes
func (o *onetext) ReadBytes(b []byte) (*onetext, error) {
	var d []Sentence
	err := json.Unmarshal(b, &d)
	if err != nil {
		return o, err
	}
	o.s = append(o.s, d...)
	return o, nil
}

// ReadFile read json from a file
func (o *onetext) ReadFile(path string) (*onetext, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return o, err
	}
	return o.ReadBytes(b)
}

// GetUrl get json from an url
func (o *onetext) GetUrl(url string) (*onetext, error) {
	resp, err := http.Get(url)
	if err != nil {
		return o, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)
	b, err := ioutil.ReadAll(resp.Body)
	return o.ReadBytes(b)
}

// Random get a random sentence
func (o *onetext) Random() Sentence {
	if len(o.s) != 0 {
		return o.s[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(o.s))]
	}
	return Sentence{}
}

// AddOne add a sentence
func (o *onetext) AddOne(sentence Sentence) *onetext {
	o.s = append(o.s, sentence)
	return o
}

// Add sentences
func (o *onetext) Add(sentences []Sentence) *onetext {
	o.s = append(o.s, sentences...)
	return o
}
