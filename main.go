package main

import (
	"fmt"
	"golang.org/x/sys/windows/svc"
	"log"
	"os"
	"github.com/leandroveronezi/exemple-golang-windows-service/controller"
	"strings"
	"time"
)

var serviceRunning = false

func usage(errmsg string) {

	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"usage: %s <command>\n"+
			"       where <command> is one of\n"+
			"       install, remove, debug, start, stop, pause or continue.\n",
		errmsg, os.Args[0])

	os.Exit(2)
}

func main() {

	var err error

	const svcName = "NomeDoMeuServico"

	isIntSess, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Println(fmt.Sprintf("failed to determine if we are running in an interactive session: %v", err))
	}

	if !isIntSess {


		controller.RunService(svcName, false)

		for {
			time.Sleep(time.Millisecond * 200)
			if serviceRunning {
				break
			}
		}

		return

	}

	if len(os.Args) < 2 {
		log.Println("usage")
		usage("no command specified")
	}

	cmd := strings.ToLower(os.Args[1])

	switch cmd {

	case "debug":
		controller.RunService(svcName, true)
		return

	case "install":
		err = controller.InstallService(svcName, "Descrição do serviço")

	case "remove":
		err = controller.RemoveService(svcName)

	case "start":
		err = controller.StartService(svcName)

	case "stop":
		err = controller.ControlService(svcName, svc.Stop, svc.Stopped)

	case "pause":
		err = controller.ControlService(svcName, svc.Pause, svc.Paused)

	case "continue":
		err = controller.ControlService(svcName, svc.Continue, svc.Running)

	default:
		usage(fmt.Sprintf("invalid command %s", cmd))
	}

	if err != nil {
		log.Println(fmt.Sprintf("failed to %s %s: %v", cmd, svcName, err))
	}

	os.Exit(0)
}
