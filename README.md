## Description

Go application with telegram bot. Consider it like test example.

## Run project

#### a) Standard  Go application
1. Create .development.env or .production.env file by example or use python script
```shell
python git_clone_init.py
```
2. Run flyway migrations with PostgreSQL
```shell
cd scripts
sh ./migrations.sh
```
3. Start project
```shell
cd build
make run_develop
```
or for production version
```shell
make run_production
```

#### b) Docker container
<em>soon</em>

## License

`Go telegram bot` is [MIT licensed](LICENSE).
