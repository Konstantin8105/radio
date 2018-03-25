package radio

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type ui struct {
	command     string
	description string
	f           func(args []string) (isExit bool, err error)
}

type terminalUI struct {
	commands []ui
}

func (u *terminalUI) register(
	command string,
	comDescription string,
	comFunction func(args []string) (isExit bool, err error)) {

	if command == "" {
		panic("Command cannot be empty")
	}
	if comDescription == "" {
		panic("Description cannot be empty")
	}
	for i := range u.commands {
		if command == u.commands[i].command {
			panic(fmt.Errorf("Dublicate of command : %v", command))
		}
	}
	u.commands = append(u.commands, ui{
		command:     command,
		description: comDescription,
		f:           comFunction,
	})
}

func (ui *terminalUI) run(com string, args ...string) (isExit bool, err error) {
	if com == "" {
		return
	}
	for _, c := range ui.commands {
		if strings.ToLower(com) == strings.ToLower(c.command) {
			return c.f(args)
		}
	}
	err = fmt.Errorf("Don`t found command: `%v`", com)
	return
}

// Run radio
func Run() (err error) {
	// show header of Xi-radio
	fmt.Println("Start : Ξ (Xi-radio)")
	fmt.Println("Enter 'help' for show all commands")

	// run VLC like a media server
	var vlc *VLC
	vlc, err = NewVLC()
	if err != nil {
		return
	}
	defer func() {
		errClose := vlc.Close()
		if err == nil {
			err = errClose
		} else {
			err = fmt.Errorf("%v. %v", err, errClose)
		}
	}()

	// run user interface
	ui := &terminalUI{}

	ui.register("info", "Return information from player about current stream",
		func(args []string) (isExit bool, err error) {
			vlc.CommandToVLC("info\n")
			return
		})

	ui.register("list", "List of allowable radio stations",
		func(args []string) (isExit bool, err error) {
			stations, err := GetStations()
			if err != nil {
				return
			}
			fmt.Printf("|%20s|%20s|%20s|", "ID", "Name", "Genre")
			for _, station := range stations {
				fmt.Println("|%20s|%20s|%20s|",
					station.ID, station.Name, station.Genre)
			}
			return
		})

	ui.register("title", "Title of the current stream",
		func(args []string) (isExit bool, err error) {
			vlc.CommandToVLC("get_title\n")
			return
		})

	ui.register("exit", "Exit from terminal radio",
		func(args []string) (isExit bool, err error) {
			return true, nil
		})

	ui.register("play", "Play [station], is [station] is empty, then playing random station",
		func(args []string) (isExit bool, err error) {
			var stationID int
			if len(args) == 0 {
				stationID = getRandomStation()
			} else {
				stationID = args[0]
			}
			url := "http://yp.shoutcast.com/sbin/tunein-station.pls?id=" + id
			p.SendCommandToVLC(fmt.Sprintf("add %s\n", url))
			fmt.Println("run station : ", stationID)
			return
		})

	ui.register("help", "Show all commands",
		func(args []string) (isExit bool, err error) {
			sort.Slice(ui.commands, func(i, j int) bool {
				return strings.Compare(
					ui.commands[i].command,
					ui.commands[j].command) < 0
			})
			_ = args
			for _, c := range ui.commands {
				fmt.Printf("%10s\t%s\n", c.command, c.description)
			}
			return
		})

	ui.register("stop", "Stop stream in player",
		func(args []string) (isExit bool, err error) {
			vlc.CommandToVLC("stop\n")
			return
		})

	ui.register("resume", "Continue playing stopped stream",
		func(args []string) (isExit bool, err error) {
			vlc.CommandToVLC("stop\n")
			return
		})

	ui.register("clear", "Clear playlist in player",
		func(args []string) (isExit bool, err error) {
			vlc.CommandToVLC("clear\n")
			return
		})

	// print help information
	ui.run("help")

	// TODO: run last used station

	for {
		fmt.Printf("Ξ> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		cmd := strings.Split(strings.TrimSpace(scanner.Text()), " ")

		// remove empty cmd elements
	checkAgain:
		for i := range cmd {
			if len(strings.TrimSpace(cmd[i])) == 0 {
				if i == len(cmd)-1 {
					cmd = cmd[:i]
				} else {
					cmd = append(cmd[:i], cmd[i+1:]...)
				}
				goto checkAgain
			}
		}

		if len(cmd) == 0 {
			continue
		}
		var isExit bool
		if len(cmd) == 1 {
			isExit, err = ui.run(cmd[0])
		} else {
			isExit, err = ui.run(cmd[0], cmd[1:]...)
		}
		if err != nil {
			fmt.Println("err = ", err)
		}
		if isExit {
			break
		}
	}

	// close radio
	fmt.Println("Close : Ξ (Xi-radio)")

	return
}
