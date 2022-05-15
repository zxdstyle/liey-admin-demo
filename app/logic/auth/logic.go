package auth

import "context"

type Logic struct {
}

func NewLogic() *Logic {
	return &Logic{}
}

func (Logic) Login(ctx context.Context) {

}
