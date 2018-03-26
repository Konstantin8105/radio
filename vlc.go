package radio

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"

	freeport "github.com/Konstantin8105/FreePort"
)

// VLC media player server
type VLC struct {
	connection net.Conn
}

// NewVLC create a new VLC server
func NewVLC() (v *VLC, err error) {
	// basic initialization
	v = &VLC{}

	// run server vlc
	var freePort int
	freePort, err = freeport.Get()
	if err != nil {
		return
	}

	// start vlc us server
	_, err = exec.LookPath("cvlc")
	if err != nil {
		err = fmt.Errorf("Please install VLC. %v", err)
		return
	}
	addr := "localhost:" + strconv.Itoa(freePort)
	go func() {
		_, err = exec.Command("cvlc", "--intf", "rc", "--rc-host",
			addr).Output()
		if err != nil {
			err = fmt.Errorf("Cannot run CVLC like server. %v", err)
			return
		}
	}()

	// check connection to vlc server
	// timeout is added, because need time for run server
	for i := 0; i < 40; i++ {
		time.Sleep(500 * time.Millisecond)
		v.connection, err = net.Dial("tcp", addr)
		if err == nil {
			break
		}
	}
	if err != nil {
		err = fmt.Errorf("Cannot connect to VLC server on port: %v", err)
		return
	}

	// start listening the server
	r := bufio.NewReader(v.connection)
	// Drop VLC greeting
	r.ReadString('\n')
	r.ReadString('\n')
	go func() {
		for {
			s, err := r.ReadString('\n')
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error in VLC: %v", err)
				return
			}
			fmt.Printf(s)
		}
	}()

	return
}

func (v *VLC) CommandToVLC(command string) (err error) {
	if len(command) == 0 {
		err = fmt.Errorf("Command for VLC cannot be empty")
		return
	}
	// send command to VLC
	_, err = v.connection.Write([]byte(command))
	return
}

// Close VLC server and close connection
func (v *VLC) Close() (err error) {
	// shutdown VLC server
	err = v.CommandToVLC("shutdown\n")
	if err != nil {
		return
	}

	// close connection
	err = v.connection.Close()
	if err != nil {
		return
	}
	return nil
}
