package termstructure

// TenorValue represents a tenor-value pair with the
// tenor expressed as year fraction.
type TenorValue struct {
	tenor float64
	value float64
}

// Tenor returns the tenor.
func (tv TenorValue) Tenor() float64 {
	return tv.tenor
}

// Value returns the value.
func (tv TenorValue) Value() float64 {
	return tv.value
}

// NewTenorValue returns a new tenor-value pair.
func NewTenorValue(tenor, value float64) TenorValue {
	return TenorValue{
		tenor: tenor,
		value: value,
	}
}
