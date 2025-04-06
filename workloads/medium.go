package workloads

import (
	"context"
	"fmt"
	"time"
)

type Medium struct{}

func (m Medium) Execute(ctx context.Context, payload string) error {
	fmt.Println("Executing MEDIUM workload with payload:", payload)
	time.Sleep(5 * time.Second)
	return nil
}
