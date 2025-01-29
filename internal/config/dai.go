package config

type DaiConfig struct {
	GoogleApplicationCredential string
	FireStoreProjectId          string
	CollectionNameDai           string
	LiveServiceBaseUrl          string
	LiveServiceApiKey           string
	FirestoreDelayDaiDelete     int
}

var DaiConfigApp DaiConfig

func initDaiConfig() {
	DaiConfigApp.GoogleApplicationCredential = mustGetString("GOOGLE_APPLICATION_CREDENTIALS")
	DaiConfigApp.FireStoreProjectId = mustGetString("FIRESTORE_PROJECT_ID")
	DaiConfigApp.CollectionNameDai = mustGetString("COLLECTION_NAME_DAI")
	DaiConfigApp.LiveServiceBaseUrl = mustGetString("LIVE_SERVICE_BASE_URL")
	DaiConfigApp.LiveServiceApiKey = mustGetString("LIVE_SERVICE_API_KEY")
	DaiConfigApp.FirestoreDelayDaiDelete = mustGetInt("FIREBASE_DELAY_DAI_DELETE")
}
