package buyin

import "testing"

func TestVideoList(t *testing.T) {
	config := ClientConfig{
		Cookie:    "qc_tt_tag=0; is_staff_user=false; __security_mc_1_s_sdk_crypt_sdk=a98483d8-4be8-a9c8; bd_ticket_guard_client_web_domain=2; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWl0ZXJhdGlvbi12ZXJzaW9uIjoxLCJiZC10aWNrZXQtZ3VhcmQtcmVlLXB1YmxpYy1rZXkiOiJCQWtlMTg3Q3hjKy9nY0hJN1V3M0ZtaXQzMEhiNnJZYXBCeXQzblUwdENubGxHaVZkMndNc2pkdTMwV2w4SGMxRHNZZHA2aUJxekQ4aXZmeCtYR2cyaDA9IiwiYmQtdGlja2V0LWd1YXJkLXdlYi12ZXJzaW9uIjoyfQ%3D%3D; s_v_web_id=verify_mfm08vs4_43g3XXFm_zRMi_4ydh_A4Sd_wEcsat9r2NS5; passport_csrf_token=bf94b57109d8402a1f0b75a104bcab39; passport_csrf_token_default=bf94b57109d8402a1f0b75a104bcab39; passport_mfa_token=CjaJktqh2NRokS%2B5eH2cJhRDb3oT8YuX7AiGaVMBHTHk9OC%2FegQWwdT%2FNmDm3XAg7COCToK%2F%2BjwaSgo8AAAAAAAAAAAAAE%2BYppJ7O7diowKrunZGM7lsd40DHk1i72SbN6nd86Mmc74U%2B0%2Bpx4utPJF1PMAvjZ4vEJfu%2Fg0Y9rHRbCACIgEDlCKVHw%3D%3D; odin_tt=f5ade0b0f57e4ac3fc4a3b160738c4beb2500b3c028b5fce8e04de09aa53f46169e4a718c284bd34b5a90005e0582443ef69cecea93c64c0d7435501eb2f389a; passport_auth_status=3d54bd69561b68ec36e84ddf07522566%2Ca8942dcb4a7d91335259f4d181c12cf8; passport_auth_status_ss=3d54bd69561b68ec36e84ddf07522566%2Ca8942dcb4a7d91335259f4d181c12cf8; uid_tt=7c338817748c7f615e0626a483e51c37; uid_tt_ss=7c338817748c7f615e0626a483e51c37; sid_tt=ae3936a51f19231f0670111188eb321d; sessionid=ae3936a51f19231f0670111188eb321d; sessionid_ss=ae3936a51f19231f0670111188eb321d; ucas_c0_compass=CkAKBTEuMC4wEJuIgJ7fxJ38aBjmJiCU8YC46ozoBiiwITDUs6D3gKwMQKjs4ccGSKignsoGUKq8oabq4YPzaFhvEhRAnmx1Py9WOOX3rRdjtLyIyxO5Sg; ucas_c0_ss_compass=CkAKBTEuMC4wEJuIgJ7fxJ38aBjmJiCU8YC46ozoBiiwITDUs6D3gKwMQKjs4ccGSKignsoGUKq8oabq4YPzaFhvEhRAnmx1Py9WOOX3rRdjtLyIyxO5Sg; LUOPAN_DT=session_7563923318974447911; COMPASS_LUOPAN_DT=session_7563923318974447911; gd_random=eyJtYXRjaCI6dHJ1ZSwicGVyY2VudCI6MC40NDY1NDYxNDEyMzQwNzY5fQ==.MxBbzM3RBgk9hqQFEoT1AYdfXFxnNGOv2Ffme7afSew=",
		EWID:      "ae3390605f3dfe70921708a0c3390b7b1",
		VerifyFp:  "verify_mfm08vs4_43g3XXFm_zRMi_4ydh_A4Sd_wEcsat9rd2NS5",
		Fp:        "verify_mfm08vs4_43g3XXFm_zRMi_4ydh_A4Sd_wEcsat9r2fNS5",
		MsToken:   "IleyX6CDVvfM_yczR69Jm-MmtFSe6wPhN8X4WaJ6-gm2Q63N8j5kzzefyaq6B9_AvcYEWOpiYl5gr0vp_Qi6juzuILtDK9f1nIzt4JgkEXWntA06rx-2Hgt-nW9iWrrGo09aKieHkdaDJvGDGLTtsKlbl3c6ao2_3XLaduXDYQNY_Nf8HnfWHXA==",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36 Edg/141.0.0.0",
	}

	// 根据 video_list.txt 中的 curl 请求参数构建 VideoListParams
	params := VideoListParams{
		PageNo:        "1",
		PageSize:      "10",
		BeginDate:     "2025/10/16 00:00:00",
		EndDate:       "2025/10/22 00:00:00",
		DateType:      "21",
		ActivityID:    "",
		AccountType:   "2",
		AuthorID:      "",
		RangeType:     "0",
		CartType:      "1",
		AdType:        "0",
		SearchInfo:    "3532241584139669146",
		IndexSelected: "watch_cnt,pay_amt,refund_amt,live_pay_amt,search_pay_amt,lead_shop_pay_amt",
		SortField:     "",
		LID:           "036579525900",
	}

	t.Logf("config:%v", config)
	t.Logf("params:%v", params)

	res, status, err := NewJXClient(config).VideoList(params)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	t.Logf("Status: %s", status)
	t.Logf("res: %s", res)
}
