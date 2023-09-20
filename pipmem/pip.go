package pipmem

import (
	"bytes"
	"context"

	"github.com/benthosdev/benthos/v4/public/service"
)

func init() {
	configSpec := service.NewConfigSpec().
		Summary("Perform PIP queries.").
		Field(service.NewStringField("latField").Default("lat")).
		Field(service.NewStringField("lngField").Default("lng"))

	constructor := func(conf *service.ParsedConfig, mgr *service.Resources) (service.Processor, error) {
		return newPIPProcessor(mgr.Logger(), mgr.Metrics()), nil
	}

	if err := service.RegisterProcessor("pip", configSpec, constructor); err != nil {
		panic(err)
	}
}

type pipProcessor struct {
	logger        *service.Logger
	insideMetrics *service.MetricCounter
}

func newPIPProcessor(logger *service.Logger, metrics *service.Metrics) *pipProcessor {
	// The logger and metrics components will already be labelled with the
	// identifier of this component within a config.
	return &pipProcessor{
		logger:        logger,
		insideMetrics: metrics.NewCounter("inside"),
	}
}

func (r *pipProcessor) Process(ctx context.Context, m *service.Message) (service.MessageBatch, error) {
	bytesContent, err := m.AsBytes()
	if err != nil {
		return nil, err
	}

	newBytes := make([]byte, len(bytesContent))
	for i, b := range bytesContent {
		newBytes[len(newBytes)-i-1] = b
	}

	if bytes.Equal(newBytes, bytesContent) {
		r.logger.Infof("Woah! This is like totally a palindrome: %s", bytesContent)
		r.insideMetrics.Incr(1)
	}

	m.SetBytes(newBytes)
	return []*service.Message{m}, nil
}

func (r *pipProcessor) Close(ctx context.Context) error {
	return nil
}
