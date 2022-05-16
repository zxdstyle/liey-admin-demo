package auth

type (
	LoginByPwd struct {
		Email    *string `json:"email" v:"required,email"`
		Password *string `json:"password" v:"required"`
	}

	LoginResp struct {
		Email  *string `json:"email"`
		Token  *string `json:"token"`
		Avatar *string `json:"avatar"`
	}
)
