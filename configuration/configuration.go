package configuration

import (
	"bytes"
	"fmt"
)

type ConfigurationItemCollection map[string]map[string]IConfigurationItem

var configProvider IConfigurationProvider = nil
var configItemCollection ConfigurationItemCollection
var configItems map[string]IConfigurationItem

func init() {
	configItemCollection = make(ConfigurationItemCollection)
	configItems = make(map[string]IConfigurationItem)
}

func setupCheck() {
	if configProvider == nil {
		panic("Configuration - No configuration provider defined. Use SetConfigurationProvider to set one.")
	}
}

func SetConfigurationProvider(_provider IConfigurationProvider) {
	if configProvider != nil {
		panic("Configuration - ConfigurationProvider has already been set")
	}

	configProvider = _provider
}

func GetConfigurationItemCollection() ConfigurationItemCollection {
	return configItemCollection
}

func AddConfigurationItems(_container IConfigurationItemContainer) error {
	var errorBuffer bytes.Buffer

	for key, item := range _container.GetConfigurationItems() {
		if ret, str := addConfigurationItem(key, item); !ret {
			errorBuffer.WriteString(str)
		}
	}

	if errorBuffer.Len() > 0 {
		return fmt.Errorf(errorBuffer.String())
	}

	return nil
}

func Initialize() {
	setupCheck()

	configProvider.Initialize()
}

func SetValue(_key string, _value interface{}) error {
	setupCheck()

	configItem, err := getConfigurationItemByKey(_key)
	if err != nil {
		return err
	}

	configProvider.SetValue(configItem, _value)

	return nil
}

func GetString(_key string) (string, error) {
	setupCheck()

	configItem, err := getConfigurationItemByKey(_key)
	if err != nil {
		return "", err
	}

	v, err := configProvider.GetString(configItem)
	return v, err
}

func GetInt(_key string) (int, error) {
	setupCheck()

	configItem, err := getConfigurationItemByKey(_key)
	if err != nil {
		return 0, err
	}

	v, err := configProvider.GetInt(configItem)
	return v, err
}

func GetFloat64(_key string) (float64, error) {
	setupCheck()

	configItem, err := getConfigurationItemByKey(_key)
	if err != nil {
		return 0, err
	}

	v, err := configProvider.GetFloat64(configItem)
	return v, err
}

func GetBool(_key string) (bool, error) {
	setupCheck()

	configItem, err := getConfigurationItemByKey(_key)
	if err != nil {
		return false, err
	}

	v, err := configProvider.GetBool(configItem)
	return v, err
}

func addConfigurationItem(_key string, _item IConfigurationItem) (bool, string) {
	// Add to global map
	if _, configFound := configItems[_key]; configFound {
		return false, fmt.Sprintf("ConfigurationItem with key '%s' has already been added.\n", _key)
	} else {
		configItems[_key] = _item
	}

	// Get section map
	sectionMap, sectionFound := configItemCollection[_item.GetSection()]
	if !sectionFound {
		sectionMap = make(map[string]IConfigurationItem)
		configItemCollection[_item.GetSection()] = sectionMap
	}

	// Check if item doesn't already exists
	_, itemFound := sectionMap[_key]
	if itemFound {
		return false, fmt.Sprintf("ConfigurationItem '%s' in section '%s' already exists.\n", _item.GetName(), _item.GetSection())
	} else {
		sectionMap[_item.GetName()] = _item
	}

	return true, ""
}

func getConfigurationItemByKey(_key string) (IConfigurationItem, error) {
	item, found := configItems[_key]
	if !found {
		return nil, fmt.Errorf("Unable to find ConfigurationItem with key '%s'\n", _key)
	}

	return item, nil
}
