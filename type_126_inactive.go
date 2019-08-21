package dmi

type Inactive struct {
	infoCommon
}

func (i Inactive) String() string {
	return "Inactive"
}
