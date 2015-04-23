/*
Engine implements...

*/
package GameEngine

import (
	"github.com/go-gl-legacy/gl"
	glfw "github.com/go-gl/glfw3"
	mgl "github.com/go-gl/mathgl/mgl64"
	"log"
	"runtime"
)

var (
	DefaultWindowHints = map[glfw.Hint]int{
		glfw.ContextVersionMajor:     3,
		glfw.ContextVersionMinor:     3,
		glfw.OpenglForwardCompatible: glfw.True,
		glfw.OpenglProfile:           glfw.OpenglCoreProfile,
	}

	DefaultSettings = Settings{
		Vsync: 1,
		Windows: []Window{
			Window{
				KeyCallbacks: []KeyCallback{
					DefaultKeyCallback,
				},
				LoopCallback: DefaultLoopCallback,
				Title:        "GameEngine",
			},
		},
	}
)

//Default Loop Function
func DefaultLoopCallback(w *Window, e *Engine) {

	for !w.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		//Do Stuff

		w.SwapBuffers()
		glfw.PollEvents()
	}
}

//Default Key Handling Function
func DefaultKeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

//Main Driver
type Engine struct {
	Settings *Settings
	Monitors []Monitor
	Windows  []Window
	Objects  []*Object
	Cameras  []*Camera

	ActiveCamera *Camera

	InitCallBacks []InitCallBack

	WindowHints map[glfw.Hint]int
}

func (e *Engine) Init() error {
	//Lock to one thread; It seems OpenGL requires all calls to come from one thread
	runtime.LockOSThread()

	e.init(e.Settings)

	//Init GLFW or exit
	if !glfw.Init() {
		err := &GeneralError{}
		err.Log("Could not initialize GLFW.")
		return err
	}

	//Vsync
	glfw.SwapInterval(e.Settings.Vsync)

	//Get all monitors
	monitors, err := glfw.GetMonitors()
	if err != nil {
		return err
	}
	e.Monitors = []Monitor{}
	for _, m := range monitors {
		e.Monitors = append(e.Monitors, Monitor{m})
	}

	//Get video mode; Assume primary is at first index
	pvm, err := e.Monitors[0].GetVideoMode()
	if err != nil {
		return err
	}

	//Hints for window creation
	e.Hint()

	e.Windows = append(e.Windows, e.Settings.Windows...)
	//Check for existing windows
	if len(e.Windows) == 0 {
		//Create main window if none has been specified
		mainWindow := Window{
			KeyCallbacks: []KeyCallback{
				DefaultKeyCallback,
			},
			LoopCallback: DefaultLoopCallback,
			Monitor:      &e.Monitors[0],
			Title:        "GameEngine",
			Width:        pvm.Width,
			Height:       pvm.Height,
		}
		e.Windows = append(e.Windows, mainWindow)
	}

	return nil
}

func (e *Engine) Run() error {

	//Get video mode; Assume primary is at first index
	pvm, err := e.Monitors[0].Monitor.GetVideoMode()
	if err != nil {
		return err
	}

	//Loop over and create windows
	for k, _ := range e.Windows {
		if e.Windows[k].Monitor == nil {
			e.Windows[k].Monitor = &e.Monitors[0]
			if e.Windows[k].Height == 0 {
				e.Windows[k].Height = pvm.Height
			}
			if e.Windows[k].Width == 0 {
				e.Windows[k].Width = pvm.Width
			}
		}
		if e.Windows[k].Title == "" {
			e.Windows[k].Title = "GameEngine"
		}

		//Create GLFW Window
		window, err := glfw.CreateWindow(e.Windows[k].Width, e.Windows[k].Height, e.Windows[k].Title, e.Windows[k].Monitor.Monitor, nil)
		if err != nil {
			return err
		} else {
			e.Windows[k].Window = window
			defer e.Windows[0].Destroy()
		}

		//Default Key Callback
		if len(e.Windows[k].KeyCallbacks) == 0 {
			e.Windows[k].KeyCallbacks = append(
				e.Windows[k].KeyCallbacks,
				DefaultKeyCallback,
			)
		}

		//Default Loop Callback
		if e.Windows[k].LoopCallback == nil {
			e.Windows[k].LoopCallback = DefaultLoopCallback
		}
	}

	//Set up on window at first index
	e.Windows[0].MakeContextCurrent()

	//Need access to window handling here

	//Initialize GLEW
	glewStatus := gl.Init()
	if glewStatus != 0 {
		err := &OGLError{}
		err.Log("Could not initialize GLFW.")
		return err
	}

	log.Println(gl.GetError())

	//Open all windows
	for _, w := range e.Windows {
		gl.ClearColor(gl.GLclampf(w.ClearColor[0]), gl.GLclampf(w.ClearColor[1]), gl.GLclampf(w.ClearColor[2]), gl.GLclampf(w.ClearColor[3]))
		//Key Callback Functions
		w.SetKeyCallbackRange()

		//Loop!
		w.LoopCallback(&w, e)
	}

	return nil
}

func (e *Engine) Finish() {
	glfw.Terminate()
}

func (e *Engine) Hint() {
	if len(e.WindowHints) == 0 {
		e.WindowHints = DefaultWindowHints
	}
	for hintType, hintValue := range e.WindowHints {
		glfw.WindowHint(hintType, hintValue)
	}
}

func (e *Engine) init(s *Settings) {
	if &s != nil {
		e.Settings = &DefaultSettings
	}
}

func (e *Engine) NewObject() (o *Object) {
	o = new(Object)
	o.Init()
	e.Objects = append(e.Objects, o)
	return
}

func (e *Engine) NewCamera() (c *Camera) {
	c = new(Camera)
	c.Init()
	//@TODO Should set up e.ActiveWindow functionality; set to first window for now
	w := e.Windows[0]

	//Perspective Aspect ratio from window dimensions
	c.Aspect = w.Aspect()

	e.Cameras = append(e.Cameras, c)
	return
}

func (e *Engine) ActivateCamera(c *Camera) {
	e.ActiveCamera = c
}

func (e *Engine) ActivateCameraIndex(i int) *Camera {
	if len(e.Cameras) > i {
		e.ActiveCamera = e.Cameras[i]
	} else {
		log.Println("Camera index is out of range.")
		/*if e.ActiveCamera == nil {
			if len(e.Cameras) > 0 {
				e.ActiveCamera = e.Cameras[len(e.Cameras)-1]
			} else {
				e.ActiveCamera = e.NewCamera()
			}
		}*/
	}

	return e.ActiveCamera
}

func (e *Engine) FromColladaFile(filepath string) {
	c := ImportColladaFile(filepath)
	e.FromCollada(c)
}

func (e *Engine) FromCollada(c *ColladaDoc) {
	geometries := c.GeometriesLibraries[0].Geometries

	for _, geo := range geometries {
		object := e.NewObject()
		mesh := object.NewMesh()
		for sk, src := range geo.Mesh.Sources {
			values := src.FloatArray.ToTListOfFloats().Values()
			end := (len(values) - 1) / 3
			switch {
			//Positions; Should be first
			case sk == 0:
				for i := 1; i < end; i++ {
					mesh.Vertices = append(
						mesh.Vertices,
						Vertex{
							Position: mgl.Vec3{
								float64(values[i*3-3]),
								float64(values[i*3-2]),
								float64(values[i*3-1]),
							},
							Index: uint32(i - 1),
						},
					)
				}
			}
		}
		for _, plist := range geo.Mesh.Polylists {
			pvalues := plist.P.Values()
			modder := len(plist.Inputs)
			for pk, pv := range pvalues {
				if pk%modder == 0 {
					mesh.Indices = append(mesh.Indices, uint32(pv))
				}
			}
		}
	}
}

type InitCallBack func(e *Engine)

type Settings struct {
	Vsync   int
	Windows []Window
}
