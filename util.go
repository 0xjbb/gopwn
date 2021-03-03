package gopwn

import "os"

func FileExists(FilePath string) bool{
	data, err := os.Stat(FilePath)

	if  os.IsNotExist(err) {
		return false
	}
	// something exists, make sure its not a directory.
	return !data.IsDir()
}
