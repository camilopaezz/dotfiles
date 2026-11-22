#!/usr/bin/env bash
set -euo pipefail

repo_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd)"
command -v claude >/dev/null
ssh -o BatchMode=yes -o ConnectTimeout=10 camilo@192.168.1.91 true

# Transfer only this provider's gateway credential, never print it.
python3 - "$repo_dir" <<'PY'
import base64
import datetime
import json
import os
from pathlib import Path
import subprocess
import sys

os.umask(0o077)
repo = Path(sys.argv[1])
provider = json.loads((repo / 'files/t3/kiro-provider.json').read_text())
userdata = Path.home() / '.t3/userdata'
settings_path = userdata / 'settings.json'
settings = json.loads(settings_path.read_text()) if settings_path.exists() else {}
secrets = userdata / 'secrets'
secrets.mkdir(parents=True, exist_ok=True)
def encoded(value):
    return base64.urlsafe_b64encode(value.encode()).decode().rstrip('=')
for variable in provider['environment']:
    if not variable.get('sensitive') or not variable.get('valueRedacted'):
        continue
    name = f"provider-env-{encoded('claudeKiro')}-{encoded(variable['name'])}.bin"
    data = subprocess.run(
        ['ssh', '-o', 'BatchMode=yes', 'camilo@192.168.1.91',
         f'cat ~/.t3/userdata/secrets/{name}'],
        check=True, capture_output=True,
    ).stdout
    if not data:
        raise RuntimeError('Remote provider credential is empty')
    target = secrets / name
    temporary = target.with_suffix('.tmp')
    temporary.write_bytes(data)
    temporary.chmod(0o600)
    temporary.replace(target)
if settings_path.exists():
    stamp = datetime.datetime.now().strftime('%Y%m%dT%H%M%S%f')
    backup = settings_path.with_name(f'settings.json.before-kiro-{stamp}')
    backup.write_bytes(settings_path.read_bytes())
settings.setdefault('providerInstances', {})['claudeKiro'] = provider
temporary = settings_path.with_suffix('.tmp')
temporary.write_text(json.dumps(settings, indent=2) + '\n')
temporary.replace(settings_path)
print('Added Kiro provider to T3 Code settings.')
PY

unit_dir="${XDG_CONFIG_HOME:-$HOME/.config}/systemd/user"
mkdir -p "$unit_dir"
ln -sfn "$repo_dir/files/systemd/user/t3-kiro-tunnel.service" "$unit_dir/t3-kiro-tunnel.service"
systemctl --user daemon-reload
systemctl --user enable --now t3-kiro-tunnel.service
printf 'Kiro uses the gateway on camilo@192.168.1.91 through localhost:18317.\n'
