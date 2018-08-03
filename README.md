# redis-horae
> unfinished


Horae is a unofficial plugin for redis (standalone & cluster)

In addition, horae provides utilities to :

 1.  Generate a Memory Report of your data across all databases and keys
 2.  Find big keys in redis , just like `redis-cli --bigkeys` , more details in `commands.bigKeysCommand`
 3.  Delete keys at one's pleasure , more details in `commands.delKeysCommand`
 4.  Check config , find bugs from your redis config (redis-cli config get `config`)
 5.  Monitor your redis, generate a usage report of your redis
 6.  Provide a dashboard page served by its embedded web server
 7.  Convert dump files to JSON

