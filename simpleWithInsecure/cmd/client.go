package cmd

import (
	"log"

	"github.com/pkg/errors"
	"github.com/serverhorror/gRPCPlayground/pb"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var opts []grpc.DialOption

		opts = append(opts, grpc.WithInsecure())
		conn, err := grpc.Dial("[::1]:2000", opts...)
		if err != nil {
			return errors.Wrap(err, "could not dial")
		}

		client := pb.NewReplyServiceClient(conn)
		ctx := context.Background()
		req := pb.MsgRequest{
			Value: "I am the request!",
		}
		resp, err := client.Reply(ctx, &req)
		if err != nil {
			return errors.Wrap(err, "problem receiving response")
		}
		log.Printf("%T: %q", resp, resp)
		return err
	},
}

func init() {
	runCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
