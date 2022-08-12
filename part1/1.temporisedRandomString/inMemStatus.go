package main

var status = "initial"

func SetStatus(_status string) {
	status = _status
}

func GetStatus() string {
	return status
}
