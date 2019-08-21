package dmi

type EndOfTable struct {
	infoCommon
}

func (e EndOfTable) String() string {
	return "End-of-Table"
}
