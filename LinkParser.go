package util

import (
	"golang.org/x/net/html"
	"io"
)

type Link struct {
	Text, Href string
}

func Parse(r io.Reader) []Link{

	doc, err := html.Parse(r)

	if err != nil{
		panic(err)
	}

	var array []Link

	for _, child := range dfs(doc){
		array = append(array, toLink(child))
	}

	return array
}

func toLink(node *html.Node) Link{

	var link Link
	link.Text = getText(node)

	for _, a := range node.Attr{
		if a.Key == "href"{
			link.Href = a.Val
			break
		}
	}

	return link
}

func getText(node *html.Node) string{

	if node.Type == html.TextNode{
		return node.Data
	}
	if node.Type != html.ElementNode{
		return ""
	}

	var str string
	for elem := node.FirstChild; elem != nil; elem = elem.NextSibling{
		str += getText(elem)
	}

	return str
}

func dfs(node *html.Node) []*html.Node{

	if node.Type == html.ElementNode && node.Data == "a"{
		return []*html.Node{node}
	}

	var array []*html.Node

	for elem := node.FirstChild; elem != nil; elem = elem.NextSibling{
		array = append(array, dfs(elem)...)
	}

	return array
}