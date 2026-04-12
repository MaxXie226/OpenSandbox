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
	"context"
	"sync"
	"testing"

	"go.opentelemetry.io/otel/attribute"

	inttelemetry "github.com/alibaba/opensandbox/internal/telemetry"
)

func TestNormalizeRoute(t *testing.T) {
	t.Parallel()

	if got := normalizeRoute(""); got != "unknown" {
		t.Fatalf("normalizeRoute(\"\") = %q, want %q", got, "unknown")
	}

	if got := normalizeRoute("/code/contexts/:contextId"); got != "/code/contexts/:contextId" {
		t.Fatalf("normalizeRoute(route) = %q, want %q", got, "/code/contexts/:contextId")
	}
}

func TestRecordHTTPRequestWithoutInit(t *testing.T) {
	t.Parallel()

	RecordHTTPRequest(context.Background(), "GET", "/ping", 200, 0.01)
}

func TestSystemMetricsReaders(t *testing.T) {
	t.Parallel()

	if got := systemProcessCount(); got < 0 {
		t.Fatalf("systemProcessCount() = %d, want >= 0", got)
	}
	if got := systemCPUUsagePercent(); got < 0 {
		t.Fatalf("systemCPUUsagePercent() = %f, want >= 0", got)
	}
	if got := systemMemoryUsageBytes(); got < 0 {
		t.Fatalf("systemMemoryUsageBytes() = %d, want >= 0", got)
	}
}

func TestExecdSharedAttrs(t *testing.T) {
	t.Setenv(envSandboxID, "sb-123")
	t.Setenv(envMetricsExtraAttr, "tenant=t1,env=dev")

	orig := execdSharedAttrs
	execdSharedAttrs = sync.OnceValue(func() []attribute.KeyValue {
		return inttelemetry.SharedAttrsFromEnv(inttelemetry.SharedAttrsEnvConfig{
			SandboxIDEnv:  envSandboxID,
			ExtraAttrsEnv: envMetricsExtraAttr,
			SandboxAttr:   "sandbox_id",
		})
	})
	t.Cleanup(func() { execdSharedAttrs = orig })

	attrs := execdSharedAttrs()
	if len(attrs) != 3 {
		t.Fatalf("attrs len = %d, want 3", len(attrs))
	}
}
