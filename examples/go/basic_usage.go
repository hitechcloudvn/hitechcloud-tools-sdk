#!/usr/bin/env go
/**
 * Basic usage example for HiTechCloud Tools Go SDK / Ví dụ sử dụng cơ bản với HiTechCloud Tools Go SDK
 *
 * Set your API key: / Đặt API key của bạn:
 *   export HITECHCLOUD_TOOLS_API_KEY=your-api-key
 */
package main

import (
    "github.com/hitechcloudvn/hitechcloud-tools-sdk/sdk/go"
)

func main() {
    client := hitechcloudtools.NewClient("your-api-key")
}

// GeoIP lookup / Tra cứu GeoIP
result, err := client.MangAndHaTang.GeoipLookup("8.8.8.8")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)

// Domain WHOIS / WHOIS tên miền
whois, err := client.TenMienAndSsl.DomainWhois("example.com")
if err != nil {
    log.Fatal(err)
}
fmt.Println(whois)
