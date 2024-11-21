package entity

type Users struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	FullName string `json:"full_name"`
	TokoID   int    `json:"toko_id"`
	Role     string `json:"role"`
}
