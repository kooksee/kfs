package types

type Status string

const CREATE Status = "CREATE"
const UPDATE Status = "UPDATE"
const DELETE Status = "DELETE"

func (s Status) CREATE() Status {
	return CREATE
}

func (s Status) UPDATE() Status {
	return CREATE
}

func (s Status) DELETE() Status {
	return CREATE
}
