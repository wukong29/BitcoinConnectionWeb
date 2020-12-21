package utils

import "encoding/base64"

/**
 * base64编码
 */
func Base64EncodeString(data string) string{
	return base64.StdEncoding.EncodeToString([]byte(data))
}