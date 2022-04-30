package main

import "fmt"

// Drawer draws on the underlying graphics device.
type Drawer interface {
	// DrawEllipseInRect draws an ellipse in rectangle
	DrawEllipseInRect(rect Rect) error
}

// OpenGL drawer.
type OpenGL struct{}

// DrawEllipseInRect draws an ellipse in rectangle.
func (gl *OpenGL) DrawEllipseInRect(r Rect) error {
	fmt.Printf("OpenGL is drawing ellipse in rect %v\n", r)
	return nil
}

// Direct2D drawer.
type Direct2D struct{}

// DrawEllipseInRect draws an ellipse in rectangle.
func (d2d *Direct2D) DrawEllipseInRect(r Rect) error {
	fmt.Printf("Direct2D is drawing ellipse in rect %v\n", r)
	return nil
}

type Rect struct {
	Location Point
	Size     Size
}

type Point struct {
	X, Y float64
}

type Size struct {
	Width  float64
	Height float64
}

// Circle represents a circle shape.
type Circle struct {
	// DrawingContext for this circle
	DrawingContext Drawer
	// Center of the circle
	Center Point
	// Radius of the circle
	Radius float64
}

// Draw draws a circle.
func (circle *Circle) Draw() error {
	rect := Rect{
		Location: Point{
			X: circle.Center.X - circle.Radius,
			Y: circle.Center.Y - circle.Radius,
		},
		Size: Size{
			Width:  2 * circle.Radius,
			Height: 2 * circle.Radius,
		},
	}

	return circle.DrawingContext.DrawEllipseInRect(rect)
}

func main() {
	openGL := &OpenGL{}
	direct2D := &Direct2D{}

	circle := &Circle{
		Center: Point{X: 100, Y: 100},
		Radius: 50,
	}

	circle.DrawingContext = openGL
	_ = circle.Draw()

	circle.DrawingContext = direct2D
	_ = circle.Draw()
}
