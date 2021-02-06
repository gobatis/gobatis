package compiler

const (
	STAT_LITERAL               = iota // 期待文本
	STAT_START_TAG                    // 期待标签开始  <mapper> | </mapper>
	STAT_END_TAG                      // 期待标签结束  <mapper> | </mapper>
	STAT_SELF_END_TAG                 // 期待标签结束  <mapper> | </mapper>
	STAT_START_TAG_NAME               // 期待解析标签名称
	STAT_END_TAG_NAME                 // 期待解析结束标签名称
	STAT_ATTRIBUTE                    // 期待解析属性开始
	STAT_ATTRIBUTE_NAME               // 期待解析属性名称
	STAT_ATTRIBUTE_EQUAL              // 期待解析属性等号
	STAT_ATTRIBUTE_VALUE_START        // 期待解析属性值
	STAT_ATTRIBUTE_VALUE_END          // 期待解析属性值
	STAT_CLOSING_TAG                  // 期待标签闭合
	STAT_IN_VALUE_NQ                  // 无引号的值
	STAT_IN_VALUE_SQ                  // 单引号值
	STAT_IN_VALUE_DQ                  // 双引号值
	STAT_CLOSING_OPEN_TAG
	STAT_OPENING_NORMAL_COMMENT
	STAT_IM_NORMAL_COMMENT
	STAT_IN_SHORT_COMMENT
	STAT_CLOSING_NORMAL_COMMENT
)
