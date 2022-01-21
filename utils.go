package onetext

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// ReadBytes 读取字节
func (o *onetext) ReadBytes(b []byte) error {
	return json.Unmarshal(b, &o.s)
}

// ReadFile 读取 json 文件
func (o *onetext) ReadFile(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &o.s)
}

// GetUrl 从 url 获取 json
func (o *onetext) GetUrl(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &o.s)
}

// Random 随机一个句子
func (o *onetext) Random() Sentense {
	if len(o.s) != 0 {
		return o.s[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(o.s))]
	}
	return Sentense{}
}
