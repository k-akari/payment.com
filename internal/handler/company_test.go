package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/k-akari/payment.com/internal/handler/mock"
	"github.com/k-akari/payment.com/internal/testutil"
	gomock "go.uber.org/mock/gomock"
)

type companyHandlerTestHelper struct {
	sub            *companyHandler
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

		h.companyUsecase.EXPECT().CreateCompany(gomock.Any(), gomock.Any()).Return(nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/companies", bytes.NewReader(testutil.LoadFile(t, "testdata/create_company/ok_req.json.golden")))
		h.sub.Create(w, r)

		resp := w.Result()
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code should be 200, but got: %d", resp.StatusCode)
		}
	})
}
