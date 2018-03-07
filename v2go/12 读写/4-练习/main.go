package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	page := Page{
		"page.md",
		[]byte("# page\n## 1.1demo\nasdfwqerqwefqwefdfasdf"),
	}
	page.save()
	var newPage Page
	newPage.load("page.md")
	fmt.Print(string(newPage.Body))
}

func (page *Page) save() error {
	return ioutil.WriteFile(page.Title, page.Body, 0666)
}

func (this *Page) load(title string) error {
	this.Title = title
	body, err := ioutil.ReadFile(title)
	this.Body = body
	return err
}
