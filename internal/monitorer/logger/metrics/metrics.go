package metrics

import (
	"errors"
	"fmt"
	"time"

	"github.com/supermarine1377/check-http-status/internal/models"
)

type Metrics struct {
	totalRequests       int
	successfulResponses int
	failedResponses     int
	totalResponseTime   time.Duration
}

func (m *Metrics) Update(res *models.Response) {
	m.totalRequests++
	m.totalResponseTime += res.ResponseTime

	if res.IsOK() {
		m.successfulResponses++
	} else {
		m.failedResponses++
	}
}

func (m *Metrics) Summarize() (string, error) {
	if m.totalRequests == 0 {
		return "", errors.New("no requests were made")
	}
	averageResponseTime := m.totalResponseTime / time.Duration(m.totalRequests)
	return fmt.Sprintf(
		"Total Requests: %d\nSuccessful Responses: %d (%.2f%%)\nFailed Responses: %d (%.2f%%)\nAverage Response Time: %s\n",
		m.totalRequests,
		m.successfulResponses,
		float64(m.successfulResponses)/float64(m.totalRequests)*100,
		m.failedResponses,
		float64(m.failedResponses)/float64(m.totalRequests)*100,
		averageResponseTime.String(),
	), nil
}
