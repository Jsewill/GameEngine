/*
Program implements...

*/

package GameEngine

import (
	"github.com/go-gl-legacy/gl"
	"log"
)

type Program struct {
	gl.Program
	Shaders []*Shader
}

//Use TBD over this method--may be unexported in the future; Reserves the program resource on the GPU
func (p *Program) Register() error {
	p.Program = gl.CreateProgram()

	if !gl.Object(p.Program).IsProgram() {
		err := &OGLProgramError{OGLProgram: p}
		err.Details = trace()
		return err
	}

	return nil
}

func (p *Program) NewShader(t gl.GLenum) (s *Shader) {
	s = &Shader{Type: t}
	s.Create()
	p.Shaders = append(p.Shaders, s)
	return
}

func (p *Program) MakeProgram() {
	for _, shader := range p.Shaders {
		shader.Shader.Source(shader.Source)
		shader.Compile()
		CheckGLError()

		if shader.Get(gl.COMPILE_STATUS) != 1 {
			log.Println(shader.GetInfoLog())
		}

		p.AttachShader(shader.Shader)
		CheckGLError()
	}

	p.Link()

	if p.Get(gl.LINK_STATUS) != 1 {
		log.Println(p.GetInfoLog())
		CheckGLError()
	}
}
