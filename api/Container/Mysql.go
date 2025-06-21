package Container

import (
	"fmt"
	"os/exec"
	"strings"
)


func PullMySQLImage() error {
	cmd := exec.Command("docker", "pull", "mysql:8.0")
	fmt.Println("📦 Image çekiliyor...")
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("image çekilemedi: %v\n%s", err, out)
	}
	return nil
}

func CreateMySQLContainer(port, password string) (string, error) {
	cmd := exec.Command("docker", "create",
		"-e", "MYSQL_ROOT_PASSWORD="+password,
		"-p", port+":3306",
		"mysql:8.0",
	)

	fmt.Println("📦 Container oluşturuluyor...")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("container oluşturulamadı: %v\n%s", err, output)
	}
	return strings.TrimSpace(string(output)), nil
}

