package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	output       string
	customDomain string
	showHelp     bool
	showProxy    bool
	configFile   = "wgee-config.txt" // 配置文件路径
)

func main() {
	flag.StringVar(&output, "O", "", "Specify output file name")
	flag.StringVar(&customDomain, "c", "", "Custom domain with trailing /")
	flag.BoolVar(&showHelp, "help", false, "Show usage information")
	flag.BoolVar(&showProxy, "proxy", false, "Show current configured domain")
	flag.Parse()

	// 如果使用 -proxy 参数，则显示当前配置的域名并退出
	if showProxy {
		currentDomain := getCurrentDomain()
		fmt.Println("Current configured domain:", currentDomain)
		os.Exit(0)
	}

	// 如果使用 -help 参数，则显示用法信息并退出
	if showHelp {
		printUsage()
		os.Exit(0)
	}

	// 如果提供了 -c 参数，则更新 customDomain 的值并持久化保存
	if customDomain != "" {
		// 检查并修正 customDomain 格式
		if !strings.HasSuffix(customDomain, "/") {
			customDomain += "/"
		}
		saveCurrentDomain(customDomain)

		// 输出配置成功信息
		fmt.Println("Configuration updated successfully.")
		fmt.Println("New configured domain:", customDomain)
		os.Exit(0)
	}

	// 如果没有提供 -c 参数，则从配置文件中读取持久化保存的 customDomain
	if customDomain == "" {
		customDomain = getCurrentDomain()
	}

	// 获取 URL 参数
	var url string
	if flag.NArg() > 0 {
		url = flag.Arg(0)
	} else {
		// 如果没有提供 URL 参数，则输出错误信息并显示用法信息
		fmt.Println("Error: Missing URL argument")
		printUsage()
		os.Exit(1)
	}

	// 构建完整的下载链接
	fullURL := strings.TrimRight(customDomain, "/") + "/" + strings.TrimLeft(url, "/")

	// 执行下载命令
	var cmd *exec.Cmd
	if output != "" {
		cmd = exec.Command("wget", "-O", output, fullURL)
	} else {
		cmd = exec.Command("wget", fullURL)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("wget failed, trying with curl...")
		cmd = exec.Command("curl", "-o", output, fullURL)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error running curl: %v\n", err)
			os.Exit(1)
		}
	}
}

func printUsage() {
	fmt.Println("Usage: ./wgee [-O output_file] [-c custom_domain] [-proxy] <URL>")
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func saveCurrentDomain(domain string) {
	// 将用户设置的 customDomain 写入配置文件
	err := os.WriteFile(configFile, []byte(domain), 0644)
	if err != nil {
		fmt.Printf("Error saving custom domain to file: %v\n", err)
		os.Exit(1)
	}
}

func getCurrentDomain() string {
	// 从配置文件中读取持久化保存的 customDomain
	data, err := os.ReadFile(configFile)
	if err != nil {
		// 如果文件不存在或出错，默认返回初始值
		return "https://newtrojan.lizhiyin.us.kg/"
	}
	return strings.TrimSpace(string(data))
}

