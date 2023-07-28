package tool

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	URL_LISTGOV = "https://dados.gov.br/api/publico/conjuntos-dados/cadastro-nacional-da-pessoa-juridica---cnpj"
)

type gov struct{}

func NewGov() *gov {
	return &gov{}
}

type RecursoForm struct {
	ID          string `json:"id"`
	Description string `json:"descricao"`
	Format      string `json:"formato"`
	Link        string `json:"link"`
	Title       string `json:"titulo"`
}

type ResourcesDetails struct {
	ID          string      `json:"id"`
	Created     string      `json:"created"`
	Description string      `json:"description"`
	Format      string      `json:"format"`
	Name        string      `json:"name"`
	PackageId   string      `json:"package_id"`
	Recurso     RecursoForm `json:"recursoForm"`
}

type Resources struct {
	Details []ResourcesDetails `json:"resources"`
}

func (g *gov) ListResources() (*Resources, error) {
	response, err := http.Get(URL_LISTGOV)
	if err != nil {
		return &Resources{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return &Resources{}, err
	}

	if response.StatusCode != 200 {
		return &Resources{}, fmt.Errorf(string(body))
	}

	var responseData Resources
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return &Resources{}, err
	}

	return &responseData, nil
}
