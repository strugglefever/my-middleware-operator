package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/wflysnow/my-middleware-operator/cmd/operator-manager/app"
	"github.com/wflysnow/my-middleware-operator/cmd/operator-manager/app/options"
	"k8s.io/apiserver/pkg/util/flag"
	"k8s.io/apiserver/pkg/util/logs"
	"k8s.io/kubernetes/pkg/version/verflag"
)

func main() {
	s := options.NewOMServer()
	s.AddFlags(pflag.CommandLine, app.KnownOperators())
	flag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()
	verflag.PrintAndExitIfRequested()
	if err := app.Run(s); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("main method execute end!\n")
}
