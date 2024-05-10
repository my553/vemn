package resource

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Login struct {
	Login    string `json:"login" db:"emailus"`
	Password string `json:"password"`
}

type TokenClaims struct {
	jwt.StandardClaims
	Login    string `json:"login"`
	Password string `json:"password"`
	FIO      string `json:"fio"`
	Access   string `json:"access"`
}

type UserAuth struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Access   string `json:"access"`
	UserSalt string `json:"userSalt"`
}

type User struct {
	ID       sql.NullInt32  `json:"id"`
	FIO      sql.NullString `json:"FIO"`
	Email    sql.NullString `json:"email"`
	Phone    sql.NullString `json:"phone"`
	Telegram sql.NullString `json:"tg"`
	IdOwner  sql.NullInt32  `json:"idOwner"`
}

type UserData struct {
	ID       int    `json:"ID"`
	FIO      string `json:"fio"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Telegram string `json:"tg"`
	IdOwner  int    `json:"idOwner"`
}

type RequestStatistic struct {
	AllServers      int `json:"allServers"`
	ErrorServers    int `json:"errorServers"`
	WorkingServers  int `json:"workingServers"`
	WithWaf         int `json:"withWaf"`
	AllCertificates int `json:"allCertificates"`
	OkCertificates  int `json:"okCertificates"`
}

type ResponseStatistic struct {
	ID              int
	Date            time.Time
	AllServers      int
	ErrorServers    int
	WorkServers     int
	WithWaf         int
	Possible        int
	WafProcPossible float64
	WafProc         float64
	WithKas         int
	WafAndKas       int
	WafAndKasProc   float64
	AllCertificate  int
	OkCertificate   int
	Users           int
}

type Weeks struct {
	Last    Week
	Current Week
}

type Week struct {
	Monday time.Time
	Friday time.Time
}

type WeeksStatistic struct {
	Last    WeekStatistic `json:"last"`
	Current WeekStatistic `json:"current"`
}

type WeekStatistic struct {
	NoResolve      int                     `json:"no_resolve"`
	NewWaf         int                     `json:"new_waf"`
	NoResResource  []WeekStatisticResource `json:"no_res_resource"`
	NewWafResource []WeekStatisticResource `json:"new_waf_resource"`
}

type WeekStatisticResource struct {
	Resource string `json:"resource"`
	Date     string `json:"date"`
}

type URL struct {
	Url        string `json:"url"`
	Email      string `json:"email"`
	Owner      string `json:"owner"`
	InteriorIP string `json:"IPIn"`
	IsId       string `json:"isName"`
	Statement  string `json:"statement"`
}

type ResourceTable struct {
	ID        sql.NullInt32  `json:"ID"`
	NameURL   sql.NullString `json:"NameURL"`
	IpFirst   sql.NullString `json:"IpFirst"`
	IpNow     sql.NullString `json:"IpNow"`
	DateFirst sql.NullString `json:"DateFirst"`
	DateNoRes sql.NullString `json:"DateNoRes"`
	WafDate   sql.NullString `json:"WafDate"`
}

type UrlTable struct {
	ID             sql.NullInt32
	URL            sql.NullString
	IP             sql.NullString
	Err            sql.NullString
	ErrBool        sql.NullBool
	WafID          sql.NullString
	WafAddress     sql.NullString
	WafBool        sql.NullBool
	IdUser         sql.NullInt32
	IdOwner        sql.NullInt32
	CertBool       sql.NullBool
	CommonName     sql.NullString
	Issuer         sql.NullString
	EndDate        sql.NullString
	FIO            sql.NullString
	OwnerShortName sql.NullString
}

type Owner struct {
	ID        sql.NullInt32  `json:"id"`
	FullName  sql.NullString `json:"nameOwner"`
	ShortName sql.NullString `json:"shortName"`
}

type NewOwner struct {
	ID        int    `json:"id"`
	FullName  string `json:"nameOwner"`
	ShortName string `json:"shortName"`
}

type CheckDataResult struct {
	UserID  int
	OwnerId int
	Result  bool
}

type CheckResource struct {
	URL       string         `json:"URL"`
	IP        string         `json:"IP"`
	Status    bool           `json:"Status"`
	WAF       bool           `json:"WAF"`
	SSLStatus bool           `json:"SSLStatus"`
	SSL       UrlCertificate `json:"SSL"`
	UserID    int            `json:"UserID"`
	OwnerID   int            `json:"OwnerID"`
}

type UpdateData struct {
	Url   string `json:"url"`
	Email string `json:"email"`
}

type GeneralStat struct {
	Date               string `json:"date"`
	Resources          int    `json:"resources"`
	DeactivateResource int    `json:"deactivateResource"`
	Owners             int    `json:"owners"`
	Waf                int    `json:"waf"`
}

type CertificateInfo struct {
	Current []Certificate `json:"current"`
	Next    []Certificate `json:"next"`
}

type Certificate struct {
	Resource string `json:"resource"`
	Date     string `json:"date"`
}

type Months struct {
	Current string
	Next    string
}

type ResolveInfo struct {
	Ip          string  `json:"ip"`
	Status      string  `json:"status"`
	ErrStatus   bool    `json:"errStatus"`
	DateNoRes   string  `json:"dateNoRes"`
	WafDate     string  `json:"wafDate"`
	Waf         int     `json:"waf"`
	WafIp       *string `json:"wafIp"`
	NameUrl     string  `json:"nameurl"`
	WafStatus   bool    `json:"wafStatus"`
	KdpStatus   bool    `json:"kdpStatus"`
	WafPossible bool    `json:"wafPossible"`
}

type UrlCertificate struct {
	CommonName string `json:"common_name"`
	Issuer     string `json:"issuer"`
	DateCert   string `json:"date_cert"`
	CertStatus bool   `json:"certStatus"`
}

type AddResourceCollection struct {
	Resolve     ResolveInfo    `json:"resolve"`
	Certificate UrlCertificate `json:"certificate"`
}

type Url struct {
	Name string `json:"urlName"`
}

type AllStats struct {
	Resource Resource `json:"resource"`
}

type Resource struct {
	Name           string `json:"nameUrl"`
	IP             string `json:"ip"`
	ErrMsg         string `json:"errMessage"`
	ErrStatus      bool   `json:"errStatus"`
	WafIP          string `json:"wafIP"`
	WafStatus      bool   `json:"wafStatus"`
	CertStatus     bool   `json:"certStatus"`
	CommonName     string `json:"commonName"`
	Issuer         string `json:"issuer"`
	DateCert       string `json:"dateCert"`
	FIO            string `json:"fio"`
	OwnerShortName string `json:"ownerShortName"`
}

type SQLChart struct {
	Date        sql.NullString `json:"date"`
	WorkServers sql.NullString `json:"workServers"`
	ErServers   sql.NullString `json:"erServers"`
	WithWAF     sql.NullString `json:"withWAF"`
}

type Year struct {
	January   Month `json:"January"`
	February  Month `json:"February"`
	March     Month `json:"March"`
	April     Month `json:"April"`
	May       Month `json:"May"`
	June      Month `json:"June"`
	July      Month `json:"July"`
	August    Month `json:"August"`
	September Month `json:"September"`
	October   Month `json:"October"`
	November  Month `json:"November"`
	December  Month `json:"December"`
}

type Month struct {
	Chart Chart `json:"chart"`
}

type Chart struct {
	AllServers string `json:"allServers"`
	ErServers  string `json:"erServers"`
	WithWAF    string `json:"withWAF"`
}

type WAFStats struct {
	WithWAF int `json:"withWAF"`
	NoWAF   int `json:"noWAF"`
}

type AttackFilter struct {
	ClientID []int     `json:"clientid"`
	VulnID   *int      `json:"!vulnid"`
	Type     []string  `json:"!type"`
	Time     [][]int64 `json:"time"`
	State    string    `json:"!state"`
}

type HitsFilter struct {
	ClientID []int     `json:"clientid"`
	VulnID   *int      `json:"vulnid"`
	Type     []string  `json:"!type"`
	Time     [][]int64 `json:"time"`
	State    string    `json:"!state"`
}

type AttackStat struct {
	Attacks float64 `json:"attacks"`
	Hits    float64 `json:"hits"`
}

type DomainAttackFilter struct {
	ClientID int       `json:"clientid"`
	Time     [][]int64 `json:"time"`
}

type DomainAttackStat struct {
	Events    int     `json:"events"`
	Incidents int     `json:"incidents"`
	Trend     float64 `json:"trend"`
	Domain    string  `json:"domain"`
}

type AttackTypeFilter struct {
	ClientID int       `json:"clientid"`
	Time     [][]int64 `json:"time"`
}

type AttackTypeStat struct {
	Hits Hits `json:"hits"`
}

type Hits struct {
	Total Total `json:"total"`
}

type Total struct {
	ByTypes Types     `json:"by_types_and_protocols"`
	ByAuth  AuthTypes `json:"by_auth_protocols"`
}

type Types struct {
	Bola     Type `json:"bola"`
	Crlf     Type `json:"crlf"`
	Infoleak Type `json:"infoleak"`
	Ldapi    Type `json:"ldapi"`
	MailInj  Type `json:"mail_injection"`
	NoSQL    Type `json:"nosqli"`
	Ptrav    Type `json:"ptrav"`
	Rce      Type `json:"rce"`
	Scanner  Type `json:"scanner"`
	SQLI     Type `json:"sqli"`
	SSI      Type `json:"ssi"`
	SSRF     Type `json:"ssrf"`
	SSTI     Type `json:"ssti"`
	Vpatch   Type `json:"vpatch"`
	XSS      Type `json:"xss"`
	XXE      Type `json:"xxe"`
}

type Type struct {
	None      int `json:"none,omitempty"`
	Rest      int `json:"rest,omitempty"`
	Soap      int `json:"soap,omitempty"`
	WebFrom   int `json:"web-form,omitempty"`
	Websocket int `json:"websocket,omitempty"`
	TotalApi  int `json:"total_api_requests,omitempty"`
	Total     int `json:"total,omitempty"`
}

type AuthTypes struct {
	None   int `json:"none"`
	Cookie int `json:"cookie"`
	Basic  int `json:"basic"`
	JWT    int `json:"jwt"`
	Bearer int `json:"bearer"`
	Oauth2 int `json:"oauth2"`
	ApiKey int `json:"api-key"`
	Total  int `json:"total"`
}
