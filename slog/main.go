package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type CompositeHandler struct {
	handlers []slog.Handler
}

func NewCompositeHandler(handlers ...slog.Handler) *CompositeHandler {
	return &CompositeHandler{handlers: handlers}
}

func (c *CompositeHandler) Enabled(_ context.Context, level slog.Level) bool {
	fmt.Println("CompositeHandler.Enabled called")
	return true
}

func (c *CompositeHandler) Handle(ctx context.Context, record slog.Record) error {
	fmt.Println("CompositeHandler.Handle called:", record.Message)
	for _, h := range c.handlers {
		if err := h.Handle(ctx, record); err != nil {
			return err
		}
	}
	return nil
}

func (c *CompositeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return c
}

func (c *CompositeHandler) WithGroup(name string) slog.Handler {
	return c
}

func main() {
	textHandler := slog.NewTextHandler(os.Stdout, nil)
	file, _ := os.Create("logs_proper.json")
	defer file.Close()
	jsonHandler := slog.NewJSONHandler(file, nil)

	composite := NewCompositeHandler(textHandler, jsonHandler)

	logger := slog.New(composite)

	logger.Info("There will be information", slog.String("user", "111"))
	logger.Warn("This is a Warning!", slog.Int("code", 123))
}
