package main

import (
	"cli-based-chat/internal/view"
)

func main() {
	if u := view.NewView(); u != nil {
		u.Show()
	}
}
