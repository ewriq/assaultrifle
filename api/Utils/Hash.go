package Utils

import (
	"encoding/base64"
	"fmt"
)

func Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
func Decode(data string)  (string, error){
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println("Decode hatası:", err)
		return "", err
	}
	return string(decoded), nil
}