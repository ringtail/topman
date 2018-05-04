package seaway

type Config struct {
	PingCorsairs       []*PingCorsair       `json:"ping"`
	TcpCorsairs        []*TcpCorsair        `json:"tcp"`
	ScreenshotCorsairs []*ScreenshotCorsair `json:"screenshotCorsair"`
}
