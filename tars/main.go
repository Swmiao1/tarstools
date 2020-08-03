package tars

const (
	//上传发布包
	uploadPatchPackage = "/api/upload_patch_package"
	uploadAndPublish   = "/api/upload_and_publish"
	serverList         = "/api/server_list"
)

type Tars struct {
	Url   string `json:"tars_url"`
	Token string `json:"token"`
}
