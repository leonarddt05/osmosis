package v25

import (
	"github.com/osmosis-labs/osmosis/v24/app/upgrades"

	store "github.com/cosmos/cosmos-sdk/store/types"

	authenticatortypes "github.com/osmosis-labs/osmosis/v24/x/authenticator/types"
)

// UpgradeName defines the on-chain upgrade name for the Osmosis v25 upgrade.
const UpgradeName = "v25"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			authenticatortypes.ManagerStoreKey,
			authenticatortypes.AuthenticatorStoreKey,
		},
		Deleted: []string{},
	},
}
