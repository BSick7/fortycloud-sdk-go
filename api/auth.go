package fortycloud

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "bytes"
    "fmt"
	"errors"
)

type Credentials struct {
    Username string
    Password string
}
type Authentication struct {
    Credentials
    Tenant string
    Token string
    Expires string
}

type authRequest struct {
    Credentials Credentials `json:"passwordCredentials"`
    Tenant string `json:"tenantName"`
}
type successResult struct {
    Access struct {
        Token struct {
            Id string
            Expires string
        }
    }
}
type failResult struct {
    IdentityFault struct {
        Code string
        Message string
        Details string
    }
}

func (a *Authentication) Do(s *Service) error {
    url := API_URL + "/v0.4/tokens"
    
    jsonStr, err := json.Marshal(&authRequest {
        Credentials: a.Credentials,
        Tenant: a.Tenant,
    })
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    
    res, err := s.Client.Do(req)
    if err != nil {
        return err
    }
    
    defer res.Body.Close()
    
    body, _ := ioutil.ReadAll(res.Body)
    
    var result successResult
    err = json.Unmarshal(body, &result)
    if err != nil {
        var result2 failResult
        err = json.Unmarshal(body, &result2)
        if err != nil {
            return err
        }
        fault := result2.IdentityFault
        return errors.New(fmt.Sprintf("[%s] '%s': %s", fault.Code, fault.Message, fault.Details))
    }
    
    a.Token = result.Access.Token.Id
    a.Expires = result.Access.Token.Expires
    
    return nil
}