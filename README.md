# Dependency

```sudo apt-get install golang```

```sudo apt-get install make```

```sudo apt-get install protoc-gen-go```

```sudo apt-get install protoc-gen-go-grpc```

# GRPC

```service-generator generate -path service-users -implement grpc -module github.com/rafailovalexey/service-users -name users```

```cd service-users```

```make tidy```

```make generate```

# HTTP

```service-generator generate -path service-users -implement http -module github.com/rafailovalexey/service-users -name users```

```cd service-users```

```make tidy```

```make generate```

# Subscribe

```service-generator generate -path service-users -implement subscribe -module github.com/rafailovalexey/service-users -name users```

```cd service-users```

```make tidy```

```make generate```

# CRON

```service-generator generate -path service-users -implement cron -module github.com/rafailovalexey/service-users -name users```

```cd service-users```

```make tidy```

```make generate```
