package onetext

// Sentense 句子数据结构
type Sentense struct {
	Text string   `json:"text"`
	By   string   `json:"by,omitempty"`
	From string   `json:"from,omitempty"`
	Time []string `json:"time"`
	Uri  string   `json:"uri,omitempty"`
}

type onetext struct {
	s []Sentense
}
