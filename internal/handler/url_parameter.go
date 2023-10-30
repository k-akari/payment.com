package handler

import (
	"context"
	"errors"
	"fmt"
	"strconv"
)

type (
	ctxKey int
)

const (
	CompanyID ctxKey = iota
	ClientID
)

func getClientIDFromCtx(ctx context.Context) (int, error) {
	sclid, ok := ctx.Value(ClientID).(string)
	if !ok {
		return 0, errors.New("clientID not found in context")
	}

	clid, err := strconv.Atoi(sclid)
	if err != nil {
		return 0, fmt.Errorf("failed to convert client ID: %w", err)
	}

	return clid, nil
}

func getCompanyIDFromCtx(ctx context.Context) (int, error) {
	scoid, ok := ctx.Value(CompanyID).(string)
	if !ok {
		return 0, errors.New("companyID not found in context")
	}

	coid, err := strconv.Atoi(scoid)
	if err != nil {
		return 0, fmt.Errorf("failed to convert company ID: %w", err)
	}

	return coid, nil
}
