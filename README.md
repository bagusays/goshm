## About The Project
A wrapper/tool for generating daily stock price historical (IDX) into CSV file, built with golang

<!-- GETTING STARTED -->
## Getting Started
### Prerequisites
* Go 1.15

### Installation
```
$ git clone https://github.com/bagusays/goshm.git
$ make build
$ chmod +x goshm
```

## Usage
```
$ ./goshm fetch --code=ASII --date_from="20-12-2020" --date_to="25-12-2020"
```

or for multiple code
```
$ ./goshm fetch --code=ASII,TLKM --date_from="20-12-2020" --date_to="25-12-2020"
```

## Roadmap Features
- [ ] JSON Exporter


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request