package models

import (
	"strings"
)

const ELLIPSE string = "ELLIPSE"
const RECTANGLE string = "RECTANGLE"
const TRIANGLE string = "TRIANGLE"

type Item struct {
	Id        string  `json:"id"`
	ShapeType string  `json:"tipo"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Creator   string  `json:"creador"`
}

func (i Item) IsValidShapeType() bool {
	shapeType := strings.ToUpper(i.ShapeType)
	return shapeType == ELLIPSE || shapeType == RECTANGLE || shapeType == TRIANGLE
}

func (i Item) IsValidData() bool {
	return i.Id != "" && i.A != 0 && i.B != 0
}
