package env

import "os"

func GetEnv(key, fallback string) string {
        value := os.Getenv(key)
        if len(value) == 0 {
                return fallback
        }
        return value
}

func GetHostname() string {
    hostname, _ := os.Hostname()
    return hostname
} 
