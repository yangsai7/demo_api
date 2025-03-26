package client

import (
	"context"
	"log"

	"github.com/olivere/elastic/v7"
	"github.com/yangsai7/demo_api/config"
)

var EsClient *elastic.Client

func InitElastic() {
	client, err := elastic.NewClient(
		elastic.SetURL(config.GlobalCfg.Elastic.Url),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(config.GlobalCfg.Elastic.Username, config.GlobalCfg.Elastic.Password),
		elastic.SetTraceLog(log.Default()),
	)
	if err != nil {
		panic(err)
	}

	if _, _, err := client.Ping(config.GlobalCfg.Elastic.Url).Do(context.Background()); err != nil {
		panic(err)
	}

	EsClient = client
}
