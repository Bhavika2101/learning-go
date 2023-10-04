package main

import (
    "errors"
    "strings"
    "testing"
    "time"
    "math/rand"
)

func TestGeneratePassword_24386dfbfc(t *testing.T) {
    tests := []struct {
        name     string
        length   int
        hasError bool
    }{
        {"Password Length", 10, false},
        {"Negative Length", -5, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            password, err := generateRandomPassword(tt.length)
            t.Log("Running: ", tt.name, " with length: ", tt.length)

            if err != nil && !tt.hasError {
                t.Error("Expected no error but got ", err.Error())
            }

            if err == nil && tt.hasError {
                t.Error("Expected error but got none")
            }

            if err == nil && len(password) != tt.length {
                t.Error("Expected password of length ", tt.length, " but got ", len(password))
            }

            if err == nil && strings.IndexAny(password, "0123456789") == -1 {
                t.Error("Password does not contain any digit")
            }

            if err == nil && strings.IndexAny(password, "~=+%^*/()[]{}/!@#$?|") == -1 {
                t.Error("Password does not contain any special character")
            }

            if err == nil && strings.IndexAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz") == -1 {
                t.Error("Password does not contain any alphabets")
            }
        })
    }
}

func generateRandomPassword(length int) (string, error) {
    if length < 0 {
        return "", errors.New("length cannot be negative")
    }

    rand.Seed(time.Now().UnixNano())
    digits := "0123456789"
    specials := "~=+%^*/()[]{}/!@#$?|"
    all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
        "abcdefghijklmnopqrstuvwxyz" +
        digits + specials

    buf := make([]byte, length)
    if length > 0 {
        buf[0] = digits[rand.Intn(len(digits))]
    }
    if length > 1 {
        buf[1] = specials[rand.Intn(len(specials))]
    }

    for i := 2; i < length; i++ {
        buf[i] = all[rand.Intn(len(all))]
    }
    rand.Shuffle(len(buf), func(i, j int) {
        buf[i], buf[j] = buf[j], buf[i]
    })

    str := string(buf)
    return str, nil
}
