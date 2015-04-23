/*
Object implements...

*/

package GameEngine

import (
	mgl "github.com/go-gl/mathgl/mgl64"
)

type Object struct {
	Meshes []*Mesh
	Translation mgl.Mat4
	Rotation mgl.Mat4
	Scale mgl.Mat4
}

func(o *Object) Init() {
	o.Translation, o.Rotation, o.Scale = mgl.Ident4(), mgl.Ident4(), mgl.Ident4()
}

func(o *Object) NewMesh() (m *Mesh) {
	m = new(Mesh)
	o.Meshes = append(o.Meshes, m)
	return
}