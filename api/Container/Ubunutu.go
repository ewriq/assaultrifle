package Container

import (
	"fmt"

	"github.com/ewriq/pouch"
)

func CreateUbuntuContainer(name, img, port string) (string, error) {
	fmt.Println("📦 Ubuntu container oluşturuluyor...")

	opt := pouch.CreateOptions{
		Name:  name,
		Image: img,
		Port:  port,
		EnvVars: map[string]string{
			"LANG": "C.UTF-8",
		},
		Labels: map[string]string{
			"type": "ubuntu",
		},
		EntryPoint: "/bin/bash",
		UserUIDGID: "1001:1001",
	}

	id, err := pouch.Create(opt)
	if err != nil {
		return "", fmt.Errorf("Ubuntu container oluşturulamadı: %v", err)
	}

	return id, nil
}
