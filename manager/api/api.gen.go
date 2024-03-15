// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// Defines values for ChargeStationInstallCertificatesCertificatesStatus.
const (
	Accepted ChargeStationInstallCertificatesCertificatesStatus = "Accepted"
	Pending  ChargeStationInstallCertificatesCertificatesStatus = "Pending"
	Rejected ChargeStationInstallCertificatesCertificatesStatus = "Rejected"
)

// Defines values for ChargeStationInstallCertificatesCertificatesType.
const (
	CSMS ChargeStationInstallCertificatesCertificatesType = "CSMS"
	MF   ChargeStationInstallCertificatesCertificatesType = "MF"
	MO   ChargeStationInstallCertificatesCertificatesType = "MO"
	V2G  ChargeStationInstallCertificatesCertificatesType = "V2G"
)

// Defines values for ChargeStationTriggerTrigger.
const (
	BootNotification               ChargeStationTriggerTrigger = "BootNotification"
	SignChargingStationCertificate ChargeStationTriggerTrigger = "SignChargingStationCertificate"
	SignCombinedCertificate        ChargeStationTriggerTrigger = "SignCombinedCertificate"
	SignV2GCertificate             ChargeStationTriggerTrigger = "SignV2GCertificate"
	StatusNotification             ChargeStationTriggerTrigger = "StatusNotification"
)

// Defines values for ConnectorFormat.
const (
	CABLE  ConnectorFormat = "CABLE"
	SOCKET ConnectorFormat = "SOCKET"
)

// Defines values for ConnectorPowerType.
const (
	AC1PHASE ConnectorPowerType = "AC_1_PHASE"
	AC3PHASE ConnectorPowerType = "AC_3_PHASE"
	DC       ConnectorPowerType = "DC"
)

// Defines values for ConnectorStandard.
const (
	CHADEMO            ConnectorStandard = "CHADEMO"
	CHAOJI             ConnectorStandard = "CHAOJI"
	DOMESTICA          ConnectorStandard = "DOMESTIC_A"
	DOMESTICB          ConnectorStandard = "DOMESTIC_B"
	DOMESTICC          ConnectorStandard = "DOMESTIC_C"
	DOMESTICD          ConnectorStandard = "DOMESTIC_D"
	DOMESTICE          ConnectorStandard = "DOMESTIC_E"
	DOMESTICF          ConnectorStandard = "DOMESTIC_F"
	DOMESTICG          ConnectorStandard = "DOMESTIC_G"
	DOMESTICH          ConnectorStandard = "DOMESTIC_H"
	DOMESTICI          ConnectorStandard = "DOMESTIC_I"
	DOMESTICJ          ConnectorStandard = "DOMESTIC_J"
	DOMESTICK          ConnectorStandard = "DOMESTIC_K"
	DOMESTICL          ConnectorStandard = "DOMESTIC_L"
	GBTAC              ConnectorStandard = "GBT_AC"
	GBTDC              ConnectorStandard = "GBT_DC"
	IEC603092Single16  ConnectorStandard = "IEC_60309_2_single_16"
	IEC603092Three16   ConnectorStandard = "IEC_60309_2_three_16"
	IEC603092Three32   ConnectorStandard = "IEC_60309_2_three_32"
	IEC603092Three64   ConnectorStandard = "IEC_60309_2_three_64"
	IEC62196T1         ConnectorStandard = "IEC_62196_T1"
	IEC62196T1COMBO    ConnectorStandard = "IEC_62196_T1_COMBO"
	IEC62196T2         ConnectorStandard = "IEC_62196_T2"
	IEC62196T2COMBO    ConnectorStandard = "IEC_62196_T2_COMBO"
	IEC62196T3A        ConnectorStandard = "IEC_62196_T3A"
	IEC62196T3C        ConnectorStandard = "IEC_62196_T3C"
	NEMA1030           ConnectorStandard = "NEMA_10_30"
	NEMA1050           ConnectorStandard = "NEMA_10_50"
	NEMA1430           ConnectorStandard = "NEMA_14_30"
	NEMA1450           ConnectorStandard = "NEMA_14_50"
	NEMA520            ConnectorStandard = "NEMA_5_20"
	NEMA630            ConnectorStandard = "NEMA_6_30"
	NEMA650            ConnectorStandard = "NEMA_6_50"
	PANTOGRAPHBOTTOMUP ConnectorStandard = "PANTOGRAPH_BOTTOM_UP"
	PANTOGRAPHTOPDOWN  ConnectorStandard = "PANTOGRAPH_TOP_DOWN"
	TESLAR             ConnectorStandard = "TESLA_R"
	TESLAS             ConnectorStandard = "TESLA_S"
	UNKNOWN            ConnectorStandard = "UNKNOWN"
)

// Defines values for LocationParkingType.
const (
	ALONGMOTORWAY     LocationParkingType = "ALONG_MOTORWAY"
	ONDRIVEWAY        LocationParkingType = "ON_DRIVEWAY"
	ONSTREET          LocationParkingType = "ON_STREET"
	PARKINGGARAGE     LocationParkingType = "PARKING_GARAGE"
	PARKINGLOT        LocationParkingType = "PARKING_LOT"
	UNDERGROUNDGARAGE LocationParkingType = "UNDERGROUND_GARAGE"
)

// Defines values for RegistrationStatus.
const (
	PENDING    RegistrationStatus = "PENDING"
	REGISTERED RegistrationStatus = "REGISTERED"
)

// Defines values for TokenCacheMode.
const (
	ALLOWED        TokenCacheMode = "ALLOWED"
	ALLOWEDOFFLINE TokenCacheMode = "ALLOWED_OFFLINE"
	ALWAYS         TokenCacheMode = "ALWAYS"
	NEVER          TokenCacheMode = "NEVER"
)

// Defines values for TokenType.
const (
	ADHOCUSER TokenType = "AD_HOC_USER"
	APPUSER   TokenType = "APP_USER"
	OTHER     TokenType = "OTHER"
	RFID      TokenType = "RFID"
)

// Certificate A client certificate
type Certificate struct {
	// Certificate The PEM encoded certificate with newlines replaced by `\n`
	Certificate string `json:"certificate"`
}

// ChargeStationAuth Connection details for a charge station
type ChargeStationAuth struct {
	// Base64SHA256Password The base64 encoded, SHA-256 hash of the charge station password
	Base64SHA256Password *string `json:"base64SHA256Password,omitempty"`

	// InvalidUsernameAllowed If set to true then an invalid username will not prevent the charge station connecting
	InvalidUsernameAllowed *bool `json:"invalidUsernameAllowed,omitempty"`

	// SecurityProfile The security profile to use for the charge station: * `0` - unsecured transport with basic auth * `1` - TLS with basic auth * `2` - TLS with client certificate
	SecurityProfile int `json:"securityProfile"`
}

// ChargeStationInstallCertificates The set of certificates to install on the charge station. The certificates will be sent
// to the charge station asynchronously.
type ChargeStationInstallCertificates struct {
	Certificates []struct {
		// Certificate The PEM encoded certificate with newlines replaced by `\n`
		Certificate string `json:"certificate"`

		// Status The status, defaults to Pending
		Status *ChargeStationInstallCertificatesCertificatesStatus `json:"status,omitempty"`
		Type   ChargeStationInstallCertificatesCertificatesType    `json:"type"`
	} `json:"certificates"`
}

// ChargeStationInstallCertificatesCertificatesStatus The status, defaults to Pending
type ChargeStationInstallCertificatesCertificatesStatus string

// ChargeStationInstallCertificatesCertificatesType defines model for ChargeStationInstallCertificates.Certificates.Type.
type ChargeStationInstallCertificatesCertificatesType string

// ChargeStationSettings Settings for a charge station
type ChargeStationSettings map[string]string

// ChargeStationTrigger Trigger a charge station action
type ChargeStationTrigger struct {
	Trigger ChargeStationTriggerTrigger `json:"trigger"`
}

// ChargeStationTriggerTrigger defines model for ChargeStationTrigger.Trigger.
type ChargeStationTriggerTrigger string

// Connector defines model for Connector.
type Connector struct {
	Format      ConnectorFormat    `json:"format"`
	Id          string             `json:"id"`
	MaxAmperage int32              `json:"max_amperage"`
	MaxVoltage  int32              `json:"max_voltage"`
	PowerType   ConnectorPowerType `json:"power_type"`
	Standard    ConnectorStandard  `json:"standard"`
}

// ConnectorFormat defines model for Connector.Format.
type ConnectorFormat string

// ConnectorPowerType defines model for Connector.PowerType.
type ConnectorPowerType string

// ConnectorStandard defines model for Connector.Standard.
type ConnectorStandard string

// Evse defines model for Evse.
type Evse struct {
	Connectors []Connector `json:"connectors"`
	EvseId     *string     `json:"evse_id"`

	// Uid Uniquely identifies the EVSE within the CPOs platform (and
	// suboperator platforms).
	Uid string `json:"uid"`
}

// GeoLocation defines model for GeoLocation.
type GeoLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Location A charge station location
type Location struct {
	Address     string               `json:"address"`
	City        string               `json:"city"`
	Coordinates GeoLocation          `json:"coordinates"`
	Country     string               `json:"country"`
	CountryCode string               `json:"country_code"`
	Evses       *[]Evse              `json:"evses"`
	Name        *string              `json:"name"`
	ParkingType *LocationParkingType `json:"parking_type"`
	PartyId     string               `json:"party_id"`
	PostalCode  *string              `json:"postal_code"`
}

// LocationParkingType defines model for Location.ParkingType.
type LocationParkingType string

// Registration Defines the initial connection details for the OCPI registration process
type Registration struct {
	// Status The status of the registration request. If the request is marked as `REGISTERED` then the token will be allowed to
	// be used to access all endpoints avoiding the need for the OCPI registration process. If the request is marked as
	// `PENDING` then the token will only be allowed to access the `/ocpi/versions`, `/ocpi/2.2` and `/ocpi/2.2/credentials`
	// endpoints.
	Status *RegistrationStatus `json:"status,omitempty"`

	// Token The token to use for communicating with the eMSP (CREDENTIALS_TOKEN_A).
	Token string `json:"token"`

	// Url The URL of the eMSP versions endpoint. If provided the CSMS will act as the sender of the versions request.
	Url *string `json:"url,omitempty"`
}

// RegistrationStatus The status of the registration request. If the request is marked as `REGISTERED` then the token will be allowed to
// be used to access all endpoints avoiding the need for the OCPI registration process. If the request is marked as
// `PENDING` then the token will only be allowed to access the `/ocpi/versions`, `/ocpi/2.2` and `/ocpi/2.2/credentials`
// endpoints.
type RegistrationStatus string

// Status HTTP status
type Status struct {
	// Error The error details
	Error *string `json:"error,omitempty"`

	// Status The status description
	Status string `json:"status"`
}

// Token An authorization token
type Token struct {
	// CacheMode Indicates what type of token caching is allowed
	CacheMode TokenCacheMode `json:"cacheMode"`

	// ContractId The contract ID (eMAID) associated with the token (with optional component separators)
	ContractId string `json:"contractId"`

	// CountryCode The country code of the issuing eMSP
	CountryCode string `json:"countryCode"`

	// GroupId This id groups a couple of tokens to make two or more tokens work as one
	GroupId *string `json:"groupId,omitempty"`

	// Issuer Issuing company, most of the times the name of the company printed on the RFID card, not necessarily the eMSP
	Issuer string `json:"issuer"`

	// LanguageCode The preferred language to use encoded as ISO 639-1 language code
	LanguageCode *string `json:"languageCode,omitempty"`

	// LastUpdated The date the record was last updated (ignored on create/update)
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`

	// PartyId The party id of the issuing eMSP
	PartyId string `json:"partyId"`

	// Type The type of token
	Type TokenType `json:"type"`

	// Uid The unique token id
	Uid string `json:"uid"`

	// Valid Is this token valid
	Valid bool `json:"valid"`

	// VisualNumber The visual/readable number/identification printed on an RFID card
	VisualNumber *string `json:"visualNumber,omitempty"`
}

// TokenCacheMode Indicates what type of token caching is allowed
type TokenCacheMode string

// TokenType The type of token
type TokenType string

// ListTokensParams defines parameters for ListTokens.
type ListTokensParams struct {
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
	Limit  *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// UploadCertificateJSONRequestBody defines body for UploadCertificate for application/json ContentType.
type UploadCertificateJSONRequestBody = Certificate

// RegisterChargeStationJSONRequestBody defines body for RegisterChargeStation for application/json ContentType.
type RegisterChargeStationJSONRequestBody = ChargeStationAuth

// InstallChargeStationCertificatesJSONRequestBody defines body for InstallChargeStationCertificates for application/json ContentType.
type InstallChargeStationCertificatesJSONRequestBody = ChargeStationInstallCertificates

// ReconfigureChargeStationJSONRequestBody defines body for ReconfigureChargeStation for application/json ContentType.
type ReconfigureChargeStationJSONRequestBody = ChargeStationSettings

// TriggerChargeStationJSONRequestBody defines body for TriggerChargeStation for application/json ContentType.
type TriggerChargeStationJSONRequestBody = ChargeStationTrigger

// RegisterLocationJSONRequestBody defines body for RegisterLocation for application/json ContentType.
type RegisterLocationJSONRequestBody = Location

// RegisterPartyJSONRequestBody defines body for RegisterParty for application/json ContentType.
type RegisterPartyJSONRequestBody = Registration

// SetTokenJSONRequestBody defines body for SetToken for application/json ContentType.
type SetTokenJSONRequestBody = Token

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Upload a certificate
	// (POST /certificate)
	UploadCertificate(w http.ResponseWriter, r *http.Request)
	// Delete a certificate
	// (DELETE /certificate/{certificateHash})
	DeleteCertificate(w http.ResponseWriter, r *http.Request, certificateHash string)
	// Lookup a certificate
	// (GET /certificate/{certificateHash})
	LookupCertificate(w http.ResponseWriter, r *http.Request, certificateHash string)
	// Register a new charge station
	// (POST /cs/{csId})
	RegisterChargeStation(w http.ResponseWriter, r *http.Request, csId string)
	// Returns the authentication details
	// (GET /cs/{csId}/auth)
	LookupChargeStationAuth(w http.ResponseWriter, r *http.Request, csId string)
	// Install certificates on the charge station
	// (POST /cs/{csId}/certificates)
	InstallChargeStationCertificates(w http.ResponseWriter, r *http.Request, csId string)
	// Reconfigure the charge station
	// (POST /cs/{csId}/reconfigure)
	ReconfigureChargeStation(w http.ResponseWriter, r *http.Request, csId string)

	// (POST /cs/{csId}/trigger)
	TriggerChargeStation(w http.ResponseWriter, r *http.Request, csId string)
	// Registers a location with the CSMS
	// (POST /location/{locationId})
	RegisterLocation(w http.ResponseWriter, r *http.Request, locationId string)
	// Registers an OCPI party with the CSMS
	// (POST /register)
	RegisterParty(w http.ResponseWriter, r *http.Request)
	// List authorization tokens
	// (GET /token)
	ListTokens(w http.ResponseWriter, r *http.Request, params ListTokensParams)
	// Create/update an authorization token
	// (POST /token)
	SetToken(w http.ResponseWriter, r *http.Request)
	// Lookup an authorization token
	// (GET /token/{tokenUid})
	LookupToken(w http.ResponseWriter, r *http.Request, tokenUid string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// UploadCertificate operation middleware
func (siw *ServerInterfaceWrapper) UploadCertificate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UploadCertificate(w, r)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteCertificate operation middleware
func (siw *ServerInterfaceWrapper) DeleteCertificate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "certificateHash" -------------
	var certificateHash string

	err = runtime.BindStyledParameterWithLocation("simple", false, "certificateHash", runtime.ParamLocationPath, chi.URLParam(r, "certificateHash"), &certificateHash)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "certificateHash", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteCertificate(w, r, certificateHash)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// LookupCertificate operation middleware
func (siw *ServerInterfaceWrapper) LookupCertificate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "certificateHash" -------------
	var certificateHash string

	err = runtime.BindStyledParameterWithLocation("simple", false, "certificateHash", runtime.ParamLocationPath, chi.URLParam(r, "certificateHash"), &certificateHash)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "certificateHash", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LookupCertificate(w, r, certificateHash)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// RegisterChargeStation operation middleware
func (siw *ServerInterfaceWrapper) RegisterChargeStation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "csId" -------------
	var csId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "csId", runtime.ParamLocationPath, chi.URLParam(r, "csId"), &csId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "csId", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RegisterChargeStation(w, r, csId)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// LookupChargeStationAuth operation middleware
func (siw *ServerInterfaceWrapper) LookupChargeStationAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "csId" -------------
	var csId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "csId", runtime.ParamLocationPath, chi.URLParam(r, "csId"), &csId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "csId", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LookupChargeStationAuth(w, r, csId)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// InstallChargeStationCertificates operation middleware
func (siw *ServerInterfaceWrapper) InstallChargeStationCertificates(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "csId" -------------
	var csId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "csId", runtime.ParamLocationPath, chi.URLParam(r, "csId"), &csId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "csId", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.InstallChargeStationCertificates(w, r, csId)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ReconfigureChargeStation operation middleware
func (siw *ServerInterfaceWrapper) ReconfigureChargeStation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "csId" -------------
	var csId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "csId", runtime.ParamLocationPath, chi.URLParam(r, "csId"), &csId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "csId", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ReconfigureChargeStation(w, r, csId)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// TriggerChargeStation operation middleware
func (siw *ServerInterfaceWrapper) TriggerChargeStation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "csId" -------------
	var csId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "csId", runtime.ParamLocationPath, chi.URLParam(r, "csId"), &csId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "csId", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.TriggerChargeStation(w, r, csId)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// RegisterLocation operation middleware
func (siw *ServerInterfaceWrapper) RegisterLocation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "locationId" -------------
	var locationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "locationId", runtime.ParamLocationPath, chi.URLParam(r, "locationId"), &locationId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "locationId", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RegisterLocation(w, r, locationId)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// RegisterParty operation middleware
func (siw *ServerInterfaceWrapper) RegisterParty(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RegisterParty(w, r)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ListTokens operation middleware
func (siw *ServerInterfaceWrapper) ListTokens(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListTokensParams

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListTokens(w, r, params)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// SetToken operation middleware
func (siw *ServerInterfaceWrapper) SetToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SetToken(w, r)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// LookupToken operation middleware
func (siw *ServerInterfaceWrapper) LookupToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "tokenUid" -------------
	var tokenUid string

	err = runtime.BindStyledParameterWithLocation("simple", false, "tokenUid", runtime.ParamLocationPath, chi.URLParam(r, "tokenUid"), &tokenUid)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tokenUid", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LookupToken(w, r, tokenUid)
	})

	for i := len(siw.HandlerMiddlewares) - 1; i >= 0; i-- {
		handler = siw.HandlerMiddlewares[i](handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/certificate", wrapper.UploadCertificate)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/certificate/{certificateHash}", wrapper.DeleteCertificate)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/certificate/{certificateHash}", wrapper.LookupCertificate)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/cs/{csId}", wrapper.RegisterChargeStation)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/cs/{csId}/auth", wrapper.LookupChargeStationAuth)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/cs/{csId}/certificates", wrapper.InstallChargeStationCertificates)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/cs/{csId}/reconfigure", wrapper.ReconfigureChargeStation)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/cs/{csId}/trigger", wrapper.TriggerChargeStation)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/location/{locationId}", wrapper.RegisterLocation)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/register", wrapper.RegisterParty)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/token", wrapper.ListTokens)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/token", wrapper.SetToken)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/token/{tokenUid}", wrapper.LookupToken)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
