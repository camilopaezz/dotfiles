---
name: computer-use
description: Computer use on this Hyprland desktop. Use when clicking or dragging a GUI, or when a click misses the target.
---

# Computer use

Match every click to the screenshot it came from. Hardware cursors often stay out of the shot — `hyprctl cursorpos` is the landing read.

## Spaces

Two spaces, one at a time:

- **Desktop** — full-screen screenshot. Click `x,y` as desktop pixels. No `relative`.
- **Frame** — window-cropped screenshot. Click with that `window_id` and `relative: true`. `(0,0)` is `list_windows` `bounds` origin, border included.

Click in the screenshot's `coordinate_width` × `coordinate_height` space. Leave `max_width` / `max_height` unset (or at least as large as the frame) so that space equals the image.

Aim at the **center** of the control, including small badges. This compositor rounds window corners.

Intended desktop point: desktop click → `(x, y)`. Frame click → `(bounds.x + x, bounds.y + y)`.

## Process

1. **Resolve the window.** `list_windows` / `get_app_state`. Done when you have `window_id` and `bounds`, or you are aiming at the bar / desktop itself.
2. **Pick the space.** One window → frame crop. Bar, desktop, or multiple windows → desktop. Electron / Helium / T3 Code often expose only a frame in AT-SPI — coordinates are the path. Named widgets in the tree → `element_index` instead of pixels.
3. **Click, then read landing.** Click in the chosen space. Run `hyprctl cursorpos`. Done when it matches the intended desktop point **and** the UI reacted (dialog, toggle, focus).
4. **On a miss, classify, then one retry.**
   - `cursorpos` matches the intended point, UI did not react → the aim was off. Nudge in the **same** space toward the control (the offset you can see: highlight, hover, or "just right of the avatar").
   - `cursorpos` is off by about `bounds.x` / `bounds.y` → the other space. Retry that same visual target with `relative` flipped.
   - Confirming a dialog while AT-SPI is a bare frame → click the button. `Enter` goes to whatever is focused.

   Done when it hits, or when a second miss means the control is not where the screenshot suggested.
