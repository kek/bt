package main

import (
	"bt/bluetooth"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	subscription chan bluetooth.ScanMsg
	devices      []string
	scanDone     bool
}

func initialModel() model {
	ch := bluetooth.CreateChannel()

	return model{
		subscription: ch,
	}
}

func (m model) nextfun() func() tea.Msg {
	return func() tea.Msg {
		msg := <-m.subscription
		return msg
	}
}

func (m model) Init() tea.Cmd {
	go bluetooth.StartScan()
	clear := func() tea.Msg {
		return tea.ClearScreen()
	}
	return tea.Sequence(
		clear,
		m.nextfun(),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m, cmd = m.handleKeyPress(msg)
	case bluetooth.DeviceFound:
		m.devices = append(m.devices, msg.String())
		cmd = m.nextfun()
	case bluetooth.ScanDone:
		m.scanDone = true
	}
	return m, cmd
}

func (m model) handleKeyPress(msg tea.KeyMsg) (model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.String() {
	case "q":
		fallthrough
	case "ctrl+c":
		cmd = tea.Quit
	case "r":
		go bluetooth.StartScan()
		m.devices = []string{}
		m.scanDone = false
		cmd = m.nextfun()
	}
	return m, cmd
}

func (m model) View() string {
	s := ""
	if len(m.devices) > 0 {
		s += fmt.Sprintf("Found %d devices.\n", len(m.devices))
	}
	if m.scanDone {
		s += menu()
	}
	for _, device := range m.devices {
		s += fmt.Sprintf("%s\n", device)
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
