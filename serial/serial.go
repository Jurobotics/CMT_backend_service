package serial_controller

import (
	"go.bug.st/serial"
)

type Response struct {
	Data   any
	Status bool
}

/*
TODO: Should find the right PORT on its own or use the mode argument and configure the serial communication
*/
func SetupSerial(mode serial.Mode)

/*
TODO: Should send the recipe data to the Arduino and return a success Response or the Error Response
*/
func SendData(data any) Response

/*
TODO: Should listen to the given PORT and return the data in the response format
*/
func ReadData() Response
