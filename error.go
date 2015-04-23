/*
Error implements...

*/
package GameEngine

import (
	"fmt"
	"github.com/go-gl-legacy/gl"
	"github.com/go-gl/glu"
	"log"
	"runtime"
	"time"
)

type GeneralError struct {
	Time    time.Time
	Details string
}

func (e *GeneralError) Error() string {
	return fmt.Sprintf("%+v\n", e)
}

func (e *GeneralError) Log(d string) {
	e.Time = time.Now()
	e.Details = fmt.Sprint(trace(), d)
}

type OGLError struct {
	*GeneralError
}

type OGLObjectError struct {
	*GeneralError
	OGLObject *gl.Object
}

type OGLProgramError struct {
	*GeneralError
	OGLProgram *Program
}

type OGLShaderError struct {
	*GeneralError
	OGLShader *Shader
}

func (e *OGLShaderError) Log(s *Shader, d string) {
	e.OGLShader = s
	e.GeneralError = new(GeneralError)
	e.GeneralError.Log(d)
}

func CheckGLError() {
	err := gl.GetError()
	if err != gl.NO_ERROR {
		log.Println(glu.ErrorString(uint(err)))
	}
}

func trace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return fmt.Sprintf("%s:%d %s\n", file, line, f.Name())
}
