/*
Vertex implements...

*/

package GameEngine

import (
	"log"
	mgl "github.com/go-gl/mathgl/mgl64"
)

type Vec4 mgl.Vec4
type Vec3 mgl.Vec3
type Vec2 mgl.Vec2

type Vertex struct {
	Position mgl.Vec3
	Color mgl.Vec4
	Normal mgl.Vec3
	UVST mgl.Vec2
	Index uint32
}

func(v *Vertex) CalculateNormal(m *Mesh) {
	v.Normal = mgl.Vec3{}
	v2,v3 := 0,0
	log.Println(len(m.Indices))
	for k,i := range m.Indices {
		if v.Index == i {
			switch(k%3) {
				default:
					switch(len(m.Indices)-k) {
						case 2:
							v2,v3 = k+1,k-1
						case 1:
							v2,v3 = k-1,k-2
						default:
							v2,v3 = k+1,k+2
					}
				case 1:
					v2,v3 = k-1, k+1
				case 2:
					v2,v3 = k-2, k-1
			}
			log.Println(k,v2,v3)
			
			v.Normal = v.Normal.Add(
				m.Vertices[m.Indices[v2]].Position.Sub(
					v.Position,
				).Cross(
					m.Vertices[m.Indices[v3]].Position.Sub(
						v.Position,
					),
				),
			)
		}
	}
	/*if v.Normal.Len() != 0 {
		v.Normal = v.Normal.Normalize()
	}*/
}


/*func(v *Vertex) toVec(property string) (f interface{}) {
	field := reflect.ValueOf(v).FieldByName(property)
	l := field.Len()
	vec := []float64{}
	
	if l > 1 {
		for i := 0; i < l; i++ {
			vec = append(vec, field.Index(i).Float())
		}
	} else {
		vec := []float64{field.Float()}
	}
	
	switch(l) {
		case 2:
			f := mgl.Vec2{vec}
		case 3:
			f := mgl.Vec3{vec}
		case 4:
			f := mgl.Vec4{vec}
		default:
			f := mgl.NewVecNFromData(vec)
	}
	
	return
}*/