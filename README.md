# go-whosonfirst-static

![](docs/images/wof-static-sf.png)

## Important

This is probably still too soon for you. If nothing else, there isn't really any documentation yet.

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.7 so let's just assume you need [Go 1.9](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Example

```
bin/wof-staticd -port 8080 -source fs -fs-root /usr/local/data/whosonfirst-data/data -cache lru -cache-arg 'CacheSize=500' -debug -mapzen-apikey ${MAPZEN_APIKEY}
```

Or with a fixed-sized cache:

```
bin/wof-staticd -port 8080 -source fs -fs-root /usr/local/data/whosonfirst-data/data -cache bigcache -cache-arg HardMaxCacheSize=100 -cache-arg MaxEntrySize=1024 -debug -mapzen-apikey ${MAPZEN_APIKEY}
2017/12/29 18:37:10 listening on localhost:8080
2017/12/29 18:37:54 REQUEST /id/85633793/
2017/12/29 18:37:54 GET 856/337/93/85633793.geojson CACHE MISS
2017/12/29 18:37:54 MISS 856/337/93/85633793.geojson
2017/12/29 18:37:54 READ 856/337/93/85633793.geojson <nil>
2017/12/29 18:37:56 SET 856/337/93/85633793.geojson entry is bigger than max shard size
```

_`HardMaxCacheSize` is measured in MB and `MaxEntrySize` in bytes._

## Docker

[Yes](Dockerfile). For example:

First, do the usual Docker `build` dance:

```
docker build -t wof-static .
```

Then:

```
docker run -it -p 6161:8080 -e HOST='0.0.0.0' -e SOURCE='http' -e HTTP_ROOT='https://whosonfirst.mapzen.com/data/' -e MAPZEN_APIKEY='mapzen-****' wof-staticd
```

Or:

```
docker run -it -p 6161:8080 -e HOST='0.0.0.0' -e SOURCE='s3' -e S3_BUCKET='whosonfirst' -e S3_PREFIX='' -e S3_REGION='us-east-1' -e S3_CREDENTIALS='env:' -e AWS_ACCESS_KEY_ID='***' -e AWS_SECRET_ACCESS_KEY='***' -e MAPZEN_APIKEY='mapzen-***' wof-staticd
```

Or even still, with caching:

```
docker run -it -p 6161:8080 -e HOST='0.0.0.0' -e SOURCE='http' -e HTTP_ROOT='https://whosonfirst.mapzen.com/data/' -e CACHE='gocache' -e CACHE_ARGS='DefaultExpiration=300 CleanupInterval=600' -e DEBUG='debug' -e MAPZEN_APIKEY='mapzen-*****' wof-static
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-readwrite
