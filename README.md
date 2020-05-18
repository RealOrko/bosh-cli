# Bosh CLI

<div>
    <a href="https://snapcraft.io/cf-bs">
        <img alt="Get it from the Snap Store" src="https://snapcraft.io/static/images/badges/en/snap-store-white.svg" />
    </a>
</div>
<div>
    <a href="https://snapcraft.io/cf-bs">
        <img alt="cf-bs" src="https://snapcraft.io/cf-bs/badge.svg" />
    </a>
</div>

* Documentation: [bosh.io/docs/cli-v2](https://bosh.io/docs/cli-v2.html)
* Slack: #bosh on <https://slack.cloudfoundry.org>
* Mailing list: [cf-bosh](https://lists.cloudfoundry.org/pipermail/cf-bosh)
* CI: <https://main.bosh-ci.cf-app.com/teams/main/pipelines/bosh:cli>
* Roadmap: [Pivotal Tracker](https://www.pivotaltracker.com/n/projects/956238)

## Install

**Mac OS X**

```sh
$ brew install cloudfoundry/tap/bosh-cli
```

**Linux**

You have to make sure you have [installed snapd](https://snapcraft.io/docs/installing-snapd) for your linux distro. 

```sh
snap install cf-bs
```

You can also alias `cf-bs` to `bosh` in your `.bashrc` file. 

```sh
echo "alias bosh='cf-bs'" >> ~/.bashrc
```

**More Info**

- [https://bosh.io/docs/cli-v2-install/](https://bosh.io/docs/cli-v2-install/)

## Client Library

This project includes [`director`](director/interfaces.go) and [`uaa`](uaa/interfaces.go) packages meant to be used in your project for programmatic access to the Director API.

See [docs/example.go](docs/example.go) for a live short usage example.

## Developer Notes

- [Workstation setup docs](docs/build.md)
- [Test docs](docs/test.md)
- [CLI workflow](docs/cli_workflow.md)
  - [Architecture docs](docs/architecture.md)
