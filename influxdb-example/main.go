package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "time"

    client "github.com/influxdata/influxdb1-client/v2"
)

// ReadConfig reads the configuration from a file and returns the username and password
func ReadConfig(filename string) (string, string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return "", "", err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var username, password string

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            return "", "", fmt.Errorf("invalid config line: %s", line)
        }
        key, value := parts[0], parts[1]
        switch key {
        case "username":
            username = value
        case "password":
            password = value
        default:
            return "", "", fmt.Errorf("unknown config key: %s", key)
        }
    }

    if err := scanner.Err(); err != nil {
        return "", "", err
    }

    return username, password, nil
}

func createClient(username, password string) (client.Client, error) {
    c, err := client.NewHTTPClient(client.HTTPConfig{
        Addr:     "http://192.168.102.133:8086", // Replace with your InfluxDB address
        Username: username,
        Password: password,
    })
    return c, err
}

func writeData(c client.Client) {
    // Create a new point batch
    bp, err := client.NewBatchPoints(client.BatchPointsConfig{
        Database:  "cbj", // Replace with your InfluxDB database
        Precision: "s",
    })
    if err != nil {
        log.Fatal(err)
    }

    // Create a point and add to batch
    tags := map[string]string{"cpu": "cpu-total"}
    fields := map[string]interface{}{
        "idle":   22.2,
        "system": 55.5,
        "user":   44.4,
    }

    pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
    if err != nil {
        log.Fatal(err)
    }
    bp.AddPoint(pt)

    // Write the batch
    if err := c.Write(bp); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Data written successfully")
}

func main() {
    // Read the configuration file
    username, password, err := ReadConfig("config.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Create a new HTTPClient
    c, err := createClient(username, password)
    if err != nil {
        log.Fatal(err)
    }
    defer c.Close()

    // Write data
    writeData(c)
}
