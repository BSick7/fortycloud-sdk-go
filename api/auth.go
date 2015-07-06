package fortycloud

type Authentication struct {
    Username string
    Password string
    Tenant string
    Token string
    Expires string
}

type authRequest struct {
    passwordCredentials struct {
        username string
        password string
    }
    tenantName string
}
type authResult struct {
    access struct {
        token struct {
            id string
            expires string
        }
    }
}
type failAuthResult struct {
    identityFault struct {
        code string
        message string
        details string
    }
}

func (a *Authentication) Do(s *Service) error {
    url := API_URL + "/v0.4/tokens"
    
    jsonStr, err := json.Marshal(&authRequest {
        passwordCredentials: &credRequest {
            username: a.Username,
            password: a.Password
        },
        tenantName: a.Tenant
    })
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    
    res, err := s.client.Do(req)
    if err != nil {
        return err
    }
    
    defer res.Body.Close()
    
    body, _ := ioutil.ReadAll(res.Body)
    
    result authResult
    err := json.Unmarshal(string(body), &result)
    if err != nil {
        result2 failAuthResult
        err := json.Unmarshal(string(body), &result2)
        if err != nil
            return err
    }
    
    return nil
}