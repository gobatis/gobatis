package dtd

import "fmt"

func ConfigElement(name string) (elem *Element, err error) {
	switch name {
	case MODULE:
		elem = Module
	case DATABASE_ID_PROVIDER:
		elem = DatabaseIdProvider
	case PROPERTIES:
		elem = Properties
	case PROPERTY:
		elem = Property
	case SETTINGS:
		elem = Settings
	case SETTING:
		elem = Setting
	case TYPE_ALIASES:
		elem = TypeAliases
	case TYPE_ALIAS:
		elem = ConfigTypeAlias
	case TYPE_HANDLERS:
		elem = TypeHandlers
	case TYPE_HANDLER:
		elem = TypeHandle
	case OBJECT_FACTORY:
		elem = ObjectFactory
	case OBJECT_WRAPPER_FACTORY:
		elem = ObjectWrapperFactory
	case REFLECT_FACTORY:
		elem = ReflectFactory
	case PLUGINS:
		elem = Plugins
	case PLUGIN:
		elem = Plugin
	case ENVIRONMENTS:
		elem = Environments
	case ENVIRONMENT:
		elem = Environment
	case TRANSACTION_MANAGER:
		elem = TransactionManager
	case DATA_SOURCE:
		elem = DataSource
	case MAPPERS:
		elem = Mappers
	case MAPPER:
		elem = ConfigMapper
	case PACKAGE:
		elem = Package
	default:
		err = fmt.Errorf("not match %s dtd element", name)
	}
	return
}

func MapperElement(name string) (elem *Element, err error) {
	switch name {
	case MAPPER:
		elem = Mapper
	case CACHE_REF:
		elem = CacheRef
	case CACHE:
		elem = Cache
	case RESULT_MAP:
		elem = ResultMap
	case CONSTRUCTOR:
		elem = Constructor
	case ID:
		elem = Id
	case RESULT:
		elem = Result
	case ID_ARG:
		elem = IdArg
	case ARG:
		elem = Arg
	case COLLECTION:
		elem = Collection
	case ASSOCIATION:
		elem = Association
	case DISCRIMINATOR:
		elem = Discriminator
	case CASE:
		elem = Case
	case TYPE_ALIAS:
		elem = TypeAlias
	case SELECT:
		elem = Select
	case INSERT:
		elem = Insert
	case SELECT_KEY:
		elem = SelectKey
	case UPDATE:
		elem = Update
	case DELETE:
		elem = Delete
	case INCLUDE:
		elem = Include
	case BIND:
		elem = Bind
	case SQL:
		elem = Sql
	case TRIM:
		elem = Trim
	case WHERE:
		elem = Where
	case SET:
		elem = Set
	case FOREACH:
		elem = Foreach
	case CHOOSE:
		elem = Choose
	case WHEN:
		elem = When
	case OTHERWISE:
		elem = Otherwise
	case IF:
		elem = If
	default:
		err = fmt.Errorf("not match %s dtd element", name)
	}

	return
}
