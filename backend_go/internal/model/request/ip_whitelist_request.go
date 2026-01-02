package request

type IPWhitelistRequest struct {
	IPs []string `json:"ips"`
}
