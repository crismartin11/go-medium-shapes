package storage

import (
	"fmt"
	"os"
)

// const PATH_TEMP = "/tmp"
//const PATH_TEMP = ""

/*func CreateFile(name string) (*os.File, error) {
	path, err := getPath(name)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	var _, errFile = os.Stat(path)
	if os.IsNotExist(errFile) {
		var file, err = os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("CreateFile. Error creando archivo %s. %s", name, err)
		}
		return file, nil
	} else {
		return OpenFile(name)
	}
}*/

func CreateFile() (*os.File, error) {
	file, err := os.CreateTemp("", "example") // TODO: cambiar nombre
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return file, nil
}

func OpenFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0644) // Abre archivo usando permisos READ & WRITE // TODO: crear otro de solo lectura
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

/*func getPath(name string) (string, error) {
	_, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("getPath. %s", err)
	}

	path := name
	if PATH_TEMP == "/tmp" {
		path = PATH_TEMP + "/" + name
	}

	return path, nil
}*/
