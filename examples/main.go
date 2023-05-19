package main

import "github.com/nex-gen-tech/nexlog"

func main() {
	logger := nexlog.NewLogger("NEXLOG")

	logger.Log(nexlog.INF, "Hello World")
	logger.LogF(nexlog.INF, "Hello %s", "World")
	logger.Info("Hello World")
	logger.InfoF("Hello %s", "World")
	logger.Debug("Hello World")
	logger.DebugF("Hello %s", "World")
	logger.Warn("Hello World")
	logger.WarnF("Hello %s", "World")
	logger.Error("Hello World")
	logger.ErrorF("Hello %s", "World")
}
