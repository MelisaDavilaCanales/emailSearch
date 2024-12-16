package storage

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DoRequest(method string, url string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		fmt.Println("-------------------------------------------------------------- ")
		fmt.Println("------------------------------------ENTRO AQUI1-------------------------- ")
		fmt.Println("-------------------------------------------------------------- ")

		return nil, err
	}

	SetHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("-------------------------------------------------------------- ")
		fmt.Println("------------------------------------ENTRO AQUI2-------------------------- ")
		fmt.Println("-------------------------------------------------------------- ")
		return resp, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("-------------------------------------------------------------- ")
		fmt.Println("------------------------------------ENTRO AQUI3-------------------------- ")
		fmt.Println("-------------------------------------------------------------- ")
		return resp, err
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))

	return resp, nil
}

func SetHeaders(req *http.Request) {

	req.SetBasicAuth(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

}
