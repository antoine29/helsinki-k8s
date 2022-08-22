package inMemo

var status = "initial status (not overwritten yet)"

func SetStatus(_status string) {
	status = _status
}

func GetStatus() string {
	return status
}
