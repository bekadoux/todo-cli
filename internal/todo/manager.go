package todo

type TaskManager struct {
	lastID   int
	tasks    []*Task
	newTasks []*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		lastID:   -1,
		tasks:    make([]*Task, 0),
		newTasks: make([]*Task, 0),
	}
}

func (m *TaskManager) NewTask(description string, done bool) *Task {
	m.lastID++
	task := &Task{id: m.lastID, description: description, done: done}
	m.newTasks = append(m.newTasks, task)
	return task
}

func (m *TaskManager) AddTaskFromStorage(id int, description string, done bool) {
	task := &Task{id: id, description: description, done: done}
	m.tasks = append(m.tasks, task)
	if id > m.lastID {
		m.lastID = id
	}
}

func (m *TaskManager) ForEachTask(fn func(*Task)) {
	for _, t := range m.tasks {
		fn(t)
	}
}

func (m *TaskManager) ForEachNewTask(fn func(*Task) error) error {
	for _, t := range m.newTasks {
		if err := fn(t); err != nil {
			return err
		}
	}
	return nil
}
