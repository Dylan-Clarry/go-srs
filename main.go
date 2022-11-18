package main

import (
	"fmt"
	"os"

    "github.com/charmbracelet/bubbles/list"
    "github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)

var doc = lipgloss.NewStyle().Margin(1, 2)

type item struct {
    title, desc string
}

func (i item) Title() string        { return i.title }
func (i item) Description() string  { return i.title }
func (i item) FilterValue() string  { return i.title }

type model struct {
    decks   list.Model
}

func getDecks() []list.Item {
    return []list.Item {
        item {title: "JavaScript", desc: "Standard Deck"},
        item {title: "TypeScript", desc: "Standard Deck"},
        item {title: "React", desc: "Standard Deck"},
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "ctrl+c" || msg.String() == "q" {
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        h, v := doc.GetFrameSize()
        m.decks.SetSize(msg.Width - h, msg.Height - v)
    }

    var cmd tea.Cmd
    m.decks, cmd = m.decks.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return doc.Render(m.decks.View())
}

func main() {
    decks := getDecks()
    m := model{decks: list.New(decks, list.NewDefaultDelegate(), 0, 0)}
        m.decks.Title = "Decks"


    p := tea.NewProgram(m, tea.WithAltScreen())

    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v", err)
        os.Exit(1)
    }
}
