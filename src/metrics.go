package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/infra-integrations-sdk/metric"
)

var metricsDefinition = map[string][]interface{}{
	"software.version":           {"zk_version", metric.ATTRIBUTE},
	"avg_latency":                {"zk_avg_latency", metric.GAUGE},
	"max_latency":                {"zk_max_latency", metric.GAUGE},
	"min_latency":                {"zk_min_latency", metric.GAUGE},
	"packets_received":           {"zk_packets_received", metric.GAUGE},
	"packets_sent":               {"zk_packets_sent", metric.GAUGE},
	"outstanding_requests":       {"zk_outstanding_requests", metric.GAUGE},
	"server_state":               {"zk_server_state", metric.ATTRIBUTE},
	"znode_count":                {"zk_znode_count", metric.GAUGE},
	"watch_count":                {"zk_watch_count", metric.GAUGE},
	"ephemerals_count":           {"zk_ephemerals_count", metric.GAUGE},
	"approximate_data_size":      {"zk_approximate_data_size", metric.GAUGE},
	"followers":                  {"zk_followers", metric.GAUGE},
	"synced_followers":           {"zk_synced_followers", metric.GAUGE},
	"pending_syncs":              {"zk_pending_syncs", metric.GAUGE},
	"open_file_descriptor_count": {"zk_open_file_descriptor_count", metric.GAUGE},
	"max_file_descriptor_count":  {"zk_max_file_descriptor_count", metric.GAUGE},
}

func asValue(value string) interface{} {
	if i, err := strconv.Atoi(value); err == nil {
		return i
	}

	if f, err := strconv.ParseFloat(value, 64); err == nil {
		return f
	}

	if b, err := strconv.ParseBool(value); err == nil {
		return b
	}
	return value
}

func populateMetrics(sample *metric.MetricSet, metrics map[string]interface{}, metricsDefinition map[string][]interface{}) error {
	if len(metrics) == 0 {
		log.Debug("Metrics data from status module not found")
	}
	for metricName, metricInfo := range metricsDefinition {
		rawSource := metricInfo[0]
		metricType := metricInfo[1].(metric.SourceType)

		var rawMetric interface{}
		var ok bool

		switch source := rawSource.(type) {
		case string:
			rawMetric, ok = metrics[source]
		default:
			log.Warn("Invalid raw source metric for %s", metricName)
			continue
		}

		if !ok {
			log.Debug("Can't find raw metrics in results for %s", metricName)
			continue
		}
		err := sample.SetMetric(metricName, rawMetric, metricType)

		if err != nil {
			log.Warn("Error setting value: %s", err)
			continue
		}
	}

	if len(*sample) < 2 {
		return fmt.Errorf("no metrics were found on the status response. Probably caused by a wrong response format")
	}
	return nil
}

func checkNCExists() {
	path, err := exec.LookPath("nc")
	if err != nil {
		log.Debug("didn't find 'nc' executable\n")
	} else {
		log.Debug("'nc' executable is in '%s'\n", path)
	}
}

func getMetricsData(sample *metric.MetricSet) error {
	checkNCExists()

	cmd := exec.Command("nc", "localhost", "2181")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Stdin = strings.NewReader("mntr")

	err := cmd.Run()
	if err != nil {
		log.Error("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if errStr != "" {
		log.Debug("err:\n%s\n", errStr)
	}

	rawMetrics := make(map[string]interface{})
	temp := strings.Split(outStr, "\n")
	for _, line := range temp {
		splitedLine := strings.Fields(line)
		if len(splitedLine) != 2 {
			continue
		}
		rawMetrics[splitedLine[0]] = asValue(strings.TrimSpace(splitedLine[1]))
	}

	return populateMetrics(sample, rawMetrics, metricsDefinition)
}
