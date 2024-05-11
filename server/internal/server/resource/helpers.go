package resource

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"vemn/configs/serverConf"
	"vemn/internal/server/postgresql/helpers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

func checkUserInDB(query string, args []any, password string) bool {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)

	defer rows.Close()

	us := UserAuth{}

	for rows.Next() {
		p := UserAuth{}
		err = rows.Scan(
			&p.ID,
			&p.Email,
			&p.Password,
			&p.UserSalt,
			&p.Access,
		)

		if err != nil {
			log.Println(err)
			return false
		}
		us = UserAuth{
			p.ID,
			p.Email,
			p.Password,
			p.Access,
			p.UserSalt,
		}

		hash := sha1.New()
		hash.Write([]byte(password))

		hashedPassword := fmt.Sprintf("%x", hash.Sum([]byte(p.UserSalt)))

		if p.Password != hashedPassword {
			log.Println("invalid password!")
			return false
		}

	}
	if us.Email == "" {
		log.Println("Email is empty!")
		return false
	}
	return true
}

func getUserAccessInDB(query string, args []any) string {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()

	us := UserAuth{}

	for rows.Next() {
		p := UserAuth{}
		err = rows.Scan(
			&p.Access,
		)

		if err != nil {
			log.Println(err)
			return ""
		}
		us = UserAuth{
			Access: p.Access,
		}
	}

	return us.Access
}

func getFioInDB(query string, args []any) string {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()

	us := UserData{}

	for rows.Next() {
		p := UserData{}
		err = rows.Scan(
			&p.FIO,
		)

		if err != nil {
			log.Println(err)
			return ""
		}
		us = UserData{
			FIO: p.FIO,
		}
	}

	return us.FIO

}

func checkResourceInDB(args []any) bool {
	rows, err := helpers.Select("select * from url where nameurl = $1", args, serverConf.DefaultConfig)
	defer rows.Close()

	res := UrlTable{}

	for rows.Next() {
		p := UrlTable{}
		err = rows.Scan(
			&p.ID,
			&p.URL,
			&p.IP,
			&p.Err,
			&p.ErrBool,
			&p.WafAddress,
			&p.WafBool,
			&p.CertBool,
			&p.CommonName,
			&p.Issuer,
			&p.EndDate,
			&p.FIO,
			&p.OwnerShortName,
		)

		if err != nil {
			log.Println(err)
			continue
		}
		res = UrlTable{
			ID: p.ID,
		}
	}
	if res.ID.Int32 == 0 {
		log.Println("ID is empty!")
		return false
	}
	return true
}

func checkData(args URL) CheckDataResult {
	user := UserData{}
	owner := NewOwner{}
	if args.Email != "" {
		user = getUserData(args.Email)
		if user.ID == 0 {
			return CheckDataResult{Result: false}
		}
	}
	if args.Owner != "" {
		owner = getOwnerData(args.Owner)
		if owner.ID == 0 {
			return CheckDataResult{Result: false}
		}
	}
	if checkResourceInDB([]any{args.Url}) == true {
		return CheckDataResult{Result: false}
	}
	if user.ID != 0 && owner.ID != 0 {
		return CheckDataResult{UserID: user.ID, OwnerId: owner.ID, Result: true}
	}
	if user.ID != 0 && owner.ID == 0 {
		return CheckDataResult{UserID: user.ID, OwnerId: owner.ID, Result: true}
	}
	if user.ID == 0 && owner.ID != 0 {
		return CheckDataResult{UserID: user.ID, OwnerId: owner.ID, Result: true}
	}
	return CheckDataResult{UserID: user.ID, OwnerId: owner.ID, Result: true}
}

func toJson(variable any) []byte {
	jsonStruct, err := json.Marshal(variable)
	if err != nil {
		log.Println("Error: ", err.Error())
	}
	return jsonStruct
}

func findWeeks(today time.Time) Weeks {
	mondayDurationDays := int(today.Weekday()) - 1
	fridayDurationDays := 5 - int(today.Weekday())

	return Weeks{
		Last: Week{
			Monday: today.Add(-(time.Duration(24*(mondayDurationDays+7)) * time.Hour)),
			Friday: today.Add(-(time.Duration(24*(7-fridayDurationDays)) * time.Hour)),
		},
		Current: Week{
			Monday: today.Add(-(time.Duration(24*mondayDurationDays) * time.Hour)),
			Friday: today.Add(time.Duration(24*fridayDurationDays) * time.Hour),
		},
	}

}

func findMonth() Months {
	today := time.Now()
	month := ""

	month = strconv.Itoa(int(today.Month()))
	if len(month) < 2 {
		month = "0" + month
	}

	day := strconv.Itoa(today.Day())
	if len(day) < 2 {
		day = "0" + day
	}

	return Months{
		strconv.Itoa(today.Year()) + "-" + month + "-" + day,
		today.AddDate(0, 2, 0).Format("2006-01-02"),
	}

}

func collector(args []any) WeekStatistic {
	NoResolve, arrayNoResolve := counter("select resource.idurl, url.nameurl, ipfirst, ipnow, datefirst, datenores, wafdate from resource join url on resource.idurl=url.idurl where datenores between $1 and $2", args)
	NewWaf, arrayNewWaf := counter("select resource.idurl, url.nameurl, ipfirst, ipnow, datefirst, datenores, wafdate from resource join url on resource.idurl=url.idurl where wafdate between $1 and $2", args)
	return WeekStatistic{
		NoResolve:      NoResolve,
		NewWaf:         NewWaf,
		NoResResource:  arrayNoResolve,
		NewWafResource: arrayNewWaf,
	}
}

func counter(query string, args []any) (int, []WeekStatisticResource) {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()
	resources := []WeekStatisticResource{}

	for rows.Next() {
		p := ResourceTable{}
		err = rows.Scan(
			&p.ID,
			&p.NameURL,
			&p.IpFirst,
			&p.IpNow,
			&p.DateFirst,
			&p.DateNoRes,
			&p.WafDate,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		if p.DateNoRes.String == "" {
			resources = append(resources, WeekStatisticResource{
				p.NameURL.String,
				p.WafDate.String,
			})
		} else {
			resources = append(resources, WeekStatisticResource{
				p.NameURL.String,
				p.DateNoRes.String,
			})
		}

	}
	return len(resources), resources
}

func getUserData(email string) UserData {
	rows, err := helpers.Select("select * from usdata where emailus = $1", []any{email}, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Println("error: ", err)
		return UserData{}
	}

	us := UserData{}
	for rows.Next() {
		p := User{}
		err = rows.Scan(
			&p.ID,
			&p.FIO,
			&p.Email,
			&p.Phone,
			&p.Telegram,
			&p.IdOwner,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		us = UserData{
			ID:       int(p.ID.Int32),
			Email:    p.Email.String,
			FIO:      p.FIO.String,
			Phone:    p.Phone.String,
			Telegram: p.Telegram.String,
			IdOwner:  int(p.IdOwner.Int32),
		}
	}
	return us
}

func getOwnerData(name string) NewOwner {
	rows, err := helpers.Select("select * from owners where nameowner = $1;", []any{name}, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Println("error: ", err)
		return NewOwner{}
	}

	owner := NewOwner{}
	for rows.Next() {
		p := Owner{}
		err = rows.Scan(
			&p.ID,
			&p.FullName,
			&p.ShortName,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		owner = NewOwner{
			ID:        int(p.ID.Int32),
			FullName:  p.FullName.String,
			ShortName: p.ShortName.String,
		}
	}
	return owner
}

func sortCertificates(month Months, data []Certificate) CertificateInfo {

	current := []Certificate{}
	next := []Certificate{}

	for i := 0; i < len(data); i++ {
		if string(data[i].Date[5])+string(data[i].Date[6]) == string(month.Current[5])+string(month.Current[6]) {
			current = append(current, data[i])
			continue
		}
		next = append(next, data[i])
	}
	return CertificateInfo{
		Current: current,
		Next:    next,
	}
}

func getDataFromResolver(url string, c *gin.Context) []AddResourceCollection {
	data := struct {
		UrlArray []Url `json:"urls"`
	}{}

	cli := &http.Client{}

	data.UrlArray = append(data.UrlArray, Url{Name: url})

	bytesRepresentation := toJson(data)

	req, err := http.NewRequest("POST", "https://localhost:8081/resolve", bytes.NewBuffer(bytesRepresentation))
	req.Header.Add("Authorization", c.GetHeader("Authorization"))
	resp, err := cli.Do(req)

	//resp, err := http.Post("https://localhost:8081/resolve", "application/json", bytes.NewBuffer(bytesRepresentation))

	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}

	resData := struct {
		Data []AddResourceCollection `json:"resolvedData"`
	}{}

	json.NewDecoder(resp.Body).Decode(&result)

	strs := result["body"].(interface{}).(string)

	json.Unmarshal([]byte(strs), &resData)

	return resData.Data
}

func generateToken(login string, fio string, access string) (string, error) {
	rows, err := helpers.Select("select * from users where emailus = $1;", []any{login}, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}

	for rows.Next() {
		p := UserAuth{}
		err = rows.Scan(
			&p.ID,
			&p.Email,
			&p.Password,
			&p.UserSalt,
			&p.Access,
		)

		if err != nil {
			log.Println(err)
			return "", err
		}
		us := UserAuth{
			p.ID,
			p.Email,
			p.Password,
			p.Access,
			p.UserSalt,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			login, us.Password, fio, access,
		})

		return token.SignedString([]byte(signingKey))
	}

	return "", err
}

func parseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.FIO, nil
}

func validateURL(url string) bool {
	if strings.Contains(url, "http://") ||
		strings.Contains(url, "https://") ||
		strings.Contains(url, "www.") ||
		strings.Contains(url, "/") ||
		strings.Contains(url, ":") ||
		strings.HasPrefix(url, ".") ||
		!strings.HasSuffix(url, ".ru") {
		return false
	}

	return true
}

func updateStatsTable() ResponseStatistic {
	stats := ResponseStatistic{}

	rows, err := helpers.Select("select count(*) from stat;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&stats.ID); err != nil {
			fmt.Println(err)
			return ResponseStatistic{}
		}
	}
	stats.ID = stats.ID + 1

	stats.Date, _ = time.Parse("2006-01-02", time.Now().Format("2006-01-02"))

	rows, err = helpers.Select("select count(*) from url;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&stats.AllServers); err != nil {
			fmt.Println(err)
			return ResponseStatistic{}
		}
	}

	rows, err = helpers.Select("select count(*) from url where errbool = false;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&stats.ErrorServers); err != nil {
			fmt.Println(err)
			return ResponseStatistic{}
		}
	}

	stats.WorkServers = stats.AllServers - stats.ErrorServers

	rows, err = helpers.Select("select count(*) from url where wafbool = true;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&stats.WithWaf); err != nil {
			fmt.Println(err)
			return ResponseStatistic{}
		}
	}

	stats.WafProc = float64((stats.WithWaf*10000)/stats.WorkServers) / 100

	rows, err = helpers.Select("select count(*) from url where wafbool = false;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&stats.Possible); err != nil {
			fmt.Println(err)
			return ResponseStatistic{}
		}
	}

	stats.WafProcPossible = float64(stats.Possible*10000/stats.WorkServers) / 100

	rows, err = helpers.Select("select count(*) from url where kdpbool = true;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&stats.WithKas); err != nil {
			fmt.Println(err)
			return ResponseStatistic{}
		}
	}

	rows, err = helpers.Select("select count(*) from url where kdpbool = true and wafbool = true;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&stats.WafAndKas); err != nil {
			fmt.Println(err)
			return ResponseStatistic{}
		}
	}

	stats.WafAndKasProc = float64((stats.WafAndKas*10000)/(stats.WorkServers+stats.WithKas)) / 100

	rows, err = helpers.Select("select count(*) from url where certbool = true;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&stats.AllCertificate); err != nil {
			fmt.Println(err)
			return ResponseStatistic{}
		}
	}

	month := findMonth()
	rows, err = helpers.Select("select count(*) from url where datecert between $1 and $2;", []any{month.Current, month.Next}, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&stats.OkCertificate); err != nil {
			fmt.Println(err)
			return ResponseStatistic{}
		}
	}

	rows, err = helpers.Select("select count(*) from owners;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&stats.Users); err != nil {
			fmt.Println(err)
			return ResponseStatistic{}
		}
	}

	return stats

}

func getDate() bool {
	var dateStr time.Time

	rows, err := helpers.Select("select datestat from stat order by idstat desc limit 1;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&dateStr); err != nil {
			fmt.Println(err)
			return false
		}
	}

	if time.Now().Month() == dateStr.Month() {
		// не нужно создавать новую запись
		return false
	} else {
		// нужно создавать новую запись
		return true
	}
}

func wallarmCollector(data []byte) float64 {

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.wallarm.ru/v1/objects/attack/count", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("err: ", err)
		return 0
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-WallarmApi-Token", "DZDnH8bd/uyQEmJUlTyexIN1cMrv4+/kwiStX3QfkZXEs26M0whNHKDjU7sdUtl2")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err: ", err)
		return 0
	}

	resData := struct {
		Body AttackStat `json:"body"`
	}{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err: ", err)
		return 0
	}

	err = json.Unmarshal(body, &resData)

	return resData.Body.Hits
}
