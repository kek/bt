package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	subscription chan ScanMsg
	devices      []string
	scanDone     bool
}

func initialModel() model {
	deviceScanChannel = make(chan ScanMsg)

	return model{
		subscription: deviceScanChannel,
	}
}

func (m model) nextfun() func() tea.Msg {
	return func() tea.Msg {
		s := <-m.subscription
		return s
	}
}

func (m model) Init() tea.Cmd {
	go StartScan()
	clear := func() tea.Msg {
		return tea.ClearScreen()
	}
	return tea.Sequence(
		clear,
		m.nextfun(),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m, c := m.handleKeyPress(msg)
		return m, c
	case DeviceFound:
		m.devices = append(m.devices, msg.String())
		return m, m.nextfun()
	case ScanDone:
		m.scanDone = true
		return m, nil
	}
	return m, nil
}

func (m model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.String() {
	case "q":
		fallthrough
	case "ctrl+c":
		cmd = tea.Quit
	case "r":
		go StartScan()
		m.devices = []string{}
		m.scanDone = false
		cmd = m.nextfun()
	}
	return m, cmd
}

func (m model) View() string {
	s := ""
	for _, device := range m.devices {
		s += fmt.Sprintf("%s\n", device)
	}
	if m.scanDone {
		s += menu()
	}
	return s
}

func menu() string {
	return "Press 'r' to restart the scan.\n" + "Press 'q' to quit.\n"
}

func main() {
	model := initialModel()
	p := tea.NewProgram(model)
	_, err := p.Run()
	if err != nil {
		fmt.Printf("Error running program: %v\n", err)
	}
}
