<img width="256px" alt="secretium logo" src="https://raw.githubusercontent.com/secretium/.github/main/images/secretium-logo-with-text.svg">

# Secretium – A smart self-hosted tool for sharing secrets to your friends, colleagues, or anyone!

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

The **Secretium** project is a self-hosted tool for sharing secrets to your friends, colleagues, or anyone. The Open Source (or Community) edition of the project is a fully community-driven solution for personal use.

Features:

- 100% **free** and **Open Source** under the [Apache 2.0][repo_license_url] license;
- For **any** level of developer's knowledge and technical expertise;
- [**Well-documented**][docs_url], with a lot of tips and assists from the authors;
- Powered by the **Go** programming language, **Templ** & **htmx** libraries and **Tailwind** utility-first CSS framework;
- Works with **AES** encryption algorithm for secure your data before storing it in the database;
- **Does not depend** on the host GNU/Linux system, it runs completely in an isolated Docker container;
- Supported automatic switching between the **light/dark** UI themes;
- ...

## ⚡️ Quick start

Here's minimal steps to run the **Secretium** on your **local** machine.

First of all, install [Docker][docker_install_url] with the [Compose][docker_compose_install_url] plugin.

For the security reasons, create the TXT files for the sensitive data to use them with the [Docker Secrets][docker_secrets_url]:

- `secretium_key.txt` for the secret key to encrypt your data in the database;
- `secretium_master_username.txt` for the master username to login to the dashboard as admin;
- `secretium_master_password.txt` for the master password to login to the dashboard as admin.

Open the TXT files and paste your sensitive data:

```bash
echo "this-is-my-secret-key-123" > secretium_key.txt
echo "this-is-my-master-username" > secretium_master_username.txt
echo "this-is-my-master-password-123" > secretium_master_password.txt
```

Run the official [`quick-start.sh`][repo_quick_start_sh_url] script from the **Secretium** repository:

```bash
wget -O - https://raw.githubusercontent.com/secretium/secretium/main/quick-start.sh | bash
```

> [!NOTE]
> This script will automatically create a minimal `docker-compose.yml` file, create a folder for the database, run `docker-compose up -d` command to start the **Secretium** container on port `8787`, and remove the TXT files with the sensitive data after running container.

Open your browser, visit `http://localhost:8787` and login to the admin dashboard with your master username and master password, which are defined in the previous steps.

That's it! 🔥 Your **Secretium** container is up and running!

### 📦 Other ways to quick start

Download a ready-made `deb` (for Debian, Ubuntu) and `rpm` (for CentOS, Fedora) packages from the [Releases][repo_releases_url] page. Install it like you normally would.

## 📖 Complete user guide

We always treasure your time and want you to start building really great web products on this awesome technology stack as soon as possible! Therefore, to get a complete guide to use and understand the basic principles of the **Secretium** project, we have prepared a comprehensive explanation of the project in this 📖 [**Complete user guide**][docs_url].

<a href="https://secretium.org" target="_blank" title="Go to the Secretium's Complete user guide"><img width="360px" alt="secretium docs banner" src="https://raw.githubusercontent.com/secretium/.github/main/images/secretium-docs-banner.svg"></a>

It is highly recommended to start exploring with short introductory articles "[**What is Secretium?**][docs_getting_started_url]" and "[**How does it work?**][docs_how_it_works_url]" to understand the basic principle and the main components built into the **Secretium** project.

Next steps are:

1. [Prepare your remote server](https://secretium.org/complete-user-guide/prepare-remote-server)
2. [Create files with sensitive data](https://secretium.org/complete-user-guide/create-files-with-sensitive-data)
3. [Run installation script](https://secretium.org/complete-user-guide/run-installation-script)
4. [Edit Docker Compose file](https://secretium.org/complete-user-guide/edit-docker-compose)
5. [Run Secretium container](https://secretium.org/complete-user-guide/run-container)

Hope you find answers to all of your questions! 😉

## 🎯 Motivation to create

Very often in the workflow there is a need to send some important information to friends or colleagues (access data, important messages and so on). Usually such data is sent in an open and unprotected form, in private messages in messenger or in a work chat room.

The data hangs around in these channels forever until someone deletes it, making it questionable whether it's safe to store. Also, you have no control over this data, you can't tell for sure if your data will be shared with third parties.

The **Secretium** project solves this problem. It is installed on your server as an isolated Docker container and only you own your secrets!

To transfer your data, all you have to do is log in to the beautiful administrative dashboard, create a secret, share the link and tell your friend the access code so that he/she can unlock it and see the data.

Each secret created by Secretium has a time limit for unlocking — you no longer have to worry that some of your data can be accessed at any time. If you want the secret to become inaccessible immediately after the first unlocking, you have that option too!

We developed this product originally for ourselves and have been using it for a long time within our development team and for communicating with external customers. That's why we took special care with the security of data storage and the look and feel of the UI and usability.

We can't wait for you to start using Secretium! ✨

> [!NOTE]
> Earlier, we have already saved the world twice, they were [Create Go App][cgapp_url] and [Gowebly][gowebly_url] (yep, that's our projects too). The GitHub stars statistics of these projects can't lie: more than **2.3k** developers of any level and different countries start a new project through these CLI tools.

## 🏆 A win-win cooperation

If you liked the **Secretium** project and found it useful for your tasks, please click a 👁️ **Watch** button to avoid missing notifications about new versions, and give it a ⭐️ **GitHub Star**!

It really **motivates** us to make this product **even** better.

...

And now, I invite you to participate in this project! Let's work **together** to create and popularize the **most useful** tool for developers on the web today.

- [Issues][repo_issues_url]: ask questions and submit your features.
- [Pull requests][repo_pull_request_url]: send your improvements to the current.
- [Discussions][repo_discussions_url]: discuss and share your ideas.
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
[repo_quick_start_sh_url]: https://github.com/secretium/secretium/blob/main/quick-start.sh
[repo_install_sh_url]: https://github.com/secretium/secretium/main/install.sh
[repo_issues_url]: https://github.com/secretium/secretium/issues
[repo_pull_request_url]: https://github.com/secretium/secretium/pulls
[repo_discussions_url]: https://github.com/secretium/secretium/discussions
[repo_releases_url]: https://github.com/secretium/secretium/releases
[repo_license_url]: https://github.com/secretium/secretium/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[repo_cc_license_url]: https://creativecommons.org/licenses/by-sa/4.0/

<!-- Docs links -->

[docs_url]: https://secretium.org
[docs_getting_started_url]: https://secretium.org/getting-started
[docs_how_it_works_url]: https://secretium.org/how-it-works

<!-- Docker links -->

[docker_install_url]: https://docs.docker.com/engine/install/#server
[docker_compose_install_url]: https://docs.docker.com/compose/install/linux/
[docker_secrets_url]: https://docs.docker.com/engine/swarm/secrets/
[docker_image_url]: https://hub.docker.com/repository/docker/secretium/secretium

<!-- Author links -->

[author_url]: https://github.com/koddr
[truewebartisans_url]: https://github.com/truewebartisans

<!-- Readme links -->

[cgapp_url]: https://github.com/create-go-app/cli
[gowebly_url]: https://github.com/gowebly/gowebly
[nginx_proxy_manager_url]: https://nginxproxymanager.org
[traefik_proxy_url]: https://traefik.io
[wsl_url]: https://learn.microsoft.com/en-us/windows/wsl/
