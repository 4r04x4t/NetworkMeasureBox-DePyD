# Image record: Debian 13.7.0 "trixie" (amd64 and arm64 netinst)

Package I-06 (work plan section 5.4). The method is work plan 5.1, and every field of 5.2 is below.
Each value names the command that produced it. A value that could not be read says "not determined" and why.
Nothing here comes from memory. Only lists, versions and checksums are recorded. No file taken out of the image
is in this repository.

- Date of the analysis: 2026-11-11.
- Where it ran: dev PC, Windows with WSL2 Ubuntu 24.04, about 1.8 GB RAM and 4 threads, no root. The QEMU guest ran
  on the same PC.
- Work folder: `~/nmb-work/I-06/` (outside every repository; one fresh sub-folder per unpack).
- Material, read only: `~/src/NMB-depythonization-firmwares/platforms/x86-64/` and `.../raspberry-pi-4-arm64/`.
- **Two kinds of source, kept apart** (work plan 5.1: an installer ISO is not the installed system):
  - **[ISO]** values read from the two netinst ISOs: their package indexes, the kernel package
    (`linux-image-6.12.107+deb13-<arch>`), the firmware folder and the boot files.
  - **[built]** values read from minimal systems built with `mmdebstrap` 1.4.3 (`--mode=unshare`, no root) from the
    signed Debian mirror `http://deb.debian.org/debian`, suite `trixie` (mirror `Release` says `Version: 13.7`,
    `Date: Sat, 12 Sep 2026`, the same point release as the ISOs). Three builds:
    - **A** amd64, `--variant=important` (Essential + Priority required + important; the mmdebstrap manual calls this
      "the default of debootstrap"), nothing else.
    - **B** amd64, A plus `linux-image-amd64` and the tools of criterion 3 (`dnsmasq-base hostapd wpasupplicant iw
      ppp wireguard-tools openvpn ethtool`). B was also booted in QEMU (section 12).
    - **C** arm64, the same set as B plus `raspi-firmware` and `firmware-brcm80211` (component
      `non-free-firmware`), built through `qemu-aarch64` user emulation (binfmt). Not booted.
  - mmdebstrap does not install Recommends. Section 6 says what the Recommends would add.

## 1. Files and checksums

| File | Size (bytes) | SHA-256 (publisher list) | Check |
| :--- | ---: | :--- | :--- |
| `debian-13.7.0-amd64-netinst.iso` | 792,723,456 | `a7ef94ac2fb9a7fec454552abd629b7cc9d5155c886165a45649f5ce6167e355` | OK |
| `debian-13.7.0-arm64-netinst.iso` | 735,358,976 | `6e93fa1759bd9d4b0fc11e938987de6967ee7de5297dac1be27c3a75cc17024b` | OK |

- **How it was checked.** The set's README lists the Debian files but not their sums. The publisher's lists were
  fetched on 2026-11-11 at 07:03 UTC from `https://cdimage.debian.org/debian-cd/13.7.0/{amd64,arm64}/iso-cd/SHA256SUMS`
  with `SHA256SUMS.sign` (HTTP 200 on the `debian-cd/13.7.0` path).
  - List SHA-256: amd64 `de6d320c52b0fa02787e4a6f58115ace89f821865baaeba86115c723f34cbab4` (3 lines: netinst,
    edu, mac), arm64 `6f768053b16f25cd8de528fc716bcbd43a365e70b44a09cb02d86e24cc2f7618` (1 line).
  - `grep " debian-13.7.0-amd64-netinst.iso$" SHA256SUMS | sha256sum -c` and `sha256sum -c SHA256SUMS` (arm64), run
    in the material folders: both `OK`, exit 0.
- **Signature of the lists.** `gpg --verify SHA256SUMS.sign SHA256SUMS` in an isolated `GNUPGHOME`: "Good signature
  from Debian CD signing key <debian-cd@lists.debian.org>" for both (exit 0), made 2026-09-12 18:20:46 / 18:20:49 CEST.
  - Key `DF9B 9C49 EAA9 2984 3258 9D76 DA87 E80D 6294 BE9B` (rsa4096, 2011-01-05), fetched from `keyring.debian.org`.
    The same fingerprint is printed on `https://www.debian.org/CD/verify` (fetched the same minute). gpg says "not
    certified with a trusted signature", because no web of trust was set up.
- **Inside the ISO.** The files taken out (section 2) were checked against the ISO's own `md5sum.txt`
  (`md5sum -c` on the 55 extracted files per ISO that the list covers): 55 of 55 `OK` for each ISO.
- **The mirror used for the builds.** apt inside mmdebstrap checked `InRelease` against the keyring taken from the
  ISO (`debian-archive-keyring_2025.1_all.deb`, covered by the md5 list above). The `Release` file was also checked
  by hand: `gpgv --keyring debian-archive-keyring.pgp Release.gpg Release` gave three good signatures (12/bookworm
  and 13/trixie automatic signing keys, 13/trixie stable release key `4158 7F7D B8C7 74BC CF13 1416 762F 67A0 B2C3 9DE4`),
  exit 0. `main/binary-amd64/Packages.xz` SHA-256 `7778d3e3…2498` equals the value in that `Release`.
- **Release identity.** `.disk/info`: `Debian GNU/Linux 13.7.0 "Trixie" - Official amd64 NETINST with firmware
  20260912-09:35` and `... arm64 NETINST with firmware 20260912-09:36`. `dists/trixie/Release`: `Suite: stable`,
  `Version: 13.7`, `Description: Debian 13.7 Released 12 September 2026`. Built roots: `/etc/debian_version` = `13.7`.

## 2. Partition and file system layout, boot method

**[ISO]** The ISOs were listed with `7z l -slt` (exit 0) and read with `fdisk -l` (exit 0) and `file`. Nothing was
mounted; selected files were extracted with `7z x -o<fresh dir>` (exit 0).

| | amd64 | arm64 |
| :--- | :--- | :--- |
| `file` | ISO 9660, volume `Debian 13.7.0 amd64 n`, DOS/MBR boot sector, bootable | ISO 9660, volume `Debian 13.7.0 arm64 n`, DOS/MBR boot sector, bootable |
| MBR (`fdisk -l`) | p1 type 0 "Empty", sectors 0–1548287 (756 MiB, the whole ISO, boot flag); p2 type `ef` EFI, sectors 4116–11475 (3.6 MiB) | p1 type 83, sectors 0–1427455 (697 MiB); p2 type `ef` EFI, sectors 1427456–1435647 (4 MiB) |
| `xorriso` options (`.disk/mkisofs`) | `-isohybrid-mbr`, El Torito BIOS (`isolinux`) + `-eltorito-alt-boot -e` EFI image, `-isohybrid-gpt-basdat`, `-isohybrid-apm-hfsplus` | `-e` EFI image only, `-append_partition` (the EFI partition) |
| Boot paths | BIOS: `isolinux/isolinux.bin`; UEFI: `EFI/boot/bootx64.efi` + `grubx64.efi`, `boot/grub/efi.img` | UEFI only: `EFI/boot/bootaa64.efi` + `grubaa64.efi`, `boot/grub/efi.img` |
| Installer kernel (`file install.*/vmlinuz`) | bzImage `6.12.107+deb13-amd64 ... #1 SMP PREEMPT_DYNAMIC Debian 6.12.107-1 (2026-08-29)` | ARM64 `Image`, little-endian, 4K pages |
| Paths in the ISO (`7z l`) | 1,485 | 1,622 |

- **Raspberry Pi 4 and this ISO.** The arm64 ISO's only boot path is EFI (`bootaa64.efi`). The Pi 4's own boot
  files (`start4*.elf`, `fixup4*.dat`, in `raspi-firmware_1.20240424+ds-6_all.deb`, which is in the ISO's
  `pool/non-free-firmware`) and the device trees (`bcm2711-rpi-4-b.dtb`, `bcm2711-rpi-400.dtb`, `bcm2711-rpi-cm4-io.dtb`
  in the kernel package, `dpkg-deb -c`) are carried as packages. Whether the netinst installer runs on a Pi 4 was
  not determined (no Pi was used, work plan 5.4; and the ISO cannot be booted under QEMU as a Pi). `flash-kernel 3.109`
  and `u-boot-tools` are in the arm64 pool.
- **[built] Layout of an installed system.** Not given by the ISO; the installer partitions at install time. The
  booted build B used one ext4 file system on `/dev/vda` (made with `mke2fs -t ext4 -d`, 2 GiB, see section 12),
  `/tmp` as tmpfs (`tmp.mount`, static in systemd 257) and no swap.
- **[built] Boot chain of an installed system:** kernel `/boot/vmlinuz-6.12.107+deb13-<arch>` with an initramfs built
  by `initramfs-tools` 0.148.4 at install time (B: `initrd.img` 36,356,871 bytes; `update-initramfs` ran in the
  build). The guest booted it with QEMU's direct kernel boot, so GRUB was not exercised.
- **[built] C: the Pi 4 boot path without UEFI.** Installing `raspi-firmware` in build C ran its hooks, which filled
  `/boot/firmware/` with 35 entries, 97,105,598 bytes (`tar -tvf`): the GPU firmware (`start4.elf`, `start4x.elf`,
  `start4db.elf`, `fixup4*.dat`, ...), the device trees (`bcm2711-rpi-4-b.dtb`, `bcm2711-rpi-400.dtb`,
  `bcm2711-rpi-cm4-io.dtb`, and Pi 3/5 ones), **copies** of `vmlinuz-6.12.107+deb13-arm64` (37,660,608 bytes) and
  `initrd.img-6.12.107+deb13-arm64` (36,930,600 bytes), and:
  - `config.txt` (non-comment lines, `tar -xOf`): `arm_64bit=1`, `enable_uart=1`, `upstream_kernel=1`,
    `kernel=vmlinuz-6.12.107+deb13-arm64`, `initramfs initrd.img-6.12.107+deb13-arm64`;
  - `cmdline.txt`: `console=tty0 root=/dev/mmcblk0p2 rw fsck.repair=yes net.ifnames=0 cma=64M rootwait`.
  - So the Debian kernel boots on the Pi 4 through the Pi's own firmware from a FAT `/boot/firmware` partition and a
    root on partition 2; the ISO's EFI path is not needed. This was read from the files, not booted.

## 3. How an update is written

- **apt/dpkg in place.** **[built]** `apt 3.0.3` and `dpkg` are in A, B and C. An update replaces files on the
  running root; there is one root file system and no A/B scheme in either ISO or build.
- **Timers enabled by default** (**[built]** B, `tar -tvf` of `etc/systemd/system/timers.target.wants/`):
  `apt-daily.timer`, `apt-daily-upgrade.timer`, `dpkg-db-backup.timer`, `e2scrub_all.timer`, `fstrim.timer`,
  `logrotate.timer`. In the guest `systemctl is-enabled apt-daily.timer apt-daily-upgrade.timer` = enabled, enabled.
  Without `unattended-upgrades` (not installed, not on the ISO) these timers download lists but install nothing.
- **Building blocks for an NMB image with A/B updates, in trixie main** (`Packages.xz` of the signed mirror,
  `awk` over the index): `rauc 1.13-3+deb13u1`, `swupdate 2024.12.1+dfsg-3+deb13u2`, `systemd-repart` and
  `systemd-boot 257.13-1~deb13u1`, `mmdebstrap 1.5.7-1+deb13u1`, `debos 1.1.5-1+deb13u1`, `mkosi 25.3-7`,
  `ostree 2025.2-1`. All Priority optional, none installed in A/B/C. Not tested in this package.
- **Security support** (criterion 2): `https://www.debian.org/releases/trixie/` (fetched 2026-11-11) says full
  Debian support until 2028-08-09 and LTS until 2030-06-30; `https://wiki.debian.org/LTS` gives
  "Debian 13 trixie ... 2028-08-09 - 2030-06-30" with the LTS architecture list "to be determined". No subscription.

## 4. Kernel and the configuration NMB needs

- **Kernel version: 6.12.107** (Debian `6.12.107-1`, ABI `6.12.107+deb13`), the same in both ISOs and in the
  mirror. Sources:
  - **[ISO]** `dpkg-deb -f` on `pool/main/l/linux-signed-<arch>/linux-image-6.12.107+deb13-<arch>_6.12.107-1_<arch>.deb`:
    `Version: 6.12.107-1`, `Source: linux-signed-<arch> (6.12.107+1)`, `Depends: kmod, initramfs-tools (>= 0.120+deb8u2)
    | linux-initramfs-tool`, `Recommends: apparmor`. Installed-Size 108,326 KiB (amd64) and 173,611 KiB (arm64).
    SHA-256 of the .deb: amd64 `7794643c…9e99` (107,898,972 bytes), arm64 `20c92f3f…9fc9` (92,902,468 bytes).
  - **[built]** guest `uname -a`: `Linux nmb-i06 6.12.107+deb13-amd64 #1 SMP PREEMPT_DYNAMIC Debian 6.12.107-1 (2026-08-29) x86_64`.
    Boot log: `x86_64-linux-gnu-gcc-14 (Debian 14.2.0-19) 14.2.0, GNU ld (GNU Binutils for Debian) 2.44`.
- **Where the config came from:** `/boot/config-6.12.107+deb13-<arch>` inside the kernel package, taken out with
  `dpkg-deb --fsys-tarfile | tar -x ./boot/config-*` (tar reported `Not found in archive` only for a second
  pattern, `./lib/modules/*/modules.builtin`; the config and `usr/lib/modules/*/modules.builtin` were extracted).
  - amd64: 11,466 lines, `# Linux/x86 6.12.107 Kernel Configuration`, SHA-256 `9b460f1003db4d417288e9c14caecfa14245dab07186df948fef0f3661e38ceb`.
  - arm64: 13,095 lines, SHA-256 `21b2f85fedc3801307af4a27a6ea15bbfdc16cd92d7a0211a5ead1f01dcc3cee`.
  - `CONFIG_IKCONFIG` is not set: there is no `/proc/config.gz`; the config file in `/boot` is the source on an
    installed system too (B has it, `ls /boot`).
- **Source of every option below: those two files** (`grep -E "^CONFIG_X=|^# CONFIG_X is not set"`). "Module" says
  whether the `.ko.xz` is in the kernel package (`grep -c "/<name>.ko.xz$"` over `dpkg-deb -c`). "Guest" is the
  `modprobe <name>; echo $?` loop in the booted amd64 build B (section 12).

| Feature (5.2) | Option(s) | amd64 | arm64 | Module in package | Guest (amd64) |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Bridge, VLAN filtering | `BRIDGE`, `BRIDGE_VLAN_FILTERING` | m, y | m, y | `bridge` both | modprobe 0 |
| Bonding | `BONDING` | m | m | `bonding` both | 0 |
| Team | `NET_TEAM` | m | m | — | — |
| VLAN 802.1Q | `VLAN_8021Q` | m | m | `8021q` both | 0 (`802.1Q VLAN Support v1.8`) |
| nftables | `NF_TABLES` m, `NF_TABLES_INET` y, `NF_TABLES_NETDEV` y, `NF_TABLES_BRIDGE` m, `NFT_NAT` m, `NFT_MASQ` m, `NFT_FIB_INET` m, `NFT_CT` m, `NF_FLOW_TABLE` m, `NFT_FLOW_OFFLOAD` m, `NFT_COMPAT` m | as listed | as listed | `nf_tables`, `nft_nat`, `nft_masq`, `nft_fib_inet` both | 0; `nft` worked (section 9) |
| iptables (legacy) | `IP_NF_IPTABLES` | m | m | — | — |
| WireGuard | `WIREGUARD` | m | m | `wireguard` both | 0 (`WireGuard 1.0.0 loaded`) |
| PPP, PPPoE | `PPP`, `PPPOE`, `PPP_ASYNC`, `PPP_SYNC_TTY`, `L2TP` | all m | all m | `ppp_generic`, `pppoe`, `pppox` both | `pppoe` 0 |
| cfg80211, mac80211 | `CFG80211`, `MAC80211` | m, m | m, m | both | 0, 0 (cfg80211 loads its built-in regulatory X.509 certificates) |
| mac80211_hwsim | `MAC80211_HWSIM` | m | m | both | 0; `iw dev` then showed 2 interfaces |
| Wi-Fi drivers (examples) | `ATH9K`, `ATH9K_HTC`, `ATH10K_PCI`, `IWLWIFI`, `MT7601U`, `MT76x2U`, `MT7921U`, `RTL8XXXU`, `RT2800USB`, `BRCMFMAC` (+`BRCMFMAC_SDIO` y) | all m | all m | 221 (amd64) / 240 (arm64) `.ko` under `drivers/net/wireless/` | — |
| USB host | `USB`, `USB_XHCI_HCD`, `USB_EHCI_HCD`, `USB_XHCI_PCI`; `USB_DWC2` | m (DWC2 not set) | m (DWC2 m) | — | — |
| USB network | `USB_USBNET`, `USB_RTL8152`, `USB_NET_AX88179_178A`, `USB_NET_AX8817X`, `USB_NET_SMSC95XX`, `USB_LAN78XX`, `USB_NET_AQC111` | all m | all m | `usbnet`, `r8152`, `ax88179_178a`, `smsc95xx`, `lan78xx` both | — |
| USB tethering (4.27) | `USB_NET_RNDIS_HOST`, `USB_NET_CDCETHER`, `USB_NET_CDC_NCM`, `USB_IPHETH`, `USB_NET_QMI_WWAN`, `USB_NET_CDC_MBIM`, `USB_NET_HUAWEI_CDC_NCM`, `USB_NET_CDC_EEM` | all m | all m | `rndis_host`, `cdc_ether`, `cdc_ncm`, `ipheth`, `qmi_wwan`, `cdc_mbim`, `huawei_cdc_ncm` both | `rndis_host cdc_ether cdc_ncm ipheth qmi_wwan cdc_mbim` all 0 |
| TUN, veth, macvlan, ipvlan, dummy, IFB | `TUN`, `VETH`, `MACVLAN`, `IPVLAN`, `DUMMY`, `IFB` | all m | all m | all present both | `veth` 0 |
| Policy routing | `IP_ADVANCED_ROUTER`, `IP_MULTIPLE_TABLES`, `IPV6_MULTIPLE_TABLES` | y | y | built in | — |
| Traffic control | `NET_SCHED` y, `NET_SCH_FQ_CODEL` y, `DEFAULT_NET_SCH="fq_codel"`, `NET_SCH_CAKE`, `NET_SCH_HTB`, `NET_SCH_NETEM`, `NET_SCH_TBF`, `NET_SCH_FQ`, `NET_SCH_INGRESS` m | as listed | as listed | `sch_cake`, `sch_netem` both | — |
| Namespaces | `NAMESPACES`, `NET_NS`, `PID_NS`, `USER_NS`, `UTS_NS`, `IPC_NS`, `TIME_NS` | all y | all y | built in | — (the PC built the roots in user namespaces, but that is the PC's kernel) |
| seccomp | `SECCOMP`, `SECCOMP_FILTER` | y, y | y, y | built in | systemd banner `+SECCOMP` |
| **Landlock** | `SECURITY` y, `SECURITY_LANDLOCK` y, `LSM="landlock,lockdown,yama,loadpin,safesetid,integrity,apparmor,selinux,smack,tomoyo,bpf,ipe"` | y | y | built in | **active**: guest `/sys/kernel/security/lsm` = `lockdown,capability,landlock,yama,apparmor,tomoyo,bpf,ipe,ima,evm` |
| Other LSMs | `SECURITY_APPARMOR` y (`DEFAULT_SECURITY_APPARMOR` y), `SECURITY_SELINUX` y, `SECURITY_YAMA` y, `SECURITY_LOCKDOWN_LSM` y, `LOCK_DOWN_KERNEL_FORCE_NONE` y | as listed | as listed | — | AppArmor active (list above); the `apparmor` userspace package is a Recommends and is not in B |
| cgroups | `CGROUPS`, `MEMCG` (`MEMCG_V1` not set), `CGROUP_PIDS`, `CPUSETS`, `CGROUP_BPF`, `BLK_CGROUP`, `CGROUP_FREEZER`, `CGROUP_DEVICE` | y (V1 not set) | y (V1 not set) | — | cgroup2 on `/sys/fs/cgroup` (`nsdelegate,memory_recursiveprot`, guest `mount`) |
| Watchdog | `WATCHDOG` y; `SOFT_WATCHDOG` m, `I6300ESB_WDT` m; `ITCO_WDT` m (amd64 only); `BCM2835_WDT` m (arm64 only) | as listed | as listed | `softdog`, `i6300esb` both; `iTCO_wdt` amd64; `bcm2835_wdt` arm64 | `iTCO_wdt` auto-loaded on the q35 machine (`Found a ICH9 TCO device ... heartbeat=30 sec`), `/dev/watchdog` and `/dev/watchdog0` present; systemd `RuntimeWatchdogUSec=0`, `RebootWatchdogUSec=10min` |
| Module signing | `MODULE_SIG` y, `MODULE_SIG_FORCE` not set, `MODULE_COMPRESS_XZ` y | as listed | as listed | `.ko.xz` | — |
| Raspberry Pi 4 platform | `ARCH_BCM2835` y, `ARCH_BCM` y, `PCIE_BRCMSTB` y, `RASPBERRYPI_FIRMWARE` y, `RASPBERRYPI_POWER` y, `RESET_RASPBERRYPI` m, `BCMGENET` m, `MMC_SDHCI_IPROC` m, `DRM_VC4` m, `HW_RANDOM_BCM2835` m, `SERIAL_8250_BCM2835AUX` y, `SENSORS_RASPBERRYPI_HWMON` not set | absent | as listed | `genet`, `bcm2835_wdt`, `bcm2711_thermal`; DTBs `bcm2711-rpi-4-b`, `-rpi-400`, `-rpi-cm4-io` | not booted |
| Other | `BPF_SYSCALL` y, `XDP_SOCKETS` y, `TCP_CONG_BBR` m, `MPTCP` y, `TLS` m, `IPV6` y, `PREEMPT_VOLUNTARY` y (`PREEMPT_DYNAMIC` y on amd64 only), `HZ=250`, `NR_CPUS` 8192 / 512, `OVERLAY_FS` m, `SQUASHFS` m, `EROFS_FS` m, `DM_VERITY` m, `EFI_STUB` y | as listed | as listed | — | — |

- `grep CGROUP_NS` found no such option in either config; cgroup namespaces were not tested separately.

## 5. Modules and firmware files carried

- **Modules, [ISO] kernel package** (`dpkg-deb -c`, `.ko.xz` files): amd64 4,229 (93,653,092 bytes compressed),
  arm64 4,350 (78,809,484 bytes). `modules.builtin`: 191 lines (amd64), 465 (arm64). arm64 also carries 1,023
  `.dtb` files. The kernel image is 12,142,528 bytes (amd64 `vmlinuz`) and 37,660,608 bytes (arm64, uncompressed
  `Image` as `vmlinuz`).
- **Modules, [built] B:** `/usr/lib/modules/6.12.107+deb13-amd64` holds the same 4,229 files, 99,582,881 bytes
  (`tar -tvf` sum). The guest had 66 lines in `lsmod` after boot (65 modules), among them `virtio_net`,
  `virtio_blk`, `iTCO_wdt`, `drm`, `bochs`, `ext4`; no netfilter module was loaded at boot.
- **Installer modules [ISO]:** `*-modules-6.12.107+deb13-<arch>-di` udebs (`nic-modules`, `nic-usb-modules`,
  `nic-wireless-modules`, `ppp-modules`, ...): 230 udebs in the amd64 index, 246 in arm64.
- **Firmware [ISO]:** the ISO carries a `firmware/` folder with 36 `.deb` packages per architecture and
  `firmware/Contents-firmware` (4,735 lines amd64, 4,291 arm64) mapping each firmware file to its package.
  - Both: `firmware-atheros`, `-brcm80211`, `-iwlwifi`, `-mediatek`, `-realtek`, `-misc-nonfree`, `-ath9k-htc`,
    `-carl9170`, `-libertas`, `-zd1211`, `-linux-free`, `-bnx2`, `-bnx2x`, `-intel-graphics`, `-amd-graphics`, ...
    (all `20250410-2` from the firmware tree, except the older single packages).
  - amd64 only: `amd64-microcode 3.20250311.1`, `intel-microcode 3.20251111.1~deb13u1`, `firmware-intel-sound`,
    `firmware-sof-signed`. arm64 only: `raspi-firmware 1.20240424+ds-6`, `firmware-qcom-soc`, `firmware-samsung`,
    `firmware-ti-connectivity`.
  - Pi 4 Wi-Fi: `Contents-firmware` maps `brcm/brcmfmac43455-sdio.bin`, `.clm_blob`, `.raspberrypi,4-model-b.txt`
    and `.raspberrypi,4-compute-module.txt` to `firmware-brcm80211_20250410-2_all.deb` (component `non-free-firmware`).
- **Firmware [built]:** A and B carry no `/usr/lib/firmware` (absent in both). The base gets firmware only from
  `firmware-*` packages. C (with `firmware-brcm80211`): 79 files, 18,741,563 bytes under `usr/lib/firmware`
  (`brcm/` and `cypress/`); the Pi GPU firmware is in `usr/lib/raspi-firmware` and `/boot/firmware` (section 2).
- **Modules, [built] C:** 4,350 files in `/usr/lib/modules/6.12.107+deb13-arm64`, 84,970,456 bytes.

## 6. Init system, service manager, C library, package manager

- **Init: systemd 257** (`systemd 257.13-1~deb13u1`, `systemctl --version` in the build; boot banner
  `systemd 257.13-1~deb13u1 running in system mode (+PAM +AUDIT +SELINUX +APPARMOR +IMA +IPE +SMACK +SECCOMP ...)`).
  `/sbin/init -> ../lib/systemd/systemd` (`systemd-sysv`, Priority important). Default target `graphical.target`
  (`systemctl get-default` in A), reached in the guest after `multi-user.target`.
- **Units enabled by default, [built] B** (`tar -tvf`, symlinks under `etc/systemd/system/*.wants/`):
  - `multi-user.target.wants`: `cron`, `e2scrub_reap`, `networking` (ifupdown), **`hostapd`**, **`wpa_supplicant`**,
    **`openvpn`**, `remote-fs.target`. The three bold ones are enabled by their packages on install. In the guest
    hostapd did not run (`ConditionFileNotEmpty=/etc/hostapd/hostapd.conf`), `wpa_supplicant.service` is `Type=dbus`
    and did not run (no D-Bus daemon in B), `openvpn.service` is a `Type=oneshot` `ExecStart=/bin/true`.
  - also `getty@tty1`, `systemd-pstore`, the six timers of section 3, and the alias
    `dbus-fi.w1.wpa_supplicant1.service`.
  - A has the same set without hostapd, wpa_supplicant and openvpn (`systemctl list-unit-files --state=enabled` in A).
- **C library: glibc 2.41** (`ldd (Debian GLIBC 2.41-12+deb13u4) 2.41`, in A and B; `libc6 2.41-12+deb13u4`
  on both ISOs and in C).
- **Package manager: apt 3.0.3 / dpkg** (A, B, C). Kernel modules are signed and xz-compressed (section 4).
- **Package counts.**
  - [ISO] amd64 index: 461 debs in `main`, 3 in `contrib`, 31 in `non-free-firmware`, 230 udebs; priorities in
    `main`: 33 required, 31 important, 22 standard, 373 optional, 2 extra. arm64: 429 / 3 / 31 / 246 udebs;
    33 required, 31 important, 22 standard, 342 optional, 1 extra (`zcat Packages.gz | grep -c '^Package: '`).
  - [built] A amd64 `important`: **144** installed (39 required, 30 important, 75 optional, from
    `var/lib/dpkg/status`). B amd64 with kernel and tools: **171**. C arm64 with the Pi firmware: **174** (section 10).
- **Recommends that mmdebstrap left out** (`Recommends:` in B's dpkg status): `dnsmasq-base: dns-root-data`,
  `iw: wireless-regdb`, `openvpn: easy-rsa`, `wireguard-tools: nftables | iptables, wireguard-modules | wireguard-dkms`,
  `systemd: default-dbus-system-bus | dbus-system-bus, linux-sysctl-defaults, systemd-timesyncd | time-daemon,
  systemd-cryptsetup`, `ifupdown: dhcpcd-base | dhcp-client`, `nftables: netbase`, `linux-image: apparmor`. An
  `apt-get install` with the default `APT::Install-Recommends` would pull in those not already installed. None of them is a Python package
  by name; their own dependencies were not resolved here.
- **Licences (criterion 7).** Sections in B (`Section:` of the status file): 85 libs, 30 admin, 24 utils, 11 net,
  5 perl, ...; every package in A and B comes from `main`. The Pi 4 boot files and Wi-Fi firmware are in
  `non-free-firmware` (licences per package, not read in this package).

## 7. Tools of 6A criterion 3, with versions

Presence: `command -v` in the build root (`chroot`), and again in the booted guest for B. Versions: the tool's own
`--version`/`-v` output in the build root, and the package version from the dpkg status. "ISO" says whether the
package is in the ISO's `main` pool (`awk` over `Packages.gz`).

| Tool | A (important) | B (amd64) version | ISO | Package (B) |
| :--- | :--- | :--- | :--- | :--- |
| systemd (init) | yes | `systemd 257 (257.13-1~deb13u1)` | yes | `systemd 257.13-1~deb13u1` |
| OpenRC, procd | no | — | — | — |
| nftables `nft` | **yes** (Priority important) | `nftables v1.1.3 (Commodore Bullmoose #4)` | yes | `nftables 1.1.3-1` |
| dnsmasq | no | `Dnsmasq version 2.91` | no | `dnsmasq-base 2.91-1+deb13u2` (binary only, no service; `dnsmasq` with its unit is a separate package) |
| hostapd | no | `hostapd v2.10` | no | `hostapd 2:2.10-24` |
| wpa_supplicant | no | `wpa_supplicant v2.10` | yes | `wpasupplicant 2:2.10-24` |
| iw | no | `iw version 6.9` | no | `iw 6.9-1` |
| pppd | no | `pppd version 2.5.2` | yes | `ppp 2.5.2-1+1` (PPPoE plugin in the same package) |
| WireGuard `wg`, `wg-quick` | no | `wireguard-tools v1.0.20210914` | no | `wireguard-tools 1.0.20210914-3`; module in the kernel |
| OpenVPN | no | `OpenVPN 2.6.14 x86_64-pc-linux-gnu [SSL (OpenSSL)] [LZO] [LZ4] [EPOLL] [PKCS11] [MH/PKTINFO] [AEAD] [DCO]` | no | `openvpn 2.6.14-1+deb13u3` |
| `ip`, `tc`, `ss`, `bridge` | **yes** | `ip utility, iproute2-6.15.0, libbpf 1.5.0` (full iproute2, not BusyBox) | yes | `iproute2 6.15.0-1` |
| ethtool | no | `ethtool version 6.14.2` | no | `ethtool 1:6.14.2-1` |
| iptables | no | absent | — | — |
| DHCP client | `dhcpcd` (`dhcpcd-base 1:10.1.0-11+deb13u4`, Priority important) | same | yes | — |
| ifupdown | yes | `ifupdown 0.8.44+deb13u1` | yes | — |
| systemd-networkd | yes (`networkctl`, part of `systemd`) | unit present, disabled | — | — |
| NetworkManager, netplan, busybox, curl, wget, ssh | no | no | network-manager, netplan not on ISO | — |
| perl | yes, `/usr/bin/perl` from `perl-base 5.40.1-6+deb13u1` (Essential) | yes | — | — |

**arm64 (build C):** the same tools are present (`command -v` loop in the build root, all found; `python3` absent)
with the same `--version` output, except `OpenVPN 2.6.14 aarch64-unknown-linux-gnu [...]`. `nft --version` failed
under user emulation (`src/mnl.c:64: Unable to initialize Netlink socket: Protocol not supported`, a limit of
`qemu-aarch64`, not of the package); the package is `nftables 1.1.3-1`. `iw` is the rebuild `6.9-1+b1`.
C enables the same units as B (`tar -tvf`, `etc/systemd/system/*.wants/`).

## 8. Python (6A criterion 8)

- **[built] A, B and C: no Python interpreter.** `command -v python3 python` = absent in the build roots and in the
  booted guest; `usr/lib/python3*` absent; the guest's `find / -xdev -name 'python3*' -path '*bin*'` found nothing.
- Leftover Python *files* without an interpreter (`find -name '*.py' -o -name '*.pyc'` and
  `grep -rlI '^#!.*python'`):
  - `/usr/bin/routel` from `iproute2` has `#! /usr/bin/env python3` (it cannot run; `routel` is a helper, NMB does
    not use it).
  - `/usr/share/gcc/python/libstdcxx/...` (4 files, gdb pretty-printers from `libstdc++6`),
    `/usr/share/gdb/auto-load/.../libstdc++.so.6.0.33-gdb.py`, `/usr/share/nano/python.nanorc` (syntax file), and in
    B `/usr/share/doc/openvpn/examples/sample-scripts/totpauth.py`. Five files in A, six in B.
  - These are data files; an image build can drop them (`dpkg` path-exclude) or leave them.
- **[ISO] Python on the ISO:** `python3 3.13.5-1`, `python3.13 3.13.5-2+deb13u5` and their libraries are in the pool
  (Priority optional). In the amd64 index, the packages that depend on `python3` are the Python packages themselves
  plus `xfsprogs`, `python3-apt`, `python3-debian`, `python3-reportbug` and **`reportbug`, Priority standard**.
  So the installer's "standard system utilities" task brings Python; a base without that task does not (A shows it).
- `unattended-upgrades` and `networkd-dispatcher` (Python, named in criterion 8) are not on the ISO and are not in
  A/B; in the mirror they are Priority optional (`unattended-upgrades 2.12`, `networkd-dispatcher 2.2.4-1.1`).
- **Can it be left out: yes.** A Debian base built with `--variant=important` (or `minbase`) plus NMB's tools has
  no interpreter, and none of the criterion-3 tools needs one. The release check of criterion 8 must still run on
  the final image, because a Recommends (section 6) or a later package can bring one.

## 9. Own network and firewall manager, and how each is switched off (6A criterion 5)

What the base brings (**[built]** B, guest `systemctl is-enabled`): `networking.service` (ifupdown) **enabled**,
`nftables.service` **disabled**, `systemd-networkd` **disabled**, `systemd-resolved` not installed (separate
package), `dnsmasq.service` not installed (only `dnsmasq-base`), `hostapd`, `wpa_supplicant`, `openvpn` **enabled**
by their packages. No NetworkManager, no netplan.

- **Default configuration.** `/etc/network/interfaces` contains only `source /etc/network/interfaces.d/*`, and
  `interfaces.d` is empty, so ifupdown brings up nothing but `lo`; the guest's NIC `enp0s2` stayed DOWN with no
  address. A real installer run writes its own `interfaces` (not given by the ISO; not determined here).
  `/etc/nftables.conf` (from `nftables 1.1.3-1`, 243 bytes) starts with `flush ruleset` and defines one empty
  `table inet filter` with input/forward/output chains and no policy.
- **`nftables.service`** (`/usr/lib/systemd/system/nftables.service`): `ExecStart=/usr/sbin/nft -f /etc/nftables.conf`,
  `ExecReload=` the same, `ExecStop=/usr/sbin/nft flush ruleset`.

Coexistence tests in the guest (probe script, exit codes as printed):

| Test | Command | Seen |
| :--- | :--- | :--- |
| Own nft table, then service start | `nft add table inet nmbtest` + input chain (rc 0); `systemctl start nftables` (rc 0); `nft list tables` | only `table inet filter`: **NMB's table deleted** (the file's `flush ruleset`) |
| Again, service restart | `nft add table inet nmbtest`; `systemctl restart nftables` (rc 0) | only `table inet filter`: **deleted** |
| Again, service stop | `nft add table inet nmbtest`; `systemctl stop nftables` (rc 0) | empty list: **everything flushed**, NMB's table included |
| Own bridge across ifupdown restart | `ip link add nmbbr0 type bridge`, up, `198.51.100.77/24` (rc 0); `systemctl restart networking` (rc 0) | `nmbbr0 ... 198.51.100.77/24` **kept** |

**What this means for criterion 5.** Both managers can be switched off cleanly, and are mostly off already:

- nftables: leave `nftables.service` disabled (the default) or `systemctl mask nftables`; any start, reload, restart
  or stop of it wipes NMB's tables, because of `flush ruleset` and `ExecStop`. NMB must either own
  `/etc/nftables.conf` or keep the unit masked.
- ifupdown: with an empty `interfaces` it touches nothing; `systemctl disable networking` or removing `ifupdown`
  (Priority important, not Essential) takes it out. It left a foreign bridge and address alone across a restart.
- systemd-networkd: present but disabled; `systemctl mask systemd-networkd` keeps it off.
- hostapd, wpa_supplicant, openvpn: NMB has to `disable`/`mask` the package units (or use only their `@` template
  units under its own control), because they are enabled on install.
- dhcpcd: only `dhcpcd-base` (binary, no unit in B: `tar -tf` found no `dhcpcd*.service`).

## 10. Installed packages

**[built] B, amd64, 171 packages** (`var/lib/dpkg/status`, Status installed). The 27 not in A are:
`dnsmasq-base dracut-install ethtool hostapd initramfs-tools initramfs-tools-bin initramfs-tools-core iw klibc-utils
libdbus-1-3 libklibc liblzo2-2 libnetfilter-conntrack3 libnfnetlink0 libnl-3-200 libnl-genl-3-200 libnl-route-3-200
libpcap0.8t64 libpcsclite1 libpkcs11-helper1t64 linux-base linux-image-6.12.107+deb13-amd64 linux-image-amd64
openvpn ppp wireguard-tools wpasupplicant`. A's 144 are the rest of the list (`comm` of the two lists: A ⊂ B).

`adduser 3.152`, `apt 3.0.3`, `apt-utils 3.0.3`, `base-files 13.8+deb13u7`, `base-passwd 3.6.7`, `bash 5.2.37-2+b10`, `bsdutils 1:2.41.5-0+deb13u1`, `coreutils 9.7-3`, `cpio 2.15+dfsg-2`, `cron 3.0pl1-197`, `cron-daemon-common 3.0pl1-197`, `dash 0.5.12-12`, `debconf 1.5.91`, `debconf-i18n 1.5.91`, `debian-archive-keyring 2025.1`, `debianutils 5.23.2`, `dhcpcd-base 1:10.1.0-11+deb13u4`, `diffutils 1:3.10-4`, `dmidecode 3.6-2`, `dnsmasq-base 2.91-1+deb13u2`, `dpkg 1.22.22`, `dracut-install 106-6`, `e2fsprogs 1.47.2-3+b12`, `ethtool 1:6.14.2-1`, `fdisk 2.41.5-0+deb13u1`, `findutils 4.10.0-3`, `gcc-14-base 14.2.0-19`, `grep 3.11-4`, `gzip 1.13-1+deb13u1`, `hostapd 2:2.10-24`, `hostname 3.25`, `ifupdown 0.8.44+deb13u1`, `init 1.69~deb13u1`, `init-system-helpers 1.69~deb13u1`, `initramfs-tools 0.148.4`, `initramfs-tools-bin 0.148.4`, `initramfs-tools-core 0.148.4`, `iproute2 6.15.0-1`, `iputils-ping 3:20240905-3`, `iw 6.9-1`, `klibc-utils 2.0.14-1`, `kmod 34.2-2`, `less 668-1`, `libacl1 2.3.2-2+b1`, `libapparmor1 4.1.0-1`, `libapt-pkg7.0 3.0.3`, `libattr1 1:2.5.2-3`, `libaudit-common 1:4.0.2-2+deb13u1`, `libaudit1 1:4.0.2-2+deb13u1`, `libblkid1 2.41.5-0+deb13u1`, `libbpf1 1:1.5.0-3`, `libbsd0 0.12.2-2`, `libbz2-1.0 1.0.8-6`, `libc-bin 2.41-12+deb13u4`, `libc6 2.41-12+deb13u4`, `libcap-ng0 0.8.5-4+b1`, `libcap2 1:2.75-10+deb13u1+b3`, `libcap2-bin 1:2.75-10+deb13u1+b3`, `libcom-err2 1.47.2-3+b12`, `libcrypt1 1:4.4.38-1`, `libdb5.3t64 5.3.28+dfsg2-9`, `libdbus-1-3 1.16.2-2`, `libdebconfclient0 0.280`, `libedit2 3.1-20250104-1`, `libelf1t64 0.192-4`, `libext2fs2t64 1.47.2-3+b12`, `libfdisk1 2.41.5-0+deb13u1`, `libgcc-s1 14.2.0-19`, `libgmp10 2:6.3.0+dfsg-3`, `libgssapi-krb5-2 1.21.3-5+deb13u1`, `libhogweed6t64 3.10.1-1`, `libidn2-0 2.3.8-2`, `libjansson4 2.14-2+b3`, `libk5crypto3 1.21.3-5+deb13u1`, `libkeyutils1 1.6.3-6`, `libklibc 2.0.14-1`, `libkmod2 34.2-2`, `libkrb5-3 1.21.3-5+deb13u1`, `libkrb5support0 1.21.3-5+deb13u1`, `liblastlog2-2 2.41.5-0+deb13u1`, `liblocale-gettext-perl 1.07-7+b1`, `liblz4-1 1.10.0-4`, `liblzma5 5.8.1-1+deb13u1`, `liblzo2-2 2.10-3+b1`, `libmd0 1.1.0-2+b1`, `libmnl0 1.0.5-3`, `libmount1 2.41.5-0+deb13u1`, `libncursesw6 6.5+20250216-2`, `libnetfilter-conntrack3 1.1.0-1`, `libnettle8t64 3.10.1-1`, `libnewt0.52 0.52.25-1`, `libnfnetlink0 1.0.2-3`, `libnftables1 1.1.3-1`, `libnftnl11 1.2.9-1`, `libnl-3-200 3.7.0-2`, `libnl-genl-3-200 3.7.0-2`, `libnl-route-3-200 3.7.0-2`, `libpam-modules 1.7.0-5`, `libpam-modules-bin 1.7.0-5`, `libpam-runtime 1.7.0-5`, `libpam0g 1.7.0-5`, `libpcap0.8t64 1.10.5-2`, `libpcre2-8-0 10.46-1~deb13u2`, `libpcsclite1 2.3.3-1`, `libpkcs11-helper1t64 1.30.0-1+b1`, `libpopt0 1.19+dfsg-2`, `libproc2-0 2:4.0.4-9`, `libreadline8t64 8.2-6`, `libseccomp2 2.6.0-2`, `libselinux1 3.8.1-1`, `libsemanage-common 3.8.1-1`, `libsemanage2 3.8.1-1`, `libsepol2 3.8.1-1`, `libslang2 2.3.3-5+b2`, `libsmartcols1 2.41.5-0+deb13u1`, `libsqlite3-0 3.46.1-7+deb13u2`, `libss2 1.47.2-3+b12`, `libssl3t64 3.5.7-1~deb13u2`, `libstdc++6 14.2.0-19`, `libsystemd-shared 257.13-1~deb13u1`, `libsystemd0 257.13-1~deb13u1`, `libtext-charwidth-perl 0.04-11+b4`, `libtext-iconv-perl 1.7-8+b4`, `libtext-wrapi18n-perl 0.06-10`, `libtinfo6 6.5+20250216-2`, `libtirpc-common 1.3.6+ds-1`, `libtirpc3t64 1.3.6+ds-1`, `libudev1 257.13-1~deb13u1`, `libunistring5 1.3-2`, `libuuid1 2.41.5-0+deb13u1`, `libxtables12 1.8.11-2`, `libxxhash0 0.8.3-2`, `libzstd1 1.5.7+dfsg-1`, `linux-base 4.12.1`, `linux-image-6.12.107+deb13-amd64 6.12.107-1`, `linux-image-amd64 6.12.107-1`, `linux-sysctl-defaults 4.12.1`, `login 1:4.16.0-2+really2.41.5-0+deb13u1`, `login.defs 1:4.17.4-2`, `logrotate 3.22.0-1`, `logsave 1.47.2-3+b12`, `mawk 1.3.4.20250131-1`, `mount 2.41.5-0+deb13u1`, `nano 8.4-1+deb13u1`, `ncurses-base 6.5+20250216-2`, `ncurses-bin 6.5+20250216-2`, `netbase 6.5`, `nftables 1.1.3-1`, `openssl-provider-legacy 3.5.7-1~deb13u2`, `openvpn 2.6.14-1+deb13u3`, `passwd 1:4.17.4-2`, `perl-base 5.40.1-6+deb13u1`, `ppp 2.5.2-1+1`, `procps 2:4.0.4-9`, `readline-common 8.2-6`, `sed 4.9-2+deb13u1`, `sensible-utils 0.0.25`, `sqv 1.3.0-3+b2`, `systemd 257.13-1~deb13u1`, `systemd-sysv 257.13-1~deb13u1`, `sysvinit-utils 3.14-4`, `tar 1.35+dfsg-3.1`, `tzdata 2026c-0+deb13u1`, `udev 257.13-1~deb13u1`, `util-linux 2.41.5-0+deb13u1`, `vim-common 2:9.1.1230-2`, `vim-tiny 2:9.1.1230-2`, `whiptail 0.52.25-1`, `wireguard-tools 1.0.20210914-3`, `wpasupplicant 2:2.10-24`, `zlib1g 1:1.3.dfsg+really1.3.1-1+b1`

**[built] C, arm64:** **174 packages**, all `install ok installed` (`var/lib/dpkg/status`). Compared with B (architecture names aligned): B's 171 plus `dosfstools 4.2-1.2` (a dependency of `raspi-firmware`), `firmware-brcm80211 20250410-2` and `raspi-firmware 1.20240424+ds-6`; the kernel is `linux-image-6.12.107+deb13-arm64`. Every other package has the same version as on amd64 except three binary rebuilds: `grep 3.11-4+b1`, `iw 6.9-1+b1`, `logrotate 3.22.0-1+b2` (`diff` of the two lists). C took 2,576 s to build under user emulation (maximum RSS 287 MB).

## 11. Size on disk

| | Bytes | Source |
| :--- | ---: | :--- |
| amd64 ISO | 792,723,456 | `stat` |
| arm64 ISO | 735,358,976 | `stat` |
| A amd64 root (`important`, no kernel) | 194,678,472 | sum of `tar -tvf` sizes; tar file 201,029,120. Of that `usr/share/locale` 39,514,506, `usr/share/doc` 10,080,801, `usr/share/man` 4,997,156 |
| B amd64 root (kernel + tools) | 360,016,018 | same method; tar 370,859,520. Of that `usr/lib/modules` 99,582,881, `/boot` 48,782,819 (initrd 36,356,871 + vmlinuz 12,142,528 + config, System.map), locale 39,834,748, doc 13,128,264, man 5,318,782 |
| B in the guest | 671,732 KiB used of 1,992,552 KiB on `/` (ext4, 36 %) | guest `df -k` |
| C arm64 root (kernel + tools + Pi firmware) | 599,503,592 (tar 611,438,080). Of that `usr/lib/modules` 84,970,456; `/boot` 172,033,118, of which `/boot/firmware` 97,105,598 (Pi firmware and the copies of kernel and initrd, section 2); `usr/lib/raspi-firmware` 22,281,230; `usr/lib/firmware` 18,741,563 (79 files) | same method |

- apt lists and the `.deb` cache are cleaned by mmdebstrap at the end (the in-hook `du` showed 110 MB in
  `var/cache/apt` and 57 MB in `var/lib/apt/lists` for A; the tar has none). A running system fills both again
  after `apt update`.
- Docs, man pages and locales (about 58 MB in B) and the kernel modules NMB does not need are the obvious cuts for
  an NMB image; not measured here.

## 12. QEMU boot test (build B, amd64)

**How it was run.** The ext4 disk was made inside the mmdebstrap namespace (`cp -ax` of the root, then
`mke2fs -q -t ext4 -L nmbroot -d <copy> disk.ext4 2G`, rc 0), with an `/etc/fstab` line for `/dev/vda`, host name
`nmb-i06` and a probe unit (`nmb-probe.service`, `Type=simple`, after `multi-user.target`) that waits for
`systemctl is-system-running --wait`, prints the figures to the serial console and powers off. The probe is a test
addition, not part of the base, and is not in the inventories of sections 6–10 (they were taken before it was
added). Each boot ran on the unchanged disk (`snapshot=on`) under `flock -w … ~/nmb-work/qemu.lock` for the whole
run:

```
qemu-system-x86_64 -machine q35 -accel tcg,tb-size=128 -cpu max -smp 1 -m 512 -display none -monitor none
  -serial file:serial.log -no-reboot
  -kernel vmlinuz-6.12.107+deb13-amd64 -initrd initrd.img-6.12.107+deb13-amd64
  -append "root=/dev/vda rw console=ttyS0,115200n8 systemd.show_status=1 systemd.mask=serial-getty@ttyS0.service"
  -drive file=disk-B-nmb-amd64.ext4,format=raw,if=virtio,snapshot=on
  -netdev user,id=n0,net=192.0.2.0/24 -device virtio-net-pci,netdev=n0,mac=02:00:00:00:06:01
```

- QEMU 8.2.2 on the PC. **TCG, no KVM:** `/dev/kvm` is `root:kvm 0660` with an ACL for the owner and group only,
  and the user is not in group `kvm` (`test -r /dev/kvm -a -w /dev/kvm` failed). Direct kernel boot, so neither
  firmware nor GRUB was exercised.

| Run | Settings | Result |
| :--- | :--- | :--- |
| 1 | as above, without `systemd.mask=serial-getty@ttyS0.service` | Booted and powered off cleanly (QEMU rc 0, flock rc 0, 07:28:55–07:33:42 UTC). The probe printed only its first line: `serial-getty@ttyS0` took over the console and hung up the probe's output. Data not used |
| 2 | as above | Booted, the probe completed, `poweroff` clean: QEMU rc 0, flock rc 0, 07:34:14–07:39:09 UTC (4 min 55 s). Used below |

Figures from run 2 (TCG, one vCPU, 512 MB; **not** a figure for real
hardware):

- **Boot time** (guest `systemd-analyze`): `Startup finished in 41.495s (kernel) + 1min 29.682s (userspace) =
  2min 11.177s`; `graphical.target reached after 1min 29.046s in userspace`. Slowest units (`blame`):
  `dev-vda.device` 68.5 s, `systemd-udev-trigger` 23.7 s, `networking` 22.5 s, `systemd-journald` 20.8 s. Kernel log:
  `Run /init as init process` at 19.79 s.
- **Memory** (guest `free -k` after boot): total 473,216 kB, used 61,412 kB, free 372,244 kB, available 411,804 kB.
  Kernel line: `Memory: 427396K/523752K available`.
- **Storage after boot:** `/dev/vda` 1,992,552 kB, 671,732 kB used, 1,199,580 kB free; `/tmp` tmpfs 236,608 kB;
  `/run` 47,324 kB.
- **Listening ports** (guest `ss -tulpnH`): **none**.
- **Running services** (guest `systemctl list-units --type=service --state=running`): `cron`, `getty@tty1`–`tty6`,
  `systemd-journald`, `systemd-udevd`, and the probe. `systemctl --failed`: none.
- **Network at boot:** `lo` up, `enp0s2` DOWN without address (ifupdown has no stanza for it).
- **Kernel features checked live:** section 4 (`modprobe` loop all 0, `mac80211_hwsim` made 2 radios, Landlock in
  the active LSM list, cgroup2, watchdog device) and section 9 (nftables and ifupdown tests).
- **Not tested in the guest:** `dnsmasq`, `hostapd`, `wpa_supplicant`, `pppd` runs, WireGuard and OpenVPN tunnels;
  only their presence and version.

## 13. Fields not determined

- **Installed layout and network configuration written by the real installer:** not determined. The ISOs were not
  installed into a guest; the installed system was judged from the mmdebstrap builds (work plan 5.1 allows either).
- **Netinst on a Raspberry Pi 4** (UEFI firmware needed or not, which kernel and DTB path the installer sets up):
  not determined; no Pi was used, and the ISO cannot be booted as a Pi under QEMU.
- **arm64 boot figures:** not determined; build C was taken apart, not booted (time; TCG for arm64 is slow).
- **Boot time and memory on real hardware or KVM:** not determined; only TCG, no access to `/dev/kvm`.
- **Effect of Recommends on Python:** not resolved transitively; only the direct Recommends are listed (section 6).
- **Licences of the firmware packages:** not read (only their component, `non-free-firmware`).
- **Ookla and RTR binaries for glibc (criterion 6):** not part of the image; not checked in this package (left to
  the comparison I-09).
- **cgroup namespaces:** no config option found by name; not tested.

## 14. Tools missing on the PC

- No sudo, so no loop mounts and no `/dev/kvm` access. `isoinfo` and `genext2fs` are missing (not needed: `7z`,
  `fdisk`, `mke2fs -d`).
- mmdebstrap's unshare mode runs as a sub-uid that cannot enter the home folder (`/home/<user>` is `drwxr-x---`). The
  keyring, the hook scripts and the hook output were therefore staged in a fresh `mktemp -d /tmp/nmb-I-06.XXXXXX`
  folder (mode 755, output folder 1777); everything else stayed in `~/nmb-work/I-06/`.
- The PC's `mmdebstrap` is 1.4.3 (Ubuntu 24.04); trixie ships 1.5.7.
- Present and used: `7z`, `fdisk`, `file`, `dpkg-deb`, `tar`, `gpg`, `gpgv`, `curl`, `mmdebstrap` 1.4.3,
  `qemu-aarch64-static` (binfmt `F` flag), `mke2fs` 1.47.0, `qemu-system-x86_64` 8.2.2 and `flock`.

## 15. Debian as the NMB base: what the native build adds and what it can drop

The owner's Raspberry Pi 4 runs the Python version on Debian 13 today. The Python version's `setup.sh` (read in
`~/src/NetworkMeasureBox`, lines 136–206) installs with `apt-get install --no-install-recommends`: `python3 iproute2
nftables iptables ethtool pciutils usbutils polkitd dnsmasq hostapd iw wireless-regdb ppp pppoe udhcpc curl
ca-certificates procps libcap2-bin wireguard-tools openvpn easy-rsa python3-qrcode qrencode traceroute
iputils-tracepath mtr-tiny iperf3 sqlite3 gnupg dirmngr gcc make libc-dev libssl-dev libcurl4-openssl-dev
pkg-config`, plus `firmware-realtek`, `miniupnpd-nftables`, OONI and Ookla from outside Debian; and it ships 25
`nmb-*` systemd units (`ls systemd/`).

On top of a minimal Debian (A, 144 packages, no Python) the native build needs:

- the kernel package (`linux-image-amd64`, or `linux-image-arm64` + `raspi-firmware` on the Pi 4) and an initramfs
  tool (B pulled `initramfs-tools`);
- the criterion-3 tools not in A: `dnsmasq-base`, `hostapd`, `wpasupplicant`, `iw` (+ `wireless-regdb`), `ppp`,
  `wireguard-tools`, `openvpn`; `nftables` and `iproute2` are already in A;
- the firmware packages for supported adapters (`firmware-realtek`, `firmware-brcm80211` for the Pi 4 Wi-Fi, ...),
  because a read-only image cannot fetch them later (roadmap 6A);
- its own systemd units for `nmbd` (fewer than 25), and masks for `nftables.service`, `systemd-networkd`, and the
  package units of `hostapd`, `wpa_supplicant`, `openvpn` (section 9).

It can drop, compared with the Python install: `python3`, `python3-qrcode`; the compilers and headers (`gcc make
libc-dev libssl-dev libcurl4-openssl-dev pkg-config`); `polkitd` (only if no D-Bus caller needs it); `iptables`
(the base kernel has nftables and `NFT_COMPAT`); `udhcpc` (`dhcpcd-base` is in A, or nmbd's own client); and, from
the base itself, `ifupdown` (if nmbd owns all interfaces), `cron` and the apt timers (updates come as images),
docs, man pages and locales. Each of these is a proposal for the image package, not tested here.

## 16. 6A criteria for this candidate

| Criterion | Finding | Where |
| :--- | :--- | :--- |
| 1. Stock kernel, Pi 4 and x86-64, network features | Met by the configs: one Debian kernel 6.12.107 for both, with bridge, bonding, VLAN, nftables, WireGuard, cfg80211/mac80211/hwsim, every tethering driver of 4.27, and the Pi 4 platform (`ARCH_BCM2835`, `PCIE_BRCMSTB`, `BCMGENET`, DTBs). Loaded live on amd64 only; the Pi 4 boot files are set up by `raspi-firmware` (read, not booted) | 4 |
| 2. Security updates for years, free | Full support to 2028-08-09, LTS to 2030-06-30 (debian.org, wiki LTS) | 3 |
| 3. Fits nmbd | systemd 257, nftables 1.1.3, dnsmasq 2.91, hostapd/wpa_supplicant 2.10, iw 6.9, pppd 2.5.2, wireguard-tools 1.0.20210914, OpenVPN 2.6.14, full iproute2 6.15.0; all from `main`, none needs Python | 7 |
| 4. Read-only A/B image, reproducible | Not in the base (apt in place). `mmdebstrap`, `rauc`, `swupdate`, `systemd-repart`, `debos` are packaged in trixie; untested | 3 |
| 5. Does not fight NMB | ifupdown idle with the default config and left NMB's bridge alone; nftables.service is off by default but wipes all tables when started, restarted or stopped, so it must stay masked; hostapd/wpa_supplicant/openvpn units come enabled | 9 |
| 6. Size, memory, glibc | amd64 B root 360 MB (A, without kernel and tools, 194.7 MB); arm64 C root 599.5 MB, of which 97 MB is `/boot/firmware` with second copies of kernel and initrd; 61 MB used after boot in a 512 MB amd64 TCG guest; glibc 2.41. Ookla/RTR not checked here | 11, 12 |
| 7. Licence | All of A/B from `main`; Pi boot files and Wi-Fi firmware from `non-free-firmware` | 6 |
| 8. No Python in the base | Met for the `important` base and the tool set; Python arrives only with the "standard" task (`reportbug`) or Recommends, so the image needs the release check | 8 |
