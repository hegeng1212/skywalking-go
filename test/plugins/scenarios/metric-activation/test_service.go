// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
	"github.com/hegeng1212/skywalking-go/toolkit/metric"
)

func testCounter() {
	requestCounter := metric.NewCounter("request_counter")
	requestCounter.Inc(1)

	requestCounterWithLabel := metric.NewCounter("request_counter_with_label", metric.WithLabels("foo", "bar"))
	requestCounterWithLabel.Inc(1)
}

func testGauge() {
	metric.NewGauge("cpu_usage_gauge", func() float64 {
		return 50
	})

	metric.NewGauge("cpu_usage_gauge_with_label", func() float64 {
		return 30
	}, metric.WithLabels("foo", "bar"))
}

func testHistogram() {
	steps := []float64{5, 20, 50, 100, 500, 1000}
	histogram := metric.NewHistogram("request_duration", steps)
	histogram.Observe(10)
	histogram.Observe(200)
	histogram.Observe(1500)
	histogram.ObserveWithCount(20, 1)
	histogram.ObserveWithCount(20, 2)

	histogramWithLabel := metric.NewHistogram("request_duration_with_label", steps, metric.WithLabels("foo", "bar"))
	histogramWithLabel.Observe(10)
	histogramWithLabel.Observe(200)
	histogramWithLabel.Observe(1500)
	histogramWithLabel.ObserveWithCount(100, 1)
	histogramWithLabel.ObserveWithCount(500, 2)
}
