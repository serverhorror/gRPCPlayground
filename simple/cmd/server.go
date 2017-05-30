package cmd

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		var lAddr net.IP
		var lPort int64
		var l net.Listener

		if lAddr = net.ParseIP(os.Getenv("IP")); lAddr == nil {
			return fmt.Errorf("could not parse IP: %v", lAddr)
			// return errors.Wrap(err, "could not get IP")
		}
		if lPort, err = strconv.ParseInt(os.Getenv("PORT"), 10, 64); err != nil {
			return errors.Wrap(err, "could not get PORT")
		}

		if l, err = net.Listen("tcp", fmt.Sprintf("[%s]:%v", lAddr, lPort)); err != nil {
			return errors.Wrap(err, "could not listen")
		}

		con, err := l.Accept()
		_, err = con.Write([]byte("bye!\n"))
		err = con.Close()

		return err
	},
}

func init() {
	runCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
