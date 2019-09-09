# sshMultiTailLog

Simple program for display logs from multi servers in one place.

## Config 

1. Create `config.json` file, e.g.

        {
          "hosts": [
            {
              "address": "127.0.0.1",
              "user": "root",
              "password": "secret",
              "logfile": "/var/log/php7.2-fpm.log"
            },
            {
              "address": "127.0.0.2",
              "user": "root",
              "password": "secret",
              "logfile": "/var/log/php7.2-fpm.log"
            },
            {
              "address": "127.0.0.3",
              "user": "root",
              "password": "secret",
              "logfile": "/var/log/php7.2-fpm.log"
            }
          ]
        }
        
2. (Build and) run program `./sshMultiTailLog`
