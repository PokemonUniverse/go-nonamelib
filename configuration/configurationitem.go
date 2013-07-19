package configuration

import ()

type ConfigurationItem struct {
	section      string
	name         string
	friendlyName string
	defaultValue interface{}
}

func NewConfigurationItem(_section, _name, _friendlyName string, _defaultValue interface{}) *ConfigurationItem {
	return &ConfigurationItem{section: _section,
		name:         _name,
		friendlyName: _friendlyName,
		defaultValue: _defaultValue}
}

func (c *ConfigurationItem) GetSection() string {
	return c.section
}

func (c *ConfigurationItem) SetSection(_section string) {
	c.section = _section
}

func (c *ConfigurationItem) GetName() string {
	return c.name
}

func (c *ConfigurationItem) SetName(_name string) {
	c.name = _name
}

func (c *ConfigurationItem) GetFriendlyName() string {
	return c.friendlyName
}

func (c *ConfigurationItem) SetFriendlyName(_friendlyName string) {
	c.friendlyName = _friendlyName
}

func (c *ConfigurationItem) GetDefaultValue() interface{} {
	return c.defaultValue
}

func (c *ConfigurationItem) SetDefaultValue(_value interface{}) {
	c.defaultValue = _value
}
