package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

type conf struct {
	R      string `yaml:"r"`
	P      string `yaml:"p"`
	Strart string `yaml:"start"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func con(key string, element string) {
	if _, err := os.Stat(key + ":\\"); os.IsNotExist(err) {

		// del := exec.Command("net", "use", key+":", "/delete", "/y")
		// if err := del.Run(); err != nil {
		// 	fmt.Println("Error: ", err)
		// }
		c := exec.Command("net", "use", key+":", element, "/p:yes")
		if err := c.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func prog(key string, element string) {
	c := exec.Command(element + key)
	c.Dir = element
	if err := c.Start(); err != nil {
		fmt.Println("Error: ", err)
	}

}

func main() {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	m := make(map[interface{}]map[string]string)

	err = yaml.Unmarshal([]byte(yamlFile), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for key, element := range m["drive"] {
		con(key, element)
	}

	for key, element := range m["command"] {
		prog(key, element)
	}
}
