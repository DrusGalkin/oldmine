package libs

import (
	"os"
	"os/signal"
	"syscall"
)

func GracefulShutdown(actions func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-sigs

	actions()
}
