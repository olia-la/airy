- [Introduction](#introduction)
- [Quick start](#quick-start)
- [How to install](#how-to-install)
- [Running the CLI](#running-the-cli)
  - [Running the app reset](#running-the-app-reset)
- [Setup auto-completion](#setup-auto-completion)

# Introduction

The Airy CLI is a program that automates various aspects of our workflow.

# Quick start

Once you [installed](#how-to-install) the Airy CLI, you can run:

```shell
$ airy -h
```

to get an overview of the available commands.

# How to install

Please run the `airycli_install.sh` script from the `scripts` directory. The script by default installs the binary in `$HOME/bin`. The dir has to be in your PATH for the app to work.

Example:
```
cd scripts
./airycli-install.sh
```
In case you don't want to install the binary in that directory, you can set your installation location with the env var as:
```
INSTALLATION_DIR=/some/path ./airycli-install.sh
```

Last step is to add your GitHub personal access token in the `~/.airy-prod.yaml` configuration file,  which was created with the `airycli_install.sh` command. 

The token needs to have the `read:org` scope.
You can generate your personal github token through `Settings -> Developer settings -> Personal access tokens` or https://github.com/settings/tokens

Putting your personal gihub token in the configuration file, can be done by editing the `~/.airy-prod.yaml` file and adding the token to the line:

```
githubtoken: 
```


# Running the cli

The help can be seen putting the `-h` or `--help` flags.

Example:
```
airy --help
airy terraform -h
```

For every command / subcommand, the help can be seen also with putting `-h` or
`--help` after the command/subcommand. However, when viewing the flags for a
particular subcommand, the flags which are local for the parent subcommand are
not displayed. Which means that after running

```
airy terraform plan -h
```

the flags which are local for `terraform` will not be displayed. They can be seen when running

```
airy terraform -h
```

## Running the app reset

To run the `app reset` option you must specify the app-name `-a` and the app-id `-i`.

```sh
airy app reset -a message-read-states-migration -i conversations.streams.MessageReadStatesMigration
```

You can change the behavior with the flags. You can get the full help with the `help` subcommand:

```sh
airy help app reset
```
or

```sh
airy app reset --help
```


# Setup auto-completion

TODO
