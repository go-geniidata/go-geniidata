package utils

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func GetWithHeader(ctx context.Context, url string, header map[string]string, response interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, response); err != nil {
		return err
	}
	return nil
}

func GetWithKey(ctx context.Context, url, apiKey string, response interface{}) error {
	header := make(map[string]string)
	header["Accept"] = "application/json"
	header["api-key"] = apiKey
	return GetWithHeader(ctx, url, header, response)
}
