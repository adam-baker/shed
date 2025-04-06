package workloads

import "context"

type Workload interface {
	Execute(ctx context.Context, payload string) error
}
