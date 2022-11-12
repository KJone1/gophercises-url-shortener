package yamlparser

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Destructured []struct {
	From string `yaml:"path"`
	To   string `yaml:"redirect"`
}

func Yaml(file string) Destructured {
	rf, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	t := Destructured{}

	log.Printf("Reading: %s\n", file)

	err = yaml.Unmarshal(rf, &t)
	if err != nil {
		log.Fatalf("Error parsing yaml: \"%s\", %v", file, err)
	}

	return t
}