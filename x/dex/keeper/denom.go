package keeper

import (
	"orderbook-interchange/x/dex/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
)

func (k Keeper) SaveVoucherDenom(ctx sdk.Context, port string, channel string, denom string) {
	voucher := VoucherDenom(port, channel, denom)

	_, saved := k.GetDenomTrace(ctx, voucher)
	if !saved {
		k.SetDenomTrace(ctx, types.DenomTrace{
			Index:   voucher,
			Port:    port,
			Channel: channel,
			Origin:  denom,
		})
	}
}

func VoucherDenom(port string, channel string, denom string) string {
	sourcePrefix := ibctransfertypes.GetDenomPrefix(port, channel)
	prefixedDenom := sourcePrefix + denom
	denomTrace := ibctransfertypes.ParseDenomTrace(prefixedDenom)
	voucher := denomTrace.IBCDenom()
	return voucher[:16]
}

func (k Keeper) OrigialDenom(ctx sdk.Context, port string, channel string, voucher string) (string, bool) {
	trace, exist := k.GetDenomTrace(ctx, voucher)
	if exist {
		if trace.Port == port && trace.Channel == channel {
			return trace.Origin, true
		}
	}
	return "", false
}
