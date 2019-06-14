package hg8347r

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"
)

// Router instance
type Router struct {
	URL string
	Req *req
}

// Device info
type Device struct {
	HostName    string
	IPAddr      string `json:"IpAddr"`
	MacAddr     string
	PortType    string
	PortID      string
	TrafficSend string
	TrafficRecv string
	DevStatus   string
	IPType      string `json:"IpType"` // DHCP or STATIC
	Time        string // online time, format is mm:ss
}

// New create a new *Router
func New(URL, username, password string) *Router {
	client := newReq()
	token := client.Post(URL + "/asp/GetRandCount.asp").String()
	// remove first character \ufeff
	token = token[len(token)-32:]
	pass := base64.StdEncoding.EncodeToString([]byte(password))
	header := http.Header{"Cookie": {"Cookie=body:Language:chinese:id=-1; MenuJumpIndex=0"}}
	body := url.Values{"UserName": {username}, "PassWord": {pass}, "x.X_HW_Token": {token}}
	client.Post(URL+"/login.cgi", header, body)
	return &Router{URL: URL, Req: client}
}

// Logout sign out
func (r *Router) Logout() {
	r.Req.Post(r.URL + "/logout.cgi?RequestFile=html/logout.html")
}

// ListDevices list devices connected to router
func (r *Router) ListDevices() []Device {
	page := r.Req.Get(r.URL + "/html/bbsp/userdevinfo/userdevinfolan.asp").String()
	regex := regexp.MustCompile(`<input type="hidden" name="onttoken" id="hwonttoken" value="(\w{32})">`)
	token := regex.FindStringSubmatch(page)[1]
	body := url.Values{
		"HostName":     {""},
		"IpAddr":       {""},
		"MacAddr":      {""},
		"PortType":     {""},
		"PortID":       {""},
		"TrafficSend":  {""},
		"TrafficRecv":  {""},
		"DevStatus":    {""},
		"time":         {""},
		"IpType":       {""},
		"x.X_HW_Token": {token},
	}
	urlSuffix := "/getajax.cgi?x=InternetGatewayDevice.LANDevice.1.X_HW_UserDev.{i}&RequestFile=html/bbsp/userdevinfo/userdevinfolan.asp"
	s := r.Req.Post(r.URL+urlSuffix, body).EscapeString()
	// fmt.Println(s)
	devices := []Device{}
	json.Unmarshal([]byte(s), &devices)
	// last one is {"result": 0}, which should be removed
	devices = devices[:len(devices)-1]
	return devices
}
