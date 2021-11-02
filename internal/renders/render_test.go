package renders

import (
	"net/http"
	"testing"

	"github.com/wycemiro/booking-site/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	result := AddDefaultData(&td, r)
	if result == nil {
		t.Error("fail")
	}

}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/XIG", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"

	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplates(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser.")
	}

	err = RenderTemplates(&ww, r, "non-existant.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("rendering template that doesnt exsist")
	}
}

func TestNewTemplates(t *testing.T) {
	CreateTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
