package models

import (
	"reflect"
	"testing"
)

func TestCar_Map(t *testing.T) {
	car := Car{
		ID:           1,
		Name:         "Test Car",
		Year:         2023,
		SellingPrice: 25000.00,
		Transmission: "Automatic",
	}

	expected := map[string]interface{}{
		"id":            "1",
		"name":          "Test Car",
		"year":          "2023",
		"selling_price": "25000",
		"transmission":  "Automatic",
	}

	result := car.Map()

	for key, value := range expected {
		if result[key] != value {
			t.Errorf("Map function produced incorrect result for key %s. Expected %s but got %v", key, value, result[key])
		}
	}
}

func TestCar_Names(t *testing.T) {
	car := Car{}

	expected := []string{"id", "name", "year", "selling_price", "transmission"}

	result := car.Names()

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Names function produced incorrect result. Expected %v but got %v", expected, result)
	}
}
