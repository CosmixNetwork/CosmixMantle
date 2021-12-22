/*
 Copyright [2019] - [2020], Cosmix TECHNOLOGIES PTE. LTD. and the cosmixMantle contributors
 SPDX-License-Identifier: Apache-2.0
*/

package application

import (
	"github.com/CosmixOne/cosmixMantle/application/internal/configurations"
	"github.com/CosmixOne/CosmixSDK/schema/applications/base"
)

var Prototype = base.NewApplication(
	configurations.Name,
	configurations.ModuleBasicManager,
	configurations.EnabledWasmProposalTypeList,
	configurations.ModuleAccountPermissions,
	configurations.TokenReceiveAllowedModules,
)
