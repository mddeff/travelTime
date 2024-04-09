package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"gopkg.in/yaml.v2"


)


type Location struct {
	Name     string `yaml:"name"`
	Location string `yaml:"location"`
}

type Config struct {
	Sources      []Location `yaml:"sources"`
	Destinations []Location `yaml:"destinations"`
}



func readConfigFile(filePath string) (*Config, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func convertToFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	default:
		return 0, fmt.Errorf("unexpected type: %T", value)
	}
}

func calculateTravelTime(sourceLocation, destinationLocation string) (float64, float64, error) {
	apiURL := fmt.Sprintf("https://router.project-osrm.org/route/v1/driving/%s;%s", sourceLocation, destinationLocation)

	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]interface{}
		err := yaml.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return 0, 0, err
		}

		routes := result["routes"].([]interface{})
		if len(routes) > 0 {
			route := routes[0].(map[interface{}]interface{})
			durationValue, _ := convertToFloat64(route["duration"])
			distanceValue, _ := convertToFloat64(route["distance"])

			return durationValue, distanceValue, nil
		}
	}

	return 0, 0, fmt.Errorf("Error: Unable to fetch travel time. Status Code: %d", resp.StatusCode)
}



func main() {
	configFilePath := "config.yaml" // Change this to the actual path of your config file
	config, err := readConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	for _, source := range config.Sources {
		fmt.Printf("Source: %s\n", source.Name)
		for _, destination := range config.Destinations {
			travelTime, distance, err := calculateTravelTime(source.Location, destination.Location)
			if err != nil {
				log.Printf("Error: %v", err)
			} else {
				hours := int(travelTime) / 3600
				minutes := int(travelTime) % 3600 / 60
				fmt.Printf("    %s: %dh%dm, %.2fMi\n", destination.Name, hours, minutes, distance/1609.34)
			}
		}
	}
}