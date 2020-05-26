package main

import(
    "fmt"
    "net/http"
    "encoding/json"
)

type loginRequest struct {
    Login string
    Password string
}

type loginResponse struct {
    Login string        `json:"login"`
    SessionId string    `json:"sessionId"`
}

type registerRequest struct {
	Name string
	Surname string
	Login string
	Password string
}

type registerResponse struct {
	Login string        `json:"login"`
    SessionId string    `json:"sessionId"`
}

func (ctx *context)handleLoginRequest(w http.ResponseWriter, r *http.Request) {
    var err error
    var req loginRequest;
	var res loginResponse;

    w.Header().Set("Content-Type", "application/json")

    err = json.NewDecoder(r.Body).Decode(&req)
    
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(fmt.Sprintf(`{"message": "Error while unmarshaling login request", "details": "%s"}`, err.Error())))
        // TODO log
        return
    }

    if req.Login == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"message": "Malformed login request", "details": "Login field is empty"}`))
        // TODO log
        return
    }

    if req.Password == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"message": "Malformed login request", "details": "Password field is empty"}`))
        // TODO log
        return
    }

    vlcres, err := ctx.validateLoginCredentials(req)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(fmt.Sprintf(`{"message": "Error while validating login credentials", "details": "%s"}`, err.Error())))
        // TODO log
        return
    }

    if vlcres != "" {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte(fmt.Sprintf(`{"message": "Invalid login credentials", "details": "%s"}`, vlcres)))
        // TODO log
        return
    }

    res, err = ctx.establishSession(req)
        
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(fmt.Sprintf(`{"message": "Error while establishing session", "details": "%s"}`, err.Error())))
        // TODO log
        return
    } 
    
    out, err := json.Marshal(&res)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(fmt.Sprintf(`{"message": "Error while marshaling response", "details": "%s"}`, err.Error())))
        // TODO log
        return
    } 

    w.Write([]byte(out))
    // TODO log
    return
}

func (ctx *context)handleRegisterRequest(w http.ResponseWriter, r *http.Request) {
    var err error
	var req registerRequest;
	var res registerResponse;

    w.Header().Set("Content-Type", "application/json")

    err = json.NewDecoder(r.Body).Decode(&req)
    
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(fmt.Sprintf(`{"message": "Error while unmarshaling register request", "details": "%s"}`, err.Error())))
        // TODO log
        return
    }

	if req.Name == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"message": "Malformed register request", "details": "Name field is empty"}`))
        // TODO log
        return
    }

    if req.Surname == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"message": "Malformed register request", "details": "Surname field is empty"}`))
        // TODO log
        return
	}
	
	if req.Login == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"message": "Malformed register request", "details": "Login field is empty"}`))
        // TODO log
        return
	}
	
	if req.Password == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"message": "Malformed register request", "details": "Password field is empty"}`))
        // TODO log
        return
	}

    vrcres, err := ctx.validateRegisterCredentials(req)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(fmt.Sprintf(`{"message": "Error while validating register credentials", "details": "%s"}`, err.Error())))
        // TODO log
        return
    }

    if vrcres != "" {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte(fmt.Sprintf(`{"message": "Invalid register credentials", "details": "%s"}`, vrcres)))
        // TODO log
        return
    }

    res, err = ctx.registerUser(req)
        
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(fmt.Sprintf(`{"message": "Error while registering user", "details": "%s"}`, err.Error())))
        // TODO log
        return
	}
	
    out, err := json.Marshal(&res)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(fmt.Sprintf(`{"message": "Error while marshaling response", "details": "%s"}`, err.Error())))
        // TODO log
        return
    } 

    w.Write([]byte(out))
    // TODO log
    return
}
