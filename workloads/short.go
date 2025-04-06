package workloads

import (
	"context"
	"fmt"
	"time"
)

type Short struct{}

func (s Short) Execute(ctx context.Context, payload string) error {
	fmt.Println("Executing SHORT workload with payload:", payload)
	time.Sleep(2 * time.Second)
	return nil
}
