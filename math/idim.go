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
package putools

// Idim returns the maximum of x-y or 0.
func Idim(x, y int) int {
	if x > y {
		return x - y
	}
	return 0
}

// Imax returns the larger of x or y.
func Imax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Imin returns the smaller of x or y.
func Imin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
