package Container

import (
	"fmt"

	"github.com/ewriq/pouch"
)

func StartContainer(id string) error {
	fmt.Println("🚀 Container başlatılıyor...")
	return pouch.Start(id)
}

func PullImage(img string) error {
	fmt.Println("📦 Image çekiliyor...")
	if err := pouch.Pull(img); err != nil {
		return fmt.Errorf("image çekilemedi: %v", err)
	}
	return nil
}

func StopContainer(id string) error {
	return pouch.Stop(id)
}

func GetContainerStatus(id string) (map[string]string, error) {
	return pouch.ContainerStats(id)
}

func DeleteContainer(id string) error {
	_, err := pouch.Remove(id, false)
	return err
}

func RestartContainer(id string) error {
	fmt.Println("🔁 Container yeniden başlatılıyor...")

	_, err := pouch.Restart(id)
	if err != nil {
		return fmt.Errorf("container yeniden başlatılamadı: %v", err)
	}

	return nil
}


func GetContainerLogs(id string) (string, error) {
	fmt.Println("📄 Container logları alınıyor...")

	logs, err := pouch.Logs(id)
	if err != nil {
		return "", fmt.Errorf("log alınamadı: %v", err)
	}

	return logs, nil
}