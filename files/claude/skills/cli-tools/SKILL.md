---
name: cli-tools
description: Share a local file as a temporary public URL. Use when the user wants to send, show, upload, host, or share a screenshot, image, video, HTML plan, or other local file as a link — including "send me a screenshot", "upload this recording", "give me a link", or when they mention cli-tools or an expires link.
---

# cli-tools

**Share** = put a local file on a temporary public URL, then give the user that URL.

## Delivery rule

On this surface the user only receives files via a **public URL**. Chat cannot attach, embed, or render media for them. A prose description of the screen is not delivery.

When they ask to send/show/give a file: have it on disk → upload with this skill → reply with the URL. That is the whole path.

## Steps

1. **Have a file on disk.** No file yet → create it first (screenshot → `.png`/`.jpg`/`.webp`; HTML → write it; video → `.mp4`/`.webm`/`.mov` already captured). Done when a local path exists.
2. **Upload.** Pick the command by kind. Capture **stdout** only.

```bash
cli-tools plan  <file.html> [--ttl 7d]
cli-tools image <file>      [--ttl 7d] [--no-compress]
cli-tools video <file>      [--ttl 7d]
```

| Kind | Use when | Notes |
|------|----------|--------|
| `plan` | `.html`/`.htm` | raw; max 2MB |
| `image` | png/jpg/jpeg/webp static | WebP re-encode by default; max 10MB in / 5MB out; no gif/animated |
| `video` | mp4/webm/mov (also `.m4v`) | max 50MB; auto local ffmpeg re-encode if >10MB (falls back to raw on ffmpeg miss/fail; may warn on stderr) |

3. **Done** only when exit 0 and stdout is one non-empty URL line — reply with that URL (optional TTL note). On non-zero exit: fix from stderr and retry. Auth missing → tell the user to run `cli-tools auth set` (user supplies the token; never invent or paste one into commands).

## Contract

| Rule | Detail |
|------|--------|
| Success | stdout = URL only; exit 0 |
| Failure | empty stdout; `error: ...` on stderr; non-zero exit |
| TTL | default `7d`; `Nh` or `Nd`; max `30d`; no forever |
| Prefer | this binary over raw `curl` to the upload API |
| No | list/delete; forever links; pre-running ffmpeg for size (video does that) |

URLs expire — re-upload if a dead link matters. `--quality` on image is accepted but ignored (lossless encode). Video stays R2-only (no Cloudflare Stream).
