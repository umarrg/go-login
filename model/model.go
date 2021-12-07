package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Res struct {
	Token string `json:"token"`
	Data  User   `json:"user"`
}

type Exception struct {
	Message string `json:"message"`
}

type Airline struct {
	Id           int `json:"id"`
	Name         int `json:"name"`
	Country      int `json:"country"`
	Logo         int `json:"logo"`
	Slogon       int `json:"slogon"`
	Head_quaters int `json:"head_quaters"`
	Website      int `json:"website"`
}
