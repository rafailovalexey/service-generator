# Dependency

```sudo apt-get install golang```

```sudo apt-get install make```

```sudo apt-get install protoc-gen-go```

```sudo apt-get install protoc-gen-go-grpc```

# GRPC

```service-generator generate -path service-employees -application grpc -module github.com/rafailovalexey/employees -name employees```

```cd service-users```

```sudo chmod +x bin/grpc-generate.sh```

```sudo chmod +x bin/mock-generate.sh```

```make generate```

```make tidy```

# HTTP

```service-generator generate -path service-employees -application http -module github.com/rafailovalexey/employees -name employees```

```cd service-users```

```sudo chmod +x bin/mock-generate.sh```

```make generate```

```make tidy```
