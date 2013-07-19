/*Pokemon Universe MMORPG
Copyright (C) 2010 the Pokemon Universe Authors

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.*/
package position

import (
	"strconv"
)

type Position struct {
	X, Y, Z int
}

// ZP is the zero position
var ZP Position

// NewPosition returns a new empty Position
func NewPosition() Position {
	return Position{}
}

// NewPositionFrom returns a new position with the given coordinates
func NewPositionFrom(_x int, _y int, _z int) Position {
	return Position{X: _x, Y: _y, Z: _z}
}

// NewPositionFromHash generates a new Position struct with 
// coordinates extracted from the hash
func NewPositionFromHash(_hash int64) Position {
	z := int(int16(_hash))

	y64 := (_hash >> 17) & 0xFFFF
	yabs := (_hash >> 33) & 0x01

	x64 := (_hash >> 34) & 0xFFFF
	xabs := (_hash >> 50) & 0x01

	var y int = int(y64)
	if yabs == 1 {
		y = 0 - y
	}

	var x int = int(x64)
	if xabs == 1 {
		x = 0 - x
	}

	return Position{X: x, Y: y, Z: z}
}

// String returns a string represntation of p like "3,9,1"
func (p Position) String() string {
	return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + "," + strconv.Itoa(p.Z)
}

// Add returns the position p+q.
func (p Position) Add(q Position) Position {
	return Position{p.X + q.X, p.Y + q.Y, p.Z + q.Z}
}

// Sub returns the position p-q.
func (p Position) Sub(q Position) Position {
	return Position{p.X - q.X, p.Y - q.Y, p.Z - q.Z}
}

// Eq returns true if Position p and q are the same
func (p Position) Eq(_q Position) bool {
	return (p.X == _q.X) && (p.Y == _q.Y) && (p.Z == _q.Z)
}

// Equals calls and returns outcome of Eq
func (p Position) Equals(_q Position) bool {
	return p.Eq(_q)
}

// Create a hash from x, y and z
// hash = [1 bit for positive/negative x][16 bits for x][1 bit for positive/negative y][16 bits for y][16 bits for z]
func (p Position) Hash() int64 {
	return Hash(p.X, p.Y, p.Z)
}

func Hash(_x int, _y int, _z int) int64 {
	var x64 int64
	if _x < 0 {
		x64 = (int64(1) << 50) | ((^(int64(_x) - 1)) << 34)
	} else {
		x64 = (int64(_x) << 34)
	}

	var y64 int64
	if _y < 0 {
		y64 = (int64(1) << 33) | ((^(int64(_y) - 1)) << 17)
	} else {
		y64 = (int64(_y) << 17)
	}

	z64 := int64(_z)
	var index int64 = int64(x64 | y64 | z64)

	return index
}
