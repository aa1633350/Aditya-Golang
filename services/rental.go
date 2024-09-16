package services

import (
	"Vehicle-Rental-App/models"
	"fmt"
)

type RentalService struct {
	vehicles map[string]models.Vehicle
}

func NewRentalService() *RentalService {
	return &RentalService{
		vehicles: map[string]models.Vehicle{
			"car-123": &models.Car{
				ID:        "123",
				Model:     "Toyota Camry",
				Available: true,
			},
			"bike-123": &models.Bike{
				ID:        "bike-123",
				Model:     "Yamaha",
				Available: true,
			},
		},
	}
}

// List all available vehicles for rent.

func (r *RentalService) ListAvailableVehicles() {
	for _, vehicle := range r.vehicles {
		if vehicle.IsAvailable() {
			fmt.Printf("ID: %s, Type: %s\n", vehicle.GetID(), vehicle.GetModel())
		}
	}
}

func (r *RentalService) RentVehicle(vehicleID string) error {
	vehicle, exists := r.vehicles[vehicleID]
	if !exists {
		return fmt.Errorf("vehicle with ID %s does not exist", vehicleID)
	}
	err := vehicle.Book()
	if err != nil {
		return err
	}
	fmt.Printf("%s rented successfully!\n", vehicle.GetModel())
	return nil
}
