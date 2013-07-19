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
package network

type INetMessageWriter interface {
	WritePacket() IPacket
}

type INetMessageReader interface {
	GetHeader() uint8
	ReadPacket(IPacket) error
}

type IPacket interface {
	Reset()
	CanAdd(_size uint16) bool

	GetHeader() uint16
	SetHeader()

	GetBuffer() [PACKET_MAXSIZE]uint8
	GetBufferSlice() []uint8
	GetMsgSize() uint16

	ReadUint8() (uint8, error)
	ReadUint16() (uint16, error)
	ReadInt16() (int16, error)
	ReadUint32() (uint32, error)
	ReadInt32() (int, error)
	ReadUint64() (uint64, error)
	ReadInt64() (int64, error)
	ReadString() (string, error)
	ReadBool() (bool, error)

	AddUint8(_value uint8) bool
	AddUint16(_value uint16) bool
	AddUint32(_value uint32) bool
	AddUint64(_value uint64) bool
	AddBool(_value bool) bool
	AddString(_value string) bool
	AddBuffer(_value []uint8) bool
}
