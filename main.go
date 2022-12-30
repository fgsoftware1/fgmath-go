package main

import (
	"fmt"

	"github.com/fgsoftware1/gomath/math"
	conv "github.com/fgsoftware1/gomath/utils"
)

func main() {
	deg := 30.0
	rad := 0.523600
	vec := math.New(1, 1)
	vec1 := math.New(1, 2)

	fmt.Println("\x1b[31m*****************")
	fmt.Println("\x1b[31m*\x1b[35m    GOMATH     \x1b[31m*")
	fmt.Println("\x1b[31m*****************")
	fmt.Println("\x1b[35mCONVERSION TESTS")
	fmt.Println("\x1b[34mAngle(deg):")
	fmt.Println("\x1b[33m" + "30")
	fmt.Println("\x1b[34mDegrees to Radians:")
	fmt.Println("\x1b[32m" + conv.FloatToString(conv.DegToRad(deg)))
	fmt.Println("\x1b[34mAngle(rad):")
	fmt.Println("\x1b[33m" + "0.523600")
	fmt.Println("\x1b[34mRadians to Degrees:")
	fmt.Println("\x1b[32m" + conv.FloatToString(conv.RadToDeg(rad)))
	fmt.Println("\x1b[35mMATH TESTS")
	fmt.Println("\x1b[34mVector1")
	fmt.Println("\x1b[33m"+"X:\x1b[32m", vec.X, "\x1b[33m"+"Y:\x1b[32m", vec.Y)
	fmt.Println("\x1b[34mVector2")
	fmt.Println("\x1b[33m"+"X:\x1b[32m", vec1.X, "\x1b[33m"+"Y:\x1b[32m", vec1.Y)
	fmt.Println("\x1b[31mAddition test:")
	fmt.Println("\x1b[33m"+"X:\x1b[32m", vec.Add(vec1).X, "\x1b[33m"+"Y:\x1b[32m", vec.Add(vec1).Y)
	fmt.Println("\x1b[31mDot test:")
	fmt.Println("\x1b[33m"+"X:\x1b[32m", vec.Dot(vec1))
	fmt.Println("\x1b[31mDistance test:")
	fmt.Println("\x1b[33m"+"X:\x1b[32m", vec.Distance(vec1))
	fmt.Println("\x1b[31mLerp test:")
	fmt.Println("\x1b[34m" + "Vector1:")
	fmt.Println("\x1b[33m"+"X:\x1b[32m", vec.Lerp(vec1, 0.0).X, "\x1b[33m"+"Y:\x1b[32m", vec.Lerp(vec1, 0.0).Y)
	fmt.Println("\x1b[34m" + "Vector2:")
	fmt.Println("\x1b[33m"+"X:\x1b[32m", vec.Lerp(vec1, 1.0).X, "\x1b[33m"+"Y:\x1b[32m", vec.Lerp(vec1, 1.0).Y)
	fmt.Println("\x1b[34m" + "Vector1/Vector2:")
	fmt.Println("\x1b[33m"+"X:\x1b[32m", vec.Lerp(vec1, 0.5).X, "\x1b[33m"+"Y:\x1b[32m", vec.Lerp(vec1, 0.5).Y)
}
