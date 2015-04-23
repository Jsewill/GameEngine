package GameEngine

import (
	_ "github.com/go-gl-legacy/gl"
	mgl "github.com/go-gl/mathgl/mgl64"
	"math"
)

const (
	ORTHOGRAPHIC = iota
	PERSPECTIVE  = iota
)

type Camera struct {
	Object
	ProjectionType uint

	//Projection Properties
	Far, Near,

	//Orthographic Properties
	Left, Right, Bottom, Top,

	//Perspective Properties
	Aspect, SensorHeight, SensorWidth, FocalLength float64
}

func (c *Camera) Init() {
	c.Object.Init()
	c.ProjectionType = PERSPECTIVE

	//Projection Type Neutral Variables
	c.Near, c.Far = 0.001, 1000.0

	//Orthographic
	c.Left, c.Right, c.Bottom, c.Top = 20.0, -20.0, 20.0, -20.0

	//Perspective
	c.Aspect = 16.0 / 9.0
	c.SensorHeight = 24.0 //24mm
	c.SensorWidth = 36.0  //36mm
	c.FocalLength = 14.0  //24mm
}

//Horizontal Field of View
func (c *Camera) FieldOfViewX() float64 {
	return mgl.RadToDeg(2.0 * math.Atan(c.SensorWidth/(c.FocalLength*2.0)))
}

//Vertical Field of View
func (c *Camera) FieldOfViewY() float64 {
	return mgl.RadToDeg(2.0 * math.Atan(c.SensorHeight/(c.FocalLength*2.0)))
}

func (c *Camera) ProjectionMatrix() mgl.Mat4 {
	switch c.ProjectionType {
	case ORTHOGRAPHIC:
		return mgl.Ortho(c.Left, c.Right, c.Bottom, c.Top, c.Near, c.Far)

	case PERSPECTIVE:
		fallthrough
	default:
		return mgl.Perspective(c.FieldOfViewY(), c.Aspect, c.Near, c.Far)
		/*m[0], m[5] = c.FieldOfView, c.FieldOfView
		m[10] = (c.Far + c.Near) / (c.Near - c.Far);
		m[14] = (2.0 * c.Far * c.Near) / (c.Near - c.Far);
		m[11] = -1.0;*/
	}
}

func (c *Camera) TransformationMatrix() mgl.Mat4 {
	return c.Translation.Mul4(c.Rotation)
}

func (c *Camera) ViewMatrix() mgl.Mat4 {
	return c.TransformationMatrix().Inv()
}
