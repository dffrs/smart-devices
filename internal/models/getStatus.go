package internal

type GetStatus struct {
	Ble     Ble     `json:"ble"`
	Bthome  Bthome  `json:"bthome"`
	Cloud   Cloud   `json:"cloud"`
	Knx     Knx     `json:"knx"`
	Matter  Matter  `json:"matter"`
	Mqtt    Mqtt    `json:"mqtt"`
	PlugsUI PlugsUI `json:"plugs_ui"`
	Switch0 Switch0 `json:"switch:0"`
	Sys     Sys     `json:"sys"`
	Wifi    Wifi    `json:"wifi"`
	Ws      Ws      `json:"ws"`
}

type (
	Ble    struct{}
	Bthome struct{}
	Cloud  struct {
		Connected bool `json:"connected"`
	}
)

type (
	Knx    struct{}
	Matter struct {
		Commissionable bool `json:"commissionable"`
		NumFabrics     int  `json:"num_fabrics"`
	}
)

type Mqtt struct {
	Connected bool `json:"connected"`
}
type (
	PlugsUI struct{}
	Aenergy struct {
		ByMinute []float64 `json:"by_minute"`
		MinuteTS int       `json:"minute_ts"`
		Total    float64   `json:"total"`
	}
)

type RetAenergy struct {
	ByMinute []int `json:"by_minute"`
	MinuteTS int   `json:"minute_ts"`
	Total    int   `json:"total"`
}
type Temperature struct {
	TC float64 `json:"tC"`
	TF float64 `json:"tF"`
}
type Switch0 struct {
	Aenergy     Aenergy     `json:"aenergy"`
	Apower      float64     `json:"apower"`
	Current     float64     `json:"current"`
	Freq        float64     `json:"freq"`
	ID          int         `json:"id"`
	Output      bool        `json:"output"`
	RetAenergy  RetAenergy  `json:"ret_aenergy"`
	Source      string      `json:"source"`
	Temperature Temperature `json:"temperature"`
	Voltage     float64     `json:"voltage"`
}

type (
	AvailableUpdates struct{}
	Sys              struct {
		AvailableUpdates AvailableUpdates `json:"available_updates"`
		BtrelayRev       int              `json:"btrelay_rev"`
		CfgRev           int              `json:"cfg_rev"`
		FsFree           int              `json:"fs_free"`
		FsSize           int              `json:"fs_size"`
		KvsRev           int              `json:"kvs_rev"`
		LastSyncTS       int              `json:"last_sync_ts"`
		Mac              string           `json:"mac"`
		RAMFree          int              `json:"ram_free"`
		RAMMinFree       int              `json:"ram_min_free"`
		RAMSize          int              `json:"ram_size"`
		ResetReason      int              `json:"reset_reason"`
		RestartRequired  bool             `json:"restart_required"`
		ScheduleRev      int              `json:"schedule_rev"`
		Time             string           `json:"time"`
		Unixtime         int              `json:"unixtime"`
		Uptime           int              `json:"uptime"`
		UtcOffset        int              `json:"utc_offset"`
		WebhookRev       int              `json:"webhook_rev"`
	}
)

type Wifi struct {
	Bssid  string   `json:"bssid"`
	Rssi   int      `json:"rssi"`
	Ssid   string   `json:"ssid"`
	StaIP  string   `json:"sta_ip"`
	StaIP6 []string `json:"sta_ip6"`
	Status string   `json:"status"`
}

type Ws struct {
	Connected bool `json:"connected"`
}
