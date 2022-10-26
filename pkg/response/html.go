package response

import (
	"os"
	"path"
)

func View(fileName string) ([]byte, error) {
	return os.ReadFile(path.Join("views", fileName))
}
