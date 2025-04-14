package main

import (
	"auth_service/interfaces/v1/connectrpc"
	"auth_service/internal/infra"
	"os"
	"os/signal"

	"auth_service/internal/repository"
	"fmt"
	"net/http"
	"runtime"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	bus "auth_service/internal/infra/bus"

	infrastructurecore "github.com/ngochuyk812/building_block/infrastructure/core"
	"github.com/ngochuyk812/building_block/infrastructure/databases"
	"github.com/ngochuyk812/building_block/pkg/config"
)

func trackMemory() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Xử lý và trả về thông tin bộ nhớ
	memUsage := fmt.Sprintf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v",
		bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)

	return memUsage
}

// Hàm hỗ trợ chuyển đổi byte sang MB
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
func healHandler(w http.ResponseWriter, r *http.Request) {
	// Theo dõi bộ nhớ mỗi khi route được gọi
	memUsage := trackMemory()

	// Trả về response kèm thông tin bộ nhớ
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("Heal API called\nMemory Usage: %s", memUsage)))
}
func main() {
	policiesPath := &map[string][]string{
		"/greet.v1.GreetService/Greet": {"user"},
	}
	config := config.NewConfigEnv()
	config.PoliciesPath = policiesPath
	infa := infrastructurecore.NewInfra(config)
	infa.InjectSQL(databases.MYSQL)
	infa.InjectCache(config.RedisConnect, config.RedisPass)
	unf := repository.NewUnitOfWork(infa.GetDatabase().GetWriteDB(), infa.GetDatabase().GetReadDB())
	cabin := infra.NewCabin(infa, unf)
	bus.InjectBus(cabin)

	app := infrastructurecore.NewServe(":"+config.Port, infa.GetLogger())
	path, handler := connectrpc.NewSiteServer(cabin)
	app.Mux.Handle(path, handler)

	app.Mux.HandleFunc("/heal", healHandler)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go app.Run()
	<-c
	fmt.Println("shutting down...")

}
