package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TurnstileVerifyRequest Turnstile 验证请求
type TurnstileVerifyRequest struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
}

// TurnstileVerifyResponse Turnstile 验证响应
type TurnstileVerifyResponse struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}

// VerifyTurnstile 验证 Turnstile token
func VerifyTurnstile(secretKey string, token string) (bool, error) {
	if secretKey == "" || token == "" {
		return false, fmt.Errorf("secret key or token is empty")
	}

	req := TurnstileVerifyRequest{
		Secret:   secretKey,
		Response: token,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(
		"https://challenges.cloudflare.com/turnstile/v0/siteverify",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var verifyResp TurnstileVerifyResponse
	if err := json.Unmarshal(body, &verifyResp); err != nil {
		return false, err
	}

	return verifyResp.Success, nil
}
