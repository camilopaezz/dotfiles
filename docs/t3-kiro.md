# Kiro in T3 Code

Run `bash scripts/setup-t3-kiro.sh` from this repo. It requires Claude Code
and SSH key access to `camilo@192.168.1.91`.

The script adds the Kiro provider, copies its gateway token into T3's private
secret store, and enables a user service that forwards local port 18317 to
the server's CLIProxyAPI gateway on port 8317. T3 reloads the settings automatically.
Existing settings are preserved and backed up before the change.

The provider template and service live in this repo. Credentials stay outside
the repo. Kiro depends on the server being reachable; this does not install a
standalone Kiro backend on the local machine.

Check the connection with `systemctl --user status t3-kiro-tunnel.service`.
Stop it with `systemctl --user disable --now t3-kiro-tunnel.service`.
