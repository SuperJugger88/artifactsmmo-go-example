### Building and running your application in Docker

First of all, you have to set and fill up environment variables:

```bash
cp .env.example .env
```

When you're ready, start your application by running:
`make up`.

### Deploying your application to the cloud

First, build your image, e.g.: `docker build --build-args=GO_VERSION=<your_version> -t myapp .`.
If your cloud uses a different CPU architecture than your development
machine (e.g., you are on a Mac M1 and your cloud provider is amd64),
you'll want to build the image for that platform, e.g.:
`docker build --build-args=GO_VERSION=<your_version> --platform=linux/amd64 -t myapp .`.

Then, push it to your registry, e.g. `docker push myregistry.com/myapp`.

Consult Docker's [getting started](https://docs.docker.com/go/get-started-sharing/)
docs for more detail on building and pushing.

### References
* [Docker's Go guide](https://docs.docker.com/language/golang/)
