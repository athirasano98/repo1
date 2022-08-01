package geometry

func Area(len, wid float64) float64 {
    area := len * wid
    return area
}

func Perimeter(len, wid float64) float64 {
    return (len + wid) * 2
}