package services

import (
    "bytes"
    "encoding/json"
    "io"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "os"
    "path/filepath"
)

type ParsedData struct {
    Skills     string `json:"skills"`
    Education  string `json:"education"`
    Experience string `json:"experience"`
    Name       string `json:"name"`
    Email      string `json:"email"`
    Phone      string `json:"phone"`
}

func ParseResume(filePath string) (ParsedData, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return ParsedData{}, err
    }
    defer file.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("resume", filepath.Base(file.Name()))
    if err != nil {
        return ParsedData{}, err
    }

    _, err = io.Copy(part, file)
    if err != nil {
        return ParsedData{}, err
    }
    writer.Close()

    req, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", body)
    if err != nil {
        return ParsedData{}, err
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("apikey", "YOUR_API_KEY_HERE")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return ParsedData{}, err
    }
    defer resp.Body.Close()

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return ParsedData{}, err
    }

    var parsedData ParsedData
    if err := json.Unmarshal(respBody, &parsedData); err != nil {
        return ParsedData{}, err
    }

    return parsedData, nil
}
