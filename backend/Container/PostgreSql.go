package Container

import (
	"fmt"

	"github.com/ewriq/pouch"
)

func CreatePostgresContainer(name, img, port, password string) (string, error) {
	fmt.Println("📦 PostgreSQL container oluşturuluyor...")

	opt := pouch.CreateOptions{
		Name:  name,
		Image: img,
		Port:  port,
		EnvVars: map[string]string{
			"POSTGRES_PASSWORD": password,
		},
		Labels: map[string]string{
			"type": "postgres",
		},
		UserUIDGID: "1001:1001",
	}

	id, err := pouch.Create(opt)
	if err != nil {
		return "", fmt.Errorf("container oluşturulamadı: %v", err)
	}

	return id, nil
}
