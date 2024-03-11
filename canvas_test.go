package canvas

import (
	"fmt"
	"strings"
	"testing"

	"github.com/B9O2/canvas/containers"
	"github.com/B9O2/canvas/pixel"
)

var HelloWorldArt = `_   _      _ _        __        __         _     _ _ 
| | | | ___| | | ___   \ \      / /__  _ __| | __| | |
| |_| |/ _ \ | |/ _ \   \ \ /\ / / _ \| '__| |/ _` + "`" + ` | |
|  _  |  __/ | | (_) |   \ V  V / (_) | |  | | (_| |_|
|_| |_|\___|_|_|\___/     \_/\_/ \___/|_|  |_|\__,_(_)`

var TriangleArt = ` _____  _  _____
(___  \( )/  ___)
  (___ | | ___)
   /")'| |'("\
  | |  | |  | |
   \ \_| |_/ /
    '._!' _.'
      / .'\
     | / | |
      \|/ /
       /.<
      (| |)
       | '
       | |
       '-'`

func TestCanvas(t *testing.T) {
	text := containers.NewTextArea(strings.Repeat("The wheel turns, nothing is ever new.", 1))
	//text.SetFrame(0, 0, 1, 1)
	//text.SetAliginLeft(true)
	//text.SetBorder(pixel.Space, pixel.Dot)

	art := containers.NewAsciiArt(strings.Split(TriangleArt, "\n"))
	//art.SetBorder(pixel.Space, pixel.Dot)

	box := containers.NewHStack()
	box.SetBorder(pixel.NewPixel('*', nil))
	box.SetPadding(pixel.Dot)

	vs := containers.NewVStack(box, box, box)
	vs.SetBorder(pixel.Space)
	vs.SetPadding(pixel.Dot)

	left := containers.NewHStack(box, text, box)
	left.SetBorder(pixel.Space)
	left.SetPadding(pixel.Dot)

	hs := containers.NewZStack(art, text)
	//hs.SetFrame(0, 0, 15, 15)
	//hs.SetBorder(pixel.Space, pixel.Dot)

	all := containers.NewVStack(hs, containers.Nil)
	//all.SetBorder(pixel.NewPixel('*', nil), pixel.Dot)

	pm, err := all.Draw(81, 15)
	if err != nil {
		t.Fatal(err)
	}
	err = pm.Display(pixel.Space)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Finished")
}
