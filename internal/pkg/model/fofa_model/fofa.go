package fofa_model

import (
	"encoding/json"
	"github.com/EwanSunn/secScan/internal/pkg/slog"
	"github.com/EwanSunn/secScan/internal/pkg/util/encode"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type FofaConfig struct {
	Email string `yaml:"email"`
	Key   string `yaml:"key"`
}

type Fofa struct {
	fofaConfig                     FofaConfig
	baseURL, loginPath, searchPath string
	fieldList                      []string
	size                           int
	results                        []FofaResult
}

type FofaResult struct {
	Host, Title, Ip, Domain, Port, Country string
	Province, City, Country_name, Protocol string
	Server, Banner, Isp, As_organization   string
	Header, Cert                           string
}

type Response struct {
	//参考：https://fofa.info/api
	Error   bool       `json:"error"`
	Size    int        `json:"size"`
	Page    int        `json:"page"`
	Mode    string     `json:"mode"`
	Query   string     `json:"query"`
	Results [][]string `json:"results"`
}

func New(email, key string) *Fofa {
	var f = &Fofa{
		fofaConfig: FofaConfig{Email: email, Key: key},
		baseURL:    "https://fofa.info",
		searchPath: "/api/v1/search/all",
		loginPath:  "/api/v1/info/my",
		fieldList: []string{
			"host",
			"title",
			"banner",
			"header",
			"ip", "domain", "port", "country", "province",
			"city", "country_name",
			"server",
			"protocol",
			"cert", "isp", "as_organization",
		},
	}
	return f
}

func (f *Fofa) SetSize(i int) {
	f.size = i
}

func (f *Fofa) Search(keyword string) (int, []FofaResult) {
	url := f.baseURL + f.searchPath
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query()
	q.Add("qbase64", encode.Base64Encode(keyword))
	q.Add("email", f.fofaConfig.Email)
	q.Add("key", f.fofaConfig.Key)
	q.Add("page", "1")
	q.Add("fields", strings.Join(f.fieldList, ","))
	q.Add("size", strconv.Itoa(f.size))
	q.Add("full", "false")
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error(err)
		return 0, nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//slog.Error(err)
		return 0, nil
	}
	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		slog.Error(body, err)
		return 0, nil
	}
	r := f.makeResult(response)
	f.results = append(f.results, r...)
	return response.Size, r
}

func (f *Fofa) makeResult(response Response) []FofaResult {
	var results []FofaResult
	var result FofaResult

	for _, row := range response.Results {
		m := reflect.ValueOf(&result).Elem()
		for index, f := range f.fieldList {
			f = encode.First2Upper(f)
			m.FieldByName(f).SetString(row[index])
		}
		results = append(results, result)
	}
	return results
}

func (f *Fofa) Results() []FofaResult {
	return f.results
}

func (r FofaResult) CreateMap() map[string]string {
	t := reflect.TypeOf(r)
	v := reflect.ValueOf(r)
	m := make(map[string]string)
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		value := v.Field(k).String()
		m[key] = value
	}
	return m
}
