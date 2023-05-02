package main

import (
	"fmt"
	"os/user"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func main() {
	// 获取电脑桌面对应的目录
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Failed to get current user:", err)
		return
	}

	desktopPath := filepath.Join(currentUser.HomeDir, "Desktop")
	fmt.Println("Desktop path:", desktopPath)
}

func changeDesktop() {
	const desktopPath = "D:\\NewDesktop" // 新的桌面路径

	// 打开 "HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Explorer\User Shell Folders" 键
	k, err := registry.OpenKey(registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Explorer\User Shell Folders`,
		registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("Failed to open registry key:", err)
		return
	}
	defer k.Close()

	// 设置 "Desktop" 值为新的桌面路径
	err = k.SetStringValue("Desktop", desktopPath)
	if err != nil {
		fmt.Println("Failed to set registry value:", err)
		return
	}

	fmt.Println("Desktop path has been changed to", desktopPath)
}
