# Dinamo
Dynamo CLI to change component values easily in various servers

### Change Value

    dinamo <component> <property> <value> [-e enviroment]
    dinamo orderManager loggingDebug true

### Invoke Method

    dinamo <component> <methodName> -i [-e enviroment]
    dinamo orderRepository invalidateCaches -i

### Aliases
The file aliases.yml is used to add all aliases for components. It has two defaults.
```yml
orderManager: /atg/commerce/order/OrderManager/
orderRepository: /atg/commerce/order/OrderRepository/
```

### Enviroments and Servers
The structure to create enviroments is the following. Change it in servers.json file. 
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
          "name": "production",
          "host": "127.0.0.1",
          "port": "10181,10281,10381"
        },
        {
          "name": "bcc",
          "host": "127.0.0.1",
          "port": "20181"
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
