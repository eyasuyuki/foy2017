package main

import (
	"bitbucket.org/mikehouston/three"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	scene := three.SceneFromJS(js.Global.Get("scene"))

	size := 32
	step := 50
	geometry := three.NewGeometry()
	for x := 0; x < size; x++ {
		for z := 0; z < size; z++ {
			geometry.Vertices().Push(three.NewVector3(
				float64(x*step-size*step/2),
				float64(z*z-x*x),
				float64(z*step-size*step/2),
			))

			if x > 0 && z > 0 {
				p1 := z + x*size
				p2 := (z - 1) + x*size
				p3 := (z - 1) + (x-1)*size
				p4 := z + (x-1)*size

				geometry.Faces().Push(three.NewFace3(p1, p2, p3))
				geometry.Faces().Push(three.NewFace3(p3, p4, p1))

				geometry.Faces().Push(three.NewFace3(p3, p2, p1))
				geometry.Faces().Push(three.NewFace3(p1, p4, p3))
			}
		}
	}
	geometry.ComputeFaceNormals()

	material := three.NewMeshNormalMaterial()
	mesh := three.NewMesh(geometry, material)
	scene.Add(mesh)
}
