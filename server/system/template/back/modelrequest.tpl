/*
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:自动生成
 */
package RequestModel
{{.ReImportString}}
import (
    "time"

    "chunDoc/system/model/CommonModel"
)
type {{.StructName}} struct { 
    CommonModel.PageInfo
    {{range .Fields }}{{ if .Search }}{{.Name}} {{.Type}} `json:"{{.Json}}" search:"like"`
    {{ end }}{{ end }}
     CreatedAt     []time.Time `json:"created_at" search:"btw"`
     UpdatedAt     []time.Time `json:"updated_at" search:"btw"`
} 

  