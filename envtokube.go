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
	"os"
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
	namespacePtr := flag.String("namespace", "default", "Specifies the Namespace for the kube secrets file.")
	flag.Parse()
	action := flag.Args()

	argCount := len(action)

	var newFilename string
	if argCount == 1 {
		newFilename = extractName(action[0])
	} else if argCount == 2 {
		newFilename = extractName(action[1])
	} else {
		log.Fatal("At minimum, and .env file must be provided")
		os.Exit(1)
	}

	// Load in the ENVs
	var myEnv map[string]string
	myEnv, err := godotenv.Read(action[0])

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}


	createNewSecrets(newFilename, myEnv, *namespacePtr)
}


func createNewSecrets(filename string, keys map[string]string, namespace string) error {
	s            := secrets{}
	s.APIVersion = "v1"
	s.Kind       = "Secret"
	s.Metadata.Namespace = namespace
	s.Type       = "Opaque"
	s.Data = make(map[string]string)

	s.Metadata.Name = filename

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

	writeErr := ioutil.WriteFile(filename+".yml", newYML, 0644)

	return writeErr
}

func extractName(filename string) string {
	ext := filepath.Ext(filename)
	// The secret name will have the directory and ext stripped
	// /tmp/test.yml becomes test
	name := filepath.Base(filename[0:len(filename)-len(ext)])

	return name
}