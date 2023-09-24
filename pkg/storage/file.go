package storage

import (
	"fmt"
	"os"
)

func CreateFile() (*os.File, error) {
	file, err := os.CreateTemp("", "fileShapes")
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return file, nil
}

func OpenFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0644) // Abre archivo usando permisos READ & WRITE
	if err != nil {
		return nil, fmt.Errorf("OpenFile. No fue posible abrir el archivo (%s). %s", path, err)
	}
	return file, nil
}

func WriteInFile(file *os.File, data string) error {
	n, err := file.WriteString(data + "\r\n")
	if err != nil {
		return fmt.Errorf("WriteInFile. Error escribiendo archivo (%d). %s", n, err)
	}

	err = file.Sync()
	if err != nil {
		return fmt.Errorf("WriteInFile. Error guardando cambios (%d). %s", n, err)
	}
	return nil
}
