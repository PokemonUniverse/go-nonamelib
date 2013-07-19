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

import math "github.com/PokemonUniverse/nonamelib/math"

// IsInRange2p checks if Position _p is in range (_delta Position) with Position _q
// Only X and Y are checked
func (_p Position) IsInRange2p(_q Position, _delta Position) bool {
	if math.Iabs(_p.X-_q.X) > _delta.X || math.Iabs(_p.Y-_q.Y) > _delta.Y {
		return false
	}

	return true
}

// IsInRange3p checks if Position _p is in range (_delta Position) with Position _q
// All values (X, Y, Z) are checked
func (_p Position) IsInRange3p(_q Position, _delta Position) bool {
	if math.Iabs(_p.X-_q.X) > _delta.X || math.Iabs(_p.Y-_q.Y) > _delta.Y || math.Iabs(_p.Z-_q.Z) > _delta.Z {
		return false
	}
	return true
}
