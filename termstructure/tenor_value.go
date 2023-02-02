package termstructure

type TenorValue struct {
	tenor float64
	value float64
}

func (tv TenorValue) Tenor() float64 {
	return tv.tenor
}

func (tv TenorValue) Value() float64 {
	return tv.value
}

func NewTenorValue(tenor, value float64) TenorValue {
	return TenorValue{
		tenor: tenor,
		value: value,
	}
}
