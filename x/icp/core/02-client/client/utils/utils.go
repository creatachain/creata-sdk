package utils

import (
	"context"

	tmtypes "github.com/creatachain/augusteum/types"

	"github.com/creatachain/creata-sdk/client"

	"github.com/creatachain/creata-sdk/codec"
	sdkerrors "github.com/creatachain/creata-sdk/types/errors"
	"github.com/creatachain/creata-sdk/x/icp/core/02-client/types"
	commitmenttypes "github.com/creatachain/creata-sdk/x/icp/core/23-commitment/types"
	host "github.com/creatachain/creata-sdk/x/icp/core/24-host"
	icpclient "github.com/creatachain/creata-sdk/x/icp/core/client"
	"github.com/creatachain/creata-sdk/x/icp/core/exported"
	icptmtypes "github.com/creatachain/creata-sdk/x/icp/light-clients/07-augusteum/types"
)

// QueryClientState returns a client state. If prove is true, it performs an MSM store query
// in order to retrieve the merkle proof. Otherwise, it uses the gRPC query client.
func QueryClientState(
	clientCtx client.Context, clientID string, prove bool,
) (*types.QueryClientStateResponse, error) {
	if prove {
		return QueryClientStateMSM(clientCtx, clientID)
	}

	queryClient := types.NewQueryClient(clientCtx)
	req := &types.QueryClientStateRequest{
		ClientId: clientID,
	}

	return queryClient.ClientState(context.Background(), req)
}

// QueryClientStateMSM queries the store to get the light client state and a merkle proof.
func QueryClientStateMSM(
	clientCtx client.Context, clientID string,
) (*types.QueryClientStateResponse, error) {
	key := host.FullClientStateKey(clientID)

	value, proofBz, proofHeight, err := icpclient.QueryAugusteumProof(clientCtx, key)
	if err != nil {
		return nil, err
	}

	// check if client exists
	if len(value) == 0 {
		return nil, sdkerrors.Wrap(types.ErrClientNotFound, clientID)
	}

	cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)

	clientState, err := types.UnmarshalClientState(cdc, value)
	if err != nil {
		return nil, err
	}

	anyClientState, err := types.PackClientState(clientState)
	if err != nil {
		return nil, err
	}

	clientStateRes := types.NewQueryClientStateResponse(anyClientState, proofBz, proofHeight)
	return clientStateRes, nil
}

// QueryConsensusState returns a consensus state. If prove is true, it performs an MSM store
// query in order to retrieve the merkle proof. Otherwise, it uses the gRPC query client.
func QueryConsensusState(
	clientCtx client.Context, clientID string, height exported.Height, prove, latestHeight bool,
) (*types.QueryConsensusStateResponse, error) {
	if prove {
		return QueryConsensusStateMSM(clientCtx, clientID, height)
	}

	queryClient := types.NewQueryClient(clientCtx)
	req := &types.QueryConsensusStateRequest{
		ClientId:       clientID,
		RevisionNumber: height.GetRevisionNumber(),
		RevisionHeight: height.GetRevisionHeight(),
		LatestHeight:   latestHeight,
	}

	return queryClient.ConsensusState(context.Background(), req)
}

// QueryConsensusStateMSM queries the store to get the consensus state of a light client and a
// merkle proof of its existence or non-existence.
func QueryConsensusStateMSM(
	clientCtx client.Context, clientID string, height exported.Height,
) (*types.QueryConsensusStateResponse, error) {
	key := host.FullConsensusStateKey(clientID, height)

	value, proofBz, proofHeight, err := icpclient.QueryAugusteumProof(clientCtx, key)
	if err != nil {
		return nil, err
	}

	// check if consensus state exists
	if len(value) == 0 {
		return nil, sdkerrors.Wrap(types.ErrConsensusStateNotFound, clientID)
	}

	cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)

	cs, err := types.UnmarshalConsensusState(cdc, value)
	if err != nil {
		return nil, err
	}

	anyConsensusState, err := types.PackConsensusState(cs)
	if err != nil {
		return nil, err
	}

	return types.NewQueryConsensusStateResponse(anyConsensusState, proofBz, proofHeight), nil
}

// QueryAugusteumHeader takes a client context and returns the appropriate
// augusteum header
func QueryAugusteumHeader(clientCtx client.Context) (icptmtypes.Header, int64, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return icptmtypes.Header{}, 0, err
	}

	info, err := node.MSMInfo(context.Background())
	if err != nil {
		return icptmtypes.Header{}, 0, err
	}

	height := info.Response.LastBlockHeight

	commit, err := node.Commit(context.Background(), &height)
	if err != nil {
		return icptmtypes.Header{}, 0, err
	}

	page := 0
	count := 10_000

	validators, err := node.Validators(context.Background(), &height, &page, &count)
	if err != nil {
		return icptmtypes.Header{}, 0, err
	}

	protoCommit := commit.SignedHeader.ToProto()
	protoValset, err := tmtypes.NewValidatorSet(validators.Validators).ToProto()
	if err != nil {
		return icptmtypes.Header{}, 0, err
	}

	header := icptmtypes.Header{
		SignedHeader: protoCommit,
		ValidatorSet: protoValset,
	}

	return header, height, nil
}

// QueryNodeConsensusState takes a client context and returns the appropriate
// augusteum consensus state
func QueryNodeConsensusState(clientCtx client.Context) (*icptmtypes.ConsensusState, int64, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return &icptmtypes.ConsensusState{}, 0, err
	}

	info, err := node.MSMInfo(context.Background())
	if err != nil {
		return &icptmtypes.ConsensusState{}, 0, err
	}

	height := info.Response.LastBlockHeight

	commit, err := node.Commit(context.Background(), &height)
	if err != nil {
		return &icptmtypes.ConsensusState{}, 0, err
	}

	page := 1
	count := 10_000

	nextHeight := height + 1
	nextVals, err := node.Validators(context.Background(), &nextHeight, &page, &count)
	if err != nil {
		return &icptmtypes.ConsensusState{}, 0, err
	}

	state := &icptmtypes.ConsensusState{
		Timestamp:          commit.Time,
		Root:               commitmenttypes.NewMerkleRoot(commit.AppHash),
		NextValidatorsHash: tmtypes.NewValidatorSet(nextVals.Validators).Hash(),
	}

	return state, height, nil
}
