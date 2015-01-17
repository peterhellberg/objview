package main

import (
	"fmt"
	"os"

	"gopkg.in/qml.v1"
	"gopkg.in/qml.v1/gl/2.0"

	"github.com/peterhellberg/wavefront"
)

var (
	modelFn = "model.obj"
	width   = 800
	height  = 600
)

func main() {
	if len(os.Args) == 2 {
		modelFn = os.Args[1]
	}

	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()

	engine.On("quit", func() {
		os.Exit(0)
	})

	model, err := wavefront.Read(modelFn)
	if err != nil {
		return err
	}

	qml.RegisterTypes("ObjView", 0, 1, []qml.TypeSpec{{
		Init: func(m *Model, obj qml.Object) {
			m.Object = obj
			m.model = model
		},
	}})

	component, err := engine.LoadString("objview.qml", qmlString)
	if err != nil {
		return err
	}

	win := component.CreateWindow(nil)

	win.On("widthChanged", func(w int) {
		width = w
	})

	win.On("heightChanged", func(h int) {
		height = h
	})

	win.Show()
	win.Wait()

	return nil
}

type Model struct {
	qml.Object

	model map[string]*wavefront.Object

	Rotation int
}

func (m *Model) SetRotation(rotation int) {
	m.Rotation = rotation

	m.Call("update")
}

func (m *Model) Paint(p *qml.Painter) {
	gl := GL.API(p)

	gl.Enable(GL.BLEND)
	gl.BlendFunc(GL.SRC_ALPHA, GL.ONE_MINUS_SRC_ALPHA)

	gl.ShadeModel(GL.SMOOTH)
	gl.Enable(GL.DEPTH_TEST)
	gl.DepthMask(true)
	gl.Enable(GL.NORMALIZE)

	gl.Clear(GL.DEPTH_BUFFER_BIT)

	w := float32(width)
	gl.Scalef(w/6, w/6, w/6)

	lka := []float32{0.3, 0.3, 0.3, 1.0}
	lkd := []float32{1.0, 1.0, 1.0, 0.0}
	lks := []float32{1.0, 1.0, 1.0, 1.0}
	lpos := []float32{-2, 6, 3, 1.0}

	gl.Enable(GL.LIGHTING)
	gl.Lightfv(GL.LIGHT0, GL.AMBIENT, lka)
	gl.Lightfv(GL.LIGHT0, GL.DIFFUSE, lkd)
	gl.Lightfv(GL.LIGHT0, GL.SPECULAR, lks)
	gl.Lightfv(GL.LIGHT0, GL.POSITION, lpos)
	gl.Enable(GL.LIGHT0)

	gl.EnableClientState(GL.NORMAL_ARRAY)
	gl.EnableClientState(GL.VERTEX_ARRAY)

	gl.Translatef(2, 2.5, 0)
	gl.Rotatef(260, 190, -135, 85)
	gl.Rotatef(float32(90+((36000+m.Rotation)%360)), 1, 0, 0)

	gl.Disable(GL.COLOR_MATERIAL)

	for _, obj := range m.model {
		for _, group := range obj.Groups {
			gl.Materialfv(GL.FRONT, GL.AMBIENT, group.Material.Ambient)
			gl.Materialfv(GL.FRONT, GL.DIFFUSE, group.Material.Diffuse)
			gl.Materialfv(GL.FRONT, GL.SPECULAR, group.Material.Specular)
			gl.Materialf(GL.FRONT, GL.SHININESS, group.Material.Shininess)
			gl.VertexPointer(3, GL.FLOAT, 0, group.Vertexes)
			gl.NormalPointer(GL.FLOAT, 0, group.Normals)
			gl.DrawArrays(GL.TRIANGLES, 0, len(group.Vertexes)/3)
		}
	}

	gl.Enable(GL.COLOR_MATERIAL)

	gl.DisableClientState(GL.NORMAL_ARRAY)
	gl.DisableClientState(GL.VERTEX_ARRAY)
}

const qmlString = `
// Start of the QML string
import QtQuick 2.0
import ObjView 0.1

Rectangle {
	id: root

	width: 800
	height: 600

	gradient: Gradient {
		GradientStop { position: 0.0; color: "#003E5F" }
		GradientStop { position: 0.1; color: "#00334E" }
		GradientStop { position: 0.2; color: "#002033" }
		GradientStop { position: 0.3; color: "#001521" }
		GradientStop { position: 0.7; color: "#00080E" }
	}

	Model {
		id: model

		x: 0
		y: 0

		width: parent.width
		height: parent.height

		NumberAnimation on rotation {
			id: anim
			from: 360
			to: 0
			duration: 8000
			loops: Animation.Infinite
		}

		MouseArea {
			anchors.fill: parent

			property real startX
			property real startR

			onPressed: {
				startX = mouse.x
				startR = model.rotation
				anim.running = false
			}
			onReleased: {
				anim.from = model.rotation + 360
				anim.to = model.rotation
				anim.running = true
			}
			onPositionChanged: {
				model.rotation = (36000 + (startR - (mouse.x - startX))) % 360
			}
		}
	}
}

// End of the QML string`

// vim: ft=qml.go
