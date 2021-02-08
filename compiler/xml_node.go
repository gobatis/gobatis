package compiler

type XMLNode struct {
	Start        *Point          `json:"-"`
	End          *Point          `json:"-"`
	Type         string          `json:"type"`
	Name         string          `json:"name,omitempty"`
	Value        string          `json:"value,omitempty"`
	RAW          string          `json:"value,omitempty"`
	Attributes   []*XMLAttribute `json:"attributes,omitempty"`
	Body         []*XMLNode      `json:"body,omitempty"`   // tree body
	Tokens       []*Token        `json:"tokens,omitempty"` // SQL tokens
	attributeMap map[string]string
	closed       bool
}

func (p *XMLNode) AppendBody(node ...*XMLNode) {
	p.Body = append(p.Body, node...)
}

type XMLAttribute struct {
	Start *Point `json:"-"`
	End   *Point `json:"-"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
