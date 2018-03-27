package radio

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestPositive(t *testing.T) {

	commands := []string{
		"help\n",
		"list\n",
		"   search \n",
		"search pop\n",
		"play\n",
		"title\n",
		"info\n",
		"play 123212312\n",
		"\n",
		"someWrongCommand\n",
		"play\n",
		"stop\n",
		"resume\n",
		"clear\n",
		"exit\n",
	}

	r, w, _ := os.Pipe()

	var isDone bool

	go func() {
		time.Sleep(5 * time.Second)
		for _, command := range commands {
			if isDone {
				// some think wrong
				break
			}
			fmt.Println("command = ", command)
			time.Sleep(1 * time.Second)
			w.WriteString(command)
		}
	}()

	go func() {
		// run radio
		err := run(r)
		if err != nil {
			t.Errorf("Not acceptable result. %v", err)
		}
		isDone = true
	}()

	time.Sleep(time.Duration(len(commands)+10) * time.Second)
	if !isDone {
		t.Fatal("Timeout")
	}
}
