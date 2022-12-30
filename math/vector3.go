package math

type Vector3 struct {
	X, Y, Z float64
}

func New(x float64, y float64, z float64) *Vector3 {
	return &Vector3{X: x, Y: y, Z: z}
}

func Dot(a *Vector3, b *Vector3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}
