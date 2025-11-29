package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	JSON_FILE_NAME = "tasks.json"
)

type Tasks struct {
	Count int    `json:"count"`
	List  []Task `json:"task_list"`
}

func loadFromFile() (*Tasks, error) {

	// Reading the JSON as byte array
	byteValue, err := os.ReadFile(JSON_FILE_NAME) // ReadFile() já faz o open, read e close do arquivo
	if err != nil {
		if os.IsNotExist(err) {
			return &Tasks{Count: 0, List: make([]Task, 0)}, nil
		}
		return nil, err
	}
	log.Default().Println("> OK: File read into bytes")

	// Decoding the bytes
	var tasks Tasks
	json.Unmarshal(byteValue, &tasks)
	log.Default().Println("> Ok: File decoded through Unmarshal")

	// Returning the memory address
	return &tasks, nil
}

// Function to store data in the JSON file
func saveToFile(tks *Tasks) error {

	byteValue, err := json.Marshal(tks)
	if err != nil {
		return err
	}
	log.Default().Println("> OK: Data coded through Marshal")

	// O writeFile é responsável por abrir, escrever e fechar o arquivo
	os.WriteFile(JSON_FILE_NAME, byteValue, 0644)
	log.Default().Println("> OK: Bytes written to file")

	return nil

}

func main() {

	// Carregando o arquivo
	tks, err := loadFromFile()
	if err != nil {
		log.Fatal(err)
	}

	count := tks.Count
	tksList := tks.List

	// Definindo as Flags
	add := flag.String("add", "", "Add task")
	description := flag.String("description", "", "Description")
	list := flag.Bool("list", false, "List tasks")
	complete := flag.Int("complete", 0, "Complete task")
	delete := flag.Int("delete", 0, "Delete task")
	flag.Parse()

	switch {
	// ADD FLAG
	case *add != "":
		log.Default().Println("> OK: Add task")
		// Criando a nova task com o novo id
		count++
		newTask := createTask(count, *add, *description)
		log.Default().Println("> OK: Added task")

		// Adicionando a nova task na lista
		tksList = append(tksList, *newTask)
		log.Default().Println("> OK: Tasks added to list")

		// Atualizando a struct principal
		tks.Count = count
		tks.List = tksList

		// Salvando no arquivo JSON
		err := saveToFile(tks)
		if err != nil {
			log.Fatal(err)
		}
		log.Default().Println("> OK: Saved to file")

	// LIST FLAG
	case *list:
		log.Default().Println("> OK: List tasks")
		if tksList != nil {
			// Retornando as tasks no terminal
			for _, tk := range tksList {
				fmt.Println(toString(tk))
			}
		} else {
			fmt.Println("Task list is empty")
		}

	case *complete != 0:
		log.Default().Println("> OK: Complete task")

		var found bool
		// Procurando a task
		for i := range tksList {
			// Usamos o tksList[i] para acessar diretamente o valor, sem fazer cópia
			if tksList[i].Id == *complete {
				// Indica que o elemento foi encontrado
				found = true

				// Setando para true
				tksList[i].Complete()
				log.Default().Println("> OK: Completed task")

				// Atualiza a struct pai
				tks.List = tksList

				// Salvando
				err := saveToFile(tks)
				if err != nil {
					log.Fatal(err)
				}
				log.Default().Println("> OK: Saved to file")

				// Saindo do loop após encontrar a task
				break
			}
		}
		if !found {
			log.Default().Println("> ERROR: Task not found")
		}

	case *delete != 0:
		var found bool
		log.Default().Println("> OK: Delete task")
		if *delete > count {
			log.Default().Println("> ERROR: Task not found")
		}
		for i := range tksList {
			if tksList[i].Id == *delete {
				found = true
				tksList = append(tksList[:i], tksList[i+1:]...)
				log.Default().Println("> OK: Deleted task")
				tks.List = tksList
				tks.Count = count
				err := saveToFile(tks)
				if err != nil {
					log.Fatal(err)
				}
				log.Default().Println("> OK: Saved to file")
				break
			}
		}

		if !found {
			log.Default().Println("> ERROR: Task not found")
		}

	}

}
