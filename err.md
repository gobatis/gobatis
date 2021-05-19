**test error mapper: 0**

data:
```
{
 "file": "mapper.xml",
 "content": "<mapper>...</mapper",
 "err": 31
}
```
error:
```
ERROR 31: syntax error
 line 1 column 0:
<mapper>...</mapper
```
**test error mapper: 1**

data:
```
{
 "file": "mapper.xml",
 "content": "<mapper</mapper",
 "err": 31
}
```
error:
```
ERROR 31: 词法分析错误
```
**test error mapper: 2**

data:
```
{
 "file": "mapper.xml",
 "content": "mapper>...</mapper",
 "err": 31
}
```
error:
```
ERROR 31: syntax error
 line 1 column 0:

```
**test error mapper: 3**

data:
```
{
 "file": "mapper.xml",
 "content": "mapper>.../mapper>",
 "err": 31
}
```
error:
```
ERROR 31: syntax error
 line 1 column 0:

```
