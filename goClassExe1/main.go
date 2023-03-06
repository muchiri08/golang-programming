package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

var raw = `<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
      <img src="yyy.jpg" width="104" height="142">
  </body>
</html>`

func visit(n *html.Node, pwords, ppics *int) {

	if n.Type == html.TextNode {
		*pwords += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*ppics++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, pwords, ppics)
	}
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pics int

	visit(doc, &words, &pics)
	return words, pics
}

func main() {
	doc, error := html.Parse(bytes.NewReader([]byte(raw)))

	if error != nil {
		fmt.Fprintf(os.Stderr, "Parse failed: %s\n", error)
		os.Exit(-1)
	}

	words, pics := countWordsAndImages(doc)

	fmt.Printf("%d words and %d images\n", words, pics)
}
