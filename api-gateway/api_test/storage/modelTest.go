package storage

type CategoryIcon struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type Category struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	UserId    string `json:"user_id"`
	IconId    int64  `json:"icon_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type Product struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Picture     string `json:"picture"`
	CategoryId  int64  `json:"category_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

type Orderproduct struct {
	Id        int64  `json:"id"`
	UserId    string `json:"user_id"`
	ProductId int64  `json:"product_id"`
	CreatedAt string `json:"created_at"`
}

type Comment struct {
	Id        int64  `json:"id"`
	Content   string `json:"content"`
	UserId    string `json:"user_id"`
	ProductId int64  `json:"product_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}


type User struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	PhoneNumber  string `json:"phone_number"`
	Bio          string `json:"bio"`
	BirthDay     string `json:"birth_day"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}


type UserLoginRequest struct {
	UserNameOrEmail string `json:"user_name_or_email"`
	Password        string `json:"password"`
}

type RegisterModel struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseMessage struct {
	Content string `json:"content"`
}

type Users struct {
	Users []*User `json:"users"`
}

type Delete struct {
	Result string
}
