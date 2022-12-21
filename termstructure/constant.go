package termstructure

type Constant struct {
	value float64
}

var (
	_ termStructure               = Constant{}
	_ differentiableTermStructure = Constant{}
)

func (c Constant) Value(_ float64) float64 {
	return c.value
}

func (c Constant) Gradient(_ float64) float64 {
	return 0.0
}

func NewConstant(value float64) *Constant {
	return &Constant{
		value: value,
	}
}
