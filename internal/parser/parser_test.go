package parser

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestParserBase(t *testing.T) {
	data, err := ioutil.ReadFile("_testdata/base.tl")
	if err != nil {
		t.Fatal(err)
	}
	schema, err := Parse(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range schema.Classes {
		t.Logf("Class %s: %s", c.Name, c.Description)
	}
	for _, d := range schema.Types {
		t.Logf("%s = %s (0x%x)", d.Definition.Name, d.Definition.Type, d.Definition.ID)
		for _, a := range d.Annotations {
			t.Logf(" %s: %s", a.Key, a.Value)
		}
	}
}

func TestParserError(t *testing.T) {
	data, err := ioutil.ReadFile("_testdata/Error.tl")
	if err != nil {
		t.Fatal(err)
	}
	schema, err := Parse(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range schema.Classes {
		t.Logf("Class %s: %s", c.Name, c.Description)
	}
	for _, d := range schema.Types {
		t.Logf("%s = %s", d.Definition.Name, d.Definition.Type)
		for _, a := range d.Annotations {
			t.Logf(" %s: %s", a.Key, a.Value)
		}
	}
}

func TestParser(t *testing.T) {
	data, err := ioutil.ReadFile("_testdata/td_api.tl")
	if err != nil {
		t.Fatal(err)
	}
	schema, err := Parse(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range schema.Classes {
		t.Logf("Class %s: %s", c.Name, c.Description)
	}
	for _, d := range schema.Types {
		t.Logf("%s = %s", d.Definition.Name, d.Definition.Type)
		for _, a := range d.Annotations {
			t.Logf(" %s: %s", a.Key, a.Value)
		}
	}
}
