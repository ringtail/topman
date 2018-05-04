# topman
```
                  ___                     ___          ___          ___
      ___        /  /\        ___        /  /\        /  /\        /  /\
     /__/\      /  /::\      /  /\      /  /::|      /  /::\      /  /::|
     \  \:\    /  /:/\:\    /  /::\    /  /:|:|     /  /:/\:\    /  /:|:|
      \__\:\  /  /:/  \:\  /  /:/\:\  /  /:/|:|__  /  /::\ \:\  /  /:/|:|__
      /  /::\/__/:/ \__\:\/  /::\ \:\/__/:/_|::::\/__/:/\:\_\:\/__/:/ |:| /\
     /  /:/\:\  \:\ /  /:/__/:/\:\_\:\__\/  /~~/:/\__\/  \:\/:/\__\/  |:|/:/
    /  /:/__\/\  \:\  /:/\__\/  \:\/:/     /  /:/      \__\::/     |  |:/:/
   /__/:/      \  \:\/:/      \  \::/     /  /:/       /  /:/      |__|::/
   \__\/        \  \::/        \__\/     /__/:/       /__/:/       /__/:/
                 \__\/                   \__\/        \__\/        \__\/
  ======================================================================
                     Yes Commander. I am on duty!
  ======================================================================
```

topman is an alarm tool using dingding. topman support ping packet loss and tcp port probe.

## Usage
1. create a topman.conf like this.

```
{
    "ping":[
        {
            "name":"www.aliyun.com",
            "host":"www.aliyun.com",
            "threshold": 30
        }
    ],
    "tcp":[
        {
            "name": "Google",
            "host": "127.0.01:80"
        }
    ]
}


```
2. run topman

```
    topman -token https://oapi.dingtalk.com/robot/send?access_token=[your token] -config topman.conf
```


## modules
ping
```
name:       the name of ping item
host:       specific the domain or ip of the site
threshold:  the threshold of the packet loss, if packet loss is more than threshold it will alert the msg
```

tcp
```
name:       the name of ping item
host:       specific the domain or ip of the site
```

## others
On macos ping module is not work very well.