package view

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type chatPage struct {
	app         *tview.Application
	userView    *tview.TextView
	messageView *tview.TextView
	inputField  *tview.InputField
	mainFlex    *tview.Flex
}

func newChatPage(app *tview.Application) *chatPage {
	page := &chatPage{
		app: app,
	}

	page.makeUserView()
	page.makeMessageView()
	page.makeInputField()
	page.makeFlex()

	return page
}

func (p *chatPage) page() tview.Primitive {
	return p.mainFlex
}

func (p *chatPage) makeUserView() {
	p.userView = tview.NewTextView()
	p.userView.SetBorder(true).SetTitle("User List")
}

func (p *chatPage) makeMessageView() {
	p.messageView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetChangedFunc(func() {
			p.app.Draw()
		}).SetMaxLines(100)

	p.messageView.SetBorder(true)
}

func (p *chatPage) makeInputField() {
	p.inputField = tview.NewInputField().SetLabel("Message: ").SetPlaceholder("Enter your message...")
	p.inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			if p.inputField.GetText() != "" {
				fmt.Fprintf(p.messageView, "[pink]You: [white]%s\n", p.inputField.GetText())
				p.messageView.ScrollToEnd()
				p.inputField.SetText("")
			}
		}
	})

	// Message view Scroll
	// Note: makeMessageView()가 먼저 호출되어야 함
	p.inputField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Modifiers() == tcell.ModCtrl {
			switch event.Key() {
			case tcell.KeyUp:
				row, _ := p.messageView.GetScrollOffset()
				newRow := row - 1
				if newRow < 0 {
					newRow = 0
				}
				p.messageView.ScrollTo(newRow, 0)
				return nil
			case tcell.KeyPgUp:
				row, _ := p.messageView.GetScrollOffset()
				newRow := row - 10
				if newRow < 0 {
					newRow = 0
				}
				p.messageView.ScrollTo(newRow, 0)
				return nil
			case tcell.KeyPgDn:
				row, _ := p.messageView.GetScrollOffset()
				p.messageView.ScrollTo(row+10, 0)
				return nil
			case tcell.KeyDown:
				row, _ := p.messageView.GetScrollOffset()
				p.messageView.ScrollTo(row+1, 0)
				return nil
			default:
				break
			}
		}

		return event
	})

	p.inputField.SetBorder(true)
}

func (p *chatPage) makeFlex() {
	p.mainFlex = tview.NewFlex().SetDirection(tview.FlexColumn)

	// 화면 왼쪽
	p.mainFlex.AddItem(p.userView, 0, 1, false)

	// 화면 오른쪽
	rightFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(p.messageView, 0, 9, false).
		AddItem(p.inputField, 0, 1, true)

	p.mainFlex.AddItem(rightFlex, 0, 2, false)
}
