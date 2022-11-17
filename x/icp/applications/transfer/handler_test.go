package transfer_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/creatachain/creata-sdk/types"
	"github.com/creatachain/creata-sdk/x/icp/applications/transfer/types"
	clienttypes "github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	channeltypes "github.com/creatachain/creata-sdk/x/icp/core/04-channel/types"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
	icptesting "github.com/creatachain/creata-sdk/x/icp/testing"
)

type TransferTestSuite struct {
	suite.Suite

	coordinator *icptesting.Coordinator

	// testing chains used for convenience and readability
	chainA *icptesting.TestChain
	chainB *icptesting.TestChain
	chainC *icptesting.TestChain
}

func (suite *TransferTestSuite) SetupTest() {
	suite.coordinator = icptesting.NewCoordinator(suite.T(), 3)
	suite.chainA = suite.coordinator.GetChain(icptesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(icptesting.GetChainID(1))
	suite.chainC = suite.coordinator.GetChain(icptesting.GetChainID(2))
}

// constructs a send from chainA to chainB on the established channel/connection
// and sends the same coin back from chainB to chainA.
func (suite *TransferTestSuite) TestHandleMsgTransfer() {
	// setup between chainA and chainB
	clientA, clientB, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Augusteum)
	channelA, channelB := suite.coordinator.CreateTransferChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
	//	originalBalance := suite.chainA.App.BankKeeper.GetBalance(suite.chainA.GetContext(), suite.chainA.SenderAccount.GetAddress(), sdk.DefaultBondDenom)
	timeoutHeight := clienttypes.NewHeight(0, 110)

	coinToSendToB := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100))

	// send from chainA to chainB
	msg := types.NewMsgTransfer(channelA.PortID, channelA.ID, coinToSendToB, suite.chainA.SenderAccount.GetAddress(), suite.chainB.SenderAccount.GetAddress().String(), timeoutHeight, 0)

	err := suite.coordinator.SendMsg(suite.chainA, suite.chainB, clientB, msg)
	suite.Require().NoError(err) // message committed

	// relay send
	fungibleTokenPacket := types.NewFungibleTokenPacketData(coinToSendToB.Denom, coinToSendToB.Amount.Uint64(), suite.chainA.SenderAccount.GetAddress().String(), suite.chainB.SenderAccount.GetAddress().String())
	packet := channeltypes.NewPacket(fungibleTokenPacket.GetBytes(), 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, timeoutHeight, 0)
	ack := channeltypes.NewResultAcknowledgement([]byte{byte(1)})
	err = suite.coordinator.RelayPacket(suite.chainA, suite.chainB, clientA, clientB, packet, ack.GetBytes())
	suite.Require().NoError(err) // relay committed

	// check that voucher exists on chain B
	voucherDenomTrace := types.ParseDenomTrace(types.GetPrefixedDenom(packet.GetDestPort(), packet.GetDestChannel(), sdk.DefaultBondDenom))
	balance := suite.chainB.App.BankKeeper.GetBalance(suite.chainB.GetContext(), suite.chainB.SenderAccount.GetAddress(), voucherDenomTrace.ICPDenom())

	coinSentFromAToB := types.GetTransferCoin(channelB.PortID, channelB.ID, sdk.DefaultBondDenom, 100)
	suite.Require().Equal(coinSentFromAToB, balance)

	// setup between chainB to chainC
	clientOnBForC, clientOnCForB, connOnBForC, connOnCForB := suite.coordinator.SetupClientConnections(suite.chainB, suite.chainC, exported.Augusteum)
	channelOnBForC, channelOnCForB := suite.coordinator.CreateTransferChannels(suite.chainB, suite.chainC, connOnBForC, connOnCForB, channeltypes.UNORDERED)

	// send from chainB to chainC
	msg = types.NewMsgTransfer(channelOnBForC.PortID, channelOnBForC.ID, coinSentFromAToB, suite.chainB.SenderAccount.GetAddress(), suite.chainC.SenderAccount.GetAddress().String(), timeoutHeight, 0)

	err = suite.coordinator.SendMsg(suite.chainB, suite.chainC, clientOnCForB, msg)
	suite.Require().NoError(err) // message committed

	// relay send
	// NOTE: fungible token is prefixed with the full trace in order to verify the packet commitment
	fullDenomPath := types.GetPrefixedDenom(channelOnCForB.PortID, channelOnCForB.ID, voucherDenomTrace.GetFullDenomPath())
	fungibleTokenPacket = types.NewFungibleTokenPacketData(voucherDenomTrace.GetFullDenomPath(), coinSentFromAToB.Amount.Uint64(), suite.chainB.SenderAccount.GetAddress().String(), suite.chainC.SenderAccount.GetAddress().String())
	packet = channeltypes.NewPacket(fungibleTokenPacket.GetBytes(), 1, channelOnBForC.PortID, channelOnBForC.ID, channelOnCForB.PortID, channelOnCForB.ID, timeoutHeight, 0)
	err = suite.coordinator.RelayPacket(suite.chainB, suite.chainC, clientOnBForC, clientOnCForB, packet, ack.GetBytes())
	suite.Require().NoError(err) // relay committed

	coinSentFromBToC := sdk.NewInt64Coin(types.ParseDenomTrace(fullDenomPath).ICPDenom(), 100)
	balance = suite.chainC.App.BankKeeper.GetBalance(suite.chainC.GetContext(), suite.chainC.SenderAccount.GetAddress(), coinSentFromBToC.Denom)

	// check that the balance is updated on chainC
	suite.Require().Equal(coinSentFromBToC, balance)

	// check that balance on chain B is empty
	balance = suite.chainB.App.BankKeeper.GetBalance(suite.chainB.GetContext(), suite.chainB.SenderAccount.GetAddress(), coinSentFromBToC.Denom)
	suite.Require().Zero(balance.Amount.Int64())

	// send from chainC back to chainB
	msg = types.NewMsgTransfer(channelOnCForB.PortID, channelOnCForB.ID, coinSentFromBToC, suite.chainC.SenderAccount.GetAddress(), suite.chainB.SenderAccount.GetAddress().String(), timeoutHeight, 0)

	err = suite.coordinator.SendMsg(suite.chainC, suite.chainB, clientOnBForC, msg)
	suite.Require().NoError(err) // message committed

	// relay send
	// NOTE: fungible token is prefixed with the full trace in order to verify the packet commitment
	fungibleTokenPacket = types.NewFungibleTokenPacketData(fullDenomPath, coinSentFromBToC.Amount.Uint64(), suite.chainC.SenderAccount.GetAddress().String(), suite.chainB.SenderAccount.GetAddress().String())
	packet = channeltypes.NewPacket(fungibleTokenPacket.GetBytes(), 1, channelOnCForB.PortID, channelOnCForB.ID, channelOnBForC.PortID, channelOnBForC.ID, timeoutHeight, 0)
	err = suite.coordinator.RelayPacket(suite.chainC, suite.chainB, clientOnCForB, clientOnBForC, packet, ack.GetBytes())
	suite.Require().NoError(err) // relay committed

	balance = suite.chainB.App.BankKeeper.GetBalance(suite.chainB.GetContext(), suite.chainB.SenderAccount.GetAddress(), coinSentFromAToB.Denom)

	// check that the balance on chainA returned back to the original state
	suite.Require().Equal(coinSentFromAToB, balance)

	// check that module account escrow address is empty
	escrowAddress := types.GetEscrowAddress(packet.GetDestPort(), packet.GetDestChannel())
	balance = suite.chainB.App.BankKeeper.GetBalance(suite.chainB.GetContext(), escrowAddress, sdk.DefaultBondDenom)
	suite.Require().Equal(sdk.NewCoin(sdk.DefaultBondDenom, sdk.ZeroInt()), balance)

	// check that balance on chain B is empty
	balance = suite.chainC.App.BankKeeper.GetBalance(suite.chainC.GetContext(), suite.chainC.SenderAccount.GetAddress(), voucherDenomTrace.ICPDenom())
	suite.Require().Zero(balance.Amount.Int64())
}

func TestTransferTestSuite(t *testing.T) {
	suite.Run(t, new(TransferTestSuite))
}
