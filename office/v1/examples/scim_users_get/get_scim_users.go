package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/antihax/optional"
	"github.com/grokify/goauth/credentials"
	"github.com/grokify/mogo/errors/errorsutil"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	rcu "github.com/grokify/go-ringcentral-client/office/v1/util"
	rs "github.com/grokify/go-scim-client"
)

const (
	envPrefix = "RINGCENTRAL_"
)

func loadEnv() error {
	envPaths := []string{}
	if len(os.Getenv("ENV_PATH")) > 0 {
		envPaths = append(envPaths, os.Getenv("ENV_PATH"))
	}
	return godotenv.Load(envPaths...)
}

type ScimApiUtil struct {
	ApiClient *rc.APIClient
	Context   context.Context
}

func (su *ScimApiUtil) GetSCIMUserById(userId string) (rc.UserInfo, error) {
	uinfo, resp, err := su.ApiClient.SCIMApi.GetUserById(su.Context, userId)
	if err != nil {
		return uinfo, err
	} else if resp.StatusCode >= 300 {
		return uinfo, fmt.Errorf("API Status %v", resp.StatusCode)
	}
	return uinfo, nil
}

func (su *ScimApiUtil) ListScimUsers(params rc.ListUsersOpts) (rc.GetUserListResponse, error) {
	info, resp, err := su.ApiClient.SCIMApi.ListUsers(su.Context, &params)
	if err != nil {
		return info, err
	} else if resp.StatusCode >= 300 {
		return info, fmt.Errorf("API Status %v", resp.StatusCode)
	}
	return info, nil
}

func mainClient() {
	err := loadEnv()
	if err != nil {
		panic(err)
	}
	apiClient, err := rcu.NewApiClientPasswordEnv(envPrefix)
	if err != nil {
		panic(err)
	}

	scimUtil := ScimApiUtil{
		ApiClient: apiClient,
		Context:   context.Background(),
	}

	// Retrieve a list of users
	info, err := scimUtil.ListScimUsers(rc.ListUsersOpts{})
	if err != nil {
		panic(err)
	}
	fmtutil.PrintJSON(info)

	for _, user := range info.Resources {
		fmt.Println(user.Id)

		if 1 == 0 {
			uinfo, err := scimUtil.GetSCIMUserById(user.Id)
			if err != nil {
				panic(err)
			}
			fmtutil.PrintJSON(uinfo)
		}
	}
	fmt.Println("DONE")
}

func getPathMembersValue(v string) string {
	return fmt.Sprintf("members[value eq \"%v\"]", v)
}

func tryPatch(scimClient *rs.APIClient, ctx context.Context, userId string) {
	body := rs.UserPatch{
		Operations: []rs.PatchOperation{
			{
				Op: "add",
				Value: rs.User{
					Addresses: []rs.Address{
						{
							Type_:         "work",
							StreetAddress: "100 Main Street",
							Locality:      "South Park",
							Region:        "CO",
							PostalCode:    "80440",
							Country:       "US",
						},
					},
				},
			},
		},
		Schemas: []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
	}
	fmtutil.PrintJSON(body)

	info, resp, err := scimClient.UserApi.PatchUser(ctx, userId,
		map[string]interface{}{"body": body},
	)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)
}

func tryReplace(scimClient *rs.APIClient, ctx context.Context, userId string) {
	userResp, resp, err := scimClient.UserApi.GetUserById(ctx, userId)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(userResp)
	user := rs.User{
		Active:       userResp.Active,
		Addresses:    userResp.Addresses,
		Emails:       userResp.Emails,
		ExternalId:   userResp.ExternalId,
		Id:           userResp.Id,
		Name:         userResp.Name,
		PhoneNumbers: userResp.PhoneNumbers,
		Photos:       userResp.Photos,
		Schemas:      userResp.Schemas,
		UserName:     userResp.UserName,
		Urnietfparamsscimschemasextensionenterprise20User: userResp.Urnietfparamsscimschemasextensionenterprise20User,
	}
	user.Name.FamilyName = "RingForceAwakens"

	userResp, resp, err = scimClient.UserApi.ReplaceUser(ctx, userId,
		map[string]interface{}{"body": user})
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(userResp)
}

func tryDelete(scimClient *rs.APIClient, ctx context.Context, userId string) {
	resp, err := scimClient.UserApi.DeleteUser(ctx, userId)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
}

func tryCreate(scimClient *rs.APIClient, ctx context.Context) {
	email := "john@example.com"
	user := rs.User{
		Name: &rs.Name{
			GivenName:  "John",
			FamilyName: "RingCentral",
		},
		Emails: []rs.Email{
			{
				Value: email,
				Type_: "work",
			},
		},
		Schemas:  []string{"urn:ietf:params:scim:schemas:core:2.0:User"},
		UserName: email,
	}

	info, resp, err := scimClient.UserApi.CreateUser(
		ctx, map[string]interface{}{"body": user})
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)
}

func tryScimClient() {
	scimClient, err := rcu.NewScimApiClient(
		credentials.NewCredentialsOAuth2Env(envPrefix))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	users, resp, err := scimClient.UserApi.SearchViaGet(ctx, map[string]interface{}{})
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(users)

	if 1 == 0 {
		userId := "1"
		tryPatch(scimClient, ctx, userId)
	}

	if 1 == 0 {
		userId := "1"
		tryDelete(scimClient, ctx, userId)
	}
	if 1 == 0 {
		tryCreate(scimClient, ctx)
	}
	if 1 == 0 {
		userId := "1"
		tryReplace(scimClient, ctx, userId)
	}
}

func tryRingCentralClient(demoConfig DemoConfig) {
	apiClient, err := rcu.NewApiClientPassword(
		credentials.NewCredentialsOAuth2Env(envPrefix))
	if err != nil {
		log.Fatal(err)
	}

	demo := ScimRingCentralDemo{
		Client:  apiClient,
		Context: context.Background(),
		Config:  demoConfig,
	}

	if demoConfig.GetServiceProviderConfig {
		demo.GetServiceProviderConfig()
	}

	demo.GetUsers()

	if demoConfig.DeleteUser {
		demo.DeleteUser(demoConfig.UserId)
		demo.GetUsers()
	}
	if demoConfig.CreateUser {
		demo.CreateUser()
	}
	if demoConfig.PatchUser {
		userId := "1"
		demo.PatchUser(userId)
	}
	if 1 == 0 {
		userId := "1"
		demo.GetUser(userId)
	}
	if demoConfig.ReplaceUser {
		userId := "1"
		demo.UpdateUser(userId)
	}
}

type ScimRingCentralDemo struct {
	Client  *rc.APIClient
	Config  DemoConfig
	Context context.Context
}

func (demo ScimRingCentralDemo) GetServiceProviderConfig() {
	cfg, resp, err := demo.Client.SCIMApi.GetServiceProviderConfig(demo.Context)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(cfg)
}

func (demo ScimRingCentralDemo) GetUsers() {
	users, resp, err := demo.Client.SCIMApi.ListUsers(
		demo.Context, &rc.ListUsersOpts{})
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(users)
}

func (demo ScimRingCentralDemo) GetUser(userId string) rc.UserInfo {
	info, resp, err := demo.Client.SCIMApi.GetUserById(demo.Context, userId)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)
	return info
}

func (demo ScimRingCentralDemo) DeleteUser(userId string) error {
	resp, err := demo.Client.SCIMApi.DeleteUser(demo.Context, userId)
	if err != nil {
		return errorsutil.Wrap(err, "err ScimRingCentralDemo.DeleteUser()")
		// log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("invalid HTTP status [%d]", resp.StatusCode)
		// log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	return httputilmore.PrintResponse(resp, true)
}

func (demo ScimRingCentralDemo) UpdateUser(userId string) {
	userInfo, resp, err := demo.Client.SCIMApi.GetUserById(demo.Context, userId)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(userInfo)

	userReq := rc.UserUpdateRequest{
		Active:       userInfo.Active,
		Addresses:    userInfo.Addresses,
		Emails:       userInfo.Emails,
		ExternalId:   userInfo.ExternalId,
		Id:           userInfo.Id,
		Name:         userInfo.Name,
		PhoneNumbers: userInfo.PhoneNumbers,
		Photos:       userInfo.Photos,
		Schemas:      userInfo.Schemas,
		UserName:     userInfo.UserName,
		Urnietfparamsscimschemasextensionenterprise20User: userInfo.Urnietfparamsscimschemasextensionenterprise20User,
	}

	userReq.Name.FamilyName = "RingForceAwakens"

	userResp, resp, err := demo.Client.SCIMApi.UpdateUser(demo.Context, userId, userReq)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(userResp)
}

func (demo ScimRingCentralDemo) CreateUser() {
	user := demo.GetNewUser()

	info, resp, err := demo.Client.SCIMApi.CreateUser(demo.Context, user)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)
}

func (demo ScimRingCentralDemo) PatchUser(userId string) {
	body := rc.ScimUserPatch{
		Operations: []rc.PatchOperationInfo{
			{
				Op: "add",
				Value: map[string]interface{}{

					"addresses": []rc.AddressInfo{
						{
							Type:          "work",
							StreetAddress: "100 Main Street",
							Locality:      "South Park",
							Region:        "CO",
							PostalCode:    "80440",
							Country:       "US",
						},
					},
				},
			},
		},
		Schemas: []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
	}
	fmtutil.PrintJSON(body)

	info, resp, err := demo.Client.SCIMApi.PatchUser(demo.Context, userId,
		&rc.PatchUserOpts{ScimUserPatch: optional.NewInterface(body)})
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal(fmt.Sprintf("StatusCode: %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)
}

func (demo ScimRingCentralDemo) GetNewUser() rc.UserCreationRequest {
	email := demo.Config.UserEmail
	if len(email) == 0 {
		email = "john@example.com"
	}
	user := rc.UserCreationRequest{
		Name: rc.NameInfo{
			GivenName:  "John",
			FamilyName: "RingCentral",
		},
		Emails: []rc.EmailInfo{
			{
				Value: email,
				Type:  "work",
			},
		},
		Schemas:  []string{"urn:ietf:params:scim:schemas:core:2.0:User"},
		UserName: email,
	}
	return user
}

type DemoConfig struct {
	GetServiceProviderConfig bool
	CreateUser               bool
	PatchUser                bool
	ReplaceUser              bool
	DeleteUser               bool
	UserEmail                string
	UserId                   string
}

func main() {
	var createUser = flag.Int("create", 0, "Create user")
	var patchUser = flag.Int("patch", 0, "Patch user")
	var replaceUser = flag.Int("replace", 0, "Replace user")
	var deleteUser = flag.Int("delete", 0, "Delete user")

	var userEmail = flag.String("email", "", "Email address")
	var userId = flag.String("userid", "", "Email address")

	flag.Parse()

	fmt.Printf("[%v]\n", *createUser)

	cfg := DemoConfig{
		GetServiceProviderConfig: true,
	}
	if *createUser != 0 {
		cfg.CreateUser = true
	}
	if *patchUser != 0 {
		cfg.PatchUser = true
	}
	if *replaceUser != 0 {
		cfg.ReplaceUser = true
	}
	if *deleteUser != 0 {
		cfg.DeleteUser = true
	}
	if len(*userEmail) > 0 {
		cfg.UserEmail = strings.TrimSpace(*userEmail)
	}
	if len(*userId) > 0 {
		cfg.UserId = strings.TrimSpace(*userId)
	}

	fmtutil.PrintJSON(cfg)

	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	if 1 == 0 {
		tryScimClient()
	}
	if 1 == 1 {
		tryRingCentralClient(cfg)
	}
	fmt.Println("DONE")
}
