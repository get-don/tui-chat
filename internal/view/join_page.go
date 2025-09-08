package view

import "github.com/rivo/tview"

type joinPage struct {
	app       *tview.Application
	form      *tview.Form
	onConnect func()
	onQuit    func()
}

func newJoinPage(app *tview.Application, onConnect func(), onQuit func()) *joinPage {
	page := &joinPage{
		app:       app,
		onConnect: onConnect,
		onQuit:    onQuit,
	}

	page.form = tview.NewForm().
		AddInputField("Server", "localhost:8000", 20, nil, nil).
		AddInputField("Name", "", 20, nil, nil).
		AddButton("접속", page.connect).
		AddButton("종료", page.quit)

	page.form.SetBorder(true).SetTitle(" Connection ")

	return page
}

func (p *joinPage) page() tview.Primitive {
	return p.form
}

func (p *joinPage) connect() {
	//serverInput := p.form.GetFormItemByLabel("Server").(*tview.InputField)
	//nameInput := p.form.GetFormItemByLabel("Name").(*tview.InputField)
	//serverAddr := serverInput.GetText()
	//name := nameInput.GetText()
	//
	// 서버 연결, 이름 설정
	//p.showErrorModal()

	if p.onConnect != nil {
		p.onConnect()
	}
}

func (p *joinPage) quit() {
	if p.onQuit != nil {
		p.onQuit()
	}

	p.app.Stop()
}

func (p *joinPage) showErrorModal() {
	modal := tview.NewModal().
		// SetText("서버 연결에 실패: " + err.Error()).
		SetText("서버 연결 실패").
		AddButtons([]string{"확인"})

	p.app.SetRoot(modal, false).SetFocus(modal)

	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "확인" {
			p.app.SetRoot(p.page(), true).SetFocus(p.form)
		}
	})
}
