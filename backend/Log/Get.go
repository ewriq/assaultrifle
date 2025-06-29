package Log

import (
	"github.com/goccy/go-json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetLog(date string) ([]LogEntry, error) {
	var logs []LogEntry

	files, err := os.ReadDir(logDir)
	if err != nil {
		return nil, fmt.Errorf("log klasörü okunamadı: %v", err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		filename := f.Name()
		if !strings.HasSuffix(filename, ".txt") {
			continue
		}

		nameWithoutExt := strings.TrimSuffix(filename, ".txt")

		if date != "*" && nameWithoutExt != date {
			continue
		}

		content, err := os.ReadFile(filepath.Join(logDir, filename))
		if err != nil {
			return nil, fmt.Errorf("log dosyası okunamadı: %v", err)
		}

		logs = append(logs, LogEntry{
			Time: nameWithoutExt,
			Data: string(content),
		})
	}

	return logs, nil
}

func Log(date string) (string, error) {
	entries, err := GetLog(date)
	if err != nil {
		return "", err
	}
	jsonBytes, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return "", fmt.Errorf("JSON'a dönüştürülemedi: %v", err)
	}
	return string(jsonBytes), nil
}
