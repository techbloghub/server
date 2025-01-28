package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/techbloghub/server/internal/testutils"
)

func TestMainIntegration(t *testing.T) {
	cfg, err := testutils.NewTestConfig(t)

	require.NoError(t, err, "config 로드 실패")

	r, client, err := createServer(cfg)
	require.NoError(t, err, "서버 생성중 에러 발생")
	require.NotNil(t, r, "gin server생성 실패")
	require.NotNil(t, client, "db client 생성 실패")

	defer client.Close()

	ts := httptest.NewServer(r)
	defer ts.Close()

	// Make a request to /ping to ensure everything is wired up
	url := fmt.Sprintf("%s/ping", ts.URL)
	resp, err := http.Get(url)
	require.NoError(t, err, "/ping 호출중 에러 발생")
	require.Equal(t, http.StatusOK, resp.StatusCode, "200 OK반환 실패")
}
