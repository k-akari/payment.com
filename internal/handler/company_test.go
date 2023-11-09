package handler

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/k-akari/golang-rest-api-sample/internal/domain"
	"github.com/k-akari/golang-rest-api-sample/internal/handler/mock"
	"github.com/k-akari/golang-rest-api-sample/internal/testutil"
	gomock "go.uber.org/mock/gomock"
)

type companyHandlerTestHelper struct {
	sub            *CompanyHandler
	companyUsecase *mock.MockcompanyUsecase
}

func createCompanyHandlerTestHelper(t *testing.T) *companyHandlerTestHelper {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	companyUsecase := mock.NewMockcompanyUsecase(ctrl)

	return &companyHandlerTestHelper{
		sub:            NewCompanyHandler(companyUsecase),
		companyUsecase: companyUsecase,
	}
}

func TestCompanyHandler_CreateCompany(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		h := createCompanyHandlerTestHelper(t)

		h.companyUsecase.EXPECT().CreateCompany(gomock.Any(), gomock.Any()).Return(domain.CompanyID(10), nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/companies", bytes.NewReader(testutil.LoadFile(t, "testdata/create_company/ok_req.json.golden")))
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
		testutil.AssertJSON(t, gb, testutil.LoadFile(t, "testdata/create_company/ok_resp.json.golden"))
	})
}
