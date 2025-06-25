package Container

import (
	"fmt"
	"os/exec"
	"strings"

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

func ListContainerFiles(containerID, containerPath string) ([]string, error) {
	fmt.Printf("📂 Container (%s) içindeki dosyalar listeleniyor: %s\n", containerID, containerPath)
	
	files, err := pouch.ListFiles(containerID, containerPath)
	if err != nil {
		return nil, fmt.Errorf("dosyalar listelenemedi: %v", err)
	}
	return files, nil
}


func DeleteContainerFile(containerID, filePath string) error {
	fmt.Printf("🗑️ Container (%s) içindeki dosya siliniyor: %s\n", containerID, filePath)
	return pouch.DeleteFile(containerID, filePath)
}


func CopyFileToContainer(containerID, localFilePath, containerTargetPath string) error {
	fmt.Printf("📤 Container (%s) içine dosya kopyalanıyor: %s -> %s\n", containerID, localFilePath, containerTargetPath)
	return pouch.CopyToContainer(containerID, localFilePath, containerTargetPath)
}

func ContainerStopAll() error {
	cmd := exec.Command("docker", "ps", "-q")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("listeleme hatası: %w", err)
	}

	ids := strings.Fields(string(output))
	if len(ids) == 0 {
		fmt.Println("📦 Hiç çalışan container yok.")
		return nil
	}

	args := append([]string{"stop"}, ids...)
	stopCmd := exec.Command("docker", args...)
	stopOutput, err := stopCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("durdurma hatası: %v\n%s", err, stopOutput)
	}

	fmt.Println("📦 Tüm container'lar durduruldu.")
	return nil
}
