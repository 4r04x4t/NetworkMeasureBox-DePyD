# I-09: Comparison of the base candidates for roadmap 6A, and proposal

Package I-09 (work plan section 5.3 point 1 and 5.4: "Comparison and proposal for 6A; from I-01 to I-08; the owner
decides"). Owner decision 4 (work plan section 7): the base of NMB's own distribution, the init system and the A/B
update tool are chosen from this comparison. This file proposes; it decides nothing.

- Date: 2026-11-11.
- Sources, read in full:
  - roadmap 6A as on the Python repository at `bb0a553` (`docs/DEPYTHONIZATION_ROADMAP.md`, "## 6A." to "## 6B.",
    with its eight criteria and *Findings from the images*), and 6B.5 (sandboxing) of the same file;
  - the work plan at `bb0a553` (`docs/NATIVE_WORK_PLAN.md`), section 5 and section 7 (points 4, 5, 6 and 10 to 24;
    here mainly 19, migration of the owner's box component by component, and 20, packages for every distribution
    the Python version supports);
  - the eight image records in this folder (table below).
- **Rule of this file.** Every value comes from a record and names it as "(I-0n §s)": the record and its section.
  A value that no record gives is written "not determined", never filled in from memory. A sentence marked
  **Proposal** or **Inference** is this package's reasoning from the cited values, not a measurement.
- QEMU figures are emulation figures (TCG unless KVM is said) on a PC that was often loaded; they compare the bases
  only roughly and say nothing about real hardware (I-01 §13, I-06 §13, I-07 §13, I-08 §13).

| Id | Record | What it covers |
| :--- | :--- | :--- |
| I-01 | `openwrt-25.12.6-x86-64.md` | OpenWrt x86-64, booted (TCG) |
| I-02 | `openwrt-25.12.6-armsr-armv8.md` | OpenWrt generic arm64 (UEFI), booted (TCG); SDK, ImageBuilder, toolchain |
| I-03 | `openwrt-25.12.6-armsr-armv7.md` | OpenWrt 32-bit ARM, booted (TCG); test guest of F-06 |
| I-04 | `openwrt-25.12.6-x86-generic.md` | OpenWrt 32-bit x86, booted (KVM and TCG); cross-check of I-01 |
| I-05 | `openwrt-25.12.6-bcm27xx-bcm2711-rpi-4.md` | OpenWrt Raspberry Pi 4, **taken apart only, not booted** |
| I-06 | `debian-13.7.0.md` | Debian 13.7 amd64 (minimal build booted, TCG) and arm64 with the Pi firmware (**built, not booted**) |
| I-07 | `alpine-3.24.4.md` | Alpine 3.24.4 x86-64 ISO (diskless, booted, KVM and TCG) and Pi archive (**not booted**) |
| I-08 | `fedora-iot-44.md` | Fedora IoT 44 aarch64 raw image and x86_64 ISO; **no guest reached `multi-user.target`** |

Only the four candidates that roadmap 6A orders first were taken apart (OpenWrt, Debian, Alpine, Fedora IoT). The
other rows of the 6A candidate table (Raspberry Pi OS Lite, VyOS, IPFire, Ubuntu Core, NixOS, Buildroot/Yocto and
the excluded ones) have no record; 6A says they are looked at only if these four fail the criteria.

## 1. The 6A criteria, one column per candidate

"OpenWrt" means 25.12.6 on every target (I-01 to I-05); where the Pi 4 image (I-05) differs, it is said.

| Criterion (6A) | OpenWrt 25.12.6 (I-01..I-05) | Debian 13.7 (I-06) | Alpine 3.24.4 (I-07) | Fedora IoT 44 (I-08) |
| :--- | :--- | :--- | :--- | :--- |
| **1. Stock kernel, Pi 4 and x86-64, NMB's network features** | Kernel 6.12.112 on every target. Bridge, VLAN built in; nftables, PPP/PPPoE modules in the image; bonding, WireGuard, USB network and every tethering driver of 4.27 as published kmods; cfg80211/mac80211 not in the kernel but from backports 6.18.55 packages (I-01 §4, I-02 §4, I-05 §4). `kmod-wireguard` and `mac80211_hwsim` installed and worked in the arm64 guest (I-02 §9a). Pi 4: GENET and `bcm2835_wdt` built in, `brcmfmac` with firmware in the image (I-05 §4, §5a, §5b) | Kernel 6.12.107 for amd64 and arm64; bridge, bonding, VLAN, nftables, WireGuard, PPPoE, cfg80211/mac80211/hwsim and every tethering driver as modules; Pi 4 platform (`ARCH_BCM2835`, `PCIE_BRCMSTB`, `BCMGENET`, DTBs) in the arm64 kernel; modules loaded live on amd64 only (I-06 §4, §16) | Kernel 6.18.60: `lts` on x86-64, `rpi` (Raspberry Pi kernel tree) on the Pi; every listed feature in both; `NET_TEAM` not set on the Pi; loaded live on x86-64 only (I-07 §4, §16) | Kernel 6.19.10 for both (7.3.3 in the F44 updates); every listed feature in both, Pi 4 platform included; `mac80211_hwsim` built but not shipped (in `kernel-modules-internal`); no module loaded live (I-08 §4, §7, §16) |
| **2. Security updates for years, free** | Not determined: no record gives a support period for 25.12 | Full support to 2028-08-09, LTS to 2030-06-30, no subscription; the LTS architecture list is "to be determined" on the Debian wiki (I-06 §3) | v3.24 end of life 2028-06-01; `main` about 2 years, `community` only to the next release; every criterion-3 tool is in `main` (I-07 §3) | **About 13 months per release**: `SUPPORT_END=2027-05-19`, schedule end of life 2027-06-02; a yearly `rpm-ostree rebase` needed, not tested (I-08 §3) |
| **3. Fits `nmbd` (init and tools)** | procd; in the image `nft` 1.1.6, dnsmasq 2.93, pppd 2.5.2; hostapd, wpa_supplicant, iw, WireGuard tools, OpenVPN 2.7.8, full iproute2 6.18.0 only in the feeds; `ip` is the BusyBox applet (I-01 §7, §9a). The Pi 4 image also has hostapd and wpa_supplicant (`wpad-basic-mbedtls`) and iw 6.17 (I-05 §7) | systemd 257; every tool from `main`: `nft` 1.1.3, dnsmasq 2.91, hostapd and wpa_supplicant 2.10, iw 6.9, pppd 2.5.2, wireguard-tools 1.0.20210914, OpenVPN 2.6.14, full iproute2 6.15.0; none needs Python (I-06 §7) | OpenRC 0.63.2; every tool from `main`: `nft` 1.1.6, dnsmasq 2.92, hostapd and wpa_supplicant 2.11, iw 6.17, pppd 2.5.3, wireguard-tools 1.0.20260223, OpenVPN 2.7.7, iproute2 7.0.0; `wg-quick` pulls `bash` (I-07 §7) | systemd 259.5; in the base `nft` 1.1.6, wpa_supplicant 2.11, iw 6.17, iproute 6.17.0; **dnsmasq, hostapd, pppd, wireguard-tools, OpenVPN not in the base**, in the F44 repositories (I-08 §7) |
| **4. Read-only image, A/B, reproducible** | Read-only squashfs root with an overlay; sysupgrade overwrites the one root; no A/B; on the Pi 4 no automatic way back after a failed boot (I-01 §2, §3; I-05 §3). Reproducibility: not determined | apt in place, one root, no A/B; RAUC 1.13, SWUpdate 2024.12.1, `systemd-repart`/`systemd-boot` 257.13, `mmdebstrap`, `debos`, `mkosi`, `ostree` packaged in trixie `main`, none tested; one independent rebuild gave the same packages but not the same bytes (the initrd differs by a few bytes) (I-06 §3, §13) | Diskless mode is a read-only base by design (squashfs modloop, RAM root, settings in one `apkovl` written by rename; commit and restore worked); kernel and modloop are replaced in place by `mv`, no second slot, no fallback; no RAUC or SWUpdate package in v3.24 (I-07 §2, §3, §12a) | **Built in**: read-only `/usr` (composefs), ostree deployments as A/B, signed commits, greenboot health checks with a GRUB boot counter (3 attempts) and rollback; reproducibility of an own compose not tried (I-08 §2, §3, §13) |
| **5. Does not fight NMB for the network** | netifd and firewall4 enabled; both off with `disable`. The firewall service's `stop` deletes every nftables table, NMB's too; plain `fw4 stop` deletes only its own; netifd keeps foreign addresses on its devices across `reload` and drops them on `restart`; it leaves devices it does not manage alone (I-01 §9, §9a; I-02 §9; I-03 §9). Pi 4: the same scripts byte for byte, live tests not run (I-05 §9) | ifupdown enabled but idle with the default configuration; left NMB's bridge alone; `nftables.service` disabled, but its start, restart and stop wipe every table, so it stays masked; `systemd-networkd` disabled; hostapd, wpa_supplicant and openvpn units enabled by their packages (I-06 §9) | Nothing enabled: no firewall installed, `networking` in no runlevel, packages do not enable their services. If installed, the nftables service wipes every table on start, reload, restart, stop; ifupdown-ng removes foreign addresses only on interfaces it manages; NMB's bridge survived (I-07 §9) | **NetworkManager, firewalld and ModemManager enabled**, `systemd-resolved` enabled; `nftables.service` off but its reload and stop flush every table; masking proposed, not tested; live behaviour of firewalld and NetworkManager towards NMB's tables and interfaces not determined (I-08 §9) |
| **6. Size, memory, C library; Ookla and RTR** | musl 1.2.5; roots 11.7 to 20.4 MB unpacked; 21,772 to 37,140 kB used after boot in the guests; Pi 4 not booted (I-01 §11, §12; I-02 §11, §12; I-03 §11, §12; I-04 §12; I-05 §11, §12). Ookla and RTR on musl: not determined | glibc 2.41; root 360 MB (amd64 with kernel and tools), 599.5 MB (arm64 with the Pi firmware); 61,412 kB used after boot in a 512 MB guest (I-06 §11, §12). Ookla and RTR: not determined (I-06 §13) | musl 1.2.6; x86-64 ISO 370 MB (modloop 308 MB, mostly firmware), Pi boot content 110 MB; 37,564 kB used at login (512 MB, KVM), 53,996 kB with the tool set; tool-set closure 27 MB (I-07 §11, §12). Ookla and RTR: not determined (I-07 §13) | glibc 2.43; aarch64 card image 6 GiB (root 2.1 GB used), OS tree 1.86 GB (aarch64) and 1.38 GB (x86_64); memory after boot not determined (I-08 §11, §12a). Ookla and RTR: not determined (I-08 §13) |
| **7. Licence lets NMB ship its image** | A licence field per package (mostly GPL-2.0; `ppp` is BSD-4-Clause); x86-64 firmware packages and six Pi 4 firmware packages have an empty licence field; the Pi VideoCore firmware carries `LICENCE.broadcom` (I-01 §6; I-05 §2a, §5, §6) | Minimal base all from `main`; the Pi boot files and Wi-Fi firmware from `non-free-firmware`, their licences not read (I-06 §6, §13) | Base packages under free licences; firmware packages `L: custom`, texts not read (I-07 §6, §13) | Distribution free; firmware `LicenseRef-Callaway-Redistributable-no-modification-permitted`, Pi boot files `LicenseRef-Fedora-Firmware`; whether that suffices not judged (I-08 §6, §13) |
| **8. The base itself brings no Python** | No interpreter in any image (I-01..I-04 §8). The Pi 4 image carries one Python script, `otpset` from `bcm27xx-utils`, without an interpreter; an NMB image must leave it out (I-05 §8) | No interpreter in the minimal base; a few Python files without an interpreter (`routel` from iproute2, gdb pretty-printers); Python arrives with the "standard" task (`reportbug`) or through Recommends, so the image needs the release check (I-06 §8) | No interpreter at boot or with the tool set, none in the dependency closure; only `/sbin/routel` has a Python shebang (I-07 §8) | **Fails**: `python3` 3.14.3 with 13 `python*` packages, pulled in by `firewalld`, `policycoreutils-python-utils`, `setools-console`; only an own compose could change that, untested (I-08 §8) |

Rows that 6B.5 and the 6A update plan add:

| | OpenWrt | Debian | Alpine | Fedora IoT |
| :--- | :--- | :--- | :--- | :--- |
| **Landlock** | Absent on every target (`CONFIG_SECURITY` not set) (I-01..I-05 §4) | `SECURITY_LANDLOCK=y` on amd64 and arm64; active in the amd64 guest; arm64 not booted (I-06 §4) | Active on x86-64; **not set in `linux-rpi`** (`LSM=""`) (I-07 §4) | `y` on both; "landlock: Up and running." in the aarch64 guest's kernel log (I-08 §4, §12a) |
| **Other confinement in the base** | procd `ujail` and seccomp installed; dnsmasq and ntpd run inside `ujail` as their own users (I-01 §6, §12) | AppArmor in the active LSM list; the `apparmor` userspace package is not installed (I-06 §4) | AppArmor built, not active (I-07 §4) | SELinux enforcing (I-08 §4, §12a) |
| **cgroup2 at boot** | Mounted (guests); on the Pi 4 the device tree's boot arguments carry `cgroup_disable=memory`, effect not determined (I-01 §4; I-05 §2, §4) | Mounted (`nsdelegate,memory_recursiveprot`) (I-06 §4) | **Not mounted**; `rc-service cgroups start` mounts it with `cpuset cpu io memory hugetlb pids dmem` (I-07 §12a) | `MEMCG`, `MEMCG_V1` and the other controllers `y`; mount not determined (I-08 §4) |
| **Update by the base** | sysupgrade, one root (I-01 §3, I-05 §3) | apt/dpkg in place (I-06 §3) | `lbu commit`, `apk` into RAM plus cache, `update-kernel` (I-07 §3) | rpm-ostree deployments, greenboot (I-08 §3) |
| **Real A/B** | No; must be built | No; must be built from packaged tools | No; must be built, no packaged tool | Yes, ready-made |

Verdict per criterion, from the two tables (Inference):

| Criterion | OpenWrt | Debian | Alpine | Fedora IoT |
| :--- | :--- | :--- | :--- | :--- |
| 1 | met (Wi-Fi stack from backports packages) | met | met | met (`hwsim` only as an extra package) |
| 2 | not determined | met (arm64 LTS not confirmed) | met, about 2 years | **not met** |
| 3 | met through the feeds | met | met | partly (5 tools to add) |
| 4 | partly (read-only, no A/B) | not in the base (blocks packaged, untested) | partly (read-only, no A/B) | met |
| 5 | met, with two cautions | met, with masks | met | not determined, three managers on by default |
| 6 | smallest; musl | largest of the three that pass 8; glibc | small; musl | largest; glibc; memory not determined |
| 7 | open: firmware licence fields empty | open: firmware licences not read | open: firmware `custom` | open: firmware licences not judged |
| 8 | met (leave out `otpset` on the Pi 4) | met (release check) | met | **not met** |

Fedora IoT fails criteria 2 and 8 as shipped; criterion 8 is the owner's Python-free end state (6A intro). It is
not proposed below.

## 2. Per target: what each base gives NMB and what it lacks

### 2.1 x86-64 (PC or appliance)

| Point | OpenWrt x86-64 (I-01; I-04 as 32-bit cross-check) | Debian amd64 (I-06, builds A and B) | Alpine x86-64 (I-07) | Fedora IoT x86_64 (I-08) |
| :--- | :--- | :--- | :--- | :--- |
| Kernel features | Every NMB option as built-in or published kmod; seccomp, namespaces, cgroup2 yes; **no Landlock**; watchdog drivers as kmods, no `/dev/watchdog` in the guest (I-01 §4) | Every option; Landlock and AppArmor active; cgroup2; `iTCO_wdt` gave `/dev/watchdog` (I-06 §4) | Every option; Landlock active; cgroups not mounted at boot; watchdog devices present (I-07 §4, §12a) | Every option in the config; Landlock `y`; SELinux; `hwsim` not shipped; nothing checked live on x86_64 (I-08 §4, §12b) |
| Tools of criterion 3 | `nft`, dnsmasq, pppd in the image; the rest from the feeds; BusyBox `ip` (full iproute2 in the feeds) (I-01 §7, §9a) | All from `main`, full iproute2 (I-06 §7) | All from `main`, full iproute2 after install (I-07 §7) | 5 of them missing from the base (I-08 §7) |
| Python-free | Yes (I-01 §8) | Yes, with the release check (I-06 §8) | Yes (I-07 §8) | No (I-08 §8) |
| Own network and firewall manager; how it is switched off | netifd and firewall4 on; `disable` both; never run the firewall service's `stop` while NMB's tables are loaded (use `fw4 stop`, or `disable` and reboot); netifd drops foreign addresses on its devices on `restart` (I-01 §9, §9a) | ifupdown idle (remove or `disable networking`); mask `nftables.service` and `systemd-networkd`; disable or mask the package units of hostapd, wpa_supplicant, openvpn (I-06 §9) | Nothing on by default; keep `nftables` and `networking` out of every runlevel, or list no NMB interface in `/etc/network/interfaces` (I-07 §9) | Mask or leave out NetworkManager, firewalld, ModemManager, `iwd`, `nftables.service`; untested (I-08 §9) |
| Update method; A/B | sysupgrade writes the disk; one root; no A/B (I-01 §3) | apt in place; A/B to be built with RAUC, SWUpdate or `systemd-repart` (packaged, untested) (I-06 §3) | lbu, apk cache, `update-kernel` by `mv`; A/B to be NMB's own (I-07 §3) | rpm-ostree with greenboot, ready (I-08 §3) |
| RAM after boot | 37,140 kB used, 154,232 kB available (256 MB guest, TCG) (I-01 §12); 21,772 kB used in the 32-bit KVM guest (I-04 §12) | 61,412 kB used, 411,804 kB available (512 MB guest, TCG) (I-06 §12) | 37,564 kB used, 401,736 kB available (512 MB, KVM); 53,996 kB with 78 packages (I-07 §12) | Not determined: the installer does not fit in 512 MB, nothing was installed (I-08 §12b) |
| Flash | Disk image 126,123,520 bytes; root 20,384,863 bytes unpacked (I-01 §11) | Root with kernel and tools 360,016,018 bytes; docs, man pages and locales about 58 MB of it, cuts not measured (I-06 §11) | ISO 370,147,328 bytes, of which modloop 307,982,336 (mostly firmware); RAM root 11,640 KiB at login (I-07 §11) | OS tree 1,382,684,853 bytes (I-08 §11) |
| Boot (emulation only) | `init complete` about 109 s after kernel start, TCG (I-01 §12); about 18 s and 22 s under KVM for 32-bit (I-04 §12) | 2 min 11 s to `graphical.target`, TCG (I-06 §12) | 29.1 s from QEMU start to `login:`, KVM (I-07 §12) | Not determined (I-08 §12b) |
| Listening after boot; defaults NMB must change | dropbear 22, uhttpd 80/443 on all addresses, dnsmasq 53; no root password; password login for root over SSH on (I-01 §12) | None listening (I-06 §12) | None listening; the diskless ISO has no root password (I-07 §12) | `sshd` enabled; root locked (I-08 §6); ports not determined |
| Licences of what ships | Mostly GPL-2.0; BSD-4-Clause `ppp` (advertising clause, for the 6B.3A list); 3 firmware packages without a licence field (I-01 §5, §6) | All from `main` (I-06 §6) | Free licences; firmware `custom`, not read (I-07 §6) | 263 of 383 packages GPL variants; firmware no-modification licence, not judged (I-08 §6) |
| Support period | Not determined | 2028-08-09, LTS 2030-06-30 (I-06 §3) | 2028-06-01 (I-07 §3) | 2027-05-19 (I-08 §3) |
| Build inputs in the set | The x86/64 tree (SDK, ImageBuilder) is not in the set (I-01 §1); the x86/generic tree is, 32-bit only (I-04 §15) | `mmdebstrap` built the roots without root (I-06 §1, §14) | `alpine-make-rootfs`, `mkinitfs`, `abuild` in `main` (I-07 §3) | `treefile.json` of the compose is in the commit (I-08 §16) |

### 2.2 Raspberry Pi 4 (arm64)

**No candidate was booted on a Pi 4.** Every run-time value below is "not determined" for that reason; the Pi 4 boot
needs the owner's box under work plan 7 point 5 (I-05 §12, I-06 §13, I-07 §13, I-08 §13).

| Point | OpenWrt bcm2711 (I-05; userland behaviour from I-02) | Debian arm64 (I-06, build C) | Alpine Pi (I-07) | Fedora IoT aarch64 (I-08) |
| :--- | :--- | :--- | :--- | :--- |
| Boot path (read from the files) | VideoCore firmware reads `config.txt`, `cmdline.txt`, starts `kernel8.img`; no UEFI, no GRUB, no U-Boot (I-05 §2) | `raspi-firmware` fills `/boot/firmware` (FAT) with the GPU firmware, DTBs and **copies** of kernel and initrd; `config.txt` names them; root on partition 2 (I-06 §2) | Firmware → `vmlinuz-rpi` + initramfs from the FAT partition; diskless RAM root (I-07 §2) | Firmware → U-Boot → shim → GRUB → BLS entry (I-08 §2) |
| Kernel features | 6.12.112 with GENET, `brcmfmac`, `bcm2835_wdt` built in; **no Landlock**; `cgroup_disable=memory` in the DT boot arguments (I-05 §2, §4) | 6.12.107 with the Pi 4 platform; Landlock `y` in the config, not seen live on arm64 (I-06 §4) | 6.18.60 Raspberry Pi tree; **no Landlock**; `NET_TEAM` not set; IPv6 as a module; `linux-lts` 6.18.63 also exists for aarch64, whether it boots on a Pi 4 not checked (I-07 §4, §13) | 6.19.10 with the Pi 4 platform; Landlock `y`, up in the QEMU `virt` guest (I-08 §4, §12a) |
| Pi Wi-Fi | `brcmfmac` (backports) with the CYW43455 firmware and the Pi 4 NVRAM file; cfg80211 installed, `mac80211` not (I-05 §5a) | `firmware-brcm80211` (non-free-firmware) carries the 43455 files (I-06 §5) | 43455 firmware for the Pi 4 in the modloop (I-07 §5) | 43455 firmware and NVRAM in the commit (I-08 §5) |
| Tools of criterion 3 | Also hostapd, wpa_supplicant (`wpad-basic-mbedtls`) and iw in the image; WireGuard tools and OpenVPN from the feeds (I-05 §7, §9a) | The same set as amd64 (`command -v` all found) (I-06 §7) | The same versions in aarch64 `main`, not run on arm64 (I-07 §7) | The same gaps as x86_64 (I-08 §7) |
| Python-free | No interpreter; `otpset` (and three Perl scripts) from `bcm27xx-utils` to leave out (I-05 §8) | No interpreter in build C (I-06 §8) | Not checked on arm64; the closure has none (I-07 §8) | No (I-08 §8) |
| Own managers | netifd and firewall4, the same scripts as I-01; live tests not determined; no WAN in the default configuration; Wi-Fi AP generated disabled (I-05 §5a, §9) | As amd64; the same units enabled (I-06 §7) | As x86-64 (I-07 §9) | As x86_64 (I-08 §9) |
| Update method; A/B | sysupgrade needs images with fwtool metadata (`REQUIRE_IMAGE_METADATA=1`), rewrites the boot partition with the root, no automatic way back; the SPI EEPROM bootloader is not handled (no `rpi-eeprom`) (I-05 §3) | apt in place; how a slot switch would work on the Pi firmware boot path: not determined; `flash-kernel` and `u-boot-tools` are in the arm64 pool (I-06 §2, §3) | `update-kernel` `mv`s new files, DTBs into the partition root (I-07 §3) | rpm-ostree and greenboot; the counter lives in GRUB, which runs under U-Boot here (I-08 §2, §3) |
| RAM after boot | Not determined (I-05 §12) | Not determined (I-06 §13) | Not determined (I-07 §13) | Not determined; before user space the QEMU guest kernel reports 271,992 kB available of 524,288 kB (I-08 §12a) |
| Flash | Squashfs factory image 82,342,912 bytes; card layout 176 MiB; root 17,645,429 bytes unpacked; boot partition 64 MiB with 20,268,541 bytes in files (I-05 §11) | Root 599,503,592 bytes, of which `/boot/firmware` 97,105,598 (I-06 §11) | Boot content 110,391,389 bytes (modloop 52,080,640); tool closure with kernel and Pi firmware 111,639,452 bytes (I-07 §7, §11) | Card image 6 GiB; root 2,100,686,848 bytes used, `/boot` 302,252,032, ESP 33,748,164 (I-08 §11) |
| Licences of what ships | `LICENCE.broadcom` for the VideoCore firmware; NVRAM GPL-2.0+; six firmware packages without a licence field (I-05 §2a, §5, §5a) | `raspi-firmware`, `firmware-brcm80211` in `non-free-firmware`, not read (I-06 §6, §13) | `raspberrypi-bootloader` `L: custom` (I-07 §6) | `LicenseRef-Fedora-Firmware` (I-08 §6) |
| Defaults NMB must change | Empty root password, dropbear password login, uhttpd on all addresses, IP forwarding on (I-05 §9) | Not determined (C not booted) | Not determined (not booted) | Accounts come from Ignition at first boot (I-08 §6) |
| Support period | Not determined | As amd64; arm64 in Debian LTS not confirmed (I-06 §3) | 2028-06-01 (I-07 §3) | 2027-05-19 (I-08 §3) |
| Relation to the owner's box | Another system on the card | **The owner's Pi 4 runs the Python version on Debian 13 today** (I-06 §15); its own kernel and package versions are not in the records | Another system on the card | Another system on the card |

**OpenWrt build trees for the Pi 4** (I-05 §13 left this comparison to I-09). The bcm27xx/bcm2711 SDK and
ImageBuilder are not in the set (I-05 §13). The Pi 4 image was built with GCC 14.3.0, GNU ld 2.44 and musl 1.2.5,
package architecture `aarch64_cortex-a72` (I-05 §1, §6, §13); the armsr/armv8 SDK and toolchain carry the same
GCC 14.3.0, binutils 2.44 and musl 1.2.5, for package architecture `aarch64_generic` (I-02 §1, §15). Whether an
`nmbd` apk built with the armv8 SDK installs on the Pi 4's feed architecture: not determined. This matters for the
router and OpenWrt packaging, not for the base choice.

## 3. Init and service manager: what each means for `nmbd` and for the 6B.5 sandbox

6B.5 asks for every component: its own unprivileged user; its own mount, PID and IPC namespace (and network
namespace where it does not need the host's); a read-only system with write access only to its own directories;
only the capabilities it needs; `no_new_privs`; a seccomp allow-list; Landlock file rules where the kernel has them;
cgroup limits for memory, CPU and processes; and a stricter profile for third-party engines and upload parsers.
On OpenWrt the profile is built with procd `ujail` and seccomp (6B.5).

What the kernels give, per base (all from section 4 of the records):

| 6B.5 element | OpenWrt (procd) | Debian (systemd 257) | Alpine (OpenRC 0.63.2) | Fedora IoT (systemd 259.5) |
| :--- | :--- | :--- | :--- | :--- |
| Namespaces (mount, PID, IPC, net, user) | `y` (I-01 §4, I-05 §4) | `y` (I-06 §4) | `y`; `/proc/self/ns` shows them (I-07 §4) | `y` (I-08 §4) |
| seccomp | `y`; `procd-seccomp` installed (I-01 §4, §6) | `y`; systemd built `+SECCOMP` (I-06 §4) | `y`; nothing in the base uses it (I-07 §4) | `y`; systemd built `+SECCOMP` (I-08 §4) |
| Landlock | none on any target | `y` both, active on amd64 | x86-64 only | `y` both |
| cgroup limits | cgroup2 mounted, `MEMCG=y`; Pi 4 memory controller may be off (`cgroup_disable=memory`, not determined) | cgroup2 mounted | not mounted until the `cgroups` service runs | controllers `y`, mount not determined |
| A mandatory access control system already on | none | AppArmor active, no profiles package (I-06 §4) | none active (I-07 §4) | SELinux enforcing; how an `nmbd` without its own policy module is confined: not determined (I-08 §4, §12a) |
| Watchdog for the box | `ubus call system watchdog` "offline" in QEMU (no device); on the Pi 4 not determined (I-01 §4, I-05 §4) | systemd `RuntimeWatchdogUSec=0` (hardware watchdog not used by default), `RebootWatchdogUSec=10min` (I-06 §4) | BusyBox `watchdog` init script present, not enabled (I-07 §4) | `WATCHDOG_HANDLE_BOOT_ENABLED=y`; greenboot has a `02_watchdog.sh` required check (I-08 §3, §4) |
| How the Python version runs there | not a Python target | 25 `nmb-*` systemd units (I-06 §15) | five OpenRC services (`nmb-db`, `nmb-netd`, `nmb-probe`, `nmb-telegram`, `nmb-web`) and a crontab (I-07 §15) | `setup.sh` Fedora branch with firewalld ports (I-08 §15) |

Consequences (Inference unless cited):

1. **`nmbd` must run under all three service managers in any case.** The routers keep OpenWrt and procd (6A,
   *Targets*); the owner's packages for Debian, Ubuntu, Fedora/RHEL and Rocky are systemd systems, and the Alpine
   package is an OpenRC system (work plan 7 point 20; init per base from I-06 §6, I-07 §6, I-08 §6). Choosing the
   image's init does not remove the work for the other two; it only decides which one the image is tested on most.
2. **Proposal: the sandbox profile is kept once in `spec/` (6B.8) and applied the same way everywhere.** Where the
   service manager has the means (procd `ujail` and seccomp on OpenWrt, as 6B.5 says), the profile is expressed in
   it; where it has not, `nmbd` applies it itself before it drops privileges. The records show no confinement
   facility of OpenRC in use (I-07 §4, §6), so on Alpine `nmbd`'s own launcher would carry the whole profile. Which
   systemd unit options would express each element was not tested in any record; that is a check for the package
   that builds the profile, against the systemd documentation, not something this file can state.
3. **Landlock must be optional in `nmbd`**, as 6B.5 already says ("where the kernel has them") and I-07 §15
   concludes for the Alpine Pi kernel: it is missing on every OpenWrt target and on Alpine's `linux-rpi`. On a base
   without it, the profile's file rules fall back to the mount namespace with a read-only system (Proposal); the
   status must show the lower level (6B.5, *The user's control*).
4. **cgroup limits need a check at start.** Alpine does not mount cgroup2 by default (I-07 §12a) and the OpenWrt Pi 4
   may run with the memory controller off (I-05 §2). `nmbd` must detect this and either mount and enable what it
   needs or report the missing limit (Proposal; the run-time detection of work plan 5.3 point 2).
5. **systemd on Debian:** one unit for `nmbd` instead of 25 (I-06 §15), masks for `nftables.service`,
   `systemd-networkd` and the package units of hostapd, wpa_supplicant and openvpn (I-06 §9, §15); the hardware
   watchdog is not used by default (I-06 §4), so NMB's watchdog design must switch it on itself (Proposal).
6. **procd on OpenWrt:** the base already runs dnsmasq and ntpd in `ujail` as their own users (I-01 §12), which
   shows the mechanism works on this kernel; whether procd can set the cgroup limits of 6B.5 is not in the records.
7. **OpenRC on Alpine:** one service instead of five (I-07 §15); the `cgroups` service or `nmbd` mounts cgroup2
   (I-07 §12a).

## 4. Proposals

Three proposals that pass criterion 8 as shipped. Fedora IoT is left out (criteria 2 and 8, section 1). A mix of
bases per target (for example Debian on the Pi 4 and OpenWrt or Alpine on x86-64) was considered: no record shows a
gain on one target that would pay for a second image pipeline, a second update scheme and a second set of
coexistence rules, so it is not proposed (Inference).

### Proposal A (recommended): Debian 13 for the Pi 4 and x86-64, systemd, A/B built with a packaged tool

- **What it gives.** One kernel line (6.12.107) for both targets with every NMB feature, the Pi 4 platform and
  Landlock in the config (I-06 §4); every criterion-3 tool from `main` with full iproute2 (I-06 §7); no Python in
  the minimal base (I-06 §8); the longest support of the four (2028-08-09, LTS to 2030-06-30) (I-06 §3); nothing
  listening after boot (I-06 §12); the network managers idle or maskable (I-06 §9); RAUC, SWUpdate,
  `systemd-repart`, `mmdebstrap`, `debos` packaged (I-06 §3).
- **Fit with the owner's decisions.** The owner's Pi 4 runs Debian 13 (I-06 §15). Point 19 migrates it component by
  component on that system, and point 20 needs a Debian package anyway; with a Debian base the image is assembled
  from the same packages the migration ran on, and the last step (card image to NMB's own image) changes the
  update scheme and the read-only layout, not the distribution (Inference).
- **What it costs.**
  - A/B, the read-only root and the image signing (work plan 7 point 6) must be built and tested: the tools are
    packaged but none was tried (I-06 §3, §16).
  - The largest image and memory of the three: root 360 MB (amd64) and 599.5 MB (arm64 with the Pi firmware), 61 MB
    used after boot, against 15 to 20 MB and 22 to 37 MB on OpenWrt and 37.6 MB on Alpine (I-06 §11, §12; I-01 §11,
    §12; I-07 §12). About 58 MB of docs, man pages and locales and the Pi boot partition's second copies of kernel
    and initrd are candidates to cut, not measured (I-06 §2, §11).
  - Units to mask and a release check for Python, because Recommends or a later package can bring it (I-06 §8, §9).
- **Risks.**
  - The arm64 build was never booted; the Pi 4 boot through `raspi-firmware` is read from files only (I-06 §2, §13).
  - Whether Debian LTS covers arm64 is "to be determined" on the Debian wiki (I-06 §3); without it the free security
    support ends 2028-08-09 on the Pi 4.
  - A rebuild is not byte-identical yet (I-06 §13); criterion 4 asks for a reproducible build.
  - `nftables.service` wipes NMB's tables if anything starts it (I-06 §9); the mask must be part of the image and of
    the self-test.

### Proposal B: OpenWrt 25.12 for the Pi 4 and x86-64 (one base for routers and boxes)

- **What it gives.** The smallest system (roots 15 to 20 MB, 22 to 37 MB RAM used) (I-01 §11, §12; I-02 §11, §12;
  I-04 §12); a read-only squashfs root (I-01 §2); procd with `ujail` and seccomp, which 6B.5 names for OpenWrt; the
  same base, scripts and coexistence rules as the routers (6A, *Targets*; I-05 §9 byte-identical scripts); an
  ImageBuilder that assembles an image from published packages plus an `nmbd` apk (I-02 §15, I-04 §15).
- **What it costs.**
  - **No Landlock on any target** (I-01..I-05 §4); the 6B.5 profile is always one level lower than on Debian or
    Fedora.
  - No A/B: one root, overwritten by sysupgrade, and on the Pi 4 no way back after a failed boot (I-01 §3, I-05 §3);
    NMB must build its own two-slot scheme on two different boot paths (GRUB on x86-64, the VideoCore firmware on
    the Pi 4).
  - The two coexistence cautions in NMB's code forever: never the firewall service's `stop`, and addresses on
    netifd's devices vanish on its `restart` (I-01 §9a).
  - BusyBox `ip`, `wpad-basic` and the other base defaults differ from what the packages of point 20 run on; full
    iproute2 and the other tools come from the feeds (I-01 §7, §9a).
  - The defaults that must change (empty root password, password SSH login, LuCI on all addresses) (I-01 §12,
    I-05 §9).
- **Risks.** The support period is not in any record; the Pi 4 image was never booted, so its memory, Wi-Fi,
  coexistence and `cgroup_disable=memory` effect are open (I-05 §2, §12); Ookla and RTR on musl are open (section
  1, criterion 6); the owner's box would get a different system at the end of the migration of point 19, which runs
  on Debian (Inference).

### Proposal C: Alpine 3.24 for the Pi 4 and x86-64, OpenRC, own A/B on the diskless layout

- **What it gives.** A small system (37.6 MB RAM used at login, 27 MB tool closure) (I-07 §7, §12); a read-only base
  by design, with settings in one file written by rename, and commit and restore shown to work (I-07 §3, §12a);
  nothing switched on that fights NMB (I-07 §9); every tool from `main`, no Python anywhere in the closure (I-07 §7,
  §8); the Python version's `setup.sh` already has an Alpine branch (I-07 §15).
- **What it costs.**
  - **No Landlock on the Pi 4** (`linux-rpi`), the first target of 6A (I-07 §4).
  - No A/B tool in the distribution; the kernel and modloop update is an in-place `mv` with no fallback; NMB builds
    its own scheme (two sets of kernel, initramfs and modloop plus a boot switch) (I-07 §3).
  - cgroup2 not mounted by default (I-07 §12a); the modloop signature check only logs a failure and still mounts
    (I-07 §3), so NMB's own signature check must come first.
  - The x86-64 modloop holds about 1.49 GB of firmware that must be narrowed (I-07 §5, §11).
- **Risks.** Support to 2028-06-01, shorter than Debian (I-07 §3); the Pi archive was never booted (I-07 §13); the
  release media lag the repository (kernel 6.18.60 on the media, 6.18.63 in `main` for `lts`) (I-07 §3); Ookla and RTR
  on musl are open (I-07 §13).

### Recommendation

**Proposal A: Debian 13 as the base for the Pi 4 and x86-64, with systemd, and the A/B update built with RAUC first
(SWUpdate as the fallback), decided finally only after the measurements of section 6.** Reasons, each from the
tables above:

1. It is the only candidate that meets criteria 1, 2, 3, 5 and 8 on both targets and has Landlock in the kernel on
   both (section 1). OpenWrt has no Landlock anywhere and no known support period; Alpine has no Landlock on the Pi 4;
   Fedora IoT fails 2 and 8.
2. Its gap is criterion 4, and the parts to close it are packaged in the same release (I-06 §3). OpenWrt and Alpine
   have the same gap (no A/B) and fewer parts to close it (I-01 §3, I-07 §3).
3. The owner's Pi 4 already runs Debian 13 (I-06 §15); the component-by-component migration (point 19) and the
   Debian package (point 20) run on the same userland the image would ship, so the image is the last step of one
   migration, not the start of a second.
4. Its costs are size and memory. 61 MB used after boot left 411 MB available in a 512 MB guest (I-06 §12); how
   that compares with the owner's Pi 4 RAM is not in the records.

The roadmap's earlier assessment named Debian as the recommendation "so far, not a decision" (6A); the records
support it and add two open points it did not have: the arm64 LTS question and the untested A/B tools.

## 5. Decision questions for the owner

Each question has its options, the recommended one first. The answers can be given one by one; question 2 and 3
follow from question 1.

**Question 1. Which base does NMB's own image use for the Raspberry Pi 4 and x86-64?**

- a. **Debian 13 for both targets (recommended).** Meets the criteria except a ready A/B; largest image (section 4,
  Proposal A).
- b. OpenWrt 25.12 for both targets, the same base as the routers. Smallest; no Landlock, no A/B, support period not
  known (Proposal B).
- c. Alpine 3.24 for both targets. Small, read-only by design; no Landlock on the Pi 4, no A/B tool, support to
  2028-06-01 (Proposal C).
- d. Fedora IoT 44. A/B ready-made; ships Python and about 13 months of support per release (criteria 8 and 2 not
  met). Not recommended.
- e. A different base per target. Not recommended (section 4).

**Question 2. Which init system and service manager does the image use?**

- a. **systemd, as Debian ships it (recommended with 1a).** It is also the service manager of the Debian, Ubuntu,
  Fedora/RHEL and Rocky packages (point 20).
- b. OpenRC (only with 1c).
- c. procd (only with 1b).

  With every answer, `nmbd` still supports all three (routers on procd, the packages on systemd and OpenRC) and the
  6B.5 sandbox stays one profile applied the same way on each (section 3). Something smaller than systemd on Debian
  (6A, *To decide*) was not examined in any record and is not offered.

**Question 3. How is the A/B update built?** (for 1a; with 1b or 1c NMB builds its own scheme, with 1d it is
rpm-ostree with greenboot)

- a. **RAUC (trixie `rauc` 1.13) as the first candidate, tested before it is fixed (recommended).** Packaged in the
  same release (I-06 §3); named first in 6A, *To decide*. The records hold no technical comparison of RAUC and
  SWUpdate, so this recommendation fixes the order of testing, not the final tool.
- b. SWUpdate (trixie 2024.12.1), the fallback if RAUC fails the tests of section 6.
- c. `systemd-repart` with `systemd-boot` (257.13), from the same systemd. Packaged, untested (I-06 §3).
- d. ostree on Debian (trixie 2025.2), the deployment model Fedora IoT shows (I-08 §3). Packaged, untested (I-06 §3).

  With every answer: the update is accepted only with NMB's own valid signature (work plan 7 point 6); whether each
  tool can check the proposed Ed25519 signature itself, or NMB checks it before handing the bundle over, is not
  determined.

## 6. What must still be measured before the decision is final

Each item names the record that left it open. Items 1 to 4 decide between the proposals; the rest complete the
chosen base. Every Pi 4 run is under work plan 7 point 5 (full backup of settings, database and the card image
first, a tried way back, the time recorded as the box's own action, the Python version put back afterwards).

1. **A real Pi 4 boot of the chosen base**, recording what the x86-64 records have: `/proc/cmdline`, boot time,
   memory and storage after boot, listening ports, running services, mounted cgroup controllers, the LSM list
   (Landlock), `/dev/watchdog`, Wi-Fi (`iw list`, firmware and NVRAM chosen, regulatory domain), GENET link and MAC,
   and the coexistence tests of I-01 §9. Open for Debian build C (I-06 §13), OpenWrt (I-05 §12, every field listed
   there), Alpine (I-07 §13) and Fedora IoT (I-08 §13).
2. **An A/B test on the chosen base, on the Pi 4 and in QEMU** (the 6A testing plan): image built, update written to
   the other slot, forced rollback, power cut during an update, full data partition; the slot switch on the Pi
   firmware boot path; signature check with the key scheme of work plan 7 point 6. Open for every base except Fedora
   IoT, whose mechanism is built in but untested (I-06 §3, I-07 §3, I-01 §3, I-05 §3, I-08 §3).
3. **Ookla and RTR binaries on the chosen C library** (criterion 6): not determined in any record (I-01 to I-08;
   I-06 §13, I-07 §13, I-08 §13 left it to I-09, and no record holds the binaries).
4. **Support periods not in the records:** OpenWrt 25.12 (no record), and whether Debian LTS covers arm64 (I-06 §3).
5. **Landlock live on the Debian arm64 kernel** (config `y`, seen active on amd64 only) (I-06 §4, §13).
6. **Debian image hygiene:** the transitive Recommends closure for Python (I-06 §13), the release check on the final
   image (I-06 §8), byte-for-byte reproducibility (I-06 §13), a read-only root and the size after cutting docs,
   locales and the second kernel copies on `/boot/firmware` (I-06 §2, §11).
7. **Runs of the criterion-3 tools** (dnsmasq, hostapd, wpa_supplicant, pppd with PPPoE, WireGuard and OpenVPN
   tunnels) on the chosen base: only presence and version were checked (I-06 §12, I-07 §12); on OpenWrt only a
   WireGuard device was created and two `mac80211_hwsim` radios were listed (I-02 §9a).
8. **Boot time and memory on real hardware or KVM** for the chosen base on x86-64: Debian and OpenWrt x86-64 ran under
   TCG only (I-06 §13, I-01 §13).
9. **Firmware licences for the 6B.3A licence list:** Debian `non-free-firmware` (I-06 §13), OpenWrt's empty licence
   fields and `LICENCE.broadcom` (I-01 §5, I-05 §2a, §5), Alpine `custom` (I-07 §13), Fedora (I-08 §13).
10. **The settings move from a `setup.sh` box** to the image (6A, *To decide*): not covered by any record.
11. **Only if Fedora IoT or OpenWrt stay in the choice:** a Fedora IoT boot that reaches `multi-user.target` on an
    idle PC or the Pi 4, and an x86_64 install in a guest with more than 512 MB (I-08 §12a, §12b, §13), plus
    whether an own compose without `firewalld`, `policycoreutils-python-utils` and `setools-console` is free of
    Python (I-08 §13); the OpenWrt Pi 4 `cgroup_disable=memory` effect (I-05 §2) and whether an apk built with the
    armsr/armv8 SDK installs on the Pi 4's `aarch64_cortex-a72` feeds (section 2.2).
