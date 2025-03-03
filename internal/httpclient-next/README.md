# Go API client for client

Documentation for all of Ory Hydra's APIs.

## Overview

This API client was generated by the
[OpenAPI Generator](https://openapi-generator.tech) project. By using the
[OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily
generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import client "github.com/ory/hydra-client-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as
defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value
`sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or
from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently
ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map
in the `Configuration`. An operation is uniquely identifield by
`"{classname}Service.{nickname}"` string. Similar rules for overriding default
operation server index and variables applies by using
`sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables`
context maps.

```
ctx := context.WithValue(context.Background(), client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to _http://localhost_

| Class         | Method                                                                                                           | HTTP request                                     | Description                                                                                            |
| ------------- | ---------------------------------------------------------------------------------------------------------------- | ------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| _AdminApi_    | [**AcceptConsentRequest**](docs/AdminApi.md#acceptconsentrequest)                                                | **Put** /oauth2/auth/requests/consent/accept     | Accept a Consent Request                                                                               |
| _AdminApi_    | [**AcceptLoginRequest**](docs/AdminApi.md#acceptloginrequest)                                                    | **Put** /oauth2/auth/requests/login/accept       | Accept a Login Request                                                                                 |
| _AdminApi_    | [**AcceptLogoutRequest**](docs/AdminApi.md#acceptlogoutrequest)                                                  | **Put** /oauth2/auth/requests/logout/accept      | Accept a Logout Request                                                                                |
| _AdminApi_    | [**CreateJsonWebKeySet**](docs/AdminApi.md#createjsonwebkeyset)                                                  | **Post** /keys/{set}                             | Generate a New JSON Web Key                                                                            |
| _AdminApi_    | [**CreateOAuth2Client**](docs/AdminApi.md#createoauth2client)                                                    | **Post** /clients                                | Create an OAuth 2.0 Client                                                                             |
| _AdminApi_    | [**DeleteJsonWebKey**](docs/AdminApi.md#deletejsonwebkey)                                                        | **Delete** /keys/{set}/{kid}                     | Delete a JSON Web Key                                                                                  |
| _AdminApi_    | [**DeleteJsonWebKeySet**](docs/AdminApi.md#deletejsonwebkeyset)                                                  | **Delete** /keys/{set}                           | Delete a JSON Web Key Set                                                                              |
| _AdminApi_    | [**DeleteOAuth2Client**](docs/AdminApi.md#deleteoauth2client)                                                    | **Delete** /clients/{id}                         | Deletes an OAuth 2.0 Client                                                                            |
| _AdminApi_    | [**DeleteOAuth2Token**](docs/AdminApi.md#deleteoauth2token)                                                      | **Delete** /oauth2/tokens                        | Delete OAuth2 Access Tokens from a Client                                                              |
| _AdminApi_    | [**DeleteTrustedJwtGrantIssuer**](docs/AdminApi.md#deletetrustedjwtgrantissuer)                                  | **Delete** /trust/grants/jwt-bearer/issuers/{id} | Delete a Trusted OAuth2 JWT Bearer Grant Type Issuer                                                   |
| _AdminApi_    | [**FlushInactiveOAuth2Tokens**](docs/AdminApi.md#flushinactiveoauth2tokens)                                      | **Post** /oauth2/flush                           | Flush Expired OAuth2 Access Tokens                                                                     |
| _AdminApi_    | [**GetConsentRequest**](docs/AdminApi.md#getconsentrequest)                                                      | **Get** /oauth2/auth/requests/consent            | Get Consent Request Information                                                                        |
| _AdminApi_    | [**GetJsonWebKey**](docs/AdminApi.md#getjsonwebkey)                                                              | **Get** /keys/{set}/{kid}                        | Fetch a JSON Web Key                                                                                   |
| _AdminApi_    | [**GetJsonWebKeySet**](docs/AdminApi.md#getjsonwebkeyset)                                                        | **Get** /keys/{set}                              | Retrieve a JSON Web Key Set                                                                            |
| _AdminApi_    | [**GetLoginRequest**](docs/AdminApi.md#getloginrequest)                                                          | **Get** /oauth2/auth/requests/login              | Get a Login Request                                                                                    |
| _AdminApi_    | [**GetLogoutRequest**](docs/AdminApi.md#getlogoutrequest)                                                        | **Get** /oauth2/auth/requests/logout             | Get a Logout Request                                                                                   |
| _AdminApi_    | [**GetOAuth2Client**](docs/AdminApi.md#getoauth2client)                                                          | **Get** /clients/{id}                            | Get an OAuth 2.0 Client                                                                                |
| _AdminApi_    | [**GetTrustedJwtGrantIssuer**](docs/AdminApi.md#gettrustedjwtgrantissuer)                                        | **Get** /trust/grants/jwt-bearer/issuers/{id}    | Get a Trusted OAuth2 JWT Bearer Grant Type Issuer                                                      |
| _AdminApi_    | [**IntrospectOAuth2Token**](docs/AdminApi.md#introspectoauth2token)                                              | **Post** /oauth2/introspect                      | Introspect OAuth2 Tokens                                                                               |
| _AdminApi_    | [**ListOAuth2Clients**](docs/AdminApi.md#listoauth2clients)                                                      | **Get** /clients                                 | List OAuth 2.0 Clients                                                                                 |
| _AdminApi_    | [**ListSubjectConsentSessions**](docs/AdminApi.md#listsubjectconsentsessions)                                    | **Get** /oauth2/auth/sessions/consent            | Lists All Consent Sessions of a Subject                                                                |
| _AdminApi_    | [**ListTrustedJwtGrantIssuers**](docs/AdminApi.md#listtrustedjwtgrantissuers)                                    | **Get** /trust/grants/jwt-bearer/issuers         | List Trusted OAuth2 JWT Bearer Grant Type Issuers                                                      |
| _AdminApi_    | [**PatchOAuth2Client**](docs/AdminApi.md#patchoauth2client)                                                      | **Patch** /clients/{id}                          | Patch an OAuth 2.0 Client                                                                              |
| _AdminApi_    | [**RejectConsentRequest**](docs/AdminApi.md#rejectconsentrequest)                                                | **Put** /oauth2/auth/requests/consent/reject     | Reject a Consent Request                                                                               |
| _AdminApi_    | [**RejectLoginRequest**](docs/AdminApi.md#rejectloginrequest)                                                    | **Put** /oauth2/auth/requests/login/reject       | Reject a Login Request                                                                                 |
| _AdminApi_    | [**RejectLogoutRequest**](docs/AdminApi.md#rejectlogoutrequest)                                                  | **Put** /oauth2/auth/requests/logout/reject      | Reject a Logout Request                                                                                |
| _AdminApi_    | [**RevokeAuthenticationSession**](docs/AdminApi.md#revokeauthenticationsession)                                  | **Delete** /oauth2/auth/sessions/login           | Invalidates All Login Sessions of a Certain User Invalidates a Subject&#39;s Authentication Session    |
| _AdminApi_    | [**RevokeConsentSessions**](docs/AdminApi.md#revokeconsentsessions)                                              | **Delete** /oauth2/auth/sessions/consent         | Revokes Consent Sessions of a Subject for a Specific OAuth 2.0 Client                                  |
| _AdminApi_    | [**TrustJwtGrantIssuer**](docs/AdminApi.md#trustjwtgrantissuer)                                                  | **Post** /trust/grants/jwt-bearer/issuers        | Trust an OAuth2 JWT Bearer Grant Type Issuer                                                           |
| _AdminApi_    | [**UpdateJsonWebKey**](docs/AdminApi.md#updatejsonwebkey)                                                        | **Put** /keys/{set}/{kid}                        | Update a JSON Web Key                                                                                  |
| _AdminApi_    | [**UpdateJsonWebKeySet**](docs/AdminApi.md#updatejsonwebkeyset)                                                  | **Put** /keys/{set}                              | Update a JSON Web Key Set                                                                              |
| _AdminApi_    | [**UpdateOAuth2Client**](docs/AdminApi.md#updateoauth2client)                                                    | **Put** /clients/{id}                            | Update an OAuth 2.0 Client                                                                             |
| _AdminApi_    | [**UpdateOAuth2ClientLifespans**](docs/AdminApi.md#updateoauth2clientlifespans)                                  | **Put** /clients/{id}/lifespans                  |
| _MetadataApi_ | [**GetVersion**](docs/MetadataApi.md#getversion)                                                                 | **Get** /version                                 | Return Running Software Version.                                                                       |
| _MetadataApi_ | [**IsAlive**](docs/MetadataApi.md#isalive)                                                                       | **Get** /health/alive                            | Check HTTP Server Status                                                                               |
| _MetadataApi_ | [**IsReady**](docs/MetadataApi.md#isready)                                                                       | **Get** /health/ready                            | Check HTTP Server and Database Status                                                                  |
| _PublicApi_   | [**DisconnectUser**](docs/PublicApi.md#disconnectuser)                                                           | **Get** /oauth2/sessions/logout                  | OpenID Connect Front-Backchannel Enabled Logout                                                        |
| _PublicApi_   | [**DiscoverOpenIDConfiguration**](docs/PublicApi.md#discoveropenidconfiguration)                                 | **Get** /.well-known/openid-configuration        | OpenID Connect Discovery                                                                               |
| _PublicApi_   | [**DynamicClientRegistrationCreateOAuth2Client**](docs/PublicApi.md#dynamicclientregistrationcreateoauth2client) | **Post** /oauth2/register                        | Register an OAuth 2.0 Client using the OpenID / OAuth2 Dynamic Client Registration Management Protocol |
| _PublicApi_   | [**DynamicClientRegistrationDeleteOAuth2Client**](docs/PublicApi.md#dynamicclientregistrationdeleteoauth2client) | **Delete** /oauth2/register/{id}                 | Deletes an OAuth 2.0 Client using the OpenID / OAuth2 Dynamic Client Registration Management Protocol  |
| _PublicApi_   | [**DynamicClientRegistrationGetOAuth2Client**](docs/PublicApi.md#dynamicclientregistrationgetoauth2client)       | **Get** /oauth2/register/{id}                    | Get an OAuth 2.0 Client using the OpenID / OAuth2 Dynamic Client Registration Management Protocol      |
| _PublicApi_   | [**DynamicClientRegistrationUpdateOAuth2Client**](docs/PublicApi.md#dynamicclientregistrationupdateoauth2client) | **Put** /oauth2/register/{id}                    | Update an OAuth 2.0 Client using the OpenID / OAuth2 Dynamic Client Registration Management Protocol   |
| _PublicApi_   | [**Oauth2Token**](docs/PublicApi.md#oauth2token)                                                                 | **Post** /oauth2/token                           | The OAuth 2.0 Token Endpoint                                                                           |
| _PublicApi_   | [**OauthAuth**](docs/PublicApi.md#oauthauth)                                                                     | **Get** /oauth2/auth                             | The OAuth 2.0 Authorize Endpoint                                                                       |
| _PublicApi_   | [**RevokeOAuth2Token**](docs/PublicApi.md#revokeoauth2token)                                                     | **Post** /oauth2/revoke                          | Revoke OAuth2 Tokens                                                                                   |
| _PublicApi_   | [**Userinfo**](docs/PublicApi.md#userinfo)                                                                       | **Get** /userinfo                                | OpenID Connect Userinfo                                                                                |
| _PublicApi_   | [**WellKnown**](docs/PublicApi.md#wellknown)                                                                     | **Get** /.well-known/jwks.json                   | JSON Web Keys Discovery                                                                                |

## Documentation For Models

- [AcceptConsentRequest](docs/AcceptConsentRequest.md)
- [AcceptLoginRequest](docs/AcceptLoginRequest.md)
- [CompletedRequest](docs/CompletedRequest.md)
- [ConsentRequest](docs/ConsentRequest.md)
- [ConsentRequestSession](docs/ConsentRequestSession.md)
- [DefaultSession](docs/DefaultSession.md)
- [FlushInactiveOAuth2TokensRequest](docs/FlushInactiveOAuth2TokensRequest.md)
- [FlushLoginConsentRequest](docs/FlushLoginConsentRequest.md)
- [GenericError](docs/GenericError.md)
- [Headers](docs/Headers.md)
- [HealthNotReadyStatus](docs/HealthNotReadyStatus.md)
- [HealthStatus](docs/HealthStatus.md)
- [IDTokenClaims](docs/IDTokenClaims.md)
- [InlineResponse200](docs/InlineResponse200.md)
- [InlineResponse2001](docs/InlineResponse2001.md)
- [InlineResponse503](docs/InlineResponse503.md)
- [JSONWebKey](docs/JSONWebKey.md)
- [JSONWebKeySet](docs/JSONWebKeySet.md)
- [JsonError](docs/JsonError.md)
- [JsonWebKeySetGeneratorRequest](docs/JsonWebKeySetGeneratorRequest.md)
- [LoginRequest](docs/LoginRequest.md)
- [LogoutRequest](docs/LogoutRequest.md)
- [OAuth2AccessRequest](docs/OAuth2AccessRequest.md)
- [OAuth2Client](docs/OAuth2Client.md)
- [OAuth2TokenIntrospection](docs/OAuth2TokenIntrospection.md)
- [Oauth2TokenResponse](docs/Oauth2TokenResponse.md)
- [OauthTokenResponse](docs/OauthTokenResponse.md)
- [OpenIDConnectContext](docs/OpenIDConnectContext.md)
- [PatchDocument](docs/PatchDocument.md)
- [PreviousConsentSession](docs/PreviousConsentSession.md)
- [RefreshTokenHookRequest](docs/RefreshTokenHookRequest.md)
- [RefreshTokenHookResponse](docs/RefreshTokenHookResponse.md)
- [RejectRequest](docs/RejectRequest.md)
- [RequestWasHandledResponse](docs/RequestWasHandledResponse.md)
- [Session](docs/Session.md)
- [TrustJwtGrantIssuerBody](docs/TrustJwtGrantIssuerBody.md)
- [TrustedJsonWebKey](docs/TrustedJsonWebKey.md)
- [TrustedJwtGrantIssuer](docs/TrustedJwtGrantIssuer.md)
- [UpdateOAuth2ClientLifespans](docs/UpdateOAuth2ClientLifespans.md)
- [UserinfoResponse](docs/UserinfoResponse.md)
- [Version](docs/Version.md)
- [WellKnown](docs/WellKnown.md)

## Documentation For Authorization

### basic

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

### oauth2

- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://hydra.demo.ory.sh/oauth2/auth
- **Scopes**:
- **offline**: A scope required when requesting refresh tokens (alias for
  `offline_access`)
- **offline_access**: A scope required when requesting refresh tokens
- **openid**: Request an OpenID Connect ID Token

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user
authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package
contains a number of utility functions to easily obtain pointers to values of
basic types. Each of these functions takes a value of the given basic type and
returns a pointer to it:

- `PtrBool`
- `PtrInt`
- `PtrInt32`
- `PtrInt64`
- `PtrFloat`
- `PtrFloat32`
- `PtrFloat64`
- `PtrString`
- `PtrTime`

## Author

hi@ory.sh
