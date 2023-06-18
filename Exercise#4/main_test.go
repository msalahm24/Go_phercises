package main

import (
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestExtractHref(t *testing.T) {
	cases := []struct{
		name string
		a string
		href string
	}{
		{
			name: "Valied Element",
			a:`<a href="/login"></a>`,
			href: "/login",
		},
		{
			name: "missing href",
			a:`<a></a>`,
			href: "",
		},
	}

	for _, c := range cases{
		t.Run(c.name,func(t *testing.T) {
			a := parse(t,c.a)
			html.Render(os.Stdout,a)
			href := extractHref(a)
			if href != c.href{
				t.Fatalf("expexted %q, got %q",c.href,href)
			}
		})
	}
}

func TestExtractText(t *testing.T) {
	cases := []struct{
		name string
		a string
		text string
	}{
		{
			name: "Valied Element",
			a:`<a href="/login">login</a>`,
			text: "login",
		},{
			name: "nested text elements",
			a:`<a href="/login">login<span>as user</span></a>`,
			text: "loginas user",
		},
		{
			name: "nested text elements",
			a:`<a href="/login">login<span>as user<!-- this is comment --!></span></a>`,
			text: "loginas user",
		},
	}

	for _, c := range cases{
		t.Run(c.name,func(t *testing.T) {
			a := parse(t,c.a)
			html.Render(os.Stdout,a)
			text := extractText(a)
			if text != c.text{
				t.Fatalf("expexted %q, got %q",c.text,text)
			}
		})
	}
}
func parse(t *testing.T,a string) *html.Node{
	root, err := html.Parse(strings.NewReader(a))

	
	if err != nil{
		t.Fatalf("Failed to parse %v",err)
	}
	return root.FirstChild.FirstChild.NextSibling.FirstChild
}
