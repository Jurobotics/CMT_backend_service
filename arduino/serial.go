package arduino

import (
	"log"

	"go.bug.st/serial"
)

var (
	port serial.Port
)

func Setup() {
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, _ = serial.Open("/dev/ttyUSB0", mode) // TODO: Change to device ID, check if working on pi

}

func SendData(port serial.Port, data []byte) Response {
	n, err := port.Write(data)
	if err != nil {
		log.Fatal(err)
		return Response{
			Data:   err.Error(),
			Status: false,
		}
	}
	return Response{
		Data:   n,
		Status: true,
	}
}

// TODO: IMPL. me
func ReadData(port serial.Port) Response {
	return Response{}
}
