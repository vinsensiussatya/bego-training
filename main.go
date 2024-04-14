package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vinsensiussatya/bego-training/cmd"
	"github.com/vinsensiussatya/bego-training/config"
	"github.com/vinsensiussatya/bego-training/internal/pkg/util"
)

func main() {
	_ = os.Setenv("BEGO_ENV", "")
	config.InitConfig()

	util.SetupLog()
	cmd.Execute()
	log.Fatal(http.ListenAndServe(os.Getenv("APP_HOST")+":"+os.Getenv("PORT"), nil))
}
