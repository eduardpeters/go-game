package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/eduardpeters/go-game/internal/game"
)

type model struct {
	g       *game.Game
	cursorX int
	cursorY int
	width   int
	height  int
	msg     string
}

func NewModel() tea.Model {
	g := game.NewGame(9)
	return model{
		g:       g,
		width:   g.Size,
		height:  g.Size,
		cursorX: 0,
		cursorY: 0,
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "n":
			m.g = game.NewGame(9)
			m.msg = fmt.Sprintf("New %[1]dx%[1]d game created.", 9)
			m.cursorX, m.cursorY = 0, 0
		case "up", "k":
			if m.cursorY > 0 {
				m.cursorY--
			}
		case "down", "j":
			if m.cursorY < m.height-1 {
				m.cursorY++
			}
		case "left", "h":
			if m.cursorX > 0 {
				m.cursorX--
			}
		case "right", "l":
			if m.cursorX < m.width-1 {
				m.cursorX++
			}
		case "enter":
			posX, posY := m.cursorX, m.cursorY
			err := m.g.PlaceStone(posX, posY)
			if err != nil {
				switch err {
				case game.ErrNotEmpty:
					m.msg = fmt.Sprintf("Illegal move at %d,%d - Intersection not empty", posX, posY)
				default:
					m.msg = fmt.Sprintf("Unable to place stone at %d,%d - Unexpected error", posX, posY)
				}
			} else {
				m.msg = fmt.Sprintf("Placed stone at %d,%d", posX, posY)
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	b.WriteString("Go Board\n\n")
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			idx := y*m.width + x
			ch := "."
			switch m.g.Board[idx] {
			case game.BLACK:
				ch = "●"
			case game.WHITE:
				ch = "○"
			}
			if x == m.cursorX && y == m.cursorY {
				b.WriteString("[" + ch + "]")
			} else {
				b.WriteString(" " + ch + " ")
			}
		}
		b.WriteByte('\n')
	}
	b.WriteString("\n")
	b.WriteString(m.msg)
	b.WriteString("\n\n")
	b.WriteString("Keys: arrows / hjkl move • enter toggles a black stone • n new game • q quit\n")
	return b.String()
}
