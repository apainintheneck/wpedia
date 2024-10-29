package main

import (
	"flag"
	"log"
	"strings"

	"github.com/apainintheneck/wpedia/model"
	tea "github.com/charmbracelet/bubbletea"
	gowiki "github.com/trietmn/go-wiki"
)

const maxSearchResults = 50

func main() {
	languagePtr := flag.String("l", "en", "2 digit language code")
	flag.Parse()

	gowiki.SetLanguage(*languagePtr)

	searchTerm := strings.Join(flag.Args(), " ")
	var searchResults []string
	var err error

	if searchTerm == "" {
		searchTerm = "Random"
		searchResults, err = gowiki.GetRandom(maxSearchResults)
	} else {
		searchResults, _, err = gowiki.Search(searchTerm, maxSearchResults, false)
	}

	if err != nil {
		log.Fatal("Error: Unable to search for wikipedia articles\n\n", err)
	}

	p := tea.NewProgram(
		model.New(searchTerm, searchResults),
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if _, err := p.Run(); err != nil {
		log.Fatal("Error: Unexpected error program error\n\n", err)
	}
}
