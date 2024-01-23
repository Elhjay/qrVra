package standardError

import (
	"log"
)

func StddErr(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}
