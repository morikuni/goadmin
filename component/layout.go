package component

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Layout struct {
	vecty.Core
	Navbar  *Navbar
	Content vecty.ComponentOrHTML
}

func (l *Layout) Render() vecty.ComponentOrHTML {
	return elem.Body(
		l.Navbar,
		elem.Div(vecty.Markup(vecty.Class("container")),
			l.Content,
		),
	)
}
