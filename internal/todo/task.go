package todo

import "strconv"

type Task struct {
	id          int
	description string
	done        bool
}

func (t *Task) ID() int             { return t.id }
func (t *Task) Description() string { return t.description }
func (t *Task) Done() bool          { return t.done }

const CSVTaskColumnCount = 3

func (t *Task) ToStringSlice() []string {
	return []string{
		strconv.Itoa(t.id),
		t.description,
		strconv.FormatBool(t.done),
	}
}

// Just for simplicity sake, maybe temporary
func GetHeader() []string {
	return []string{"ID", "Description", "Done"}
}
