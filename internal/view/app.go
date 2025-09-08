package view

import (
	"github.com/rivo/tview"
)

type View struct {
	app      *tview.Application
	pages    *tview.Pages
	joinPage *joinPage
	chatPage *chatPage
}

func NewView() *View {
	view := &View{}

	view.app = tview.NewApplication()
	view.pages = tview.NewPages()

	onConnect := func() {
		view.pages.SwitchToPage("chatPage")
		view.app.SetFocus(view.chatPage.inputField)
	}

	view.joinPage = newJoinPage(view.app, onConnect, nil)
	view.chatPage = newChatPage(view.app)

	view.pages.AddPage("joinPage", view.joinPage.page(), true, true)
	view.pages.AddPage("chatPage", view.chatPage.page(), true, false)

	return view
}

func (v *View) Show() {
	if err := v.app.SetRoot(v.pages, true).Run(); err != nil {
		panic(err)
	}
}
