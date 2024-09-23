package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestRule(t *testing.T) {
	// a sample yaml paylod
	payload := `
unique_id: "NginxErrorLog"
enabled: true
es:
  addresses:
    - "http://localhost:9200"
  username: ""
  password: ""
  conn_timeout: 300
  version: "v7"
index: "nginx-error-*"
run_every:
  seconds: 5
query:
  type: "frequency"
  query_string: '$query_string'
  config:
    timeframe:
      minutes: 3
    num_events: 2
  labels:
    alertname: "NginxErrorLog"
    instance: "localhost"
    severity: "warning"
    for_time: "2min"
    threshold: "3"
  annotations:
    description: "Nginx error日志条数 {{ .value }} > {{ .threshold }}"
    summary: "Nginx错误日志告警"
`
	message := Rule{}
	err := yaml.Unmarshal([]byte(payload), &message)
	assert.Nil(t, err)
}
