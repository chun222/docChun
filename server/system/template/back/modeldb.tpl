/*
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:自动生成{{.StructName}}
 */
package DbModel 
 {{.DbImportString}}
type {{.StructName}} struct { 
    {{range .Fields }}{{.Name}} {{.Type}} `json:"{{.Json}}" gorm:"{{.Gorm}}"`
    {{ end }}
} 


  