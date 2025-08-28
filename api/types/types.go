package types

type Task struct {
	Id     string `json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
}

type CreateTaskRequest struct {
	Text string `json:"text"`
}

type DeleteTaskRequest struct {
	Id string `json:"id"`
}

type CompleteTaskRequest struct {
	Id string `json:"id"`
}

type TasksList struct {
	Tasks []Task `json:"tasks"`
}
