package main

import (
	sdk_args "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/infra-integrations-sdk/sdk"
)

type argumentList struct {
	sdk_args.DefaultArgumentList
	Host string `default:"localhost" help:"Hostname or IP where zookeeper node is running."`
	Port int    `default:"2181" help:"Port on which zookeeper node is listening."`
	Cmd  string `default:"nc" help:"Command to use for tcp connection to zookeeper [nc | tcp | telnet"`
}

const (
	integrationName    = "com.newrelic.zookeeper"
	integrationVersion = "1.0.0"
)

var args argumentList

func main() {
	log.Debug("Starting Zookeeper integration")
	defer log.Debug("Zookeeper integration exited")

	integration, err := sdk.NewIntegration(integrationName, integrationVersion, &args)
	fatalIfErr(err)

	/*
		if args.All || args.Inventory {
			log.Debug("Fetching data for '%s' integration", integrationName+"-inventory")
			fatalIfErr(setInventory(integration.Inventory))
		}
	*/

	if args.All || args.Metrics {
		log.Debug("Fetching data for '%s' integration", integrationName+"-metrics")
		ms := integration.NewMetricSet("ZookeeperSample")

		fatalIfErr(getMetricsData(ms))
	}

	fatalIfErr(integration.Publish())
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
