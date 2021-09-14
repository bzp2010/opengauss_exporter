package scrape

type Manager struct {
	tasks   []*Task
	stopSig chan bool
}

func NewManager() *Manager {
	return &Manager{
		tasks:   []*Task{},
		stopSig: make(chan bool, 100),
	}
}

func (m *Manager) AddTask(t *Task) {
	m.tasks = append(m.tasks, t)
}

func (m *Manager) Start() {
	for _, task := range m.tasks {
		go task.Start()
	}
}

func (m *Manager) Stop() {
	for _, task := range m.tasks {
		task.Stop()
	}
}
