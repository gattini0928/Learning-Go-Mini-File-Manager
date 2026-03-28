package main

import ("fmt"
	"os"
	"strings"
	"path/filepath"
	"strconv"
	"bufio"
	"io"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	loop:
	for {
	fmt.Println("Bem-vindo ao organizador de arquivos")
	codeFolder, txtFolder, imagesFolder := createDirectories()
	fmt.Println("Pastas code, txt e images criadas!")
	fmt.Println("1 - Mover arquivo(.go)")
	fmt.Println("2 - Mover arquivo(.txt)")
	fmt.Println("3 - Mover arquivo(.png, .jpeg, .jpg, .webp)")
	fmt.Println("4 - Listar arquivos de uma pasta: (./code, ./txt, ./images)")
	fmt.Println("5 - Deletar arquivo de uma pasta: (./code, ./txt, ./images)")
	fmt.Println("6 - Sair do sistema")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("❌ Opção Inválida. Tente Novamente")
		continue
	}

	switch choice{
	case 1:
		fmt.Println("Insira o arquivo .go: ")
		file, _ := reader.ReadString('\n')
		file = strings.TrimSpace(file)

		msg, err := insertFileToFolder(file, codeFolder, txtFolder, imagesFolder)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
	case 2 :
		fmt.Println("Insira o arquivo .txt: ")
		file, _ := reader.ReadString('\n')
		file = strings.TrimSpace(file)

		msg, err := insertFileToFolder(file, codeFolder, txtFolder, imagesFolder)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
	case 3:
		fmt.Println("Insira o arquivo: ")
		file, _ := reader.ReadString('\n')
		file = strings.TrimSpace(file)

		msg, err := insertFileToFolder(file, codeFolder, txtFolder, imagesFolder)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
	case 4:
		fmt.Println("Insira o nome da pasta que deseja listar os arquivos: ")
		folder, _ := reader.ReadString('\n')
		folder = strings.TrimSpace(folder)

		listAllFiles(folder)
	case 5:
		fmt.Println("Insira o nome da pasta que deseja deletar o arquivo: ")
		folder, _ := reader.ReadString('\n')
		folder = strings.TrimSpace(folder)

		fmt.Println("Qual arquivo deseja deletar: ")
		file, _ := reader.ReadString('\n')
		file = strings.TrimSpace(file)

		listAllFiles(folder)
		deleteFileFromFolder(folder, file)
	case 6:
		fmt.Println("Até logo ✌️")
		break loop
	}
	}
}

func createDirectories() (string, string, string){
	folders := []string{"code", "txt", "images"}
	for _, folderName := range folders {
		err := os.MkdirAll(folderName, os.ModePerm)
		if err != nil {
			fmt.Printf("Erro ao criar diretório: %v, %v", folderName, err) 
		}
	}

	return "./code", "./txt", "./images"
}

func moveFileToFolder(fileName string, folder string) error {
	newPath := filepath.Join(folder, fileName)
	return os.Rename(fileName, newPath)
}

func insertFileToFolder(file string, codeFolder string, txtFolder string, imagesFolder string) (string, error){
	var msg string
	var folder string
		
	ext := filepath.Ext(file)

	if ext == ".go" {
		folder = codeFolder
	} else if ext == ".txt" {
		folder = txtFolder
	} else if ext == ".png" || ext == ".webp" || ext == ".jpeg" || ext == ".jpg" {
		folder = imagesFolder
	} else {
		return "Tipo de arquivo não suportado", nil
	}

	err := moveFileToFolder(file, folder)
	if err != nil {
		return fmt.Sprintf("Erro ao mover arquivo: %v", err), err
	}

	msg = fmt.Sprintf("Arquivo %s, movido para pasta %s", file, folder)
	return msg, nil
}

func listAllFiles(folder string) {
	empty, err := IsEmpty(folder)
	if err != nil {
		fmt.Printf("Erro ao checar diretório vazio: %v", err)
	}
	if empty {
		fmt.Println("A Pasta está vazia, mova arquivos para lista-los")
	}

	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Printf("Arquivo: %v \n, %v", file.Name(), file.IsDir())
	}
}

func deleteFileFromFolder(folder string, fileName string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		fmt.Printf("Pasta %v não existe", folder)
		return
	}

	empty, err := IsEmpty(folder)
	if err != nil {
		fmt.Printf("Erro ao checar diretório vazio: %v", err)
	}
	if empty {
		fmt.Println("A Pasta está vazia, mova arquivos para lista-los")
	}
	path := filepath.Join(folder, fileName)
	err = os.Remove(path)
	if err != nil {
		fmt.Printf("Erro ao deletar arquivo %v: , %v", fileName, err)
	}

	fmt.Printf("Arquivo %v deletado com sucesso!", fileName)
}

func IsEmpty(folder string) (bool, error) {
	f, err := os.Open(folder)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)

	if err == io.EOF {
		return true, nil 
	}

	return false, err
}