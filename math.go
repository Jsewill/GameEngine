package GameEngine

import (
	mgl64 "github.com/go-gl/mathgl/mgl64"
	mgl32 "github.com/go-gl/mathgl/mgl32"
)

func Mat4dToMat4f(m mgl64.Mat4) (M mgl32.Mat4) {
	for k,v := range m {
		M[k] = float32(v)
	}
	return
}