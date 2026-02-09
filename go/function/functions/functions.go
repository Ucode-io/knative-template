package helper

import (
	"function/models"
)

// here you can add your own functions
var Handlers = map[string]models.HandlerFunc{

	// Healt Check
	"health_check": HealthCheck,
}
