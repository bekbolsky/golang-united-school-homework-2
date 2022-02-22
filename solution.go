package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SidesInt int

const (
	SidesCircle   SidesInt = 0
	SidesTriangle SidesInt = 3
	SidesSquare   SidesInt = 4
)

func CalcSquare(sideLen float64, sidesNum SidesInt) float64 {
	switch sidesNum {
	case SidesCircle:
		S := math.Pi * sideLen * sideLen
		return S
	case SidesTriangle:
		S := (math.Sqrt(3) * math.Pow(sideLen, 2)) / 4
		return S
	case SidesSquare:
		S := sideLen * sideLen
		return S
	default:
		return 0.0
	}
}
