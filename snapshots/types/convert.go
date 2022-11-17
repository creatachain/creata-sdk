package types

import (
	msm "github.com/creatachain/augusteum/msm/types"
	proto "github.com/gogo/protobuf/proto"

	sdkerrors "github.com/creatachain/creata-sdk/types/errors"
)

// Converts an MSM snapshot to a snapshot. Mainly to decode the SDK metadata.
func SnapshotFromMSM(in *msm.Snapshot) (Snapshot, error) {
	snapshot := Snapshot{
		Height: in.Height,
		Format: in.Format,
		Chunks: in.Chunks,
		Hash:   in.Hash,
	}
	err := proto.Unmarshal(in.Metadata, &snapshot.Metadata)
	if err != nil {
		return Snapshot{}, sdkerrors.Wrap(err, "failed to unmarshal snapshot metadata")
	}
	return snapshot, nil
}

// Converts a Snapshot to its MSM representation. Mainly to encode the SDK metadata.
func (s Snapshot) ToMSM() (msm.Snapshot, error) {
	out := msm.Snapshot{
		Height: s.Height,
		Format: s.Format,
		Chunks: s.Chunks,
		Hash:   s.Hash,
	}
	var err error
	out.Metadata, err = proto.Marshal(&s.Metadata)
	if err != nil {
		return msm.Snapshot{}, sdkerrors.Wrap(err, "failed to marshal snapshot metadata")
	}
	return out, nil
}
