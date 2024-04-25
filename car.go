package autokatolog

type Car struct {
	ID      int    `json:"id" db:"id"`
	RegNums string `json:"reg_num" db:"reg_num"`
	Mark    string `json:"mark" db:"mark"`
	Model   string `json:"model" db:"model"`
	Year    int    `json:"year" db:"year"`
	Owner   People `json:"owner"  db:"owner"`
}

type People struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}
