package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	yaml "github.com/goccy/go-yaml"
)


func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func WriteYaml(y interface{}, file string) error {
	yMarshalled, err := yaml.Marshal(y)
	if err != nil {
		return err
	}
	fmt.Printf("---\n%s", string(yMarshalled))
	if file != "" {
		err = ioutil.WriteFile(file, yMarshalled, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

// Return Env var if present, or default if not
func EnvORDefault(v, def string) string {
	val, present := os.LookupEnv(v)
	if present {
		return val
	}
	return def
}

// Split keys by prefix
func SplitKeys(data map[string]string, keyPrefix string) (map[string]string, map[string]string) {
	data1 := map[string]string{}
	data2 := map[string]string{}

	for k, v := range data {
		if keyPrefix != "" && strings.HasPrefix(k, keyPrefix) {
			data2[k] = EnvORDefault(k, v)
		} else {
			data1[k] = EnvORDefault(k, v)
		}
	}

	return data1, data2
}
