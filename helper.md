# 辅助标签

```xml


<insert id="" parameter="row">
    <bind name="table" value="row"></bind>
    <inserter table="'users'" entity="row" multiple="false">
        <field>*</field>
        <field ignore="true">age</field>
    </inserter>

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
    </selecter> where "id" > 0;
    
    <query table="'users'" limit="100" count="false" select="" ignore="">
        <field>*</field>
        <field>age</field>
        <field ignore="true">age</field>
    </query>

</insert>
```