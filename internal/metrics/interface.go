package metrics

type MetricsRecorder interface {
	GetValuesForMetrics() (pvzs int, receptions int, products int, err error)
}
