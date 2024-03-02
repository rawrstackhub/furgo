package furgo 

import (
	"fmt"
	"bytes"
	"html/template"
)

type UIElement interface {
	Render() (string, error)
}

type DivElement struct {
	ID	string
	Class	string
	Content string
}

func (d *DivElement) Render() (string, error) {
	t := `<div id="{{.ID}}" class="{{.Class}}">{{.Content}}</div>`
	tmpl, err := template.New("div").Parse(t)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	err = tmpl.Execute(&out, d)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func NewDivElement(id, class, content string) *DivElement {
	return &DivElement{
		ID:		id,
		Class:		class,
		Content:	content,
	}
}

func potato {
	fmt.Println("Potato")
} 
