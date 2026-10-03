---
name: cli-tools
description: Share a local file as a temporary public URL. Use when the user wants to send, show, upload, host, or share a screenshot, image, video, HTML plan, or other local file as a link — including "send me a screenshot", "upload this recording", "give me a link", or when they mention cli-tools or an expires link.
---

# cli-tools

Share means put a local file on a temporary public URL, then give the user that URL. In this chat that URL is the only delivery. A prose description of the screen is not delivery.

## Steps

1. Have a nonempty file on disk. Done when a local path exists. If none exists yet, write a screenshot as `.png`, `.jpg`, or `.webp`, write HTML to `.html` or `.htm`, or use a captured `.mp4`, `.webm`, `.mov`, or `.m4v`.

2. Run one `cli-tools` command. Read stdout only. The command sends the bearer token and, for `file`, `Content-Length`.

```bash
cli-tools plan  <file.html> [--ttl 7d]
cli-tools image <file>      [--ttl 7d] [--no-compress]
cli-tools video <file>      [--ttl 7d]
cli-tools file  <file>      [--ttl 24h]
```

Pick the command for how the user should receive the file. Extensions match case-insensitively. Flags may follow the path.

| Command | Use when | Behavior |
|---------|----------|----------|
| `plan` | HTML the user should open in a browser. `.html` or `.htm` only. | Raw upload. Max 2 MiB. Default TTL `7d`. |
| `image` | A static png, jpg, jpeg, or webp the user should see. | Re-encodes to lossless WebP unless `--no-compress`. Input max 10 MiB. Encoded output max 5 MiB. Rejects gif and animated images. `--quality` is accepted and ignored. Default TTL `7d`. |
| `video` | An mp4, m4v, webm, or mov the user should play. | Max 50 MiB. Over 10 MiB, re-encodes with local ffmpeg and uploads the original if ffmpeg is missing, fails, or does not shrink the file. Pass the captured file. Default TTL `7d`. |
| `file` | Any other nonempty local file, when a download is enough. | Unchanged bytes as `application/octet-stream`. Max 100 MiB. Default TTL `24h`. The Worker requires `Content-Length`; the command sets it and preserves the original download filename and extension. |

Viewable HTML, images, and video go through `plan`, `image`, or `video`. `file` always downloads.

TTL override is `Nh` or `Nd`, minimum 1, maximum `30d`.

3. Reply with the URL. A short TTL note is optional. Done only when the exit code is 0 and stdout is one non-empty URL line. That line is the URL. A stderr `warning:` from a skipped video re-encode can appear on success. Stderr is not the URL.

On a non-zero exit, stdout is empty and stderr is one `error:` line. Fix from that line and retry.

If the error is `no token: set CLI_TOOLS_TOKEN or run auth set`, tell the user to run `cli-tools auth set` and supply the token themselves. Do not invent a token or put one in a command. When `CLI_TOOLS_TOKEN` is set, it wins over the config file. On `upload failed (401)`, the stored token was rejected. Ask the user to run `auth set` again.

URLs expire. Upload again when a dead link matters. There is no list or delete command.
