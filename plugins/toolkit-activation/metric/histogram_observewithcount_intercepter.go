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

package metric

import (
	"github.com/hegeng1212/skywalking-go/plugins/core/metrics"
	"github.com/hegeng1212/skywalking-go/plugins/core/operator"
)

type HistogramObserveWithCountInterceptor struct{}

func (h *HistogramObserveWithCountInterceptor) BeforeInvoke(_ operator.Invocation) error {
	return nil
}

func (h *HistogramObserveWithCountInterceptor) AfterInvoke(invocation operator.Invocation, result ...interface{}) error {
	enhanced, ok := invocation.CallerInstance().(operator.EnhancedInstance)
	if !ok {
		return nil
	}

	histogram, ok := enhanced.GetSkyWalkingDynamicField().(metrics.Histogram)
	if ok && histogram != nil {
		histogram.ObserveWithCount(invocation.Args()[0].(float64), invocation.Args()[1].(int64))
	}
	return nil
}
