package selector

import (
	"errors"
	"fmt"
	"strings"

	termbox "github.com/nsf/termbox-go"
)

func printLine(s string, y int) {
	for x, r := range []rune(s) {
		termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.Flush()
}

func clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

type Selector struct {
	Query    string
	Cursor   int
	ViewList []string
	List     []string
}

func New(list []string) *Selector {
	viewList := []string{}

	viewList = append(viewList, list...)

	return &Selector{
		ViewList: viewList,
		List:     list,
	}
}

func (s *Selector) Render() {
	clear()
	printLine(fmt.Sprintf("Query> %v", s.Query), 0)

	for i, e := range s.ViewList {
		if i == s.Cursor {
			printLine("* "+e, i+1)
		} else {
			printLine("  "+e, i+1)
		}
	}
}

func checkChar(r rune) bool {
	if byte(r) < 32 || byte(r) == 0x7f {
		return false
	}
	return true
}

func (s *Selector) execQuery() {
	viewList := []string{}
	for _, e := range s.List {
		if strings.Contains(e, s.Query) {
			viewList = append(viewList, e)
		}
	}
	s.ViewList = viewList
	if len(viewList) <= s.Cursor {
		s.Cursor = len(viewList) - 1
	}
}

func (s *Selector) RunSelector() (string, error) {
	if err := termbox.Init(); err != nil {
		return "", err
	}
	defer termbox.Close()
	if err := termbox.Flush(); err != nil {
		return "", err
	}

	s.Render()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return "", errors.New("escape")
			} else if ev.Key == termbox.KeyEnter {
				if len(s.ViewList) == 0 {
					return "", errors.New("no element")
				}
				return s.ViewList[s.Cursor], nil
			} else if len(s.Query) > 0 && (ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2) {
				s.Query = string([]rune(s.Query)[:len([]rune(s.Query))-1])
			} else if ev.Key == termbox.KeyArrowUp {
				if s.Cursor > 0 {
					s.Cursor--
				} else {
					continue
				}
			} else if ev.Key == termbox.KeyArrowDown {
				s.Cursor++
				if len(s.ViewList) <= s.Cursor {
					s.Cursor--
					continue
				}
			}
			if checkChar(ev.Ch) {
				s.Query += string([]rune{ev.Ch})
			}

			s.execQuery()
			s.Render()

		case termbox.EventResize:

		}
	}

	return "", nil
}
