package resource

type PVCParams struct {
	Name     string       `json:"name" xml:"name" form:"name" query:"name"`
	Capacity ResourceList `json:"capacity" xml:"capacity" form:"capacity" query:"capacity"`
}