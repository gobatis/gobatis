package compiler

const EOF = 0xFF

const (
	TS_LITERAL               = iota // 期待文本
	TS_START_TAG                    // 期待标签开始  <mapper> | </mapper>
	TS_SELF_END_TAG                 // 期待标签结束  <mapper> | </mapper>
	TS_START_TAG_NAME               // 期待解析标签名称
	TS_END_TAG_NAME                 // 期待解析结束标签名称
	TS_ATTRIBUTE                    // 期待解析属性开始
	TS_ATTRIBUTE_NAME               // 期待解析属性名称
	TS_ATTRIBUTE_EQUAL              // 期待解析属性等号
	TS_ATTRIBUTE_VALUE_START        // 期待解析属性值
	TS_ATTRIBUTE_VALUE_END          // 期待解析属性值
	TS_COMMENT_START                // 解析注释开始
	TS_COMMENT_END                  // 解析注释结束
	TS_DOCTYPPE                     // 解析文档声明
	TS_STATEMENT                    // 解析声明类型
)

const (
	TS_BLANK               = iota // 过滤空白
	TS_ID                         // 期待ID
	TS_SQL_VAR_START              // 期待解析模板变量开始
	TS_SQL_VAR_VALUE_START        // 期待解析模板变量内容开始
	TS_SQL_VAR_VALUE_END          // 期待解析模板变量内容结束
	TS_SQL_VAR_END                // 期待解析模板变量结束
)

const (
	TT_TEXT         = "text"
	TT_START_TAG    = "tag:start"
	TT_END_TAG      = "tag:end"
	TT_SELF_END_TAG = "tag:self-end"
	TT_ATTR_VALUE   = "attr:value"
	TT_ATTR_NAME    = "attr:name"
)

const (
	TT_ID           = "id"
	TT_COMMA        = ","
	TT_EQUAL        = "="
	TT_SQL_STRUCT   = "sql:struct"
	TT_SQL_DOT      = "sql:."
	TT_SQL_VAR      = "sql:var"
	TT_PROPERTY_VAR = "property:var"
	TT_PROPERTY_DS  = "property:ds"
	TT_PROPERTY_VAL = "property:val"
)

const (
	ST_TEXT = "text"
	ST_NODE = "node"
)

const (
	DOCTYPE = "doctype"
)
