 package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func StartMySQLContainer(containerName string, dbPort string, rootPassword string) (string, error) {
	image := "mysql:8.0"

	fmt.Println("📦 MySQL image çekiliyor...")
	cmdPull := exec.Command("docker", "pull", image)
	if out, err := cmdPull.CombinedOutput(); err != nil {
		return "", fmt.Errorf("image çekilemedi: %v\n%s", err, out)
	}

	fmt.Println("🚀 MySQL container başlatılıyor...")

	cmd := exec.Command("docker", "run", "-d",
		"-e", "MYSQL_ROOT_PASSWORD="+rootPassword,
		"-p", dbPort+":3306",
		"--name", containerName,
		image,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("container başlatılamadı: %v\n%s", err, string(output))
	}

	containerID := string(output)
	return containerID, nil
}

func main() {
	containerName := fmt.Sprintf("mysql-server-%d", time.Now().Unix())
	dbPort := "3307"      // localhost:3307 → container:3306
	rootPassword := "123" // root şifresi

	containerID, err := StartMySQLContainer(containerName, dbPort, rootPassword)
	if err != nil {
		fmt.Println("HATA:", err)
		os.Exit(1)
	}

	fmt.Println("✅ MySQL container başarıyla başlatıldı.")
	fmt.Println("🆔 ID:", containerID)
	fmt.Println("🔗 Bağlantı için:")
	fmt.Printf("    Host: localhost\n")
	fmt.Printf("    Port: %s\n", dbPort)
	fmt.Printf("    Kullanıcı: root\n")
	fmt.Printf("    Şifre: %s\n", rootPassword)
}
