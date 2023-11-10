# Secretium – A smart self-hosted tool for sharing secrets to your friends

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

A full description for the **Secretium** ...

Features:

- 100% **free** and **Open Source** under the [Apache 2.0][repo_license_url] license;
- Powered by the **Go** programming language, **Templ** & **htmx** libraries and **Tailwind CSS** utility-first framework;
- Works with **AES** encryption algorithm for secure your secrets before storing it in the database;
- **Does not depend** on the host OS, it runs completely in an isolated Docker container;
- Supported automatic switching between the **light/dark** UI themes;
- **Well-documented**, with a lot of tips and assists from the authors;
- For **any** level of developer's knowledge and technical expertise;
- ...

## ⚡️ Quick start

Here's a (*extremely*) minimal version of the steps to run the **Secretium** on your remote server:

- Login to your remote server.
- Install [Docker][docker_install_url] with the [Compose][docker_compose_install_url] plugin.
- For the security reasons, create the TXT files for the sensitive data to use them with the [Docker Secrets][docker_secrets_url]:
  - `secretium_key.txt` for the secret key to encrypt data;
  - `secretium_master_username.txt` for the master username to login to the dashboard as admin;
  - `secretium_master_password.txt` for the master password to login to the dashboard as admin;
  - `secretium_domain.txt` for the domain name to share links.
- Open the TXT files and paste your sensitive data:

```bash
echo "this-is-my-secret-key-123" > secretium_key.txt
echo "this-is-my-master-username" > secretium_master_username.txt
echo "this-is-my-master-password-123" > secretium_master_password.txt
echo "example.com" > secretium_domain.txt
```

- Run the official [`quick-start.sh`][repo_quick_start_sh_url] script from the **Secretium** repository:

```bash
wget -O - https://raw.githubusercontent.com/secretium/secretium/quick-start.sh | bash
```

- This script will automatically create a minimal `docker-compose.yml` file, create a folder for the database, run `docker-compose up -d` command to start the **Secretium** container on port `8787`, and remove the TXT files with the sensitive data after running container.
- Link the container to a web/proxy server (via [Nginx Proxy Manager][nginx_proxy_manager_url], for example).
- Get [Let's Encrypt][lets_encrypt_url] SSL certificate for your domain and add it to the web/proxy server.
- Open your browser, visit `https://example.com` and login to the admin dashboard with your master password.

That's it! 🔥 Your smart self-hosted Secretium is ready to use!

### 📦 Other ways to quick start

Feel free to using the **Secretium** project from the [official Docker image][docker_image_url] and run it manually in the isolated container in your pipelines. This Docker image is available for GNU/Linux only (`amd64` and `arm64`).

Also, the ready-made `exe` files for Windows, `deb`, `rpm`, `apk` or Arch Linux packages can be downloaded from the [Releases][repo_releases_url] page of this repository.

## 📖 Complete user guide

To get a complete guide to use and understand the basic principles of the **Secretium** project, we have prepared a comprehensive explanation of each command at once in this README file.

> [!NOTE]
> We always treasure your time and want you to start using the really great products as soon as possible!

We hope you find answers to all of your questions! 👌 But, if you do not find needed information, feel free to create an [issue][repo_issues_url] or send a [PR][repo_pull_request_url] to this repository.

### Install on the remote server

Here's a complete version of the steps to run the **Secretium** on your remote server:

- Login to your remote server.
- Install [Docker][docker_install_url] with the [Compose][docker_compose_install_url] plugin.
- Run the official [`install.sh`][repo_install_sh_url] script from the **Secretium** repository:

```bash
wget -O - https://raw.githubusercontent.com/secretium/secretium/install.sh | bash
```

- The installation script will create a folder for the database (`./secretium-data`) and the default `docker-compose.yml` file with the following configuration:

```yaml
version: '3.8'

# Define services.
services:
  # Service for the backend.
  secretium:
    # Configuration for the Docker image for the service.
    image: 'secretium/secretium:latest'
    # Set restart rules for the container.
    restart: unless-stopped
    # Forward the exposed port 8787 on the container to port 8787 on the host machine.
    ports:
      - '8787:8787'
    # Set required environment variables for the backend.
    environment:
      SECRET_KEY: /run/secrets/secretium_key
      MASTER_USERNAME: /run/secrets/secretium_master_username
      MASTER_PASSWORD: /run/secrets/secretium_master_password
      DOMAIN: /run/secrets/secretium_domain
      DOMAIN_SCHEMA: https
      SERVER_PORT: 8787 # same as the exposed container port
      SERVER_TIMEZONE: Europe/Moscow
      SERVER_READ_TIMEOUT: 5
      SERVER_WRITE_TIMEOUT: 10
    # Set volumes for the container with SQLite data and the root SSL certificates.
    volumes:
      - ./secretium-data:/secretium-data
      - /etc/ssl/certs:/etc/ssl/certs:ro

# Define the secrets.
secrets:
  # Key for the Secretium.
  secretium_key:
    # Path to the file with your secret key.
    file: secretium_key.txt
  # Master username used for the Secretium dashboard.
  secretium_master_username:
    # Path to the file with your master username.
    file: secretium_master_username.txt
  # Master password used for the Secretium dashboard.
  secretium_master_password:
    # Path to the file with your master password.
    file: secretium_master_password.txt
  # Domain name for the Secretium's links.
  secretium_domain:
    # Path to the file with your domain name.
    file: secretium_domain.txt
```

- Edit the `docker-compose.yml` file with your parameters (see the [Configuration](#configuration) section).
- For the security reasons, create the TXT files for the sensitive data to use them with the [Docker Secrets][docker_secrets_url]:
  - `secretium_key.txt` for the secret key to encrypt data;
  - `secretium_master_username.txt` for the master username to login to the dashboard as admin;
  - `secretium_master_password.txt` for the master password to login to the dashboard as admin;
  - `secretium_domain.txt` for the domain name to share links.
- Open the TXT files and paste your sensitive data:

```bash
echo "this-is-my-secret-key-123" > secretium_key.txt
echo "this-is-my-master-username" > secretium_master_username.txt
echo "this-is-my-master-password-123" > secretium_master_password.txt
echo "example.com" > secretium_domain.txt
```

- Run `docker-compose up -d` command to start the **Secretium** container.

> [!WARNING]
> Don't forget to **remove** the TXT files with the sensitive data after running container from your server:
>
> ```bash
> rm secretium_key.txt secretium_master_username.txt secretium_master_password.txt secretium_domain.txt
> ```

#### Configuration

A complete list of parameters for the `docker-compose.yml` file:

| Environment variable   | Description                                                           | Default value     |
| ---------------------- | --------------------------------------------------------------------- | ----------------- |
| `SECRET_KEY`           | Set your secret key to the encrypt data on the backend                | `""`              |
| `MASTER_USERNAME`      | Set your master username to login to the backend's dashboard as admin | `""`              |
| `MASTER_PASSWORD`      | Set your master password to login to the backend's dashboard as admin | `""`              |
| `DOMAIN`               | Set your domain name for sharing links to the secrets                 | `""`              |
| `DOMAIN_SCHEMA`        | Set the HTTP scheme for your domain name                              | `"https"`         |
| `SERVER_PORT`          | Set the server port for the backend                                   | `8787`            |
| `SERVER_TIMEZONE`      | Set the [Timezone][timezone_url] to the backend                       | `"Europe/Moscow"` |
| `SERVER_READ_TIMEOUT`  | Set the server read timeout to the backend                            | `5`               |
| `SERVER_WRITE_TIMEOUT` | Set the server write timeout to the backend                           | `10`              |

## 🎯 Motivation to create

...

> [!NOTE]
> Earlier, we have already saved the world twice, it was [Create Go App][cgapp_url] and [gowebly][gowebly_url] (yep, that's our projects too). The GitHub stars statistics of these projects can't lie: more than **2.3k** developers of any level and different countries start a new project through these CLI tools.

## 🏆 A win-win cooperation

If you liked the **Secretium** project and found it useful for your tasks, please click a 👁️ **Watch** button to avoid missing notifications about new versions, and give it a ⭐️ **GitHub Star**!

It really **motivates** us to make this product **even** better.

...

And now, I invite you to participate in this project! Let's work **together** to create and popularize the **most useful** tool for developers on the web today.

- [Issues][repo_issues_url]: ask questions and submit your features.
- [Pull requests][repo_pull_request_url]: send your improvements to the current.
- Say a few words about the project on your social networks and blogs (Dev.to, Medium, Хабр, and so on).

Your PRs, issues & any words are welcome! Thank you 😘

### 🌟 Stargazers

...

## ⚠️ License

[`Secretium`][repo_url] is free and open-source software licensed under the [Apache 2.0 License][repo_license_url], created and supported by [Vic Shóstak][author_url] and the [True web artisans][truewebartisans_url] team with 🩵 for people and robots. Official logo distributed under the [Creative Commons License][repo_cc_license_url] (CC BY-SA 4.0 International).

<!-- Go links -->

[go_report_url]: https://goreportcard.com/report/github.com/secretium/secretium
[go_dev_url]: https://pkg.go.dev/github.com/secretium/secretium
[go_version_img]: https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go
[go_code_coverage_url]: https://codecov.io/gh/koddr/secretium
[go_code_coverage_img]: https://img.shields.io/codecov/c/gh/koddr/secretium.svg?logo=codecov&style=for-the-badge
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none

<!-- Repository links -->

[repo_url]: https://github.com/secretium/secretium
[repo_complete_user_guide_url]: https://github.com/secretium/secretium#-complete-user-guide
[repo_quick_start_sh_url]: https://github.com/secretium/secretium/blob/main/quick-start.sh
[repo_install_sh_url]: https://github.com/secretium/secretium/main/install.sh
[repo_issues_url]: https://github.com/secretium/secretium/issues
[repo_pull_request_url]: https://github.com/secretium/secretium/pulls
[repo_releases_url]: https://github.com/secretium/secretium/releases
[repo_license_url]: https://github.com/secretium/secretium/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[repo_cc_license_url]: https://creativecommons.org/licenses/by-sa/4.0/

<!-- Docker links -->

[docker_install_url]: https://docs.docker.com/engine/install/#server
[docker_compose_install_url]: https://docs.docker.com/compose/install/linux/
[docker_secrets_url]: https://docs.docker.com/engine/swarm/secrets/
[docker_image_url]: https://hub.docker.com/repository/docker/secretium/secretium

<!-- Author links -->

[author_url]: https://github.com/koddr
[truewebartisans_url]: https://github.com/truewebartisans

<!-- Readme links -->

[timezone_url]: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
[cgapp_url]: https://github.com/create-go-app/cli
[gowebly_url]: https://github.com/gowebly/gowebly
[nginx_proxy_manager_url]: https://nginxproxymanager.org
[lets_encrypt_url]: https://letsencrypt.org
