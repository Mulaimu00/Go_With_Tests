package structs

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{
			name: "Rectangle Perimeter",
			shape: Rectangle{Width: 10.0, Height: 10.0},
			want: 40.0,
		},
		{
			name: "Circle Circumfence",
			shape: Circle{Radius: 5},
			want: 31.41592653589793,
		},
		{
			name: "Triangle Perimeter",
			shape: Triangle{Height:5, Base: 5, Hypotenuse: 5},
			want: 15,
		},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{
			name : "Rectangle Area",
			shape: Rectangle{Width: 12.0, Height: 6.0},
			want: 72.0,
		},
		{
			name: "Circle Area",
			shape: Circle{Radius: 10},
			want: 314.1592653589793,
		},
		{
			name: "Triangle Area",
			shape: Triangle{Height: 5, Base: 5, Hypotenuse: 5},
			want: 10.825317547305483,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
		})
	}
}
