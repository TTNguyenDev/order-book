package keeper

import (
	"context"
	"errors"
	"orderbook-interchange/x/dex/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CancelSellOrder(goCtx context.Context, msg *types.MsgCancelSellOrder) (*types.MsgCancelSellOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairIndex := types.OrderBookIndex(msg.Port, msg.Channel, msg.AmountDenom, msg.PriceDenom)
	s, found := k.GetSellOrderBook(ctx, pairIndex)
	if !found {
		return &types.MsgCancelSellOrderResponse{}, errors.New("the pair doesn't exist")
	}

	order, err := s.Book.GetOrderFromID(msg.OrderId)
	if err != nil {
		return &types.MsgCancelSellOrderResponse{}, errors.New("canceller must be creator")
	}

	if err := s.Book.RemoveOrderFromID(msg.OrderId); err != nil {
		return &types.MsgCancelSellOrderResponse{}, err
	}

	k.SetSellOrderBook(ctx, s)

	seller, err := sdk.AccAddressFromBech32(order.creator)
	if err != nil {
		return &types.MsgCancelSellOrderResponse{}, err
	}

	if err := k.SafeMint(ctx, msg.Port, msg.Channel, seller, msg.AmountDenom, order.Amount); err != nil {
		return &types.MsgCancelSellOrderResponse{}, err
	}

	return &types.MsgCancelSellOrderResponse{}, nil
}
