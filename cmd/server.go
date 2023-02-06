package cmd

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	brickwallv1 "github.com/programmablemike/ngrok-go-example/gen/go/brickwall/v1"
	"github.com/programmablemike/ngrok-go-example/gen/go/brickwall/v1/brickwallv1connect"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewServerCommand() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "starts the server",
		Action: func(cCtx *cli.Context) error {
			return RunServer(context.Background())
		},
	}
}

type BrickwallServer struct{}

func (bs *BrickwallServer) CheckBlocklist(ctx context.Context, req *connect.Request[brickwallv1.CheckBlocklistRequest]) (*connect.Response[brickwallv1.CheckBlocklistResponse], error) {
	log.Info().Msgf("Called CheckblockList with request %v", req)
	// TODO(mlee): Implement the block list checking logic
	return connect.NewResponse(&brickwallv1.CheckBlocklistResponse{
		Id:      "myid",
		Blocked: brickwallv1.BlockStatus_NOT_BLOCKED,
	}), nil

}

func RunServer(ctx context.Context) error {
	l, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}
	log.Info().Msgf("Listening at %s", l.Addr())
	brickwall := &BrickwallServer{}
	mux := http.NewServeMux()
	path, handler := brickwallv1connect.NewBrickwallServiceHandler(brickwall)
	mux.Handle(path, handler)
	return http.Serve(l, h2c.NewHandler(mux, &http2.Server{}))
}
