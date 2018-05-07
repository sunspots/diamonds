package vector2

import "fmt"

// Vector2 represents a position in a cartesian
type Vector2 struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

// New creates a new Vector2 initialized with x and y
func New(x, y int) Vector2 {
	return Vector2{x, y}
}

// Zero returns the vector with 0:0
func Zero() Vector2 {
	return Vector2{0, 0}
}

// Default direction vectors

// North returns the cardinal directional vector for north, 0:-1
func North() Vector2 {
	return Vector2{0, -1}
}

// East returns the cardinal directional vector for east, 1:0
func East() Vector2 {
	return Vector2{1, 0}
}

// South returns the cardinal directional vector for south, 0:1
func South() Vector2 {
	return Vector2{0, 1}
}

// West returns the cardinal directional vector for west, -1:0
func West() Vector2 {
	return Vector2{-1, 0}
}

// Add another Vector2, returns a new Vector2 as the result
func (v Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{v.X + v2.X, v.Y + v2.Y}
}

// Sub subtracts another Vector2, returns a new Vector2 as the result
func (v Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{v.X - v2.X, v.Y - v2.Y}
}

// Multiply with another Vector3, returns a new Vector2 as the result
func (v Vector2) Multiply(v2 Vector2) Vector2 {
	return Vector2{v.X * v2.X, v.Y * v2.Y}
}

// Divide with another vector, returns a new Vector2 as the result
func (v Vector2) Divide(v2 Vector2) Vector2 {
	return Vector2{v.X / v2.X, v.Y / v2.Y}
}

// MultiplyScalar multiplies the vector with a scalar on both axes,
// returns a new Vector2 as the result
func (v Vector2) MultiplyScalar(s int) Vector2 {
	return Vector2{v.X * s, v.Y * s}
}

// DivideScalar multiplies the vector with a scalar on both axes,
// returns a new Vector2 as the result
func (v Vector2) DivideScalar(s int) Vector2 {
	return Vector2{v.X / s, v.Y / s}
}

// Distance returns the Manhattan distance from v to v2
func (v Vector2) Distance(v2 Vector2) int {
	r := v.Sub(v2)
	return r.X + r.Y
}

func (v Vector2) String() string {
	return fmt.Sprintf("%v:%v", v.X, v.Y)
}
