barge
=====
[![Build Status](https://travis-ci.org/thedodd/barge.svg?branch=master)](https://travis-ci.org/thedodd/barge)
[![codecov.io](https://codecov.io/github/thedodd/barge/coverage.svg?branch=master)](https://codecov.io/github/thedodd/barge?branch=master)
[![codebeat badge](https://codebeat.co/badges/5fb376bb-a22e-45b4-b305-4f60c381ac39)](https://codebeat.co/projects/github-com-thedodd-barge)
[![GoDoc](https://godoc.org/github.com/thedodd/barge?status.svg)](https://godoc.org/github.com/thedodd/barge)

Development and deployment for docker based apps made easy.

This project is currently in an `alpha` state. Inspiration for this project has been drawn form years of experience developing the [ObjectRocket](https://objectrocket.com) devtools system, as well as from [Otto](https://www.ottoproject.io/). The original goal was to implement Otto drivers to support the docker ecosystem as the primary means of development and deployment, but I am unsure as to wheather that fits into the Otto vision.

### install
```bash
go get github.com/thedodd/barge
```
This will install the `barge` command-line tool.

### usage
```
barge [--version] [--help] <command> [<args>]

Commands:
    init         Initialize a Bargefile in the working directory.
    dev up       Spin up a new docker machine according to the specs found in
                 the working directory's Bargefile.
    dev destroy  Destroy the docker machine specified in the working directory's
                 Bargefile.
    dev rebuild  Rebuild the docker machine specified in the working directory's
                 Bargefile.
```
`--help` followed by a sub-command will display the help text for that sub-command.

### bargefile
The `Bargefile` is a configuration file written in [HCL](https://github.com/hashicorp/hcl#syntax), with the following structure:

```yaml
development:
    cpus: int — the number of CPUs to allocate to the docker machine.
    disk: int — the disk size (MiB) to allocate to the docker machine.
    machineName: string — the name to assign to the docker machine.
    driver: string — the driver to use for the creation of the docker machine.
    ram: int — the amount of RAM (MiB) to allocate to the docker machine.
```

Currently, the goal is to use only explicit configuration in the Barge ecosystem. Generally speaking, favoring explicit over implicit patterns yields better predictability in software systems, IMHO.
