# Image record: Alpine Linux 3.24.4 (x86-64 standard ISO and Raspberry Pi aarch64 archive)

Package I-07 (work plan section 5.4). The method is work plan 5.1, and every field of 5.2 is below.
Each value names the command that produced it. A value that could not be read says "not determined" and why.
Nothing here comes from memory. Only lists, versions and checksums are recorded. No file taken out of the images
is in this repository.

- Date of the analysis: 2026-11-11.
- Where it ran: dev PC, Windows with WSL2 Ubuntu 24.04, about 1.8 GB RAM and 4 threads, no root. The QEMU guests
  ran on the same PC.
- Work folder: `~/nmb-work/I-07/r2/` (outside every repository; one fresh sub-folder per unpack). An earlier,
  interrupted attempt had left files in `~/nmb-work/I-07/`; none of them was used, every value below was produced
  again in `r2/`.
- Material, read only: `~/src/NMB-depythonization-firmwares/platforms/x86-64/alpine-standard-3.24.4-x86_64.iso` and
  `.../raspberry-pi-4-arm64/alpine-rpi-3.24.4-aarch64.tar.gz`.
- **Three kinds of source, kept apart:**
  - **[image]** values read from the two files: the boot files, the kernel configs, the `modloop` squashfs images,
    the initramfs lists, the `apks` folders and their `APKINDEX`.
  - **[booted]** values read in the x86-64 ISO booted diskless in QEMU (section 12). Alpine's standard ISO *is*
    the running system (diskless mode, root in RAM), not an installer for another system, so the booted ISO is the
    installed system for diskless use.
  - **[repo]** values from the signed online repositories `https://dl-cdn.alpinelinux.org/alpine/v3.24/{main,community}`
    (index `v3.24.4-60-g922ef2b92d7` / `v3.24.4-61-g48794b399db`), read by `apk` in the guest (`apk update`, exit 0, with the keys in
    `/etc/apk/keys`) and from the same `APKINDEX.tar.gz` files fetched on the PC on 2026-11-11 11:37 UTC. The tools
    of criterion 3 that are not on the media come from there.
- The Raspberry Pi archive was taken apart only; it was not booted (work plan 5.4: no Pi used in this package;
  `qemu-system-aarch64 -M help` on the PC lists `raspi0` to `raspi3b`, no Pi 4).

## 1. Files and checksums

| File | Size (bytes) | SHA-256 (publisher) | Check |
| :--- | ---: | :--- | :--- |
| `alpine-standard-3.24.4-x86_64.iso` | 370,147,328 | `92e3e4e4c6b8e265a889278de2a633c0f1456ca52820952666ecb395c3d23582` | OK |
| `alpine-rpi-3.24.4-aarch64.tar.gz` | 88,017,396 | `47524dfd4dc91f7039e1fea3cd70dff2d4c7d355f8bc57033b6d495d30d5a214` | OK |

- **How it was checked.** The set's README lists the Alpine files but no sums. The publisher's files were fetched
  on 2026-11-11 at 11:22 UTC from `https://dl-cdn.alpinelinux.org/alpine/v3.24/releases/{x86_64,aarch64}/<file>.{sha256,sha512,asc}`
  (all HTTP 200), with `latest-releases.yaml` of both architectures.
  - `sha256sum -c <file>.sha256` and `sha512sum -c <file>.sha512`, run in the material folders: all four `OK`, exit 0.
  - `latest-releases.yaml` gives the same SHA-256 and size for both files, `version: 3.24.4`, `date: 2026-11-05`.
- **Signatures.** `gpg --verify <file>.asc <file>` in an isolated `GNUPGHOME`: "Good signature" from the Alpine release signing key (the key holder's name and address are not copied) for both, exit 0, made Thu Nov 5 12:43:37 2026 CET (ISO) and 12:43:16 CET (Pi archive).
  - Key `0482 D840 22F5 2DF1 C4E7 CD43 293A CD09 07D9 495A` (rsa4096, 2014-12-10), fetched from the official
    Alpine keys page (`https://alpinelinux.org/keys/`). The same fingerprint is printed on `https://alpinelinux.org/downloads/`
    (fetched the same minute). gpg says "not certified with a trusted signature" (no web of trust set up).
- **Release identity.** `.alpine-release`: `alpine-standard-3.24.4 261105` (ISO), `alpine-rpi-3.24.4 261105`
  (archive). ISO volume id `alpine-std 3.24.4 x86_64` (`file`). Booted: `/etc/alpine-release` = `3.24.4`.
- **Inside the media.** Every `.apk` carries its own signature (`.SIGN.RSA.alpine-devel@lists.alpinelinux.org-6165ee59.rsa.pub`
  on x86_64, `-616ae350` on aarch64, seen when unpacking `APKINDEX.tar.gz` and four packages with `tar -xzf`). The modloop is signed too
  (section 3).

## 2. Partition and file system layout, boot method

**[image] x86-64 ISO.** Read with `file`, `fdisk -l` (exit 0), `xorriso -indev <iso> -report_el_torito plain` and
`7z l` (exit 0, 116 files, 8 folders). Extracted with `7z x -o<fresh dir> -x'![BOOT]'` (exit 0, 114 files). Nothing
was mounted.

| | Value |
| :--- | :--- |
| `file` | ISO 9660, volume `alpine-std 3.24.4 x86_64`, DOS/MBR boot sector, bootable |
| MBR (`fdisk -l`) | p1 type 0 "Empty", sectors 0–722943 (353 MiB, the whole ISO, boot flag); p2 type `ef` EFI, sectors 220–3099 (1.4 MiB) |
| El Torito (`xorriso`) | "MBR isohybrid cyl-align-on GPT"; image 1 BIOS `/boot/syslinux/isolinux.bin` (boot-info-table, isohybrid); image 2 UEFI `/boot/grub/efi.img` (1,474,560 bytes) |
| BIOS path | `boot/syslinux/syslinux.cfg`: `DEFAULT lts`, `KERNEL /boot/vmlinuz-lts`, `INITRD /boot/initramfs-lts`, `FDTDIR /boot/dtbs-lts`, `APPEND modules=loop,squashfs,sd-mod,usb-storage quiet`, `TIMEOUT 10` |
| UEFI path | `efi/boot/bootx64.efi` (847,872 bytes); `boot/grub/grub.cfg`: `set timeout=1`, `linux /boot/vmlinuz-lts modules=loop,squashfs,sd-mod,usb-storage quiet`, `initrd /boot/initramfs-lts`. The guest printed `GNU GRUB version 2.14` |
| Boot files | `vmlinuz-lts` 14,607,360; `initramfs-lts` 22,545,663 (gzip, 74,287,196 unpacked); `modloop-lts` 307,982,336 (squashfs 4.0, xz, 8,510 inodes, block 128 KiB); `config-6.18.60-0-lts` 281,567; `System.map-6.18.60-0-lts` 7,389,713; `dtbs-lts/` empty |
| Packages | `apks/x86_64/`: 95 `.apk` + `APKINDEX.tar.gz`, 13,918,953 bytes (`du -sb`) |

**[image] Raspberry Pi archive.** Listed with `tar -tvzf` (exit 0, 530 entries) and unpacked with
`tar -xzf ... --no-same-owner --no-same-permissions` into a fresh folder (exit 0, 110,391,389 bytes, `du -sb`).
It is the content of a FAT boot partition: the user copies it onto a FAT-formatted SD card (the publisher also
offers `alpine-rpi-3.24.4-aarch64.img.gz`, 88,115,056 bytes, per `latest-releases.yaml`; that image is not in the set).

| | Value |
| :--- | :--- |
| Pi firmware | `bootcode.bin` 52,624, `start.elf` 3,023,104, `fixup.dat` 7,381, **`start4.elf` 2,298,912, `fixup4.dat` 5,510** (Pi 4); no `start4x`/`start4cd` variants |
| `config.txt` (whole file) | comment "do not modify this file as it will be overwritten on upgrade. create and/or modify usercfg.txt instead."; `kernel=boot/vmlinuz-rpi`, `initramfs boot/initramfs-rpi`, `arm_64bit=1`, `include usercfg.txt` (no `usercfg.txt` in the archive) |
| `cmdline.txt` | `modules=loop,squashfs,sd-mod,usb-storage quiet console=tty1` (no serial console) |
| Device trees | 28 `.dtb` in the root, among them `bcm2711-rpi-4-b.dtb`, `bcm2711-rpi-400.dtb`, `bcm2711-rpi-cm4.dtb`, `-cm4-io`, `-cm4s`, `-4-a`, `-zero-3-w`, plus Pi 2/3/Zero 2/5 ones; `overlays/` 387 files, 954,175 bytes (with `disable-wifi.dtbo`, `disable-bt.dtbo`, `miniuart-bt.dtbo`, `dwc2.dtbo`, `vc4-kms-v3d.dtbo`) |
| Boot files | `boot/vmlinuz-rpi` 24,218,112 (ARM64 `Image`, uncompressed, 4K pages); `boot/initramfs-rpi` 6,999,726 (gzip, 18,419,168 unpacked); `boot/modloop-rpi` 52,080,640 (squashfs 4.0, xz, 2,819 inodes); `config-6.18.60-0-rpi` 255,655; `System.map-6.18.60-0-rpi` 4,711,304 |
| Packages | `apks/aarch64/`: 95 `.apk` + index, 14,409,855 bytes |

- **Boot chain, both:** firmware → kernel + initramfs (Alpine `mkinitfs` 3.14.1-r0, `VERSION=` in `/init`) →
  `nlplug-findfs` finds the boot media, the `modloop` and any `*.apkovl.tar.gz` → root is a **tmpfs** →
  `apk add --initdb` of the packages in the overlay's `/etc/apk/world` from the media's `apks` → OpenRC. Read from
  the initramfs `init` script (1,187 lines; taken out with `zcat | cpio -id init` into a fresh folder, not run) and
  seen in the guest (section 12).
- **Diskless mode** (the default for both media): `mount -t tmpfs -o mode=0755 tmpfs $sysroot` (no size option: the
  kernel default, half of RAM; guest: 233,520 kB of 467,040 kB). The boot media is mounted read-only (`/dev/sr0 on /media/sr0 type iso9660 (ro,...)` in the guest).
- **Other install modes**, from `alpine-conf 3.22.0-r0` (`setup-disk`, 1,927 lines, `setup-bootable`, 450 lines,
  unpacked from the ISO's `apks` with `tar -xzf`): `setup-disk -m data` (diskless root, `/var` on a disk) and
  `setup-disk -m sys` (classic installed root on disk). Not exercised here.

## 3. How an update is written

Diskless mode has three separately written parts. Read from `alpine-conf 3.22.0-r0` (`/usr/sbin/lbu`, 777 lines;
`/usr/sbin/update-kernel`, 414 lines) and `openrc 0.63.2-r0` (`/etc/init.d/modloop`), unpacked from the ISO; the
`lbu` path was also run in the guest (section 12).

- **Configuration: `lbu commit`.** `lbu` collects changed files with `apk audit --backup --recursive --check-permissions`
  (plus the `lbu include` list, `$LBU_LIST`), writes `<hostname>.apkovl.tar.gz` to a temporary
  file, copies it to `<file>.new` on the media and **renames** it over the old one (`cp "$tmppkg" "$pkg.new"` then
  `mv "$pkg.new" "$pkg"`), then `sync`. Before that the old file is renamed to `<hostname>.<UTC date>.tar.gz`; with
  the default `BACKUP_LIMIT` of 0 that backup is deleted after a successful commit. The media is mounted (or remounted
  read-write) for the commit and unmounted or remounted read-only at the end (`mount_once_rw`; `cleanup` runs on
  every exit). Optional encryption with
  `openssl enc` (`/etc/lbu/lbu.conf`: `DEFAULT_CIPHER=aes-256-cbc`, `# ENCRYPTION=$DEFAULT_CIPHER` commented out). If two `*.apkovl.tar.gz` are on
  the media, commit refuses ("this is a security concern").
- **Packages: `apk add`/`apk upgrade` into RAM**, and the `.apk` files into a cache on the media
  (`setup-apkcache`, `/etc/apk/cache` link). At the next boot the initramfs reinstalls `/etc/apk/world` from the
  media and the cache. So a package update is only kept if the cache and the overlay are committed.
- **Kernel, initramfs and modloop: `update-kernel`.** It builds a new modloop in a temporary folder from the kernel
  package (`mksquashfs ... -comp xz -Xbcj x86` / `arm,armthumb`), copies only the firmware files that some module's
  `modinfo -F firmware` names (plus `-e` extras), optionally signs it (`-s`), builds a new initramfs with
  `mkinitfs`, then **unmounts `/.modloop`, remounts the boot media read-write and `mv`s the new files over the old
  ones** (`mv $STAGING/* "$DESTDIR"/boot/`), `sync`, remounts read-only and loop-mounts the new modloop. There is no
  second slot and no automatic fallback: a failure or power cut between the `mv` of the kernel, initramfs and modloop
  can leave a mismatched set. On the Pi the DTBs are copied into the partition root (`yes,rpi*) _dtb="$DESTDIR/"`).
- **Modloop signature.** The initramfs carries `var/cache/misc/modloop-lts.SIGN.RSA.alpine-devel@lists.alpinelinux.org-6165ee59.rsa.pub`
  (Pi: `modloop-rpi.SIGN.RSA...-616ae350.rsa.pub`) (`cpio -it` lists); `/etc/init.d/modloop` verifies the modloop
  with `openssl dgst -sha1 -verify <key> -signature <sig>` unless `modloop_verify=no` is on the kernel command line,
  and only logs an error (`eerror "Failed to verify signature..."`) and still mounts it if verification fails.
- **The release media lag the repository.** [repo] `linux-lts` is `6.18.63-r0` while the ISO's kernel is
  `6.18.60-0-lts`; `linux-rpi` is still `6.18.60-r0` (`APKINDEX` of `main`, both architectures).
- **A/B building blocks** ([repo], `apkq.py info` over both indexes): `rauc` and `swupdate` are **absent** from
  v3.24 main and community; `alpine-make-rootfs 0.8.1-r0`, `mkinitfs 3.14.1-r0`, `abuild 3.17.0-r0` are in `main`.
  An A/B scheme would have to be NMB's own (two boot partitions or two sets of `vmlinuz`/`initramfs`/`modloop`
  plus a boot-loader switch).
- **Security support** (criterion 2): `https://alpinelinux.org/releases.json` (fetched 2026-11-11) gives
  `v3.24` `eol_date 2028-06-01` (3.23: 2027-11-01, 3.22: 2027-05-01). `https://alpinelinux.org/releases/` says the
  main repository "is typically supported for 2 years and the community repository is supported until next stable
  release". No subscription. Every criterion-3 tool is in **main** (section 7), so the 2-year span applies to them.

## 4. Kernel and the configuration NMB needs

- **Kernel version: 6.18.60** in both.
  - x86-64: `Linux kernel x86 boot executable bzImage, version 6.18.60-0-lts (buildozer@build-3-24-x86_64) #1-Alpine SMP PREEMPT_DYNAMIC 2026-11-01 07:26:31` (`file`);
    guest `uname -a`: `Linux localhost 6.18.60-0-lts #1-Alpine SMP PREEMPT_DYNAMIC 2026-11-01 07:26:31 x86_64 Linux`.
  - Pi: `Linux version 6.18.60-0-rpi (buildozer@build-3-24-aarch64) (cc (Alpine 15.2.0) 15.2.0, GNU ld (GNU Binutils) 2.45.1) #1-Alpine SMP PREEMPT`
    (`strings` on `vmlinuz-rpi`). It is the Raspberry Pi kernel tree, not mainline: the config has
    `BCM_VCIO=y`, `BCM2835_VCHIQ=y`, `RASPBERRYPI_GPIOMEM=m`, `MFD_RP1=y`, `PINCTRL_RP1=y`, `DRM_RP1_DSI=m`.
- **Where the configs came from:** `boot/config-6.18.60-0-lts` (ISO) and `boot/config-6.18.60-0-rpi` (archive),
  both `CONFIG_CC_VERSION_TEXT="cc (Alpine 15.2.0) 15.2.0"`.
  - x86-64: 11,288 lines, `# Linux/x86_64 6.18.60 Kernel Configuration`, 3,604 `=m`, 2,429 `=y`, SHA-256
    `b1f96fd4f6b51f9cee1ff97081cb98439a20dca8cbb672d0d6ee09300a1e0789`.
  - Pi: 9,378 lines, `# Linux/arm64 6.18.60 Kernel Configuration`, 1,911 `=m`, 1,966 `=y`, SHA-256
    `5f5bf20ad5ec50080f7104985304675fa0e83344a557d011eab8f22499302f4c`.
  - `IKCONFIG=m`, `IKCONFIG_PROC=y` in both. In the guest `modprobe configs; zcat /proc/config.gz | sha256sum`
    gave `b1f96fd4f6b51f9cee1ff97081cb98439a20dca8cbb672d0d6ee09300a1e0789`: the running kernel's config is byte-identical to the file on the ISO.
- **Source of every option below: those two files**, read with `opts.sh` (`grep -E "^(CONFIG_X=|# CONFIG_X is not set)"`
  per option). "Module" says whether the `.ko` is in the modloop (`grep -c "/<name>.ko$"` over `unsquashfs -lls`)
  or built in (`modules.builtin`, taken out with `unsquashfs -d <fresh> <modloop> modules/<kver>/modules.builtin`:
  171 lines x86-64, 426 Pi). "Guest" is the `modprobe <name>; echo $?` loop in the booted x86-64 ISO.

| Feature (5.2) | Option(s) | x86-64 (`lts`) | Pi (`rpi`) | Module | Guest (x86-64) |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Bridge, VLAN filtering | `BRIDGE`, `BRIDGE_VLAN_FILTERING`, `BRIDGE_NETFILTER` | m, y, m | m, y, m | `bridge` both | 0 |
| Bonding | `BONDING` | m | m | `bonding` both | 0 (created `bond0`) |
| Team | `NET_TEAM` | m | **not set** | — | — |
| VLAN 802.1Q | `VLAN_8021Q` | m | m | `8021q` both | 0 |
| nftables | `NF_TABLES` m, `NF_TABLES_INET` y, `NF_TABLES_NETDEV` y, `NF_TABLES_BRIDGE` m, `NFT_NAT` m, `NFT_MASQ` m, `NFT_FIB_INET` m, `NFT_CT` m, `NFT_COMPAT` m, `NF_FLOW_TABLE` m, `NFT_FLOW_OFFLOAD` m | as listed | as listed | `nf_tables`, `nft_nat`, `nft_masq`, `nft_fib_inet`, `nft_ct` both | `nf_tables nft_nat nft_masq nft_fib_inet nft_ct nft_compat` all 0; `nft` worked (section 9) |
| iptables (legacy) | `IP_NF_IPTABLES` | m | m | — | — |
| WireGuard | `WIREGUARD` | m | m | `wireguard` both | 0 |
| PPP, PPPoE | `PPP`, `PPPOE`, `PPP_ASYNC`, `PPP_SYNC_TTY`, `L2TP`, `PPPOL2TP` | all m | all m | `ppp_generic`, `pppoe`, `pppox`, `ppp_async`, `l2tp_ppp` both | `ppp_generic pppoe ppp_async l2tp_ppp` 0 |
| cfg80211, mac80211 | `CFG80211`, `MAC80211` | m, m | m, m | both | 0, 0 |
| mac80211_hwsim | `MAC80211_HWSIM` | m | m | both | `modprobe mac80211_hwsim radios=2` 0: `phy0 phy1`, `wlan0 wlan1 hwsim0` |
| Wi-Fi drivers (examples) | `ATH9K`, `ATH9K_HTC`, `ATH10K_PCI`, `IWLWIFI`, `MT7601U`, `MT76x2U`, `MT7921U`, `RTL8XXXU`, `RT2800USB`, `BRCMFMAC` (+`BRCMFMAC_SDIO` y) | all m | all m except `ATH10K_PCI` absent | 171 (x86-64) / 106 (Pi) `.ko` under `drivers/net/wireless/` | — |
| USB host | `USB`, `USB_XHCI_HCD`, `USB_XHCI_PCI`, `USB_EHCI_HCD`; `USB_DWC2` | m, m, m, m; DWC2 not set | y, y, y, EHCI not set; DWC2 m | Pi: `xhci-hcd`, `xhci-pci` built in | — |
| USB network | `USB_USBNET`, `USB_RTL8152`, `USB_NET_AX88179_178A`, `USB_NET_AX8817X`, `USB_NET_SMSC95XX`, `USB_LAN78XX`, `USB_NET_AQC111` | all m, `AQC111` not set | `USBNET` y, `RTL8152` m, `AX88179_178A` m, `AX8817X` m, `SMSC95XX` y, `LAN78XX` y, `AQC111` m | x86-64: all as `.ko` (no `aqc111`); Pi: `usbnet`, `smsc95xx`, `lan78xx` built in, `r8152`, `ax88179_178a`, `aqc111` `.ko` | `usbnet r8152 ax88179_178a smsc95xx lan78xx` 0 |
| USB tethering (4.27) | `USB_NET_RNDIS_HOST`, `USB_NET_CDCETHER`, `USB_NET_CDC_NCM`, `USB_IPHETH`, `USB_NET_QMI_WWAN`, `USB_NET_CDC_MBIM`, `USB_NET_HUAWEI_CDC_NCM`, `USB_NET_CDC_EEM` | all m | all m | all present both | `rndis_host cdc_ether cdc_ncm ipheth qmi_wwan cdc_mbim huawei_cdc_ncm` all 0 |
| TUN, veth, macvlan, ipvlan, dummy, IFB | `TUN`, `VETH`, `MACVLAN`, `IPVLAN`, `DUMMY`, `IFB` | all m | all m | all present both | all 0 |
| Policy routing | `IP_ADVANCED_ROUTER`, `IP_MULTIPLE_TABLES`, `IPV6_MULTIPLE_TABLES` | y | y | built in | — |
| IPv6 | `IPV6` | y | **m** (module on the Pi) | — | — |
| Traffic control | `NET_SCHED` y, `NET_SCH_FQ_CODEL`, `NET_SCH_CAKE`, `NET_SCH_HTB`, `NET_SCH_NETEM`, `NET_SCH_TBF`, `NET_SCH_FQ`, `NET_SCH_INGRESS` m; `NET_SCH_DEFAULT` not set | as listed | as listed | `sch_cake`, `sch_netem`, `sch_fq_codel` both | `sch_cake sch_netem sch_fq_codel tcp_bbr` 0; `eth0` came up with `qdisc pfifo_fast` |
| Namespaces | `NAMESPACES`, `NET_NS`, `PID_NS`, `USER_NS`, `UTS_NS`, `IPC_NS`, `TIME_NS` | all y | all y | built in | `/proc/self/ns`: cgroup, ipc, mnt, net, pid, time, user, uts |
| seccomp | `SECCOMP`, `SECCOMP_FILTER` | y, y | y, y | built in | — (nothing in the base uses it: `Seccomp: 0` for the shell) |
| **Landlock** | `SECURITY` y; `SECURITY_LANDLOCK`; `LSM` | **y**; `LSM="landlock,lockdown,yama,loadpin,safesetid,integrity"` | **not set**; `LSM=""` | built in (x86-64) | **active on x86-64**: guest `/sys/kernel/security/lsm` = `lockdown,capability,landlock,yama` |
| Other LSMs | `SECURITY_APPARMOR` y in both (x86-64 `DEFAULT_SECURITY_DAC` y, Pi `DEFAULT_SECURITY_APPARMOR` y); `SECURITY_YAMA`, `SECURITY_LOCKDOWN_LSM` y (x86-64), not set (Pi); `SECURITY_SELINUX` not set; `BPF_LSM` y (x86-64), not set (Pi) | as listed | as listed | — | AppArmor built but **not** in the active list (x86-64) |
| cgroups | `CGROUPS`, `MEMCG` (`MEMCG_V1` not set), `CGROUP_PIDS`, `CPUSETS`, `CGROUP_BPF`, `BLK_CGROUP`, `CGROUP_FREEZER`, `CGROUP_DEVICE`, `CGROUP_SCHED` | all y | all y | — | **not mounted at boot**: no cgroup line in `mount`, OpenRC's `cgroups` service is not in any runlevel (section 12) |
| Watchdog | `WATCHDOG` y; `SOFT_WATCHDOG`, `I6300ESB_WDT`, `ITCO_WDT` m (x86-64); `BCM2835_WDT` **y** (Pi), `SOFT_WATCHDOG` not set (Pi) | as listed | as listed | `softdog`, `i6300esb`, `iTCO_wdt` (x86-64); `bcm2835_wdt` built in (Pi) | `softdog i6300esb iTCO_wdt` 0; `/dev/watchdog`, `/dev/watchdog0`, `/dev/watchdog1` present, `watchdog0` identity `iTCO_wdt`; BusyBox `watchdog` init script present, not enabled |
| Module signing, compression | `MODULE_SIG` y, `MODULE_SIG_SHA256` y, `MODULE_SIG_ALL` y, `MODULE_SIG_FORCE` not set, `MODULE_COMPRESS_GZIP` y (x86-64); `MODULE_SIG` not set, `MODULE_COMPRESS_XZ` y (Pi). The modloop holds plain `.ko` files in both | as listed | as listed | — | — |
| Firmware loader | `FW_LOADER_COMPRESS` y, `FW_LOADER_COMPRESS_ZSTD` y (both), `_XZ` y (Pi only) | as listed | as listed | firmware files are `.zst` | — |
| Raspberry Pi 4 platform | `ARCH_BCM2835` y, `ARCH_BCM` y, `PCIE_BRCMSTB` y, `RASPBERRYPI_FIRMWARE` y, `RASPBERRYPI_POWER` y, `RESET_RASPBERRYPI` y, `BCMGENET` y, `MMC_SDHCI_IPROC` y, `MMC_BCM2835` y, `HW_RANDOM_BCM2835` y, `SERIAL_8250_BCM2835AUX` y, `BCM2711_THERMAL` y, `DRM_VC4` m, `SENSORS_RASPBERRYPI_HWMON` m | absent (`BCMGENET` m) | as listed | `genet`, `bcm2835_wdt` built in | not booted |
| Other | `BPF_SYSCALL` y; `XDP_SOCKETS` y (x86-64) / not set (Pi); `TCP_CONG_BBR` m; `MPTCP` y; `TLS` m; `PREEMPT_DYNAMIC` + `PREEMPT_VOLUNTARY` (x86-64) / `PREEMPT` y (Pi); `HZ` 1000 / 250; `NR_CPUS` 256 / 4; `SQUASHFS` m (+`SQUASHFS_XZ` y), `OVERLAY_FS` m, `EROFS_FS` m, `DM_VERITY` m, `ZRAM` m; `EXT4_FS`, `VFAT_FS`, `F2FS_FS` m (x86-64) / y (Pi); `EFI_STUB` y | as listed | as listed | — | — |

- `grep CGROUP_NS` found no such option in either config; the guest's `/proc/self/ns/cgroup` shows cgroup
  namespaces exist.

## 5. Modules and firmware files carried

All from `unsquashfs -no-xattrs -lls <modloop>` (exit 0, no errors) and `zcat <initramfs> | cpio -it --quiet` (exit 0).

| | x86-64 `modloop-lts` | Pi `modloop-rpi` |
| :--- | :--- | :--- |
| Squashfs file | 307,982,336 bytes | 52,080,640 bytes |
| SHA-256 | `49e09b02671ab555e7dc96d0006a2e8ab50a91ed4636b4a7eb98de7660fff8ee` | `008d8a541591e7b4d686d7b17e4bba519aa34eea7f838310d3c552d482883a92` |
| Paths | 8,510 | 2,819 |
| Kernel modules | 3,800 `.ko` (plain), 535,049,792 bytes | 1,915 `.ko`, 99,318,616 bytes |
| `modules.builtin` | 171 lines | 426 lines |
| Firmware (`modules/firmware/`) | 3,570 files (5 symlinks), 1,490,722,116 bytes as stored (`.zst`) | 444 files, 37,008,528 bytes |
| Largest firmware folders | `cirrus` 1,412, `amdgpu` 662, `nvidia` 650, `radeon` 233, `brcm` 185, `i915` 50, `rtl_bt` 43, `rtl_nic` 34, `ath11k` 32, `rtlwifi` 31 entries | `brcm` 166, `rtl_bt` 43, `rtl_nic` 34, `mediatek` 19, `rtlwifi` 15, `libertas` 12, `mrvl` 10, `rtw88` 9 |
| Initramfs | 610 entries, 373 `.ko`, 32 firmware files | 305 entries, 152 `.ko`, 6 firmware files |

- **The firmware set is selected, not complete.** `update-kernel` (section 3) puts into the modloop only firmware
  that a module names in `modinfo -F firmware`. Firmware NMB cares about, all present in **both** modloops
  (`grep` over the listings): the Pi 4 Wi-Fi `brcm/brcmfmac43455-sdio.raspberrypi,4-model-b.{bin,clm_blob,txt}.zst`
  (and `,4-compute-module`, `,500`, `,5-model-b`); `rtl_nic/rtl8153a-2/-3/-4`, `rtl8153b-2`, `rtl8153c-1`,
  `rtl8156a-2`, `rtl8156b-2.fw.zst` (r8152 USB Ethernet); `mt7601u.bin.zst`, `mt7662*.bin.zst`;
  `ath9k_htc/htc_7010-1.4.0`, `htc_9271-1.4.0.fw.zst`; `regulatory.db` and `regulatory.db.p7s`.
- **Firmware not in the modloop** comes as packages: [repo] `linux-firmware-brcm 20260519-r0`, `linux-firmware-rtl_nic`,
  `linux-firmware-cypress`, ... in `main` (both architectures), and `update-kernel -p <pkg>`/`-e <file>` puts them
  into a new modloop. The x86-64 ISO page says "Includes AMD and Intel microcode updates" (`latest-releases.yaml`);
  no `microcode` path was found in the initramfs list or the ISO list (`grep -i microcode`), so where it is was
  not determined.
- **Booted:** `/lib/modules -> /.modloop/modules`, `/lib/firmware -> /lib/modules/firmware` (symlinks made by
  `/etc/init.d/modloop`, owned by `openrc-0.63.2-r0`); the modloop is loop-mounted read-only on `/.modloop`
  (`/dev/loop0 ... squashfs (ro,relatime,errors=continue)`). `lsmod` had 133 lines after the probe loop.

## 6. Init system, service manager, C library, package manager

- **Init: BusyBox `init` + OpenRC 0.63.2** (`openrc-0.63.2-r0`, `openrc-user-0.63.2-r0`). PID 1 is `/sbin/init`
  (BusyBox); OpenRC runs the runlevels. [image] `alpine-base 3.24.4-r0` depends on `alpine-baselayout alpine-conf
  alpine-release apk-tools busybox busybox-mdev-openrc busybox-openrc busybox-suid musl-utils openrc ifupdown-any`
  (`.PKGINFO`). Device manager: BusyBox `mdev` (`mdev-conf 4.10-r0`), not udev. `rc.conf`: `#rc_parallel="NO"`,
  `#rc_cgroup_mode="unified"` (all commented: defaults).
- **Services the diskless boot enables** (initramfs `init`, `rc_add` when no overlay or `/etc/.default_boot_services`;
  guest `rc-update show -v`): sysinit `devfs dmesg mdev hwdrivers modloop`; boot `modules sysctl hostname bootmisc
  syslog` (+ `hwclock`); shutdown `mount-ro killprocs savecache`; default `firstboot`. `sshd` only with the
  `ssh_key=` boot option; `tiny-cloud` only with `ds=`/`tinycloud`.
- **C library: musl 1.2.6** (guest `/lib/ld-musl-x86_64.so.1`: `musl libc (x86_64) Version 1.2.6`, owned by
  `musl-1.2.6-r3`; Pi archive and initramfs: `ld-musl-aarch64.so.1`, `musl 1.2.6-r3` in its index). No glibc.
- **Package manager: apk-tools 3.0.9** (`apk-tools 3.0.9-r0, compiled for x86_64.`). The index has dependency
  fields `D:` (77 of 95 packages), `i:` install-if (12) and `p:` provides (61), and no recommends field
  (`cut -c1-2 APKINDEX | sort | uniq -c`), so only declared dependencies and install-if rules pull packages.
- **Package counts.**
  - [image] Each `apks` folder: **95** packages, the same names and versions on both media (`diff` of the
    `P:`/`V:` pairs of the two `APKINDEX`: no difference). Sum of `S:` 13,906,967 (x86-64) / 14,397,865 (aarch64)
    bytes, of `I:` (installed) 32,209,965 / 39,253,539 bytes.
  - [booted] at the login prompt: **28** installed (`apk info | wc -l`; world = `alpine-base openssl`);
    **78** after adding the criterion-3 tools (section 7).
  - [repo] v3.24: `main` 5,961 (x86_64) / 5,983 (aarch64), `community` 22,708 / 22,588 (`grep -c '^P:'`);
    `apk update` in the guest: "OK: 28670 distinct packages available".
- **The 95 packages on the media** (`APKINDEX`; `name version`): `alpine-base 3.24.4-r0`, `alpine-baselayout 3.7.2-r1`,
  `alpine-baselayout-data 3.7.2-r1`, `alpine-conf 3.22.0-r0`, `alpine-keys 2.6-r0`, `alpine-release 3.24.4-r0`,
  `apk-cron 1.0-r3`, `apk-tools 3.0.9-r0`, `bonding 2.6-r5`, `bridge 1.5-r5`, `busybox 1.37.0-r31` (+`-binsh`,
  `-mdev-openrc`, `-openrc`, `-suid`), `ca-certificates-bundle 20261020-r0`, `chrony 4.8-r7` (+`-common`, `-openrc`),
  `dbus-libs 1.16.2-r2`, `dhcpcd 10.5.2-r1` (+`-openrc`), `doas 6.8.2-r8`, `e2fsprogs 1.47.4-r0` (+`-extra`, `-libs`),
  `ifupdown-ng 0.13.0-r0` (+`-ppp`, `-wifi`), `iw 6.17-r0`, `kbd-bkeymaps 2.8.0-r0`, `libapk 3.0.9-r0`,
  `libblkid 2.42.3-r1`, `libcap2 2.78-r0`, `libcom_err 1.47.4-r0`, `libcrypto3 3.5.9-r0`, `libeconf 0.8.3-r0`,
  `libedit 20260508.3.1-r1`, `libfdisk 2.42.3-r1`, `libidn2 2.3.8-r0`, `libncursesw 6.6_p20260516-r0`,
  `libnl3 3.11.0-r0`, `libpcap 1.10.7-r0`, `libseccomp 2.6.1-r0`, `libsmartcols 2.42.3-r1`, `libssl3 3.5.9-r0`,
  `libunistring 1.4.2-r0`, `libusb 1.0.30-r0`, `libuuid 2.42.3-r1`, `mdev-conf 4.10-r0`, `musl 1.2.6-r3`,
  `musl-utils 1.2.6-r3`, `ncurses-terminfo-base 6.6_p20260516-r0`, `network-extras 2.1-r0`, `openntpd 6.8_p1-r12`
  (+`-openrc`), `openrc 0.63.2-r0`, `openrc-user 0.63.2-r0`, `openssh 10.3_p1-r2` (+ 7 subpackages), `openssl 3.5.9-r0`,
  `partx 2.42.3-r1`, `pcre2 10.49-r0`, `pcsc-lite-libs 2.4.0-r4`, `ppp-atm`, `ppp-chat`, `ppp-daemon`, `ppp-l2tp`,
  `ppp-minconn`, `ppp-openrc`, `ppp-passwordfd`, `ppp-pppoe`, `ppp-radius`, `ppp-winbind` (all `2.5.3-r0`),
  `scanelf 1.3.9-r1`, `sfdisk 2.42.3-r1`, `ssl_client 1.37.0-r31`, `tiny-cloud 3.4.1-r2` (+`-nocloud`, `-openrc`),
  `tzdata 2026f-r0`, `usb-modeswitch 2.6.1-r5`, `vlan 2.3-r1`, `wget 1.25.0-r3`, `wpa_supplicant 2.11-r4`
  (+`-openrc`), `yaml 0.2.5-r2`, `yx 1.0.2-r0`, `zlib 1.3.2-r1`.
- **Licences (criterion 7).** `L:` fields of the 95: GPL-2.0-only 14, MIT 12, GPL-2.0-or-later 11, SSH-OpenSSH 8,
  ISC 7, LGPL-2.1-or-later 6, BSD-3-Clause 6, mixed GPL/LGPL/BSD/MIT 4, BSD-2-Clause 4, Apache-2.0 3, the rest
  single (Zlib, MPL-2.0 AND MIT, GPL-3.0-or-later WITH OpenSSL-Exception, ...). The firmware packages are `L: custom`
  ([repo] `linux-firmware-*`, `raspberrypi-bootloader 1.20260915-r0`); their texts were not read.

## 7. Tools of 6A criterion 3, with versions

On the media only `wpa_supplicant`, `iw`, `ppp-*`, `bridge`, `bonding`, `vlan`, `dhcpcd` and `usb-modeswitch` are
present as packages; the rest comes from [repo] `main`. In the guest: `command -v` before and after
`apk add nftables dnsmasq hostapd wpa_supplicant iw ppp-daemon ppp-pppoe wireguard-tools openvpn iproute2 ethtool`
(exit 0, "(50/50) ... OK: 26.8 MiB in 78 packages"); versions from each tool's own `--version`/`-v`.
"Repo" is the repository `apk policy` names; "Media" says whether the package is in the `apks` folder.

| Tool | At boot (28 pkgs) | Version after `apk add` | Package, repo | Media |
| :--- | :--- | :--- | :--- | :--- |
| OpenRC (init) | yes | `openrc 0.63.2` | `openrc 0.63.2-r0`, main | yes |
| systemd, procd | no | — | not in Alpine | — |
| nftables `nft` | no | `nftables v1.1.6 (Commodore Bullmoose #7)` | `nftables 1.1.6-r1` (+`nftables-openrc`), main | no |
| dnsmasq | no | `Dnsmasq version 2.92rel2` | `dnsmasq 2.92_p2-r0` (+`-common`, `-openrc`), main | no |
| hostapd | no | `hostapd v2.11` | `hostapd 2.11-r4` (+`-openrc`), main | no |
| wpa_supplicant | no (package on media, not installed) | `wpa_supplicant v2.11` | `wpa_supplicant 2.11-r4`, main | yes |
| iw | no | `iw version 6.17` | `iw 6.17-r0`, main | yes |
| pppd (+PPPoE) | no | `pppd version 2.5.3`; plugins `pppoe.so`, `dhcpv6relay.so` | `ppp-daemon`, `ppp-pppoe 2.5.3-r0`, main | yes |
| WireGuard `wg`, `wg-quick` | no | `wireguard-tools v1.0.20260223` | `wireguard-tools 1.0.20260223-r0`, main; `wg-quick` pulls `bash 5.3.9-r1` and `openresolv 3.17.4-r0` | no |
| OpenVPN | no | `OpenVPN 2.7.7 x86_64-alpine-linux-musl [SSL (OpenSSL)] [LZO] [LZ4] [EPOLL] [PKCS11] [MH/PKTINFO] [AEAD] [DCO]` | `openvpn 2.7.7-r0`, main | no |
| `ip`, `tc`, `ss` | `ip` = BusyBox applet (`/sbin/ip -> /bin/busybox`); `tc`, `ss`, `bridge` absent | `ip utility, iproute2-v7.0.0` (`/sbin/ip` 614,992 bytes) | `iproute2 7.0.0-r0` (+`-minimal`, `-tc`, `-ss`), main | no |
| ethtool | no | `ethtool version 7.0` | `ethtool 7.0-r0`, main | no |
| iptables | no | not installed | `iptables 1.8.13-r0`, main | no |
| DHCP client | BusyBox `udhcpc` (used by ifupdown-ng) | — | `dhcpcd 10.5.2-r1` on media, not installed | yes |
| Network config | `ifupdown-ng 0.13.0` (`/sbin/ifup -> ifupdown`), `brctl`, `vconfig`, `ifenslave` (BusyBox / `bridge`, `vlan`, `bonding`) | + `ifupdown-ng-ethtool`, `-iproute2`, `-ppp`, `-wifi`, `-wireguard`, `-wireguard-quick` pulled in | main | yes |
| NTP | BusyBox `ntpd` | — | `chrony 4.8-r7`, `openntpd` on media | yes |
| NetworkManager | no | not installed | `networkmanager 1.52.2-r0`, **community** | no |
| awall (firewall front end) | no | not installed | `awall 1.14.0-r1`, main (Lua, section 8) | no |

- **aarch64:** [repo] the same tool versions exist in `aarch64` `main` (`apkq.py info`: `nftables 1.1.6-r1`,
  `hostapd 2.11-r4`, `openvpn 2.7.7-r0`, `wireguard-tools 1.0.20260223-r0`); the Pi kernel and boot files are
  `linux-rpi 6.18.60-r0` and `raspberrypi-bootloader 1.20260915-r0` in `main`. Not run on arm64.
- **Closure of the NMB tool set** ([repo], `apkq.py closure alpine-base` + the 11 tool packages, first provider for
  virtual names): 62 packages, 27,076,493 bytes installed (x86-64); with `linux-rpi raspberrypi-bootloader
  linux-firmware-brcm` on aarch64: 79 packages, 111,639,452 bytes. The guest got 78 because `alpine-base` was
  already installed with `ifupdown-ng`, `openssl` and the ISO set.

## 8. Python (6A criterion 8)

- **[booted] At boot: no interpreter.** `command -v python python3 perl lua lua5.4 lua5.1` = all absent;
  `find / -xdev \( -iname 'python*' -o -name '*.py' -o -name '*.pyc' \)` found nothing; the same `find` over
  `/.modloop` found 0; `grep -rlI '^#!.*python'` over `/bin /sbin /usr /etc /lib` found nothing.
- **[booted] After `apk add` of the criterion-3 tools: still no interpreter** (runs 1 and 2: `command -v` loop all
  absent; `find / -xdev` for `python*`, `*.py`, `*.pyc` printed nothing; `apk info -v | grep -i -E 'python|perl|lua'`
  printed nothing). One leftover file: **`/sbin/routel`** (from `iproute2`, `grep -rlI '^#!.*python'`) has a Python
  shebang; it cannot run and NMB does not use it.
- **[repo] Closure check:** the dependency closure of `alpine-base` + the 11 tool packages (section 7) contains no
  package with `python`, `perl` or `lua` in its name, on x86_64 and aarch64 (`apkq.py closure`).
- **Interpreters in the base:** BusyBox `ash` and, with `wg-quick`, `bash`. `awall`, Alpine's own firewall front
  end, is **Lua 5.4** (its closure has 14 `lua*` packages), not Python; it is not installed by default.
- **On the media:** `tiny-cloud 3.4.1-r2` (Alpine's cloud-init replacement) is on the ISO; it is shell, not Python
  (no Python package among the 95). `python3 3.14.10-r1` exists in [repo] main, not installed.
- **Can it be left out: yes, and nothing to remove.** No Alpine base package needs Python. The release check of
  criterion 8 must still run on the final image (and `routel` can be dropped).

## 9. Own network and firewall manager, and how each is switched off (6A criterion 5)

What the base brings ([booted]): **ifupdown-ng 0.13.0** (`/sbin/ifup` owned by `ifupdown-ng-0.13.0-r0`) driven by
OpenRC's `networking` service (`/etc/init.d/networking`, from `openrc`); **no `/etc/network/interfaces`** and
`networking` in no runlevel on a fresh diskless boot: `lo` and `eth0` stayed DOWN without addresses. No firewall:
`nftables` is not installed; after installing it, `rc-update show` listed **no** `nftables`, `iptables`, `dnsmasq`,
`hostapd`, `wpa_supplicant`, `openvpn` or `networking` entries (Alpine packages do not enable their services).
No NetworkManager, no netplan, no systemd-networkd, no resolved. `udhcpc` writes `/etc/resolv.conf` when ifupdown-ng
runs DHCP.

- **Default configuration after `setup-interfaces -a -r`** (exit 0): `auto lo`, `iface lo inet loopback`,
  `iface lo inet6 loopback`, `auto eth0`, `iface eth0 inet dhcp`, `iface eth0 inet6 auto`; `udhcpc -b -R -p
  /var/run/udhcpc.eth0.pid -i eth0 -x hostname:localhost` got 192.0.2.15/24.
- **`nftables` service** (`/etc/init.d/nftables` from `nftables-openrc 1.1.6-r1`): `start` runs `nft -f
  /etc/nftables.nft`, `stop` runs **`nft flush ruleset`**, `reload` = `start`; `provide firewall`, `before net`;
  `/etc/conf.d/nftables`: `save_objects="counters limits quotas"`; `save_on_stop` default `no`.
  `/etc/nftables.nft` starts with `flush ruleset` and defines `table inet filter` with **input policy drop**
  (accepts lo, established/related, ICMP/ICMPv6 basics, DHCPv6 replies), **forward policy drop**, output accept,
  then `include "/var/lib/nftables/*.nft"` and `include "/etc/nftables.d/*.nft"` (folder empty).

Coexistence tests in the guest (probe script, exit codes as printed):

| Test | Command | Seen |
| :--- | :--- | :--- |
| Own nft table, then service start | `nft add table inet nmbtest` + chain; `rc-service nftables start` (rc 0) | only `table inet filter`: **NMB's table deleted** (`flush ruleset` in the file) |
| Again, reload | `nft add table inet nmbtest`; `rc-service nftables reload` (rc 0) | only `table inet filter`: **deleted** |
| Again, restart | same; `rc-service nftables restart` (rc 0) | only `table inet filter`: **deleted** |
| Again, stop | same; `rc-service nftables stop` (rc 0) | empty: **everything flushed** |
| Own bridge and address across ifupdown-ng | `ip link add nmbbr0 type bridge`, `198.51.100.77/24`, up; extra `203.0.113.9/24` on `eth0`; then `rc-service networking restart` (rc 0) | `nmbbr0 198.51.100.77/24` **kept**; the extra address on `eth0` **removed** (eth0 was taken down and up again by ifupdown-ng) |
| `ifdown eth0; ifup eth0` | rc 0 | `nmbbr0` kept; `eth0` back to its DHCP address only |
| `rc-service networking stop` | rc 0 | `lo` and `eth0` DOWN, `lo` lost `::1`; `nmbbr0` **kept** with its address |

**What this means for criterion 5.** Both can be switched off cleanly, and both are off by default:

- nftables: not installed and not enabled by default. If installed, keep it out of every runlevel
  (`rc-update del nftables`) or let NMB own `/etc/nftables.nft`; any start, reload, restart or stop wipes all
  tables. `awall` (not installed) would generate the same ruleset; leave it out.
- ifupdown-ng / `networking`: touches only interfaces listed in `/etc/network/interfaces`, and on those it
  removes addresses it did not set. Leave the service out of the runlevels (default on a fresh diskless boot), or
  list no interface NMB manages. `ifupdown-any` is a dependency of `alpine-base`, so the binary stays installed.
- BusyBox `udhcpc`, `ntpd`, `crond`, `acpid`, `watchdog`, `syslogd` exist as init scripts and are not enabled
  except `syslog` (boot runlevel).

## 10. Installed packages

**[booted] 28 at the login prompt** (`apk info -v`): `alpine-base-3.24.4-r0 alpine-baselayout-3.7.2-r1
alpine-baselayout-data-3.7.2-r1 alpine-conf-3.22.0-r0 alpine-keys-2.6-r0 alpine-release-3.24.4-r0 apk-tools-3.0.9-r0
bridge-1.5-r5 busybox-1.37.0-r31 busybox-binsh-1.37.0-r31 busybox-mdev-openrc-1.37.0-r31 busybox-openrc-1.37.0-r31
busybox-suid-1.37.0-r31 ca-certificates-bundle-20261020-r0 ifupdown-ng-0.13.0-r0 libapk-3.0.9-r0 libcap2-2.78-r0
libcrypto3-3.5.9-r0 libssl3-3.5.9-r0 mdev-conf-4.10-r0 musl-1.2.6-r3 musl-utils-1.2.6-r3 openrc-0.63.2-r0
openrc-user-0.63.2-r0 openssl-3.5.9-r0 scanelf-1.3.9-r1 ssl_client-1.37.0-r31 zlib-1.3.2-r1`.

**[booted] 78 after the tool install.** The 50 added: `bash-5.3.9-r1 dbus-libs-1.16.2-r2 dnsmasq-2.92_p2-r0
dnsmasq-common-2.92_p2-r0 dnsmasq-openrc-2.92_p2-r0 ethtool-7.0-r0 gmp-6.3.0-r4 hostapd-2.11-r4 hostapd-openrc-2.11-r4
ifupdown-ng-ethtool-0.13.0-r0 ifupdown-ng-iproute2-0.13.0-r0 ifupdown-ng-ppp-0.13.0-r0 ifupdown-ng-wifi-0.13.0-r0
ifupdown-ng-wireguard-0.13.0-r0 ifupdown-ng-wireguard-quick-0.13.0-r0 iproute2-7.0.0-r0 iproute2-minimal-7.0.0-r0
iproute2-ss-7.0.0-r0 iproute2-tc-7.0.0-r0 iw-6.17-r0 jansson-2.15.0-r0 libcap-ng-0.8.5-r2 libelf-0.195-r0
libmnl-1.0.5-r2 libncursesw-6.6_p20260516-r0 libnftnl-1.3.1-r0 libnl3-3.11.0-r0 libpcap-1.10.7-r0
libxtables-1.8.13-r0 lz4-libs-1.10.0-r1 lzo-2.10-r5 ncurses-terminfo-base-6.6_p20260516-r0 nftables-1.1.6-r1
nftables-openrc-1.1.6-r1 openresolv-3.17.4-r0 openvpn-2.7.7-r0 openvpn-openrc-2.7.7-r0 pcsc-lite-libs-2.4.0-r4
pkcs11-helper-1.31.0-r0 ppp-daemon-2.5.3-r0 ppp-openrc-2.5.3-r0 ppp-pppoe-2.5.3-r0 readline-8.3.3-r1
wireguard-tools-1.0.20260223-r0 wireguard-tools-openrc-1.0.20260223-r0 wireguard-tools-wg-1.0.20260223-r0
wireguard-tools-wg-quick-1.0.20260223-r0 wpa_supplicant-2.11-r4 wpa_supplicant-openrc-2.11-r4 zstd-libs-1.5.7-r2`
(set difference of the two `apk info -v` lists).

## 11. Size on disk

| | Bytes | Source |
| :--- | ---: | :--- |
| x86-64 ISO | 370,147,328 | `stat`; of that modloop 307,982,336, initramfs 22,545,663, kernel 14,607,360, `apks` 13,918,953 |
| Pi archive (`.tar.gz`) | 88,017,396 | `stat` |
| Pi archive unpacked (the FAT boot partition's content) | 110,391,389 | `du -sb`; modloop 52,080,640, kernel 24,218,112, `apks` 14,409,855, initramfs 6,999,726, Pi firmware files 5,387,531, overlays 954,175 |
| x86-64 RAM root at login | 11,640 KiB used of 233,520 KiB tmpfs | guest `df -k /` |
| x86-64 RAM root after the tool install | 31,884 KiB used | guest `df -k /` (apk: "26.8 MiB in 78 packages") |
| Tool-set closure (repo) | 27,076,493 (x86_64), 111,639,452 with kernel + Pi firmware (aarch64) | `apkq.py closure`, sum of `I:` |

- In diskless mode the persistent storage is only: the boot media (read-only), the overlay `*.apkovl.tar.gz` and the
  optional apk cache (section 12 gives their sizes for the tool set).
- Cuts for an NMB image (not measured): most of the x86-64 modloop's 1.49 GB of firmware (GPU firmware `amdgpu`,
  `nvidia`, `radeon`, `i915`, `cirrus` audio ...), via `update-kernel` with a narrower firmware list.

## 12. QEMU boot tests (x86-64 ISO, diskless)

**How it was run.** The ISO itself was booted (read only, `media=cdrom,readonly=on`), with one 64 MiB raw disk
(`truncate` + `mkfs.vfat -n NMBCFG` on the PC; `/dev/vda` in the guest) for the `lbu` overlay and the apk cache.
A serial-console driver (`drive.py`, run with `python3 -I`) waited for `login:`, logged in as `root` (the diskless
ISO has no root password) and ran probe blocks, each ending with a marker and its exit code, then `poweroff`.
Every boot ran under `flock -w 1800 ~/nmb-work/qemu.lock` for its whole run:

```
qemu-system-x86_64 -machine q35 <accel> -smp 1 -m 512 -nographic -monitor none
  [-drive if=pflash,...OVMF_CODE_4M.fd,readonly=on -drive if=pflash,...OVMF_VARS_4M.fd copy]
  -serial unix:serial.sock,server=on,wait=on
  -drive file=alpine-standard-3.24.4-x86_64.iso,media=cdrom,readonly=on,if=ide
  -drive file=cfg.img,format=raw,if=virtio
  -netdev user,id=n0,net=192.0.2.0/24 -device virtio-net-pci,netdev=n0,mac=02:00:00:00:07:01 -no-reboot
```

QEMU 8.2.2 on the PC. No kernel command line was added: the ISO's own `grub.cfg`/`syslinux.cfg` entry booted, and
Alpine's initramfs started a getty on `ttyS0` by itself (`Kernel 6.18.60-0-lts on x86_64 (/dev/ttyS0)`; the
`autodetect_serial` logic in `/init`).

| Run | Settings | Result |
| :--- | :--- | :--- |
| 1 | **KVM** (`-accel kvm -cpu host`, `test -w /dev/kvm` true at the start), UEFI (OVMF), 25 probe blocks | 11:27:57–11:35:38 UTC; driver rc 0, QEMU rc 0 (clean `poweroff`), flock rc 0. Probe `python-after` passed its 120 s driver limit but finished (marker rc 0); `net-default` and `nft-service` ended rc 1 only from a `cat` of absent files and an empty `grep`. The `lbu` block used `/dev/vdb`, but the only virtio disk was `/dev/vda`: **no data from that block**. All other figures below are from this run |
| — | KVM again, 11:36:40 UTC | QEMU: `Could not access KVM kernel module: Permission denied`. `/dev/kvm` had been re-created (time stamp 11:35 UTC) without the user's ACL; `test -w /dev/kvm` was false from then on, so the later runs used **TCG** |
| 2 | TCG (`-accel tcg,tb-size=128 -cpu max`), UEFI | 11:37:33–12:10:20 UTC; login after 418.6 s. The probes fell behind (a `grep -r` over `/usr` ran for minutes under TCG); the boot script stopped QEMU after the driver's limits (driver rc 0, QEMU rc 0 after the kill). Used only for the `apk policy` repository names (section 7) and the Python check after the install (section 8) |
| — | TCG, BIOS (SeaBIOS) | SeaBIOS booted the FAT disk first ("This is not a bootable disk"); stopped by hand, fixed with `-boot order=d`. No data |
| 3 | TCG, BIOS, `-boot order=d`, 9 probe blocks | 12:27:49–12:38:12 UTC; login after 357.4 s; all 9 blocks rc 0; QEMU rc 0, flock rc 0. Tools, `hwsim` with `iw`, cgroups, LSM, `lbu commit` (section 12a) |
| 4 | TCG, BIOS, the same disk after run 3 | 12:38:27–12:46:19 UTC; login after 367.7 s; all 4 blocks rc 0; QEMU rc 0, flock rc 0. Restore from the overlay (section 12a) |

Figures from **run 1 (KVM, UEFI, one vCPU, 512 MB)**:

- **Boot time:** QEMU start to `login:` on the serial console **29.1 s** (driver clock; includes OVMF, the GRUB
  menu's 1 s timeout and the kernel); guest `/proc/uptime` at the first probe **13.62 s**; kernel log
  `Run /init as init process` at **1.70 s**. The diskless boot unpacks the 28 packages from the ISO into RAM in
  that time.
- **Memory** (guest `free -k` at login): total 467,040 kB, used 37,564 kB, free 88,612 kB, buff/cache 340,864 kB,
  available 401,736 kB. Kernel line: `Memory: 392396K/517916K available`. No swap. At the end of run 1 (78 packages
  installed): used 53,996 kB, available 356,152 kB.
- **Storage after boot:** `/` tmpfs 233,520 kB, 11,640 kB used; `/run` tmpfs 93,408 kB; `/dev/shm` 233,520 kB;
  `/dev/sr0` on `/media/sr0` 361,472 kB (ro); `/.modloop` 300,800 kB (ro squashfs); `devtmpfs` 10,240 kB.
- **Listening ports** (guest `netstat -tulnp`): **none**.
- **Running services / processes** (`rc-status -a`, `ps`): `syslogd -t -n`, six `getty` on `tty1`–`tty6`, the
  serial `login`; no `sshd`, no cron, no DHCP client until `setup-interfaces` ran. All sysinit and boot services
  `started`; `firstboot` started (manual).
- **Network at boot:** `lo` and `eth0` DOWN, no `/etc/network/interfaces`, no `/etc/resolv.conf`. After
  `setup-interfaces -a -r`: `eth0` 192.0.2.15/24 from `udhcpc`, default route via 192.0.2.2, DNS 192.0.2.3.
- **Kernel features checked live:** section 4 (`modprobe` loop all 0, `mac80211_hwsim` made 2 radios, Landlock in
  the active LSM list, watchdog devices, `/proc/config.gz` = the ISO's config) and section 9 (nftables and
  ifupdown-ng tests).
- **Not tested in the guest:** `dnsmasq`, `hostapd`, `wpa_supplicant`, `pppd` runs, WireGuard and OpenVPN tunnels;
  only their presence and version.

### 12a. Diskless mode: commit and restore (runs 3 and 4, TCG)

- **Run 3, after `apk add` of the tool set (exit 0, 78 packages):** `mount /media/vda` (rc 0; fstab line
  `/dev/vda /media/vda vfat noauto,rw 0 0` added by the probe), `setup-hostname nmbtest`, a marker file
  `/etc/nmb-marker`, `rc-update add networking boot`, `setup-apkcache /media/vda/cache` (rc 0; `/etc/apk/cache ->
  /media/vda/cache`), `apk cache sync` (rc 0, 36 packages downloaded into the cache: 38 entries, 7,304 KiB), then
  `lbu commit -v vda` (rc 0): "Created /media/vda/nmbtest.apkovl.tar.gz", "Successfully saved apk overlay files".
  - The overlay was **3,604 bytes** with 35 paths (`tar -tzf`): `etc/apk/{arch,cache,repositories,world}`, the three
    `etc/apk/keys/*.pub`, `etc/fstab`, `etc/group`, `etc/hostname`, `etc/inittab`, `etc/network/interfaces`,
    `etc/nmb-marker`, `etc/passwd`, `etc/resolv.conf`, `etc/shadow`, `etc/shells`, their `-` backups, and the
    `etc/runlevels/*` links of every enabled service.
  - A second `lbu commit vda` after changing the marker (rc 0) left **one** file, `nmbtest.apkovl.tar.gz` (3,609
    bytes): the dated backup was removed (`BACKUP_LIMIT` 0, section 3). The disk then had only `cache/` and that file
    (`mdir` on the PC: 59,475,968 bytes free of the 64 MiB).
  - End of run 3: RAM used 42,572 kB of 472,164 kB, available 380,996 kB; `/` tmpfs 31,892 kB used of 236,080 kB.
- **Run 4, same ISO and disk, nothing else changed:** the initramfs found and applied the overlay:
  `hostname` = `nmbtest`, `/etc/nmb-marker` = `nmb-marker-2`, **78** packages installed (`apk info | wc -l`),
  `/etc/apk/world` = `alpine-base dnsmasq ethtool hostapd iproute2 iw nftables openssl openvpn ppp-daemon ppp-pppoe
  wireguard-tools wpa_supplicant`, all seven tools found by `command -v`, `networking` started in the boot runlevel
  with `eth0` 192.0.2.15/24. The kernel command line had no `ip=` option, and the initramfs only brings up a network
  for `ip=`, `nbd=`, or a URL in `apkovl=`/`alpine_repo=` (`/init` line 693), so the packages not on the ISO came
  from the cache on `/dev/vda`.
  - After the boot `/dev/vda` stayed mounted **read-write** on `/media/vda` (`vfat (rw,...)`), the ISO read-only on
    `/media/sr0`, root a tmpfs. RAM used 40,980 kB, available 385,584 kB; `/` 28,896 kB used. No listening ports.
  - `modprobe mac80211_hwsim radios=2` then `iw dev`: 2 interfaces.
- **Run 3 kernel checks with the tools installed:**
  - `iw dev` after `modprobe mac80211_hwsim radios=2` (rc 0): `phy#0`/`wlan0` and `phy#1`/`wlan1`, type managed;
    interface modes `IBSS, managed, AP, AP/VLAN, monitor, P2P-client, P2P-GO, P2P-device, outside context of a BSS,
    NAN`; `iw reg get`: `global`, `country 00: DFS-UNSET`.
  - **cgroups:** no cgroup mount before; `rc-service cgroups start` (rc 0) mounted `cgroup2 on /sys/fs/cgroup type
    cgroup2 (rw,nosuid,nodev,noexec,relatime,nsdelegate)`, controllers `cpuset cpu io memory hugetlb pids dmem`.
    So NMB must enable (or do) the cgroup2 mount itself on Alpine.
  - **Landlock:** boot log `LSM: initializing lsm=lockdown,capability,landlock,yama` and `landlock: Up and running.`;
    `/proc/sys/user/max_user_namespaces` = 1734.
- **TCG boot time** (not a hardware figure): QEMU start to `login:` 357.4 s (run 3) and 367.7 s (run 4, with the
  overlay and 50 packages reinstalled from the cache), BIOS; 418.6 s with UEFI (run 2). Under KVM (run 1) 29.1 s.

## 13. Fields not determined

- **Raspberry Pi archive booted:** not determined. No Pi was used in this package (work plan 5.4) and the PC's QEMU 8.2.2
  has no Pi 4 machine; so Pi boot time, memory, the GENET/Wi-Fi drivers and the Pi firmware's use of `usercfg.txt` were
  read from the files only.
- **arm64 run-time checks** (`modprobe` loop, LSM list, cgroups on the Pi kernel): not determined, same reason.
- **Landlock on the Pi 4 with another Alpine kernel:** [repo] `linux-lts 6.18.63-r0` exists for aarch64 in `main`;
  whether it boots on a Pi 4 and has Landlock was not checked.
- **Boot time and memory on real hardware:** not determined; KVM (run 1) and TCG (runs 2–4) only.
- **Where the x86-64 microcode is:** the release page says the ISO includes AMD and Intel microcode; no
  `microcode` path in the ISO or initramfs lists. Not determined.
- **`update-kernel`, `setup-disk -m sys|data`, `setup-bootable`:** read, not run. Power cut during `lbu commit` or
  `update-kernel`: not tested.
- **Runs of `dnsmasq`, `hostapd`, `wpa_supplicant`, `pppd`, WireGuard and OpenVPN:** not tested; presence and version
  only.
- **`wireless-regdb` package and `iw` regulatory data beyond `regulatory.db` in the modloop:** not checked.
- **Licences of the firmware packages (`L: custom`):** not read.
- **Ookla and RTR binaries for musl (criterion 6):** not part of the image; not checked in this package (left to
  the comparison I-09).
- **Python files after the tool install, run 1:** the `find` printed nothing and the probe ended rc 0, but only after
  the driver's 120 s limit; run 2 repeated it with the same result. Taken as "none" (section 8).

## 14. Tools missing on the PC

- No sudo, so no loop mounts; the ISO and squashfs were read with `7z`, `xorriso -report_el_torito`, `unsquashfs`
  and `cpio -it`, all present now (`cpio`, `bsdtar`, `xorriso` were missing on 2026-11-10 and are installed).
- KVM: usable for run 1 only (section 12); `/dev/kvm` lost the user's access at 11:35 UTC and runs 2–4 used TCG.
- Present and used: `7z`, `fdisk`, `file`, `xorriso`, `unsquashfs` 4.6.1, `cpio`, `tar`, `gpg`, `curl`,
  `mkfs.vfat`, `mtools`, `qemu-system-x86_64` 8.2.2, OVMF (`OVMF_CODE_4M.fd`), `flock`, `python3 -I` (index query
  and serial driver only, both scripts outside the unpacked data).

## 15. Alpine as the NMB base: what the native build adds and what it can drop

The Python version's `setup.sh` (read in `~/src/NetworkMeasureBox`, lines 209–216) already has an Alpine branch:
`apk add --no-cache python3 iproute2 nftables iptables ethtool dnsmasq hostapd iw wireless-regdb ppp curl
ca-certificates busybox-extras libcap traceroute iputils mtr iperf3 sqlite openvpn wireguard-tools build-base gcc
make openssl-dev curl-dev pkgconfig`, plus Ookla and OONI from outside Alpine, and five OpenRC services
(`os/openrc/`: `nmb-db nmb-netd nmb-probe nmb-telegram nmb-web`, `rc-update add ... default`) and a BusyBox crontab.

On top of the diskless base (28 packages, no Python) the native build needs:

- the criterion-3 tools not on the media: `nftables`, `dnsmasq`, `hostapd`, `wireguard-tools` (or only
  `wireguard-tools-wg`, which avoids `bash` and `openresolv` from `wg-quick`), `openvpn`, `iproute2`, `ethtool`;
  `wpa_supplicant`, `iw`, `ppp-daemon`/`ppp-pppoe` are on the media; `wireless-regdb` for `iw` regulatory data
  (not checked here; the modloop carries `regulatory.db`);
- the firmware for supported adapters in the modloop (`update-kernel -p linux-firmware-<x>` or `-e <file>`), because
  a read-only image cannot fetch it later;
- its own OpenRC service for `nmbd` (one instead of five) and, for the A/B update, its own scheme (no RAUC/SWUpdate
  package in Alpine 3.24, section 3);
- on the Pi 4: **no Landlock in `linux-rpi`** (section 4). NMB either runs without Landlock there or uses another
  kernel; the roadmap's "kernel not our own" rule means the stock `linux-rpi` stays, so Landlock must be optional
  in `nmbd`.

It can drop, compared with the Python install: `python3`; `build-base gcc make openssl-dev curl-dev pkgconfig`;
`iptables` (the kernels have `NFT_COMPAT`); `busybox-extras`; `curl` (if nmbd does its own HTTP). From the base:
`tiny-cloud`, `openssh`, `chrony`/`openntpd` (unless NMB wants them), and most of the x86-64 modloop firmware.
Each of these is a proposal for the image package, not tested here.

## 16. 6A criteria for this candidate

| Criterion | Finding | Where |
| :--- | :--- | :--- |
| 1. Stock kernel, Pi 4 and x86-64, network features | Met by the configs, with one gap: kernel 6.18.60 (`lts` x86-64, `rpi` from the Raspberry Pi tree with `BCMGENET`, `PCIE_BRCMSTB`, `BCM2835_WDT`, Pi 4 DTBs and `start4.elf`); bridge, bonding, VLAN, nftables, WireGuard, PPP/PPPoE, cfg80211/mac80211/hwsim and every tethering driver of 4.27 in both. **Landlock is not built into `linux-rpi`** (x86-64 has it active). `NET_TEAM` absent on the Pi. Loaded live on x86-64 only | 4, 5 |
| 2. Security updates for years, free | v3.24 until 2028-06-01 (`releases.json`): main about 2 years, community only to the next release; every criterion-3 tool is in main | 3 |
| 3. Fits nmbd | OpenRC 0.63.2, nftables 1.1.6, dnsmasq 2.92, hostapd/wpa_supplicant 2.11, iw 6.17, pppd 2.5.3, wireguard-tools 1.0.20260223, OpenVPN 2.7.7, iproute2 7.0.0; all from main, none needs Python | 7 |
| 4. Read-only A/B image, reproducible | Diskless mode is a read-only base by design (squashfs modloop, RAM root, config in one `apkovl` written by rename; commit and restore worked in runs 3–4). The kernel/modloop update is an in-place `mv` with no second slot or fallback; no RAUC/SWUpdate package. A/B would be NMB's own | 2, 3, 12 |
| 5. Does not fight NMB | Nothing enabled by default: no firewall installed, `networking` in no runlevel, packages do not enable services. nftables service wipes all tables on start/reload/restart/stop; ifupdown-ng removes foreign addresses only on interfaces it manages; NMB's bridge survived | 9 |
| 6. Size, memory, musl | ISO 370 MB (modloop 308 MB, mostly firmware); Pi boot content 110 MB; 37.6 MB RAM used and 11.6 MB root after boot in a 512 MB KVM guest, 41 MB used with the tool set restored from the overlay (TCG); tool closure 27 MB; musl 1.2.6, so Ookla/RTR need musl or static builds (not checked here) | 11, 12 |
| 7. Licence | Base packages free licences (GPL/MIT/BSD/ISC ...); firmware packages `custom` | 6 |
| 8. No Python in the base | Met: no interpreter at boot or after the tool set, none in the dependency closure; only `/sbin/routel` (iproute2) has a Python shebang | 8 |
