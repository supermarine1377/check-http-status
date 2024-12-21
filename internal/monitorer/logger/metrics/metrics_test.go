package metrics

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/supermarine1377/check-http-status/internal/models"
)

func TestMetrics_Update(t *testing.T) {
	m := Metrics{}
	res := &models.Response{
		ResponseTime: 100,
		Status:       "200 OK",
	}

	m.Update(res)

	want := Metrics{
		totalRequests:       1,
		successfulResponses: 1,
		failedResponses:     0,
		totalResponseTime:   100,
	}
	assert.Equal(t, want, m)

	res = &models.Response{
		ResponseTime: 200,
		Status:       "500 Internal Server Error",
	}

	m.Update(res)

	want = Metrics{
		totalRequests:       2,
		successfulResponses: 1,
		failedResponses:     1,
		totalResponseTime:   300,
	}
	assert.Equal(t, want, m)
}

func TestMetrics_Summarize(t *testing.T) {
	type fields struct {
		totalRequests       int
		successfulResponses int
		failedResponses     int
		totalResponseTime   time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name:    "no requests were made",
			fields:  fields{},
			want:    "",
			wantErr: true,
		},
		{
			name: "one request was made",
			fields: fields{
				totalRequests:       1,
				successfulResponses: 1,
				failedResponses:     0,
				totalResponseTime:   time.Second,
			},
			want:    "Total Requests: 1\nSuccessful Responses: 1 (100.00%)\nFailed Responses: 0 (0.00%)\nAverage Response Time: 1s\n",
			wantErr: false,
		},
		{
			name: "two request was made",
			fields: fields{
				totalRequests:       2,
				successfulResponses: 1,
				failedResponses:     1,
				totalResponseTime:   time.Second,
			},
			want:    "Total Requests: 2\nSuccessful Responses: 1 (50.00%)\nFailed Responses: 1 (50.00%)\nAverage Response Time: 500ms\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Metrics{
				totalRequests:       tt.fields.totalRequests,
				successfulResponses: tt.fields.successfulResponses,
				failedResponses:     tt.fields.failedResponses,
				totalResponseTime:   tt.fields.totalResponseTime,
			}
			got, err := m.Summarize()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
