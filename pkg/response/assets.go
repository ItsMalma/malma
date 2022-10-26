package response

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Assets(filePath ...string) ([]byte, string, error) {
	file, err := os.OpenFile("assets/"+filepath.Join(filePath...), os.O_RDONLY, 0644)
	if err != nil {
		return nil, "", err
	}
	extension := ""
	splittedFileName := strings.Split(file.Name(), ".")
	switch splittedFileName[len(splittedFileName)-1] {
	case "css":
		extension = "text/css"
	case "js":
		extension = "text/javascript"
	}
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, "", err
	}
	return fileContent, extension, nil
}
