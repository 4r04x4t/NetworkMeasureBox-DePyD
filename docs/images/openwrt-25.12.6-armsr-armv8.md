# Image record: OpenWrt 25.12.6 armsr/armv8 (64-bit ARM, squashfs and ext4 combined-efi, with SDK, ImageBuilder, toolchain)

Package I-02 (work plan section 5.4). The method is work plan 5.1, and every field of 5.2 is below, in the order
of the I-01 record (`openwrt-25.12.6-x86-64.md`) so that I-09 can compare them; section 15 adds the build trees
(5.3 point 4). Each value names the command that produced it. A value that could not be read says "not determined"
and why. Nothing here comes from memory. Only lists, versions and checksums are recorded. No file taken out of the
image is in this repository.

- Date of the analysis: 2026-11-11.
- Where it ran: dev PC, Windows with WSL2 Ubuntu 24.04, about 1.8 GB RAM and 4 threads, no root. The QEMU guest ran
  on the same PC, under TCG (an ARM guest on an x86 host; KVM does not apply).
- Work folder: `~/nmb-work/I-02/` (outside every repository; one fresh sub-folder per unpack).
- Material, read only: `~/src/NMB-depythonization-firmwares/openwrt-target-trees/openwrt-25.12.6-armsr-armv8/`.

## 1. Files and checksums

| File | Size (bytes) | SHA-256 (publisher list) | Check |
| :--- | ---: | :--- | :--- |
| `openwrt-25.12.6-armsr-armv8-generic-squashfs-combined-efi.img.gz` | 13,425,520 | `da639ebe2205532dce3ffacd2e91657c15d0ac77234c26bda1f63efbbd8efba1` | OK |
| `openwrt-25.12.6-armsr-armv8-generic-ext4-combined-efi.img.gz` | 14,669,158 | `77a3802b583a6964ccb868c80f83a6ea0eb5847d3835350b5f0bcbf6f141d1d3` | OK |
| `openwrt-25.12.6-armsr-armv8-generic-kernel.bin` | 20,222,464 | `2b9ce48df5536e993ee691339057828488e4a31be9c2bb5d3f6120f1135fcbb0` | OK |
| `openwrt-25.12.6-armsr-armv8-generic-initramfs-kernel.bin` | 35,754,496 | `2b1f0796ec31201b609643beca48de20d75f2879104a20adba327cbaf64b5f2b` | OK |
| `openwrt-25.12.6-armsr-armv8-generic-squashfs-rootfs.img.gz` | 4,723,333 | `5ebf47db2d19410571b6f760ea7717891a681d3fda836978e9750e7a575ed709` | OK |
| `openwrt-25.12.6-armsr-armv8-generic-ext4-rootfs.img.gz` | 5,967,598 | `3f25a4ee0537f0a4336833047d7933e884c29450eefb7ac758155cbd0f5278d3` | OK |
| `openwrt-25.12.6-armsr-armv8-generic-targz-rootfs.tar.gz` | 5,851,547 | `b781548d30886845176fee8256e115188029de0c42af252ffc6ea94a62794381` | OK |
| `openwrt-25.12.6-armsr-armv8-rootfs.tar.gz` | 4,504,215 | `8fcb00eb886d9add37e111a4ec43d47d57aa74742281c6a8f348a7a7ec37edb7` | OK |
| `openwrt-25.12.6-armsr-armv8-rootfs.cpio.gz` | 4,501,948 | `3b2650b0ec84980258a28e45c869d7915d7dc33c2b66d850a01f2b1edc761379` | OK |
| `openwrt-sdk-25.12.6-armsr-armv8_gcc-14.3.0_musl.Linux-x86_64.tar.zst` | 259,484,411 | `cbf3343bc3d6dc6180fc84077b966c13596a79904adaadbfe68409d563b0b596` | OK |
| `openwrt-imagebuilder-25.12.6-armsr-armv8.Linux-x86_64.tar.zst` | 37,060,608 | `3d0128148395314c510fe7adfc5996cecec51f3011ea2acc560a9c21772301cf` | OK |
| `openwrt-toolchain-25.12.6-armsr-armv8_gcc-14.3.0_musl.Linux-x86_64.tar.zst` | 70,140,391 | `3daaf41e6d7416183dfaa7ddb35c5c6743c5101200046e3fc26af8a89dab1148` | OK |
| `kernel-debug.tar.zst` | 124,945,890 | `a002937a21fdcdc577ab77a56f58eb505474c871115e80e32ddec69ec4ec5403` | OK |
| `llvm-bpf-21.1.6.Linux-x86_64.tar.zst` | 51,187,706 | `3f07164c83d2965bc8685bbb1f1efda4539e9c86e42779f52808655b039df574` | OK |
| `u-boot-qemu_armv8/u-boot.bin` | 1,148,512 | `174fe659a2cf28a45304a88555b2f026750f3561bb1e88cba5123924d462b57c` | OK |
| `profiles.json`, `*.manifest`, `*.bom.cdx.json`, `*.buildinfo` | small | in the list | OK |
| `kmods/6.12.112-1-67e205f26d0458c4d12f1fd315216398/` (1,145 `.apk`, `index.json`, `packages.adb`) | — | in the list | OK |
| `packages/` (88 `.apk`, `index.json`, `packages.adb`) | — | in the list | OK |

- **How it was checked.** The tree carries the publisher's `sha256sums` (1,258 lines, SHA-256
  `eef9bcd2a1c517436a3bf9a1cee060d22e959e29a3e725ff108c92ee60ca5ec8`). Every file of the tree except the list itself
  and the four index signature files (`packages.adb.asc`/`.sig` in `kmods/…` and `packages/`) is named in it, and every
  listed file is present (`find` + `comm`). `sha256sum -c --quiet sha256sums` gave exit 0; 1,258 of 1,258 lines `OK`.
- **Signature of the list (gpg).** Checked offline in an isolated `GNUPGHOME`, `gpg --verify sha256sums.asc sha256sums`:
  "Good signature from OpenWrt Build System (Nitrokey3)", exit 0, signature made 2026-10-18.
  - Signing subkey `92C5 61DE 55AE 6552 F3C7 36B8 2B01 5109 0606 D1D9`, primary key
    `8A8B C12F 46B8 36C0 F9CD B36F 1D53 D187 7742 E911`.
  - No key was fetched for I-02: the key file that I-01 fetched and compared with OpenWrt's keyring repository
    (`gpg/0x1D53D1877742E911.asc`) was reused from `~/nmb-work/I-01/pub/`. `gpg --show-keys` gives the same
    fingerprints. gpg says "not certified with a trusted signature", because no web of trust was set up.
  - The two index signatures were checked the same way: `gpg --verify packages.adb.asc packages.adb` for `packages/`
    and for `kmods/…/`, both "Good signature", exit 0.
- **Signature of the list (usign): not checked.** `sha256sums.sig` is a usign signature ("verify with usign.pub").
  `usign`/`signify` is not installed on the PC, and the image carries no usign public key (`/etc/opkg/keys/` is empty;
  `/etc/apk/keys/` holds only the apk key `openwrt-25.12.pem`), so it cannot be checked without fetching a key.
- **Release identity.** `version.buildinfo` = `r33311-981c978eff`; the image's `/etc/openwrt_release` gives the same,
  with `DISTRIB_TARGET='armsr/armv8'` and `DISTRIB_ARCH='aarch64_generic'`. `profiles.json`: kernel `6.12.112`,
  release 1, vermagic `67e205f26d0458c4d12f1fd315216398`, one profile `generic` ("Generic EFI Boot").
- **Decompressed images.** `gzip -dc` gave exit 0 for both images and printed no warning. Each decompressed image is
  243,564,032 bytes (475,711 sectors).
  - squashfs: `7d1e1d58c2a9c9177e51029bddc39e47fd6037bdc51725dfcc2fb29feb4eb09c`
  - ext4: `d0304cc325bf33e46ed27dc7cfecd15b3c6bedd97d266cc2a3e0e8974618ce33`

## 2. Partition and file system layout, boot method

The layout was read with `sgdisk -p disk.img` (exit 0), `sgdisk -i N` and `fdisk -l`. It is the same in both images.

| No. | Sectors (512 B) | Size | GPT type | Contents (`file`) |
| :--- | :--- | :--- | :--- | :--- |
| 128 | 34–511 | 239 KiB | BIOS boot (`EF02`) | all zero bytes (`tr -d '\0' \| wc -c` = 0): unused on ARM |
| 1 | 512–262655 | 128.0 MiB | **EFI System** (`C12A7328-…`, name "EFI System Partition") | FAT16, label `kernel` |
| 2 | 262656–475647 | 104.0 MiB | Linux filesystem (`0FC63DAF-…`) | squashfs image: squashfs 4.0, xz, 4,619,820 bytes, 1,308 inodes. ext4 image: ext4, label `rootfs`, 26,624 blocks of 4 KiB, 20,553 free (`debugfs -R stats`) |

- **Partition table.** GPT behind a protective MBR. As on x86-64, the backup GPT header is missing: sgdisk reports
  "invalid backup GPT header", and the guest kernel logs `GPT:Alternate GPT header not at the end of the disk`
  (`475680 != 475710`) and continues.
- **Differences from x86-64 (I-01).** Partition 1 is a real ESP here (type EF00, 128 MiB instead of 16 MiB), and
  partition 128 holds no boot code.
- **Boot partition.** Read with `mdir -/ -i p1.img ::` and `mcopy -s` (both exit 0). It holds
  `/efi/boot/bootaa64.efi` and `/efi/openwrt/bootaa64.efi` (the same file, 729,088 bytes, GRUB 2.12 by `strings`),
  `/efi/openwrt/grub.cfg` (525 bytes) and `/efi/openwrt/vmlinux` (20,222,464 bytes). 112,246,784 bytes are free.
  - Every boot file is identical in the two images (`sha256sum`).
  - `vmlinux` is an uncompressed `Linux kernel ARM64 boot executable Image, little-endian, 4K pages` (`file`) with the
    EFI stub, and is byte-identical to the publisher's `openwrt-25.12.6-armsr-armv8-generic-kernel.bin`.
- **Boot method: UEFI only.** The firmware loads `\efi\boot\bootaa64.efi` (console: `BdsDxe: starting Boot0001 "UEFI
  Misc Device"`), GRUB loads `/efi/openwrt/vmlinux` after the 5 s timeout, and the kernel's EFI stub boots
  ("EFI stub: Booting Linux Kernel…", "Generating empty DTB"): on QEMU `virt` with edk2 the kernel ran on **ACPI**
  (`psci: probing for conduit method from ACPI`, console from the `SPCR` table: pl011 at `0x9000000`).
  - GRUB prints `error: serial port 'com0' isn't found` and `terminal 'serial' isn't found` (the `serial --unit=0`
    line of grub.cfg is x86-style) and falls back to the EFI console, which reaches the same serial port.
  - Kernel command line (guest `cat /proc/cmdline`): `BOOT_IMAGE=/efi/openwrt/vmlinux root=PARTUUID=<disk GUID>-02
    rootwait earlycon noinitrd`. There is no `console=`: the console is `ttyAMA0` from SPCR. No initramfs.
  - A second GRUB entry boots with `failsafe=true`.
  - The tree also carries `u-boot-qemu_armv8/u-boot.bin` (a U-Boot for QEMU `virt`); it was not used (section 12).
- **Root file system at run time (squashfs image).** From the guest's `mount` and `df -k`:
  - The squashfs is mounted read-only on `/rom`.
  - The rest of partition 2 is a loop device (`loop0`, 203,904 sectors), formatted ext4 on the first boot
    (`mount_root: overlay filesystem in /dev/loop0 has not been formatted yet`, then `switching to ext4 overlay`).
  - The overlay is mounted on `/overlay`, and `overlayfs:/overlay` is `/`. `/tmp` is tmpfs.
  - **The ESP is not mounted at run time** (no `/boot`, unlike x86-64); `efivarfs` is mounted on
    `/sys/firmware/efi/efivars`.
- **Root file system (ext4 image).** Taken apart, not booted. Unpacked with `debugfs -R "rdump / out" p2.img` (exit 0;
  warnings about changing ownership, because there is no root). Its file list equals the squashfs root's
  (`find -type f | sort`, `diff` exit 0: 957 files), and so does its package list (`diff` exit 0).
- **Unpacking the squashfs.** `unsquashfs -no-xattrs -d … p2.img` gave exit 2; the one failure was
  `could not create character device …/dev/console, because you're not superuser`, expected without root. It
  created 957 files, 118 directories and 232 symlinks.
- **Other root forms in the tree.** `…-generic-targz-rootfs.tar.gz` (1,307 entries, 205 packages: the same root as the
  disk images) and the profile-less `…-rootfs.tar.gz` (1,075 entries, 131 packages, the list in
  `openwrt-25.12.6-armsr-armv8.manifest`: the default set without the device packages of `profiles.json`); counts from
  `tar -tzf … | wc -l` and `grep -c '^P:'` on `lib/apk/db/installed`. `rootfs.cpio.gz` is a `070701` (newc) cpio.

## 3. How an update is written

- **sysupgrade.** Present (`/sbin/sysupgrade`, guest `command -v`). The armsr rule is `/lib/upgrade/platform.sh`,
  read in the unpacked root:
  - `platform_check_image` compares the image's partition table (first 63 sectors) with the disk; a changed layout
    means "Full image will be written".
  - `platform_do_upgrade`: with `-p` (keep partitions) unset and a changed layout, it `dd`s the whole image and reruns
    `partx`. Otherwise it writes partition by partition: for the ESP it **only replaces `efi/openwrt` and `efi/boot`**
    and keeps other firmware files on the ESP ("ubootefi.var or firmware/X/ aka EBBR"); other partitions are `dd`;
    then it rewrites the `PARTUUID=` in `/efi/openwrt/grub.cfg`.
  - `platform_copy_config` copies the saved settings onto partition 1.
- **No A/B scheme.** One root partition, overwritten in place; settings travel in the backup file.
- **Other update helpers.** `owut` and `attendedsysupgrade-common` are installed (package list).
- **Packages at run time.** `apk` installs into the overlay (section 9a: 20 packages added in the guest).

## 4. Kernel and the configuration NMB needs

- **Kernel version: 6.12.112.** Sources:
  - `strings` of `vmlinux`: `Linux version 6.12.112 (builder@buildhost) (aarch64-openwrt-linux-musl-gcc (OpenWrt GCC
    14.3.0 r33311-981c978eff) 14.3.0, GNU ld (GNU Binutils) 2.44) #0 SMP Sat Oct 17 21:45:38 2026`
  - guest `uname -a` (`aarch64`) and `ubus call system board` (`"kernel": "6.12.112"`, `"rootfs_type": "squashfs"`)
  - installed package `kernel 6.12.112~67e205f26d0458c4d12f1fd315216398-r1`
- **Where the config came from.** Not in the image: `vmlinux` contains no `IKCFG_ST` marker (`grep -c -a` = 0), the
  guest has no `/proc/config.gz`, and the root carries no `.config`. Two sources in the tree instead, both local:
  - **Primary: `kmod-ikconfig-6.12.112-r1.apk`** from `kmods/6.12.112-1-67e205f26d0458c4d12f1fd315216398/`
    (`CONFIG_IKCONFIG=m`). SHA-256 `8edf79fe5549af98690f1edf3d1d1bbb45049c710f7ad1fdb1b5a5c2b09eed7b`, in the signed
    list. The kmods folder hash equals the `kernel` package's version hash. Inflated with the same 7-line Python
    script as I-01 (`python3 -I`, script outside the data folder; ADB `ADBd` raw deflate, then the gzip blob between
    `IKCFG_ST` and `IKCFG_ED`), exit 0. Result: `# Linux/arm64 6.12.112 Kernel Configuration`,
    `CONFIG_CC_VERSION_TEXT="aarch64-openwrt-linux-musl-gcc (OpenWrt GCC 14.3.0 r33311-981c978eff) 14.3.0"`;
    SHA-256 of the extracted config `63f4c2ac917082dc45eaf1f9ad862c94fb9c6c368c086aa8c192b866ed49fda8`.
  - **Cross-check: the SDK's and ImageBuilder's
    `build_dir/target-aarch64_generic_musl/linux-armsr_armv8/linux-6.12.112/.config`** (9,492 lines, SHA-256
    `d90ce6993c9f36f9104e1c1bef0720cb2053d59dbd33b373a5df1f67cc361afe`; the two are identical, `diff -q`). `diff`
    against the ikconfig shows only the initramfs lines (`CONFIG_INITRAMFS_SOURCE`, `_ROOT_UID/GID`,
    `INITRAMFS_COMPRESSION_NONE`): the trees keep the config of the last (initramfs) kernel build. Every option
    below is the same in both.
  - "In image" says whether the module is in `lib/modules/6.12.112/`; "Package" whether a kmod is in the tree's
    `kmods/` folder (1,145 packages, `index.json`).

| Feature (5.2) | Option(s) | Value | In image | Package |
| :--- | :--- | :--- | :--- | :--- |
| Bridge, VLAN filtering | `BRIDGE`, `BRIDGE_VLAN_FILTERING` | y, y | built in | — |
| Bonding | `BONDING` | m | no | `kmod-bonding 6.12.112-r1` |
| Team | `NET_TEAM` | m | no | `kmod-team` and five mode packages |
| VLAN 802.1Q | `VLAN_8021Q` | y | built in | — |
| nftables | `NF_TABLES` m, `NF_TABLES_INET` y, `NF_TABLES_NETDEV` y, `NF_TABLES_BRIDGE` m, `NFT_NAT` m, `NFT_MASQ` m, `NFT_FIB_INET` m, `NF_FLOW_TABLE` m | as listed | yes (`nf_tables.ko`, 29 `nft_*`/`nf_*` modules) | `kmod-nft-*`, `kmod-nft-bridge` |
| WireGuard | `WIREGUARD` | m | no | `kmod-wireguard 6.12.112-r1`; **installed and used in the guest** (section 9a) |
| PPP, PPPoE | `PPP` m, `PPPOE` m, `PPP_ASYNC` m | m | yes (`ppp_generic`, `pppoe`, `pppox`, `ppp_async`, `slhc`) | `kmod-ppp`, `kmod-pppoe` |
| cfg80211, mac80211 | `CFG80211` | **not set in the kernel** | no | `kmod-cfg80211`, `kmod-mac80211` from backports 6.18.55 (`6.12.112.6.18.55-r2`) |
| mac80211_hwsim | — | from backports | no | `kmod-mac80211-hwsim 6.12.112.6.18.55-r2`; **installed in the guest, two radios came up** (section 9a) |
| Wi-Fi drivers (examples) | — | backports / mt76 | no | `kmod-ath9k(-htc)`, `kmod-ath10k(-ct)`, `kmod-iwlwifi`, `kmod-mt76*`, `kmod-mt7601u`, `kmod-mt7921u`, `kmod-rtw88-*` |
| USB host | `USB` y, `USB_XHCI_HCD` y, `USB_XHCI_PLATFORM` y, `USB_EHCI_HCD` y, `USB_DWC3` y | built in | — | `kmod-usb-core`, `kmod-usb2`, `kmod-usb3` also published |
| USB network | `USB_USBNET` m, `USB_RTL8152` m, `USB_NET_AX88179_178A` m, `USB_NET_SMSC95XX` m, `USB_NET_AQC111` m | m | no | `kmod-usb-net-rtl8152`, `-asix-ax88179`, `-smsc95xx`, `-aqc111` |
| USB tethering (4.27) | `USB_NET_RNDIS_HOST`, `USB_NET_CDCETHER`, `USB_NET_CDC_NCM`, `USB_IPHETH`, `USB_NET_QMI_WWAN`, `USB_NET_CDC_MBIM`, `USB_NET_HUAWEI_CDC_NCM` | all m | no | `kmod-usb-net-rndis`, `-cdc-ether`, `-cdc-ncm`, `-ipheth`, `-qmi-wwan`, `-cdc-mbim`, `-huawei-cdc-ncm` |
| TUN, veth, macvlan, dummy, IFB | `TUN`, `VETH`, `MACVLAN`, `DUMMY`, `IFB` | all m | no | `kmod-tun`, `kmod-veth`, `kmod-macvlan`, `kmod-dummy`, `kmod-ifb` |
| Policy routing | `IP_ADVANCED_ROUTER`, `IP_MULTIPLE_TABLES`, `IPV6_MULTIPLE_TABLES` | y | built in | — |
| Traffic control | `NET_SCHED` y, `NET_SCH_FQ_CODEL` y, `NET_SCH_CAKE` m, `NET_SCH_HTB` m, `NET_SCH_NETEM` m, `NET_SCH_TBF` m | as listed | fq_codel built in | `kmod-sched-cake`, `kmod-netem` |
| Namespaces | `NAMESPACES`, `NET_NS`, `PID_NS`, `USER_NS`, `UTS_NS`, `IPC_NS` | all y | built in | — |
| seccomp | `SECCOMP`, `SECCOMP_FILTER` | y, y | built in | (`procd-seccomp` installed) |
| **Landlock** | `SECURITY` | **not set**, so `SECURITY_LANDLOCK` is absent | — | not available without an own kernel |
| cgroups | `CGROUPS` y, `MEMCG` y (`MEMCG_V1` not set), `CGROUP_PIDS` y, `CPUSETS` y, `CGROUP_BPF` y, `BLK_CGROUP` y; `CGROUP_FREEZER`, `CGROUP_DEVICE` not set | as listed | cgroup2 on `/sys/fs/cgroup` (guest `mount`, `nsdelegate`) | — |
| Watchdog | `WATCHDOG` y; `ARM_SP805_WATCHDOG` m, `SOFT_WATCHDOG` m, `I6300ESB_WDT` m | as listed | `sp805_wdt.ko` in the image (`kmod-wdt-sp805`), but no `/dev/watchdog*` in the guest (QEMU `virt` has no SP805); `ubus call system watchdog`: `"status": "offline", "timeout": 30` | `kmod-softdog`, `kmod-wdt-sp805` |
| Module signing | `MODULE_SIG` | not set | — | — |
| Guest/test helpers | `VIRTIO_NET`, `VIRTIO_BLK`, `VIRTIO_PCI`, `VIRTIO_MMIO`, `NET_9P`, `NET_9P_VIRTIO`, `9P_FS` y; `VIRTIO_FS` not set; `FUSE_FS` m | as listed | built in | `kmod-fs-9p`, `kmod-9pvirtio` also published |
| Other | `BPF_SYSCALL` y, `XDP_SOCKETS` not set, `TCP_CONG_BBR` m, `IPV6` y, `PREEMPT_NONE` y, `HZ=100`, `NR_CPUS=256`, `ARM64_VA_BITS=48`, 4K pages, `EFI_STUB` y, `KVM` y | as listed | — | `kmod-tcp-bbr` |

## 5. Modules and firmware files carried

- **Modules.** There are 111 `.ko` files in `lib/modules/6.12.112/` (`ls *.ko | wc -l`), 4,881,889 bytes, plus
  `modules.builtin` (460 lines). Guest `lsmod | wc -l` = 110 lines.
  - Groups: SoC and server NICs for the ARM boards of the `generic` profile (`fsl-dpaa2-eth`, `fsl-enetc*`, `fec`,
    `mvneta`, `mvpp2`, `stmmac` with `dwmac-imx`/`-rk`/`-sun8i`, `genet`, `ravb`, `thunder_*`/`nic*`,
    `rvu_*` OcteonTX2, `atlantic`, `ena`, `e1000e`, `vmxnet3`), PHY and MDIO drivers (`aquantia`, `broadcom`,
    `marvell`, `realtek`, `smsc`, `sfp`, `phylink`, …), nftables and conntrack, PPP/PPPoE, `macsec`, crypto helpers,
    I2C mux/GPIO expander, `rtc-rx8025`, `sp805_wdt`, and `vfat`/NLS.
  - virtio-net, virtio-blk and 9p are built in (section 4).
- **Firmware.** `lib/firmware` exists and is **empty** (`find lib/firmware -type f | wc -l` = 0). No NIC or Wi-Fi
  firmware is in the image.

## 6. Init system, service manager, C library, package manager

- **Init:** procd (`/sbin/procd` is PID 1 in the guest's `ps w`) with ubus, `ujail` (`/sbin/ujail`, package
  `procd-ujail`) and seccomp (`procd-seccomp`). Init scripts in `/etc/init.d`, linked in `/etc/rc.d`.
- **Enabled services** (guest loop `$s enabled`): boot, cron, dnsmasq, done, dropbear, firewall, gpio_switch, led, log,
  network, odhcpd, packet_steering, rpcd, sysctl, sysfixtime, sysntpd, system, ucitrack, uhttpd, umount,
  urandom_seed, urngd (22 of 22, the same set as x86-64).
- **Services known to procd** (guest `ubus call service list`): cron, dnsmasq, dropbear, firewall, gpio_switch, log,
  network, odhcpd, packet_steering, rpcd, sysntpd, system, ubus, ucitrack, uhttpd, urandom_seed, urngd.
- **C library: musl 1.2.5.** The guest's `/lib/libc.so` prints `musl libc (aarch64)` / `Version 1.2.5`; package
  `libc 1.2.5-r5`; `/lib/ld-musl-aarch64.so.1` is a link to `libc.so`.
- **Package manager: apk-tools 3.0.5** (guest `apk --version`: `apk-tools 3.0.5, compiled for aarch64.`). opkg is
  absent (guest `command -v opkg`).
  - Feeds in `/etc/apk/repositories.d/distfeeds.list`: target `armsr/armv8/packages`, kmods
    `6.12.112-1-67e205f26d0458c4d12f1fd315216398`, and `aarch64_generic` base, luci, packages, routing, telephony,
    video for 25.12.6. Trust key: `/etc/apk/keys/openwrt-25.12.pem`.
- **Installed packages: 205.** Same count and list in the squashfs and ext4 images and the `targz` root. Sources:
  `awk '/^P:/…' lib/apk/db/installed`; guest `apk info | wc -l` = 205 (its cache warnings appear before
  `apk update`). The list is in section 10.
  - Against x86-64 (182; `comm` of the two name lists): 144 in both, 61 only here (ARM NIC/PHY/MDIO drivers,
    `kmod-macsec`, crypto helpers, `kmod-wdt-sp805`, `kmod-rtc-rx8025`, `grub2-efi-arm`, `blkid`), 38 only on x86-64
    (i915/DRM/fb stack, x86 NICs, `grub2`, `grub2-bios-setup`, `grub2-efi`, three firmware packages).
- **Licences of the installed packages** (`L:` fields, `sort | uniq -c`): GPL-2.0-only 104, ISC 29, GPL-2.0 26,
  Apache-2.0 15 (LuCI), LGPL-2.1 7, MIT 4, GPL-3.0-with-GCC-exception 4 (libc, libgcc, …), GPL-2.0-or-later 4,
  LGPL-2.1-or-later 3, BSD-4-Clause 2, and one each of Zlib, LGPL-2.1+, GPL-3.0-or-later (GRUB),
  "GPL-2.0-or-later MPL-2.0", GPL-2.0+, "GPL-2.0 BSD-3-Clause", BSD-3-Clause. Every package has a licence field.

## 7. Tools of 6A criterion 3, with versions

Presence from the guest's `command -v` loop (and `ls` in the unpacked root, same result). Versions from the guest's
own output where it exists, otherwise the installed package.

| Tool | Present | Version |
| :--- | :--- | :--- |
| procd (init) | yes, `/sbin/procd` | `procd 2026.03.13~58eb263d-r1` (package) |
| systemd, OpenRC | no | — |
| nftables `nft` | yes, `/usr/sbin/nft` | `nftables v1.1.6 (Commodore Bullmoose #7)`; package `nftables-json 1.1.6-r2` |
| dnsmasq | yes | `Dnsmasq version 2.93`; package `dnsmasq 2.93-r1` (not `dnsmasq-full`) |
| hostapd | **no** | feed: section 9a |
| wpa_supplicant | **no** | feed: section 9a |
| iw | **no** | feed: section 9a (installed in the guest for the hwsim test) |
| pppd | yes | `pppd version 2.5.2`; `ppp 2.5.2-r4`, `ppp-mod-pppoe 2.5.2-r4` |
| WireGuard (`wg`, module) | **no** | feed: section 9a (module installed in the guest) |
| OpenVPN | **no** | feed: section 9a |
| `ip` | yes, but the **BusyBox applet** (`/sbin/ip -> ../bin/busybox`) | BusyBox v1.37.0; `ip -br addr` printed the BusyBox usage. `ip link add … type bridge` and `type wireguard` worked |
| `tc`, `ethtool`, `iptables`, `ss` | no | — (`iptables-nft` and friends are in the target feed `packages/`) |
| DHCP client | `udhcpc` (BusyBox), `odhcp6c 2026.06.27~10a52220-r3` | — |
| DHCP/RA server | `odhcpd-ipv6only 2026.06.29~5d7be43f-r1` and dnsmasq | — |
| SSH | `dropbear` | `Dropbear v2025.89` |
| Web | `uhttpd 2026.08.24~373145f7-r1` with LuCI `26.288.56989~75a306d` | — |
| HTTP client | `uclient-fetch` (`wget` is a link to it) | `uclient-fetch 2026.06.07~daad21fa-r1` |
| BusyBox | yes | `busybox 1.37.0-r6`; banner `BusyBox v1.37.0 (2026-10-17 21:45:38 UTC)` |
| Script runtime | `ucode` (`/usr/bin/ucode`) | `ucode 2026.01.16~85922056-r2`; `lua`, `perl` absent |
| `nproc` | **no** (guest: `-ash: nproc: not found`) | — |

## 8. Python (6A criterion 8)

**No Python is in the image.** The checks:

- In the unpacked squashfs root, `find . -iname "python*" -o -name "*.py" -o -name "*.pyc"` found 0 files;
  `grep -rlI "^#!.*python" .` found 0 and `grep -rlI "^#!.*perl" .` found 0.
- In the guest, `find / -xdev …` and `find /rom -xdev …` both found 0, and `command -v python python3 perl lua`
  printed absent for all four.
- `python3` / `python3-base` `3.13.18-r1` are in the aarch64 packages feed but not installed (section 9a).

## 9. Own network and firewall manager, and how each is switched off (6A criterion 5)

The guest runs **netifd** (`/sbin/netifd`, package `netifd 2026.02.26~cbb83a18-r2`) and **firewall4**
(`/sbin/fw4`, ucode, package `firewall4 2025.03.17~b6e51575-r2`). The ruleset is one table, `table inet fw4`
(guest `nft list tables`), 145 lines (`nft list ruleset | wc -l`), the same as on x86-64.

**Default configuration** (guest `uci show network` and `uci show firewall`, `/etc/board.d/99-default_network`):

- `eth0` (first virtio NIC) is bridged into `br-lan`, static `192.168.1.1/24`, `ip6assign 60`.
- `eth1` is `wan` (DHCP) and `wan6` (DHCPv6), only if `eth1` exists at first boot (`[ -d /sys/class/net/eth1 ]`).
- A ULA prefix is generated on the first boot; dnsmasq serves DHCP and DNS on the LAN.
- Firewall: defaults input/forward `REJECT`; zone `lan` accept; zone `wan` (`wan`, `wan6`) input `REJECT`, forward
  `DROP`, `masq` and `mtu_fix`; the usual rules Allow-DHCP-Renew, -Ping, -IGMP, -DHCPv6, -MLD, -ICMPv6-Input/-Forward,
  -IPSec-ESP, -ISAKMP.

Coexistence tests in the guest (one boot, all probe lines exit 0):

| Test | Command | Seen |
| :--- | :--- | :--- |
| Own nft table across fw4 | `nft add table inet nmbtest` + chain; `fw4 stop`; `fw4 start`; `fw4 reload` | after `fw4 stop` only `inet nmbtest`; after `start` and `reload` both tables: **kept** |
| Firewall service restart | `/etc/init.d/firewall restart` | both tables: **kept** |
| Firewall service stop | `/etc/init.d/firewall stop` | **empty list, NMB's table gone too** (the service stop flushes the whole ruleset, as on x86-64). `start` brought back only `inet fw4` |
| Firewall disable | `disable`; `enabled; echo $?`; `ls /etc/rc.d \| grep -c firewall` | `enabled` gave 1, 0 links: off at boot. `enable` gave 1 link back |
| Own bridge and an address on `br-lan` | `ip link add nmbbr0 type bridge`, `198.51.100.77/24`; `203.0.113.9/24` on `br-lan` | both present |
| netifd reload | `/etc/init.d/network reload`, 10 s | both **kept** |
| netifd restart | `/etc/init.d/network restart`, 15 s | `nmbbr0` address kept; **`203.0.113.9` on `br-lan` gone** (netifd rebuilds `br-lan`) |
| netifd stop / start | `network stop`; `pgrep -l netifd`; `network start` | no netifd (pgrep rc 1), `nmbbr0` still addressed; netifd back after `start` |
| netifd disable | `network disable`; `ls /etc/rc.d \| grep -c network` | 0 links: off at boot. `enable` gave 2 links back |

**What this means for criterion 5.** The same as on x86-64: both managers can be switched off cleanly with `disable`
(plus `stop` for netifd); the firewall service's `stop` flushes the **whole** nftables ruleset, so NMB must use
`fw4 stop`, or `disable` and a reboot; netifd keeps devices it does not manage and foreign addresses on its own
devices across `reload`, and drops them on `restart`.

### 9a. Feed lookup, and kernel modules added in the guest

The WAN side reached the internet over user-mode networking (`ping 192.0.2.2` 2/2, DNS resolved
`downloads.openwrt.org`). `apk update` printed `OK: 11171 distinct packages available` (rc 0). Then, for each name,
`apk search $p | grep "^$p-[0-9]"`. **None of these packages is installed in the image.**

| Package | Version in the 25.12.6 aarch64 feeds |
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
| `kmod-veth`, `kmod-bonding` | `6.12.112-r1` |
| `python3` / `python3-base` | `3.13.18-r1` (feed only; never installed for NMB, criterion 8) |

Then, in the same guest (a throw-away copy of the image), `apk add kmod-mac80211-hwsim iw kmod-wireguard` installed 20
packages (`OK: 16.9 MiB in 225 packages`, rc 0). This shows that the published kmods load into the image's kernel:

- `mac80211_hwsim` loaded through kmodloader on install (`compat: loading out-of-tree module taints kernel`,
  `Loading modules backported from Linux version v6.18.55`, `mac80211_hwsim: initializing netlink`). `iw list` showed
  two radios, `phy0` and `phy1`; `iw dev` listed no interface (none is created until one is configured).
- `wireguard` loaded (`WireGuard 1.0.0 loaded`); `ip link add nmbwg0 type wireguard` created the device (rc 0).
- `lsmod` afterwards: `cfg80211`, `mac80211`, `mac80211_hwsim`, `compat`, `wireguard`, `udp_tunnel`,
  `ip6_udp_tunnel`, `libchacha20poly1305`, `libcurve25519_generic`.

## 10. Installed packages (205, `lib/apk/db/installed`; the same in both images)

`apk-mbedtls 3.0.5-r3`, `attendedsysupgrade-common 10`, `base-files 1718~981c978eff`, `blkid 2.41.5-r1`, `busybox 1.37.0-r6`, `ca-bundle 20260816-r1`, `cgi-io 2026.07.21~31cb3c89-r1`, `dnsmasq 2.93-r1`, `dropbear 2025.89-r3`, `e2fsprogs 1.47.3-r1`, `firewall4 2025.03.17~b6e51575-r2`, `fstools 2026.05.23~16718b6e-r1`, `fwtool 2025.10.03~04cd252e-r1`, `getrandom 2025.10.30~6f78fa49-r1`, `grub2-efi-arm 2.12-r1`, `jansson4 2.14.1-r1`, `jshn 2026.06.19~7dd12784-r1`, `jsonfilter 2026.03.16~b9034210-r1`, `kernel 6.12.112~67e205f26d0458c4d12f1fd315216398-r1`, `kmod-acpi-mdio 6.12.112-r1`, `kmod-amazon-ena 6.12.112-r1`, `kmod-atlantic 6.12.112-r1`, `kmod-bcmgenet 6.12.112-r1`, `kmod-crypto-aead 6.12.112-r1`, `kmod-crypto-crc32 6.12.112-r1`, `kmod-crypto-crc32c 6.12.112-r1`, `kmod-crypto-ctr 6.12.112-r1`, `kmod-crypto-gcm 6.12.112-r1`, `kmod-crypto-geniv 6.12.112-r1`, `kmod-crypto-gf128 6.12.112-r1`, `kmod-crypto-ghash 6.12.112-r1`, `kmod-crypto-hash 6.12.112-r1`, `kmod-crypto-hmac 6.12.112-r1`, `kmod-crypto-manager 6.12.112-r1`, `kmod-crypto-null 6.12.112-r1`, `kmod-crypto-rng 6.12.112-r1`, `kmod-crypto-seqiv 6.12.112-r1`, `kmod-crypto-sha3 6.12.112-r1`, `kmod-crypto-sha512 6.12.112-r1`, `kmod-dwmac-imx 6.12.112-r1`, `kmod-dwmac-rockchip 6.12.112-r1`, `kmod-dwmac-sun8i 6.12.112-r1`, `kmod-e1000e 6.12.112-r1`, `kmod-fixed-phy 6.12.112-r1`, `kmod-fs-vfat 6.12.112-r1`, `kmod-fsl-dpaa1-net 6.12.112-r1`, `kmod-fsl-dpaa2-net 6.12.112-r1`, `kmod-fsl-enetc-net 6.12.112-r1`, `kmod-fsl-fec 6.12.112-r1`, `kmod-fsl-mc-dpio 6.12.112-r1`, `kmod-fsl-pcs-lynx 6.12.112-r1`, `kmod-fsl-xgmac-mdio 6.12.112-r1`, `kmod-gpio-pca953x 6.12.112-r1`, `kmod-hwmon-core 6.12.112-r1`, `kmod-i2c-core 6.12.112-r1`, `kmod-i2c-mux 6.12.112-r1`, `kmod-i2c-mux-pca954x 6.12.112-r1`, `kmod-lib-crc-ccitt 6.12.112-r1`, `kmod-lib-crc-itu-t 6.12.112-r1`, `kmod-lib-crc16 6.12.112-r1`, `kmod-lib-crc32c 6.12.112-r1`, `kmod-libphy 6.12.112-r1`, `kmod-macsec 6.12.112-r1`, `kmod-marvell-mdio 6.12.112-r1`, `kmod-mdio-bcm-unimac 6.12.112-r1`, `kmod-mdio-bus-mux 6.12.112-r1`, `kmod-mdio-devres 6.12.112-r1`, `kmod-mdio-gpio 6.12.112-r1`, `kmod-mdio-thunder 6.12.112-r1`, `kmod-mii 6.12.112-r1`, `kmod-mvneta 6.12.112-r1`, `kmod-mvpp2 6.12.112-r1`, `kmod-net-selftests 6.12.112-r1`, `kmod-nf-conntrack 6.12.112-r1`, `kmod-nf-conntrack6 6.12.112-r1`, `kmod-nf-flow 6.12.112-r1`, `kmod-nf-log 6.12.112-r1`, `kmod-nf-log6 6.12.112-r1`, `kmod-nf-nat 6.12.112-r1`, `kmod-nf-reject 6.12.112-r1`, `kmod-nf-reject6 6.12.112-r1`, `kmod-nfnetlink 6.12.112-r1`, `kmod-nft-core 6.12.112-r1`, `kmod-nft-fib 6.12.112-r1`, `kmod-nft-nat 6.12.112-r1`, `kmod-nft-offload 6.12.112-r1`, `kmod-nls-base 6.12.112-r1`, `kmod-nls-cp437 6.12.112-r1`, `kmod-nls-iso8859-1 6.12.112-r1`, `kmod-nls-utf8 6.12.112-r1`, `kmod-octeontx2-net 6.12.112-r1`, `kmod-of-mdio 6.12.112-r1`, `kmod-pcs-xpcs 6.12.112-r1`, `kmod-phy-aquantia 6.12.112-r1`, `kmod-phy-bcm7xxx 6.12.112-r1`, `kmod-phy-broadcom 6.12.112-r1`, `kmod-phy-marvell 6.12.112-r1`, `kmod-phy-marvell-10g 6.12.112-r1`, `kmod-phy-realtek 6.12.112-r1`, `kmod-phy-smsc 6.12.112-r1`, `kmod-phylib-broadcom 6.12.112-r1`, `kmod-phylink 6.12.112-r1`, `kmod-ppp 6.12.112-r1`, `kmod-pppoe 6.12.112-r1`, `kmod-pppox 6.12.112-r1`, `kmod-pps 6.12.112-r1`, `kmod-ptp 6.12.112-r1`, `kmod-regmap-core 6.12.112-r1`, `kmod-regmap-i2c 6.12.112-r1`, `kmod-renesas-net-avb 6.12.112-r1`, `kmod-rtc-rx8025 6.12.112-r1`, `kmod-sfp 6.12.112-r1`, `kmod-slhc 6.12.112-r1`, `kmod-stmmac-core 6.12.112-r1`, `kmod-thunderx-net 6.12.112-r1`, `kmod-vmxnet3 6.12.112-r1`, `kmod-wdt-sp805 6.12.112-r1`, `libblkid1 2.41.5-r1`, `libblobmsg-json20260213 2026.06.19~7dd12784-r1`, `libc 1.2.5-r5`, `libcomerr0 1.47.3-r1`, `libe2p2 1.47.3-r1`, `libext2fs2 1.47.3-r1`, `libf2fs6 1.16.0-r4`, `libgcc1 14.3.0-r5`, `libiwinfo-data 2026.05.26~66bdd1a0-r1`, `libiwinfo20230701 2026.05.26~66bdd1a0-r1`, `libjson-c5 0.18-r1`, `libjson-script20260213 2026.06.19~7dd12784-r1`, `liblucihttp-ucode 2023.03.15~9b5b683f-r1`, `liblucihttp0 2023.03.15~9b5b683f-r1`, `libmbedtls21 3.6.7-r1`, `libmnl0 1.0.5-r1`, `libnftnl11 1.3.1-r1`, `libnl-tiny1 2025.12.02~40493a65-r1`, `libpthread 1.2.5-r5`, `librt 1.2.5-r5`, `libsmartcols1 2.41.5-r1`, `libss2 1.47.3-r1`, `libubox20260213 2026.06.19~7dd12784-r1`, `libubus20251202 2026.06.28~24864e78-r2`, `libuci20250120 2025.12.02~66127cd7-r1`, `libuclient20201210 2026.06.07~daad21fa-r1`, `libucode20230711 2026.01.16~85922056-r2`, `libudebug 2026.01.16~875e1a7a`, `libustream-mbedtls20201210 2026.03.01~99f1c0db-r1`, `libuuid1 2.41.5-r1`, `logd 2025.10.30~6f78fa49-r1`, `luci 26.288.56989~75a306d`, `luci-app-attendedsysupgrade 26.288.56989~75a306d`, `luci-app-firewall 26.288.56989~75a306d`, `luci-app-package-manager 26.288.56989~75a306d`, `luci-base 26.288.56989~75a306d`, `luci-lib-uqr 26.288.56989~75a306d`, `luci-light 26.288.56989~75a306d`, `luci-mod-admin-full 26.288.56989~75a306d`, `luci-mod-network 26.288.56989~75a306d`, `luci-mod-status 26.288.56989~75a306d`, `luci-mod-system 26.288.56989~75a306d`, `luci-proto-ipv6 26.288.56989~75a306d`, `luci-proto-ppp 26.288.56989~75a306d`, `luci-ssl 26.288.56989~75a306d`, `luci-theme-bootstrap 26.288.56989~75a306d`, `mkf2fs 1.16.0-r4`, `mtd 27`, `netifd 2026.02.26~cbb83a18-r2`, `nftables-json 1.1.6-r2`, `odhcp6c 2026.06.27~10a52220-r3`, `odhcpd-ipv6only 2026.06.29~5d7be43f-r1`, `openwrt-keyring 2025.12.10~f0670054-r1`, `owut 2026.08.30~9b51eeec-r1`, `partx-utils 2.41.5-r1`, `ppp 2.5.2-r4`, `ppp-mod-pppoe 2.5.2-r4`, `procd 2026.03.13~58eb263d-r1`, `procd-seccomp 2026.03.13~58eb263d-r1`, `procd-ujail 2026.03.13~58eb263d-r1`, `px5g-mbedtls 11`, `rpcd 2026.10.04~d99f703e-r1`, `rpcd-mod-file 2026.10.04~d99f703e-r1`, `rpcd-mod-iwinfo 2026.10.04~d99f703e-r1`, `rpcd-mod-luci 20240305-r1`, `rpcd-mod-rpcsys 2026.10.04~d99f703e-r1`, `rpcd-mod-rrdns 20170710`, `rpcd-mod-ucode 2026.10.04~d99f703e-r1`, `ubox 2025.10.30~6f78fa49-r1`, `ubus 2026.06.28~24864e78-r2`, `ubusd 2026.06.28~24864e78-r2`, `uci 2025.12.02~66127cd7-r1`, `uclient-fetch 2026.06.07~daad21fa-r1`, `ucode 2026.01.16~85922056-r2`, `ucode-mod-fs 2026.01.16~85922056-r2`, `ucode-mod-html 1`, `ucode-mod-log 2026.01.16~85922056-r2`, `ucode-mod-math 2026.01.16~85922056-r2`, `ucode-mod-ubus 2026.01.16~85922056-r2`, `ucode-mod-uci 2026.01.16~85922056-r2`, `ucode-mod-uclient 2026.06.07~daad21fa-r1`, `ucode-mod-uloop 2026.01.16~85922056-r2`, `uhttpd 2026.08.24~373145f7-r1`, `uhttpd-mod-ubus 2026.08.24~373145f7-r1`, `urandom-seed 3`, `urngd 2025.10.03~f17e33d9-r1`, `usign 2025.10.03~c4c72b1b-r1`, `zlib 1.3.1-r1`

## 11. Size on disk

- Download: 13,425,520 bytes (squashfs) and 14,669,158 bytes (ext4), as `.img.gz`.
- Disk image: 243,564,032 bytes (232.3 MiB) for each; the ESP is 128 MiB of it.
- Kernel `vmlinux` (uncompressed Image): 20,222,464 bytes. The initramfs kernel: 35,754,496 bytes.
- squashfs root: 4,619,820 bytes compressed. Unpacked, the root is 15,398,567 bytes (`du -sb .`), the same for the
  ext4 root. Of that: `usr` 6,313,736, `lib` 6,228,652 (`lib/modules` 4,881,889, `lib/firmware` 0), `www` (LuCI)
  1,337,317, `sbin` 721,273, `bin` 474,911, `etc` 322,556.
- Guest `df -k` after `init complete` (squashfs image):
  - `/rom`: 4,608 kB, 100 % used
  - `/overlay` (ext4 on loop0): 89,911 kB, 61 kB used, 82,714 kB free
  - `/tmp`: tmpfs, 243,736 kB
  - At the end of the run (after `apk update` and 20 added packages): `/overlay` 2,774 kB used, 80,001 kB free;
    `/tmp` 1,776 kB used.
- ext4 image: the 104 MiB root has 20,553 of 26,624 4-KiB blocks free (about 80.3 MiB), from `debugfs -R stats`.

## 12. QEMU boot test (squashfs image)

**How it was run.** One boot, on a fresh copy of the decompressed image, under `flock -w 3600 ~/nmb-work/qemu.lock`
for the whole run (06:23:45–06:36:34 UTC). The serial console was driven by a small Python script (`python3 -I`,
kept outside the data folder) over a Unix socket; input lines were sent in 32-byte chunks. The command line, which
**F-06 can reuse as the ARM64 guest** (`<work>` is the run folder; `<xfer>` holds the transfer disk and share):

```
qemu-system-aarch64 -M virt -cpu cortex-a72 -accel tcg,thread=multi,tb-size=128 -smp 2 -m 512
  -nographic -monitor none -serial unix:<work>/serial.sock,server=on,wait=on
  -bios /usr/share/qemu-efi-aarch64/QEMU_EFI.fd
  -drive id=d0,file=<work>/disk.img,format=raw,if=none -device virtio-blk-pci,drive=d0
  -drive id=d1,file=<xfer>/xfer.img,format=raw,if=none,readonly=on -device virtio-blk-pci,drive=d1
  -fsdev local,id=fs0,path=<xfer>/share,security_model=none,readonly=on -device virtio-9p-pci,fsdev=fs0,mount_tag=nmbshare
  -netdev user,id=lan0,net=198.51.100.0/24,restrict=on -device virtio-net-pci,netdev=lan0,mac=02:00:00:00:02:01
  -netdev user,id=wan0,net=192.0.2.0/24             -device virtio-net-pci,netdev=wan0,mac=02:00:00:00:02:02
```

- QEMU 8.2.2 (Debian 1:8.2.2+ds-0ubuntu1.18), firmware `QEMU_EFI.fd` from package `qemu-efi-aarch64 2024.02-2ubuntu0.10`
  (UEFI banner "version 2024.02-2ubuntu0.10"). TCG with two vCPUs, 512 MB. The `u-boot-qemu_armv8/u-boot.bin` of the
  tree was not needed and not tried.
- **Guest devices.** `vda` is the image (`vda1`, `vda2`, `vda128`), `vdb` the transfer disk; `eth0` is the first NIC
  (LAN, `br-lan`), `eth1` the second (WAN, DHCP lease `192.0.2.15` from the user-mode network).
- **Putting a file into the guest: both ways work** (each probe rc 0, SHA-256 equal on both sides).
  - 9p (built into the kernel): `mount -t 9p -o trans=virtio,version=9p2000.L,ro nmbshare /mnt/x9`. Easiest for a
    fresh binary on every test run; the share is read-only for the guest.
  - Second virtio disk: a FAT image made on the PC without root (`mkfs.vfat -C -n NMBXFER xfer.img 16384`, then
    `mcopy -i xfer.img <file> ::/`), in the guest `mount -t vfat -o ro /dev/vdb /mnt/xb` (`kmod-fs-vfat` is installed).
  - Results come back over the serial console, or by writing to a writable second disk (not tried).
- **Clean end.** `poweroff` gave `reboot: Power down`; QEMU exit 0, driver exit 0, `flock` exit 0; 767.8 s from QEMU
  start to power-down. The host load average was about 3.5–4.7 during the run (other packages were working).

Figures (TCG on an x86 host, two vCPUs, 512 MB, loaded host; **not** a performance figure for real hardware):

- **Boot time** (host clock from QEMU start):
  - kernel start (`Linux version` on the console): 14.7 s, including UEFI and the 5 s GRUB timeout;
  - `Please press Enter to activate this console`: 76.4 s; root prompt 86.7 s;
  - kernel timestamps: `Run /sbin/init` 2.32 s, `init: - preinit -` 5.38 s, overlay formatted at 31–33.6 s,
    `procd: - early -` 56.19 s, `procd: - init -` 60.67 s; `/etc/modules.d` loaded 77.6–82.9 s; `br-lan` forwarding at
    156.0 s.
  - `procd: - init complete -` (`logread`): found by a 5 s polling loop that ended at kernel uptime 287.8 s
    (`/proc/uptime`). The guest log has 177 s between `procd: - init -` and `init complete`, which would put it at
    about 238 s, but the guest clock may have been set by ntpd in between (not checked), so only "at most 288 s" is
    certain. `logread` failed with `Failed to find log object` until logd was up.
  - The first boot includes formatting the overlay.
- **Memory.** Kernel line `Memory: 450944K/524288K available`. Guest `free` after `init complete`: total 487,472 kB,
  used 36,032 kB, free 424,944 kB, available 413,520 kB. At the end (uptime 715 s, after `apk update`, 20 added
  packages, mac80211/hwsim and WireGuard loaded): used 42,592 kB, available 400,752 kB. No swap.
- **Storage after boot:** section 11.
- **CPU seen by the guest** (`/proc/cpuinfo`): 2 × part `0xd08` (Cortex-A72), features `fp asimd evtstrm aes pmull
  sha1 sha2 crc32 cpuid`.
- **Listening ports** (guest `netstat -tulnp`):
  - TCP 22: dropbear, all addresses (IPv4 and IPv6)
  - TCP 80 and 443: uhttpd, all addresses
  - TCP and UDP 53: dnsmasq, on loopback, the LAN address, the WAN address `192.0.2.15`, IPv6 link-local and the ULA
  - UDP 546: odhcp6c; UDP 547: odhcpd
- **Running services** (guest `ps w`): procd, ubusd, logd, rpcd, urngd, dropbear, netifd, odhcpd, udhcpc on eth1,
  odhcp6c on eth1, uhttpd, dnsmasq (inside `ujail`, user `dnsmasq`), ntpd (inside `ujail`, user `ntp`), and
  askfirst with the login shell on `ttyAMA0`.
- **Security observations** (same as x86-64):
  - The banner says "There is no root password defined on this device!".
  - `/etc/config/dropbear` has `PasswordAuth 'on'` and `RootPasswordAuth 'on'`.
  - uhttpd listens on `0.0.0.0:80/443` and `[::]:80/443`. The default firewall rejects input from the wan zone.
- **Kernel messages:** no RCU stall and no oops in the console log (`grep`); the GPT backup warning of section 2.

## 13. Fields not determined

- **Boot time and memory on real ARM hardware:** not determined. Only TCG emulation on a loaded x86 host was possible;
  the figures above are emulation figures.
- **Exact time of `init complete`:** only an upper bound (288 s kernel uptime), see section 12.
- **Ext4 image boot:** not run. Taken apart only; its file and package lists equal the squashfs image's.
- **Initramfs kernel boot:** not run; the combined-efi squashfs image booted, which the task asked for first.
- **usign signature of `sha256sums`:** not checked (no usign tool, no usign public key without network). The gpg
  signature was checked.
- **Running the SDK/toolchain compilers:** not done. Their versions come from `info.mk`, `version.mk`, `version.h`,
  the kernel's version string and `strings`; no downloaded binary was executed on the PC.
- **`/dev/watchdog` behaviour:** not determined; QEMU `virt` offers no SP805 watchdog, and no soft watchdog was loaded.

## 14. Tools missing on the PC

`cpio`, `bsdtar`, `xorriso`, `parted`, `expect`, `socat`, `extract-ikconfig` and `usign`/`signify` are missing, and
there is no sudo, so no loop mounts.

Present and used: `gzip`, `zstd`, `tar`, `sgdisk`, `fdisk`, `mtools`, `mkfs.vfat`, `debugfs`, `unsquashfs`
(squashfs-tools), `7z`, `file`, `strings`, `gpg`, `qemu-system-aarch64` 8.2.2, `qemu-efi-aarch64` 2024.02, `flock`.

## 15. SDK, ImageBuilder and toolchain (work plan 5.3 point 4)

The three archives were not unpacked whole: each was listed with `zstd -dc … | tar -tv` (exit 0 for both stages),
and only the files named here were extracted into fresh folders (then deleted).

| Tree | Entries | Unpacked size (sum of `tar -tv`) | What is in it |
| :--- | ---: | ---: | :--- |
| SDK `openwrt-sdk-25.12.6-armsr-armv8_gcc-14.3.0_musl.Linux-x86_64` | 36,721 | 1,217,986,617 bytes | `staging_dir/host` (19,446 entries), `build_dir/target-aarch64_generic_musl` (13,113, with the kernel tree's `.config` and headers), `staging_dir/toolchain-aarch64_generic_gcc-14.3.0_musl` (3,328), `feeds.conf.default` pinned to the release commits |
| ImageBuilder `openwrt-imagebuilder-25.12.6-armsr-armv8.Linux-x86_64` | 6,072 | 286,972,688 bytes | `build_dir` (5,158), `staging_dir` (612), `packages/` with only `base-files`, `kernel` and `libc`, `keys/openwrt-25.12.pem`, `repositories` (the 8 release feeds), `.config` (`CONFIG_USE_APK=y`, `CONFIG_SIGNATURE_CHECK=y`, `TARGET_KERNEL_PARTSIZE=128`, `TARGET_ROOTFS_PARTSIZE=104`, squashfs, ext4, targz, cpiogz, initramfs, `GRUB_EFI_IMAGES`, `GRUB_TIMEOUT="5"`) |
| Toolchain `openwrt-toolchain-25.12.6-armsr-armv8_gcc-14.3.0_musl.Linux-x86_64` | 3,300 | 331,343,829 bytes | `toolchain-aarch64_generic_gcc-14.3.0_musl` (3,288), `version.mk` (`REVISION:=r33311-981c978eff`), `LICENSES/` |
| `kernel-debug.tar.zst` | 1,364 | 482,848,320 bytes | `debug/vmlinux` (93,627,008 bytes, unstripped) and 1,361 unstripped `.ko` under `debug/modules/` |
| `llvm-bpf-21.1.6.Linux-x86_64.tar.zst` | 405 | not summed | LLVM 21.1.6 for BPF objects (listed only) |

**Toolchain versions** (the same in the SDK and the stand-alone toolchain):

- Cross prefix `aarch64-openwrt-linux-musl-` (`info.mk`: `TARGET_CROSS`).
- **GCC 14.3.0** (`info.mk` `GCC_VERSION=14.3.0`; `strings` of `.aarch64-openwrt-linux-musl-gcc-14.3.0.bin`:
  "(OpenWrt GCC 14.3.0 r33311-981c978eff)"). The compiler is an x86-64 ELF; the visible `aarch64-openwrt-linux-musl-gcc`
  is a 5-line wrapper that runs it through the SDK's own `staging_dir/host/lib/ld-linux-x86-64.so.2`.
- **musl 1.2.5** (`info.mk` `LIBC_TYPE=musl`, `LIBC_VERSION=1.2.5`; `strings` of `lib/libc.so` contains `1.2.5`).
  `lib/libc.so` is byte-identical in the SDK and the toolchain (`cmp`) and matches the image's `libc 1.2.5-r5`.
- **GNU Binutils 2.44** (`strings` of `aarch64-openwrt-linux-musl/bin/.ld.bin`: "(GNU Binutils) 2.44"; the kernel was
  linked with "GNU ld (GNU Binutils) 2.44").
- Kernel headers 6.12.112 (`include/linux/version.h`: `LINUX_VERSION_CODE 396400`, 6.12.112).
- Also present: `g++` with `libstdc++.so.6.0.33`, `libatomic`, `gdb`, `lto-dump`. The toolchain's file list differs from
  the SDK copy's in 168 `diff` lines (3,288 against 3,328 entries); not examined further.

**What they would be used for** (5.3 point 4):

- **SDK:** building a first `nmbd` apk package for `aarch64_generic` against the same feeds (`feeds.conf.default`
  pins the release's commits), and the later C11/musl builds with the exact compiler and musl of the image.
- **Toolchain alone:** cross-compiling C11/musl binaries for aarch64 musl without the package build system.
- **ImageBuilder:** assembling an NMB test image from the published packages plus an `nmbd` apk, without rebuilding
  the base; its `repositories` point at the release feeds, so a build needs network access or a local mirror.
- **Kernel config** (ikconfig and the trees' `.config`): the options NMB needs are there as built-ins or published
  kmods, except **Landlock** (no `CONFIG_SECURITY`) and in-kernel `cfg80211` (backports packages instead). No own
  kernel is needed for anything else in section 4.
- **`kernel-debug`:** symbols for kernel crash analysis on this exact build, if one is ever needed.
