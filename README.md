# taogo

```
    01__
    _/  |______    ____   ____   ____
    \   __\__  \  /  _ \ / ___\ /  _ \
     |  |  / __ \(  <_> ) /_/  >  <_> )
     |__| (____  /\____/\___  / \____/
               \/      /_____/
```

A util to generate the universe of tao!

```shell
go install github.com/taouniverse/taogo@v0.0.2
taogo --help
taogo version
```

## unit

generate tao unit

```shell
taogo help unit
taogo unit -m github.com/taouniverse/tao-mysql -n tao-mysql -r gorm.io/gorm,gorm.io/driver/mysql
```

## project

generate tao project

```shell
taogo help project
taogo project -m github.com/taouniverse/home -n home -r github.com/taouniverse/tao-gin
```

