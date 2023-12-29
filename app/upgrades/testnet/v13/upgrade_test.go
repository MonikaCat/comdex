package v13_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/suite"

// 	"github.com/MonikaCat/comdex/v13/app"
// 	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
// 	v13 "github.com/MonikaCat/comdex/v13/app/upgrades/testnet/v13"
// )

// type UpgradeTestSuite struct {
// 	app.KeeperTestHelper
// }

// func (s *UpgradeTestSuite) SetupTest() {
// 	s.Setup()
// }

// func TestKeeperTestSuite(t *testing.T) {
// 	suite.Run(t, new(UpgradeTestSuite))
// }

// // Ensures the test does not error out.
// func (s *UpgradeTestSuite) TestUpgrade() {
// 	s.Setup()

// 	preUpgradeChecks(s)

// 	upgradeHeight := int64(5)
// 	s.ConfirmUpgradeSucceeded(v13.UpgradeName, upgradeHeight)

// 	postUpgradeChecks(s)
// }

// func preUpgradeChecks(s *UpgradeTestSuite) {

// 	mp := s.App.MintKeeper.GetParams(s.Ctx)
// 	s.Require().Equal(mp.BlocksPerYear, uint64(6311520))

// 	sp := s.App.SlashingKeeper.GetParams(s.Ctx)
// 	s.Require().Equal(sp.SignedBlocksWindow, int64(100))

// }

// func postUpgradeChecks(s *UpgradeTestSuite) {

// 	// Ensure the gov params have MinInitialDepositRatio added
// 	gp := s.App.GovKeeper.GetParams(s.Ctx)
// 	s.Require().Equal(gp.MinInitialDepositRatio, "0.200000000000000000")

// 	// Ensure the mint params have doubled
// 	mp := s.App.MintKeeper.GetParams(s.Ctx)
// 	s.Require().Equal(mp.BlocksPerYear, uint64(6311520*2))

// 	// Ensure the slashing params have doubled
// 	sp := s.App.SlashingKeeper.GetParams(s.Ctx)
// 	s.Require().Equal(sp.SignedBlocksWindow, int64(100*2))

// 	// Ensure the wasm Permissionless
// 	wp := s.App.WasmKeeper.GetParams(s.Ctx)
// 	s.Require().Equal(wp.CodeUploadAccess, wasmtypes.AllowEverybody)
// }
