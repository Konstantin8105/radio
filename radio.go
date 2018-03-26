package radio

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type cliCommand struct {
	command     string
	description string
	action      func(args []string) (isExit bool, err error)
}

type cli struct {
	commands []cliCommand
}

func (c *cli) register(
	command string,
	description string,
	action func(args []string) (isExit bool, err error)) {

	if command == "" {
		panic("Command cannot be empty")
	}
	if description == "" {
		panic("Description cannot be empty")
	}
	for i := range c.commands {
		if command == c.commands[i].command {
			panic(fmt.Errorf("Dublicate of command : %v", command))
		}
	}
	c.commands = append(c.commands, cliCommand{
		command:     command,
		description: description,
		action:      action,
	})
}

func (c *cli) run(com string, args ...string) (isExit bool, err error) {
	// empty command
	if com == "" {
		return
	}
	// is command exists
	for _, command := range c.commands {
		if strings.ToLower(com) == strings.ToLower(command.command) {
			return command.action(args)
		}
	}
	// command is not found
	err = fmt.Errorf("Don`t found command: `%v`", com)
	return
}

func showWithMaxLength(str string, length int) string {
	if len(str) < length {
		return str
	}
	return str[:length]
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

	// search stations
	stations, err := getStations()
	if err != nil {
		return
	}

	// add seef for random number
	rand.Seed(time.Now().UTC().UnixNano())

	// run user interface
	ui := &cli{}

	ui.register("info", "Return information from player about current stream",
		func(args []string) (isExit bool, err error) {
			err = vlc.CommandToVLC("info\n")
			return
		})

	ui.register("search", "Search [query] a specific name",
		func(args []string) (isExit bool, err error) {
			if len(args) != 1 {
				err = fmt.Errorf("Not allowable command. Example: `search pop`")
				return
			}
			fmt.Println("Search by query : ", args[0])
			fmt.Printf("|%20s|%20s|%20s|\n", "ID", "Name", "Genre")
			for _, station := range stations {
				var add bool
				if strings.Contains(
					strings.ToLower(station.Name),
					strings.ToLower(args[0])) {
					add = true
				}
				if strings.Contains(
					strings.ToLower(station.Genre),
					strings.ToLower(args[0])) {
					add = true
				}
				if !add {
					continue
				}
				fmt.Printf("|%20s|%20s|%20s|\n",
					strconv.Itoa(station.ID),
					showWithMaxLength(station.Name, 20),
					showWithMaxLength(station.Genre, 20))
			}
			return
		})

	ui.register("list", "List of all allowable radio stations",
		func(args []string) (isExit bool, err error) {
			fmt.Printf("|%20s|%20s|%20s|\n", "ID", "Name", "Genre")
			for _, station := range stations {
				fmt.Printf("|%20s|%20s|%20s|\n",
					strconv.Itoa(station.ID),
					showWithMaxLength(station.Name, 20),
					showWithMaxLength(station.Genre, 20))
			}
			return
		})

	ui.register("title", "Title of the current stream",
		func(args []string) (isExit bool, err error) {
			err = vlc.CommandToVLC("get_title\n")
			return
		})

	ui.register("exit", "Exit from terminal radio",
		func(args []string) (isExit bool, err error) {
			return true, nil
		})

	ui.register("play", "Play [station], is [station] is empty, then playing random station",
		func(args []string) (isExit bool, err error) {
			var stationID string
			if len(args) == 0 {
				stationID = strconv.Itoa(stations[rand.Intn(len(stations))].ID)
			} else {
				stationID = args[0]
			}
			url := "http://yp.shoutcast.com/sbin/tunein-station.pls?id=" + stationID
			err = vlc.CommandToVLC(fmt.Sprintf("add %s\n", url))
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
			err = vlc.CommandToVLC("stop\n")
			return
		})

	ui.register("resume", "Continue playing stopped stream",
		func(args []string) (isExit bool, err error) {
			err = vlc.CommandToVLC("stop\n")
			return
		})

	ui.register("clear", "Clear playlist in player",
		func(args []string) (isExit bool, err error) {
			err = vlc.CommandToVLC("clear\n")
			return
		})

	// print help information
	_, _ = ui.run("help")

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
