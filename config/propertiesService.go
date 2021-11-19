package config

import (
	"github.com/magiconair/properties"
	"log"
	"os"
)

// LoadProperties Loads the Properties from properties.yaml file
func LoadProperties() *properties.Properties {
	myDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	log.Println(myDir)

	propertiesInstance := properties.MustLoadFile(myDir+"/resources/properties.yaml", properties.UTF8)
	return propertiesInstance
}
