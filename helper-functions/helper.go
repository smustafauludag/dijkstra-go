package helperfunctions 
import "math"

// Rounds the float value with given precision
func RoundFloat(val float64, precision uint) float64 {
  ratio := math.Pow(10, float64(precision))
  return math.Round(val*ratio) / ratio
}

// Returns the distance between two x,y coordinates
func DistanceXY(a [2]float64, b [2]float64) float64 {
  dist := math.Sqrt(math.Pow((a[0]-b[0]),2) + math.Pow((a[1]-b[1]),2))
  return dist
}

