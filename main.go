package main

import (
	"Vehicle-Rental-App/services"
	"fmt"
)

func main() {

	rentalService := services.NewRentalService()
	fmt.Println("Available Vehicles !!")
	rentalService.ListAvailableVehicles()

	fmt.Println("Renting a vehicle !!")
	err := rentalService.RentVehicle("bike-123")
	if err != nil {
		fmt.Println("Error : ", err)
	}

	fmt.Println("Available Vehicles After Renting !!")
	rentalService.ListAvailableVehicles()

}
