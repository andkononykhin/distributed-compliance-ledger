package pki

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

type DclauthKeeperMock struct {
	mock.Mock
}

func (m *DclauthKeeperMock) HasRole(
	ctx sdk.Context,
	addr sdk.AccAddress,
	roleToCheck dclauthtypes.AccountRole,
) bool {
	args := m.Called(ctx, addr, roleToCheck)
	return args.Bool(0)
}

var _ types.DclauthKeeper = &DclauthKeeperMock{}

type TestSetup struct {
	T *testing.T
	// Cdc         *amino.Codec
	Ctx           sdk.Context
	Wctx          context.Context
	Keeper        *keeper.Keeper
	DclauthKeeper *DclauthKeeperMock
	Handler       sdk.Handler
	// Querier     sdk.Querier
	Trustee sdk.AccAddress
}

func (setup *TestSetup) AddAccount(
	accAddress sdk.AccAddress,
	roles []dclauthtypes.AccountRole,
) {
	dclauthKeeper := setup.DclauthKeeper

	for _, role := range roles {
		dclauthKeeper.On("HasRole", mock.Anything, accAddress, role).Return(true)
	}
	dclauthKeeper.On("HasRole", mock.Anything, accAddress, mock.Anything).Return(false)
}

func GenerateAccAddress() sdk.AccAddress {
	_, _, accAddress := testdata.KeyTestPubAddr()
	return accAddress
}

func Setup(t *testing.T) *TestSetup {
	dclauthKeeper := &DclauthKeeperMock{}
	// FIXME: Add dclauthKeeper parameter to testkeeper.PkiKeeper
	keeper, ctx := testkeeper.PkiKeeper(t /*, dclauthKeeper*/)

	trustee := GenerateAccAddress()

	setup := &TestSetup{
		T:             t,
		Ctx:           ctx,
		Wctx:          sdk.WrapSDKContext(ctx),
		Keeper:        keeper,
		DclauthKeeper: dclauthKeeper,
		Handler:       NewHandler(*keeper),
		Trustee:       trustee,
	}

	setup.AddAccount(trustee, []dclauthtypes.AccountRole{dclauthtypes.Vendor})

	return setup
}

func queryProposedCertificate(
	setup *TestSetup,
	subject string,
	subjectKeyId string,
) (*types.ProposedCertificate, error) {

	// query proposed certificate
	req := &types.QueryGetProposedCertificateRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
	}

	resp, err := setup.Keeper.ProposedCertificate(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)
		return nil, err
	}

	require.NotNil(setup.T, resp)
	require.NotNil(setup.T, resp.ProposedCertificate)
	return resp.ProposedCertificate, nil
}

func queryAllApprovedCertificates(setup *TestSetup) ([]types.ApprovedCertificates, error) {

	// query all certificates
	req := &types.QueryAllApprovedCertificatesRequest{}

	resp, err := setup.Keeper.ApprovedCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)
		return nil, err
	}

	require.NotNil(setup.T, resp)
	return resp.ApprovedCertificates, nil
}

func querySingleApprovedCertificate(
	setup *TestSetup,
	subject string,
	subjectKeyId string,
) (*types.Certificate, error) {

	certificates, err := queryApprovedCertificates(setup, subject, subjectKeyId)
	if err != nil {
		return nil, err
	}

	if len(certificates.Certs) > 1 {
		require.Fail(setup.T, "More than 1 certificate returned")
	}

	return certificates.Certs[0], nil
}

func queryApprovedCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyId string,
) (*types.ApprovedCertificates, error) {

	// query certificate
	req := &types.QueryGetApprovedCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
	}

	resp, err := setup.Keeper.ApprovedCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)
		return nil, err
	}

	require.NotNil(setup.T, resp)
	require.NotNil(setup.T, resp.ApprovedCertificates)
	return resp.ApprovedCertificates, nil
}

func queryAllProposedCertificateRevocations(setup *TestSetup) ([]types.ProposedCertificateRevocation, error) {

	// query all proposed certificate revocations
	req := &types.QueryAllProposedCertificateRevocationRequest{}

	resp, err := setup.Keeper.ProposedCertificateRevocationAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)
		return nil, err
	}

	require.NotNil(setup.T, resp)
	return resp.ProposedCertificateRevocation, nil
}

func queryProposedCertificateRevocation(
	setup *TestSetup,
	subject string,
	subjectKeyId string,
) (*types.ProposedCertificateRevocation, error) {

	// query proposed certificate revocation
	req := &types.QueryGetProposedCertificateRevocationRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
	}

	resp, err := setup.Keeper.ProposedCertificateRevocation(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)
		return nil, err
	}

	require.NotNil(setup.T, resp)
	require.NotNil(setup.T, resp.ProposedCertificateRevocation)
	return resp.ProposedCertificateRevocation, nil
}

func queryAllRevokedCertificates(setup *TestSetup) ([]types.RevokedCertificates, error) {

	// query all revoked certificates
	req := &types.QueryAllRevokedCertificatesRequest{}

	resp, err := setup.Keeper.RevokedCertificatesAll(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)
		return nil, err
	}

	require.NotNil(setup.T, resp)
	return resp.RevokedCertificates, nil
}

func querySingleRevokedCertificate(
	setup *TestSetup,
	subject string,
	subjectKeyId string,
) (*types.Certificate, error) {

	certificates, err := queryRevokedCertificates(setup, subject, subjectKeyId)
	if err != nil {
		return nil, err
	}

	if len(certificates.Certs) > 1 {
		require.Fail(setup.T, "More than 1 certificate returned")
	}

	return certificates.Certs[0], nil
}

func queryRevokedCertificates(
	setup *TestSetup,
	subject string,
	subjectKeyId string,
) (*types.RevokedCertificates, error) {

	// query revoked certificate
	req := &types.QueryGetRevokedCertificatesRequest{
		Subject:      subject,
		SubjectKeyId: subjectKeyId,
	}

	resp, err := setup.Keeper.RevokedCertificates(setup.Wctx, req)
	if err != nil {
		require.Nil(setup.T, resp)
		return nil, err
	}

	require.NotNil(setup.T, resp)
	require.NotNil(setup.T, resp.RevokedCertificates)
	return resp.RevokedCertificates, nil
}
