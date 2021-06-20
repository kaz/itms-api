package itms

import (
	"os"
)

var (
	credential = map[string]string{
		"inputUserId":   os.Getenv("ITMS_USER_ID"),
		"inputPassword": os.Getenv("ITMS_PASSWORD"),
	}
)
