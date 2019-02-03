// Copyright 2019, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	gokitLog "github.com/go-kit/kit/log"
	"github.com/prometheus/prometheus/config"

	"github.com/orijtech/promreceiver"

	agentmetricspb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/metrics/v1"
)

func main() {
	source := `
scrape_configs:
  - job_name: 'ocservice_example'

    scrape_interval: 5s
    static_configs:
        - targets: ['localhost:9988']

  - job_name: 'ocjdbc'

    scrape_interval: 4s
    static_configs:
        - targets: ['localhost:8889']
`
	cfg, err := config.Load(source)
	if err != nil {
		log.Fatalf("Failed to load the Prometheus YAML configuration: %v", err)
	}
	nopLogger := gokitLog.NewNopLogger()

	jems := new(jsonEncodingMetricsSink)

	recv, _ := promreceiver.ReceiverFromConfig(context.Background(), jems, cfg, nopLogger, promreceiver.WithBufferPeriod(5*time.Second))
	defer recv.Cancel()

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt)

	// Now wait until Ctrl-C.
	<-shutdownCh
}

type jsonEncodingMetricsSink int

var _ promreceiver.MetricsSink = (*jsonEncodingMetricsSink)(nil)

func (jems *jsonEncodingMetricsSink) ReceiveMetrics(ctx context.Context, ereq *agentmetricspb.ExportMetricsServiceRequest) error {
	blob, _ := json.MarshalIndent(ereq, "", "  ")
	log.Printf("%s\n\n", blob)
	return nil
}
