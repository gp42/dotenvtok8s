package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
	yaml "github.com/goccy/go-yaml"
)


func TestEnvOrDefault(t *testing.T) {
	fmt.Println("OK")
	variable := "TESTVAR"
	val := "foo"

	os.Unsetenv(variable)
	got := EnvORDefault(variable, val)
    if got != val {
        t.Errorf("EnvORDefault(%s, \"%s\") = %s; want \"%s\"", variable, val, got, val)
    }

	val2 := "bar"
	os.Setenv(variable, val2)
	got = EnvORDefault(variable, val)
    if got != val2 {
        t.Errorf("OS Env %s=%s, EnvORDefault(%s, \"%s\") = %s; want \"%s\"", variable, val2, variable, val, got, val2)
    }
}

func TestSplitKeys(t *testing.T) {
	data := map[string]string {
		"KEY1": "val1",
		"SECRET_KEY1": "secret1",
	}

	data1, data2 := SplitKeys(data, "SECRET_")

	r := map[string]string{ "KEY1": "val1" }
	if ! reflect.DeepEqual(data1, r) {
		t.Errorf("SplitKeys(\"%v\", \"SECRET_\") data1 = \"%v\"; want \"%v\"", data, data1, r)
	}

	r = map[string]string{ "SECRET_KEY1": "secret1" }
	if ! reflect.DeepEqual(data2, r) {
		t.Errorf("SplitKeys(\"%v\", \"SECRET_\") data2 = \"%v\"; want \"%v\"", data, data2, r)
	}
}

func TestWriteYaml(t *testing.T) {
	type metadata struct {
		Name string
	}

	type Cm struct {
		Name		string
		Metadata	metadata
		Data	map[string]string
	}

	cm := &Cm {
		Name: "Testcm",
		Metadata: metadata {
			Name: "Testmeta",
		},
		Data: map[string]string {
			"FOO": "BAR",
		},
	}

	fname := fmt.Sprintf("TestWriteYaml_%v.yaml", time.Now().Unix())
	defer os.Remove(fname)

	err := WriteYaml(cm, fname)
	if err != nil {
		t.Errorf("TestWriteYaml failed to write file: %s", err)
	}

	dat, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Errorf("TestWriteYaml failed to read file: %s", err)
	}

	var c Cm
	err = yaml.Unmarshal(dat, &c)
	if err != nil {
		t.Errorf("TestWriteYaml failed to unmarshal read file: %s", err)
	}

	if ! reflect.DeepEqual(&c, cm) {
		t.Errorf("TestWriteYaml failed - files do not match: %s", err)
	}
}
