# Image record: OpenWrt 25.12.6 x86-64 (squashfs and ext4, combined-efi)

Package I-01 (work plan section 5.4). The method is work plan 5.1, and every field of 5.2 is below.
Each value names the command that produced it. A value that could not be read says "not determined" and why.
Nothing here comes from memory. Only lists, versions and checksums are recorded. No file taken out of the image
is in this repository.

- Date of the analysis: 2026-11-11.
- Where it ran: dev PC, Windows with WSL2 Ubuntu 24.04, about 1.8 GB RAM and 4 threads, no root. The QEMU guest ran
  on the same PC.
- Work folder: `~/nmb-work/I-01/` (outside every repository; one fresh sub-folder per unpack).
- Material, read only: `~/src/NMB-depythonization-firmwares/platforms/x86-64/`.

## 1. Files and checksums

| File | Size (bytes) | SHA-256 (publisher list) | Check |
| :--- | ---: | :--- | :--- |
| `openwrt-25.12.6-x86-64-generic-squashfs-combined-efi.img.gz` | 12,498,825 | `61a550501581023f5f27817a5ec36edeb574e549483c47720e41d366be0bbf05` | OK |
| `openwrt-25.12.6-x86-64-generic-ext4-combined-efi.img.gz` | 13,813,344 | `1ef06b324d04490a90de26d6f2a292111854f05e3c946648fad70c339b52f739` | OK |

- **How it was checked.** The set's README states that the files were checked against the publisher's checksums, but
  it does not list the x86-64 OpenWrt sums. The x86/64 tree is not in the set either: `openwrt-target-trees/` holds
  only `x86/generic`, `armsr/armv8` and `armsr/armv7`. So the publisher's list was fetched on 2026-11-11 at
  01:34 UTC from `https://downloads.openwrt.org/releases/25.12.6/targets/x86/64/sha256sums`, together with
  `sha256sums.asc`. The list itself has 1,303 lines and SHA-256 `5515522e025c507802d03d25853c753e278d877dac04f64ef40578112f1d0319`.
  - `grep -E "x86-64-generic-(squashfs|ext4)-combined-efi.img.gz" sha256sums | sha256sum -c` printed both files
    `OK` (exit 0).
- **Signature of the list.** It was verified in an isolated `GNUPGHOME` with `gpg --verify sha256sums.asc sha256sums`.
  The result was "Good signature from OpenWrt Build System (Nitrokey3)" (exit 0).
  - Signing subkey: `92C5 61DE 55AE 6552 F3C7 36B8 2B01 5109 0606 D1D9`. Primary key:
    `8A8B C12F 46B8 36C0 F9CD B36F 1D53 D187 7742 E911`. The signature was made on 2026-10-18.
  - The key came from keys.openpgp.org. Its fingerprints were compared with OpenWrt's keyring repository
    (`github.com/openwrt/keyring`, file `gpg/0x1D53D1877742E911.asc`) using `gpg --show-keys`, and they are the same.
    gpg says "not certified with a trusted signature", because no web of trust was set up.
- **Release identity.** The release is `r33311-981c978eff`. It comes from `version.buildinfo`, which was fetched from
  the same place, and the image's own `/etc/openwrt_release` gives the same value.
- **Decompressed images.** `gzip -dc` gave exit 0 for both images and printed no "trailing garbage" warning. Each
  decompressed image is 126,123,520 bytes.
  - squashfs: `abc99c60179a289f081bbbc1d8e84c046cf7b298e24f5c82cbdde6be708f6430`
  - ext4: `86f040223f2e6a2a134189862b71351114d426ae0cdbe904347f337e7f8f2dca`

## 2. Partition and file system layout, boot method

The layout was read with `sgdisk -p disk.img` (exit 0), `sgdisk -i N`, `fdisk -l` (exit 0) and `fdisk -t dos -l`.
It is the same in both images.

| No. | Sectors (512 B) | Size | GPT type | Contents (`file`) |
| :--- | :--- | :--- | :--- | :--- |
| 128 | 34–511 | 239 KiB | BIOS boot (`21686148-…`, name "Hah!IdontNeedEFI") | GRUB core for legacy BIOS boot |
| 1 | 512–33279 | 16.0 MiB | Linux filesystem (`0FC63DAF-…`), **not** the EFI System type | FAT16, label `kernel` |
| 2 | 33280–246271 | 104.0 MiB | Linux filesystem | squashfs image: squashfs 4.0, xz, 6,096,242 bytes, 1,291 inodes. ext4 image: ext4, label `rootfs`, 26,624 blocks of 4 KiB, 19,361 blocks free (`debugfs -R stats`) |

- **Partition table.** It is a GPT behind a protective MBR (one `ee` entry). The backup GPT header is missing from the
  image: the bytes at the alternate LBA 246304 are zero.
  - sgdisk reports this as "invalid backup GPT header".
  - The guest kernel logs `GPT:Alternate GPT header not at the end of the disk` and continues.
- **Boot partition.** Its contents were read with `mdir -/ -i p1.img ::` and `mcopy -s` (both exit 0). It holds
  `/boot/vmlinuz` (5,706,752 bytes), `/boot/grub/{boot.img,core.img,grub.cfg}` and `/efi/boot/bootx64.efi`
  (638,976 bytes, GRUB 2.12).
  - Every boot file is identical in the two images (`sha256sum`).
  - `vmlinuz` has SHA-256 `8f70595f…7ead4`. This equals the publisher's `openwrt-25.12.6-x86-64-generic-kernel.bin`.
- **Boot method.**
  - UEFI: OVMF loads `\efi\boot\bootx64.efi` (GRUB 2.12) from partition 1, even though the GPT type is not ESP. The
    boot log shows `BdsDxe: starting Boot0001 "UEFI Misc Device"`.
  - Legacy BIOS: through the BIOS boot partition.
  - GRUB prints `error: can't find command 'search'` twice and then boots the default entry after the 5 s timeout.
  - Kernel command line, from the guest's `cat /proc/cmdline`: `root=PARTUUID=<disk GUID>-02 rootwait
    console=tty1 console=ttyS0,115200n8 noinitrd`. There is no initramfs.
  - A second GRUB entry boots with `failsafe=true`.
- **Root file system at run time (squashfs image).** It comes from the guest's `mount` and `df -k`.
  - The squashfs is mounted read-only on `/rom`.
  - The rest of partition 2 is a loop device (`/dev/loop0`, 200,960 sectors), formatted ext4 on the first boot. The
    log shows `mount_root: overlay filesystem in /dev/loop0 has not been formatted yet` and then
    `switching to ext4 overlay`.
  - The overlay is mounted on `/overlay`, and `overlayfs:/overlay` is `/`.
  - Partition 1 is mounted read-write on `/boot` (vfat).
  - `/tmp` is tmpfs.
- **Root file system (ext4 image).** It was not booted, only taken apart. The root is the ext4 partition itself, with
  no overlay.
  - It was unpacked with `debugfs -R "rdump / out" p2.img` (exit 0; three warnings about changing ownership, because
    there is no root).
  - Its file list is identical to the squashfs root (`find -type f | sort`, then `diff`, exit 0), and so is its package
    list (`diff`, exit 0).
- **Unpacking the squashfs.** `unsquashfs -no-xattrs -d … p2.img` gave exit 2. The one failure was
  `could not create character device …/dev/console, because you're not superuser`, which is expected without root.
  It created 940 files, 121 directories and 229 symlinks.

## 3. How an update is written

- **sysupgrade.** It is present: `/sbin/sysupgrade` (guest: `command -v sysupgrade`). The x86 rule is
  `/lib/upgrade/platform.sh`, read in the unpacked root.
  - `platform_check_image` accepts the magic `eb48`/`eb63`, which is the boot sector of a combined image. It compares
    the image's partition table with the disk. If the layout changed, it writes the full image.
  - `platform_do_upgrade` writes the disk.
  - `platform_copy_config` copies the saved settings into partition 1.
  - `platform_do_bootloader_upgrade` reruns `grub-bios-setup`.
- **No A/B scheme.** There is one root partition. An update overwrites it in place, and the settings are carried
  across in the backup file.
- **Other update helpers.** `owut` and `attendedsysupgrade-common` are installed (the package list). They are online
  image-rebuild helpers.
- **Packages at run time.** `apk` installs into the overlay.

## 4. Kernel and the configuration NMB needs

- **Kernel version: 6.12.112.** Sources:
  - `file vmlinuz`
  - `strings` of the decompressed kernel: `Linux version 6.12.112 (builder@buildhost) (x86_64-openwrt-linux-musl-gcc
    (OpenWrt GCC 14.3.0 r33311-981c978eff) 14.3.0, GNU ld (GNU Binutils) 2.44) #0 SMP Sat Oct 17 21:45:38 2026`
  - guest `uname -a`
  - guest `ubus call system board`, which gives `"kernel": "6.12.112"`
  - The package `kernel 6.12.112~b4125196ecec75391093e1e2edeba490-r1` is installed.
- **Where the config came from: not in the image.**
  - The kernel payload was decompressed (xz stream at offset 17092, `tail -c +17093 vmlinuz | xz -dc --single-stream`,
    exit 0). It contains no `IKCFG_ST` marker (`grep -c -a IKCFG_ST` printed 0).
  - In the guest, `/proc/config.gz` does not exist.
  - The image carries no `.config`.
- **Source used instead: the official `kmod-ikconfig` module for this exact kernel build.** The config says
  `CONFIG_IKCONFIG=m`, so the module carries the configuration.
  - Address: `https://downloads.openwrt.org/releases/25.12.6/targets/x86/64/kmods/6.12.112-1-b4125196ecec75391093e1e2edeba490/kmod-ikconfig-6.12.112-r1.apk`,
    fetched 2026-11-11.
  - SHA-256 `2c87ccb2…c6cb`. It was checked against the signed `sha256sums` above (`sha256sum -c`, exit 0).
  - The kmods folder hash is the same as the hash in the installed `kernel` package's version.
  - The apk (ADB format, `ADBd` = raw deflate) was inflated with a 10-line Python script run with `-I`, outside the
    data folder. The gzip blob between `IKCFG_ST` and `IKCFG_ED` was then decompressed (exit 0).
  - Result: 8,537 lines, starting `# Linux/x86 6.12.112 Kernel Configuration` and
    `CONFIG_CC_VERSION_TEXT="x86_64-openwrt-linux-musl-gcc (OpenWrt GCC 14.3.0 r33311-981c978eff) 14.3.0"`.
    SHA-256 of the extracted config: `c936e1f6818ce8c5729a4c1dd5e77860214548fc44d6b5fe10cf107b4e3ddf3a`.
  - **Source of every option below: this config.** The "in image" column says whether the module is in the image
    (`ls lib/modules/6.12.112/`). The "package" column says whether a kmod package is published in the kmods folder
    (the file names in the signed `sha256sums`). That folder holds 1,184 kmod packages.

| Feature (5.2) | Option(s) | Value | In image | Package |
| :--- | :--- | :--- | :--- | :--- |
| Bridge, VLAN filtering | `BRIDGE`, `BRIDGE_VLAN_FILTERING` | y, y | built in | — |
| Bonding | `BONDING` | m | no | `kmod-bonding` |
| Team | `NET_TEAM` | m | no | `kmod-team` and its modes |
| VLAN 802.1Q | `VLAN_8021Q` | y | built in | — |
| nftables | `NF_TABLES` m, `NF_TABLES_INET` y, `NF_TABLES_NETDEV` y, `NF_TABLES_BRIDGE` m, `NFT_NAT` m, `NFT_MASQ` m, `NFT_FIB_INET` m, `NF_FLOW_TABLE` m | as listed | yes (`nf_tables.ko`, `nft_*.ko`, 71 `.ko` files in all; guest `lsmod \| wc -l` = 71 lines) | `kmod-nft-*` |
| WireGuard | `WIREGUARD` | m | no | `kmod-wireguard` |
| PPP, PPPoE | `PPP` m, `PPPOE` m, `PPP_ASYNC` m | m | yes (`ppp_generic`, `pppoe`, `pppox`, `ppp_async`) | `kmod-ppp`, `kmod-pppoe` |
| cfg80211, mac80211 | `CFG80211` | **not set in the kernel** | no | `kmod-cfg80211`, `kmod-mac80211` built out of tree from backports `6.18.55` (versions `6.12.112.6.18.55-r2`) |
| mac80211_hwsim | — | from backports | no | `kmod-mac80211-hwsim-6.12.112.6.18.55-r2` |
| Wi-Fi drivers (examples) | — | from backports / mt76 | no | `kmod-ath9k`, `kmod-ath10k(-ct)`, `kmod-iwlwifi`, `kmod-mt76*`, `kmod-mt7601u` |
| USB host | `USB` y, `USB_XHCI_HCD` y, `USB_EHCI_HCD` y | built in | — | `kmod-usb3`, `kmod-usb2` also published |
| USB network | `USB_USBNET` m, `USB_RTL8152` m, `USB_NET_AX88179_178A` m, `USB_NET_SMSC95XX` m, `USB_NET_AQC111` m | m | no | `kmod-usb-net-*` |
| USB tethering (4.27) | `USB_NET_RNDIS_HOST`, `USB_NET_CDCETHER`, `USB_NET_CDC_NCM`, `USB_IPHETH`, `USB_NET_QMI_WWAN`, `USB_NET_CDC_MBIM`, `USB_NET_HUAWEI_CDC_NCM` | all m | no | `kmod-usb-net-rndis`, `-cdc-ether`, `-cdc-ncm`, `-ipheth`, `-qmi-wwan`, `-cdc-mbim` |
| TUN, veth, macvlan, dummy, IFB | `TUN`, `VETH`, `MACVLAN`, `DUMMY`, `IFB` | all m | no | `kmod-tun`, `kmod-veth`, `kmod-macvlan`, `kmod-dummy` |
| Policy routing | `IP_ADVANCED_ROUTER`, `IP_MULTIPLE_TABLES`, `IPV6_MULTIPLE_TABLES` | y | built in | — |
| Traffic control | `NET_SCHED` y, `NET_SCH_FQ_CODEL` y, `NET_SCH_CAKE` m, `NET_SCH_HTB` m, `NET_SCH_NETEM` m, `NET_SCH_TBF` m | as listed | fq_codel built in | `kmod-sched-core`, `kmod-sched-cake`, `kmod-netem` |
| Namespaces | `NAMESPACES`, `NET_NS`, `PID_NS`, `USER_NS`, `UTS_NS`, `IPC_NS` | all y | built in | — |
| seccomp | `SECCOMP`, `SECCOMP_FILTER` | y, y | built in | (`procd-seccomp` installed) |
| **Landlock** | `SECURITY` | **not set**, so `SECURITY_LANDLOCK` is absent | — | not available without an own kernel |
| cgroups | `CGROUPS` y, `MEMCG` y (`MEMCG_V1` not set), `CGROUP_PIDS` y, `CPUSETS` y, `CGROUP_BPF` y, `BLK_CGROUP` y; `CGROUP_FREEZER`, `CGROUP_DEVICE` not set | as listed | cgroup2 mounted on `/sys/fs/cgroup` (guest `mount`) | — |
| Watchdog | `WATCHDOG` y; `SOFT_WATCHDOG`, `I6300ESB_WDT`, `ITCO_WDT` m | as listed | no `/dev/watchdog*` in the guest; `ubus call system watchdog` gives `"status": "offline", "timeout": 30` | `kmod-softdog`, `kmod-i6300esb-wdt`, `kmod-itco-wdt` |
| Module signing | `MODULE_SIG` | not set | — | — |
| Other | `BPF_SYSCALL` y, `XDP_SOCKETS` not set, `TCP_CONG_BBR` m, `IPV6` y, `PREEMPT_NONE` y, `HZ=100`, `NR_CPUS=512` | as listed | — | — |

## 5. Modules and firmware files carried

- **Modules.** There are 71 `.ko` files in `lib/modules/6.12.112/` (`ls *.ko | wc -l`), plus `modules.builtin`
  (305 lines).
  - Groups: NIC drivers (`e1000`, `e1000e`, `igb`, `igc`, `ixgbe`, `r8169`, `tg3`, `bnx2`, `forcedeth`, `amd-xgbe`,
    `ena`, `stmmac`/`dwmac-intel`), nftables and conntrack, PPP/PPPoE, `i915` with DRM helpers, and `vfat`/NLS.
  - virtio-net and virtio-blk are built in (`VIRTIO_NET=y`, `VIRTIO_BLK=y`).
- **Firmware.** There are 70 files under `lib/firmware`, 1,411,732 bytes in all (`find -type f`, `du -sb`): `bnx2/`
  (5), `i915/` DMC (37) and `rtl_nic/` (28). No Wi-Fi firmware is in the image.
  - The three firmware packages (`bnx2-firmware`, `i915-firmware-dmc`, `r8169-firmware`) have an empty licence field
    in the apk database.

## 6. Init system, service manager, C library, package manager

- **Init:** procd (`/sbin/procd` is PID 1 in the guest's `ps w`) with ubus, ujail and seccomp (`procd-ujail`,
  `procd-seccomp`). Init scripts are in `/etc/init.d` and are linked in `/etc/rc.d`.
- **Enabled services** (guest loop `$s enabled`): boot, cron, dnsmasq, done, dropbear, firewall, gpio_switch, led, log,
  network, odhcpd, packet_steering, rpcd, sysctl, sysfixtime, sysntpd, system, ucitrack, uhttpd, umount,
  urandom_seed, urngd (22 of 22).
- **Services known to procd** (guest `ubus call service list`): cron, dnsmasq, dropbear, firewall, gpio_switch, log,
  network, odhcpd, packet_steering, rpcd, sysntpd, system, ubus, ucitrack, uhttpd, urandom_seed, urngd.
- **C library: musl 1.2.5.** The guest's `/lib/libc.so` prints `musl libc (x86_64)` / `Version 1.2.5`.
  - The package is `libc 1.2.5-r5`.
  - `/lib/ld-musl-x86_64.so.1` is a link to `libc.so`.
  - The toolchain is GCC 14.3.0, from the kernel's version string.
- **Package manager: apk-tools 3.0.5** (guest `apk --version`: `apk-tools 3.0.5, compiled for x86_64.`). opkg is
  absent (guest `command -v opkg`).
  - Feeds are in `/etc/apk/repositories.d/distfeeds.list`: target, kmods, base, luci, packages, routing, telephony and
    video for 25.12.6.
- **Installed packages: 182.** The squashfs and ext4 images give the same count and the same list. Sources:
  - `awk '/^P:/…' lib/apk/db/installed`
  - guest `apk info | wc -l` = 182
  - The list is in section 10.
- **Licences of the installed packages** (`L:` fields): GPL-2.0(-only/-or-later) for most of them, then ISC, Apache-2.0
  (LuCI), LGPL-2.1, MIT, GPL-3.0-with-GCC-exception (libc, libgcc), GPL-3.0-or-later (GRUB), BSD and Zlib. Three
  firmware packages have no licence field.

## 7. Tools of 6A criterion 3, with versions

The presence check is the guest's `command -v` loop. Versions come from the guest's own `--version` output where it
exists, otherwise from the installed package's version.

| Tool | Present | Version |
| :--- | :--- | :--- |
| procd (init) | yes, `/sbin/procd` | `procd 2026.03.13~58eb263d-r1` (package) |
| systemd, OpenRC | no | — |
| nftables `nft` | yes, `/usr/sbin/nft` | `nftables v1.1.6 (Commodore Bullmoose #7)`; package `nftables-json 1.1.6-r2` |
| dnsmasq | yes | `Dnsmasq version 2.93`; package `dnsmasq 2.93-r1` (not `dnsmasq-full`) |
| hostapd | **no** | (feeds: see section 9) |
| wpa_supplicant | **no** | (feeds: see section 9) |
| iw | **no** | (feeds: see section 9) |
| pppd | yes | `pppd version 2.5.2`; `ppp 2.5.2-r4`, `ppp-mod-pppoe 2.5.2-r4` |
| WireGuard (`wg`, module) | **no** | (feeds: see section 9) |
| OpenVPN | **no** | (feeds: see section 9) |
| `ip` | yes, but the **BusyBox applet** (`/sbin/ip -> ../bin/busybox`) | BusyBox v1.37.0. It has no `-br` and no `-V` (guest: `ip -br addr` printed the BusyBox usage and exit 1). `ip link add … type bridge` worked |
| `tc`, `ethtool`, `iptables`, `ss` | no | — |
| DHCP client | `udhcpc` (BusyBox), `odhcp6c 2026.06.27~10a52220-r3` | — |
| DHCP/RA server | `odhcpd-ipv6only 2026.06.29~5d7be43f-r1` and dnsmasq | — |
| SSH | `dropbear` | `Dropbear v2025.89` |
| Web | `uhttpd 2026.08.24~373145f7-r1` with LuCI `26.288.56989~75a306d` | — |
| HTTP client | `uclient-fetch` (`wget` is a link to it) | `uclient-fetch 2026.06.07~daad21fa-r1` |
| BusyBox | yes | `busybox 1.37.0-r6` (the banner reads `BusyBox v1.37.0 (2026-10-17 21:45:38 UTC)`) |

## 8. Python (6A criterion 8)

**No Python is in the image.** The checks:

- In the unpacked squashfs root, `find . -iname "python*" -o -name "*.py" -o -name "*.pyc"` found 0 files.
- `grep -rlI "^#!.*python" .` found 0 files.
- In the guest, `find / -xdev …` and `find /rom -xdev …` both found 0.
- In the guest, `command -v python python3 perl` printed absent for all three.
- The base's own tools are shell, ucode and C: `fw4` is ucode, and netifd and procd are C.
- `python3` is available in the packages feed but is not installed (section 9).

## 9. Own network and firewall manager, and how each is switched off (6A criterion 5)

The guest runs **netifd** (`/sbin/netifd`, package `netifd 2026.02.26~cbb83a18-r2`) and **firewall4**
(`/sbin/fw4`, ucode, package `firewall4 2025.03.17~b6e51575-r2`). The firewall's ruleset is one table,
`table inet fw4` (guest `nft list tables`), with 145 lines (`nft list ruleset | wc -l`).

**Default configuration.** It comes from the guest's `uci show network`, plus `/etc/board.d/99-default_network` and
`bin/config_generate`.

- `eth0` is bridged into `br-lan`, which gets the static built-in LAN default address.
- `eth1` is `wan` with DHCP, and `wan6` with DHCPv6.
- A ULA prefix is generated on the first boot.
- dnsmasq serves DHCP and DNS on the LAN.

These are the coexistence tests run in the guest. The exit codes are those of the probe line.

| Test | Command | Seen |
| :--- | :--- | :--- |
| Own nft table across a firewall restart | `nft add table inet nmbtest` + chain; `/etc/init.d/firewall restart`; `nft list tables`; `fw4 reload`; `nft list tables` | `inet nmbtest` **kept** after both (exit 0) |
| Firewall service stop | `/etc/init.d/firewall stop`; `nft list tables` | **all tables gone, NMB's included** (empty list). `stop_service` calls `fw4 flush`, which deletes every table (`nft list tables \| while read …; nft delete table`, read in `/sbin/fw4`). Plain `fw4 stop` deletes only `inet fw4` (same file) |
| Firewall disable | `/etc/init.d/firewall disable`; `… enabled; echo $?`; `ls /etc/rc.d \| grep -c firewall` | `enabled` gave 1 and 0 links were left, so it is off at boot. `enable` + `start` brought back `inet fw4` |
| netifd stop | `/etc/init.d/network stop`; `pgrep -l netifd` | no netifd process (pgrep rc=1) |
| netifd disable | `/etc/init.d/network disable`; `ls /etc/rc.d \| grep -c network` | 0 links, so it is off at boot. `enable` + `start` brought the network back |
| Own bridge across netifd reload/restart/stop/start | `ip link add nmbbr0 type bridge`, address from 198.51.100.0/24; then `network reload`, `restart`, `stop`, `start` | after all of them, `ip route` still showed `198.51.100.0/24 dev nmbbr0` (run 3, probe 25). netifd leaves a device it does not manage alone |
| Own address on netifd's `br-lan` | `ip addr add 203.0.113.9/24 dev br-lan`, then reload and restart | run 4 (9a): **kept** across `network reload`, **removed** by `network restart`. In run 3 the read-out failed, because BusyBox `ip` has no `-br` |

**What this means for criterion 5.** Both managers can be switched off cleanly with `disable`, plus `stop` for netifd.
Two cautions:

- The firewall's `stop` action flushes the **whole** nftables ruleset, so NMB must not run it while its own tables
  are loaded. Use `fw4 stop`, or `disable` followed by a reboot.
- netifd owns the devices that are in its configuration.

### 9a. Feed lookup and repeated `br-lan` test (run 4)

Run 4 was the last boot. All 22 probes ended with exit 0. The feeds were read in the guest over user-mode networking:
`apk update` gave rc 0 (in run 3 it also gave `OK: 11234 distinct packages available`). Then, for each package
name, `apk search $p | grep "^$p-[0-9]"` gave these versions. **None of these packages is installed in the image.**

| Package | Version in the 25.12.6 feeds |
| :--- | :--- |
| `hostapd-mbedtls` / `wpad-basic-mbedtls` / `wpa-supplicant-mbedtls` | `2025.08.26~ca266cc2-r7` |
| `iw` / `iw-full` | `6.17-r1` |
| `wireguard-tools` | `1.0.20260223-r1` |
| `kmod-wireguard` | `6.12.112-r1` |
| `openvpn-mbedtls` | `2.7.8-r1` |
| `ip-full` / `tc-full` (iproute2) | `6.18.0-r2` |
| `ethtool` | `6.15-r1` |
| `iperf3` | `3.22-r1` |
| `kmod-mac80211-hwsim` | `6.12.112.6.18.55-r2` |
| `kmod-veth`, `kmod-bonding`, `kmod-usb-net-rndis`, `kmod-ikconfig` | `6.12.112-r1` |
| `python3` / `python3-base` | `3.13.18-r1` (in the feed only; never installed for NMB, criterion 8) |

The `br-lan` and firewall tests were repeated with BusyBox `ip addr show`. Each step gave exit 0.

| Step | `nmbbr0` (own bridge, 198.51.100.77/24) | 203.0.113.9/24 added on netifd's `br-lan` |
| :--- | :--- | :--- |
| after adding | present | present |
| `/etc/init.d/network reload` + 10 s | present | **present** |
| `/etc/init.d/network restart` + 15 s | present | **gone** (netifd rebuilds `br-lan`) |

| Step | `nft list tables` afterwards |
| :--- | :--- |
| `nft add table inet nmbtest` | `inet fw4`, `inet nmbtest` |
| `fw4 stop` | `inet nmbtest` only: plain `fw4 stop` removes only its own table |
| `fw4 start` | `inet nmbtest`, `inet fw4` |
| `/etc/init.d/firewall stop` | **empty**: the service stop runs `fw4 flush` and deletes NMB's table too |
| `/etc/init.d/firewall start` | `inet fw4` |

What follows for NMB:

- netifd leaves devices it does not manage alone, and keeps foreign addresses on its own devices across a `reload`.
  It drops them when it rebuilds the device on `restart`.
- `fw4` keeps foreign tables across `start`, `stop`, `restart` and `reload`. The init script's `stop` flushes
  everything.

## 10. Installed packages (182, `lib/apk/db/installed`; the same in both images)

`apk-mbedtls 3.0.5-r3`, `attendedsysupgrade-common 10`, `base-files 1718~981c978eff`, `bnx2-firmware 20260221-r1`, `busybox 1.37.0-r6`, `ca-bundle 20260816-r1`, `cgi-io 2026.07.21~31cb3c89-r1`, `dnsmasq 2.93-r1`, `dropbear 2025.89-r3`, `e2fsprogs 1.47.3-r1`, `firewall4 2025.03.17~b6e51575-r2`, `fstools 2026.05.23~16718b6e-r1`, `fwtool 2025.10.03~04cd252e-r1`, `getrandom 2025.10.30~6f78fa49-r1`, `grub2 2.12-r1`, `grub2-bios-setup 2.12-r1`, `grub2-efi 2.12-r1`, `i915-firmware-dmc 20260221-r1`, `jansson4 2.14.1-r1`, `jshn 2026.06.19~7dd12784-r1`, `jsonfilter 2026.03.16~b9034210-r1`, `kernel 6.12.112~b4125196ecec75391093e1e2edeba490-r1`, `kmod-acpi-video 6.12.112-r1`, `kmod-amazon-ena 6.12.112-r1`, `kmod-amd-xgbe 6.12.112-r1`, `kmod-backlight 6.12.112-r1`, `kmod-bnx2 6.12.112-r1`, `kmod-button-hotplug 6.12.112-r3`, `kmod-crypto-crc32c 6.12.112-r1`, `kmod-crypto-hash 6.12.112-r1`, `kmod-dma-buf 6.12.112-r1`, `kmod-drm 6.12.112-r1`, `kmod-drm-buddy 6.12.112-r1`, `kmod-drm-display-helper 6.12.112-r1`, `kmod-drm-exec 6.12.112-r1`, `kmod-drm-i915 6.12.112-r1`, `kmod-drm-kms-helper 6.12.112-r1`, `kmod-drm-suballoc-helper 6.12.112-r1`, `kmod-drm-ttm 6.12.112-r1`, `kmod-drm-ttm-helper 6.12.112-r1`, `kmod-dwmac-intel 6.12.112-r1`, `kmod-e1000 6.12.112-r1`, `kmod-e1000e 6.12.112-r1`, `kmod-fb 6.12.112-r1`, `kmod-fb-cfb-copyarea 6.12.112-r1`, `kmod-fb-cfb-fillrect 6.12.112-r1`, `kmod-fb-cfb-imgblt 6.12.112-r1`, `kmod-fb-sys-fops 6.12.112-r1`, `kmod-fb-sys-ram 6.12.112-r1`, `kmod-forcedeth 6.12.112-r1`, `kmod-fs-vfat 6.12.112-r1`, `kmod-hwmon-core 6.12.112-r1`, `kmod-i2c-algo-bit 6.12.112-r1`, `kmod-i2c-core 6.12.112-r1`, `kmod-igb 6.12.112-r1`, `kmod-igc 6.12.112-r1`, `kmod-input-core 6.12.112-r1`, `kmod-ixgbe 6.12.112-r1`, `kmod-lib-crc-ccitt 6.12.112-r1`, `kmod-lib-crc32c 6.12.112-r1`, `kmod-libphy 6.12.112-r1`, `kmod-mdio 6.12.112-r1`, `kmod-mdio-devres 6.12.112-r1`, `kmod-mii 6.12.112-r1`, `kmod-nf-conntrack 6.12.112-r1`, `kmod-nf-conntrack6 6.12.112-r1`, `kmod-nf-flow 6.12.112-r1`, `kmod-nf-log 6.12.112-r1`, `kmod-nf-log6 6.12.112-r1`, `kmod-nf-nat 6.12.112-r1`, `kmod-nf-reject 6.12.112-r1`, `kmod-nf-reject6 6.12.112-r1`, `kmod-nfnetlink 6.12.112-r1`, `kmod-nft-core 6.12.112-r1`, `kmod-nft-fib 6.12.112-r1`, `kmod-nft-nat 6.12.112-r1`, `kmod-nft-offload 6.12.112-r1`, `kmod-nls-base 6.12.112-r1`, `kmod-nls-cp437 6.12.112-r1`, `kmod-nls-iso8859-1 6.12.112-r1`, `kmod-nls-utf8 6.12.112-r1`, `kmod-pcs-xpcs 6.12.112-r1`, `kmod-phy-realtek 6.12.112-r1`, `kmod-phylink 6.12.112-r1`, `kmod-ppp 6.12.112-r1`, `kmod-pppoe 6.12.112-r1`, `kmod-pppox 6.12.112-r1`, `kmod-pps 6.12.112-r1`, `kmod-ptp 6.12.112-r1`, `kmod-r8169 6.12.112-r1`, `kmod-slhc 6.12.112-r1`, `kmod-stmmac-core 6.12.112-r1`, `kmod-tg3 6.12.112-r1`, `libblkid1 2.41.5-r1`, `libblobmsg-json20260213 2026.06.19~7dd12784-r1`, `libc 1.2.5-r5`, `libcomerr0 1.47.3-r1`, `libe2p2 1.47.3-r1`, `libext2fs2 1.47.3-r1`, `libf2fs6 1.16.0-r4`, `libgcc1 14.3.0-r5`, `libiwinfo-data 2026.05.26~66bdd1a0-r1`, `libiwinfo20230701 2026.05.26~66bdd1a0-r1`, `libjson-c5 0.18-r1`, `libjson-script20260213 2026.06.19~7dd12784-r1`, `liblucihttp-ucode 2023.03.15~9b5b683f-r1`, `liblucihttp0 2023.03.15~9b5b683f-r1`, `libmbedtls21 3.6.7-r1`, `libmnl0 1.0.5-r1`, `libnftnl11 1.3.1-r1`, `libnl-tiny1 2025.12.02~40493a65-r1`, `libpthread 1.2.5-r5`, `librt 1.2.5-r5`, `libsmartcols1 2.41.5-r1`, `libss2 1.47.3-r1`, `libubox20260213 2026.06.19~7dd12784-r1`, `libubus20251202 2026.06.28~24864e78-r2`, `libuci20250120 2025.12.02~66127cd7-r1`, `libuclient20201210 2026.06.07~daad21fa-r1`, `libucode20230711 2026.01.16~85922056-r2`, `libudebug 2026.01.16~875e1a7a`, `libustream-mbedtls20201210 2026.03.01~99f1c0db-r1`, `libuuid1 2.41.5-r1`, `logd 2025.10.30~6f78fa49-r1`, `luci 26.288.56989~75a306d`, `luci-app-attendedsysupgrade 26.288.56989~75a306d`, `luci-app-firewall 26.288.56989~75a306d`, `luci-app-package-manager 26.288.56989~75a306d`, `luci-base 26.288.56989~75a306d`, `luci-lib-uqr 26.288.56989~75a306d`, `luci-light 26.288.56989~75a306d`, `luci-mod-admin-full 26.288.56989~75a306d`, `luci-mod-network 26.288.56989~75a306d`, `luci-mod-status 26.288.56989~75a306d`, `luci-mod-system 26.288.56989~75a306d`, `luci-proto-ipv6 26.288.56989~75a306d`, `luci-proto-ppp 26.288.56989~75a306d`, `luci-ssl 26.288.56989~75a306d`, `luci-theme-bootstrap 26.288.56989~75a306d`, `mkf2fs 1.16.0-r4`, `mtd 27`, `netifd 2026.02.26~cbb83a18-r2`, `nftables-json 1.1.6-r2`, `odhcp6c 2026.06.27~10a52220-r3`, `odhcpd-ipv6only 2026.06.29~5d7be43f-r1`, `openwrt-keyring 2025.12.10~f0670054-r1`, `owut 2026.08.30~9b51eeec-r1`, `partx-utils 2.41.5-r1`, `ppp 2.5.2-r4`, `ppp-mod-pppoe 2.5.2-r4`, `procd 2026.03.13~58eb263d-r1`, `procd-seccomp 2026.03.13~58eb263d-r1`, `procd-ujail 2026.03.13~58eb263d-r1`, `px5g-mbedtls 11`, `r8169-firmware 20260221-r1`, `rpcd 2026.10.04~d99f703e-r1`, `rpcd-mod-file 2026.10.04~d99f703e-r1`, `rpcd-mod-iwinfo 2026.10.04~d99f703e-r1`, `rpcd-mod-luci 20240305-r1`, `rpcd-mod-rpcsys 2026.10.04~d99f703e-r1`, `rpcd-mod-rrdns 20170710`, `rpcd-mod-ucode 2026.10.04~d99f703e-r1`, `ubox 2025.10.30~6f78fa49-r1`, `ubus 2026.06.28~24864e78-r2`, `ubusd 2026.06.28~24864e78-r2`, `uci 2025.12.02~66127cd7-r1`, `uclient-fetch 2026.06.07~daad21fa-r1`, `ucode 2026.01.16~85922056-r2`, `ucode-mod-fs 2026.01.16~85922056-r2`, `ucode-mod-html 1`, `ucode-mod-log 2026.01.16~85922056-r2`, `ucode-mod-math 2026.01.16~85922056-r2`, `ucode-mod-ubus 2026.01.16~85922056-r2`, `ucode-mod-uci 2026.01.16~85922056-r2`, `ucode-mod-uclient 2026.06.07~daad21fa-r1`, `ucode-mod-uloop 2026.01.16~85922056-r2`, `uhttpd 2026.08.24~373145f7-r1`, `uhttpd-mod-ubus 2026.08.24~373145f7-r1`, `urandom-seed 3`, `urngd 2025.10.03~f17e33d9-r1`, `usign 2025.10.03~c4c72b1b-r1`, `zlib 1.3.1-r1`

## 11. Size on disk

- Download: 12,498,825 bytes (squashfs) and 13,813,344 bytes (ext4), as `.img.gz`.
- Disk image: 126,123,520 bytes (120.3 MiB) for each.
- squashfs root: 6,096,242 bytes compressed. Unpacked, the root is 20,384,863 bytes (`du -sb`), the same for the
  ext4 root. Of that:
  - `lib/modules`: 7,971,111 bytes
  - `lib/firmware`: 1,411,732 bytes
  - `www` (LuCI): 1,337,317 bytes
- Guest `df -k` after the first boot (squashfs image):
  - `/rom`: 6,144 kB, 100 % used
  - `/overlay` (ext4 on loop0): 88,517 kB, 62 kB used, 81,422 kB free
  - `/boot`: 16,334 kB, 6,362 kB used, 9,972 kB free
  - `/tmp`: tmpfs, 115,236 kB
- ext4 image: the 104 MiB root has 19,361 of 26,624 4-KiB blocks free (about 75.6 MiB), from `debugfs -R stats`.

## 12. QEMU boot test (squashfs image)

**How it was run.** Each boot ran on a fresh copy of the decompressed image, under
`flock -w … ~/nmb-work/qemu.lock` for the whole run, one guest at a time. The QEMU command line:

```
qemu-system-x86_64 -machine q35 -accel tcg,tb-size=128 -cpu max -smp 1 -m 256 -nographic -monitor none
  -serial unix:<work>/serial.sock,server=on,wait=on
  -drive if=pflash,format=raw,readonly=on,file=OVMF_CODE_4M.fd  -drive if=pflash,format=raw,file=OVMF_VARS_4M.fd(copy)
  -drive file=disk.img,format=raw,if=virtio
  -netdev user,id=lan0,net=198.51.100.0/24,restrict=on -device virtio-net-pci,netdev=lan0,mac=02:00:00:00:01:01
  -netdev user,id=wan0,net=192.0.2.0/24             -device virtio-net-pci,netdev=wan0,mac=02:00:00:00:01:02
```

- QEMU 8.2.2 (Debian 1:8.2.2+ds-0ubuntu1.18) with OVMF from package `ovmf 2024.02-2ubuntu0.10`.
- **TCG, no KVM.** `/dev/kvm` exists but belongs to group `kvm`, and the user is not in that group (`test -r -a -w`
  failed). There is no sudo.
- The serial console was driven by a small Python script (`python3 -I`), with the script kept outside the data
  folder.
- 256 MB of RAM was used instead of 512 MB. When runs 2 and 3 started, the PC had about 250 MB available, because
  other packages were compiling (load average 8–18). Run 4 kept 256 MB so that the figures stay comparable.

| Run | Settings | Result |
| :--- | :--- | :--- |
| 1 | `-smp 2 -m 512` | Booted, but the host was overloaded: RCU stall warnings, and the console reached `Please press Enter` only after 183 s. The probes went out of step and the run took 2.5 h. QEMU was killed after a 120 s power-off timeout. Data not used, except as a boot proof |
| 2 | `-smp 1 -m 256` | Booted cleanly; console after 62.2 s. The driver's byte/character offset bug (GRUB's UTF-8 menu) made probe 7 miss its marker. `poweroff` was clean: `reboot: Power down`, QEMU exit 0, driver+flock exit 0 |
| 3 | `-smp 1 -m 256`, offset bug fixed | Probes 0–25 complete, used in this record. Probe 26 (`apk search`, about 500 characters) was cut on the serial input and waited at a `>` continuation prompt. QEMU was stopped with SIGTERM: QEMU says "terminating on signal 15", the driver exited 1, and the lock was released (`flock -n` rc 0 afterwards) |
| 4 | `-smp 1 -m 256`, short input lines sent in 32-byte chunks | Host idle (load 0.09 at the start). Booted cleanly; console after 39.7 s. All 22 probes gave exit 0. `poweroff` was clean: `reboot: Power down`, QEMU exit 0, driver+flock exit 0, 258.8 s in all (05:19:56–05:24:15 UTC) |

Figures from run 4 (TCG, one vCPU, 256 MB, idle host; **not** a performance figure for real hardware). The
values from run 3, on a loaded host, are given in brackets.

- **Boot time.**
  - Host clock from QEMU start to `Please press Enter to activate this console`: 39.7 s [67.0 s]. This includes
    OVMF and the 5 s GRUB timeout.
  - Kernel timestamps: `Run /sbin/init` at 4.75 s [7.54 s], `procd: - init -` at 22.57 s [36.78 s].
  - `procd: - init complete -` (from `logread`) was logged about 109 s [162 s] after kernel start. This is derived
    from the RTC line `setting system clock to … 05:20:14` at 2.60 s and the log entry at 05:22:01.
  - The first boot includes formatting the overlay.
- **Memory after `init complete`** (run 4, guest `free`): used 37,140 kB, free 165,900 kB, available 154,232 kB of
  230,476 kB. At the end of the run (uptime 221 s, after `apk update` and the tests) it was 41,520 kB used and
  147,652 kB available. `/overlay` then had 65 kB used, and `/tmp` 1,756 kB (the apk indexes).

Further detail from run 3:

- **Kernel memory line** (256 MB guest): `Memory: 191676K/255772K available`.
- **Storage after boot:** see section 11.
- **Listening ports** (guest `netstat -tulnp`):
  - TCP 22: dropbear, all addresses
  - TCP 80 and 443: uhttpd, all addresses
  - TCP and UDP 53: dnsmasq, on loopback, the LAN address, the WAN address 192.0.2.15, and IPv6 link-local and ULA
  - UDP 546: odhcp6c
  - UDP 547: odhcpd
  - Ephemeral UDP ports: dnsmasq and ntpd
- **Running services** (guest `ps w`): procd, ubusd, logd, rpcd, urngd, dropbear, netifd, odhcpd, udhcpc on eth1,
  odhcp6c on eth1, uhttpd, dnsmasq (inside `ujail`, user `dnsmasq`), ntpd (inside `ujail`, user `ntp`), and askfirst
  with the login shell on the console.
- **Security observations.**
  - The banner says "There is no root password defined on this device!".
  - `/etc/config/dropbear` has `PasswordAuth 'on'` and `RootPasswordAuth 'on'`.
  - LuCI listens on 80 and 443 on all addresses. The default firewall rejects input from the wan zone.
- **Kernel messages under TCG:** the TSC was marked unstable and the clock moved to HPET; `hrtimer: interrupt took …`.
  These come from the emulation and host load, not from the image.

## 13. Fields not determined

- **Boot time and memory on real hardware or KVM:** not determined. Only TCG on a loaded host was possible, because
  there is no access to `/dev/kvm`.
- **Free memory with 512 MB:** not determined. The usable run used 256 MB, because of host memory; run 1 (512 MB) was
  not usable.
- **Ext4 image boot:** not run. It was taken apart only, and its file and package lists equal the squashfs image's.
- **Licence field of the three firmware packages:** empty in the apk database.
- **Package versions from a PC-side read of the feed indexes:** not used. A `strings` read of the ADB files was not
  reliable, so the versions in 9a come from `apk search` in the guest.

## 14. Tools missing on the PC

`cpio`, `bsdtar`, `xorriso`, `parted`, `expect`, `socat` and `extract-ikconfig` are missing, and there is no sudo, so
no loop mounts and no access to `/dev/kvm`.

Present and used: `gzip`, `xz`, `sgdisk`, `fdisk`, `mtools`, `debugfs`, `unsquashfs` (squashfs-tools), `binwalk`,
`7z`, `zstd`, `file`, `strings`, `gpg`, `qemu-system-x86_64` 8.2.2, OVMF 2024.02 and `flock`.
