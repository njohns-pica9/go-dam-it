package dam

import (
    "bytes"
    "net/http"
    "net/url"
    "strconv"
)

func (d *DamConnection) NewContainer(name string) bool {
    var query = make(map[string]string)
    query["api_key"] = d.Key

    target, err := BuildUrl(d.Url, "container-tags", query)

    if err != nil {
        panic(err)
    }

    var data = url.Values{}
    data.Add("name", name)

    stream := bytes.NewBufferString(data.Encode())
    req, err := http.NewRequest("POST", target, stream)

    if err != nil {
        panic(err)
    }

    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    res, err := d.client.Do(req)

    if err != nil || res.StatusCode >= 400 {
        return false
    }

    return true
}
