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

	fmt.Println(lista.tasks[0])
	lista.addTaskList(t4)
	fmt.Println(len(lista.tasks))
	lista.deleteTask(1)
	fmt.Println(len(lista.tasks))

}
