package Container

import (
	"fmt"
	"os/exec"
	"strings"
)

func PullPostgresImage() error {
	cmd := exec.Command("docker", "pull", "postgres:16")
	fmt.Println("📦 PostgreSQL image çekiliyor...")
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("image çekilemedi: %v\n%s", err, out)
	}
	return nil
}

func CreatePostgresContainer(port, password string) (string, error) {
	cmd := exec.Command("docker", "create",
		"-e", "POSTGRES_PASSWORD="+password,
		"-p", port+":5432",
		"postgres:16",
	)

	fmt.Println("📦 PostgreSQL container oluşturuluyor...")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("container oluşturulamadı: %v\n%s", err, output)
	}
	return strings.TrimSpace(string(output)), nil
}
