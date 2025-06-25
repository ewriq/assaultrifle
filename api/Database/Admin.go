package Database

import (
	"assaultrifle/Form"
)

func ValidateAdminAccess(token string) (string, error) {
	var user Form.User
	result := DB.Where("token = ?", token).First(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.Perm, nil
}