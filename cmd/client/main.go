package main

import (
	"github.com/get-don/tui-chat/internal/view"
)

func main() {
	if u := view.NewView(); u != nil {
		u.Show()
	}
}
