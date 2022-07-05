package fofa

import (
	"fmt"
	fofa_model "github.com/EwanSunn/secScan/internal/pkg/model/fofa_model"
	"github.com/EwanSunn/secScan/internal/pkg/slog"
	"github.com/EwanSunn/secScan/internal/pkg/util/color"
	"github.com/EwanSunn/secScan/internal/pkg/util/encode"
	"regexp"
	"strconv"
	"strings"
)

var fofa *fofa_model.Fofa

func initConfig(email, key string, size int) {
	fofa = fofa_model.New(email, key)
	fofa.SetSize(size)
}

func Run(email, key, target string, size int) {
	initConfig(email, key, size)
	keywords := loadFofa(target)
	for _, keyword := range keywords {
		slog.Infof("本次搜索关键字为：%v", keyword)
		size, results := fofa.Search(keyword)
		displayResponse(results)
		slog.Infof("本次搜索，返回结果总条数为：%d，此次返回条数为：%d", size, len(results))
	}
}

func displayResponse(results []fofa_model.FofaResult) {
	for _, row := range results {
		Fix(&row)
		m := row.CreateMap()
		m["Header"] = ""
		m["Cert"] = ""
		m["Title"] = row.Title
		m["Host"] = ""
		m["As_organization"] = ""
		m["Ip"] = ""
		m["Port"] = ""
		m["Country_name"] = ""
		m = encode.FixMap(m)
		if m["Banner"] != "" {
			m["Banner"] = encode.FixLine(m["Banner"])
			m["Banner"] = encode.StrRandomCut(m["Banner"], 20)
		}

		line := fmt.Sprintf("%-30v %-"+strconv.Itoa(encode.AutoWidth(row.Title, 26))+"v %v",
			row.Host,
			row.Title,
			color.StrMapRandomColor(m, false, []string{"Server"}, []string{}),
		)
		fmt.Println(line)
	}
}

func Fix(r *fofa_model.FofaResult) {
	//修复title
	if r.Title == "" && r.Protocol != "" {
		r.Title = strings.ToUpper(r.Protocol)
	}
	r.Title = encode.FixLine(r.Title)
	//修改host
	if r.Host == "" {
		r.Host = r.Ip
	}

	if regexp.MustCompile("\\w+://.*").MatchString(r.Host) == false {
		if r.Host == "" {
			r.Protocol = "http"
		}
		r.Host = r.Protocol + "://" + r.Host
	}
}

func loadFofa(target string) []string {
	//对象是否为多个
	if strArr := strings.ReplaceAll(target, "\\,", "[DouHao]"); strings.Count(strArr, ",") > 0 {
		var passArr []string
		for _, str := range strings.Split(strArr, ",") {
			passArr = append(passArr, strings.ReplaceAll(str, "[DouHao]", ","))
		}
		return passArr
	}
	//对象为单个且不为空时直接返回
	if target != "" {
		return []string{target}
	}
	return []string{}
}

func ShowSyntax() {
	const syntax = `
	title="beijing"			从标题中搜索"北京"			-
	header="elastic"		从http头中搜索"elastic"			-
	body="网络空间测绘"		从html正文中搜索"网络空间测绘"		-
	domain="qq.com"			搜索根域名带有qq.com的网站。		-
	icp="京ICP证030173号"		查找备案号为"京ICP证030173号"的网站	搜索网站类型资产
	js_name="js/jquery.js"		查找包含js/jquery.js的资产		搜索网站类型资产
	js_md5="82ac3f14327a8b7ba49baa208d4eaa15"	查找js源码与之匹配的资产	-
	icon_hash="-247388890"		搜索使用此icon的资产。			仅限FOFA高级会员使用
	host=".gov.cn"			从url中搜索".gov.cn"			搜索要用host作为名称
	port="6379"			查找对应"6379"端口的资产		-
	ip="1.1.1.1"			从ip中搜索包含"1.1.1.1"的网站		搜索要用ip作为名称
	ip="220.181.111.1/24"		查询IP为"220.181.111.1"的C网段资产	-
	status_code="402"		查询服务器状态为"402"的资产		-
	protocol="quic"			查询quic协议资产			搜索指定协议类型(在开启端口扫描的情况下有效)
	country="CN"			搜索指定国家(编码)的资产。		-
	region="Xinjiang"		搜索指定行政区的资产。			-
	city="Changsha"			搜索指定城市的资产。			-
	cert="baidu"			搜索证书中带有baidu的资产。		-
	cert.subject="Oracle"		搜索证书持有者是Oracle的资产		-
	cert.issuer="DigiCert"		搜索证书颁发者为DigiCert Inc的资产	-
	cert.is_valid=true		验证证书是否有效			仅限FOFA高级会员使用
	type=service			搜索所有协议资产			搜索所有协议资产
	os="centos"			搜索CentOS资产。			-
	server=="Microsoft-IIS"		搜索IIS 10服务器。			-
	app="Oracle"			搜索Microsoft-Exchange设备		-
	after="2017" && before="2017-10-01"	时间范围段搜索			-
	asn="19551"			搜索指定asn的资产。			-
	org="Amazon.com, Inc."	搜索指定org(组织)的资产。			-
	base_protocol="udp"		搜索指定udp协议的资产。			-
	is_fraud=falsenew		排除仿冒/欺诈数据			-
	is_honeypot=false		排除蜜罐数据				仅限FOFA高级会员使用
	is_ipv6=true			搜索ipv6的资产				搜索ipv6的资产,只接受true和false。
	is_domain=true			搜索域名的资产				搜索域名的资产,只接受true和false。
	port_size="6"			查询开放端口数量等于"6"的资产		仅限FOFA会员使用
	port_size_gt="6"		查询开放端口数量大于"6"的资产		仅限FOFA会员使用
	port_size_lt="12"		查询开放端口数量小于"12"的资产		仅限FOFA会员使用
	ip_ports="80,161"		搜索同时开放80和161端口的ip		搜索同时开放80和161端口的ip资产(以ip为单位的资产数据)
	ip_country="CN"			搜索中国的ip资产。			搜索中国的ip资产
	ip_region="Zhejiang"		搜索指定行政区的ip资产。		索指定行政区的资产
	ip_city="Hangzhou"		搜索指定城市的ip资产。			搜索指定城市的资产
	ip_after="2021-03-18"		搜索2021-03-18以后的ip资产。		搜索2021-03-18以后的ip资产
	ip_before="2019-09-09"		搜索2019-09-09以前的ip资产。		搜索2019-09-09以前的ip资产
`
	fmt.Print(syntax)
}
