---
title: The Airy CLI
sidebar_label: Airy CLI
---

This document provides an overview of the commands of the Airy CLI and the different configuration parameters.

## Help

Lists all available commands.

The help can be seen putting the `-h` or `--help` flags.

## Version

Returns the version number from the VERSION file.

## Core

### Start

Launches the Airy Core Platform locally by running the bootstrap script.

```
--beta will load the app images with the beta tag
```

### Login

Creates a default user and returns a login token.

### Demo

Opens the chatplugin demo page in the browser if the platform is running locally.

### Stop

Suspends the Vagrant machine.

### Destroy

Destroys the Vagrant machine.

