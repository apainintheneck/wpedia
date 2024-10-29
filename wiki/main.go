package wiki

import (
	"fmt"
	"strings"

	gowiki "github.com/trietmn/go-wiki"
)

func FetchContent(title string) string {
	page, err := gowiki.GetPage(title, -1, false, true)
	if err != nil {
		fmt.Println(err)
	}

	content, err := page.GetContent()
	if err != nil {
		fmt.Println(err)
	}

	header := ""
	indent := ""
	lines := strings.Split(content, "\n")
	var builder strings.Builder

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
			header = line
			continue
		case strings.HasPrefix(line, "=== "):
			indent = "    "
		case strings.HasPrefix(line, "==== "):
			indent = "      "
		}

		if header != "" {
			builder.WriteRune('\n')
			builder.WriteString("  ")
			builder.WriteString(header)
			builder.WriteRune('\n')
			header = ""
		}

		builder.WriteRune('\n')
		builder.WriteString(indent)
		builder.WriteString(line)
		builder.WriteRune('\n')
	}

	return builder.String()
}
