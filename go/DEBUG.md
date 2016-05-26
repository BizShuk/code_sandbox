# DEBUG

### invalid type for composite literal



### why does copy slice need twice cap()
```go
t := make([]byte, len(s), (cap(s)+1)*2)
copy(t, s)
s = t
```

### slice copy by value
how?

### Could not determine kind of name for C.xxxxxx
A: Need feader file , shared object file , ldconfig , pkg-config path(*.pc file in pkg_config_path )
