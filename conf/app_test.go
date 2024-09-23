package conf

import (
	"os"
	"testing"

	"github.com/jessevdk/go-flags"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestEsConf(t *testing.T) {
	// a sample yaml payload
	payload := `
es:
  addresses:
    - "http://es1:9200"
    - "https://es2:9200"
  username: ""
  password: ""
  conn_timeout: 300
  version: "v7"
`
	message := EsConfig{}
	err := yaml.Unmarshal([]byte(payload), &message)
	assert.Nil(t, err)
}

func TestAppConf(t *testing.T) {
	// a sample yaml payload
	payload := `
exporter:
  enabled: true
  listen_addr: ":9003"
loader:
  type: "FileLoader"
  config:
    rules_folder: "rules"
    rules_folder_recursion: false
alert:
  alertmanager:
    url: "http://alertmanager:9093/api/v2/alerts"
    basic_auth:
      username: ""
      password: ""
  generator:
    base_url: "http://localhost:9003/alert/message"
    expire:
      days: 1
redis:
  addr: "docker.for.mac.host.internal"
  port: 6379
  password: ""
  db: 0
run_every:
  seconds: 10
buffer_time:
  minutes: 10
alert_time_limit:
  minutes: 10
max_scrolling_count: 5
`
	message := AppConfig{}
	err := yaml.Unmarshal([]byte(payload), &message)
	assert.Nil(t, err)
}

func TestFlagOption(t *testing.T) {
	var opts FlagOption
	assert := assert.New(t)

	// defalut args
	os.Args = []string{"cmd"}
	_, err := flags.ParseArgs(&opts, os.Args)
	assert.Nil(err)
	assert.Equal("./config.yaml", opts.ConfigPath)
	assert.False(opts.Debug)
	assert.Equal("info", opts.Verbose)
	assert.Equal("", opts.Rule)
	assert.Equal("PRC", opts.Zone)

	// custom args
	os.Args = []string{"cmd", "--config=config.yaml", "--debug", "--zone=UTC"}
	_, err = flags.ParseArgs(&opts, os.Args)
	assert.Nil(err)
	assert.Equal("config.yaml", opts.ConfigPath)
	assert.True(opts.Debug)
	assert.Equal("UTC", opts.Zone)
}
