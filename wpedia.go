package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/apainintheneck/wpedia/model"
	tea "github.com/charmbracelet/bubbletea"
	gowiki "github.com/trietmn/go-wiki"
)

const maxSearchResults = 50

func main() {
	searchTerm := strings.Join(os.Args[1:], " ")
	var searchResults []string
	var err error

	if searchTerm == "" {
		searchTerm = "Random"
		searchResults, err = gowiki.GetRandom(maxSearchResults)
	} else {
		searchResults, _, err = gowiki.Search(searchTerm, maxSearchResults, false)
	}

	if err != nil {
		fmt.Println(err)
	}

	p := tea.NewProgram(
		model.New(searchTerm, searchResults),
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
