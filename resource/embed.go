package resource

import (
	"embed"
)

//嵌入web模板目录
//go:embed templates
var Templates embed.FS
