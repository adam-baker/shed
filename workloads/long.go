package workloads

import (
	"context"
	"fmt"
	"time"
)

type Long struct{}

func (l Long) Execute(ctx context.Context, payload string) error {
	fmt.Println("Executing LONG workload with payload:", payload)
	time.Sleep(10 * time.Second)
	return nil
}
