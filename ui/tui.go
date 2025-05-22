package ui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zeshi09/ipenrich/internal/enrich"
	// "github.com/zeshi09/ipenrich/cmd"
)

var (
	headerStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))
	progressStyle = progress.New(progress.WithDefaultGradient())
	rowStyle      = lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("252"))
	cellStyle     = lipgloss.NewStyle().Padding(0, 2).Foreground(lipgloss.Color("249"))
	tableHeader   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("245")).Render
	tableRow      = lipgloss.NewStyle().Foreground(lipgloss.Color("250")).Render
	abuseLow      = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))  // 🟢 зелёный
	abuseMid      = lipgloss.NewStyle().Foreground(lipgloss.Color("214")) // 🟠 жёлтый
	abuseHigh     = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))   // 🔴 красный
	vtLow         = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))  // green
	vtMid         = lipgloss.NewStyle().Foreground(lipgloss.Color("214")) // yellow
	vtHigh        = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))   // red

)

func InitialModel(logPath string, ips []string) tea.Model {
	vp := viewport.New(160, 30)
	vp.SetContent("[Enriching IPs...]")
	return model{
		logFile:  logPath,
		progress: progressStyle,
		percent:  0.0,
		ips:      ips,
		index:    0,
		viewport: vp,
		geo:      make(map[string]string),
		abuse:    make(map[string]string),
		vt:       make(map[string]string),
	}
}

type model struct {
	logFile  string
	progress progress.Model
	percent  float64
	done     bool
	index    int
	ips      []string
	viewport viewport.Model
	geo      map[string]string
	abuse    map[string]string
	vt       map[string]string
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.done {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "q":
				return m, tea.Quit
			}
			m.viewport, _ = m.viewport.Update(msg)
		}
		return m, nil
	}

	ip := m.ips[m.index]
	m.geo[ip] = enrich.FetchGeoInfo(ip)
	m.abuse[ip] = enrich.FetchAbuseScore(ip)
	m.vt[ip] = enrich.FetchVTStats(ip)
	m.index++
	m.percent = float64(m.index) / float64(len(m.ips))
	if m.index >= len(m.ips) {
		m.percent = 1.0
		m.done = true
		m.viewport.SetContent(renderIPTable(m.ips, m.geo, m.abuse, m.vt))
		return m, nil
	}
	return m, tick()
}

func (m model) View() string {
	if m.done {
		return headerStyle.Render("Completed! 🎉") +
			headerStyle.Render(fmt.Sprintf("\n%d IP's was found\n\n", len(m.ips))) +
			"\n" +
			m.viewport.View() +
			"\nYou can scroll with j,k keys, q to quit\n"
	}

	bar := m.progress.ViewAs(m.percent)
	status := fmt.Sprintf(
		"%s\n\n File: %s\n Preparation: %d / %d\n\n%s",
		headerStyle.Render("[IPEnricher TUI]"),
		m.logFile,
		m.index,
		len(m.ips),
		bar,
	)
	return status
}

func tick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return struct{}{}
	})
}

func padRight(str string, length int) string {
	visibleLen := lipgloss.Width(str)
	if visibleLen >= length {
		return str
	}
	return str + strings.Repeat(" ", length-visibleLen)
}

func renderIPTable(ips []string, geo map[string]string, abuse map[string]string, vt map[string]string) string {
	var b strings.Builder
	b.WriteString(tableHeader(fmt.Sprintf("%-3s  │ %-15s │ %-30s │ %-30s │ %-7s │ %-9s │", "#", "IP", "Location", "Organization", "Abuse", "VTStats")))
	header := fmt.Sprintf("%-3s  │ %-15s │ %-30s │ %-30s │ %-7s │ %-9s │", "#", "IP", "Location", "Organization", "Abuse", "VTStats")
	b.WriteString("\n" + strings.Repeat("─", lipgloss.Width(header)) + "\n")

	for i, ip := range ips {
		parts := strings.SplitN(geo[ip], "(", 2)
		loc := strings.TrimSpace(parts[0])
		org := ""
		if len(parts) > 1 {
			org = strings.TrimRight(strings.TrimSpace(parts[1]), ")")
		}

		vt := vt[ip]
		coloredVT := vt
		if parts := strings.SplitN(vt, "/", 2); len(parts) == 2 {
			if count, err := strconv.Atoi(parts[0]); err == nil {
				switch {
				case count == 0:
					coloredVT = vtLow.Render(vt)
				case count <= 4:
					coloredVT = vtMid.Render(vt)
				default:
					coloredVT = vtHigh.Render(vt)
				}
			}
		}

		score := strings.TrimSpace(abuse[ip])
		coloredScore := score
		if s, err := strconv.Atoi(score); err == nil {
			switch {
			case s >= 80:
				coloredScore = abuseHigh.Render(score)
			case s >= 40:
				coloredScore = abuseMid.Render(score)
			default:
				coloredScore = abuseLow.Render(score)
			}
		}
		// row := fmt.Sprintf("%-3d  │ %-15s │ %-30s │ %-30s │ %-7s │ %-9s │", i+1, ip, loc, org, coloredScore, coloredVT)
		row := fmt.Sprintf("%-3d  │ %s │ %s │ %s │ %s │ %s │",
			i+1,
			padRight(ip, 15),
			padRight(loc, 30),
			padRight(org, 30),
			padRight(coloredScore, 7),
			padRight(coloredVT, 9),
		)

		b.WriteString(tableRow(row) + "\n")
	}
	return b.String()
}
