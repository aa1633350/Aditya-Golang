package models

import "fmt"

type Vehicle interface {
	GetID() string
	GetModel() string
	IsAvailable() bool
	Book() error
}
type Car struct {
	ID        string
	Model     string
	Available bool
}

type Bike struct {
	ID        string
	Model     string
	Available bool
}

// Getters of Car

func (c *Car) GetID() string {
	return c.ID
}

func (c *Car) GetModel() string {
	return c.Model
}

func (c *Car) IsAvailable() bool {
	return c.Available
}

func (c *Car) Book() error {
	if !c.IsAvailable() {
		return fmt.Errorf("car %s is not available", c.GetID())
	}
	c.Available = false
	return nil
}

// Getters of Bike

func (b *Bike) GetID() string {
	return b.ID
}

func (b *Bike) GetModel() string {
	return b.Model
}

func (b *Bike) IsAvailable() bool {
	return b.Available
}

func (b *Bike) Book() error {
	if !b.IsAvailable() {
		return fmt.Errorf("car %s is not available", b.GetID())
	}
	b.Available = false
	return nil
}
