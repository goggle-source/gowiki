package createenv

import "os"

func IsDirExists(pathDirectory string) (bool, error) {
	result, err := os.Stat(pathDirectory)
	if err != nil {
		return false, err
	}
	return result.IsDir(), err
}

func CreateDirectory(path string) error {

	err := os.Mkdir(path, 0755)

	return err
}
