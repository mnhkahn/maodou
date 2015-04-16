package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"strconv"
	"strings"
	// "github.com/PuerkitoBio/goquery"
	"io"
)

var (
	SCORE_TAG        = "maodou"
	COMPUTE_FUNC     = []func(*html.Node){init_f, include_f, exclude_f, text_f, res_f}
	INCLUDE_ID_CLASS = []string{"content"}                                    //className或id包含content…加分
	EXCLUDE_ID_CLASS = []string{"header", "foot", "sidebar", "about", "logo"} // className或id为header|foot|sidebar…减分
	SET_PUNCTUATION  = []string{",", ".", "，", "。"}

	largest_score            = 0
	largest_node  *html.Node = nil
)

func Content(r io.Reader) string {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	for _, f := range COMPUTE_FUNC {
		Tranverse(doc, f)
	}

	res := getNodeText(largest_node)
	b := bufio.NewReader(strings.NewReader(res))
	line, err := b.ReadString('\n')
	max_enter := 0

	var buf bytes.Buffer
	for ; max_enter < 4 && err == nil; line, err = b.ReadString('\n') {
		buf.WriteString(line)
		if line == "\n" {
			max_enter++
		} else {
			max_enter = 0
		}
	}
	return buf.String()
}

func init_f(n *html.Node) {
	SetAttr(n, SCORE_TAG, "0")
}

func include_f(n *html.Node) {
	id := Attr(n, "id")
	class := Attr(n, "class")
	for _, include := range INCLUDE_ID_CLASS {
		if strings.Contains(strings.ToLower(id), include) {
			score, _ := strconv.Atoi(Attr(n, SCORE_TAG))
			SetAttr(n, SCORE_TAG, fmt.Sprintf("%d", score+1))
		}
		if strings.Contains(strings.ToLower(class), include) {
			score, _ := strconv.Atoi(Attr(n, SCORE_TAG))
			SetAttr(n, SCORE_TAG, fmt.Sprintf("%d", score+1))
		}
	}
}

func exclude_f(n *html.Node) {
	id := Attr(n, "id")
	class := Attr(n, "class")
	for _, exclude := range EXCLUDE_ID_CLASS {
		if strings.Contains(strings.ToLower(id), exclude) {
			score, _ := strconv.Atoi(Attr(n, SCORE_TAG))
			SetAttr(n, SCORE_TAG, fmt.Sprintf("%d", score-1))
		}
		if strings.Contains(strings.ToLower(class), exclude) {
			score, _ := strconv.Atoi(Attr(n, SCORE_TAG))
			SetAttr(n, SCORE_TAG, fmt.Sprintf("%d", score-1))
		}
	}
}

func text_f(n *html.Node) {
	text := getNodeText(n)
	if len(text) > 150 {
		score, _ := strconv.Atoi(Attr(n, SCORE_TAG))
		SetAttr(n, SCORE_TAG, fmt.Sprintf("%d", score+1))
	}
	for _, punctuation := range SET_PUNCTUATION {
		if strings.Contains(text, punctuation) {
			score, _ := strconv.Atoi(Attr(n, SCORE_TAG))
			SetAttr(n, SCORE_TAG, fmt.Sprintf("%d", score+1))
		}
	}
}

func res_f(n *html.Node) {
	score, _ := strconv.Atoi(Attr(n, "maodou"))
	if score > largest_score {
		largest_score = score
		largest_node = n
	}
}

func Tranverse(n *html.Node, f func(*html.Node)) {
	if n.Type == html.ElementNode {
		f(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Tranverse(c, f)
	}
}

func Attr(n *html.Node, attrName string) string {
	if a := getAttributePtr(attrName, n); a != nil {
		return a.Val
	}
	return ""
}

func getAttributePtr(attrName string, n *html.Node) *html.Attribute {
	if n == nil {
		return nil
	}

	for i, a := range n.Attr {
		if a.Key == attrName {
			return &n.Attr[i]
		}
	}
	return nil
}

func SetAttr(n *html.Node, attrName, val string) {
	attr := getAttributePtr(attrName, n)
	if attr == nil {
		n.Attr = append(n.Attr, html.Attribute{Key: attrName, Val: val})
	} else {
		attr.Val = val
	}
}

func getNodeText(node *html.Node) string {
	if node.Type == html.TextNode {
		// Keep newlines and spaces, like jQuery
		return node.Data
	} else if node.FirstChild != nil {
		var buf bytes.Buffer
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if c.Data != "script" && c.Data != "noscript" {
				buf.WriteString(getNodeText(c))
			}
		}
		return buf.String()
	}

	return ""
}