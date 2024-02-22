# Dependency

```sudo apt-get install golang```

```sudo apt-get install make```

```sudo apt-get install protoc-gen-go```

```sudo apt-get install protoc-gen-go-grpc```

```go install github.com/pressly/goose/v3/cmd/goose@latest```

# GRPC

```service-generator generate -path employees -application grpc -module github.com/rafailovalexey/employees -name employees```

```cd service-users```

```make generate```

```make tidy```

```make download```

# HTTP

```service-generator generate -path employees -application http -module github.com/rafailovalexey/employees -name employees```

```cd service-users```

```make generate```

```make tidy```

```make download```

# CRON

```service-generator generate -path employees -application cron -module github.com/rafailovalexey/employees -name employees```

```cd service-users```

```make generate```

```make tidy```

```make download```