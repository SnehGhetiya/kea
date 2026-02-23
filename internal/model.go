package internal

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	primaryColor = lipgloss.Color("#d04f99")

	logoStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			MarginLeft(2)

	greetingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8acfd1")).
			Bold(true).
			MarginLeft(2).
			MarginTop(1)

	subtextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#fbe2a7")).
			MarginLeft(2).
			MarginTop(1)

	helperTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f96f70")).
			MarginLeft(2).
			MarginTop(1)

	goodbyeTextStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#f6e6ee")).
				MarginLeft(2).
				MarginTop(1)
)

type Model struct {
	spinner  spinner.Model
	cursor   int
	quitting bool
	err      error
}

func New() Model {
	s := spinner.New()
	s.Spinner = spinner.Ellipsis
	s.Style = lipgloss.NewStyle().Foreground(primaryColor)
	return Model{
		spinner: s,
	}
}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}
	case error:
		m.err = msg
		return m, nil
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	renderedLogo := logoStyle.Render(KeaLogo)
	greeting := greetingStyle.Render("KEA Agent initialized. System standing by.")
	waitingText := subtextStyle.Render("Waiting for user input " + m.spinner.View())
	helperText := helperTextStyle.Render("Press 'q' to quit.")
	goodbyeText := goodbyeTextStyle.Render("Goodbye!")

	finalStr := fmt.Sprintf("\n%s\n%s\n%s\n%s\n", renderedLogo, greeting, waitingText, helperText)

	if m.quitting {
		return finalStr + goodbyeText
	}
	return finalStr
}
