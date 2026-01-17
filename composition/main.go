package main

import (
	"slices"
	"fmt"
	"time"
)

// --- CAPABILITY 1: Logging ---
type Logger struct {
	Prefix string
}

func (l Logger) Log(message string) {
	fmt.Printf(
		"[%s] %s: %s\n",
		time.Now().Format("19:04:05"),
		l.Prefix,
		message,
	)
}

// --- CAPABILITY 2: Permissions ---
type PermissionHandler struct {
    Roles []string 
}

func (ph PermissionHandler) HasRole(targetRole string) bool{
    return slices.Contains(ph.Roles, targetRole)
}

// --- THE COMPOSITE OBJECT: User --- 
type User struct {
    Username string     
    // Embedding both capabilities 
    Logger 
    PermissionHandler
}

func main() {
    // Initialize the user with its capabilities 
    u := User{
        Username: "jane_doe",
        Logger: Logger{Prefix: "USER_LOG"},
        PermissionHandler: PermissionHandler{
            Roles: []string{"admin", "editor"},
        },
    }

    // 1. Using the 'Promoted' Log method 
    u.Log("Logged in successfully")

    // 2. Using the 'Promoted' Permission method 
    if u.HasRole("admin") {
        u.Log("Accessing sensitive admin panel...")
    } else {
        fmt.Println("Access Denied")
    }

    // 3. Accessing the User's own fields
    fmt.Printf("User: %s is now active.\n", u.Username)
}
