package consul

import (
	"errors"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"strconv"
	"time"

	"github.com/vladopajic/go-actor/actor"
	"link-society.com/flowg/internal/utils/proctree"

	"github.com/hashicorp/consul/api"
)

type procHandler struct {
	consulClient *api.Client
	logger       *slog.Logger
	opts         *ConsulServiceOptions
}

const (
	getNodesMaxRetries  = 10
	healthCheckPath     = "/health"
	healthCheckInterval = 5 * time.Second
	healthCheckTimeout  = 1 * time.Second
	shutdownTimeout     = 5 * time.Second
)

func (h *procHandler) Init(ctx actor.Context) proctree.ProcessResult {
	// If no consul url is provided then stop the consul service as it is not needed
	if h.opts.ConsulUrl == "" {
		h.logger.InfoContext(ctx, "no consul url provided")
		return proctree.Terminate(nil)
	}

	// Register node with Consul
	if err := h.registerNode(ctx); err != nil {
		h.logger.ErrorContext(
			ctx,
			"failed to start Consul service",
			slog.Any("error", err),
		)
		return proctree.Terminate(err)
	}

	// Get other nodes from consul to form a cluster with memberlist
	_, err := h.getNodes(ctx)
	if err != nil {
		h.logger.WarnContext(
			ctx,
			"failed to get service nodes from consul",
			slog.Any("error", err),
		)
	}

	// Form cluster using memberlist

	return proctree.Continue()
}

func (h *procHandler) DoWork(ctx actor.Context) proctree.ProcessResult {
	<-ctx.Done()
	return proctree.Terminate(ctx.Err())
}

func (h *procHandler) Terminate(ctx actor.Context, err error) error {
	h.logger.InfoContext(ctx, "Deregistering service with consul")

	if deregisterErr := h.consulClient.Agent().ServiceDeregister(h.opts.NodeId); err != nil {
		h.logger.ErrorContext(
			ctx,
			"Failed to shutdown HTTP server",
			slog.String("error", deregisterErr.Error()),
		)
		err = errors.Join(err, deregisterErr)
	}

	return err
}

func (h *procHandler) registerNode(ctx actor.Context) error {
	// Create a Consul client
	config := api.DefaultConfig()
	config.Address = h.opts.ConsulUrl
	client, err := api.NewClient(config)
	if err != nil {
		h.logger.ErrorContext(
			ctx,
			"failed to create Consul client",
			slog.Any("error", err),
		)
		return err
	}
	h.consulClient = client
	var port int
	port, err = strconv.Atoi(h.opts.NodePort)
	if err != nil {
		h.logger.ErrorContext(
			ctx,
			"error converting port from string to int",
			slog.Any("error", err),
		)
		return err
	}

	// Define the service registration
	registration := &api.AgentServiceRegistration{
		ID:      h.opts.NodeId,
		Name:    h.opts.ServiceName,
		Address: h.opts.NodeHost,
		Port:    port,
		Check: &api.AgentServiceCheck{
			Interval: healthCheckInterval.String(),
			HTTP:     fmt.Sprintf("http://%s:%d%s", h.opts.NodeHost, port, healthCheckPath),
			Timeout:  healthCheckTimeout.String(),
		},
	}

	// Register the service with Consul
	if err = client.Agent().ServiceRegister(registration); err != nil {
		h.logger.ErrorContext(
			ctx,
			"failed to register service with Consul",
			slog.Any("error", err),
		)
		return err
	}

	return nil
}

// getNodes() retries with exponential backoff with jitter to fetch other nodes in the cluster using consul
func (h *procHandler) getNodes(ctx actor.Context) ([]string, error) {
	retryCount := 0
	delay := 100 * time.Millisecond

	for retryCount <= getNodesMaxRetries {
		nodes, _, err := h.consulClient.Health().Service(h.opts.ServiceName, "", false, nil)
		if err != nil {
			h.logger.ErrorContext(
				ctx,
				"failed to get nodes from consul",
				slog.Any("error", err),
			)
			return nil, err
		}

		otherNodes := []string{}
		for _, node := range nodes {
			if node.Service.ID != h.opts.NodeId {
				otherNodes = append(otherNodes, node.Node.Address)
			}
		}

		if len(otherNodes) >= 1 {
			return otherNodes, nil
		}

		retryCount++
		if retryCount <= getNodesMaxRetries {
			h.logger.InfoContext(ctx, "did not find other nodes, will try again with a delay")
			time.Sleep(delay)
			// Exponential backoff
			delay = delay * 2
			// Add jitter to the delay
			delay += time.Duration(rand.IntN(int(delay / 4)))
		}
	}

	return nil, errors.New("failed to find other nodes")

}
