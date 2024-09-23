package xelastic

import (
	"crypto/tls"
	"net/http"

	"github.com/dream-mo/prom-elastic-alert/conf"
	"github.com/dream-mo/prom-elastic-alert/utils/logger"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
)

type ElasticClient interface {
	FindByDSL(index string, dsl string, source []string) ([]any, int, int)
	CountByDSL(index string, dsl string) (int, int)
}

func NewElasticClient(esConfig conf.EsConfig, version string) ElasticClient {
	// skip es https ca checks
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	client, err := elasticsearch7.NewClient(elasticsearch7.Config{
		Addresses: esConfig.Addresses,
		Username:  esConfig.Username,
		Password:  esConfig.Password,
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	})
	if err != nil {
		logger.Logger.Errorln(err)
		return nil
	}
	// check es connects（if user or passwd is wrong, print error）
	res, err := client.Info()
	if err != nil {
		logger.Logger.Errorln(err)
	}
	defer res.Body.Close()
	if res.IsError() {
		logger.Logger.Errorln(res.String())
		return nil
	}

	c := &ElasticClientV7{
		client: client,
	}
	return c
}
