package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

// SetRevokedRootCertificates set revokedRootCertificates in the store
func (k Keeper) SetRevokedRootCertificates(ctx sdk.Context, revokedRootCertificates types.RevokedRootCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedRootCertificatesKey))
	b := k.cdc.MustMarshal(&revokedRootCertificates)
	store.Set([]byte{0}, b)
}

// GetRevokedRootCertificates returns revokedRootCertificates
func (k Keeper) GetRevokedRootCertificates(ctx sdk.Context) (val types.RevokedRootCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedRootCertificatesKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRevokedRootCertificates removes revokedRootCertificates from the store
func (k Keeper) RemoveRevokedRootCertificates(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedRootCertificatesKey))
	store.Delete([]byte{0})
}

// Add revoked root certificate to the list
func (k Keeper) AddRevokedRootCertificate(ctx sdk.Context, certId types.CertificateIdentifier) {
	rootCertificates, _ := k.GetRevokedRootCertificates(ctx)

	// Check if the root cert is already there
	for _, existingСertId := range rootCertificates.Certs {
		if *existingСertId == certId {
			return
		}
	}

	rootCertificates.Certs = append(rootCertificates.Certs, &certId)

	k.SetRevokedRootCertificates(ctx, rootCertificates)
}

// Remove revoked root certificate from the list
func (k Keeper) RemoveRevokedRootCertificate(ctx sdk.Context, certId types.CertificateIdentifier) {
	rootCertificates, _ := k.GetRevokedRootCertificates(ctx)

	certIDIndex := -1
	for i, existingIdentifier := range rootCertificates.Certs {
		if *existingIdentifier == certId {
			certIDIndex = i
			break
		}
	}
	if certIDIndex == -1 {
		return
	}

	rootCertificates.Certs =
		append(rootCertificates.Certs[:certIDIndex], rootCertificates.Certs[certIDIndex+1:]...)

	k.SetRevokedRootCertificates(ctx, rootCertificates)
}