package models

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)


type ProjectModel struct {
	header    string
	textInput textinput.Model
	err       error
	program   *Program
}


func InitializeProjectModel(header string, program *Program) ProjectModel {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.Placeholder = "my-project.."

	return ProjectModel{
		textInput: ti,
		header:    titleStyle.Render(header),
		program:   program,
	}
}

func (m ProjectModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m ProjectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if len(m.textInput.Value()) > 1 {
				m.program.Project = m.textInput.Value()
				return m, tea.Quit
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			m.program.Exit = true
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}


func (m ProjectModel) View() string {
	return fmt.Sprintf("%s\n\n%s\n\n",
		m.header,
		m.textInput.View(),
	)
}

