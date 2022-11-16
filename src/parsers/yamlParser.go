package yamlparser

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Destructured struct {
	Mapping []struct {
		From string `yaml:"from"`
		To   string `yaml:"to"`
	} `yaml:"Mapping"`
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
