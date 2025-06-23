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
    defer func() {
        if err := c.Close(); err != nil {
            log.Println("WebSocket close error:", err)
        }
    }()

    containerID := c.Params("id")
    userToken := c.Query("token")

    ok, err := Database.ValidateContainerAccess(containerID, userToken)
    if err != nil || !ok {
        if writeErr := c.WriteMessage(websocket.TextMessage, []byte("❌ Yetkisiz erişim")); writeErr != nil {
            log.Println("WebSocket write error:", writeErr)
        }
        return
    }

    dirStateLock.Lock()
    if _, ok := dirState[containerID]; !ok {
        dirState[containerID] = "/"
    }
    currentDir := dirState[containerID]
    dirStateLock.Unlock()

    sendPrompt := func() {
        if writeErr := c.WriteMessage(websocket.TextMessage, []byte(currentDir+" ➜ ")); writeErr != nil {
            log.Println("WebSocket write error:", writeErr)
        }
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
                if writeErr := c.WriteMessage(websocket.TextMessage, []byte("⚠️ Dizin bulunamadı: "+newPath+"\n")); writeErr != nil {
                    log.Println("WebSocket write error:", writeErr)
                }
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
            if writeErr := c.WriteMessage(websocket.TextMessage, []byte(currentDir+"\n")); writeErr != nil {
                log.Println("WebSocket write error:", writeErr)
            }
            sendPrompt()
            continue
        }

        if command == "clear" {
            if writeErr := c.WriteMessage(websocket.TextMessage, []byte("__CLEAR__")); writeErr != nil {
                log.Println("WebSocket write error:", writeErr)
            }
            sendPrompt()
            continue
        }

        if command == "exit" {
            if writeErr := c.WriteMessage(websocket.TextMessage, []byte("👋 Görüşmek üzere.\n")); writeErr != nil {
                log.Println("WebSocket write error:", writeErr)
            }
            break
        }

        fullCommand := fmt.Sprintf("cd %s && %s", currentDir, command)
        out, err := pouch.Exec(containerID, []string{"sh", "-c", fullCommand})
        if err != nil {
            if writeErr := c.WriteMessage(websocket.TextMessage, []byte("⚠️ "+err.Error()+"\n")); writeErr != nil {
                log.Println("WebSocket write error:", writeErr)
            }
        } else {
            if writeErr := c.WriteMessage(websocket.TextMessage, []byte(out)); writeErr != nil {
                log.Println("WebSocket write error:", writeErr)
            }
        }

        sendPrompt()
    }
}