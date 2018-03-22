package commands


import (
	"github.com/spf13/cobra"
	"time"
)

var(

	host string
	interval int
)


func init(){
	pollCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Specify Host to poll")
	pollCmd.Flags().IntVar(&interval,"interval",1,"Interval to poll")




}


