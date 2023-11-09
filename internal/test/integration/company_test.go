package integration_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/k-akari/golang-rest-api-sample/internal/domain"
	"github.com/k-akari/golang-rest-api-sample/internal/handler"
	"github.com/k-akari/golang-rest-api-sample/internal/infrastructure/database"
	"github.com/k-akari/golang-rest-api-sample/internal/infrastructure/repository"
	"github.com/k-akari/golang-rest-api-sample/internal/testutil"
	"github.com/k-akari/golang-rest-api-sample/internal/usecase"
)

type companyHandlerTestHelper struct {
	sub *handler.CompanyHandler
}

func newClientHandlerTestHelper(t *testing.T) *companyHandlerTestHelper {
	sqlxDB, dbCloseFunc := testutil.InitDB()

	dba := database.NewClient(sqlxDB)
	cor := repository.NewCompanyRepository(dba)
	cou := usecase.NewCompanyUsecase(cor)
	coh := handler.NewCompanyHandler(cou)

	t.Cleanup(dbCloseFunc)

	return &companyHandlerTestHelper{
		sub: coh,
	}
}

func TestCompanyHandler_Create(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		h := newClientHandlerTestHelper(t)

		company := domain.Company{
			Name:            "company_name",
			Representative:  "company_representative",
			TelephoneNumber: "company_telephone_number",
			PostalCode:      "company_postal_code",
			Address:         "company_address",
		}
		jc, err := json.Marshal(company)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/companies", bytes.NewReader(jc))
		h.sub.Create(w, r)

		resp := w.Result()
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code should be 200, but got: %d", resp.StatusCode)
		}

		gb, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		wr := struct {
			ID domain.CompanyID `json:"id"`
		}{ID: 1}
		wb, err := json.Marshal(wr)
		if err != nil {
			t.Fatal(err)
		}

		testutil.AssertJSON(t, gb, wb)
	})
}
