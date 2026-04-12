// Copyright 2026 Alibaba Group Holding Ltd.
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

package telemetry

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

func systemProcessCount() int64 {
	pids, err := process.Pids()
	if err != nil {
		return 0
	}
	return int64(len(pids))
}

func systemCPUUsagePercent() float64 {
	usage, err := cpu.Percent(0, false)
	if err != nil || len(usage) == 0 {
		return 0
	}
	return usage[0]
}

func systemMemoryUsageBytes() int64 {
	stats, err := mem.VirtualMemory()
	if err != nil {
		return 0
	}
	return int64(stats.Used)
}
