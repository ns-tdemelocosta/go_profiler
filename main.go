package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"twitter/gopsutil"
	"twitter/models"
	"twitter/repository"
)

type Vehicle struct {
	Type          string
	Wheels        int
	NumberOfDoors int `json:"number_of_doors"` // JSON tag is required if the JSON string is different than the Field name
}

func main() {

	//get process info
	process, _ := gopsutil.GetProcessesInfo()

	for _, p := range process {
		fmt.Printf("Process Name: %s, CPU Usage: %f, Memory Usage: %f, Process ID: %d\n", p.Name, p.CPUUsage, p.Memory, p.ProcessId)
	}
	fmt.Println("Hello, World!")
	url := "https://gorest.co.in/public/v2/users"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	response := `[{"type": "Car","wheels": 4, "number_of_doors": 4},{"type": "Motorcycle","wheels": 2, "number_of_doors": 0}]`
	var vehicles []Vehicle

	// And we unmarshal our JSON, assigning the result to vehicles
	json.Unmarshal([]byte(response), &vehicles)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// body to array or users
	var users []models.User
	err = json.Unmarshal(body, &users)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Gender)
	}

	//Kafka consumer
	kafka := repository.NewKafka()
	kafka.Consume("quickstart-events", true)

}
