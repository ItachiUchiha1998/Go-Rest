package store

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type Product struct {
	ID     int    `bson:"_id"`
	Title  string `json:"title"`
	Image  string `json:"image"`
	Price  uint64 `json:"price"`
	Rating uint8  `json:"rating"`
}

type Products []Product
