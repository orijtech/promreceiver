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

package promreceiver

import (
	"testing"
)

func TestReceiver(t *testing.T) {
	// Start a Prometheus server with various configurations
}

var pages = []string{
	`
# HELP opdemo_latency The various latencies of the methods
# TYPE opdemo_latency histogram
opdemo_latency_bucket{client="cli",method="repl",le="0"} 0
opdemo_latency_bucket{client="cli",method="repl",le="10"} 56
opdemo_latency_bucket{client="cli",method="repl",le="50"} 272
opdemo_latency_bucket{client="cli",method="repl",le="100"} 482
opdemo_latency_bucket{client="cli",method="repl",le="200"} 497
opdemo_latency_bucket{client="cli",method="repl",le="400"} 535
opdemo_latency_bucket{client="cli",method="repl",le="800"} 588
opdemo_latency_bucket{client="cli",method="repl",le="1000"} 609
opdemo_latency_bucket{client="cli",method="repl",le="1400"} 627
opdemo_latency_bucket{client="cli",method="repl",le="2000"} 630
opdemo_latency_bucket{client="cli",method="repl",le="5000"} 653
opdemo_latency_bucket{client="cli",method="repl",le="10000"} 680
opdemo_latency_bucket{client="cli",method="repl",le="15000"} 695
opdemo_latency_bucket{client="cli",method="repl",le="+Inf"} 702
opdemo_latency_sum{client="cli",method="repl"} 669824.6063739995
opdemo_latency_count{client="cli",method="repl"} 702
# HELP opdemo_process_counts The various counts
# TYPE opdemo_process_counts counter
opdemo_process_counts{client="cli",method="repl"} 702
`,
}
