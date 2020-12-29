package mergedusers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	rc "github.com/grokify/go-ringcentral/office/v1/client"
	"github.com/grokify/simplego/mime/multipartutil"
	"github.com/nyaruka/phonenumbers"
)

type MergedUserSet struct {
	UserIdOrder   []string
	MergedUserMap map[string]MergedUser
}

func NewMergedUserSet() MergedUserSet {
	return MergedUserSet{
		UserIdOrder:   []string{},
		MergedUserMap: map[string]MergedUser{}}
}

type MergedUser struct {
	MainNumber     string
	GlipPersonInfo rc.GlipPersonInfo
	ExtensionInfo  rc.GetExtensionInfoResponse
}

func (mu *MergedUser) TrimNumbers() {
	mu.MainNumber = strings.TrimSpace(mu.MainNumber)
	mu.ExtensionInfo.Contact.BusinessPhone = strings.TrimSpace(mu.ExtensionInfo.Contact.BusinessPhone)
	mu.ExtensionInfo.ExtensionNumber = strings.TrimSpace(mu.ExtensionInfo.ExtensionNumber)
}

func (mu *MergedUser) PhoneBusinessOrMainNumberExt() string {
	mu.TrimNumbers()

	if len(mu.ExtensionInfo.Contact.BusinessPhone) > 0 {
		return mu.ExtensionInfo.Contact.BusinessPhone
	}
	if len(mu.MainNumber) > 0 &&
		len(mu.ExtensionInfo.ExtensionNumber) > 0 {
		return fmt.Sprintf("%s;ext=%s", mu.MainNumber, mu.ExtensionInfo.ExtensionNumber)
	} else if len(mu.MainNumber) > 0 {
		return mu.MainNumber
	} else if len(mu.ExtensionInfo.ExtensionNumber) > 0 {
		return fmt.Sprintf("ext=%s", mu.ExtensionInfo.ExtensionNumber)
	}
	return ""
}

func (mu *MergedUser) DisplayNumber() (string, error) {
	mu.TrimNumbers()

	if len(mu.ExtensionInfo.Contact.BusinessPhone) > 0 {
		num, err := phonenumbers.Parse(mu.ExtensionInfo.Contact.BusinessPhone, "")
		if err != nil {
			return "", err
		}
		return phonenumbers.Format(num, phonenumbers.NATIONAL), nil
	} else if len(mu.MainNumber) > 0 {
		num, err := phonenumbers.Parse(mu.MainNumber, "")
		if err != nil {
			return "", err
		}
		main := phonenumbers.Format(num, phonenumbers.NATIONAL)
		if len(mu.ExtensionInfo.ExtensionNumber) > 0 {
			main = fmt.Sprintf("%s x%s", main, mu.ExtensionInfo.ExtensionNumber)
		}
		return main, nil
	} else if len(mu.ExtensionInfo.ExtensionNumber) > 0 {
		return "x" + mu.ExtensionInfo.ExtensionNumber, nil
	}
	return "", errors.New("No phone number")
}

type MergedUserThin struct {
	Id              string
	FirstName       string
	LastName        string
	Email           string
	Avatar          string
	CompanyId       string
	JobTitle        string
	MainNumber      string
	ExtensionNumber string
	BusinessPhone   string
	DisplayNumber   string
}

func (mu *MergedUser) ToMergedUserThin() MergedUserThin {
	per := mu.GlipPersonInfo
	ext := mu.ExtensionInfo
	thin := MergedUserThin{
		Id:              per.Id,
		FirstName:       per.FirstName,
		LastName:        per.LastName,
		Email:           per.Email,
		Avatar:          per.Avatar,
		CompanyId:       per.CompanyId,
		JobTitle:        per.JobTitle,
		MainNumber:      mu.MainNumber,
		ExtensionNumber: ext.ExtensionNumber,
		BusinessPhone:   ext.Contact.BusinessPhone}
	if disp, err := mu.DisplayNumber(); err == nil {
		thin.DisplayNumber = disp
	}
	return thin
}

func NewMergedUserFromMimePart(part *multipart.Part) (MergedUser, error) {
	mu := MergedUser{}
	if slurp, err := ioutil.ReadAll(part); err != nil {
		return mu, err
	} else if err = json.Unmarshal(slurp, &mu.GlipPersonInfo); err != nil {
		return mu, err
	}
	return mu, nil
}

func NewGlipPersonInfoFromMimePart(part *multipart.Part) (rc.GlipPersonInfo, error) {
	var person rc.GlipPersonInfo
	if slurp, err := ioutil.ReadAll(part); err != nil {
		return person, err
	} else {
		return person, json.Unmarshal(slurp, &person)
	}
}

func NewExtensionInfoFromMimePart(part *multipart.Part) (rc.GetExtensionInfoResponse, error) {
	var ext rc.GetExtensionInfoResponse
	if slurp, err := ioutil.ReadAll(part); err != nil {
		return ext, err
	} else {
		return ext, json.Unmarshal(slurp, &ext)
	}
}

func AddBatchGlipPersonInfosBodyBoundary(mergedUserSet MergedUserSet, body []byte, boundary string) (MergedUserSet, error) {
	mr := multipartutil.NewReaderBodyBytes([]byte(body), boundary)
	return AddBatchGlipPersonInfosMultipartReader(mergedUserSet, mr)
}

func AddBatchGlipPersonInfosMultipartReader(mergedUserSet MergedUserSet, mr *multipart.Reader) (MergedUserSet, error) {
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return mergedUserSet, err
		}
		if person, err := NewGlipPersonInfoFromMimePart(part); err != nil {
			return mergedUserSet, err
		} else if len(person.Id) > 0 {
			if mu, ok := mergedUserSet.MergedUserMap[person.Id]; ok {
				mu.GlipPersonInfo = person
				mergedUserSet.MergedUserMap[person.Id] = mu
			} else {
				mergedUserSet.MergedUserMap[person.Id] = MergedUser{GlipPersonInfo: person}
				mergedUserSet.UserIdOrder = append(mergedUserSet.UserIdOrder, person.Id)
			}
		}
	}
	return mergedUserSet, nil
}

func AddBatchExtensionInfosHttpResponse(mergedUserSet MergedUserSet, resp *http.Response) (MergedUserSet, error) {
	mr, err := multipartutil.NewMultipartReaderForHttpResponse(resp)
	if err != nil {
		return mergedUserSet, err
	}
	return AddBatchExtensionInfosMultipartReader(mergedUserSet, mr)
}

func AddBatchExtensionInfosMultipartReader(mergedUserSet MergedUserSet, mr *multipart.Reader) (MergedUserSet, error) {
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return mergedUserSet, err
		}
		if ext, err := NewExtensionInfoFromMimePart(part); err != nil {
			return mergedUserSet, err
		} else if ext.Id > 0 {
			extIdString := strconv.Itoa(int(ext.Id))
			if mu, ok := mergedUserSet.MergedUserMap[extIdString]; ok {
				mu.ExtensionInfo = ext
				mergedUserSet.MergedUserMap[extIdString] = mu
			} else {
				mergedUserSet.MergedUserMap[extIdString] = MergedUser{ExtensionInfo: ext}
				mergedUserSet.UserIdOrder = append(mergedUserSet.UserIdOrder, extIdString)
			}
		}
	}
	return mergedUserSet, nil
}
