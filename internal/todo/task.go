package todo

import "strconv"

type Task struct {
	ID          int
	Description string
	Done        bool
}

// TODO replace with something that doesn't depend on manual editing maybe?
func (t Task) ToStringSlice() []string {
	return []string{
		strconv.Itoa(t.ID),
		t.Description,
		strconv.FormatBool(t.Done),
	}
}

// TODO replace with something that doesn't depend on manual editing maybe?
// Just for simplicity sake, maybe temporary
func GetHeader() []string {
	return []string{"ID", "Description", "Done"}
}
