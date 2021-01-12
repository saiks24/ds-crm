package app

type DsCrm struct {
	svc ServiceProvider
}

func NewDsCrm(svc ServiceProvider) *DsCrm {
	return &DsCrm{svc: svc}
}
