package cli

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/windmilleng/tilt/internal/container"
	"github.com/windmilleng/tilt/internal/k8s"
	"github.com/windmilleng/tilt/internal/logger"
	"github.com/windmilleng/tilt/internal/options"
	"github.com/windmilleng/tilt/internal/synclet"
	"github.com/windmilleng/tilt/internal/synclet/proto"
	"github.com/windmilleng/tilt/internal/tracer"
)

type SyncletCmd struct {
	port        int
	debug       bool
	verbose     bool
	criEndpoint string
}

func (s *SyncletCmd) Register() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "synclet",
		Short: "starts the tilt synclet daemon",
		Run: func(c *cobra.Command, args []string) {
			s.run()
		},
	}

	cmd.Flags().BoolVarP(&s.debug, "debug", "d", false, "Enable debug logging")
	cmd.Flags().BoolVarP(&s.verbose, "verbose", "v", false, "Enable verbose logging")
	cmd.Flags().IntVar(&s.port, "port", synclet.Port, "Server port")
	cmd.Flags().StringVar(&s.criEndpoint, "cri-endpoint", "", "Endpoint of CRI container runtime service")

	return cmd
}

func (sc *SyncletCmd) run() {
	ctx := logger.WithLogger(
		context.Background(),
		logger.NewLogger(logLevel(sc.verbose, sc.debug), os.Stdout))

	closer, err := tracer.Init(ctx)
	if err != nil {
		log.Fatalf("error initializing tracer: %v", err)
	}
	defer func() {
		err := closer()
		if err != nil {
			log.Fatalf("error closing tracer: %v", err)
		}
	}()

	addr := fmt.Sprintf("127.0.0.1:%d", sc.port)
	log.Printf("Running synclet listening on %s", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TODO(matt) figure out how to reconcile this with opt-in tracing
	t := opentracing.GlobalTracer()

	opts := options.MaxMsgServer()
	opts = append(opts, options.TracingInterceptorsServer(t)...)

	serv := grpc.NewServer(opts...)

	var s synclet.Synclet
	if sc.criEndpoint == "" {
		s, err = synclet.WireDockerSynclet(ctx, k8s.EnvUnknown, container.RuntimeDocker)
	} else {
		s, err = synclet.WireCriSynclet(ctx, synclet.Endpoint(sc.criEndpoint))
	}
	if err != nil {
		log.Printf("failed to wire synclet: %v", err)
		time.Sleep(3600 * time.Second)
		log.Fatalf("failed to wire synclet: %v", err)
	}

	proto.RegisterSyncletServer(serv, synclet.NewGRPCServer(s))

	err = serv.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
