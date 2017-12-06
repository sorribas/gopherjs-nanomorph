package nanomorph

import "github.com/gopherjs/gopherjs/js"

type App interface {
	Render() string
	OnChange(func())
}

type BaseApp struct {
	ReRender func()
}

func (app *BaseApp) OnChange(fn func()) {
	fn()
	app.ReRender = fn
}

func MountApp(app App) {
	html := app.Render()
	node := js.Global.Get("gopherjsNanomorphDeps").Call("htmlToNode", html)
	js.Global.Get("document").Get("body").Call("appendChild", node)

	rerender := func() {
		newNode := js.Global.Get("gopherjsNanomorphDeps").Call("htmlToNode", app.Render())
		js.Global.Get("gopherjsNanomorphDeps").Call("nanomorph", node, newNode)
	}
	app.OnChange(rerender)
}
