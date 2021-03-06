package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/microcosm-cc/bluemonday"
	"github.com/slimsag/blackfriday"
	router "marwan.io/vecty-router"
)

func main() {
	vecty.SetTitle("Hello LogRocket!")
	vecty.RenderBody(&pageView{})
}

type pageView struct {
	vecty.Core
	Input string
}

// Render implements the vecty.Component interface.
func (p *pageView) Render() vecty.ComponentOrHTML {
	exactMatchOpts := router.NewRouteOpts{ExactMatch: true}
	return elem.Body(
		elem.Div(
			vecty.Markup(vecty.Class("container")),
			router.NewRoute("/", &home{}, exactMatchOpts),
		),
	)
}

// Markdown is a simple component which renders the Input markdown as sanitized
// HTML into a div.
type Markdown struct {
	vecty.Core
	Input string `vecty:"prop"`
}

// Render implements the vecty.Component interface.
func (m *Markdown) Render() vecty.ComponentOrHTML {
	// Render the markdown input into HTML using Blackfriday.
	unsafeHTML := blackfriday.Run([]byte(m.Input))

	// Sanitize the HTML.
	safeHTML := string(bluemonday.UGCPolicy().SanitizeBytes(unsafeHTML))

	// Return the HTML, which we guarantee to be safe / sanitized.
	return elem.Div(
		vecty.Markup(
			vecty.UnsafeHTML(safeHTML),
		),
	)
}
