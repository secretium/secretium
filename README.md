# {{PROJECT}} – a small description

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

**English** | [Русский][repo_readme_ru_url] | [简体中文][repo_readme_cn_url] | 
[Español][repo_readme_es_url]

A full description for the **{{PROJECT}}** ...

Features:

- 100% **free** and **open source** under the [Apache 2.0][repo_license_url] 
  license;
- For **any** level of developer's knowledge and technical expertise;
- **Well-documented**, with a lot of tips and assists from the authors;
- ...

## ⚡️ Quick start

First, [download][go_download_url] and install **Go**. Version `1.21` (or 
higher) is required.

Now, you can use `{{PROJECT}}` without installation. Just [`go run`][go_run_url] 
it to create a new [default][repo_default_config] config file:

```console
go run github.com/koddr/{{PROJECT}}@latest init
```

Edit a config file with your settings and options. 

Next, run a generation process to build Docker files to deploy your project:

```console
go run github.com/koddr/{{PROJECT}}@latest generate
```

And now, deploy your project by the generated `Dockerfile` and 
`docker-compose.yml` files in the current folder to the remote server.

> 👆 Tip: We are strongly recommended to use the awesome 
> [Portainer][portainer_url] platform (or a self-hosted Community Edition) 
> for the deploying process.

That's it! 🔥 A wonderful ...

### 🔹 A full Go-way to quick start

If you still want to install `{{PROJECT}}` CLI to your system by Golang, use the 
[`go install`][go_install_url] command:

```console
go install github.com/koddr/{{PROJECT}}@latest
```

### 🍺 Homebrew-way to quick start

GNU/Linux and Apple macOS users available way to install `{{PROJECT}}` CLI via 
[Homebrew][brew_url].

Tap a new formula:

```console
brew tap koddr/tap
```

Install to your system:

```console
brew install koddr/tap/{{PROJECT}}
```

### 🐳 Docker-way to quick start

Feel free to using `{{PROJECT}}` CLI from our 
[official Docker image][docker_image_url] and run it in the isolated container:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} koddr/{{PROJECT}}:latest [COMMAND]
```

### 📦 Other way to quick start

Download ready-made `exe` files for Windows, `deb`, `rpm`, `apk` or Arch 
Linux packages from the [Releases][repo_releases_url] page.

## 📖 Complete user guide

To get a complete guide to use and understand the basic principles of the
`{{PROJECT}}` CLI, we have prepared a comprehensive explanation of each 
command at once in this README file.

> 💬 From the authors: We always treasure your time and want you to start 
> building really great web products on this awesome technology stack as 
> soon as possible!

We hope you find answers to all of your questions! 👌 But, if you do not find 
needed information, feel free to create an [issue][repo_issues_url] or send a 
[PR][repo_pull_request_url] to this repository.

Don't forget to switch this page for your language (current is
**English**): [Русский][repo_readme_ru_url], [简体中文][repo_readme_cn_url],
[Español][repo_readme_es_url].

### `init`

Command to create a **default** config file 
([`.{{PROJECT}}.yml`][repo_default_config]) in the current folder.

```console
{{PROJECT}} init
```

Typically, a created config file contains the following options:

```yaml
...
```

...

| Option name | Description |
|-------------|-------------|
| `default`   | Use the ... |

## 🎯 Motivation to create

...

> 💬 From the authors: Earlier, we have already saved the world once, it was 
> [Create Go App][cgapp_url] (yep, that's our project too). The 
> [GitHub stars][cgapp_stars_url] statistics of this project can't lie: 
> more than **2.2k** developers of any level and different countries start a 
> new project through this CLI tool.

## 🏆 A win-win cooperation

If you liked the `{{PROJECT}}` project and found it useful for your tasks, 
please click a 👁️ **Watch** button to avoid missing notifications about new 
versions, and give it a 🌟 **GitHub Star**!

It really **motivates** us to make this product **even** better.

...

And now, I invite you to participate in this project! Let's work **together** to
create and popularize the **most useful** tool for developers on the web today.

- [Issues][repo_issues_url]: ask questions and submit your features.
- [Pull requests][repo_pull_request_url]: send your improvements to the current.
- Say a few words about the project on your social networks and blogs
  (Dev.to, Medium, Хабр, and so on).

Your PRs, issues & any words are welcome! Thank you 😘

### 🌟 Stargazers

...

## ⚠️ License

[`{{PROJECT}}`][repo_url] is free and open-source software licensed 
under the [Apache 2.0 License][repo_license_url], created and supported by 
[Vic Shóstak][author_url] with 🩵 for people and robots. Official logo 
distributed under the [Creative Commons License][repo_cc_license_url] (CC BY-SA 
4.0 International).

<!-- Go links -->

[go_download_url]: https://golang.org/dl/
[go_run_url]: https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program
[go_install_url]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[go_report_url]: https://goreportcard.com/report/github.com/koddr/{{PROJECT}}
[go_dev_url]: https://pkg.go.dev/github.com/koddr/{{PROJECT}}
[go_version_img]: https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go
[go_code_coverage_url]: https://codecov.io/gh/koddr/{{PROJECT}}
[go_code_coverage_img]: https://img.shields.io/codecov/c/gh/koddr/{{PROJECT}}.svg?logo=codecov&style=for-the-badge
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none

<!-- Repository links -->

[repo_url]: https://github.com/koddr/{{PROJECT}}
[repo_issues_url]: https://github.com/koddr/{{PROJECT}}/issues
[repo_pull_request_url]: https://github.com/koddr/{{PROJECT}}/pulls
[repo_releases_url]: https://github.com/koddr/{{PROJECT}}/releases
[repo_license_url]: https://github.com/koddr/{{PROJECT}}/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[repo_cc_license_url]: https://creativecommons.org/licenses/by-sa/4.0/
[repo_readme_ru_url]: https://github.com/koddr/{{PROJECT}}/blob/main/README_RU.md
[repo_readme_cn_url]: https://github.com/koddr/{{PROJECT}}/blob/main/README_CN.md
[repo_readme_es_url]: https://github.com/koddr/{{PROJECT}}/blob/main/README_ES.md
[repo_default_config]: https://github.com/koddr/{{PROJECT}}/blob/main/internal/attachments/configs/default.yml

<!-- Author links -->

[author_url]: https://github.com/koddr

<!-- Readme links -->

[cgapp_url]: https://github.com/create-go-app/cli
[cgapp_stars_url]: https://github.com/create-go-app/cli/stargazers
[docker_image_url]: https://hub.docker.com/repository/docker/koddr/{{PROJECT}}
[portainer_url]: https://docs.portainer.io
[brew_url]: https://brew.sh
