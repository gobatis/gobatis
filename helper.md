# 辅助标签

```xml

<insert id="" parameter="row">

    insert into fns.users(*) values(
    <>)

        <bind name="table" value="row"></bind>

        <inserter id="" table="'users'" fields="'fields'" data="user" parameter="user" batch="100">
            on conflict(id,name) do noting;
        </inserter>

        <updater id="" table="'fns.users'" fields="'name,age'" data="user" paramter="id,user">
            where id = ${id}
        </updater>

        <deleter id="" table="'test'" parameter="id,name">
            where id > ${id} and name = ${name}
        </deleter>

        <queryer id="table" parameter="age" result="rows,count">
            <select>
                *
            </select>
            <count>
                count(1)
            </count>
            <from>
                users left join users
                <where>
                    epoch > 10
                </where>
            </from>
            <paging limit="10">
                age > ${age}  
            </paging>
        </queryer>

        <inserter table="'users'" index="index" item="item" data="row" multiple="false">
            <field>name</field>
            <field ignore="false">age</field>
        </inserter>
        returning id;

        <updater table="'users'" entity="row"">
        <field>*</field>
        <field ignore="true">age</field>
        <field ignore="true">name</field>
        <field ignore="true">bod</field>
    </updater>

    <selecter table="'users'">
        <field>*</field>
        <field>name</field>
        <field>age</field>
    </selecter>
    where "id" > 0;

    <query table="'users'" limit="100" count="false" select="" ignore="">
        <field>*</field>
        <field>age</field>
        <field ignore="true">age</field>
    </query>

</insert>
```