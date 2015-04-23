/*
Shader implements...

*/

package GameEngine

import (
	"github.com/go-gl-legacy/gl"
	"io/ioutil"
	"runtime"
)

type Shader struct {
	gl.Shader
	Type   gl.GLenum
	Source string
}

//Use TBD over this method--may be unexported in the future; Reserves the shader resource on the GPU
func (s *Shader) Register() error {
	s.Create()

	if !gl.Object(s.Shader).IsShader() {
		err := &OGLShaderError{OGLShader: s}
		return err
	}
	runtime.SetFinalizer(&s, func(s *Shader) { s.Delete() })

	return nil
}

func (s *Shader) Create() {
	s.Shader = gl.CreateShader(s.Type)
}

func (s *Shader) SourceFile(p string) error {
	source, err := ioutil.ReadFile(p)
	s.Source = string(source)
	if err != nil {
		//erro := new(OGLShaderError)
		//erro.Log(s, err.Error())
		return err //erro
	}

	return nil
}
