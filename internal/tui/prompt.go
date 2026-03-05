package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type PromptModel struct {
	question string
	cursor   int
	options  []string
	chosen   string
	quitting bool
}

func NewPromptModel(question string) PromptModel {
	return PromptModel{
		question: question,
		options:  []string{"Yes", "No"},
	}
}

func (m PromptModel) Init() tea.Cmd {
	return nil
}

func (m PromptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit
		case "left", "h":
			if m.cursor > 0 {
				m.cursor--
			}
		case "right", "l":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter":
			m.chosen = m.options[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m PromptModel) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder
	b.WriteString(titleStyle.Render("  paco") + "\n\n")
	b.WriteString(fmt.Sprintf("  %s\n\n  ", m.question))

	for i, opt := range m.options {
		if m.cursor == i {
			b.WriteString(promptSelectedStyle.Render(fmt.Sprintf("[ %s ]", opt)))
		} else {
			b.WriteString(promptNormalStyle.Render(fmt.Sprintf("  %s  ", opt)))
		}
		b.WriteString("  ")
	}

	b.WriteString(helpStyle.Render("\n\n  arrows to select | enter to confirm"))
	return b.String()
}

func (m PromptModel) Accepted() bool {
	return m.chosen == "Yes"
}
