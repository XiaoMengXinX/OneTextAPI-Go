package onetext

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// New OneText
func New() *OneText {
	return &OneText{nil}
}

// ReadBytes read json from bytes
func (o *OneText) ReadBytes(b []byte) (*OneText, error) {
	return o, json.Unmarshal(b, &o.s)
}

// ReadFile read json from a file
func (o *OneText) ReadFile(path string) (*OneText, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return o, err
	}
	return o, json.Unmarshal(b, &o.s)
}

// GetUrl get json from an url
func (o *OneText) GetUrl(url string) (*OneText, error) {
	resp, err := http.Get(url)
	if err != nil {
		return o, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return o, err
	}
	return o, json.Unmarshal(b, &o.s)
}

// Random get a random sentence
func (o *OneText) Random() Sentence {
	if len(o.s) != 0 {
		return o.s[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(o.s))]
	}
	return Sentence{}
}

// AddOne add a sentence
func (o *OneText) AddOne(sentence Sentence) *OneText {
	o.s = append(o.s, sentence)
	return o
}

// Add sentences
func (o *OneText) Add(sentences []Sentence) *OneText {
	o.s = append(o.s, sentences...)
	return o
}
