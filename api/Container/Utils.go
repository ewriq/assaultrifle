package Container

import (
	"fmt"
	"os/exec"
	"strings"
)

func StartContainer(ID string) error {
	cmd := exec.Command("docker", "start", ID)
	fmt.Println("🚀 Container başlatılıyor...")
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("container başlatılamadı: %v\n%s", err, out)
	}
	return nil
}

func StopContainer(ID string) error {
	cmd := exec.Command("docker", "stop", ID)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("container durdurulamadı: %v\n%s", err, out)
	}
	return nil
}

func GetContainerStatus(ID string) (string, error) {
	cmd := exec.Command("docker", "inspect", "-f", "{{.State.Status}}", ID)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("durum alınamadı: %v\n%s", err, output)
	}
	return strings.TrimSpace(string(output)), nil
}

func DeleteContainer(ID string) error {
	cmd := exec.Command("docker", "rm", ID)
	fmt.Println("🗑️ Container siliniyor...")
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("container silinemedi: %v\n%s", err, out)
	}
	return nil
}
