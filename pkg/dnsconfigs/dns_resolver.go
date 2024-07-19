// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package dnsconfigs

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"github.com/iwind/TeaGo/types"
)

type DNSResolver struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

/*
// DNS解析设计 #2
type DNSResolver struct {
	Host	string `json:"host"` // DNS Host
 	Port	int	`json:"port"` // 端口
  	Protocol string `json:"protocol"` // 协议白名单 tls udp https
   	TLSPort	int	`json:"tlsPort"` // TLS 端口
    	DefaultUDP	int `json:"defaultUDP"` // 回滚端口
}

// 添加 error 强制规范化
func (this *DNSResolver) Addr() (string, error) {
    // 不回应空白域
    if this.Host == "" {
        return "", errors.New("host cannot be empty")
    }
    // 不回应非白名单协议
    if !isValidProtocol(this.Protocol) {
        return "", errors.New("unsupported protocol: " + this.Protocol)
    }
    // 端口切换
    var port = this.Port
    if port <= 0 || port > 65535 {
        switch this.Protocol {
        case "tls":
            if this.TLSPort > 0 && this.TLSPort <= 65535 {
                port = this.TLSPort
            } else {
                port = 853 // Default TLS port
            }
        default:
            port = 53 // Default UDP port
        }
    }

    if port == 853 && this.DefaultUDP > 0 && this.DefaultUDP <= 65535 {
        port = this.DefaultUDP // UDP853回滚端口
    }

    return configutils.QuoteIP(this.Host) + ":" + types.String(port), nil
}
*/

func (this *DNSResolver) Addr() string {
	var port = this.Port
	if port <= 0 {
		// 暂时不支持DoH
		// 实际应用中只支持udp
		switch this.Protocol {
		case "tls":
			port = 853
		default:
			port = 53
		}
	}
	return configutils.QuoteIP(this.Host) + ":" + types.String(port)
}
