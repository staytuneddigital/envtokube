package main

import (
	"github.com/joho/godotenv"
	"log"
	"fmt"
	"flag"
	"io/ioutil"
	"github.com/go-yaml/yaml"
	"path/filepath"
	"encoding/base64"
)


type secrets struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Type       string `yaml:"type"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	}  `yaml:"metadata"`
	Data map[string]string `yaml:"data"`
	updateValue string
}



func main() {
	flag.Parse()
	action := flag.Args()

	fmt.Println(action)

	var myEnv map[string]string
	myEnv, err := godotenv.Read(action[0])

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	createNewSecrets(action[1], myEnv)
}


func createNewSecrets(filename string, keys map[string]string) error {
	s            := secrets{}
	s.APIVersion = "v1"
	s.Kind       = "Secret"
	s.Type       = "Opaque"
	s.Data = make(map[string]string)

	ext := filepath.Ext(filename)
	// The secret name will have the directory and ext stripped
	// /tmp/test.yml becomes test
	s.Metadata.Name = filepath.Base(filename[0:len(filename)-len(ext)])


	for key, value := range keys {
		updateValue := base64.StdEncoding.EncodeToString([]byte(value))
		s.Data[key] = updateValue
	}

	s.writeSecrets(filename)

	return nil
}


func (s *secrets) writeSecrets(filename string) error {
	// If no namespace is set, use default
	if s.Metadata.Namespace == "" {
		s.Metadata.Namespace = "default"
	}

	newYML, err := yaml.Marshal(&s)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	writeErr := ioutil.WriteFile(filename, newYML, 0644)

	return writeErr
}