package requests

type SetCurrentSceneCollection struct {
	RequestBase
	SceneCollectionName string `json:"sc-name"`
}
