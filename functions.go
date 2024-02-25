package main

import (
    _ "github.com/go-sql-driver/mysql"
)

func createUser(username, email string) error {
    db, err := connectDB()
    if err != nil {
        return err
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
    if err != nil {
        return err
    }

    return nil
}

func readUser(userID int) (string, string, error) {
    db, err := connectDB()
    if err != nil {
        return "", "", err
    }
    defer db.Close()

    var username, email string
    err = db.QueryRow("SELECT username, email FROM users WHERE id=?", userID).Scan(&username, &email)
    if err != nil {
        return "", "", err
    }

    return username, email, nil
}

func updateUser(userID int, newEmail string) error {
    db, err := connectDB()
    if err != nil {
        return err
    }
    defer db.Close()

    _, err = db.Exec("UPDATE users SET email=? WHERE id=?", newEmail, userID)
    if err != nil {
        return err
    }

    return nil
}

func deleteUser(userID int) error {
    db, err := connectDB()
    if err != nil {
        return err
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM users WHERE id=?", userID)
    if err != nil {
        return err
    }

    return nil
}
