package fortycloud

import (
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "bytes"
    "fmt"
	"errors"
)

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
type Authentication struct {
    Credentials
    Tenant string
    Token string
    Expires string
}

type authPost struct {
    Auth authRequest `json:"auth"`
}
type authRequest struct {
    Credentials Credentials `json:"passwordCredentials"`
    Tenant string `json:"tenantName"`
}
type authResult struct {
    Access struct {
        Token struct {
            Id string
            Expires string
        }
    }
    IdentityFault struct {
        Code string
        Message string
        Details string
    }
}

func (a *Authentication) Do(s *Service) error {
    jsonStr, err := json.Marshal(&authPost {
        Auth: authRequest {
            Credentials: a.Credentials,
            Tenant: a.Tenant,
        },
    })
    if err != nil {
        return err
    }
    
    log.Printf("%s", jsonStr)
    
    req, err := http.NewRequest("POST", s.Url + "/tokens", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    
    res, err := s.Client.Do(req)
    if err != nil {
        return err
    }
    
    defer res.Body.Close()
    
    body, _ := ioutil.ReadAll(res.Body)
    
    log.Printf("%s", body)
    
    var result authResult
    err = json.Unmarshal(body, &result)
    if err != nil {
        return err
    }
    
    fault := result.IdentityFault
    if len(fault.Code) > 0 {
        return errors.New(fmt.Sprintf("[%s] '%s': %s", fault.Code, fault.Message, fault.Details))   
    }
    
    a.Token = result.Access.Token.Id
    a.Expires = result.Access.Token.Expires
    
    return nil
}