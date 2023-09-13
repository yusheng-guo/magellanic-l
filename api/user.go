package api

/* ---------------------------- request ---------------------------- */

type LoginReq struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterReq struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

/* ---------------------------- response ---------------------------- */

type LoginResp struct {
	Token string `json:"token" form:"token"`
}
type RegisterResp struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
