package grpc_logger

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func PrintStart(host, port string) {
	fmt.Println("" +
		"   _____   _____    _____    _____   " +
		"\n  / ____| |  __ \\  |  __ \\  |  __ \\  " +
		"\n | |  __  | |__) | | |__) | | |__) | " +
		"\n | | |_ | |  _  /  |  ___/  |  _  /  " +
		"\n | |__| | | | \\ \\  | |      | | \\ \\  " +
		"\n  \\_____| |_|  \\_\\ |_|      |_|  \\_\\ ")

	fmt.Println("--------------------------------------------------")
	fmt.Printf("INFO Sever started on:    \t http://%s:%s\n", host, port)
}

func PrintError(method string, code codes.Code, err error) error {
	fmt.Printf("| %s | ---------- %s: %v\n", code.String(), method, err)
	return status.Errorf(code, method+" :", err.Error())
}

func Print(method string, code codes.Code) error {
	fmt.Printf("| %s | ---------- %s\n", code.String(), method)
	return nil
}
