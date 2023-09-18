package draw

import (
	"github.com/go-gl/gl/v2.1/gl"
)

var lineDrawn bool = false
var rectDrawn bool = false

func DrawLine(x1, y1, x2, y2 float32, color [3]float32) {
	gl.Begin(gl.LINES)
	gl.Color3f(color[0], color[1], color[2])
	gl.Vertex2f(x1, y1)
	gl.Vertex2f(x2, y2)
	gl.End()
}

func DrawRect(x, y float32, height, width float32, color [3]float32) {
	gl.Begin(gl.QUADS)
	gl.Color3f(color[0], color[1], color[2])
	gl.Vertex2f(x, y)
	gl.Vertex2f(x+width, y)
	gl.Vertex2f(x+width, y+height)
	gl.Vertex2f(x, y+height)
	gl.End()
}

func SetBackgroundColor(color [3]float32) {
	gl.ClearColor(color[0], color[1], color[2], 1)

}
