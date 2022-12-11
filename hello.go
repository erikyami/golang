package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		// if option == 0 {
		// 	fmt.Println("Exiting...")
		// } else if option == 1 {
		// 	fmt.Println("Monitoring...")
		// } else if option == 2 {
		// 	fmt.Println("Logs...")
		// } else {
		// 	fmt.Println("Command not found")
		// }

		option := leComando()

		switch option {
		case 1:
			initMonitoring()
		case 2:
			fmt.Println("Logs...")
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Command not found")
			os.Exit(-1)
		}
	}
}

func exibeMenu() {
	fmt.Println("1 - Init Monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}
func exibeIntroducao() {
	var name string = "John"
	var version float32 = 1.1

	fmt.Println("Hello, mr.", name)
	fmt.Println("This program is in version", version)
}

func leComando() int {
	var option int
	fmt.Scan(&option)

	fmt.Println("The option choosed was", option)
	return option
}

func initMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://www.google.com.br/"
	resp, _ := http.Get(site)
	//fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site ", site, "está com problemas. Status Code: ", resp.StatusCode)
	}
}
