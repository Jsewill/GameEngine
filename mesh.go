/*
Mesh implements...

*/

package GameEngine

import (
	"github.com/go-gl-legacy/gl"
)

type Mesh struct {
	Programs []*Program
	Vertices []Vertex
	Indices  []uint32
}

func (m *Mesh) NewProgram() (p *Program) {
	p = new(Program)
	p.Program = gl.CreateProgram()
	m.Programs = append(m.Programs, p)
	return
}

func (m *Mesh) CalculateNormals() {
	for k := range m.Vertices {
		m.Vertices[k].CalculateNormal(m)
	}
}
