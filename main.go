package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/iki-rumondor/assignment3-GLNG-KS-08-08/internal/adapter/database"
	"github.com/iki-rumondor/assignment3-GLNG-KS-08-08/internal/adapter/http/response"
	"github.com/iki-rumondor/assignment3-GLNG-KS-08-08/internal/application"
	"github.com/iki-rumondor/assignment3-GLNG-KS-08-08/internal/repository"
)

var condService *application.ConditionServices

func main() {

	StartConditionServices()
	fmt.Println("condition services is running...")

	go UpdateWaterWind()

	select {}

}

func StartConditionServices() {
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal("failed to connect postgres database")
		return
	}

	// db.Debug().AutoMigrate(domain.Condition{})

	condRepo := repository.NewConditionRepo(db)
	condService = application.NewConditionServices(condRepo)
}

func UpdateWaterWind() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		res, err := condService.UpdateConditon(rand.Intn(100)+1, rand.Intn(100)+1)
		if err != nil {
			log.Fatal("failed to update condition:", err.Error())
			return
		}
		printConditionInformation(*res)
		time.Sleep(2 * time.Second)
	}

}

func printConditionInformation(res response.Condition) {

	var statusWater string
	var statusWind string

	switch {
	case res.Water <= 5:
		statusWater = "aman"
	case res.Water > 5 && res.Water <= 8:
		statusWater = "siaga"
	default:
		statusWater = "bahaya"
	}

	switch {
	case res.Wind <= 6:
		statusWind = "aman"
	case res.Wind > 6 && res.Wind <= 15:
		statusWind = "siaga"
	default:
		statusWind = "bahaya"
	}

	condMarshaled, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
	}

	fmt.Println(string(condMarshaled))
	fmt.Println("status water:", statusWater)
	fmt.Println("status wind:", statusWind)

}
