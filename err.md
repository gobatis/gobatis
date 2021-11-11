**test parse expression: 4**

data:
```
{
 "in": [
  {}
 ],
 "parameter": "a",
 "expr": "a.b",
 "result": 1,
 "err": 36
}
```
error:
```
expr.xml line 0:1 [1] reflect.Value.Interface: cannot return value obtained from unexported field or method at ''
```
**test parse expression: 28**

data:
```
{
 "in": [
  1,
  1
 ],
 "parameter": "a,b",
 "expr": "a + b",
 "result": 2,
 "err": 36
}
```
error:
```
expr.xml line 1:1 [15] operand types are different and cant't convert at 'a+b'
```
**test parse expression: 29**

data:
```
{
 "in": [
  1,
  1
 ],
 "parameter": "a,b",
 "expr": "a + b",
 "result": 2,
 "err": 36
}
```
error:
```
expr.xml line 1:1 [15] operand types are different and cant't convert at 'a+b'
```
**test parse expression: 84**

data:
```
{
 "in": [
  1
 ],
 "parameter": "a",
 "expr": "++1",
 "result": 1,
 "err": 36
}
```
error:
```
expr.xml line 0:1 [1] interface conversion: *antlr.CommonTokenStream is not antlr.CharStream: missing method GetText at ''
```
**test parse expression: 85**

data:
```
{
 "in": [
  1
 ],
 "parameter": "a",
 "expr": "+++1",
 "result": 1,
 "err": 36
}
```
error:
```
expr.xml line 0:1 [1] interface conversion: *antlr.CommonTokenStream is not antlr.CharStream: missing method GetText at ''
```
**test parse expression: 87**

data:
```
{
 "in": [
  1
 ],
 "parameter": "a",
 "expr": "--1",
 "result": -1,
 "err": 36
}
```
error:
```
expr.xml line 0:1 [1] interface conversion: *antlr.CommonTokenStream is not antlr.CharStream: missing method GetText at ''
```
**test parse expression: 88**

data:
```
{
 "in": [
  1
 ],
 "parameter": "a",
 "expr": "---1",
 "result": -1,
 "err": 36
}
```
error:
```
expr.xml line 0:1 [1] interface conversion: *antlr.CommonTokenStream is not antlr.CharStream: missing method GetText at ''
```
