package Container

import (
	"assaultrifle/Log"
	"fmt"

	"github.com/ewriq/pouch"
)

func CreatePostgresContainer(name, img, port, password string) (string, error) {
	Log.Set("📦 PostgreSQL container oluşturuluyor...")
	opt := pouch.CreateOptions{
		Name:        name,
		Image:       img,
		Port:        port,
		MemoryLimit: "512m",  
		CPULimit:    0.3,
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
