package main

import "fmt"

//Genero Clase

type taskList struct {
	tasks []*task
}

type task struct {
	name        string
	description string
	complete    bool
}

//Genero metodo
func (t *taskList) addTaskList(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *task) checkTask() {
	t.complete = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *taskList) deleteTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) printList() {
	fmt.Println("------------------------------ LISTA DE TAREAS -----------------------------")
	for _, task := range t.tasks {
		fmt.Println("----------------------------------------------------------------------------")
		fmt.Println(" ")
		fmt.Println("Nombre: ", task.name)
		fmt.Println(" ")
		fmt.Println("Description: ", task.description)
		fmt.Println(" ")
	}
	fmt.Println("----------------------------------------------------------------------------")
}

func (t *taskList) printListCompleted() {
	fmt.Println("------------------------- LISTA DE TAREAS COMPLETADAS ----------------------")
	for _, task := range t.tasks {
		if task.complete {
			fmt.Println("----------------------------------------------------------------------------")
			fmt.Println(" ")
			fmt.Println("Nombre: ", task.name)
			fmt.Println(" ")
			fmt.Println("Description: ", task.description)
			fmt.Println(" ")
			fmt.Println("Completed: ", task.complete)
			fmt.Println(" ")
		}
	}
	fmt.Println("----------------------------------------------------------------------------")
}

func (t *taskList) imprimirLista() {
	for _, e := range t.tasks {
		fmt.Println("\nNombre", e.name, "\nDescripcion", e.description, "\nCompletado", e.complete)
	}
}

func main() {
	t1 := &task{ //Aqui creamos una referencia nueva a task
		name:        "Completar mi curso de GO",
		description: "Completar el curso de GO en esta semana",
	}
	t2 := &task{ //Aqui creamos una referencia nueva a task
		name:        "Completar mi curso de JS",
		description: "Completar este mes ",
	}
	t3 := &task{ //Aqui creamos una referencia nueva a task
		name:        "Completar mi curso de Habilidad blanda",
		description: "Completar este mes ",
	}
	t4 := &task{ //Aqui creamos una referencia nueva a task
		name:        "Completar mi curso de English",
		description: "Completar este mes ",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2, t3,
		},
	}

	lista2 := &taskList{
		tasks: []*task{
			t3, t2,
		},
	}

	fmt.Println(lista.tasks[0])
	lista.addTaskList(t4)
	fmt.Println(len(lista.tasks))
	lista.deleteTask(1)
	fmt.Println(len(lista.tasks))

	// Sintaxis clásica para definir un ciclo for
	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("Index:", i, "task:", lista.tasks[i].name)
	}

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Carlos"] = lista
	mapaTareas["Nestor"] = lista2

	fmt.Println("Tareas de Carlos")
	mapaTareas["Carlos"].imprimirLista()

	fmt.Println("Tareas de Nestor")
	mapaTareas["Nestor"].imprimirLista()

}
