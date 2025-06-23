package Handler

import (
	"fmt"
	"log"

	"strings"
	"sync"

	"assaultrifle/Database"

	"github.com/ewriq/pouch"
	"github.com/gofiber/websocket/v2"
)

var dirState = make(map[string]string)   
var dirStateLock = sync.Mutex{}

func WebSocket(c *websocket.Conn) {
	defer c.Close()

	containerID := c.Params("id")
	userToken := c.Query("token")

	ok, err := Database.ValidateContainerAccess(containerID, userToken)
	if err != nil || !ok {
		c.WriteMessage(websocket.TextMessage, []byte("❌ Yetkisiz erişim"))
		return
	}

	
	dirStateLock.Lock()
	if _, ok := dirState[containerID]; !ok {
		dirState[containerID] = "/"
	}
	currentDir := dirState[containerID]
	dirStateLock.Unlock()

	sendPrompt := func() {
		c.WriteMessage(websocket.TextMessage, []byte(currentDir + " ➜ "))
	}

	sendPrompt()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		command := strings.TrimSpace(string(msg))
		if command == "" {
			sendPrompt()
			continue
		}

	
		if strings.HasPrefix(command, "cd") {
			parts := strings.Fields(command)
			newPath := "/"

			if len(parts) > 1 {
				arg := parts[1]

				if strings.HasPrefix(arg, "/") {
					newPath = arg
				} else if currentDir == "/" {
					newPath = "/" + arg
				} else {
					newPath = currentDir + "/" + arg
				}
			}


			checkCmd := fmt.Sprintf("[ -d \"%s\" ] && echo OK || echo FAIL", newPath)
			out, _ := pouch.Exec(containerID, []string{"sh", "-c", checkCmd})

			if strings.TrimSpace(out) == "FAIL" {
				c.WriteMessage(websocket.TextMessage, []byte("⚠️ Dizin bulunamadı: " + newPath + "\n"))
				sendPrompt()
				continue
			}

			dirStateLock.Lock()
			dirState[containerID] = newPath
			currentDir = newPath
			dirStateLock.Unlock()

			sendPrompt()
			continue
		}

		  
		if command == "pwd" {
			c.WriteMessage(websocket.TextMessage, []byte(currentDir + "\n"))
			sendPrompt()
			continue
		}


		if command == "clear" {
			c.WriteMessage(websocket.TextMessage, []byte("__CLEAR__"))
			sendPrompt()
			continue
		}

	
		if command == "exit" {
			c.WriteMessage(websocket.TextMessage, []byte("👋 Görüşmek üzere.\n"))
			break
		}


		fullCommand := fmt.Sprintf("cd %s && %s", currentDir, command)
		out, err := pouch.Exec(containerID, []string{"sh", "-c", fullCommand})
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte("⚠️ " + err.Error() + "\n"))
		} else {
			c.WriteMessage(websocket.TextMessage, []byte(out))
		}

		sendPrompt()
	}
}
