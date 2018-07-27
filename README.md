[![badge: Supported by Clockwork](https://img.shields.io/badge/Supported%20by-Clockwork-ffcc00.svg)](https://www.clockwork.com/)

# envtokube

Conversion tool to assist creating Kubernetes Secrets YAML files by providing dotenv.

## How it Works & Usage

`envtokube [--namespace=<name>] <input> [output]`
* `--namespace=` is optional, and will default to the output filename (without file extension)
* An .env file 'input' is required, as the source of secrets to encode for the YAML output file
* If 'output' (a YAML file) is not provided, the input filename will be used, with a `.yml` extension

`envtokube sample.env`
* Will convert a `sample.env` to `sample.yml`
* It will auto-set the metadata.name to be 'sample'
* It will auto-set the metadata.namespace to be 'default'


`envtokube sample.env out.yml`
* Will convert `sample.env` to `out.yml`
* It will auto-set the metadata.name to be 'out'
* It will auto-set the metadata.namespace to be 'default'


`envtokube --namespace=hello-world sample.env out.yml`
* Will convert `sample.env` to `out.yml`
* It will auto-set the metadata.name to be 'out'
* It will set the metadata.namespace to be 'hello-world'


## ToDo

* Reverse, to "decrypt" to an dotenv
* Inherit namespace, if you're overwriting an existing file
* Providing a warning if you're overwriting an existing file

## Additional Resources

Based on: [kube-secrets a helpful utility for dealing with K8s Secretes](https://github.com/jengo/kube-secrets), huge thanks!


* [The Go dotenv port (what I used to help build this)](https://github.com/joho/godotenv)
* [The original Ruby dotenv](https://github.com/bkeepers/dotenv)
* [The PHP dotenv port (why I built this)](https://github.com/vlucas/phpdotenv)

## License

- [LICENSE](LICENSE) (Expat/[MIT License][MIT])

[MIT]: http://www.opensource.org/licenses/MIT "The MIT License (MIT)"