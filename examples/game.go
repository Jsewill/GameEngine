package main

import (
	"fmt"
	ge "github.com/Jsewill/GameEngine"
	"github.com/go-gl-legacy/gl"
	glfw "github.com/go-gl/glfw3"
	mgl "github.com/go-gl/mathgl/mgl64"
	"log"
	"os"
	"path/filepath"
	"unsafe"
)

func main() {
	e := new(ge.Engine)
	settings := ge.DefaultSettings
	settings.Windows[0].LoopCallback = func(w *ge.Window, e *ge.Engine) {

		gl.Enable(gl.DEPTH_TEST)
		gl.DepthFunc(gl.LESS)

		camera := e.NewCamera()
		camera.Translation = camera.Translation.Mul4(mgl.Translate3D(0.0, -5.0, 10))
		camera.Rotation = camera.Rotation.Mul4(mgl.HomogRotate3DX(mgl.DegToRad(27.5)))

		e.ActivateCamera(camera)
		/*object := e.NewObject()
		mesh := object.NewMesh()
		mesh.Vertices = append(mesh.Vertices,
			//Bottom
			ge.Vertex{Position:mgl.Vec3{-0.5, -0.5, 0.0}, Index:0},
			ge.Vertex{Position:mgl.Vec3{-0.5, 0.5, 0.0}, Index:1},
			ge.Vertex{Position:mgl.Vec3{0.5, 0.5, 0.0}, Index:2},
			ge.Vertex{Position:mgl.Vec3{0.5, -0.5, 0.0}, Index:3},
		)
		mesh.Indices = append(mesh.Indices,
			0,1,2,
			0,2,3,
		)*/

		cwd, err := os.Getwd()
		e.FromColladaFile(fmt.Sprintf("%v/%v", cwd, "shape.dae"))
		object := e.Objects[0]
		mesh := object.Meshes[0]

		//mesh.CalculateNormals()
		//log.Println(mesh.Vertices)

		program := mesh.NewProgram()
		defer program.Delete()
		vshader := program.NewShader(gl.VERTEX_SHADER)
		defer vshader.Delete()
		fshader := program.NewShader(gl.FRAGMENT_SHADER)
		defer fshader.Delete()

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}

		err = vshader.SourceFile(dir + "/shaders/vertex.glsl")
		if err != nil {
			log.Fatalln(err)
		}

		err = fshader.SourceFile(dir + "/shaders/fragment.glsl")
		if err != nil {
			log.Fatalln(err)
		}

		program.MakeProgram()
		ge.CheckGLError()
		program.Use()
		log.Println(program.GetInfoLog())
		log.Println(gl.GetError())

		pu := program.GetUniformLocation("Projection")
		vu := program.GetUniformLocation("View")
		ru := program.GetUniformLocation("Rotation")
		tu := program.GetUniformLocation("Translation")
		su := program.GetUniformLocation("Scale")

		vao := gl.GenVertexArray()
		vbo := gl.GenBuffer()
		ebo := gl.GenBuffer()

		defer vao.Delete()
		defer vbo.Delete()
		defer ebo.Delete()

		vao.Bind()
		vbo.Bind(gl.ARRAY_BUFFER)
		gl.BufferData(gl.ARRAY_BUFFER, len(mesh.Vertices)*int(unsafe.Sizeof(ge.Vertex{})), mesh.Vertices, gl.STATIC_DRAW)

		positionAttrib := program.GetAttribLocation("position")
		colorAttrib := program.GetAttribLocation("color")
		normalAttrib := program.GetAttribLocation("normal")

		positionAttrib.EnableArray()
		defer positionAttrib.DisableArray()

		colorAttrib.EnableArray()
		defer colorAttrib.DisableArray()

		normalAttrib.EnableArray()
		defer normalAttrib.DisableArray()

		positionAttrib.AttribPointer(3, gl.DOUBLE, false, int(unsafe.Sizeof(ge.Vertex{})), nil)
		colorAttrib.AttribPointer(4, gl.DOUBLE, false, int(unsafe.Sizeof(ge.Vertex{})), &mesh.Vertices[0].Color)
		normalAttrib.AttribPointer(3, gl.DOUBLE, false, int(unsafe.Sizeof(ge.Vertex{})), &mesh.Vertices[0].Normal)

		ebo.Bind(gl.ELEMENT_ARRAY_BUFFER)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(mesh.Indices)*int(unsafe.Sizeof(mesh.Indices[0])), mesh.Indices, gl.STATIC_DRAW)

		gl.ClearColor(0.333, 0.333, 0.333, 1.0)
		for !w.ShouldClose() {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
			//Do Stuff

			object.Rotation = object.Rotation.Mul4(
				mgl.HomogRotate3DZ(mgl.DegToRad(0.5)).Mul4(
					mgl.HomogRotate3DX(mgl.DegToRad(0.3)).Mul4(
						mgl.HomogRotate3DZ(mgl.DegToRad(0.2)),
					),
				),
			)

			pu.UniformMatrix4fv(false, ge.Mat4dToMat4f(e.ActiveCamera.ProjectionMatrix()))
			vu.UniformMatrix4fv(false, ge.Mat4dToMat4f(e.ActiveCamera.ViewMatrix()))
			//mu.UniformMatrix4fv(false, Model)
			ru.UniformMatrix4fv(false, ge.Mat4dToMat4f(object.Rotation))
			tu.UniformMatrix4fv(false, ge.Mat4dToMat4f(object.Translation))
			su.UniformMatrix4fv(false, ge.Mat4dToMat4f(object.Scale))

			gl.DrawElements(gl.TRIANGLES, len(mesh.Indices), gl.UNSIGNED_INT, nil)

			w.SwapBuffers()
			glfw.PollEvents()
		}
	}

	e.Settings = &settings //Initialize with default settings; Unnecessary, just to show it works
	e.Init()
	defer e.Finish()
	err := e.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
