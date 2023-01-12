package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUsername(t *testing.T) {
	tests := []struct {
		name       string
		redisValue string
	}{
		{
			name:       "success",
			redisValue: `"get v2.local.kthrSvwfC0rXmnRbekhKapOxD8SUzfv6gbhER0c-3cicQbUPFVh_U2ZGE3cbRHAAP1APflg99wrqqyck86SjpHwP3ggW7aHvCvtA1f4kFocueAPB1cjiLxWZrjgaqqIz-lJjbfnePov7eVPFnfFOLIg_hvKJ8F33BgxExxlpTODDuBo_tQbQfES1b-3dBJfatySslP0O6yjCmm8db0nUHjqCI-QUiQH1H1RSMlN-SLY8mnW1Arv0fBNA7_IWqYNsIAYLbsD2LRwi3Y1Hp6FwqU2w99I.bnVsbA: {"id":"ec98b8f7-7cee-46ac-92d1-9e0482ba41c3","token":"srrrs","subject":"API","issued_at":"2023-01-13T06:16:49.053369+09:00","expired_at":"2023-01-13T06:46:49.053369+09:00"}"`,
		},
		{
			name:       "failed",
			redisValue: `"get v2.local.kthrSvwfC0rXmnRbekhKapOxD8SUzfv6gbhER0c-3cicQbUPFVh_U2ZGE3cbRHAAP1APflg99wrqqyck86SjpHwP3ggW7aHvCvtA1f4kFocueAPB1cjiLxWZrjgaqqIz-lJjbfnePov7eVPFnfFOLIg_hvKJ8F33BgxExxlpTODDuBo_tQbQfES1b-3dBJfatySslP0O6yjCmm8db0nUHjqCI-QUiQH1H1RSMlN-SLY8mnW1Arv0fBNA7_IWqYNsIAYLbsD2LRwi3Y1Hp6FwqU2w99I.bnVsbA: {"id":"ec98b8f7-7cee-46ac-92d1-9e0482ba41c3","token":"fail","subject":"API","issued_at":"2023-01-13T06:16:49.053369+09:00","expired_at":"2023-01-13T06:46:49.053369+09:00"}"`,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			username := GetUsernameTest(tt.redisValue)
			if i == 0 {
				require.NotEqual(t, "", username)
				require.Equal(t, "srrrs", username)
			} else if i == 1 {
				require.NotEqual(t, "", username)
				require.Equal(t, "fail", username)
			}
		})
	}
}
