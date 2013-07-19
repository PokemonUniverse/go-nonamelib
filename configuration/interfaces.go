package configuration

import ()

type IConfigurationProvider interface {
	Initialize()
	SetValue(_item IConfigurationItem, _value interface{})
	GetString(_item IConfigurationItem) (string, error)
	GetInt(_item IConfigurationItem) (int, error)
	GetFloat64(_item IConfigurationItem) (float64, error)
	GetBool(_item IConfigurationItem) (bool, error)
}

type IConfigurationItemContainer interface {
	GetConfigurationItems() map[string]IConfigurationItem
}

type IConfigurationItem interface {
	GetSection() string
	GetName() string
	GetFriendlyName() string
	GetDefaultValue() interface{}
}
