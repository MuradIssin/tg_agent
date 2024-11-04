package files

import (
	"fmt"
	"os"
)

func ReadFile(nameFile string) ([]byte, error) {
	data, err := os.ReadFile(nameFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// fmt.Println(string(data))
	return data, nil
}

func WriteFile(content []byte, nameFile string) {
	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("запись успешна")
}
