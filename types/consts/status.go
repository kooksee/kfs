package consts

var MetaStatus = status{}

type status struct{ string }

const sCREATE = "CREATE"
const sUPDATE = "UPDATE"
const sDELETE = "DELETE"

func (s status) CREATE() string {
	return sCREATE
}

func (s status) UPDATE() string {
	return sUPDATE
}

func (s status) DELETE() string {
	return sDELETE
}
