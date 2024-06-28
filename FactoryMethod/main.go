package main

import "fmt"

// abstract product
type Button interface {
	onClick()
}

// concrete product
type HTMLButon struct{}

func (b *HTMLButon) onClick() {
	fmt.Println("Clicks on web")
}

// concrete product
type WindownButton struct{}

func (b *WindownButton) onClick() {
	fmt.Println("Clicks on windows")
}

// abstract creator
type Dialog interface {
	newButton() Button
}

// concreate creator
type WindownDialog struct {
}

func (wd WindownDialog) newButton() Button {
	return &WindownButton{}
}

// concreate creator
type WebDialog struct {
}

func (wed *WebDialog) newButton() Button {
	return &HTMLButon{}
}

type Application struct {
	d Dialog
}

func newApp(d Dialog) *Application {
	return &Application{d}
}

func main() {
	d := WindownDialog{}
	app := newApp(d)
	app.d.newButton().onClick()
}
