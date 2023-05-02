package main

import (
	"context"
	"fmt"
	"os/user"
	"path/filepath"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// NewComputerEssentials 新电脑必备
func (a *App) new_computer_essentials() {
	// 获取电脑桌面对应的目录
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Failed to get current user:", err)
		return
	}

	desktopPath := filepath.Join(currentUser.HomeDir, "Desktop")
	fmt.Println("Desktop path:", desktopPath)
}
