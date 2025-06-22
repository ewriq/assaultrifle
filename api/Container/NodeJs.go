package Container

import (
	"fmt"

	"github.com/ewriq/pouch"
)

func CreateNodeContainer(name, img, port string) (string, error) {
	fmt.Println("📦 Node.js container oluşturuluyor...")

	opt := pouch.CreateOptions{
		Name:  name,
		Image: img,
		Port:  port,
		EnvVars: map[string]string{
			"NODE_ENV": "production",
		},
		Labels: map[string]string{
			"type": "nodejs",
		},
	}

	id, err := pouch.Create(opt)
	if err != nil {
		return "", fmt.Errorf("container oluşturulamadı: %v", err)
	}

	return id, nil
}
