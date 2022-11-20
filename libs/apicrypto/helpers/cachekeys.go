package helpers

import "fmt"

func GetPublicKeyId(clientId string) string {
	return fmt.Sprintf("%s_public", clientId)
}

func GetPrivateKeyId(clientId string) string {
	return fmt.Sprintf("%s_private", clientId)
}
