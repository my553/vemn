package resource

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"
	"vemn/configs/serverConf"
	"vemn/internal/server/postgresql/helpers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func (service *PgService) Login(c *gin.Context) {
	var data Login
	var access string
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	if !checkUserInDB("select * from users where emailus = $1;", []any{data.Login}, data.Password) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}

	fio := getFioInDB("select fio from usdata where emailus = $1;", []any{data.Login})

	access = getUserAccessInDB("select access from users where emailus = $1;", []any{data.Login})

	token, _ := generateToken(data.Login, fio, access)

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"fio":   fio,
		"token": token,
	})

}

func (service *PgService) GetShortStat(c *gin.Context) {
	rows, err := helpers.Select("select datestat, workservers, errservers, withwaf from stat order by idstat desc limit 6;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	chart := Year{}

	for rows.Next() {
		p := SQLChart{}
		err := rows.Scan(
			&p.Date,
			&p.WorkServers,
			&p.ErServers,
			&p.WithWAF,
		)

		if err != nil {
			continue
		}

		t, _ := time.Parse(time.RFC3339, p.Date.String)

		if t.Month().String() == "January" {
			chart.January = Month{Chart{AllServers: p.WorkServers.String, ErServers: p.ErServers.String, WithWAF: p.WithWAF.String}}
		}
		if t.Month().String() == "February" {
			chart.February = Month{Chart{AllServers: p.WorkServers.String, ErServers: p.ErServers.String, WithWAF: p.WithWAF.String}}
		}
		if t.Month().String() == "March" {
			chart.March = Month{Chart{AllServers: p.WorkServers.String, ErServers: p.ErServers.String, WithWAF: p.WithWAF.String}}
		}
		if t.Month().String() == "April" {
			chart.April = Month{Chart{AllServers: p.WorkServers.String, ErServers: p.ErServers.String, WithWAF: p.WithWAF.String}}
		}
		if t.Month().String() == "May" {
			chart.May = Month{Chart{AllServers: p.WorkServers.String, ErServers: p.ErServers.String, WithWAF: p.WithWAF.String}}
		}
		if t.Month().String() == "June" {
			chart.June = Month{Chart{AllServers: p.WorkServers.String, ErServers: p.ErServers.String, WithWAF: p.WithWAF.String}}
		}
		if t.Month().String() == "July" {
			chart.July = Month{Chart{AllServers: p.WorkServers.String, ErServers: p.ErServers.String, WithWAF: p.WithWAF.String}}
		}
		if t.Month().String() == "August" {
			chart.August = Month{Chart{AllServers: p.WorkServers.String, ErServers: p.ErServers.String, WithWAF: p.WithWAF.String}}
		}
		if t.Month().String() == "September" {
			chart.September = Month{Chart{AllServers: p.WorkServers.String, ErServers: p.ErServers.String, WithWAF: p.WithWAF.String}}
		}
		if t.Month().String() == "October" {
			chart.October = Month{Chart{AllServers: p.WorkServers.String, ErServers: p.ErServers.String, WithWAF: p.WithWAF.String}}
		}
		if t.Month().String() == "November" {
			chart.November = Month{Chart{AllServers: p.WorkServers.String, ErServers: p.ErServers.String, WithWAF: p.WithWAF.String}}
		}
		if t.Month().String() == "December" {
			chart.December = Month{Chart{AllServers: p.WorkServers.String, ErServers: p.ErServers.String, WithWAF: p.WithWAF.String}}
		}
	}

	rows, err = helpers.Select("select withwaf from stat order by idstat desc limit 1;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	stats := WAFStats{}
	for rows.Next() {
		if err = rows.Scan(&stats.WithWAF); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}

	rows, err = helpers.Select("select workservers-withwaf as nowaf from stat order by idstat desc limit 1;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&stats.NoWAF); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}

	data := struct {
		Year  Year     `json:"year"`
		Stats WAFStats `json:"wafStat"`
	}{
		Year:  chart,
		Stats: stats,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(data)),
	})
}

func (service *PgService) GetWeekStat(c *gin.Context) {
	days := findWeeks(time.Now())
	format := "2006-01-02"

	collections := WeeksStatistic{
		Last:    collector([]any{days.Last.Monday.Format(format), days.Last.Friday.Format(format)}),
		Current: collector([]any{days.Current.Monday.Format(format), days.Current.Friday.Format(format)}),
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(collections)),
	})
}

func (service *PgService) AddResource(c *gin.Context) {
	data := URL{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	if !validateURL(data.Url) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	result := checkData(data)
	if result.Result == false {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	resolvedResource := getDataFromResolver(data.Url, c)

	res := helpers.Exec("insert into vulnerability (crits, medium, light, inform, lastdate) values ($1,$2,$3,$4,$5)",
		[]any{
			0,
			0,
			0,
			0,
			time.Now().Format("2006-01-02"),
		},
		serverConf.DefaultConfig,
	)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	fmt.Println("insert into vuln")

	if resolvedResource[0].Resolve.WafStatus == true {
		resolvedResource[0].Resolve.WafDate = time.Now().Format("2006-01-02")
	}

	if resolvedResource[0].Resolve.ErrStatus == false {
		resolvedResource[0].Resolve.DateNoRes = time.Now().Format("2006-01-02")
	}

	if resolvedResource[0].Resolve.Status == "Ok" {
		resolvedResource[0].Resolve.Status = "true"
	} else {
		resolvedResource[0].Resolve.Status = "false"
	}

	res = helpers.Exec("INSERT INTO resource (ipfirst,ipnow,datefirst,datenores,wafdate) VALUES ($1,$2,$3,$4,$5)",
		[]any{
			resolvedResource[0].Resolve.Ip,
			resolvedResource[0].Resolve.Ip,
			time.Now().Format("2006-01-02"),
			resolvedResource[0].Resolve.DateNoRes,
			resolvedResource[0].Resolve.WafDate,
		},
		serverConf.DefaultConfig,
	)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}
	fmt.Println("insert into resource")

	res = helpers.Exec(
		"INSERT INTO url (nameurl,ip,interiorip,isid,err,errbool,wafid,wafbool,idusd,idowner,certbool,commonname,issuer,datecert, kdpbool,statement) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)",
		[]any{
			resolvedResource[0].Resolve.NameUrl,
			resolvedResource[0].Resolve.Ip,
			data.InteriorIP, //interior IP
			0,               //infsystem id
			resolvedResource[0].Resolve.Status,
			resolvedResource[0].Resolve.ErrStatus,
			resolvedResource[0].Resolve.Waf,
			resolvedResource[0].Resolve.WafStatus,
			result.UserID,
			result.OwnerId,
			resolvedResource[0].Certificate.CertStatus,
			resolvedResource[0].Certificate.CommonName,
			resolvedResource[0].Certificate.Issuer,
			resolvedResource[0].Certificate.DateCert,
			resolvedResource[0].Resolve.KdpStatus,
			data.Statement,
		},
		serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}
	fmt.Println("insert into url")

	stats := updateStatsTable()
	if getDate() == true {
		res = helpers.Exec("INSERT INTO stat (idstat, datestat, allservers, errservers, workservers, withwaf, wafproc, possible, wafprocpossible, withkas, wafandkas, wafandkasproc, allcertificates, okcertificates, owners) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)",
			[]any{
				stats.ID,
				stats.Date,
				stats.AllServers,
				stats.ErrorServers,
				stats.WorkServers,
				stats.WithWaf,
				stats.WafProc,
				stats.Possible,
				stats.WafProcPossible,
				stats.WithKas,
				stats.WafAndKas,
				stats.WafAndKasProc,
				stats.AllCertificate,
				stats.OkCertificate,
				stats.Users,
			},
			serverConf.DefaultConfig,
		)
		if !res {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
			})
			return
		}
	}
	if !getDate() {
		res = helpers.Exec("update stat set datestat=$1, allservers=$2, errservers=$3, workservers=$4, withwaf=$5, wafproc=$6, possible=$7, wafprocpossible=$8, withkas=$9, wafandkas=$10, wafandkasproc=$11, allcertificates=$12, okcertificates=$13, owners=$14 where idstat = $15",
			[]any{
				stats.Date,
				stats.AllServers,
				stats.ErrorServers,
				stats.WorkServers,
				stats.WithWaf,
				stats.WafProc,
				stats.Possible,
				stats.WafProcPossible,
				stats.WithKas,
				stats.WafAndKas,
				stats.WafAndKasProc,
				stats.AllCertificate,
				stats.OkCertificate,
				stats.Users,
				stats.ID - 1,
			},
			serverConf.DefaultConfig,
		)
		if !res {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (service *PgService) CheckResource(c *gin.Context) {
	var data string
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	if !checkResourceInDB([]any{data}) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}

	rows, err := helpers.Select("select * from url where nameurl = $1", []any{data}, serverConf.DefaultConfig)
	defer rows.Close()

	if err != nil {
		fmt.Println("err")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	resourceStr := CheckResource{}

	for rows.Next() {
		p := UrlTable{}
		err := rows.Scan(
			&p.ID,
			&p.URL,
			&p.IP,
			&p.Err,
			&p.ErrBool,
			&p.WafID,
			&p.WafBool,
			&p.IdUser,
			&p.IdOwner,
			&p.CertBool,
			&p.CommonName,
			&p.Issuer,
			&p.EndDate,
		)

		if err != nil {
			fmt.Println(err)
			continue
		}
		resourceStr = CheckResource{
			URL:       p.URL.String,
			IP:        p.IP.String,
			Status:    p.ErrBool.Bool,
			WAF:       p.WafBool.Bool,
			SSLStatus: p.CertBool.Bool,
			SSL:       UrlCertificate{CommonName: p.CommonName.String, Issuer: p.Issuer.String, DateCert: p.EndDate.String, CertStatus: p.CertBool.Bool},
			UserID:    int(p.IdUser.Int32),
			OwnerID:   int(p.IdOwner.Int32),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(resourceStr)),
	})
}

func (service *PgService) DeleteResource(c *gin.Context) {
	var data string
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	if !checkResourceInDB([]any{data}) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}

	var id int

	rows, err := helpers.Select("select idurl from url where nameurl = $1;", []any{data}, serverConf.DefaultConfig)
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			fmt.Println(err)
			return
		}
	}

	res := helpers.Exec("delete from vulnerability where idurl = $1", []any{id}, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
	}
	res = helpers.Exec("delete from resource where idurl = $1", []any{id}, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
	}
	res = helpers.Exec("delete from url where idurl = $1", []any{id}, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": res,
	})

}

func (service *PgService) UpdateResource(c *gin.Context) {
	var data UpdateData
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	if !validateURL(data.Url) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	user := getUserData(data.Email)

	res := helpers.Exec("update url set idusd = $1 where nameurl = $2", []any{user.ID, data.Url}, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (service *PgService) GetGeneralStat(c *gin.Context) {
	var stats GeneralStat
	var owners []NewOwner
	var users []UserData

	var date sql.NullTime
	rows, err := helpers.Select("select datestat from stat order by idstat desc limit 1;", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&date); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}
	stats.Date = date.Time.Format("2006-01-02")

	rows, err = helpers.Select("select allservers from stat order by idstat desc limit 1;", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&stats.Resources); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}

	rows, err = helpers.Select("select owners from stat order by idstat desc limit 1;", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&stats.Owners); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}

	rows, err = helpers.Select("select withwaf from stat order by idstat desc limit 1;", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&stats.Waf); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}

	rows, err = helpers.Select("select errservers from stat order by idstat desc limit 1;", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&stats.DeactivateResource); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}

	rows, err = helpers.Select("select * from usdata;", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		p := User{}
		err := rows.Scan(
			&p.ID,
			&p.FIO,
			&p.Email,
			&p.Phone,
			&p.Telegram,
			&p.IdOwner,
		)

		if err != nil {
			continue
		}

		users = append(users, UserData{ID: int(p.ID.Int32), FIO: p.FIO.String, Email: p.Email.String, Phone: p.Phone.String, Telegram: p.Telegram.String, IdOwner: int(p.IdOwner.Int32)})
	}

	rows, err = helpers.Select("select * from owners;", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		p := Owner{}
		err := rows.Scan(
			&p.ID,
			&p.FullName,
			&p.ShortName,
		)

		if err != nil {
			continue
		}

		owners = append(owners, NewOwner{ID: int(p.ID.Int32), FullName: p.FullName.String, ShortName: p.ShortName.String})
	}

	data := struct {
		Owners []NewOwner `json:"owners"`
		Users  []UserData `json:"users"`
	}{
		Owners: owners,
		Users:  users,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"body":  string(toJson(stats)),
		"cache": string(toJson(data)),
	})

}

func (service *PgService) GetCertificates(c *gin.Context) {
	month := findMonth()
	rows, err := helpers.Select("select * from url where datecert between $1 and $2;", []any{month.Current, month.Next}, serverConf.DefaultConfig)
	defer rows.Close()

	if err != nil {
		fmt.Println("err")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	certificates := []Certificate{}

	for rows.Next() {
		p := UrlTable{}
		err := rows.Scan(
			&p.ID,
			&p.URL,
			&p.IP,
			&p.Err,
			&p.ErrBool,
			&p.WafID,
			&p.WafBool,
			&p.IdUser,
			&p.IdOwner,
			&p.CertBool,
			&p.CommonName,
			&p.Issuer,
			&p.EndDate,
		)

		if err != nil {
			fmt.Println(err)
			continue
		}
		certificates = append(certificates, Certificate{p.URL.String, p.EndDate.String[0:10]})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(sortCertificates(month, certificates))),
	})
}

func (service *PgService) UserIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		fmt.Println("empty auth header")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 1 {
		fmt.Println("invalid auth string")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}

	userFIO, err := parseToken(headerParts[0])
	if err != nil {
		fmt.Println("invalid auth string")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
		})
		return
	}

	c.Set("userFIO", userFIO)
}

func (service *PgService) GetStatistic(c *gin.Context) {
	rows, err := helpers.Select("select idurl, nameurl,ip,err,errbool,waf.address,wafbool,certbool,commonname,issuer,datecert, usdata.fio, owners.shortname from url join waf on url.wafid=waf.id join usdata on url.idusd=usdata.idusd join owners on url.idowner=owners.id order by idurl;", nil, serverConf.DefaultConfig)
	defer rows.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	stats := []AllStats{}

	for rows.Next() {
		p := UrlTable{}
		err := rows.Scan(
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
			fmt.Println(err)
		}

		rows, err := helpers.Select("select crits,medium,light,inform,date from vulnerability where idurl=$1;", []any{p.ID.Int32}, serverConf.DefaultConfig)
		defer rows.Close()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
			})
			return
		}

		vulns := Vulnerabilites{}

		for rows.Next() {
			v := VulnTable{}
			err := rows.Scan(
				&v.Crits,
				&v.Medium,
				&v.Light,
				&v.Inform,
				&v.Date,
			)
			if err != nil {
				fmt.Println(err)
			}

			vulns = Vulnerabilites{Crits: int(v.Crits.Int32), Medium: int(v.Medium.Int32), Light: int(v.Light.Int32), Inform: int(v.Inform.Int32), Date: v.Date.String}

		}

		stats = append(stats, AllStats{Resource{Name: p.URL.String, IP: p.IP.String, ErrMsg: p.Err.String, ErrStatus: p.ErrBool.Bool, WafStatus: p.WafBool.Bool, WafIP: p.WafAddress.String, CertStatus: p.CertBool.Bool, CommonName: p.CommonName.String, Issuer: p.Issuer.String, DateCert: p.EndDate.String, Vulnerabilites: vulns, FIO: p.FIO.String, OwnerShortName: p.OwnerShortName.String}})
	}

	data := struct {
		Stats []AllStats `json:"stats"`
	}{
		Stats: stats,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(data)),
	})
}

func (service *PgService) GetUserInfo(c *gin.Context) {
	fio, _ := c.Get("userFIO")

	data := struct {
		FIO any `json:"userFIO"`
	}{
		FIO: fio,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(data)),
	})
}

func (service *PgService) AddOwner(c *gin.Context) {
	owner := NewOwner{}
	err := c.BindJSON(&owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	rows, err := helpers.Select("select count(*) from owners;", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&owner.ID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}

	res := helpers.Exec("INSERT INTO owners (id, nameowner, shortname) VALUES ($1, $2, $3)",
		[]any{
			owner.ID + 1,
			owner.FullName,
			owner.ShortName,
		},
		serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (service *PgService) AddUsData(c *gin.Context) {
	usdata := UserData{}
	err := c.BindJSON(&usdata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	rows, err := helpers.Select("select count(*) from usdata;", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&usdata.ID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}

	res := helpers.Exec("INSERT INTO usdata (idusd, fio, emailus, phoneus, tgus, idowner) VALUES ($1, $2, $3, $4, $5, $6)",
		[]any{
			usdata.ID + 1,
			usdata.FIO,
			usdata.Email,
			usdata.Phone,
			usdata.Telegram,
			usdata.IdOwner,
		},
		serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (service *PgService) ChangeUser(c *gin.Context) {
	data := struct {
		UrlName string `json:"urlName"`
		UserID  int    `json:"userID"`
		OwnerID int    `json:"ownerID"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	res := helpers.Exec("update url set idusd = $1, idowner = $2 where nameurl = $3;",
		[]any{
			data.UserID,
			data.OwnerID,
			data.UrlName,
		},
		serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}
