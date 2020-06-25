// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package requests

import (
	"github.com/AzureAD/microsoft-authentication-library-for-go/src/internal/msalbase"
	"github.com/AzureAD/microsoft-authentication-library-for-go/src/internal/wstrust"
	"github.com/stretchr/testify/mock"
)

type MockWebRequestManager struct {
	mock.Mock
}

func (mock *MockWebRequestManager) GetUserRealm(authParameters *msalbase.AuthParametersInternal) (*msalbase.UserRealm, error) {
	args := mock.Called(authParameters)
	return args.Get(0).(*msalbase.UserRealm), args.Error(1)
}

func (mock *MockWebRequestManager) GetMex(federationMetadataURL string) (*wstrust.WsTrustMexDocument, error) {
	args := mock.Called(federationMetadataURL)
	return args.Get(0).(*wstrust.WsTrustMexDocument), args.Error(1)
}

func (mock *MockWebRequestManager) GetWsTrustResponse(authParameters *msalbase.AuthParametersInternal,
	cloudAudienceURN string,
	endpoint *wstrust.WsTrustEndpoint) (*wstrust.WsTrustResponse, error) {
	args := mock.Called(authParameters, cloudAudienceURN, endpoint)
	return args.Get(0).(*wstrust.WsTrustResponse), args.Error(1)
}

func (mock *MockWebRequestManager) GetAccessTokenFromSamlGrant(authParameters *msalbase.AuthParametersInternal,
	samlGrant *wstrust.SamlTokenInfo) (*msalbase.TokenResponse, error) {
	args := mock.Called(authParameters, samlGrant)
	return args.Get(0).(*msalbase.TokenResponse), args.Error(1)
}

func (mock *MockWebRequestManager) GetAccessTokenFromUsernamePassword(authParameters *msalbase.AuthParametersInternal) (*msalbase.TokenResponse, error) {
	args := mock.Called(authParameters)
	return args.Get(0).(*msalbase.TokenResponse), args.Error(1)
}

func (mock *MockWebRequestManager) GetAccessTokenFromAuthCode(authParameters *msalbase.AuthParametersInternal,
	authCode string,
	codeVerifier string) (*msalbase.TokenResponse, error) {
	args := mock.Called(authParameters, authCode, codeVerifier)
	return args.Get(0).(*msalbase.TokenResponse), args.Error(1)
}

func (mock *MockWebRequestManager) GetAccessTokenFromRefreshToken(authParameters *msalbase.AuthParametersInternal,
	refreshToken string) (*msalbase.TokenResponse, error) {
	args := mock.Called(authParameters, refreshToken)
	return args.Get(0).(*msalbase.TokenResponse), args.Error(1)
}

func (mock *MockWebRequestManager) GetAccessTokenWithCertificate(authParameters *msalbase.AuthParametersInternal,
	certificate *msalbase.ClientCertificate) (*msalbase.TokenResponse, error) {
	args := mock.Called(authParameters, certificate)
	return args.Get(0).(*msalbase.TokenResponse), args.Error(1)
}

func (mock *MockWebRequestManager) GetDeviceCodeResult(authParameters *msalbase.AuthParametersInternal) (*msalbase.DeviceCodeResult, error) {
	args := mock.Called(authParameters)
	return args.Get(0).(*msalbase.DeviceCodeResult), args.Error(1)
}

func (mock *MockWebRequestManager) GetAccessTokenFromDeviceCodeResult(authParameters *msalbase.AuthParametersInternal,
	deviceCodeResult *msalbase.DeviceCodeResult) (*msalbase.TokenResponse, error) {
	args := mock.Called(authParameters, deviceCodeResult)
	return args.Get(0).(*msalbase.TokenResponse), args.Error(1)
}

func (mock *MockWebRequestManager) GetTenantDiscoveryResponse(openIDConfigurationEndpoint string) (*TenantDiscoveryResponse, error) {
	args := mock.Called(openIDConfigurationEndpoint)
	return args.Get(0).(*TenantDiscoveryResponse), args.Error(1)
}

func (mock *MockWebRequestManager) GetAadinstanceDiscoveryResponse(authorityInfo *msalbase.AuthorityInfo) (*InstanceDiscoveryResponse, error) {
	args := mock.Called(authorityInfo)
	return args.Get(0).(*InstanceDiscoveryResponse), args.Error(1)
}

func (mock *MockWebRequestManager) GetProviderConfigurationInformation(authParameters *msalbase.AuthParametersInternal) (*ProviderConfigurationInformation, error) {
	args := mock.Called(authParameters)
	return args.Get(0).(*ProviderConfigurationInformation), args.Error(1)
}