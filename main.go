package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	a := App{}

	viper.SetConfigFile("./config.json")
	// Searches for config file in given paths and read it
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	testData := viper.Sub("test")
	err := testData.Unmarshal(&conn)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	// fmt.Println(conn.Host)

	a.Initialize("steven", "here", "api_medium_kelvin")

	a.Run(":8000")
}
