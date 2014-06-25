package dam

import (
    "crypto/tls"
    "fmt"
    "net/http"
    "net/url"
)

type DamConnection struct {
    // This is the base url with trailing slash
    // Example: https://cd-dam/
    Url string

    // This is the querystring key.. i don't feel like making an HMAC signature right now
    // Example: pica9testing_APPLICATION
    Key string

    // net/http client type.
    client *http.Client
}

func NewConnection(url, key string) *DamConnection {
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    c := &http.Client{Transport: tr}

    return &DamConnection{client: c, Url: url, Key: key}
}

func BuildUrl(base, uri string, query map[string]string) (string, error) {
    slash := ""

    if !SlashCheck(base) {
        slash = "/"
    }

    temp_url := fmt.Sprintf("%s%s%s", base, slash, uri)

    parsed_url, err := url.Parse(temp_url)

    if err != nil {
        return "", err
    }

    q := parsed_url.Query()

    for k, v := range query {
        q.Set(k, v)
    }

    parsed_url.RawQuery = q.Encode()

    return parsed_url.String(), nil
}

func SlashCheck(base string) bool {
    substring := base[len(base)-1:]
    if substring == "/" {
        return true
    }

    return false
}

func (d *DamConnection) Ping() bool {
    var headers = make(map[string]string)
    headers["api_key"] = d.Key

    target, _ := BuildUrl(d.Url, "", headers)
    res, err := d.client.Get(target)

    if err != nil || res.StatusCode >= 400 {
        return false
    }

    return true
}
