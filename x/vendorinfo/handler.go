package vendorinfo

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateVendorInfoType:
			res, err := msgServer.CreateVendorInfoType(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUpdateVendorInfoType:
			res, err := msgServer.UpdateVendorInfoType(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgDeleteVendorInfoType:
			res, err := msgServer.DeleteVendorInfoType(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
