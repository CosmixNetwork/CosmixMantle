/*
 Copyright [2019] - [2020], Cosmix TECHNOLOGIES PTE. LTD. and the cosmixMantle contributors
 SPDX-License-Identifier: Apache-2.0
*/

package initialize

import (
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func ValidateGenesisCommand(
	serverContext *server.Context,
	codec *codec.Codec,
	moduleBasicManager module.BasicManager,
) *cobra.Command {
	return cli.ValidateGenesisCmd(
		serverContext,
		codec,
		moduleBasicManager,
	)
}
