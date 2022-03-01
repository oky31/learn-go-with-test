package shape

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangel Rectangle) float64 {
	return 2 * (rectangel.Width + rectangel.Height)
}

func Area(rectangel Rectangle) float64 {
	return rectangel.Width * rectangel.Height
}
