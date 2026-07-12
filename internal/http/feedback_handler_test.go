package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"ms-feedbacks/feedback"
	feedbackmocks "ms-feedbacks/feedback/mocks"

	"github.com/gin-gonic/gin"
)

func TestFeedbackHandler_GetFeedbacksByFigureID_InvalidID(t *testing.T) {
	t.Parallel()

	router := setupTestRouter(&FeedbackHandler{})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ms-feedback/get/abc", nil)
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status code = %d, want %d", rec.Code, http.StatusBadRequest)
	}

	if !bytes.Contains(rec.Body.Bytes(), []byte("invalid figure id")) {
		t.Fatalf("response body = %s, want invalid figure id", rec.Body.String())
	}
}

func TestFeedbackHandler_GetFeedbacksByFigureID_Success(t *testing.T) {
	t.Parallel()

	useCase := feedbackmocks.NewUseCase(t)
	expected := []feedback.Feedback{
		{
			ID:          "1",
			CustomerID:  "customer-1",
			Rating:      4,
			Description: "good",
			CreatedAt:   time.Date(2026, time.July, 12, 10, 0, 0, 0, time.UTC),
			IdFigure:    10,
			IdUser:      20,
		},
	}

	useCase.EXPECT().GetFeedbacksByFigureID(10).Return(expected, nil)

	router := setupTestRouter(NewFeedbackHandler(useCase))

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ms-feedback/get/10", nil)
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status code = %d, want %d", rec.Code, http.StatusOK)
	}

	var got []feedback.Feedback
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}

	if len(got) != 1 || got[0].ID != expected[0].ID {
		t.Fatalf("response = %#v, want %#v", got, expected)
	}
}

func TestFeedbackHandler_GetFeedbacksByFigureID_Error(t *testing.T) {
	t.Parallel()

	useCase := feedbackmocks.NewUseCase(t)
	useCase.EXPECT().GetFeedbacksByFigureID(10).Return(nil, errors.New("boom"))

	router := setupTestRouter(NewFeedbackHandler(useCase))

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ms-feedback/get/10", nil)
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("status code = %d, want %d", rec.Code, http.StatusInternalServerError)
	}

	if !bytes.Contains(rec.Body.Bytes(), []byte("boom")) {
		t.Fatalf("response body = %s, want boom", rec.Body.String())
	}
}

func TestFeedbackHandler_CreateFeedback_InvalidBody(t *testing.T) {
	t.Parallel()

	router := setupTestRouter(&FeedbackHandler{})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/ms-feedback", bytes.NewBufferString("not-json"))
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status code = %d, want %d", rec.Code, http.StatusBadRequest)
	}

	if !bytes.Contains(rec.Body.Bytes(), []byte("invalid request body")) {
		t.Fatalf("response body = %s, want invalid request body", rec.Body.String())
	}
}

func TestFeedbackHandler_CreateFeedback_Success(t *testing.T) {
	t.Parallel()

	useCase := feedbackmocks.NewUseCase(t)
	input := validFeedbackRequest()
	output := input
	output.ID = "feedback-1"
	output.CreatedAt = time.Date(2026, time.July, 12, 10, 30, 0, 0, time.UTC)
	output.UpdatedAt = time.Date(2026, time.July, 12, 10, 31, 0, 0, time.UTC)

	useCase.EXPECT().CreateFeedback(input).Return(output, nil)

	router := setupTestRouter(NewFeedbackHandler(useCase))

	payload, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/ms-feedback", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status code = %d, want %d", rec.Code, http.StatusCreated)
	}

	var got feedback.Feedback
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}

	if got.ID != output.ID {
		t.Fatalf("response = %#v, want %#v", got, output)
	}
}

func TestFeedbackHandler_CreateFeedback_Error(t *testing.T) {
	t.Parallel()

	useCase := feedbackmocks.NewUseCase(t)
	useCase.EXPECT().CreateFeedback(validFeedbackRequest()).Return(feedback.Feedback{}, errors.New("boom"))

	router := setupTestRouter(NewFeedbackHandler(useCase))

	payload, err := json.Marshal(validFeedbackRequest())
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/ms-feedback", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("status code = %d, want %d", rec.Code, http.StatusInternalServerError)
	}

	if !bytes.Contains(rec.Body.Bytes(), []byte("boom")) {
		t.Fatalf("response body = %s, want boom", rec.Body.String())
	}
}

func setupTestRouter(handler *FeedbackHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/ms-feedback/get/:idFigure", handler.GetFeedbacksByFigureID())
	router.POST("/ms-feedback", handler.CreateFeedback())

	return router
}

func validFeedbackRequest() feedback.Feedback {
	return feedback.Feedback{
		ID:          "feedback-input-1",
		CustomerID:  "customer-1",
		Rating:      5,
		Description: "great product",
		CreatedAt:   time.Date(2026, time.July, 12, 10, 0, 0, 0, time.UTC),
		IdFigure:    10,
		IdUser:      20,
	}
}
