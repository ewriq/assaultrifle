package Log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const logDir = "./Log"

func Set(text string) error {
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return fmt.Errorf("log klasörü oluşturulamadı: %v", err)
	}

	today := time.Now().Format("2006-01-02")
	filePath := filepath.Join(logDir, today+".txt")

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("log dosyası açılamadı: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(text + "\n")
	return err
}

