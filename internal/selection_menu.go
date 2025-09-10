package internal

import (
	"fmt"
	"os"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// SelectionOption holds the label and the internal key
type SelectionOption struct {
	Label string
	Key   string
}

// Model represents the application state for the selection prompt
type Model struct {
	options        map[string]string // id -> name mapping
	filter         string
	filteredKeys   []SelectionOption
	selected       int
	terminalWidth  int
	terminalHeight int
	scrollOffset   int
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles user input and updates the model
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if wsm, ok := msg.(tea.WindowSizeMsg); ok {
		m.terminalWidth = wsm.Width
		m.terminalHeight = wsm.Height
	}

	updateFilter := false

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.filteredKeys[m.selected] = SelectionOption{Label: "quit", Key: "-1"}
			return m, tea.Quit
		case "enter":
			return m, tea.Quit
		case "backspace":
			if len(m.filter) > 0 {
				m.filter = m.filter[:len(m.filter)-1]
				updateFilter = true
			}
		case "down":
			if m.selected < len(m.filteredKeys)-1 {
				m.selected++
			}
			if m.selected >= m.scrollOffset+m.visibleItemsCount() {
				m.scrollOffset++
			}
		case "up":
			if m.selected > 0 {
				m.selected--
			}
			if m.selected < m.scrollOffset {
				m.scrollOffset--
			}
		default:
			if len(msg.String()) == 1 && msg.String() >= " " && msg.String() <= "~" {
				m.filter += msg.String()
				updateFilter = true
			}
		}
	}

	if updateFilter {
		m.filterOptions()
		m.selected = 0
		m.scrollOffset = 0
	}

	return m, nil
}

// View renders the UI
func (m Model) View() string {
	var b strings.Builder

	b.WriteString("Search (Press Ctrl+C to quit):\n")
	b.WriteString("Filter: " + m.filter + "\n\n")

	if len(m.filteredKeys) == 0 {
		b.WriteString("No matches found.\n")
	} else {
		visibleItems := m.visibleItemsCount()
		start := m.scrollOffset
		end := start + visibleItems
		if end > len(m.filteredKeys) {
			end = len(m.filteredKeys)
		}

		for i := start; i < end; i++ {
			entry := m.filteredKeys[i]
			prefix := "  "
			if i == m.selected {
				prefix = "▶ "
			}
			label := entry.Label
			b.WriteString(fmt.Sprintf("%s%s\n", prefix, label))
		}
	}

	return b.String()
}

func isKittyTerminal() bool {
	term := os.Getenv("TERM")
	// Kitty typically sets TERM to xterm-kitty
	return strings.Contains(term, "kitty")
}

func (m Model) visibleItemsCount() int {
	return m.terminalHeight - 4
}

func (m *Model) filterOptions() {
	m.filteredKeys = []SelectionOption{}

	for id, name := range m.options {
		if strings.Contains(strings.ToLower(name), strings.ToLower(m.filter)) {
			m.filteredKeys = append(m.filteredKeys, SelectionOption{
				Label: name,
				Key:   id,
			})
		}
	}

	// Sort by label
	sort.Slice(m.filteredKeys, func(i, j int) bool {
		return m.filteredKeys[i].Label < m.filteredKeys[j].Label
	})

	// Add quit option at the end
	m.filteredKeys = append(m.filteredKeys, SelectionOption{
		Label: "Quit",
		Key:   "-1",
	})
}

// removed external icat placement logic

// truncateMiddle retained (currently unused, but harmless); removed image rendering
func truncateMiddle(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 10 {
		return s[:max]
	}
	keep := (max - 3) / 2
	return s[:keep] + "..." + s[len(s)-keep:]
}

// DynamicSelect shows a selection menu and returns the selected option
// DynamicSelectWithImages allows passing an id->imageURL map for kitty icat previews.
// Removed DynamicSelectWithImages (CLI image preview feature disabled)

func DynamicSelect(options map[string]string) (SelectionOption, error) {
	config := GetGlobalConfig()
	if config != nil && config.RofiSelection {
		return RofiSelect(options, false)
	}
	model := &Model{
		options:      options,
		filteredKeys: make([]SelectionOption, 0),
	}

	model.filterOptions()
	p := tea.NewProgram(model)

	finalModel, err := p.Run()
	if err != nil {
		return SelectionOption{}, err
	}

	finalSelectionModel, ok := finalModel.(*Model)
	if !ok {
		return SelectionOption{}, fmt.Errorf("unexpected model type")
	}

	if finalSelectionModel.selected < len(finalSelectionModel.filteredKeys) {
		selected := finalSelectionModel.filteredKeys[finalSelectionModel.selected]
		return selected, nil
	}
	return SelectionOption{}, nil
}
