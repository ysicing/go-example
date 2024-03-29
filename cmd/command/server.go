// Copyright (c) 2023 ysicing All rights reserved.
// Use of this source code is governed by WTFPL LICENSE
// license that can be found in the LICENSE file.

package command

import (
	"app/internal/app/server"
	"context"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func ServerCommand() *cobra.Command {
	s := &cobra.Command{
		Use:   "server",
		Short: "core server",
		RunE:  core,
	}
	return s
}

func core(cmd *cobra.Command, args []string) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-ctx.Done()
		stop()
	}()
	return server.Serve(ctx)
}
