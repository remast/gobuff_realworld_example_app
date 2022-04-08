# Login to Postgres Container

```shell script
$ sudo docker exec -it rw_db bash
$ psql -U postgres
# \c gobuff_realworld_example_app_development
# select * from users;
```

Or:
1. Ubuntu setup: 
   ```shell script
   $ curl -o ~/.psqlrc https://raw.githubusercontent.com/mate-academy/fed/master/mate-scripts/config-files/.psqlrc
   $ sudo apt update
   $ sudo apt install postgresql postgresql-contrib
   
   $ make db
   password: postgres
   # select * from users;
   ```
1. macOS setup:
   ```shell script
   $ curl -o ~/.psqlrc https://raw.githubusercontent.com/mate-academy/fed/master/mate-scripts/config-files/.psqlrc
   $ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   $ brew install postgresql
   
   $ make db
   password: postgres
   # select * from users;
   ```