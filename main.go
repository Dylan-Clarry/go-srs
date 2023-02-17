package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbles/list"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type deck struct {
    title, desc     string
}

func(d deck) Title() string       { return d.title }
func(d deck) Description() string { return d.desc }
func(d deck) FilterValue() string { return d.title }

func getDecks() []list.Item {
    decks := []list.Item {
        deck{title: "JavaScript", desc: "JS mock interview questions"},
        deck{title: "Japanese", desc: "Japanese word deck"},
    }
    return decks
}

type model struct {
    list    list.Model
}

func(m model) Init() tea.Cmd {
    return nil
}

func(m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "ctrl+c" {
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        h, v := docStyle.GetFrameSize()
        m.list.SetSize(msg.Width-h, msg.Height-v)
    }

    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func(m model) View() string {
    return docStyle.Render(m.list.View())
}

func main() {
    decks := getDecks()
    m := model{list: list.New(decks, list.NewDefaultDelegate(), 0, 0)}

    p := tea.NewProgram(m, tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("There's been an error, oopsie woopsie!\nerror: %v", err)
        os.Exit(1)
    }
}
