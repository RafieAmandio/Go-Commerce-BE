package dto

type CustomerRegisterRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type CustomerLoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type CustomerLoginResponse struct {
    Token string `json:"token"`
}
