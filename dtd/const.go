package dtd

const (
	ANY_TIMES     = iota // * 任意次
	AT_MOST_ONCE         // ? 最多一次
	AT_LEAST_ONCE        // + 至少一次
	ONCE                 // 必须出现一次
)

const (
	IMPLIED = iota
	REQUIRED
)

const (
	PCDATA = "PCDATA" // 元素内容可以同时出现文本和标签
	CDATA  = "CDATA"  // 属性值以字符串来解析
)

const (
	CONFIGURATION          = "configuration"
	PROPERTIES             = "properties"
	MODULE                 = "module"
	BASE                   = "base"
	SETTINGS               = "settings"
	TYPE_ALIASES           = "typeAliases"
	TYPE_HANDLERS          = "typeHandlers"
	OBJECT_FACTORY         = "objectFactory"
	OBJECT_WRAPPER_FACTORY = "objectWrapperFactory"
	REFLECT_FACTORY        = "reflectFactory"
	PLUGINS                = "plugins"
	PLUGIN                 = "plugin"
	INTERCEPTOR            = "interceptor"
	ENVIRONMENTS           = "environments"
	DATABASE_ID_PROVIDER   = "databaseIdProvider"
	MAPPERS                = "mappers"
	TYPE                   = "type"
	PROPERTY               = "property"
	RESOURCE               = "resource"
	URL                    = "url"
	NAME                   = "name"
	VALUE                  = "value"
	SETTING                = "setting"
	TYPE_ALIAS             = "typeAlias"
	ALIAS                  = "alias"
	TYPE_HANDLER           = "typeHandler"
	PACKAGE                = "package"
	GO_TYPE                = "goType"
	HANDLER                = "handler"
	ENVIRONMENT            = "environment"
	DEFAULT                = "default"
	TRANSACTION_MANAGER    = "transactionManager"
	DATA_SOURCE            = "dataSource"
	STRUCT                 = "struct"
)

const (
	MAPPER            = "mapper"
	CACHE_REF         = "cache-ref"
	CACHE             = "cache"
	RESULT_MAP        = "resultMap"
	PARAMETER_MAP     = "parameterMap"
	PARAMETER         = "parameter"
	PARAMETER_TYPE    = "parameterType"
	SQL               = "sql"
	INSERT            = "insert"
	UPDATE            = "update"
	DELETE            = "delete"
	SELECT            = "select"
	NAMESPACE         = "namespace"
	EVICTION          = "eviction"
	FLUSH_INTERVAL    = "flushInterval"
	SIZE              = "size"
	READ_ONLY         = "readOnly"
	BLOCKING          = "blocking"
	MODE              = "mode"
	SCALE             = "scale"
	ID                = "id"
	EXTENDS           = "extends"
	AUTO_MAPPING      = "autoMapping"
	CONSTRUCTOR       = "constructor"
	ID_ARG            = "idArg"
	ARG               = "arg"
	COLUMN            = "column"
	COLUMN_PREFIX     = "columnPrefix"
	COLLECTION        = "collection"
	RESULT            = "result"
	ASSOCIATION       = "association"
	DISCRIMINATOR     = "discriminator"
	OF_TYPE           = "ofType"
	NOT_NULL_COLUMN   = "notNullColumn"
	RESULT_SET        = "resultSet"
	FOREIGN_COLUMN    = "foreignColumn"
	FETCH_TYPE        = "fetchType"
	CASE              = "case"
	INCLUDE           = "include"
	REFID             = "refid"
	TRIM              = "trim"
	WHERE             = "where"
	SET               = "set"
	FOREACH           = "foreach"
	CHOOSE            = "choose"
	IF                = "if"
	BIND              = "bind"
	RESULT_TYPE       = "resultType"
	RESULT_SET_TYPE   = "resultSetType"
	STATEMENT_TYPE    = "statementType"
	FETCH_SIZE        = "fetchSize"
	TIMEOUT           = "timeout"
	FLUSH_CACHE       = "flushCache"
	USE_CACHE         = "useCache"
	DATABASE_ID       = "databaseId"
	LANG              = "lang"
	RESULT_ORDERED    = "resultOrdered"
	RESULT_SETS       = "resultSets"
	SELECT_KEY        = "selectKey"
	KEY_PROPERTY      = "keyProperty"
	USE_GENERATE_KEYS = "useGenerateKeys"
	KEY_COLUMN        = "keyColumn"
	ORDER             = "order"
	PREFIX            = "prefix"
	PREFIX_OVERRIDES  = "prefixOverrides"
	SUFFIX            = "suffix"
	SUFFIX_OVERRIDES  = "suffixOverrides"
	ITEM              = "item"
	INDEX             = "index"
	OPEN              = "open"
	CLOSE             = "close"
	SEPARATOR         = "separator"
	WHEN              = "when"
	OTHERWISE         = "otherwise"
	TEST              = "TEST"
)
