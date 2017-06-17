package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"bitbucket.org/mozillazg/go-cos"
)

func main() {
	u, _ := url.Parse("https://test-1253846586.cn-north.myqcloud.com")
	b := &cos.BaseURL{
		BucketURL: u,
	}
	c := cos.NewClient(os.Getenv("COS_SECRETID"), os.Getenv("COS_SECRETKEY"), b, nil)
	c.Client.Transport = &cos.DebugRequestTransport{
		RequestHeader:  true,
		RequestBody:    true,
		ResponseHeader: true,
		ResponseBody:   true,
	}

	//	`<CORSConfiguration>
	//	<CORSRule>
	//		<AllowedOrigin>www.qq.com</AllowedOrigin>
	//		<AllowedMethod>PUT</AllowedMethod>
	//		<MaxAgeSeconds>100</MaxAgeSeconds>
	//	</CORSRule>
	//</CORSConfiguration>
	//`
	opt := &cos.BucketPutCORSOptions{
		Rules: []*cos.BucketCORSRule{
			{
				//ID:            "1234",
				AllowedOrigins: []string{"http://www.qq.com"},
				AllowedMethods: []string{"PUT", "GET"},
				AllowedHeaders: []string{"x-cos-meta-test", "x-cos-xx"},
				MaxAgeSeconds:  500,
				ExposeHeaders:  []string{"x-cos-meta-test1"},
			},
		},
	}
	_, err := c.Bucket.PutCORS(context.Background(), cos.NewAuthTime(time.Hour), opt)
	if err != nil {
		fmt.Println(err)
	}
}
