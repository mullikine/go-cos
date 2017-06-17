package cos

import (
	"context"
	"encoding/xml"
	"net/http"
)

// BucketCORSRule ...
type BucketCORSRule struct {
	ID             string   `xml:"ID,omitempty"`
	AllowedMethods []string `xml:"AllowedMethod"`
	AllowedOrigins []string `xml:"AllowedOrigin"`
	AllowedHeaders []string `xml:"AllowedHeader,omitempty"`
	MaxAgeSeconds  int      `xml:"MaxAgeSeconds,omitempty"`
	ExposeHeaders  []string `xml:"ExposeHeader,omitempty"`
}

// BucketGetCORSResult ...
type BucketGetCORSResult struct {
	XMLName xml.Name          `xml:"CORSConfiguration"`
	Rules   []*BucketCORSRule `xml:"CORSRule,omitempty"`
}

// GetCORS Get Bucket CORS实现跨域访问配置读取。
//
// https://www.qcloud.com/document/product/436/8274
func (s *BucketService) GetCORS(ctx context.Context,
	authTime *AuthTime) (*BucketGetCORSResult, *Response, error) {
	u := "/?cors"
	baseURL := s.client.BaseURL.BucketURL
	var res BucketGetCORSResult
	resp, err := s.client.sendNoBody(ctx, baseURL, u, http.MethodGet, authTime, nil, nil, &res)
	return &res, resp, err
}

type BucketPutCORSOptions struct {
	XMLName xml.Name          `xml:"CORSConfiguration"`
	Rules   []*BucketCORSRule `xml:"CORSRule,omitempty"`
}

// PutCORS Put Bucket CORS实现跨域访问设置，您可以通过传入XML格式的配置文件实现配置，文件大小限制为64 KB。
//
// https://www.qcloud.com/document/product/436/8279
func (s *BucketService) PutCORS(ctx context.Context,
	authTime *AuthTime, opt *BucketPutCORSOptions) (*Response, error) {
	u := "/?cors"
	baseURL := s.client.BaseURL.BucketURL
	resp, err := s.client.sendWithBody(ctx, baseURL, u, http.MethodPut, authTime, opt, nil, nil, nil)
	return resp, err
}

// DeleteCORS Delete Bucket CORS实现跨域访问配置删除。
//
// https://www.qcloud.com/document/product/436/8283
func (s *BucketService) DeleteCORS(ctx context.Context,
	authTime *AuthTime) (*Response, error) {
	u := "/?cors"
	baseURL := s.client.BaseURL.BucketURL
	resp, err := s.client.sendNoBody(ctx, baseURL, u, http.MethodDelete, authTime, nil, nil, nil)
	return resp, err
}
