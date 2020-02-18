package controller

import (
	"fmt"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"log"
	"os"
	"strings"
	"time"
)

func RunService(name string, isDebug bool) {

	if isDebug {
		log.Println("isDebug")

	} else {

		log.Println("not Debug")

	}

	log.Println(fmt.Sprintf("starting service %s", name))

	run := svc.Run
	if isDebug {
		run = debug.Run
	}

	err := run(name, &s{})

	if err != nil {
		log.Println(fmt.Sprintf("%s service failed: %v", name, err))
		return
	}

	log.Println(fmt.Sprintf("%s service stopped", name))
	os.Exit(0)

}

type s struct{}

func (m *s) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {

	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue
	changes <- svc.Status{State: svc.StartPending}

	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	go RunMeuServico()

loop:

	for {
		select {

		case c := <-r:
			switch c.Cmd {

			case svc.Interrogate:

				changes <- c.CurrentStatus
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus

			case svc.Stop, svc.Shutdown:

				testOutput := strings.Join(args, "-")
				testOutput += fmt.Sprintf("-%d", c.Context)
				log.Println("Output")
				break loop

			case svc.Pause:

				//changes <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}

			case svc.Continue:

				//changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

			default:
				log.Println(fmt.Sprintf("unexpected control request #%d", c))

			}
		}
	}

	changes <- svc.Status{State: svc.StopPending}

	return
}


func RunMeuServico() {

	/*
		Aqui é implementado o código do meu programa
	 */

}
