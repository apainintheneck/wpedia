package main

import (
	"fmt"
	"os"
	"strings"

	gowiki "github.com/trietmn/go-wiki"
)

func main() {
	searchTerm := strings.Join(os.Args[1:], " ")
	var searchResult []string
	var err error

	if searchTerm == "" {
		searchResult, err = gowiki.GetRandom(1)
	} else {
		searchResult, _, err = gowiki.Search(searchTerm, 1, false)
	}

	if err != nil {
		fmt.Println(err)
	}

	title := searchResult[0]

	page, err := gowiki.GetPage(title, -1, false, true)
	if err != nil {
		fmt.Println(err)
	}

	content, err := page.GetContent()
	if err != nil {
		fmt.Println(err)
	}

	indent := ""
	lines := strings.Split(content, "\n")

	fmt.Printf("= %s =\n", title)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if line == "== See also ==" || line == "== References ==" {
			break
		}

		switch {
		case strings.HasPrefix(line, "== "):
			indent = "  "
		case strings.HasPrefix(line, "=== "):
			indent = "    "
		case strings.HasPrefix(line, "==== "):
			indent = "      "
		}

		fmt.Printf("\n%s%s\n", indent, line)
	}
}
