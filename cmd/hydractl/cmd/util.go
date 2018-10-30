package cmd

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func checkYAMLKind(data []byte) (string, error) {
	m := make(map[string]interface{})
	err := yaml.Unmarshal(data, &m)
	checkError(err)

	kind, ok := m["kind"]
	if !ok {
		return "", fmt.Errorf("No kind is specified in the configuration file")
	}
	switch kind.(type) {
	default:
		return "", fmt.Errorf("Bad kind value")
	case string:
		return kind.(string), nil
	}
}
