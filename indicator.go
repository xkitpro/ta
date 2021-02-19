package ta

type Indicator interface {
	Write(float64) float64
}
