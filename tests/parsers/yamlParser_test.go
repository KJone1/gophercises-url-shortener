package yamlparser

import (
	"reflect"
	"testing"

	parser "github.com/KJone1/gophercises-url-shortener/src/parsers"
)

func TestYamlType(t *testing.T) {

	test := parser.Yaml("../../routeFile.yaml")
	got := reflect.TypeOf(test)
	want := reflect.TypeOf(parser.Destructured{})

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestYamlValidate(t *testing.T) {

	test := parser.Yaml("../../routeFile.yaml")
	got := test.Mapping

	if got == nil {
		t.Error("Yaml is not valid -> missing 'Mapping' field")
	}
	for _, v := range got {

		if v.To == "" {
			t.Error("Yaml is not valid -> missing 'to' field")
		}
		if v.From == "" {
			t.Error("Yaml is not valid -> missing 'from' field")
		}
	}

}
