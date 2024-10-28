package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func registerWithServer(deviceID string, serverURL string) error {
    resp, err := http.Get(fmt.Sprintf("%s/register?id=%s", serverURL, deviceID))
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    return nil
}

func getDeviceList(serverURL string) error {
    resp, err := http.Get(fmt.Sprintf("%s/devices", serverURL))
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Devices connected:")
    fmt.Println(string(body))
    return nil
}

func main() {
    serverURL := "http://localhost:8080" // URL of the signaling server
    deviceID := "Device1"                  // Unique ID for this device

    registerWithServer(deviceID, serverURL)
    getDeviceList(serverURL)
}
