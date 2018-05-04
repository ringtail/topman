package seaway

type ScreenshotCorsair struct {
	Corsair
	Name string `json:"name"`
}

func (sc *ScreenshotCorsair) GetName() string {
	return sc.Name
}

func (sc *ScreenshotCorsair) Wigwag() (spotted bool, err error) {
	return
}

func (sc *ScreenshotCorsair) Msg() (corsairInfo *CorsairInfo, err error) {
	return
}
