package resource

type Namespace struct {
	Name      string `json:"name" xml:"name" form:"name" query:"name"`
	Namespace string `json:"namespace" xml:"namespace" form:"namespace" query:"namespace"`
	Status string `json:"status" xml:"status" form:"status" query:"status"`
	Age    int64  `json:"age" xml:"age" form:"age" query:"age"`
}