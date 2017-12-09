package main

// User - user model
type User struct {
	// Field names must begin with captials (be exported), or else packages
	// such as `json`  or `c.JSON` cannot read from them
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func isEmailExists(email string) bool {
	for _, user := range userList {
		if user.Email == email {
			return false
		}
	}
	return true
}

func isUserValid(user User) bool {
	return user.Email != "" && user.FirstName != "" && user.LastName != ""
}
