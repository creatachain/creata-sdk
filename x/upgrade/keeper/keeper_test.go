package keeper_test

import (
	"path/filepath"
	"testing"
	"time"

	tmproto "github.com/creatachain/augusteum/proto/augusteum/types"
	"github.com/stretchr/testify/suite"

	"github.com/creatachain/creata-sdk/creataapp"
	store "github.com/creatachain/creata-sdk/store/types"
	sdk "github.com/creatachain/creata-sdk/types"
	clienttypes "github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	commitmenttypes "github.com/creatachain/creata-sdk/x/icp/core/23-commitment/types"
	icpexported "github.com/creatachain/creata-sdk/x/icp/core/exported"
	icptmtypes "github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
	"github.com/creatachain/creata-sdk/x/upgrade/keeper"
	"github.com/creatachain/creata-sdk/x/upgrade/types"
)

type KeeperTestSuite struct {
	suite.Suite

	homeDir string
	app     *creataapp.CreataApp
	ctx     sdk.Context
}

func (s *KeeperTestSuite) SetupTest() {
	app := creataapp.Setup(false)
	homeDir := filepath.Join(s.T().TempDir(), "x_upgrade_keeper_test")
	app.UpgradeKeeper = keeper.NewKeeper( // recreate keeper in order to use a custom home path
		make(map[int64]bool), app.GetKey(types.StoreKey), app.AppCodec(), homeDir,
	)
	s.T().Log("home dir:", homeDir)
	s.homeDir = homeDir
	s.app = app
	s.ctx = app.BaseApp.NewContext(false, tmproto.Header{
		Time:   time.Now(),
		Height: 10,
	})
}

func (s *KeeperTestSuite) TestReadUpgradeInfoFromDisk() {
	// require no error when the upgrade info file does not exist
	_, err := s.app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	s.Require().NoError(err)

	expected := store.UpgradeInfo{
		Name:   "test_upgrade",
		Height: 100,
	}

	// create an upgrade info file
	s.Require().NoError(s.app.UpgradeKeeper.DumpUpgradeInfoToDisk(expected.Height, expected.Name))

	ui, err := s.app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	s.Require().NoError(err)
	s.Require().Equal(expected, ui)
}

func (s *KeeperTestSuite) TestScheduleUpgrade() {
	clientState := &icptmtypes.ClientState{ChainId: "creatachain"}
	cs, err := clienttypes.PackClientState(clientState)
	s.Require().NoError(err)

	altClientState := &icptmtypes.ClientState{ChainId: "ethermint"}
	altCs, err := clienttypes.PackClientState(altClientState)
	s.Require().NoError(err)

	consState := icptmtypes.NewConsensusState(time.Now(), commitmenttypes.NewMerkleRoot([]byte("app_hash")), []byte("next_vals_hash"))
	consAny, err := clienttypes.PackConsensusState(consState)
	s.Require().NoError(err)

	cases := []struct {
		name    string
		plan    types.Plan
		setup   func()
		expPass bool
	}{
		{
			name: "successful time schedule",
			plan: types.Plan{
				Name: "all-good",
				Info: "some text here",
				Time: s.ctx.BlockTime().Add(time.Hour),
			},
			setup:   func() {},
			expPass: true,
		},
		{
			name: "successful height schedule",
			plan: types.Plan{
				Name:   "all-good",
				Info:   "some text here",
				Height: 123450000,
			},
			setup:   func() {},
			expPass: true,
		},
		{
			name: "successful icp schedule",
			plan: types.Plan{
				Name:                "all-good",
				Info:                "some text here",
				Height:              123450000,
				UpgradedClientState: cs,
			},
			setup:   func() {},
			expPass: true,
		},
		{
			name: "successful overwrite",
			plan: types.Plan{
				Name:   "all-good",
				Info:   "some text here",
				Height: 123450000,
			},
			setup: func() {
				s.app.UpgradeKeeper.ScheduleUpgrade(s.ctx, types.Plan{
					Name:   "alt-good",
					Info:   "new text here",
					Height: 543210000,
				})
			},
			expPass: true,
		},
		{
			name: "successful ICP overwrite",
			plan: types.Plan{
				Name:                "all-good",
				Info:                "some text here",
				Height:              123450000,
				UpgradedClientState: cs,
			},
			setup: func() {
				s.app.UpgradeKeeper.ScheduleUpgrade(s.ctx, types.Plan{
					Name:                "alt-good",
					Info:                "new text here",
					Height:              543210000,
					UpgradedClientState: altCs,
				})
			},
			expPass: true,
		},
		{
			name: "successful ICP overwrite with non ICP plan",
			plan: types.Plan{
				Name:   "all-good",
				Info:   "some text here",
				Height: 123450000,
			},
			setup: func() {
				s.app.UpgradeKeeper.ScheduleUpgrade(s.ctx, types.Plan{
					Name:                "alt-good",
					Info:                "new text here",
					Height:              543210000,
					UpgradedClientState: altCs,
				})
			},
			expPass: true,
		},
		{
			name: "unsuccessful schedule: invalid plan",
			plan: types.Plan{
				Height: 123450000,
			},
			setup:   func() {},
			expPass: false,
		},
		{
			name: "unsuccessful time schedule: due date in past",
			plan: types.Plan{
				Name: "all-good",
				Info: "some text here",
				Time: s.ctx.BlockTime(),
			},
			setup:   func() {},
			expPass: false,
		},
		{
			name: "unsuccessful height schedule: due date in past",
			plan: types.Plan{
				Name:   "all-good",
				Info:   "some text here",
				Height: 1,
			},
			setup:   func() {},
			expPass: false,
		},
		{
			name: "unsuccessful schedule: schedule already executed",
			plan: types.Plan{
				Name:   "all-good",
				Info:   "some text here",
				Height: 123450000,
			},
			setup: func() {
				s.app.UpgradeKeeper.SetUpgradeHandler("all-good", func(_ sdk.Context, _ types.Plan) {})
				s.app.UpgradeKeeper.ApplyUpgrade(s.ctx, types.Plan{
					Name:   "all-good",
					Info:   "some text here",
					Height: 123450000,
				})
			},
			expPass: false,
		},
		{
			name: "unsuccessful ICP schedule: UpgradedClientState is not valid client state",
			plan: types.Plan{
				Name:                "all-good",
				Info:                "some text here",
				Height:              123450000,
				UpgradedClientState: consAny,
			},
			setup:   func() {},
			expPass: false,
		},
	}

	for _, tc := range cases {
		tc := tc

		s.Run(tc.name, func() {
			// reset suite
			s.SetupTest()

			// setup test case
			tc.setup()

			err := s.app.UpgradeKeeper.ScheduleUpgrade(s.ctx, tc.plan)

			if tc.expPass {
				s.Require().NoError(err, "valid test case failed")
				if tc.plan.UpgradedClientState != nil {
					got, err := s.app.UpgradeKeeper.GetUpgradedClient(s.ctx, tc.plan.Height)
					s.Require().NoError(err)
					s.Require().Equal(clientState, got, "upgradedClient not equal to expected value")
				} else {
					// check that upgraded client is empty if latest plan does not specify an upgraded client
					got, err := s.app.UpgradeKeeper.GetUpgradedClient(s.ctx, tc.plan.Height)
					s.Require().Error(err)
					s.Require().Nil(got)
				}
			} else {
				s.Require().Error(err, "invalid test case passed")
			}
		})
	}
}

func (s *KeeperTestSuite) TestSetUpgradedClient() {
	var (
		clientState icpexported.ClientState
	)
	cases := []struct {
		name   string
		height int64
		setup  func()
		exists bool
	}{
		{
			name:   "no upgraded client exists",
			height: 10,
			setup:  func() {},
			exists: false,
		},
		{
			name:   "success",
			height: 10,
			setup: func() {
				clientState = &icptmtypes.ClientState{ChainId: "creatachain"}
				s.app.UpgradeKeeper.SetUpgradedClient(s.ctx, 10, clientState)
			},
			exists: true,
		},
	}

	for _, tc := range cases {
		// reset suite
		s.SetupTest()

		// setup test case
		tc.setup()

		gotCs, err := s.app.UpgradeKeeper.GetUpgradedClient(s.ctx, tc.height)
		if tc.exists {
			s.Require().Equal(clientState, gotCs, "valid case: %s did not retrieve correct client state", tc.name)
			s.Require().NoError(err, "valid case: %s returned error")
		} else {
			s.Require().Nil(gotCs, "invalid case: %s retrieved valid client state", tc.name)
			s.Require().Error(err, "invalid case: %s did not return error", tc.name)
		}
	}

}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
