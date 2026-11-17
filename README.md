# demo-monorepo

A small monorepo with two services, used to demo [badil](https://github.com/badilhq/badil)'s
path filters: a PR that changes only `api/` pins only the api service in its preview, and
one that changes only `web/` pins only web.

## Layout

```
api/                 Go HTTP service (GET /health, GET /hello), no database
  Dockerfile
  chart/             wrapper chart on badil's app chart (alias app)
web/                 static site served by nginx-unprivileged on port 8080
  Dockerfile
  chart/             wrapper chart on badil's app chart (alias app)
.github/workflows/   CI: tests api, builds and pushes both images on every run
```

## Images

CI builds both images for `linux/amd64` and `linux/arm64` on every push to `main` and every
pull request, so every commit has both:

- `ghcr.io/badilhq/demo-monorepo-api`
- `ghcr.io/badilhq/demo-monorepo-web`

Each is tagged with the full commit SHA (the PR head SHA on pull requests) and `sha-<short>`.

## badil.yaml catalog entries

A gitops repo declares the two services under `services` of its `badil.yaml`. Both charts
alias the `app` chart, so the default `imageTagParam` (`app.image.tag`) and `hostParam`
(`app.ingress.host`) apply and only `chart.path` is set:

```yaml
services:
  mono-api:
    repo: https://github.com/badilhq/demo-monorepo.git
    chart: {path: api/chart}
    paths: ["api/**"]
  mono-web:
    repo: https://github.com/badilhq/demo-monorepo.git
    chart: {path: web/chart}
    paths: ["web/**"]
```

To give them preview URLs, add them to `preview.expose`:

```yaml
preview:
  expose:
    mono-api: {}
    mono-web: {}
```

## Local checks

```sh
(cd api && go test ./...)
docker build -t demo-monorepo-api api
docker build -t demo-monorepo-web web
for c in api/chart web/chart; do helm dependency build $c && helm lint $c && helm template mono $c; done
```

Previews of this repo are built by badil from its pull requests.
