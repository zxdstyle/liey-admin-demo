package subscribers

import (
	"context"
	"fmt"
)

type SendNotification struct {
}

func (s SendNotification) Handle(ctx context.Context, payload interface{}) error {
	fmt.Printf("SendNotification: %+v\n", payload)
	return nil
}

func (s SendNotification) Async() bool {
	return true
}
