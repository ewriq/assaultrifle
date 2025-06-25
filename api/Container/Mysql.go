package Container

import (
	"fmt"

	"github.com/ewriq/pouch"
)

func CreateMySQLContainer(name, img, port, password string) (string, error) {
	fmt.Println("📦 MySQL container oluşturuluyor...")

	opt := pouch.CreateOptions{
		Name:  name,
		Image: img,
		Port:  port,
		EnvVars: map[string]string{
			"MYSQL_ROOT_PASSWORD": password,
		},
		Labels: map[string]string{
			"type": "mysql",
		},
		UserUIDGID: "1001:1001",
	}

	id, err := pouch.Create(opt)
	if err != nil {
		return "", fmt.Errorf("container oluşturulamadı: %v", err)
	}

	return id, nil
}
