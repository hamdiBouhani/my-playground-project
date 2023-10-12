package taskes

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("init")
}

func TestTasks(t *testing.T) {
	t.Skip("skipping test in short mode.")
	tasks := New()

	tasks.Add(func() error {
		fmt.Println("job 1")
		return nil
	}, 1)

	tasks.Add(func() error {
		fmt.Println("job 2")
		return nil
	}, 1)

	tasks.Run()

}
