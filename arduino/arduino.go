package arduino

type Response struct {
	Data   any
	Status bool
}

type ArduinoConnector interface {
	Setup() bool
	SendData(port any, data []byte) Response
	ReadData(port any) Response
}
