package log

type Info struct {
	HumanContext  string                 `json:"humanContext"`
	MessageName   InfoMessage            `json:"messageName"`
	Payload       map[string]interface{} `json:"payloadObject"`
	TransactionId string                 `json:"transactionId"`
}

type InfoMessage string

const (
	BootSequenceReportMsg     InfoMessage = "Client.BootSequenceReport"
	ConnectedMsg              InfoMessage = "Client.Connected"
	InventoryReportMsg        InfoMessage = "Client.InventoryReport"
	PerformanceReportMsg      InfoMessage = "Client.PerformanceReport"
	PregameSequenceReportMsg  InfoMessage = "Client.PregameSequenceReport"
	ProgressionTrackViewedMsg InfoMessage = "Client.ProgressionTrackViewed"
	PurchaseFunnelMsg         InfoMessage = "Client.PurchaseFunnel"
	SceneChangeMsg            InfoMessage = "Client.SceneChange"
	SetAvatarSelectoinMsg     InfoMessage = "Client.SetAvatarSelection"
	SystemMessageViewMsg      InfoMessage = "Client.SystemMessageView"
	UserDeviceSpecsMsg        InfoMessage = "Client.UserDeviceSpecs"
	EventNavigationMsg        InfoMessage = "Client.Home.EventNavigation"

	GameStartMsg        InfoMessage = "DuelScene.GameStart"
	GameStopMsg         InfoMessage = "DuelScene.GameStop"
	EndOfMatchReportMsg InfoMessage = "DuelScene.EndOfMatchReport"
	EmotesUsedReportMsg InfoMessage = "DuelScene.EmotesUsedReport"
)
