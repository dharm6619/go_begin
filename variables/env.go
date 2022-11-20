package variables

import (
	"os"
)

func GetLoggedInUser() string {

	return os.Getenv("USER")
}
