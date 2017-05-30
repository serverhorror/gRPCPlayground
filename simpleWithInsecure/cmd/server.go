package cmd

import (
	"net"

	"github.com/pkg/errors"
	"github.com/serverhorror/gRPCPlayground/pb"
	"github.com/serverhorror/gRPCPlayground/simpleWithInsecure/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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
		s := server.EchoServer{}

		l, err := net.Listen("tcp", "[::1]:2000")
		if err != nil {
			return errors.Wrap(err, "could not start listener")
		}

		grpcServer := grpc.NewServer()
		pb.RegisterReplyServiceServer(grpcServer, &s)
		grpcServer.Serve(l)
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
