package rest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/keeper"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliancetest/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"net/http"
)

func getTestingResultHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()

		vid, err_ := conversions.ParseVID(vars[vid])
		if err_ != nil {
			restCtx.WriteErrorResponse( http.StatusBadRequest, err_.Error())
			return
		}

		pid, err_ := conversions.ParsePID(vars[pid])
		if err_ != nil {
			restCtx.WriteErrorResponse( http.StatusBadRequest, err_.Error())
			return
		}

		res, height, err := restCtx.QueryStore(keeper.TestingResultId(vid, pid), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound, types.ErrTestingResultDoesNotExist(vid, pid).Error())
			return
		}

		var testingResult types.TestingResults
		cliCtx.Codec.MustUnmarshalBinaryBare(res, &testingResult)

		restCtx.EncodeAndRespondWithHeight(testingResult, height)
	}
}
