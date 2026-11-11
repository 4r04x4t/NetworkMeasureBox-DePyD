# Known Python faults

The shared fault register of the native build (roadmap section 6B.8 point 3; *Register of Python Faults the Native
Build Must Not Repeat*). Every fault found in the Python version, fixed or not, is a requirement for every native
code base: the component is designed so that the fault cannot happen, and it has a test that fails if it does
(Quality Bar points 2, 6 and 7). A native component is not accepted while one of its entries has no test.

- **Sources:** the roadmap's register, `docs/DEPYTHONIZATION_ROADMAP.md` of the Python repository at commit
  `ad3758c`, lines 610-619 (one entry per fault, the wording of the source kept; square brackets mark words added
  or left out where one sentence of the source holds several faults). Package F-03 adds the faults of
  the private audit and data-quality lists and of the Quality Bar examples, writes the correct behaviour where the
  source does not state it, and assigns the package that must prove each entry.
- **Open security holes** of the Python version are recorded here only by category until they are fixed, never
  with details. None of the entries below is one: the source lists them under "Found and fixed in the Python
  version up to 2026-11-10" (line 608), the measurement-data faults of line 617 as "being fixed".
- **Status per code base:** Go (step 1), C (step 2, the router tiers) and the Python reference. "fixed (test
  passes)" is written only when the named native test has run and passed in the gate.

## Entry format

Each entry has exactly these lines (machine-checked later by F-07):

```
### KF-0000: <short title>

- Area: <area as in the source>
- What went wrong: <copied from the source>
- Correct behaviour: <copied from the source, or "not stated in the source; F-03 writes it">
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line <n>
- Python fix commit: <id, or "not recorded in the source">
- Package that proves it: <id, or "to be assigned by F-03">
- Spec vector: <id, or "none yet (F-07)">
- Native test, Go: <test name, or "none yet">
- Native test, C: <test name, or "none yet (step 2)">
- Status, Go: open | fixed (test passes) | not applicable (reason)
- Status, C: not started | open | fixed (test passes) | not applicable (reason)
- Status, Python: fixed | being fixed | open (as the source says)
```

## Entries

### KF-0001: the update's test run as root overwrote the real settings

- Area: Privileged helper and updates
- What went wrong: the update's test run as root overwrote the real settings
- Correct behaviour: tests run unprivileged in a sandbox with a copy of the settings
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 610
- Python fix commit: not recorded in the source
- Package that proves it: F-00 for the gate (the update's own test run comes with the update package, assigned by F-03)
- Spec vector: none yet (F-07)
- Native test, Go: `tools/gate-selftest.sh` case `root`: every gate script refuses uid 0 with exit 2 (`tools/lib.sh` `nmb_init`)
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0002: root helper actions took paths, packages and files from the caller ...

- Area: Privileged helper and updates
- What went wrong: root helper actions took paths, packages and files from the caller without an allow-list
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 610
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0003: a stray PYTHONPATH or working directory could make root code ...

- Area: Privileged helper and updates
- What went wrong: a stray `PYTHONPATH` or working directory could make root code import the wrong module
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 610
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0004: a plaintext fallback wrote the settings to the boot partition

- Area: Privileged helper and updates
- What went wrong: a plaintext fallback wrote the settings to the boot partition
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 610
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0005: the diagnostic bundle, crash reports and the boot-partition copy ...

- Area: Secrets and personal data
- What went wrong: the diagnostic bundle, crash reports and the boot-partition copy carried tokens, passwords, user names and multi-line values
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 611
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0006: /backup send put the plain archive with every secret into Telegram

- Area: Secrets and personal data
- What went wrong: `/backup send` put the plain archive with every secret into Telegram
- Correct behaviour: encrypted by default
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 611
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0007: it checked the chat, not the person (in a group every member was ...

- Area: Telegram bot
- What went wrong: it checked the chat, not the person (in a group every member was an admin)
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 612
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0008: button data could carry any command

- Area: Telegram bot
- What went wrong: button data could carry any command
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 612
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0009: edited messages ran again

- Area: Telegram bot
- What went wrong: edited messages ran again
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 612
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0010: commands queued while the box was off ran hours later

- Area: Telegram bot
- What went wrong: commands queued while the box was off ran hours later
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 612
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0011: a reload kept the old token

- Area: Telegram bot
- What went wrong: a reload kept the old token
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 612
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0012: an OpenVPN profile name could leave its directory and read the ...

- Area: VPN
- What went wrong: an OpenVPN profile name could leave its directory and read the server's keys
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 613
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0013: a deleted client could still connect (no revocation)

- Area: VPN
- What went wrong: a deleted client could still connect (no revocation)
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 613
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0014: a config line without a final newline glued the next directive ...

- Area: VPN
- What went wrong: a config line without a final newline glued the next directive onto it and stopped OpenVPN
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 613
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0015: standalone and transparent mode had no firewall for the box itself

- Area: Firewall
- What went wrong: standalone and transparent mode had no firewall for the box itself
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 614
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0016: PPPoE came up before the router's rules

- Area: Firewall
- What went wrong: PPPoE came up before the router's rules
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 614
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0017: forwarding was switched on before the rules and stayed on when ...

- Area: Firewall
- What went wrong: forwarding was switched on before the rules and stayed on when they failed
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 614
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0018: WAN2 and IPv6 in the fallback were unfiltered

- Area: Firewall
- What went wrong: WAN2 and IPv6 in the fallback were unfiltered
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 614
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0019: teardown flushed other software's NAT rules

- Area: Firewall
- What went wrong: teardown flushed other software's NAT rules
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 614
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0020: fallback rules piled up

- Area: Firewall
- What went wrong: fallback rules piled up
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 614
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0021: a firewall that followed only the settings would have cut off a ...

- Area: Firewall
- What went wrong: a firewall that followed only the settings would have cut off a VPN server that was running
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 614
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0022: any datagram reaching the upstream port was taken as the answer ...

- Area: DNS ad-blocker
- What went wrong: any datagram reaching the upstream port was taken as the answer and cached for everyone (cache poisoning)
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 615
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0023: malformed names were read as other names

- Area: DNS ad-blocker
- What went wrong: malformed names were read as other names
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 615
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0024: one socket error stopped DNS for the whole network

- Area: DNS ad-blocker
- What went wrong: one socket error stopped DNS for the whole network
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 615
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0025: pings that could not even be sent (no route, link down) were not ...

- Area: Measurement honesty
- What went wrong: pings that could not even be sent (no route, link down) were not counted as lost, so such outages were never recorded
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 616
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0026: an answer from the local router ended an internet outage

- Area: Measurement honesty
- What went wrong: an answer from the local router ended an internet outage
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 616
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0027: the Web UI showed made-up data (fake clients, another device's ...

- Area: Measurement honesty
- What went wrong: the Web UI showed made-up data (fake clients, another device's firmware number, invented rates)
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 616
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0028: Ookla's terms were accepted silently for the user

- Area: Measurement honesty
- What went wrong: Ookla's terms were accepted silently for the user
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 616
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0029: the automatic-engine badge did not follow the setting

- Area: Measurement honesty
- What went wrong: the automatic-engine badge did not follow the setting
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 616
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed

### KF-0030: packet loss written as 0 % by engines that do not measure it

- Area: Measurement data
- What went wrong: packet loss written as 0 % by engines that do not measure it
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 617
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: being fixed

### KF-0031: outages caused by the box's own mode changes, updates and settings ...

- Area: Measurement data
- What went wrong: outages caused by the box's own mode changes, updates and settings counted against the ISP
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 617
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: being fixed

### KF-0032: ISP name and line type taken from the settings instead of the ...

- Area: Measurement data
- What went wrong: ISP name and line type taken from the settings instead of the measurement
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 617
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: being fixed

### KF-0033: invented contract and guaranteed speeds used for verdicts when ...

- Area: Measurement data
- What went wrong: invented contract and guaranteed speeds used for verdicts when none were set
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 617
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: being fixed

### KF-0034: missing values exported as 0 ms / 0 %

- Area: Measurement data
- What went wrong: missing values exported as `0 ms` / `0 %`
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 617
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: being fixed

### KF-0035: times without a time zone

- Area: Measurement data
- What went wrong: times without a time zone
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 617
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: being fixed

### KF-0036: the public IP missing or replaced by the LAN address

- Area: Measurement data
- What went wrong: the public IP missing or replaced by the LAN address
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 617
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: being fixed

### KF-0037: all engines of one mixed run sharing one timestamp

- Area: Measurement data
- What went wrong: all engines of one mixed run sharing one timestamp
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 617
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: being fixed

### KF-0038: a generic web address given as the result link

- Area: Measurement data
- What went wrong: a generic web address given as the result link
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 617
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: being fixed

### KF-0039: the diagnostic bundle concluding "double NAT" from a PPP peer ...

- Area: Measurement data
- What went wrong: the diagnostic bundle concluding "double NAT" from a PPP peer address and "LAN fault" from outage bursts
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 617
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: being fixed

### KF-0040: the diagnostic bundle not knowing its own version

- Area: Measurement data
- What went wrong: the diagnostic bundle [...] not knowing its own version
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 617
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: being fixed

### KF-0041: a result did not say whether the box measured behind another ...

- Area: Measurement setup
- What went wrong: a result did not say whether the box measured behind another router or directly on the ISP
- Correct behaviour: roadmap section 4.29 (the source refers to it; F-03 writes it out)
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 618
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed (2026-11-10)

### KF-0042: the mode was stored only from 2026-11-10 noon on, read after the ...

- Area: Measurement setup
- What went wrong: the mode was stored only from 2026-11-10 noon on (read after the test, so a switch during it went unseen)
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 618
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed (2026-11-10)

### KF-0043: the link and the position were never stored

- Area: Measurement setup
- What went wrong: [stored:] the link and the position never
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 618
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed (2026-11-10)

### KF-0044: skipped rows had no context at all

- Area: Measurement setup
- What went wrong: skipped rows had no context at all
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 618
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed (2026-11-10)

### KF-0045: the diagnostic bundle reported the configured mode (auto read as ...

- Area: Measurement setup
- What went wrong: the diagnostic bundle reported the configured mode (`auto` read as "LAN probe behind a router" also when the box was the router)
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 618
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed (2026-11-10)

### KF-0046: outages and ping minutes had no setup

- Area: Measurement setup
- What went wrong: outages and ping minutes had no setup
- Correct behaviour: not stated in the source; F-03 writes it
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 618
- Python fix commit: not recorded in the source
- Package that proves it: to be assigned by F-03
- Spec vector: none yet (F-07)
- Native test, Go: none yet
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed (2026-11-10)

### KF-0047: tests depended on the live network mode, router state file, line ...

- Area: Tests that depended on the machine
- What went wrong: [tests depended on] the live network mode, router state file, line traffic, existing tunnel interfaces, the wall or monotonic clock, a cache in `/tmp` and third-party websites
- Correct behaviour: native tests are hermetic by design
- Source: DEPYTHONIZATION_ROADMAP.md@ad3758c line 619
- Python fix commit: not recorded in the source
- Package that proves it: F-00 for the test runner, F-06 for the fakes (clock, network, servers)
- Spec vector: none yet (F-07)
- Native test, Go: `tools/test.sh` runs A-G (offline namespace, empty HOME and TMPDIR, no test cache, repository unchanged); `tools/gate-selftest.sh` cases `clean` (probe test `TestHermetic`), `testwrites`
- Native test, C: none yet (step 2)
- Status, Go: open
- Status, C: not started
- Status, Python: fixed
