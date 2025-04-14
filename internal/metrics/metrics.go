package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
)

var HTTPRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "http_request_total",
	Help: "Number of request",
}, []string{"path"})

var HTTPRequestTime = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name:       "http_request_time",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
}, []string{"path"})

var PVZCounter = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "num_of_pvzs",
	Help: "Number of pvzs in db",
}, []string{"pvz"})

var ReceptionCounter = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "num_of_receptions",
	Help: "Number of receptions in db",
}, []string{"reception"})

var ProductCounter = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "num_of_products",
	Help: "Number of products in db",
}, []string{"product"})

type Metrics struct {
	metricsRecorder MetricsRecorder
}

func NewMetrics(metricsRecorder MetricsRecorder) *Metrics {
	return &Metrics{
		metricsRecorder: metricsRecorder,
	}
}

func (r *Metrics) CheckMetricsFromDBCircle() {
	const op = "internal.metrics.CheckMetricsFromDB()"
	for {
		pvzs, receptions, products, err := r.metricsRecorder.GetValuesForMetrics()
		if err != nil {
			log.Error().Err(err).Msgf("%s: failed get number of pvz", op)
		}

		PVZCounter.WithLabelValues("pvz").Set(float64(pvzs))

		ReceptionCounter.WithLabelValues("reception").Set(float64(receptions))

		ProductCounter.WithLabelValues("products").Set(float64(products))

		time.Sleep(15 * time.Second)
	}
}
