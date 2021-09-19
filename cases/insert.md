# Insert test cases



## Common insert

```xml
<insert id="TestInsert" parameter="row">
    insert t_table(t_char, t_bool) values( #{row.TChar}, #{row.TBool} );
</insert>
```

## Insert bind return field

```xml
<insert id="TestInsertBindFiled" parameter="row" selectKey="id">
    insert t_table(t_char, t_bool) values( #{row.TChar}, #{row.TBool} ) returning id;
</insert>
```

## Insert and return fields
```xml
<insert id="TestInsertReturnFields" parameter="row" result="id,t_char">
    insert t_table(t_char, t_bool) values( #{row.TChar}, #{row.TBool} ) returning id,t_char;
</insert>
```



## Insert with default value

```xml
<insert id="TestInsertReturnFields" parameter="row" result="id,t_char">
    insert t_table(
  		t_char, 
  		t_bool
  	)
  	values(
  		#{ row.TChar == '' ? nil : row.TChar },
		  #{row.TBool}
  	) returning id,t_char;
</insert>
```

## Common inserter

```xml
<insert id="TestInserter" parameter="row">
  	<inserter table="'t_table'" data="row">
      	<field name="'t_char'">#{row.TChar}</field>
        <field name="'t_bool'">#{row.TBool}</field>
  	</inserter>
</insert>
```

## Inserter bind return field 

```xml
<insert id="TestInserter" parameter="row" selectKey="id">
  	<inserter table="'t_table'" data="row">
      	<field name="'t_char'">#{row.TChar}</field>
        <field name="'t_bool'">#{row.TBool}</field>
  	</inserter> returning id;
</insert>
```

## Inserter returning fields

```xml
<insert id="TestInserter" parameter="row" result="id, t_char, t_bool">
  	<inserter table="'t_table'" data="row">
      	<field name="'t_char'">#{row.TChar}</field>
        <field name="'t_bool'">#{row.TBool}</field>
  	</inserter> returning id, t_char, t_bool;
</insert>
```



