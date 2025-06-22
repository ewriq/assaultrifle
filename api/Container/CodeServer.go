package Container

import (
	"fmt"

	"github.com/ewriq/pouch"
)

func CreateNextcloudContainer(name, img, port string) (string, error) {
	fmt.Println("☁️ Nextcloud container oluşturuluyor...")

	opt := pouch.CreateOptions{
		Name:  name,
		Image: img, 
		Port:  port,
		EnvVars: map[string]string{
			"PUID": "1000",
			"PGID": "1000",
			"TZ":   "Europe/Istanbul",
		},
		Labels: map[string]string{
			"type": "nextcloud",
		},
	}

	id, err := pouch.Create(opt)
	if err != nil {
		return "", fmt.Errorf("nextcloud container oluşturulamadı: %v", err)
	}

	return id, nil
}
