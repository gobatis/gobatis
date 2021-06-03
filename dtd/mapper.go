package dtd

var Mapper = &Element{
	Name: MAPPER,
	Nodes: map[string]int{
		CACHE_REF:     AT_MOST_ONCE, // TODO 和协议有区别
		CACHE:         AT_MOST_ONCE, // TODO 和协议有区别
		PARAMETER_MAP: ANY_TIMES,
		SQL:           ANY_TIMES,
		INSERT:        ANY_TIMES,
		UPDATE:        ANY_TIMES,
		DELETE:        ANY_TIMES,
		SELECT:        ANY_TIMES,
	},
	Attributes: map[string]int{
		NAMESPACE: IMPLIED,
	},
}

var CacheRef = &Element{
	Name: CACHE_REF,
	Attributes: map[string]int{
		NAMESPACE: REQUIRED,
	},
}

var Cache = &Element{
	Name: CACHE,
	Nodes: map[string]int{
		PROPERTY: ANY_TIMES,
	},
	Attributes: map[string]int{
		TYPE:           IMPLIED,
		EVICTION:       IMPLIED,
		FLUSH_INTERVAL: IMPLIED,
		SIZE:           IMPLIED,
		READ_ONLY:      IMPLIED,
		BLOCKING:       IMPLIED,
	},
}

// 废弃
//var ParameterMap = &Element{
//	Name:       PARAMETER_MAP,
//	Nodes: map[string]int{
//		PARAMETER:A
//	},
//	Attributes: nil,
//}

var ResultMap = &Element{
	Name: RESULT_MAP,
	Nodes: map[string]int{
		CONSTRUCTOR:   AT_MOST_ONCE,
		ID:            ANY_TIMES,
		RESULT:        ANY_TIMES,
		ASSOCIATION:   ANY_TIMES,
		COLLECTION:    ANY_TIMES,
		DISCRIMINATOR: AT_MOST_ONCE,
	},
	Attributes: map[string]int{
		ID:           REQUIRED,
		TYPE:         REQUIRED,
		EXTENDS:      IMPLIED,
		AUTO_MAPPING: IMPLIED,
	},
}

var Constructor = &Element{
	Name: CONSTRUCTOR,
	Nodes: map[string]int{
		ID_ARG: ANY_TIMES,
		ARG:    ANY_TIMES,
	},
}

var Id = &Element{
	Name: ID,
	Attributes: map[string]int{
		PROPERTY:     IMPLIED,
		GO_TYPE:      IMPLIED,
		COLUMN:       IMPLIED,
		TYPE_HANDLER: IMPLIED,
	},
}

var Result = &Element{
	Name: RESULT,
	Attributes: map[string]int{
		PROPERTY:     IMPLIED,
		GO_TYPE:      IMPLIED,
		COLUMN:       IMPLIED,
		TYPE_HANDLER: IMPLIED,
	},
}

var IdArg = &Element{
	Name: ID_ARG,
	Attributes: map[string]int{
		GO_TYPE:       IMPLIED,
		COLUMN:        IMPLIED,
		TYPE_HANDLER:  IMPLIED,
		SELECT:        IMPLIED,
		RESULT_MAP:    IMPLIED,
		NAME:          IMPLIED,
		COLUMN_PREFIX: IMPLIED,
	},
}

var Arg = &Element{
	Name: ARG,
	Attributes: map[string]int{
		GO_TYPE:       IMPLIED,
		COLUMN:        IMPLIED,
		TYPE_HANDLER:  IMPLIED,
		SELECT:        IMPLIED,
		RESULT_MAP:    IMPLIED,
		NAME:          IMPLIED,
		COLUMN_PREFIX: IMPLIED,
	},
}

var Collection = &Element{
	Name: COLLECTION,
	Nodes: map[string]int{
		CONSTRUCTOR:   AT_MOST_ONCE,
		ID:            ANY_TIMES,
		RESULT:        ANY_TIMES,
		ASSOCIATION:   ANY_TIMES,
		COLLECTION:    ANY_TIMES,
		DISCRIMINATOR: AT_MOST_ONCE,
	},
	Attributes: map[string]int{
		PROPERTY:        REQUIRED,
		COLUMN:          IMPLIED,
		GO_TYPE:         IMPLIED,
		OF_TYPE:         IMPLIED,
		SELECT:          IMPLIED,
		RESULT_MAP:      IMPLIED,
		TYPE_HANDLER:    IMPLIED,
		NOT_NULL_COLUMN: IMPLIED,
		COLUMN_PREFIX:   IMPLIED,
		RESULT_SET:      IMPLIED,
		FOREIGN_COLUMN:  IMPLIED,
		AUTO_MAPPING:    IMPLIED,
		FETCH_TYPE:      IMPLIED,
	},
}

var Association = &Element{
	Name: ASSOCIATION,
	Nodes: map[string]int{
		CONSTRUCTOR:   AT_MOST_ONCE,
		ID:            ANY_TIMES,
		RESULT:        ANY_TIMES,
		ASSOCIATION:   ANY_TIMES,
		COLLECTION:    ANY_TIMES,
		DISCRIMINATOR: AT_MOST_ONCE,
	},
	Attributes: map[string]int{
		PROPERTY:        REQUIRED,
		COLUMN:          IMPLIED,
		GO_TYPE:         IMPLIED,
		SELECT:          IMPLIED,
		RESULT_MAP:      IMPLIED,
		TYPE_HANDLER:    IMPLIED,
		NOT_NULL_COLUMN: IMPLIED,
		COLUMN_PREFIX:   IMPLIED,
		RESULT_SET:      IMPLIED,
		FOREIGN_COLUMN:  IMPLIED,
		AUTO_MAPPING:    IMPLIED,
		FETCH_TYPE:      IMPLIED,
	},
}

var Discriminator = &Element{
	Name: DISCRIMINATOR,
	Nodes: map[string]int{
		CASE: AT_LEAST_ONCE,
	},
	Attributes: map[string]int{
		COLUMN:       IMPLIED,
		GO_TYPE:      REQUIRED,
		TYPE_HANDLER: IMPLIED,
	},
}

var Case = &Element{
	Name: CASE,
	Nodes: map[string]int{
		CONSTRUCTOR:   AT_MOST_ONCE,
		ID:            ANY_TIMES,
		RESULT:        ANY_TIMES,
		ASSOCIATION:   ANY_TIMES,
		COLLECTION:    ANY_TIMES,
		DISCRIMINATOR: AT_MOST_ONCE,
	},
	Attributes: map[string]int{
		VALUE:       REQUIRED,
		RESULT_MAP:  IMPLIED,
		//RESULT_TYPE: IMPLIED,
	},
}

var TypeAlias = &Element{
	Name: TYPE_ALIAS,
	Attributes: map[string]int{
		ALIAS: REQUIRED,
		TYPE:  REQUIRED,
	},
}

var Select = &Element{
	Name: SELECT,
	Nodes: map[string]int{
		PCDATA:  ANY_TIMES,
		INCLUDE: ANY_TIMES,
		TRIM:    ANY_TIMES,
		WHERE:   ANY_TIMES,
		SET:     ANY_TIMES,
		FOREACH: ANY_TIMES,
		CHOOSE:  ANY_TIMES,
		IF:      ANY_TIMES,
		BIND:    ANY_TIMES,
	},
	Attributes: map[string]int{
		ID:              REQUIRED,
		PARAMETER_MAP:   IMPLIED,
		PARAMETER:       IMPLIED,
		RESULT_MAP:      IMPLIED,
		//RESULT_TYPE:     IMPLIED,
		RESULT:          IMPLIED,
		RESULT_SET_TYPE: IMPLIED,
		STATEMENT_TYPE:  IMPLIED,
		FETCH_SIZE:      IMPLIED,
		TIMEOUT:         IMPLIED,
		FLUSH_CACHE:     IMPLIED,
		USE_CACHE:       IMPLIED,
		DATABASE_ID:     IMPLIED,
		LANG:            IMPLIED,
		RESULT_ORDERED:  IMPLIED,
		RESULT_SETS:     IMPLIED,
	},
}

var Insert = &Element{
	Name: INSERT,
	Nodes: map[string]int{
		PCDATA:     ANY_TIMES,
		SELECT_KEY: ANY_TIMES,
		INCLUDE:    ANY_TIMES,
		TRIM:       ANY_TIMES,
		WHERE:      ANY_TIMES,
		SET:        ANY_TIMES,
		FOREACH:    ANY_TIMES,
		CHOOSE:     ANY_TIMES,
		IF:         ANY_TIMES,
		BIND:       ANY_TIMES,
	},
	Attributes: map[string]int{
		ID:                REQUIRED,
		PARAMETER_MAP:     IMPLIED,
		PARAMETER:         IMPLIED,
		TIMEOUT:           IMPLIED,
		FLUSH_CACHE:       IMPLIED,
		STATEMENT_TYPE:    IMPLIED,
		KEY_PROPERTY:      IMPLIED,
		USE_GENERATE_KEYS: IMPLIED,
		KEY_COLUMN:        IMPLIED,
		DATABASE_ID:       IMPLIED,
		LANG:              IMPLIED,
	},
}

var SelectKey = &Element{
	Name: SELECT_KEY,
	Nodes: map[string]int{
		PCDATA:  ANY_TIMES,
		INCLUDE: ANY_TIMES,
		TRIM:    ANY_TIMES,
		WHERE:   ANY_TIMES,
		SET:     ANY_TIMES,
		FOREACH: ANY_TIMES,
		CHOOSE:  ANY_TIMES,
		IF:      ANY_TIMES,
		BIND:    ANY_TIMES,
	},
	Attributes: map[string]int{
		//RESULT_TYPE:    IMPLIED,
		STATEMENT_TYPE: IMPLIED,
		KEY_PROPERTY:   IMPLIED,
		KEY_COLUMN:     IMPLIED,
		ORDER:          IMPLIED,
		DATABASE_ID:    IMPLIED,
	},
}

var Update = &Element{
	Name: UPDATE,
	Nodes: map[string]int{
		PCDATA:     ANY_TIMES,
		SELECT_KEY: ANY_TIMES,
		INCLUDE:    ANY_TIMES,
		TRIM:       ANY_TIMES,
		WHERE:      ANY_TIMES,
		SET:        ANY_TIMES,
		FOREACH:    ANY_TIMES,
		CHOOSE:     ANY_TIMES,
		IF:         ANY_TIMES,
		BIND:       ANY_TIMES,
	},
	Attributes: map[string]int{
		ID:                REQUIRED,
		PARAMETER_MAP:     IMPLIED,
		PARAMETER:         IMPLIED,
		TIMEOUT:           IMPLIED,
		FLUSH_CACHE:       IMPLIED,
		STATEMENT_TYPE:    IMPLIED,
		KEY_PROPERTY:      IMPLIED,
		USE_GENERATE_KEYS: IMPLIED,
		KEY_COLUMN:        IMPLIED,
		DATABASE_ID:       IMPLIED,
		LANG:              IMPLIED,
	},
}

var Delete = &Element{
	Name: DELETE,
	Nodes: map[string]int{
		PCDATA:  ANY_TIMES,
		INCLUDE: ANY_TIMES,
		TRIM:    ANY_TIMES,
		WHERE:   ANY_TIMES,
		SET:     ANY_TIMES,
		FOREACH: ANY_TIMES,
		CHOOSE:  ANY_TIMES,
		IF:      ANY_TIMES,
		BIND:    ANY_TIMES,
	},
	Attributes: map[string]int{
		ID:             REQUIRED,
		PARAMETER_MAP:  IMPLIED,
		PARAMETER:      IMPLIED,
		TIMEOUT:        IMPLIED,
		FLUSH_CACHE:    IMPLIED,
		STATEMENT_TYPE: IMPLIED,
		DATABASE_ID:    IMPLIED,
		LANG:           IMPLIED,
	},
}

var Include = &Element{
	Name: INCLUDE,
	Nodes: map[string]int{
		PROPERTY: ANY_TIMES, // TODO 和协议有区别
	},
	Attributes: map[string]int{
		REFID: REQUIRED,
	},
}

var Bind = &Element{
	Name: BIND,
	Attributes: map[string]int{
		NAME:  REQUIRED,
		VALUE: REQUIRED,
	},
}

var Sql = &Element{
	Name: SQL,
	Nodes: map[string]int{
		PCDATA:  ANY_TIMES,
		INCLUDE: ANY_TIMES,
		TRIM:    ANY_TIMES,
		WHERE:   ANY_TIMES,
		SET:     ANY_TIMES,
		FOREACH: ANY_TIMES,
		CHOOSE:  ANY_TIMES,
		IF:      ANY_TIMES,
		BIND:    ANY_TIMES,
	},
	Attributes: map[string]int{
		ID:          REQUIRED,
		LANG:        IMPLIED,
		DATABASE_ID: IMPLIED,
	},
}

var Trim = &Element{
	Name: TRIM,
	Nodes: map[string]int{
		PCDATA:  ANY_TIMES,
		INCLUDE: ANY_TIMES,
		TRIM:    ANY_TIMES,
		WHERE:   ANY_TIMES,
		SET:     ANY_TIMES,
		FOREACH: ANY_TIMES,
		CHOOSE:  ANY_TIMES,
		IF:      ANY_TIMES,
		BIND:    ANY_TIMES,
	},
	Attributes: map[string]int{
		PREFIX:           IMPLIED,
		PREFIX_OVERRIDES: IMPLIED,
		SUFFIX:           IMPLIED,
		SUFFIX_OVERRIDES: IMPLIED,
	},
}

var Where = &Element{
	Name: WHERE,
	Nodes: map[string]int{
		PCDATA:  ANY_TIMES,
		INCLUDE: ANY_TIMES,
		TRIM:    ANY_TIMES,
		WHERE:   ANY_TIMES,
		SET:     ANY_TIMES,
		FOREACH: ANY_TIMES,
		CHOOSE:  ANY_TIMES,
		IF:      ANY_TIMES,
		BIND:    ANY_TIMES,
	},
}

var Set = &Element{
	Name: SET,
	Nodes: map[string]int{
		PCDATA:  ANY_TIMES,
		INCLUDE: ANY_TIMES,
		TRIM:    ANY_TIMES,
		WHERE:   ANY_TIMES,
		SET:     ANY_TIMES,
		FOREACH: ANY_TIMES,
		CHOOSE:  ANY_TIMES,
		IF:      ANY_TIMES,
		BIND:    ANY_TIMES,
	},
}

var Foreach = &Element{
	Name: FOREACH,
	Nodes: map[string]int{
		PCDATA:  ANY_TIMES,
		INCLUDE: ANY_TIMES,
		TRIM:    ANY_TIMES,
		WHERE:   ANY_TIMES,
		SET:     ANY_TIMES,
		FOREACH: ANY_TIMES,
		CHOOSE:  ANY_TIMES,
		IF:      ANY_TIMES,
		BIND:    ANY_TIMES,
	},
	Attributes: map[string]int{
		COLLECTION: REQUIRED,
		ITEM:       IMPLIED,
		INDEX:      IMPLIED,
		OPEN:       IMPLIED,
		CLOSE:      IMPLIED,
		SEPARATOR:  IMPLIED,
	},
}
var Choose = &Element{
	Name: CHOOSE,
	Nodes: map[string]int{
		WHEN:      ANY_TIMES,
		OTHERWISE: AT_MOST_ONCE,
	},
}

var When = &Element{
	Name: WHEN,
	Nodes: map[string]int{
		PCDATA:  ANY_TIMES,
		INCLUDE: ANY_TIMES,
		TRIM:    ANY_TIMES,
		WHERE:   ANY_TIMES,
		SET:     ANY_TIMES,
		FOREACH: ANY_TIMES,
		CHOOSE:  ANY_TIMES,
		IF:      ANY_TIMES,
		BIND:    ANY_TIMES,
	},
	Attributes: map[string]int{
		TEST: REQUIRED,
	},
}

var Otherwise = &Element{
	Name: OTHERWISE,
	Nodes: map[string]int{
		PCDATA:  ANY_TIMES,
		INCLUDE: ANY_TIMES,
		TRIM:    ANY_TIMES,
		WHERE:   ANY_TIMES,
		SET:     ANY_TIMES,
		FOREACH: ANY_TIMES,
		CHOOSE:  ANY_TIMES,
		IF:      ANY_TIMES,
		BIND:    ANY_TIMES,
	},
}

var If = &Element{
	Name: IF,
	Nodes: map[string]int{
		PCDATA:  ANY_TIMES,
		INCLUDE: ANY_TIMES,
		TRIM:    ANY_TIMES,
		WHERE:   ANY_TIMES,
		SET:     ANY_TIMES,
		FOREACH: ANY_TIMES,
		CHOOSE:  ANY_TIMES,
		IF:      ANY_TIMES,
		BIND:    ANY_TIMES,
	},
	Attributes: map[string]int{
		TEST: REQUIRED,
	},
}
