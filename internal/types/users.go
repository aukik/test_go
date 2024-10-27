package types

import "time"

type CreateUserRequest struct {
    Username string `json:"username"`
}

type CreateUserResponse struct {
    ID        string    `json:"id"`
    Username  string    `json:"username"`
    CreatedAt time.Time `json:"created_at"`
}
