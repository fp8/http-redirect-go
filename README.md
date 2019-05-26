# http-redirect

A very simple program to redirect to URL defined but also allow for definition of health check
url.  This is designed to be used by a Kubernetes Ingress.  The existing solution like `schmunk42/docker-nginx-redirect`
works great but does not allow for definition of `/healthz` URL.

The environment variable used.

* `SERVER_PORT`: Server and port `http-redirect` should listen to.  Defaults to `:8080`.
* `SERVER_REDIRECT`: Url to redirect to, defaults to `https://www.google.com/`
* `SERVER_REDIRECT_CODE`: Http response code for rediect, default to `301` 
* `HEALTH_ENDPOINT`: Health check endpoint, default to `/healthz`

P.S.: The main reason for this project is simply my desire to write a program in `golang`...

## build

One of the main reason to use `golang` is that binary it generate requires only linux kernel.  The multi-staged
build in the Dockerfile actually create docker image from `scratch` which does not contain even simple `sh`.  The
resulting image is 7.33MB.  As the compiled `goapp` is actually 7.1MB, docker only addes 0.23 MB of overhead.

## References

* https://weberc2.bitbucket.io/posts/golang-docker-scratch-app.html
* https://github.com/schmunk42/docker-nginx-redirect
