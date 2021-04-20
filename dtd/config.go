package dtd

var Configuration = &Element{
	Name:  CONFIGURATION,
	Nodes: map[string]int{
		PROPERTIES:             AT_MOST_ONCE,
		MODULE:                 AT_MOST_ONCE,
		SETTINGS:               AT_MOST_ONCE,
		TYPE_ALIASES:           AT_MOST_ONCE,
		TYPE_HANDLERS:          AT_MOST_ONCE,
		OBJECT_FACTORY:         AT_MOST_ONCE,
		OBJECT_WRAPPER_FACTORY: AT_MOST_ONCE,
		REFLECT_FACTORY:        AT_MOST_ONCE,
		PLUGINS:                AT_MOST_ONCE,
		ENVIRONMENTS:           AT_MOST_ONCE,
		DATABASE_ID_PROVIDER:   AT_MOST_ONCE,
		MAPPERS:                AT_MOST_ONCE,
	},
}

var Module = &Element{
	Name: MODULE,
	Attributes: map[string]int{
		BASE: REQUIRED,
	},
}

var DatabaseIdProvider = &Element{
	Name: DATABASE_ID_PROVIDER,
	Nodes: map[string]int{
		PROPERTY: ANY_TIMES,
	},
	Attributes: map[string]int{
		TYPE: REQUIRED,
	},
}

var Properties = &Element{
	Name: PROPERTIES,
	Nodes: map[string]int{
		PROPERTY: ANY_TIMES,
	},
	Attributes: map[string]int{
		RESOURCE: IMPLIED,
		URL:      IMPLIED,
	},
}

var Property = &Element{
	Name: PROPERTY,
	Attributes: map[string]int{
		NAME:  REQUIRED,
		VALUE: REQUIRED,
	},
}

var Settings = &Element{
	Name: SETTINGS,
	Nodes: map[string]int{
		SETTING: ANY_TIMES,
	},
}

var Setting = &Element{
	Name: SETTING,
	Attributes: map[string]int{
		NAME:  REQUIRED,
		VALUE: REQUIRED,
	},
}

var TypeAliases = &Element{
	Name: TYPE_ALIASES,
	Nodes: map[string]int{
		TYPE_ALIAS: ANY_TIMES,
		PACKAGE:    ANY_TIMES,
	},
}

var ConfigTypeAlias = &Element{
	Name: TYPE_ALIAS,
	Attributes: map[string]int{
		TYPE:  REQUIRED,
		ALIAS: IMPLIED,
	},
}

var TypeHandlers = &Element{
	Name: TYPE_HANDLERS,
	Nodes: map[string]int{
		TYPE_HANDLER: ANY_TIMES,
		PACKAGE:      ANY_TIMES,
	},
}

var TypeHandle = &Element{
	Name: TYPE_HANDLER,
	Attributes: map[string]int{
		GO_TYPE: IMPLIED,
		HANDLER: REQUIRED,
	},
}

var ObjectFactory = &Element{
	Name: OBJECT_FACTORY,
	Nodes: map[string]int{
		PROPERTY: ANY_TIMES,
	},
	Attributes: map[string]int{
		TYPE: REQUIRED,
	},
}

var ObjectWrapperFactory = &Element{
	Name: OBJECT_WRAPPER_FACTORY,
	Attributes: map[string]int{
		TYPE: REQUIRED,
	},
}

var ReflectFactory = &Element{
	Name: REFLECT_FACTORY,
	Attributes: map[string]int{
		TYPE: REQUIRED,
	},
}

var Plugins = &Element{
	Name: PLUGINS,
	Nodes: map[string]int{
		PLUGIN: AT_LEAST_ONCE,
	},
}

var Plugin = &Element{
	Name: PLUGIN,
	Nodes: map[string]int{
		PROPERTY: ANY_TIMES,
	},
	Attributes: map[string]int{
		INTERCEPTOR: REQUIRED,
	},
}

var Environments = &Element{
	Name: ENVIRONMENTS,
	Nodes: map[string]int{
		ENVIRONMENT: AT_LEAST_ONCE,
	},
	Attributes: map[string]int{
		DEFAULT: REQUIRED,
	},
}

var Environment = &Element{
	Name: ENVIRONMENT,
	Nodes: map[string]int{
		TRANSACTION_MANAGER: ONCE,
		DATA_SOURCE:         ONCE,
	},
	Attributes: map[string]int{
		ID: REQUIRED,
	},
}

var TransactionManager = &Element{
	Name: TRANSACTION_MANAGER,
	Nodes: map[string]int{
		PROPERTY: ANY_TIMES,
	},
	Attributes: map[string]int{
		TYPE: REQUIRED,
	},
}

var DataSource = &Element{
	Name: DATA_SOURCE,
	Nodes: map[string]int{
		PROPERTY: ANY_TIMES,
	},
	Attributes: map[string]int{
		TYPE: REQUIRED,
	},
}

var Mappers = &Element{
	Name: MAPPERS,
	Nodes: map[string]int{
		MAPPER:  ANY_TIMES,
		PACKAGE: ANY_TIMES,
	},
}

var ConfigMapper = &Element{
	Name: MAPPER,
	Attributes: map[string]int{
		RESOURCE: IMPLIED,
		URL:      IMPLIED,
		STRUCT:   IMPLIED,
	},
}

var Package = &Element{
	Name: PACKAGE,
	Attributes: map[string]int{
		NAME: REQUIRED,
	},
}
