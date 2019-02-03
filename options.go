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
	"time"
)

type Option interface {
	withCustomAppender(*customAppender)
}

type withBufferPeriod time.Duration

func (wbp *withBufferPeriod) withCustomAppender(ca *customAppender) {
	ca.metricsBufferPeriod = time.Duration(*wbp)
}

var _ Option = (*withBufferPeriod)(nil)

func WithBufferPeriod(period time.Duration) Option {
	wbp := withBufferPeriod(period)
	return &wbp
}

type withBufferCount int

func (wbc *withBufferCount) withCustomAppender(ca *customAppender) {
	ca.metricsBufferCount = int(*wbc)
}

var _ Option = (*withBufferCount)(nil)

func WithBufferCount(count int) Option {
	wbc := withBufferCount(count)
	return &wbc
}
