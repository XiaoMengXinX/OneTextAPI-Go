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

// New OneText
func New() *OneText {
	return &OneText{nil}
}

// ReadBytes read json from bytes
func (o *OneText) ReadBytes(b []byte) error {
	var d []Sentence
	err := json.Unmarshal(b, &d)
	if err != nil {
		return err
	}
	o.s = append(o.s, d...)
	return nil
}

// ReadFile read json from a file
func (o *OneText) ReadFile(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return o.ReadBytes(b)
}

// GetUrl get json from an url
func (o *OneText) GetUrl(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)
	b, err := ioutil.ReadAll(resp.Body)
	return o.ReadBytes(b)
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
