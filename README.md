[![badge: Supported by Clockwork](https://img.shields.io/badge/Supported%20by-Clockwork-ffcc00.svg)](https://www.clockwork.com/)

# envtokube

Conversion tool to assist creating Kubernetes Secrets YAML files by providing dotenv.

## How it Works & Usage

`envtokube sample.env sample.yml`
* Will convert `sample.env` to `sample.yml`


## ToDo

* Input a namespace, instead of always defaulting to "default"
* Reverse, to "decrypt" to an dotenv
* Inherit namespace, if you're overwriting an existing file
* Providing a warning if you're overwriting an existing file

## Additional Resources

* [The original Ruby dotenv](https://github.com/bkeepers/dotenv)
* [The PHP dotenv port (why I built this)](https://github.com/vlucas/phpdotenv)
* [The Go dotenv port (what I used to help build this)](https://github.com/joho/godotenv)
* [kube-secrets a helpful utility for dealing with K8s Secretes](https://github.com/jengo/kube-secrets)
  * Huge thanks, got me 90% of the way there for this tool

## License

- [LICENSE](LICENSE) (Expat/[MIT License][MIT])

[MIT]: http://www.opensource.org/licenses/MIT "The MIT License (MIT)"