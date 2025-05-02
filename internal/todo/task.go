package todo

import "strconv"

type Task struct {
	ID          int
	Description string
	Done        bool
}

func (t Task) getCSVData() []string {
	return []string{
		strconv.Itoa(t.ID),
		t.Description,
		strconv.FormatBool(t.Done),
	}
}

// Just for simplicity sake
func getHeader() []string {
	return []string{"ID", "Description", "Done"}
}
