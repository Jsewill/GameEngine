/*
Window implements...

*/
package GameEngine

import (
	mgl "github.com/go-gl/mathgl/mgl64"
	glfw "github.com/go-gl/glfw3"
)

type KeyCallback func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)

type LoopCallback func(w *Window, e *Engine)

type Window struct {
	*glfw.Window
	*Monitor
	KeyCallbacks []KeyCallback
	LoopCallback
	Title string
	Width int
	Height int
	ClearColor mgl.Vec4
}

func(w *Window) SetKeyCallbackRange() {
	w.Window.SetKeyCallback(
		func(win *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
			for _,f := range w.KeyCallbacks {
				f(win, key, scancode, action, mods)
			}
		},
	)
}

//Horizontal Field of View
func(w *Window) Aspect() float64 {
	return float64(w.Width)/float64(w.Height)
}