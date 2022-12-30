package conv

import (
	"strconv"

	"github.com/fgsoftware1/gomath/constants"
)

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 6, 64)
}

func DegToRad(angle float64) float64 {
	return angle*constants.Pi/180
}

func RadToDeg(angle float64) float64 {
	return angle*180/constants.Pi
}