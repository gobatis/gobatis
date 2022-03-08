create table if not exists public.test_entities
(
    id                            bigserial,
    t_int8                        smallint,
    t_bigint                      integer,
    t_int                         bigint,
    t_decimal                     numeric(32, 18),
    t_numeric                     numeric(32, 18),
    t_real                        real,
    t_double_precision            double precision,
    t_small_serial                smallserial,
    t_serial                      serial,
    t_big_serial                  bigserial,
    t_money                       money,
    t_char                        varchar,
    t_text                        text,
    t_timestamp_without_time_zone timestamp,
    t_timestamp_with_time_zone    timestamp with time zone,
    t_date                        date,
    t_time_without_time_zone      time,
    t_time_with_time_zone         time with time zone,
    t_interval                    interval,
    t_boolean                     boolean
);

create table if not exists public.users
(
    id         serial
        constraint users_pk
            primary key,
    name       varchar,
    age        integer,
    "from"     varchar,
    vip        boolean,
    created_at timestamp default CURRENT_TIMESTAMP,
    tags       character varying[]
);

create table if not exists public.orders
(
    id              bigserial
        constraint orders_pk
            primary key,
    member_id       bigint,
    shop_id         bigint,
    amount          numeric,
    products        bigint[],
    pay_channel     varchar,
    pay_status      integer,
    dispatch_status integer,
    created_at      time default CURRENT_TIME,
    updated_at      time
);

create table if not exists public.members
(
    id         bigserial
        constraint members_pk
            primary key,
    username   varchar,
    email      varchar,
    mobile     varchar,
    password   varchar,
    status     integer default 0,
    created_at time    default CURRENT_TIME,
    updated_at time
);

create table if not exists public.shops
(
    id          bigint not null
        constraint shops_pk
            primary key,
    name        varchar,
    description text,
    created_at  time default CURRENT_TIME,
    updated_at  time
);

create unique index if not exists shops_id_uindex
    on public.shops (id);

create table if not exists public.products
(
    id          bigserial
        constraint products_pk
            primary key,
    name        varchar,
    tags        character varying[],
    description text,
    stock       bigint,
    up          boolean,
    sell        integer,
    price       numeric,
    created_at  time default CURRENT_TIME,
    updated_at  time
);

create table if not exists public.types
(
    id                        serial
        constraint types_pk
            primary key,
    sid                       varchar,
    source                    varchar,
    t_bigint                  bigint,
    t_int8                    bigint,
    t_bigserial               bigserial,
    t_serial8                 bigserial,
    t_bit                     bit,
    t_bit_varying             bit varying,
    t_varbit                  bit varying,
    t_boolean                 boolean,
    t_bool                    boolean,
    t_box                     box,
    t_bytea                   bytea,
    t_character               char(24),
    t_char                    char(24),
    t_character_varying       varchar,
    t_varchar                 varchar,
    t_cidr                    cidr,
    t_circle                  circle,
    t_date                    date,
    t_double_precision        double precision,
    t_float8                  double precision,
    t_inet                    inet,
    t_integer                 integer,
    t_int                     integer,
    t_int4                    integer,
    t_interval                interval,
    t_json                    json,
    t_jsonb                   jsonb,
    t_line                    line,
    t_lseg                    lseg,
    t_macaddr                 macaddr,
    t_macaddr8                macaddr8,
    t_money                   money,
    t_numeric                 numeric,
    t_decimal                 numeric,
    t_path                    path,
    t_pg_lsn                  pg_lsn,
    t_pg_snapshot             pg_snapshot,
    t_point                   point,
    t_polygon                 polygon,
    t_real                    real,
    t_float4                  real,
    t_smallint                smallint,
    t_int2                    smallint,
    t_smallserial             smallserial,
    t_serial2                 smallserial,
    t_serial                  serial,
    t_serial4                 serial,
    t_text                    text,
    t_time                    time,
    t_time_with_timezone      time with time zone,
    t_timetz                  time with time zone,
    t_timestamp               timestamp,
    t_timestamp_with_timezone timestamp with time zone,
    t_timestamptz             timestamp with time zone,
    t_tsquery                 tsquery,
    t_tsvector                tsvector,
    t_txid_snapshot           txid_snapshot,
    t_uuid                    uuid,
    t_xml                     xml,
    deleted                   boolean
);

create unique index if not exists types_sid_uindex
    on public.types (sid);

create table if not exists public.array_types
(
    id                        serial
        constraint array_types_pk
            primary key,
    sid                       varchar,
    source                    varchar,
    t_bigint                  bigint[],
    t_int8                    bigint[],
    t_bit                     bit[],
    t_bit_varying             bit varying[],
    t_varbit                  bit varying[],
    t_boolean                 boolean[],
    t_bool                    boolean[],
    t_box                     box[],
    t_bytea                   bytea[],
    t_character               char(24)[],
    t_char                    char(24)[],
    t_character_varying       character varying[],
    t_varchar                 character varying[],
    t_cidr                    cidr[],
    t_circle                  circle[],
    t_date                    date[],
    t_double_precision        double precision[],
    t_float8                  double precision[],
    t_inet                    inet[],
    t_integer                 integer[],
    t_int                     integer[],
    t_int4                    integer[],
    t_interval                interval[],
    t_json                    json[],
    t_jsonb                   jsonb[],
    t_line                    line[],
    t_lseg                    lseg[],
    t_macaddr                 macaddr[],
    t_macaddr8                macaddr8[],
    t_money                   money[],
    t_numeric                 numeric[],
    t_decimal                 numeric[],
    t_path                    path[],
    t_pg_lsn                  pg_lsn[],
    t_pg_snapshot             pg_snapshot[],
    t_point                   point[],
    t_polygon                 polygon[],
    t_real                    real[],
    t_float4                  real[],
    t_smallint                smallint[],
    t_int2                    smallint[],
    t_text                    text[],
    t_time                    time without time zone[],
    t_time_with_timezone      time with time zone[],
    t_timetz                  time with time zone[],
    t_timestamp               timestamp without time zone[],
    t_timestamp_with_timezone timestamp with time zone[],
    t_timestamptz             timestamp with time zone[],
    t_tsquery                 tsquery[],
    t_tsvector                tsvector[],
    t_txid_snapshot           txid_snapshot[],
    t_uuid                    uuid[],
    t_xml                     xml[],
    deleted                   boolean
);

create unique index if not exists array_types_sid_uindex
    on public.array_types (sid);

