package api

type StatusJobResponse struct {
	ID            int     `json:"id"`
	Progress      float64 `json:"progress"`
	TimeRemaining int     `json:"time_remaining"`
	TimePrinting  int     `json:"time_printing"`
}

type StatusPrinterResponse struct {
	State string `json:"state"`
}

type StatusResponse struct {
	Job     StatusJobResponse     `json:"job"`
	Printer StatusPrinterResponse `json:"printer"`
}

type JobFileResponse struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

type JobResponse struct {
	ID            int             `json:"id"`
	State         string          `json:"state"`
	Progress      float64         `json:"progress"`
	TimeRemaining int             `json:"time_remaining"`
	TimePrinting  int             `json:"time_printing"`
	File          JobFileResponse `json:"file"`
}

type StorageInfoResponse struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Path      string `json:"path"`
	Available bool   `json:"available"`
	ReadOnly  bool   `json:"read_only"`
}

type StorageResponse struct {
	StorageList []StorageInfoResponse `json:"storage_list"`
}

type FileItemResponse struct {
	Name              string `json:"name"`
	ReadOnly          bool   `json:"ro"`
	Type              string `json:"type"`
	DisplayName       string `json:"display_name"`
	ModifiedTimestamp int64  `json:"m_timestamp"`
}

type FilesResponse struct {
	Type     string             `json:"type"`
	ReadOnly bool               `json:"ro"`
	Name     string             `json:"name"`
	Children []FileItemResponse `json:"children"`
}
