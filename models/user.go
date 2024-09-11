package models

type User struct {
    ID   int    `json:"user_id"`
    Name string `json:"user_name"`
}

func GetUserByID(userID int) (*User, error) {
    user := User{}
    query := `SELECT id, name FROM users WHERE id = ?`
    err := DB.QueryRow(query, userID).Scan(&user.ID, &user.Name)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
