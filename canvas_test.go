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
	//text := containers.NewTextArea(strings.Repeat("The wheel turns, nothing is ever new.", 1))
	//text.SetAliginLeft(true)
	//text.SetBorder(pixel.Space, pixel.Dot)

	art := containers.NewHStack(containers.NewAsciiArt(strings.Split(TriangleArt, "\n")))
	art.SetFrame(0, 0, 15, 0)
	//art.SetBorder(pixel.Space, pixel.Dot)

	box := containers.NewHStack()
	box.SetPadding(pixel.Dot)
	//box.SetPadding(pixel.Dot)

	// vs := containers.NewVStack(box, box, box)
	// vs.SetBorder(pixel.Space)
	// vs.SetPadding(pixel.Dot)

	// left := containers.NewHStack(box, text, box)
	// //left.SetBorder(pixel.Space)
	// left.SetPadding(pixel.Dot)

	hs := containers.NewVStack(nil, nil)
	//hs.SetFrame(0, 0, 15, 15)
	hs.SetBorder(pixel.Dot)

	var vs *containers.VStack
	vs = containers.NewVStack() // Initialize vs

	all := containers.NewHStack(vs, vs, hs)
	//all.SetBorder()
	//all.SetPadding(pixel.Dot)
	// all.SetVPadding(1)
	// all.SetHPadding(1)
	//all.SetBorder(pixel.NewPixel('*', nil), pixel.Dot)

	pm, err := all.Draw(80, 25)
	if err != nil {
		t.Fatal(err)
	}
	err = pm.Display(pixel.Space)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Finished")
}

func TestTextAreaUnicode(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		width   uint
		height  uint
		wantErr bool
	}{
		{"Chinese short text fits", "你好", 10, 1, false},
		{"Chinese long text wraps", "你好世界欢迎你", 6, 3, false}, // "你好世界" (width 8) "欢迎你" (width 6)
		{"Mixed emoji and text", "🚀abc你好", 10, 2, false},     // Rocket(2)abc(3)你好(4) = 9
		{"Text wider than area", "你好世界", 2, 2, false},       // "你","好","世","界"
		{"Text taller than area", "你好\n世界\n欢迎", 4, 2, false}, // 3 lines, height 2, should truncate
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := containers.NewTextArea(tt.text)
			// Set a border to ensure DrawBorder logic is also exercised
			ta.SetBorder(pixel.NewPixel('-', nil))
			_, err := ta.Draw(tt.width, tt.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("TextArea.Draw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAsciiArtUnicode(t *testing.T) {
	tests := []struct {
		name    string
		lines   []string
		width   uint
		height  uint
		wantErr bool
	}{
		{"Chinese art fits", []string{"你好", "世界"}, 10, 2, false},
		{"Chinese art truncate width", []string{"你好世界", "欢迎你"}, 4, 2, false}, // "你好" (w4), "欢迎" (w4)
		{"Emoji art truncate width", []string{"🚀abc", "你好🚀"}, 3, 2, false},   // "🚀a", "你好" (rocket w2, a w1; 你好 w4 truncated to 你 w2)
		{"Art taller than area", []string{"一", "二", "三"}, 5, 2, false},     // Should truncate to first 2 lines
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			art := containers.NewAsciiArt(tt.lines)
			art.SetBorder(pixel.NewPixel('|', nil))
			_, err := art.Draw(tt.width, tt.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("AsciiArt.Draw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
