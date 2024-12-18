package storage

import (
	"backend/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func DoRequest(method string, url string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	SetHeaders(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	// Crear un buffer para almacenar el cuerpo original
	var bodyBuffer bytes.Buffer
	tee := io.TeeReader(resp.Body, &bodyBuffer) // Clonamos el contenido para leerlo sin consumirlo

	// Leer el cuerpo para impresión
	bodyContent, readErr := io.ReadAll(tee)
	if readErr != nil {
		return resp, fmt.Errorf("error reading response body: %s", readErr)
	}

	// Imprimir detalles del response
	fmt.Println("=========================================")
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	fmt.Println("Response Body (as string):", string(bodyContent)) // Mostrar como texto

	// Intentar decodificar JSON si es válido
	var jsonBody interface{}
	if jsonErr := json.Unmarshal(bodyContent, &jsonBody); jsonErr == nil {
		prettyJSON, _ := json.MarshalIndent(jsonBody, "", "  ") // Formatear bonito
		fmt.Println("Response Body (as JSON):", string(prettyJSON))
	} else {
		fmt.Println("Response Body is not valid JSON.")
	}
	fmt.Println("Response ContentLength:", resp.ContentLength)
	fmt.Println("=========================================")

	// Restaurar el cuerpo al Response para su posterior procesamiento
	resp.Body = io.NopCloser(&bodyBuffer)

	if resp.StatusCode >= 100 && resp.StatusCode < 200 {
		return resp, fmt.Errorf("informational response: %s", string(bodyContent))
	}

	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		return resp, fmt.Errorf("redirection %s ", string(bodyContent))
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return resp, fmt.Errorf("%s", string(bodyContent))
	}

	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		return resp, fmt.Errorf("internal server error %s - statusCode %d", string(bodyContent), resp.StatusCode)
	}

	return resp, nil
}

func SetHeaders(req *http.Request) {

	req.SetBasicAuth(config.DB_USER, config.DB_PASSWORD)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

}
