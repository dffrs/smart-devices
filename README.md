# smart-devices

A minimal Go CLI for interacting with **Shelly smart outlets** via their HTTP API.

The project currently focuses on **basic, direct device interaction**: sending simple commands to a Shelly outlet (gen 2+) and printing the responses in a human‑readable format. It is intentionally small and opinionated, acting more as a lightweight control/debugging tool than a full‑featured device manager.

---

## What it does (currently)

The tool allows you to:

- Turn a Shelly outlet **on** or **off**
- Query the outlet **status** (power, relay state, and related metadata)
- List the **supported methods** exposed by the device

All responses are returned directly from the device and printed as formatted JSON for easy inspection.

---

## Supported flags

The CLI exposes a small set of flags for interacting with the device:

- `-get status` - fetches and prints the current device status
- `-get methods` - lists the methods supported by the device
- `-set on` - turns the outlet on
- `-set off` - turns the outlet off

---

## Run

```bash
# build - this produces an executable file
make

# run
./sd [-options] [-args]
```

---

## Status

🚧 Work in progress.

