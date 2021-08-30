**test error mapper: 0**

data:
```
{
 "file": "mapper.xml",
 "content": "<mapper>...</mapper",
 "err": 33
}
```
error:
```
[ERROR 33]: syntax error
[file]: mapper.xml near line 1 column 1:
[context]: <mapper>...</mapper
```
**test error mapper: 1**

data:
```
{
 "file": "mapper.xml",
 "content": "<mapper</mapper",
 "err": 33
}
```
error:
```
[ERROR 33]: 词法分析错误
```
**test error mapper: 2**

data:
```
{
 "file": "mapper.xml",
 "content": "mapper>...</mapper",
 "err": 33
}
```
error:
```
[ERROR 33]: syntax error
[file]: mapper.xml near line 1 column 1:
[context]: 
```
**test error mapper: 3**

data:
```
{
 "file": "mapper.xml",
 "content": "mapper>.../mapper>",
 "err": 33
}
```
error:
```
[ERROR 33]: syntax error
[file]: mapper.xml near line 1 column 1:
[context]: 
```
