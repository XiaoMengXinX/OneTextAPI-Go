package onetext

// Sentence data struct
type Sentence struct {
	Text string   `json:"text"`
	By   string   `json:"by,omitempty"`
	From string   `json:"from,omitempty"`
	Time []string `json:"time"`
	Uri  string   `json:"uri,omitempty"`
}

type onetext struct {
	s []Sentence
}
