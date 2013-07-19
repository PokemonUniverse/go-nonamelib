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
package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	L_FILE    int = 1 // Write to file
	L_CONSOLE int = 2 // Print to terminal/console

	F_DEBUG int = 3 // Print debug lines
)

var (
	Log         *log.Logger
	Flags       int
	LogFilename string
)

// A Logger represents an object which inherits from io.Writer
// it's used in combination with log.Logger to write text to a file
type Logger struct {
	file *os.File
}

// Logger init function will be called automatically
// http://golang.org/doc/effective_go.html#init
func init() {
	Log = log.New(&Logger{}, "", log.Ltime)
}

// Use the io.Writer interface
func (l *Logger) Write(p []byte) (n int, err error) {
	if Flags&L_CONSOLE != 0 { // Print string to terminal before writing to file
		fmt.Print(string(p))
	}

	if Flags&L_FILE != 0 && len(LogFilename) > 0 { // Write to file
		filename := fmt.Sprintf("%s-%s.log", LogFilename, time.Now().Format("20060102"))

		logFile, fileErr := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if fileErr == nil {
			defer logFile.Close()
			logFile.Write(p)
		}
	}

	return len(p), nil
}

func Println(v ...interface{}) {
	Log.Output(2, fmt.Sprintln(v...))
}

func Printf(format string, v ...interface{}) {
	Log.Output(2, fmt.Sprintf(format, v...))
}

func Debug(_struct, _method, _message string, v ...interface{}) {
	if Flags&F_DEBUG != 0 {
		if len(v) > 0 {
			_message = fmt.Sprintf(_message, v...)
		}
		Printf("[D] %v - %v - %v\n", _struct, _method, _message)
	}
}

func Verbose(_struct, _method, _message string, v ...interface{}) {
	if len(v) > 0 {
		_message = fmt.Sprintf(_message, v...)
	}
	Printf("[V] %v - %v - %v\n", _struct, _method, _message)
}

func Info(_struct, _method, _message string, v ...interface{}) {
	if len(v) > 0 {
		_message = fmt.Sprintf(_message, v...)
	}
	Printf("[I] %v - %v - %v\n", _struct, _method, _message)
}

func Warning(_struct, _method, _message string, v ...interface{}) {
	if len(v) > 0 {
		_message = fmt.Sprintf(_message, v...)
	}
	Printf("[W] %v - %v - %v\n", _struct, _method, _message)
}

func Error(_struct, _method, _message string, v ...interface{}) {
	if len(v) > 0 {
		_message = fmt.Sprintf(_message, v...)
	}
	Printf("[E] %v - %v - %v\n", _struct, _method, _message)
}
