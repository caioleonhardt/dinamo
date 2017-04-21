# Dinamo
Dynamo CLI to change component values easily in various servers

### Change Value

    dinamo <component> <property> <value> [-e enviroment]
    dinamo orderManager loggingDebug true

### Invoke Method

    dinamo <component> <methodName> -i [-e enviroment]
    dinamo orderRepository invalidateCaches -i

### Aliases
The file aliases.json is used to add all aliases for components. It has two defaults.
```json
{
    "orderManager": "/atg/commerce/order/OrderManager/",
    "orderRepository": "/atg/commerce/order/OrderRepository/"
}
```

### Enviroments and Servers
The structure to create enviroments is the following.  
The property default is used to call *dinamo* command without type the flag enviroment (-e)

```json
{
  "default": "local",
  "enviroments" : 
  [
    {
      "name": "local",
      "user": "admin",
      "password": "admin",
      "servers": [
        {
          "name": "ps01",
          "host": "127.0.0.1",
          "port": "10181"
        }
      ]
    }
  ]
}
```
### Install and Use

- Windows:  
  Clone the repository and run the dynamo.exe  
  Obs: you should have the files servers.json and aliases.json to Dinamo work