package component

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type Navbar struct {
	vecty.Core
	Title string       `vecty:"prop"`
	Items []NavbarItem `vecty:"prop"`
}

func (n *Navbar) Render() vecty.ComponentOrHTML {
	l := make(vecty.List, len(n.Items))
	for i := range n.Items {
		l[i] = n.Items[i]
	}

	return elem.Navigation(vecty.Markup(vecty.Class("navbar", "navbar-expand-sm", "navbar-dark", "bg-dark")),
		elem.Anchor(vecty.Markup(vecty.Class("navbar-brand"), prop.Href("#")),
			vecty.Text(n.Title),
		),
		elem.Button(vecty.Markup(vecty.Class("navbar-toggler"), vecty.Attribute("data-toggle", "collapse"), vecty.Attribute("data-target", "#navbarItems")),
			elem.Span(vecty.Markup(vecty.Class("navbar-toggler-icon"))),
		),
		elem.Div(vecty.Markup(vecty.Class("collapse", "navbar-collapse"), prop.ID("navbarItems")),
			elem.UnorderedList(vecty.Markup(vecty.Class("navbar-nav", "mr-auto")),
				l,
			),
		),
	)
}

type NavbarItem interface {
	vecty.ComponentOrHTML
	navbarItem()
}

type NavbarItemLink struct {
	vecty.Core
	Text string `vecty:"prop"`
	Link string `vecty:"prop"`
}

func (ni *NavbarItemLink) Render() vecty.ComponentOrHTML {
	return elem.ListItem(vecty.Markup(vecty.Class("nav-item")),
		elem.Anchor(vecty.Markup(vecty.Class("nav-link"), prop.Href(ni.Link)),
			vecty.Text(ni.Text),
		),
	)
}

func (ni *NavbarItemLink) navbarItem() {}

type NavbarItemDropdown struct {
	vecty.Core
	ID       string       `vecty:"prop"`
	Text     string       `vecty:"prop"`
	Children []LinkedText `vecty:"props"`
}

func (ni *NavbarItemDropdown) Render() vecty.ComponentOrHTML {
	l := make(vecty.List, len(ni.Children))
	for i, child := range ni.Children {
		l[i] = elem.Anchor(vecty.Markup(vecty.Class("dropdown-item"), prop.Href(child.Link)),
			vecty.Text(child.Text),
		)
	}
	return elem.ListItem(vecty.Markup(vecty.Class("nav-item", "dropdown")),
		elem.Anchor(vecty.Markup(vecty.Class("nav-link", "dropdown-toggle"), prop.ID(ni.ID), prop.Href("#"), vecty.Attribute("role", "button"), vecty.Attribute("data-toggle", "dropdown")),
			vecty.Text(ni.Text),
		),
		elem.Div(vecty.Markup(vecty.Class("dropdown-menu")),
			l,
		),
	)
}

func (ni *NavbarItemDropdown) navbarItem() {}
