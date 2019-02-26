package responses

type GetCurrentSceneCollection struct {
	*ResponseBase
	SceneCollectionName string `json:"sc-name"`
}

type ListSceneCollections struct {
	*ResponseBase
	SceneCollections []string `json:"scene-collections"`
}