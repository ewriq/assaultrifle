package Container

import (
	"assaultrifle/Log"
	"fmt"

	"github.com/ewriq/pouch"
)

func CreateNodeContainer(name, img, port string) (string, error) {
	Log.Set("📦 Node.js container oluşturuluyor...")
	
	opt := pouch.CreateOptions{
		Name:        name,
		Image:       img,
		Port:        port,
		MemoryLimit: "512m",   
		CPULimit:    0.3,      
		EnvVars: map[string]string{
			"NODE_ENV": "production",
		},
		Labels: map[string]string{
			"type": "nodejs",
		},
		UserUIDGID: "1001:1001",
	}

	id, err := pouch.Create(opt)
	if err != nil {
		return "", fmt.Errorf("container oluşturulamadı: %v", err)
	}

	return id, nil
}
