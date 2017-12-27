package db

type User struct {
	Id *int
	Username *string
}

func GetUserById(id string) (*User, error) {
	row := db.QueryRow("select id, username from auth_user where id=$1", id)
	user := User{}
	err := row.Scan(&user.Id, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}