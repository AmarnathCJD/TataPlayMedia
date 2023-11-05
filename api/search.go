package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SearchReqHandler(w http.ResponseWriter, r *http.Request) {
	xTime := time.Now()
	query := r.URL.Query().Get("q")
	getAll := r.URL.Query().Get("all")
	if r.Method == "POST" {
		query = r.FormValue("q")
		getAll = r.FormValue("all")
	}

	if getAll == "true" {
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		enc.SetEscapeHTML(false)

		var channels []Channel
		for _, channel := range MainTV.Channels {
			ch := channel
			ch.Id = ch.GetID()
			ch.URL = ""
			channels = append(channels, ch)
		}

		if err := enc.Encode(channels); err != nil {
			http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
			return
		}

		return
	}

	if query == "" {
		http.Error(w, `{"error":"Query is empty"}`, http.StatusBadRequest)
		return
	}

	channels, err := MainTV.Search(query)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	if len(channels) == 0 {
		http.Error(w, `{"error":"No results found, try with a more refined query"}`, http.StatusBadRequest)
		return
	}

	var channelsQ []Channel
	for _, channel := range channels {
		ch := channel
		ch.Id = ch.GetID()
		ch.URL = ""
		channelsQ = append(channelsQ, ch)
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)

	w.Header().Set("Content-Type", "application/json")

	if err := enc.Encode(channelsQ); err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
	}

	w.Header().Set("X-Search-Time", fmt.Sprintf("%f", time.Since(xTime).Seconds()))
}

type TV struct {
	Channels []Channel `json:"channels,omitempty"`
}

type Channel struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"link,omitempty"`
	Image string `json:"image,omitempty"`
	Genre string `json:"genre,omitempty"`
	Id    string `json:"id,omitempty"`
}

var MainTV = &TV{}

func __load_channels() {
	if err := json.Unmarshal([]byte(ChannelList), &MainTV.Channels); err != nil {
		log.Fatal(err)
	}
}

func (t *TV) Search(query string) ([]Channel, error) {
	var channels []Channel
	for _, channel := range t.Channels {
		if strings.Contains(strings.ToLower(channel.Name), strings.ToLower(query)) {
			channels = append(channels, channel)
		} else if strings.Contains(strings.ToLower(channel.Genre), strings.ToLower(query)) {
			channels = append(channels, channel)
		}
	}

	for _, channel := range channels {
		channel.Id = channel.GetID()
	}
	return channels, nil
}

type DRM struct {
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
	MPD   string `json:"mpd,omitempty"`
	KeyID string `json:"key_id,omitempty"`
	Key   string `json:"key,omitempty"`
	Host  string `json:"host,omitempty"`
	Curr  string `json:"current_program,omitempty"`
}

func (t *TV) GetChannel(id string) (*Channel, error) {
	for _, channel := range t.Channels {
		if strings.EqualFold(channel.GetID(), id) {
			return &channel, nil
		}
	}
	return nil, fmt.Errorf("Channel not found")
}

func (c *Channel) GetID() string {
	return strings.Split(c.URL, "id=")[1]
}

func (c *Channel) GetDRMWithoutEnc() (*DRM, error) {
	var drm DRM
	if c.URL == "" {
		return nil, fmt.Errorf("URL is empty")
	}

	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	drm.MPD = extractValue(string(body), `file: "`, `"`)
	drm.KeyID = extractValue(string(body), `"keyId":"`, `"`)
	drm.Key = extractValue(string(body), `key":"`, `"`)
	drm.Curr = extractValue(string(body), `program-name">`, `<`)
	host, err := url.Parse(drm.MPD)
	if err != nil {
		return nil, err
	}

	drm.Host = host.Host
	drm.Name = c.Name
	drm.Image = c.Image

	return &drm, nil
}

func (c *Channel) GetDRM() (*DRM, error) {
	var drm DRM
	if c.URL == "" {
		return nil, fmt.Errorf("URL is empty")
	}

	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	prg := extractProgramName(body)
	bdy := extractScriptBody(body)

	if jsq, err := compileJS(bdy); err != nil {
		return nil, err
	} else {
		drm.MPD = extractValue(jsq, `file: "`, `"`)
		drm.KeyID = extractValue(jsq, `keyId":"`, `"`)
		drm.Key = extractValue(jsq, `key":"`, `"`)

		urlHost, _ := url.Parse(drm.MPD)
		drm.Host = urlHost.Host
		drm.Curr = prg
		drm.Name = c.Name
		drm.Image = c.Image
	}

	return &drm, nil
}

func extractProgramName(body []byte) string {
	prg := strings.Split(strings.Split(string(body), `program-name">`)[1], `<br>`)[0]
	return prg
}

func extractScriptBody(body []byte) string {
	bdy := strings.Split(strings.Split(string(body), `<div id="jwplayerDiv"></div>`)[1], `<script>`)[1]
	bdy = strings.Replace(strings.Split(bdy, `</script>`)[0], `eval`, `console.log`, -1)
	return bdy
}

func extractValue(input, prefix, suffix string) string {
	value := strings.Split(strings.Split(input, prefix)[1], suffix)[0]
	return value
}

func escapeCode(c string) string {
	return strings.NewReplacer(`\`, `\\`, `"`, `\"`, `'`, `\'`, "/", `\/`, "\n", `\n`, "\r", `\r`, "\t", `\t`).Replace(c)
}

func compileJS(code string) (string, error) {
	rq, _ := http.Get("https://onecompiler.com/api/getIdAnd")
	var dq struct {
		Id    string `json:"id"`
		Token string `json:"token"`
	}

	json.NewDecoder(rq.Body).Decode(&dq)
	var data = strings.NewReader(`{"name":"JavaScript","title":"` + dq.Id + `","version":"ES6","mode":"javascript","description":null,"extension":"js","languageType":"programming","active":true,"properties":{"language":"javascript","docs":true,"tutorials":true,"cheatsheets":true,"filesEditable":true,"filesDeletable":true,"files":[{"name":"index.js","content":"` + escapeCode(code) + `"}],"newFileOptions":[{"helpText":"New JS file","name":"script${i}.js","content":"/**\n *  In main file\n *  let script${i} = require('./script${i}');\n *  console.log(script${i}.sum(1, 2));\n */\n\nfunction sum(a, b) {\n    return a + b;\n}\n\nmodule.exports = { sum };"},{"helpText":"Add Dependencies","name":"package.json","content":"{\n  \"name\": \"main_app\",\n  \"version\": \"1.0.0\",\n  \"description\": \"\",\n  \"main\": \"HelloWorld.js\",\n  \"dependencies\": {\n    \"lodash\": \"^4.17.21\"\n  }\n}"}]},"_id":"` + dq.Id + `","user":null,"idToken":"` + dq.Token + `","visibility":"public"}`)
	req, _ := http.NewRequest("POST", "https://onecompiler.com/api/code/exec", data)

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res struct {
		Stdout string `json:"stdout"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	return strings.TrimSpace(res.Stdout), nil
}

func init() {
	__load_channels()
}
