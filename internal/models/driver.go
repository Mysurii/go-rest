package models

import (
	"fmt"
	"go-rest/internal/steps"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Change this
var (
	titleStyle            = lipgloss.NewStyle().Background(lipgloss.Color("#818cf8")).Foreground(lipgloss.Color("#3730a3")).Bold(true)
	focusedStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#e879f9")).Bold(true)
	selectedItemStyle     = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#e879f9")).Bold(true)
	selectedItemDescStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#e879f9"))
	nameStype			  = lipgloss.NewStyle().Foreground(lipgloss.Color("#134e4a")).Bold(true)
	descriptionStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#134e4a"))
)

type DriverModel struct {
	header string
	cursor int
	options []steps.Option
	selected int
}

func (m DriverModel) Init() tea.Cmd {
	return nil
}

func InitialDriverModel(header string) DriverModel {
	return DriverModel{
		header: header,
		options:  steps.DriverOptions,
		selected: 0 ,
	}
}

func (m DriverModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selected = m.cursor
		case "y":
			return m, tea.Quit
			
				
		}
	}
	return m, nil
}

func (m DriverModel) View() string {
	s := titleStyle.Render(m.header)
	s += "\n\n"

	for i, option := range m.options {
		tpl := " "
		if m.cursor == i {
			tpl = focusedStyle.Render(">")
			option.Title = selectedItemStyle.Render(option.Title)
			option.Description = selectedItemDescStyle.Render(option.Description)

		}

		checked := " "
		if m.selected == i {
			checked = focusedStyle.Render("X")
		}

		title := nameStype.Render(option.Title)
		description := descriptionStyle.Render(option.Description)

		s += fmt.Sprintf("%s [%s] %s\n%s\n\n", tpl, checked, title, description)
	}

	s += fmt.Sprintf("Press %s to confirm choice.\n", focusedStyle.Render("y"))
	return s
}

func (m DriverModel) GetDriver() string {
	return m.options[m.selected].Title
}