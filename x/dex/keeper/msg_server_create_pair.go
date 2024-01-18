package keeper

import (
	"context"
	"errors"
	"orderbook-interchange/x/dex/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
)

func (k msgServer) SendCreatePair(goCtx context.Context, msg *types.MsgSendCreatePair) (*types.MsgSendCreatePairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairIndex := types.OrderBookIndex(msg.Port, msg.ChannelID, msg.SourceDenom, msg.TargetDenom)

	_, found := k.GetSellOrderBook(ctx, pairIndex)
	if found {
		return &types.MsgSendCreatePairResponse{}, errors.New("the pair already exist")
	}

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.CreatePairPacketData

	packet.SourceDenom = msg.SourceDenom
	packet.TargetDenom = msg.TargetDenom

	// Transmit the packet
	_, err := k.TransmitCreatePairPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendCreatePairResponse{}, nil
}
