# 辅助标签

## inserter

```xml

<insert id="" parameter="row,name">
    <!--    <bind name="table" value="row"></bind>-->
    <!-- 当出现 * 时，要求 data 解析结果必须是 struct  -->
    <inserter table="'users'" data="row" item="row" index="index">
        <field>*</field>
        <field name="'name'">age</field>
    </inserter>

    <inserter table="'users'" index="index" item="item" data="row" multiple="false">
        <field name="'name'">name</field>
    </inserter>
    returning id;
</insert>
```

## query

```go
type Mapper struct{
QueryUsers func (page, limit int)(count int64, rows []*User, err error)
}
```

```xml

<query id="QueryUsers" parameter="page,limit">
    <count>
        select count(1)
    </count>
    <collect>
        select name,age
    </collect>
    <source>
        from users where age > 10
    </source>
    <!--  order by "age" desc limit #{limit}, #{ (page-1) * limit }  -->
    <pager page="page" limit="limit" order='""age" desc"'/>
</query>
```

## save

> PGSQL 可以通过 on conflict 实现，但是不能做到多条件判断

只支持参数映射返回，不支持 result 返回

```xml
<save id="SaveUser" parameter="user" result="id">
    <!--  该用法会引起严格模式判断  -->
    <bind name="table" values="'users'"/>

    <!-- 只接收一个 int* 返回参数 -->
    <!--    <exist result="count">-->
    <!--        select count(1) as count from ${table}-->
    <!--        where id = #{user.Id}-->
    <!--        or username=#{user.Username}-->
    <!--        or email = #{user.Email}-->
    <!--        or mobile = #{user.Mobile}-->
    <!--    </exist>-->
    <!--    <insert>-->
    <!--        <inserter table="table" data="user">-->
    <!--            <filed name="*"/>-->
    <!--        </inserter>-->
    <!--        returning id;-->
    <!--    </insert>-->
    <!--    <conflict handler="user.handler"/>-->

    <if test="user.Id > 0">
        <inserter table="table" data="user">
            <filed name="*"/>
        </inserter>
        returning id;
    </if>

    <if test="user.Id == 0">
        update $table set
        "name" = #{user.Name},
        age = #{user.Age}
        where id = #{user.Id}
    </if>
</save>
```