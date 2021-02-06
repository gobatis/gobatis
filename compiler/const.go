package compiler

const EOF = 0xFF

// token type
const (

)

// config element type
const (
	NT_CONFIGURATION       = "configuration"
	NT_PROPERTIES          = "properties"
	NT_SETTINGS            = "settings"
	NT_TYPE_ALIASES        = "typeAliases"
	NT_TYPE_HANDLERS       = "typeHandlers"
	NT_OBJECT_FACTORY      = "objectFactory"
	NT_PLUGINS             = "plugins"
	NT_ENVIRONMENTS        = "environments"
	NT_ENVIRONMENT         = "environment"
	NT_TRANSACTION_MANAGER = "transactionManager"
	NT_DATA_SOURCE         = "dataSource"
	NT_DATABASEldProvider  = "databaseld_provider"
	NT_MAPPERS             = "mappers"
)

// mapper element type
const (
	NT_CACHE      = "cache"
	NT_CACHE_REF  = "cache-ref"
	NT_RESULT_MAP = "resultMap"
	NT_SQL        = "sql"
	NT_INSERT     = "insert"
	NT_UPDATE     = "update"
	NT_DELETE     = "delete"
	NT_SELECT     = "select"
)

// sql word
const (
	SW_SELECT = "select"
	SW_WHERE  = "where"
	SW_FROM   = "from"
	SW_INSERT = "insert"
	SW_UPDATE = "update"
	SW_DELETE = "delete"
	SW_JOIN   = "join"
	SW_LEFT   = "left"
)
