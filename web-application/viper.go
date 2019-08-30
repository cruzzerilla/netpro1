package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// Set loc config file
	viper.SetConfigFile("config.json")

	// Show error jika file config tidak ada
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error config file!, &s", err)
	}

	// Set route ketika web diakses
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %q", html.EscapeString(r.URL.Path))
	})

	// Run server sesuai port pada file config
	http.ListenAndServe(":"+viper.GetString("server.port"), nil)
}
