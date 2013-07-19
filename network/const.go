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

const (
	HEADER_PING            = 0x00
	HEADER_LOGIN           = 0x01
	HEADER_LOGOUT          = 0x02
	HEADER_REFRESHCOMPLETE = 0x03

	HEADER_CHAT     = 0x10
	HEADER_TEXT_MSG = 0x11

	HEADER_DIALOG = 0x12

	HEADER_FRIENDLIST   = 0xA0
	HEADER_FRIENDUPDATE = 0xA1
	HEADER_QUESTLIST    = 0xA2
	HEADER_QUESTUPDATE  = 0xA3

	HEADER_IDENTITY = 0xAA

	HEADER_WALK       = 0xB1
	HEADER_CANCELWALK = 0xB2
	HEADER_WARP       = 0xB3
	HEADER_TURN       = 0xB4

	HEADER_TILES          = 0xC1
	HEADER_ADDCREATURE    = 0xC2
	HEADER_REMOVECREATURE = 0xC3
	HEADER_REFRESHWORLD   = 0xC4

	// Pokemon headers
	HEADER_POKEMONPARTY = 0xD0

	// Battle
	HEADER_BATTLEMESSAGE       = 0xE0
	HEADER_BATTLE_UPDATEMYPOKE = 0xE1
)

const (
	BATTLEEVENT_SLEEP              uint8 = 0
	BATTLEEVENT_TEXTID                   = 1
	BATTLEEVENT_TEXT                     = 2
	BATTLEEVENT_CHANGEHP                 = 3
	BATTLEEVENT_ANIMATION                = 4
	BATTLEEVENT_ALLOWCONTROL             = 5
	BATTLEEVENT_CHANGEPOKEMON_SELF       = 6
	BATTLEEVENT_CHANGEPOKEMON            = 7
	BATTLEEVENT_CHANGESELECTION          = 8
	BATTLEEVENT_CHANGEPP                 = 9
	BATTLEEVENT_CHANGESTATUS             = 10
	BATTLEEVENT_CHANGELEVELSELF          = 11
	BATTLEEVENT_CHANGELEVEL              = 12
	BATTLEEVENT_CHANGEATTACK             = 13
	BATTLEEVENT_CHANGESCREEN             = 14
	BATTLEEVENT_DIALOGUE                 = 15
	BATTLEEVENT_REMOVEPLAYER             = 16
	BATTLEEVENT_CHANGEEXP                = 17
	BATTLEEVENT_CHANGEMAXHP              = 18

	BATTLEEVENT_STOPBATTLE = 999
)

const (
	SPEAK_NORMAL    = 1
	SPEAK_YELL      = 2
	SPEAK_WHISPER   = 3
	SPEAK_PRIVATE   = 6
	SPEAK_CHANNEL   = 7
	SPEAK_BROADCAST = 255
)

const (
	DIALOG_CLOSE   = 1
	DIALOG_NPC     = 2
	DIALOG_NPCTEXT = 3
	DIALOG_OPTIONS = 4
)

const (
	CHANNEL_LOCAL   = 0
	CHANNEL_WORLD   = 1
	CHANNEL_TRADE   = 2
	CHANNEL_BATTLE  = 3
	CHANNEL_GUILD   = 4
	CHANNEL_PARTY   = 5
	CHANNEL_PRIVATE = 65535
)

const (
	MSG_STATUS_WARNING = 0 // Red message in game window and in the console
	MSG_STATUS_DEFAULT = 1 // White message at the bottom of the game window and in the console
	MSG_INFO_DESCR     = 2 // Green message in game window and in the console
	MSG_STATUS_SMALL   = 3 // White message at the bottom of the game window
)
