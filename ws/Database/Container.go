
func ValidateContainerAccess(containerToken, userToken string) (bool, error) {
	var container Form.Container
	if err := DB.Where("token = ?", containerToken).First(&container).Error; err != nil {
		return false, errors.New("container bulunamadı")
	}

	if container.User != userToken {
		return false, errors.New("yetkisiz erişim")
	}

	return true, nil
}
