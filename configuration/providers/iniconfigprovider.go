package providers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PokemonUniverse/go-nonamelib/configuration"
)

const (
	// Get Errors
	SectionNotFound = iota
	OptionNotFound
	MaxDepthReached

	// Read Errors
	BlankSection

	// Get and Read Errors
	CouldNotParse
)

var (
	DefaultSection = "default" // Default section name (must be lower-case).
	DepthValues    = 200       // Maximum allowed depth when recursively substituing variable names.

	// Strings accepted as bool.
	BoolStrings = map[string]bool{
		"t":     true,
		"true":  true,
		"y":     true,
		"yes":   true,
		"on":    true,
		"1":     true,
		"f":     false,
		"false": false,
		"n":     false,
		"no":    false,
		"off":   false,
		"0":     false,
	}

	varRegExp = regexp.MustCompile(`%\(([a-zA-Z0-9_.\-]+)\)s`)
)

type IniConfigProvider struct {
	iniPath string

	data map[string]map[string]string // Maps sections to options to values.
}

func NewIniConfigProvider(_file string) *IniConfigProvider {
	return &IniConfigProvider{iniPath: _file,
		data: make(map[string]map[string]string)}
}

func (p *IniConfigProvider) Initialize() {
	f, err := os.OpenFile(p.iniPath, os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	if err := p.read(f); err != nil {
		panic(err)
	}

	newItemsAdded := false
	for section, items := range configuration.GetConfigurationItemCollection() {
		p.addSection(section)

		for _, item := range items {
			if _, found := p.data[section][item.GetName()]; !found {
				p.addOption(section, item.GetName(), fmt.Sprintf("%v", item.GetDefaultValue()))
				newItemsAdded = true
			}
		}
	}

	if newItemsAdded {
		p.write()
	}
}

func (p *IniConfigProvider) SetValue(_item configuration.IConfigurationItem, _value interface{}) {
	section := _item.GetSection()
	option := _item.GetName()

	if section == "" {
		section = "default"
	}

	section = strings.ToLower(section)
	option = strings.ToLower(option)

	if _, ok := p.data[section]; ok {
		p.data[section][option] = fmt.Sprintf("%v", _value)

		p.write()
	} else {
		println("Section not found")
	}
}

func (p *IniConfigProvider) GetString(_item configuration.IConfigurationItem) (string, error) {
	return p.getRawString(_item.GetSection(), _item.GetName())
}

func (p *IniConfigProvider) GetInt(_item configuration.IConfigurationItem) (int, error) {
	var value int

	sv, err := p.getRawString(_item.GetSection(), _item.GetName())
	if err == nil {
		value, err = strconv.Atoi(sv)
		if err != nil {
			err = GetError{CouldNotParse, "int", sv, _item.GetSection(), _item.GetName()}
		}
	}

	return value, err
}

func (p *IniConfigProvider) GetFloat64(_item configuration.IConfigurationItem) (float64, error) {
	var value float64

	sv, err := p.getRawString(_item.GetSection(), _item.GetName())
	if err == nil {
		value, err = strconv.ParseFloat(sv, 32)
		if err != nil {
			err = GetError{CouldNotParse, "float", sv, _item.GetSection(), _item.GetName()}
		}
	}

	return value, err
}

func (p *IniConfigProvider) GetBool(_item configuration.IConfigurationItem) (bool, error) {
	sv, err := p.getRawString(_item.GetSection(), _item.GetName())
	if err != nil {
		return false, err
	}

	value, ok := BoolStrings[strings.ToLower(sv)]
	if !ok {
		return false, GetError{CouldNotParse, "bool", sv, _item.GetSection(), _item.GetName()}
	}

	return value, nil
}

// AddSection adds a new section to the configuration.
// It returns true if the new section was inserted, and false if the section already existed.
func (p *IniConfigProvider) addSection(section string) bool {
	section = strings.ToLower(section)

	if _, ok := p.data[section]; ok {
		return false
	}
	p.data[section] = make(map[string]string)

	return true
}

// AddOption adds a new option and value to the configuration.
// It returns true if the option and value were inserted, and false if the value was overwritten.
// If the section does not exist in advance, it is created.
func (p *IniConfigProvider) addOption(section string, option string, value string) bool {
	p.addSection(section) // make sure section exists

	section = strings.ToLower(section)
	option = strings.ToLower(option)

	_, ok := p.data[section][option]
	p.data[section][option] = value

	return !ok
}

// GetRawString gets the (raw) string value for the given option in the section.
// The raw string value is not subjected to unfolding, which was illustrated in the beginning of this documentation.
// It returns an error if either the section or the option do not exist.
func (p *IniConfigProvider) getRawString(section string, option string) (value string, err error) {
	if section == "" {
		section = "default"
	}

	section = strings.ToLower(section)
	option = strings.ToLower(option)

	if _, ok := p.data[section]; ok {
		if value, ok = p.data[section][option]; ok {
			return value, nil
		}
		return "", GetError{OptionNotFound, "", "", section, option}
	}
	return "", GetError{SectionNotFound, "", "", section, option}
}

func (p *IniConfigProvider) read(_reader io.Reader) (err error) {
	buf := bufio.NewReader(_reader)

	var section, option string
	section = "default"
	for {
		l, buferr := buf.ReadString('\n') // parse line-by-line
		l = strings.TrimSpace(l)

		if buferr != nil {
			if buferr != io.EOF {
				return err
			}

			if len(l) == 0 {
				break
			}
		}

		// switch written for readability (not performance)
		switch {
		case len(l) == 0: // empty line
			continue

		case l[0] == '#': // comment
			continue

		case l[0] == ';': // comment
			continue

		case len(l) >= 3 && strings.ToLower(l[0:3]) == "rem": // comment (for windows users)
			continue

		case l[0] == '[' && l[len(l)-1] == ']': // new section
			option = "" // reset multi-line value
			section = strings.TrimSpace(l[1 : len(l)-1])
			p.addSection(section)

		case section == "": // not new section and no section defined so far
			return ReadError{BlankSection, l}

		default: // other alternatives
			i := strings.IndexAny(l, "=:")
			switch {
			case i > 0: // option and value
				i := strings.IndexAny(l, "=:")
				option = strings.TrimSpace(l[0:i])
				value := strings.TrimSpace(stripComments(l[i+1:]))
				p.addOption(section, option, value)

			case section != "" && option != "": // continuation of multi-line value
				prev, _ := p.getRawString(section, option)
				value := strings.TrimSpace(stripComments(l))
				p.addOption(section, option, prev+"\n"+value)

			default:
				return ReadError{CouldNotParse, l}
			}
		}

		// Reached end of file
		if buferr == io.EOF {
			break
		}
	}
	return nil
}

// Writes the configuration file to the io.Writer.
func (p *IniConfigProvider) write() (err error) {
	f, err := os.OpenFile(p.iniPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	buf := bytes.NewBuffer(nil)

	for section, sectionmap := range p.data {
		if section == DefaultSection && len(sectionmap) == 0 {
			continue // skip default section if empty
		}
		if _, err = buf.WriteString("[" + section + "]\n"); err != nil {
			return err
		}
		for option, value := range sectionmap {
			if _, err = buf.WriteString(option + "=" + value + "\n"); err != nil {
				return err
			}
		}
		if _, err = buf.WriteString("\n"); err != nil {
			return err
		}
	}

	buf.WriteTo(f)

	return nil
}

type GetError struct {
	Reason    int
	ValueType string
	Value     string
	Section   string
	Option    string
}

func (err GetError) Error() string {
	switch err.Reason {
	case SectionNotFound:
		return fmt.Sprintf("section '%s' not found", err.Section)
	case OptionNotFound:
		return fmt.Sprintf("option '%s' not found in section '%s'", err.Option, err.Section)
	case CouldNotParse:
		return fmt.Sprintf("could not parse %s value '%s'", err.ValueType, err.Value)
	case MaxDepthReached:
		return fmt.Sprintf("possible cycle while unfolding variables: max depth of %d reached", DepthValues)
	}

	return "invalid get error"
}

type ReadError struct {
	Reason int
	Line   string
}

func (err ReadError) Error() string {
	switch err.Reason {
	case BlankSection:
		return "empty section name not allowed"
	case CouldNotParse:
		return fmt.Sprintf("could not parse line: %s", err.Line)
	}

	return "invalid read error"
}

func stripComments(l string) string {
	// comments are preceded by space or TAB
	for _, c := range []string{" ;", "\t;", " #", "\t#"} {
		if i := strings.Index(l, c); i != -1 {
			l = l[0:i]
		}
	}
	return l
}
