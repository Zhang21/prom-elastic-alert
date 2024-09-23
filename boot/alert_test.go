package boot

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlertSampleMessage(t *testing.T) {
	// a smaple json payload, the contents of redis key
	payload := `
	{
	  "es": {
	    "address": ["http://es1:9200", "https://es2:9200"],
	    "username": "user01",
	    "password": "pwd01",
	    "conn_timeout": 60,
	    "version": "v7"
	  },
	  "index": "index-test01-*",
	  "ids": ["id1", "id2", "id3"]
	}
	`
	message := AlertSampleMessage{}
	err := json.Unmarshal([]byte(payload), &message)
	assert.Nil(t, err)
}
