# entrypoint-demoter

[![Upstream release](https://img.shields.io/github/v/release/itzg/entrypoint-demoter?label=Upstream%20release)](https://github.com/itzg/entrypoint-demoter/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

[Deutsch](#deutsch) · [English](#english)

## Deutsch

`entrypoint-demoter` startet einen Container-Unterprozess mit passenden Benutzer- und Gruppenrechten. Die IDs kommen aus `UID` und `GID` oder aus den Eigentümerrechten eines eingebundenen Verzeichnisses. Das hilft, wenn Container und Host dieselben Dateien verwenden.

Dieses Repository ist ein Fork von [itzg/entrypoint-demoter](https://github.com/itzg/entrypoint-demoter). Das ursprüngliche Projekt stammt von Geoff Bourne. Fertige Programme und Prüfsummen findest du in den [Upstream-Releases](https://github.com/itzg/entrypoint-demoter/releases).

### Funktionen

- Benutzer- und Gruppen-ID über Umgebungsvariablen festlegen.
- Dateieigentümer mit `--match` als Vorgabe übernehmen.
- Den Unterprozess mit seinen Argumenten starten und dessen Rückgabecode weitergeben.
- Bei `SIGTERM` einen Stoppbefehl über die Standardeingabe senden.
- Optional vorher eine Ankündigung senden und eine Wartezeit einhalten.

### Voraussetzungen

Du brauchst eine zur Zielplattform passende Programmdatei. Für den Einsatz im Container müssen Anwendung, eingebundene Verzeichnisse und Dateirechte zueinander passen. Ein bereits ohne Root-Rechte gestarteter Prozess behält seine vorhandenen Benutzer- und Gruppenrechte.

### Schnellstart

Wenn das Programm installiert ist, startet dieser Aufruf `/entry.sh` mit den Eigentümerrechten von `/data`:

```bash
entrypoint-demoter --match /data /entry.sh
```

Die Pfade sind Beispiele. `/data` muss vorhanden sein; `/entry.sh` ist dein eigenes Startskript.

### Konfiguration

| Variable | Bedeutung |
| --- | --- |
| `UID` | Benutzer-ID des Unterprozesses. Hat Vorrang vor der Benutzer-ID aus `--match`. |
| `GID` | Gruppen-ID des Unterprozesses. Hat Vorrang vor der Gruppen-ID aus `--match`. |

Ohne Vorgabe über Umgebungsvariablen oder `--match` wird jeweils die ID `0` verwendet. Der Aufruf allein erzwingt daher keinen Wechsel zu einem Benutzer ohne Root-Rechte.

| Option | Bedeutung |
| --- | --- |
| `--match PATH` | Eigentümer des angegebenen Pfads für Benutzer- und Gruppen-ID verwenden. |
| `--stdin-on-term MESSAGE` | Bei `SIGTERM` die Nachricht mit einem Zeilenumbruch an die Standardeingabe des Unterprozesses senden. |
| `--stdin-on-term-announce MESSAGE` | Vor der Wartezeit eine Ankündigung senden. `%delay%` wird durch die Wartezeit in ganzen Sekunden ersetzt. |
| `--stdin-on-term-delay DURATION` | Vor dem Stoppbefehl warten, zum Beispiel `30s`. Ein zweites `SIGTERM` beendet die verbleibende Wartezeit. |
| `--debug` | Debug-Protokollierung aktivieren. |
| `--version` | Versionsinformationen anzeigen und beenden. |

Ankündigung und Wartezeit werden zusammen mit `--stdin-on-term` verwendet. Die Wartezeit funktioniert auch ohne Ankündigung. Sie muss kürzer als die Stoppfrist des Containers sein, damit der Stoppbefehl vor einem erzwungenen Ende noch ankommt. Die Schreibweise folgt [Go-Dauerangaben](https://pkg.go.dev/time#ParseDuration).

### Einbindung in ein Dockerfile

Lade eine passende Linux-Programmdatei aus einem [Upstream-Release](https://github.com/itzg/entrypoint-demoter/releases) und prüfe sie anhand der dortigen `checksums.txt`. Lege die entpackte Datei als `entrypoint-demoter` neben dein Dockerfile. Dieser Ausschnitt ergänzt sie und dein eigenes Startskript:

```dockerfile
COPY entrypoint-demoter /usr/local/bin/entrypoint-demoter
COPY entry.sh /entry.sh
RUN chmod 0755 /usr/local/bin/entrypoint-demoter /entry.sh

ENTRYPOINT ["/usr/local/bin/entrypoint-demoter", "--match", "/data", "/entry.sh"]
```

Wähle die Programmdatei nach der Architektur des Ziel-Images, etwa AMD64 oder ARM64. Weitere Einbindungsbeispiele findest du im [Dockerfile des Bedrock-Servers](https://github.com/itzg/docker-minecraft-bedrock-server/blob/master/Dockerfile).

### Lokaler Build und Tests

Für einen lokalen Build brauchst du Go in der zu [go.mod](go.mod) passenden Version sowie Make. Führe die Befehle im Repository-Ordner aus:

```bash
make build
make test
```

Die zusätzlichen Containerprüfungen benötigen Bash und Docker:

```bash
bash test.sh
```

[Makefile](Makefile) und [test.sh](test.sh) enthalten die jeweiligen Aufrufe. Die Containerprüfung baut ein Test-Image und prüft unter anderem Benutzerrechte und Stoppnachrichten.

### Lizenz

[MIT License](LICENSE). Copyright © 2019 Geoff Bourne. Die ursprünglichen Lizenz- und Urheberhinweise gelten auch in diesem Fork.

## English

`entrypoint-demoter` starts a container subprocess with suitable user and group permissions. IDs come from `UID` and `GID` or from the ownership of a mounted directory. This helps when the container and host share files.

This repository is a fork of [itzg/entrypoint-demoter](https://github.com/itzg/entrypoint-demoter), originally created by Geoff Bourne. Binaries and checksums are available in the [upstream releases](https://github.com/itzg/entrypoint-demoter/releases).

### Features

- Set the user and group IDs through environment variables.
- Use file ownership as the default with `--match`.
- Start a subprocess with its arguments and propagate its exit code.
- Send a stop command through standard input on `SIGTERM`.
- Optionally send an announcement first and wait before stopping.

### Requirements

Use a binary that matches the target platform. The application, mounted directories and file permissions must match your container setup. A process already running without root privileges keeps its current user and group credentials.

### Quickstart

Once the program is installed, this command starts `/entry.sh` using the ownership of `/data`:

```bash
entrypoint-demoter --match /data /entry.sh
```

These paths are examples. `/data` must exist; `/entry.sh` is your own entrypoint script.

### Configuration

| Variable | Purpose |
| --- | --- |
| `UID` | User ID for the subprocess. Takes precedence over the user ID from `--match`. |
| `GID` | Group ID for the subprocess. Takes precedence over the group ID from `--match`. |

An ID defaults to `0` if neither its environment variable nor `--match` supplies a value. Calling the tool without these settings does not force a switch to a non-root user.

| Option | Purpose |
| --- | --- |
| `--match PATH` | Use the given path's ownership to determine the user and group IDs. |
| `--stdin-on-term MESSAGE` | On `SIGTERM`, write the message and a newline to the subprocess's standard input. |
| `--stdin-on-term-announce MESSAGE` | Send an announcement before the delay. `%delay%` is replaced with the delay in whole seconds. |
| `--stdin-on-term-delay DURATION` | Wait before sending the stop command, for example `30s`. A second `SIGTERM` skips the remaining delay. |
| `--debug` | Enable debug logging. |
| `--version` | Show version information and exit. |

The announcement and delay are used with `--stdin-on-term`. A delay also works without an announcement. Keep it shorter than the container's stop grace period so the stop command arrives before the process is forcibly killed. Values use [Go duration syntax](https://pkg.go.dev/time#ParseDuration).

### Dockerfile integration

Download a suitable Linux binary from an [upstream release](https://github.com/itzg/entrypoint-demoter/releases) and verify it against that release's `checksums.txt`. Place the extracted file next to your Dockerfile as `entrypoint-demoter`. This fragment adds it together with your own entrypoint script:

```dockerfile
COPY entrypoint-demoter /usr/local/bin/entrypoint-demoter
COPY entry.sh /entry.sh
RUN chmod 0755 /usr/local/bin/entrypoint-demoter /entry.sh

ENTRYPOINT ["/usr/local/bin/entrypoint-demoter", "--match", "/data", "/entry.sh"]
```

Choose the binary for the target image architecture, such as AMD64 or ARM64. The [Bedrock server Dockerfile](https://github.com/itzg/docker-minecraft-bedrock-server/blob/master/Dockerfile) provides another integration example.

### Local build and tests

A local build requires the Go version specified in [go.mod](go.mod) and Make. Run these commands from the repository directory:

```bash
make build
make test
```

The additional container checks require Bash and Docker:

```bash
bash test.sh
```

See [Makefile](Makefile) and [test.sh](test.sh) for the commands. The container checks build a test image and verify user permissions and stop messages, among other cases.

### License

[MIT License](LICENSE). Copyright © 2019 Geoff Bourne. The original license and attribution also apply to this fork.
