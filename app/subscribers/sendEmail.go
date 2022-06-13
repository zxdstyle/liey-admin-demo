package subscribers

import (
	"context"
	"fmt"
)

type SendEmail struct {
}

func (s SendEmail) Handle(ctx context.Context, payload interface{}) error {
	fmt.Printf("SendEmail: %+v\n", payload)
	return nil
}

func (SendEmail) Async() bool {
	return true
}
