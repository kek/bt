package main

import (
	"bt/bluetooth"
	"fmt"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	subscription chan bluetooth.ScanMsg
	devices      []bluetooth.Device
	scanDone     bool
	selected     int
	screen       screen
	debugMsg     []string
}

type screen int

const (
	Scan screen = iota
	Connect
)

func initialModel() model {
	return model{
		subscription: bluetooth.CreateChannel(),
		screen:       Scan,
		debugMsg:     nil,
	}
}

func (m model) Init() tea.Cmd {
	go bluetooth.StartScan()

	clear := func() tea.Msg {
		return tea.ClearScreen()
	}
	return tea.Sequence(
		clear,
		m.nextScanResult(),
	)
}

type connectToDevice struct {
	device bluetooth.Device
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m, cmd = m.handleKeyPress(msg)
	case bluetooth.Device:
		m.devices = append(m.devices, msg)
		cmd = m.nextScanResult()
	case bluetooth.ScanDone:
		m.scanDone = true
	case connectToDevice:
		m.screen = Connect
		m.debug(msg.device.Identifier)
	}
	return m, cmd
}

func (m *model) debug(identifier string) {
	m.debugMsg = append(m.debugMsg, identifier)
}

func (m model) handleKeyPress(msg tea.KeyMsg) (model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.String() {
	case "q":
		fallthrough
	case "ctrl+c":
		cmd = tea.Quit
	case "r":
		m.subscription = bluetooth.CreateChannel()
		m.devices = []bluetooth.Device{}
		m.scanDone = false
		m.selected = 0
		m.screen = Scan
		cmd = m.nextScanResult()
		go bluetooth.StartScan()
	case "j":
		fallthrough
	case "down":
		if m.selected < len(m.devices)-1 {
			m.selected++
		}
	case "k":
		fallthrough
	case "up":
		if m.selected > 0 {
			m.selected--
		}
	case "enter":
		if m.selected < len(m.devices) {
			cmd = func() tea.Msg {
				return connectToDevice{device: m.devices[m.selected]}
			}
		}
	}
	return m, cmd
}

func (m model) nextScanResult() func() tea.Msg {
	return func() tea.Msg {
		msg := <-m.subscription
		return msg
	}
}

func (m model) View() string {
	s := ""
	switch m.screen {
	case Scan:
		if len(m.devices) > 0 {
			s += fmt.Sprintf("Found %d devices.\n", len(m.devices))
		}
		if m.scanDone {
			s += menu()
		}
		for i, device := range m.devices {
			if i == m.selected {
				s += fmt.Sprintf("-> %s\n", device)
			} else {
				s += fmt.Sprintf("   %s\n", device)
			}
		}
	case Connect:
		s += "Connecting to "
		if m.selected < len(m.devices) {
			s += m.devices[m.selected].Identifier
		}
	default:
		s += "Unknown screen"
	}
	s += "\n\n"
	log := slices.Clone(last(m.debugMsg, 5))
	slices.Reverse(log)
	for _, msg := range log {
		s += fmt.Sprintf("%s\n", msg)
	}
	return s
}

func last[T any](slice []T, n int) []T {
	if len(slice) < n {
		return slice
	}
	return slice[len(slice)-n:]
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
