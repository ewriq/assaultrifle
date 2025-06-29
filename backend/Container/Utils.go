package Container

import (
	"assaultrifle/Log"
	"fmt"
	"os/exec"
	"strings"

	"github.com/ewriq/pouch"
)

func StartContainer(id string) error {
	Log.Set("🚀 Container başlatılıyor..."+ id)
	return pouch.Start(id)
}

func PullImage(img string) error {
	Log.Set("📦 Image çekiliyor..."+ img)
	if err := pouch.Pull(img); err != nil {
		return fmt.Errorf("image çekilemedi: %v", err)
	}
	return nil
}

func StopContainer(id string) error {
	Log.Set("🚀 Container durduruluyor"+ id);
	return pouch.Stop(id)
}

func GetContainerStatus(id string) (map[string]string, error) {
	Log.Set("🚀 Container status çekildi"+ id);
	return pouch.ContainerStats(id)
}

func DeleteContainer(id string) error {
	Log.Set("🚀 Container silindi"+ id);
	_, err := pouch.Remove(id, false)
	return err
}

func RestartContainer(id string) error {
	
	_, err := pouch.Restart(id)
	if err != nil {
		return fmt.Errorf("container yeniden başlatılamadı: %v", err)
	}

	return nil
}


func GetContainerLogs(id string) (string, error) {
	Log.Set("📄 Container logları alınıyor..."+ id);
    logs, err := pouch.Logs(id)
    if err != nil {
        return "", fmt.Errorf("log alınamadı: %v", err)
    }


    return logs, nil
}
func ListContainerFiles(containerID, containerPath string) ([]string, error) {
	Log.Set("📂 Container içindeki dosyalar listeleniyor:"+ containerID + containerPath);
	
	files, err := pouch.ListFiles(containerID, containerPath)
	if err != nil {
		return nil, fmt.Errorf("dosyalar listelenemedi: %v", err)
	}
	return files, nil
}


func DeleteContainerFile(containerID, filePath string) error {
	Log.Set("🗑️ Container  içindeki dosya siliniyor: "+ containerID+  filePath)
	return pouch.DeleteFile(containerID, filePath)
}


func CopyFileToContainer(containerID, localFilePath, containerTargetPath string) error {
	Log.Set("📤 Container içine dosya kopyalanıyor: "+ containerID + localFilePath +  containerTargetPath)
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

	Log.Set("📦 Tüm container'lar durduruldu.")
	return nil
}


func DeleteAllContainer()  error {
	Log.Set("📦 Tüm container'lar silindi.")
	err := DeleteContainer("$(docker ps -aq)")
	if err != nil {
		return err
	}
	return err
}