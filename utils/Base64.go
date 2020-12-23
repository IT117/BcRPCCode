package utils

import "encoding/base64"

func Base64() string {
	authStr := RPCUSER + ":" + RPCPASSSWORD
	authStrByte := []byte(authStr)
	authBase := base64.StdEncoding.EncodeToString(authStrByte)

	return authBase

}
