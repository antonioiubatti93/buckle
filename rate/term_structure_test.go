package rate

type termStructure float64

var _ TermStructure = termStructure(0.0)

func (t termStructure) Value(_ float64) float64 {
	return float64(t)
}

func newTermStructure(c float64) termStructure {
	return termStructure(c)
}
