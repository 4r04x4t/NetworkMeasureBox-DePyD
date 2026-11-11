# Image record: OpenWrt 25.12.6 armsr/armv7 (32-bit ARM, squashfs and ext4, combined-efi, with its tree)

Package I-03 (work plan section 5.4). The method is work plan 5.1, and every field of 5.2 is below.
Each value names the command that produced it. A value that could not be read says "not determined" and why.
Nothing here comes from memory. Only lists, versions and checksums are recorded. No file taken out of the image
is in this repository. The headings and their order are the same as in the I-01 record (x86-64), so that I-09 can
compare the two; section 15 is extra and covers the SDK, ImageBuilder and toolchain of the tree (5.3 point 4).

- Date of the analysis: 2026-11-11.
- Where it ran: dev PC, Windows with WSL2 Ubuntu 24.04, about 1.8 GB RAM and 4 threads, no root. The QEMU guests ran
  on the same PC (TCG, because an ARM guest cannot use KVM on an x86 host).
- Work folder: `~/nmb-work/I-03/` (outside every repository; one fresh sub-folder per unpack and per boot; the
  helper scripts in their own sub-folder `tools/`, away from the unpacked data).
- Material, read only: `~/src/NMB-depythonization-firmwares/openwrt-target-trees/openwrt-25.12.6-armsr-armv7/`
  (the full target folder of the release: images, `packages/`, `kmods/`, `u-boot-qemu_armv7/`, SDK, ImageBuilder,
  toolchain, `kernel-debug`, `llvm-bpf`).

## 1. Files and checksums

| File | Size (bytes) | SHA-256 (publisher list) | Check |
| :--- | ---: | :--- | :--- |
| `openwrt-25.12.6-armsr-armv7-generic-squashfs-combined-efi.img.gz` | 8,006,271 | `79f647e56253e9cf450c8debc21a284156bf5661bf37b5632fb90b7f921efb6a` | OK |
| `openwrt-25.12.6-armsr-armv7-generic-ext4-combined-efi.img.gz` | 9,135,764 | `875de39040292f90e27c259c2e1d2690623d994699e2fe383bc0eb4cbac31c36` | OK |
| `openwrt-25.12.6-armsr-armv7-generic-squashfs-rootfs.img.gz` | 3,939,035 | `43c5db4f1b3743ac81675e2421b5131243b0cb457caf7aec80117f6153f406a8` | OK |
| `openwrt-25.12.6-armsr-armv7-generic-ext4-rootfs.img.gz` | 5,070,300 | `744ee8a26bcd6f738f9c675a13d798e2d7c8e840661ec33e0aee688560687580` | OK |
| `openwrt-25.12.6-armsr-armv7-generic-kernel.bin` | 3,568,128 | `d47c4d36e5bd907dec8e22d2c13229cb7351d554dced5792135035686162bf44` | OK |
| `openwrt-25.12.6-armsr-armv7-generic-initramfs-kernel.bin` | 7,000,576 | `2c9ee0f4ad9f825b90c034d39d4a3b2baa31583ab98255e9e4b36239522fd0ff` | OK |
| `openwrt-25.12.6-armsr-armv7-generic-targz-rootfs.tar.gz` | 4,953,339 | `1b9950dd8248040afc8c399b20f4abceeba11a34d167f8b6fb1977ee69f647b3` | OK |
| `openwrt-25.12.6-armsr-armv7-rootfs.tar.gz` | 4,214,260 | `f0e25d1b7abd6d317e416aed915820d85b8f1e2c457c1b2c38522e91c1e31aca` | OK |
| `openwrt-25.12.6-armsr-armv7-rootfs.cpio.gz` | 4,212,217 | `91b4f0a06084af3a4902bbd9929887c846fbff100765cbf20bff7a43605690a1` | OK |
| `u-boot-qemu_armv7/u-boot.bin` | 1,198,016 | `7904a6930e68b26b66477563aabf3a977a150423038f48c57164e01133a3d704` | OK |
| `openwrt-sdk-25.12.6-armsr-armv7_gcc-14.3.0_musl_eabi.Linux-x86_64.tar.zst` | 231,153,820 | `b29c5b96d757c3518a4f4df18e3187e16de618381625847b8874461935358179` | OK |
| `openwrt-toolchain-25.12.6-armsr-armv7_gcc-14.3.0_musl_eabi.Linux-x86_64.tar.zst` | 55,615,731 | `cc84c103e3c43fec14af2eff74e61d6180d0b187d6e97eef79a38fcfde179771` | OK |
| `openwrt-imagebuilder-25.12.6-armsr-armv7.Linux-x86_64.tar.zst` | 50,135,792 | `6109c6719b04a62dfc244e0a9754ed7e06dfb7d45f9dad8bdd40d57d08527ee2` | OK |
| `kernel-debug.tar.zst` | 99,771,731 | `709c7270ff95cf0febf3929f3560733aed80aadb53e24f21d73659e126d10384` | OK |
| `llvm-bpf-21.1.6.Linux-x86_64.tar.zst` | 51,187,706 | `3f07164c83d2965bc8685bbb1f1efda4539e9c86e42779f52808655b039df574` | OK |
| `kmods/…/kmod-ikconfig-6.12.112-r1.apk` | — | `e835c6fe3aedc160af6270b8a54b0c392f4492c49ebb6b0b9192fccef686ae9f` | OK |

- **How it was checked.** The tree carries the publisher's own `sha256sums` (1,234 lines, SHA-256
  `d1e5bcaf1d46b356b231084c84d60a80165a4ddf33e12369d93ec6023c6385ae`), with `sha256sums.asc` (OpenPGP) and
  `sha256sums.sig` (usign). No network fetch was needed for the list.
  - Top-level files: `grep -v / sha256sums | sha256sum -c` printed `OK` for all 20 listed files (exit 0). These are
    the 14 top-level files of the table (all rows but `u-boot.bin` and the kmod), plus `config.buildinfo`,
    `feeds.buildinfo`, `version.buildinfo`, `profiles.json`, the `.manifest` and the `.bom.cdx.json`.
  - Sub-folders: all 1,214 listed entries under `kmods/`, `packages/` and `u-boot-qemu_armv7/` exist on disk and
    were checked with `sha256sum -c --quiet` (exit 0, no failure). On disk there are 1,218 files there; the four
    not in the list are the signatures `packages.adb.asc` and `packages.adb.sig` of the two package indexes, which
    the list cannot contain.
- **OpenPGP signature of the list.** `gpg --verify sha256sums.asc sha256sums` in a fresh, isolated `GNUPGHOME` gave
  "Good signature from OpenWrt Build System (Nitrokey3)" (exit 0). The signature was made on 2026-10-18 by
  subkey `92C5 61DE 55AE 6552 F3C7 36B8 2B01 5109 0606 D1D9`, primary key
  `8A8B C12F 46B8 36C0 F9CD B36F 1D53 D187 7742 E911`.
  - **Where the key came from.** No key was fetched in this package. The public key was exported from I-01's
    keyring (`~/nmb-work/I-01/gnupg`), where I-01 had fetched it from keys.openpgp.org and compared its
    fingerprints with OpenWrt's keyring repository. So this check depends on that earlier network fetch.
    gpg says "not certified with a trusted signature", because no web of trust was set up.
  - The same key also verified both package indexes: `gpg --verify packages.adb.asc packages.adb` in `packages/`
    and in `kmods/6.12.112-1-30028b278c7723afe2edd04be365895c/` (both "Good signature", exit 0).
- **usign signature: not checked.** There is no `usign` or `signify` on the PC. The signature's key number is
  `b5043e70f9a75cde` (bytes 3–10 of the decoded `sha256sums.sig`). The only usign key in the image is
  `/etc/opkg/keys/8a11255d14aef6c8` ("ASU CA pubkey 2022"), which is a different key. The matching public key was
  not available offline.
- **Release identity.** `version.buildinfo` says `r33311-981c978eff`. The image's `/etc/openwrt_release` gives the
  same revision with `DISTRIB_TARGET='armsr/armv7'` and `DISTRIB_ARCH='arm_cortex-a15_neon-vfpv4'`. The guest's
  `ubus call system board` agrees (section 12).
- **Decompressed images.** `gzip -t` on every `.gz` gave exit 0, and `gzip -dc` printed no "trailing garbage"
  warning. Each combined image is 243,564,032 bytes (475,711 sectors):
  - squashfs: `fb18c500f9e1990139fedeadc283226a0e7b278fe82f7a02aa4b7c641c7bb70b`
  - ext4: `f0d7ef9a698245b31e7aa2c741f02bcb9b9bfc663160203e0acd253673175fdd`
- **What the other files are** (compared with `sha256sum` and file lists):
  - `…-squashfs-rootfs.img.gz` and `…-ext4-rootfs.img.gz` decompress to exactly partition 2 of the matching combined
    image (109,051,904 bytes; same SHA-256 as the cut-out partition).
  - `…-generic-kernel.bin` is the `zImage` on the boot partition (same SHA-256).
  - `…-generic-targz-rootfs.tar.gz` has the same 914 files as the squashfs root (`diff` of the sorted file lists,
    exit 0).
  - `…-rootfs.tar.gz` and `…-rootfs.cpio.gz` are a smaller **base** root: the same 1,072 entries in both (`diff`
    exit 0) and 131 packages, which are exactly the 131 lines of the `.manifest` (`diff` exit 0). The 72 extra
    packages of the generic image are the device packages of `profiles.json` and their dependencies (section 10).
  - `…-initramfs-kernel.bin` carries the generic root as an uncompressed cpio inside the kernel (section 2).

## 2. Partition and file system layout, boot method

The layout was read with `sgdisk -p disk.img` (exit 0), `sgdisk -i N` and `fdisk -l` (exit 0). It is the same in
both images, including the disk GUID (`0F9DC41D-18C7-F917-0F85-DEF5FA7E4400`, fixed by the build).

| No. | Sectors (512 B) | Size | GPT type | Contents (`file`) |
| :--- | :--- | :--- | :--- | :--- |
| 128 | 34–511 | 239 KiB | BIOS boot (`EF02`) | all zero bytes (`tr -d '\0' \| wc -c` = 0); unused on ARM |
| 1 | 512–262655 | 128.0 MiB | **EFI System Partition** (`C12A7328-…`, name "EFI System Partition") | FAT16, label `kernel` |
| 2 | 262656–475647 | 104.0 MiB | Linux filesystem (`0FC63DAF-…`) | squashfs image: squashfs 4.0, xz, 3,836,034 bytes, 1,263 inodes. ext4 image: ext4, label `rootfs`, 26,624 blocks of 4 KiB, 21,515 free (`debugfs -R stats`) |

- **Partition table.** GPT behind a protective MBR. The backup GPT header is missing: the bytes at the alternate
  LBA are zero (`od`), sgdisk reports "invalid backup GPT header", and the guest kernel logs
  `Alternate GPT is invalid, using primary GPT.` and continues. Unlike x86-64 (I-01), partition 1 has the real
  ESP type here.
- **Boot partition.** Read with `mdir -/ -i p1.img ::` and `mcopy -s` (both exit 0). It holds
  `/efi/boot/bootarm.efi` and `/efi/openwrt/bootarm.efi` (both 471,040 bytes, identical; `file`: PE32 EFI
  application, ARM Thumb; GRUB 2.12, package `grub2-efi-arm 2.12-r1`), `/efi/openwrt/grub.cfg` (523 bytes) and
  `/efi/openwrt/zImage` (3,568,128 bytes). Partition 1 is byte-identical in the two images (`sha256sum`).
- **Boot method.** UEFI only. The tree's `u-boot-qemu_armv7/u-boot.bin` ("U-Boot 2025.04-OpenWrt-r33311-981c978eff",
  `strings`) acts as the UEFI firmware in QEMU. The boot log (section 12) shows U-Boot's `efi_mgr` boot method
  `Booting: virtio 0`, then GRUB 2.12, then `EFI stub: Booting Linux Kernel...` and
  `EFI stub: Using DTB from configuration table` (the device tree comes from QEMU through U-Boot).
  - `grub.cfg`: serial console 115200 8N1, timeout 5 s, two entries "OpenWrt" and "OpenWrt (failsafe)"
    (`failsafe=true`), both `linux /efi/openwrt/zImage root=PARTUUID=…-02 rootwait earlycon noinitrd`. There is no
    `console=` argument: the kernel takes its console from the device tree (`ttyAMA0`, PL011).
  - GRUB prints `error: serial port 'com0' isn't found.` and `error: terminal 'serial' isn't found.` and still
    shows its menu on the U-Boot console and boots the default entry after 5 s.
  - Guest `cat /proc/cmdline`: `BOOT_IMAGE=/efi/openwrt/zImage root=PARTUUID=0f9dc41d-18c7-f917-0f85-def5fa7e4402
    rootwait earlycon noinitrd`. No initramfs.
- **Kernel image format.** `file`: "Linux kernel ARM boot executable zImage (little-endian)"; `binwalk` finds a PE
  header (EFI stub) and xz data at offsets 54688 and 58548. The kernel proper is the xz stream at 58548
  (`tail -c +58549 | xz -dc --single-stream`, exit 0, 15,303,948 bytes).
- **Initramfs kernel.** `…-initramfs-kernel.bin` is the same kind of zImage. Its xz stream at 58548 decompresses to
  25,789,708 bytes and contains an uncompressed SVR4 cpio at offset 12,885,320 (`binwalk`). `cpio -it` lists 1,271
  entries; `cpio -id` gave the same 914 files as the squashfs root plus `/init` (276 bytes), and the same 203
  packages. `CONFIG_CMDLINE` is empty, so the command line comes from the boot loader or from QEMU's `-append`.
- **Root file system at run time (squashfs image)**, from the guest's `mount` and `df -k`:
  - the squashfs (`/dev/root`, 3,840 kB) is mounted read-only on `/rom`;
  - the rest of partition 2 is `/dev/loop0` (`loop0: detected capacity change from 212992 to 205440`), and on the
    first boot it is formatted **f2fs**, not ext4 as on x86-64: the log says
    `mount_root: overlay filesystem in /dev/loop0 has not been formatted yet` and then
    `mount_root: switching to f2fs overlay`;
  - `/overlay` is that f2fs, and `overlayfs:/overlay` is `/`;
  - unlike x86-64, the boot partition is **not** mounted at run time (no `/boot` in `mount`);
  - `/tmp` is tmpfs; `efivarfs` is mounted read-only on `/sys/firmware/efi/efivars`.
- **Root file system (ext4 image).** Taken apart, not booted. `debugfs -R "rdump / un/root" p2.img` (exit 0; three
  warnings about changing ownership, because there is no root). Its file list equals the squashfs root's
  (`find -type f | sort`, `diff` exit 0, 914 files), and so does its package list (`diff` exit 0). Features
  (`debugfs -R stats`): `has_journal ext_attr resize_inode filetype extent sparse_super large_file uninit_bg`.
- **Unpacking the squashfs.** `unsquashfs -no-xattrs -d … p2.img` gave exit 2; the one failure was
  `could not create character device …/dev/console, because you're not superuser`, expected without root.
  It created 914 files, 118 directories and 230 symlinks.

## 3. How an update is written

- **sysupgrade.** Present: `/sbin/sysupgrade`. The armsr rule is `/lib/upgrade/platform.sh`, read in the unpacked root:
  - `platform_check_image` reads the image's boot sector, compares its partition table with the disk, and if the
    layout changed writes the full image (`Partition layout has changed. Full image will be written.`).
  - `platform_do_upgrade` writes the whole disk only when the layout differs and `-p` was not given; otherwise it
    writes each partition separately. For partition 1 (the ESP) it does **not** overwrite the file system: it
    loop-mounts the new ESP and copies only `efi/openwrt` and `efi/boot` into the existing ESP, "to avoid writing
    over any firmware files (e.g ubootefi.var or firmware/X/ aka EBBR)" (comment in the file). Partition 2 is
    written with `dd`. Afterwards the `PARTUUID` in `grub.cfg` is rewritten to the real root partition.
  - `platform_copy_config` copies the saved settings onto partition 1 (vfat or ext4).
- **No A/B scheme.** One root partition, overwritten in place; settings carried across in the backup file.
- **Other update helpers.** `owut` and `attendedsysupgrade-common` are installed (package list), the same online
  image-rebuild helpers as on x86-64.
- **Packages at run time.** `apk` installs into the f2fs overlay.

## 4. Kernel and the configuration NMB needs

- **Kernel version: 6.12.112.** Sources:
  - the decompressed kernel (`grep -a -o 'Linux version [^%]*'`): `Linux version 6.12.112 (builder@buildhost)
    (arm-openwrt-linux-muslgnueabi-gcc (OpenWrt GCC 14.3.0 r33311-981c978eff) 14.3.0, GNU ld (GNU Binutils) 2.44)
    #0 SMP Sat Oct 17 21:45:38 2026`; the boot log shows the same line;
  - guest `uname -a`: `Linux OpenWrt 6.12.112 #0 SMP Sat Oct 17 21:45:38 2026 armv7l GNU/Linux`;
  - guest `ubus call system board`: `"kernel": "6.12.112"`;
  - package `kernel 6.12.112~30028b278c7723afe2edd04be365895c-r1` (the hash is the `vermagic` in `profiles.json`
    and the name of the `kmods/6.12.112-1-30028b278c7723afe2edd04be365895c/` folder).
- **Where the config came from: not in the image.** The decompressed kernel contains no `IKCFG_ST` marker
  (`grep -c -a IKCFG_ST` printed 0); the guest has no `/proc/config.gz`; the image carries no `.config`.
- **Source used: the `kmod-ikconfig` module of this exact build, from the tree** (`CONFIG_IKCONFIG=m`, so the module
  carries the configuration). `kmods/…/kmod-ikconfig-6.12.112-r1.apk`, SHA-256 checked against the signed list
  (section 1). The apk (ADB, `ADBd` = raw deflate) was inflated with the same 7-line Python script as in I-01, run
  with `python3 -I` from `tools/`, outside the data folder; the gzip blob between `IKCFG_ST` and `IKCFG_ED` was then
  decompressed (exit 0).
  - Result: 7,633 lines, starting `# Linux/arm 6.12.112 Kernel Configuration` and
    `CONFIG_CC_VERSION_TEXT="arm-openwrt-linux-muslgnueabi-gcc (OpenWrt GCC 14.3.0 r33311-981c978eff) 14.3.0"`.
    SHA-256 `9636d0c23549cb1b2060910187adbc9ada28e916cd9c9b42226ab8ce65db6fbc`.
  - **Cross-check with the SDK.** The SDK carries the kernel tree's `.config`
    (`build_dir/target-arm_cortex-a15+neon-vfpv4_musl_eabi/linux-armsr_armv7/linux-6.12.112/.config`, extracted
    alone, SHA-256 `b59abe68…1cf83`). The sorted non-comment lines of both differ only in the initramfs options
    (`INITRAMFS_SOURCE` and three `INITRAMFS_*` lines in the SDK copy, which was left by the initramfs build); every
    other option is the same.
  - **Source of every option below: this config.** "In image" says whether the module is in
    `lib/modules/6.12.112/` (or built in). "Package" says whether a kmod package is in the tree's `kmods/` folder
    (its `index.json`, 1,121 packages; every file checked against the signed list).

| Feature (5.2) | Option(s) | Value | In image | Package |
| :--- | :--- | :--- | :--- | :--- |
| Bridge, VLAN filtering | `BRIDGE`, `BRIDGE_VLAN_FILTERING` | y, y | built in (`modules.builtin`: `bridge.ko`) | — |
| Bonding | `BONDING` | m | no | `kmod-bonding 6.12.112-r1` |
| Team | `NET_TEAM` | m | no | `kmod-team 6.12.112-r1` |
| VLAN 802.1Q | `VLAN_8021Q` | y | built in (`8021q.ko`) | — |
| nftables | `NF_TABLES` m, `NF_TABLES_INET` y, `NF_TABLES_NETDEV` y, `NF_TABLES_BRIDGE` m, `NFT_NAT` m, `NFT_MASQ` m, `NFT_FIB_INET` m, `NF_FLOW_TABLE` m | as listed | yes (`nf_tables.ko` and 28 more `nft_*`/`nf_*` modules) | `kmod-nft-*` |
| WireGuard | `WIREGUARD` | m | no | `kmod-wireguard 6.12.112-r1` |
| PPP, PPPoE | `PPP` m, `PPPOE` m, `PPP_ASYNC` m | m | yes (`ppp_generic`, `pppoe`, `pppox`, `ppp_async`, `slhc`) | `kmod-ppp`, `kmod-pppoe` |
| cfg80211, mac80211 | `CFG80211` | **not set in the kernel** (`MAC80211` absent) | no | `kmod-cfg80211`, `kmod-mac80211` from backports, version `6.12.112.6.18.55-r2` |
| mac80211_hwsim | — | from backports | no | `kmod-mac80211-hwsim 6.12.112.6.18.55-r2` |
| Wi-Fi drivers (examples) | `WLAN` y | from backports / mt76 | no | `kmod-ath9k 6.12.112.6.18.55-r2`, `kmod-ath10k-ct`, `kmod-iwlwifi`, `kmod-mt7601u`, `kmod-mt76x2u`, `kmod-rt2800-usb` |
| USB host | `USB` y, `USB_XHCI_HCD` y, `USB_XHCI_PLATFORM` y, `USB_EHCI_HCD` m, `USB_DWC3` m | as listed | xHCI built in | `kmod-usb2`, `kmod-usb3` |
| USB network | `USB_USBNET` m, `USB_RTL8152` m, `USB_NET_AX88179_178A` m, `USB_NET_SMSC95XX` m, `USB_NET_AQC111` m | m | no | `kmod-usb-net-*` (e.g. `-rtl8152`, `-asix-ax88179`) |
| USB tethering (4.27) | `USB_NET_RNDIS_HOST`, `USB_NET_CDCETHER`, `USB_NET_CDC_NCM`, `USB_IPHETH`, `USB_NET_QMI_WWAN`, `USB_NET_CDC_MBIM`, `USB_NET_HUAWEI_CDC_NCM` | all m | no | `kmod-usb-net-rndis`, `-cdc-ether`, `-cdc-ncm`, `-ipheth`, `-qmi-wwan`, `-cdc-mbim` (all `6.12.112-r1`) |
| TUN, veth, macvlan, dummy, IFB | `TUN`, `VETH`, `MACVLAN`, `DUMMY`, `IFB` | all m | no | `kmod-tun`, `kmod-veth`, `kmod-macvlan`, `kmod-dummy` |
| Policy routing | `IP_ADVANCED_ROUTER`, `IP_MULTIPLE_TABLES`, `IPV6_MULTIPLE_TABLES` | y | built in | — |
| Traffic control | `NET_SCHED` y, `NET_SCH_FQ_CODEL` y, `NET_SCH_CAKE` m, `NET_SCH_HTB` m, `NET_SCH_NETEM` m, `NET_SCH_TBF` m | as listed | fq_codel built in (guest `ip addr`: `qdisc fq_codel`) | `kmod-sched-cake`, `kmod-netem` |
| Namespaces | `NAMESPACES`, `NET_NS`, `PID_NS`, `USER_NS`, `UTS_NS`, `IPC_NS` | all y | built in | — |
| seccomp | `SECCOMP`, `SECCOMP_FILTER` | y, y | built in | (`procd-seccomp` installed) |
| **Landlock** | `SECURITY` | **not set**, so `SECURITY_LANDLOCK` is absent | — | not available without an own kernel |
| cgroups | `CGROUPS` y, `MEMCG` y (`MEMCG_V1` not set), `CGROUP_PIDS` y, `CPUSETS` y, `CGROUP_BPF` y, `BLK_CGROUP` y; `CGROUP_FREEZER`, `CGROUP_DEVICE` not set | as listed | cgroup2 on `/sys/fs/cgroup` (guest `mount`, `nsdelegate`) | — |
| Watchdog | `WATCHDOG` y; `ARM_SP805_WATCHDOG` m, `SOFT_WATCHDOG` m, `I6300ESB_WDT` m | as listed | `sp805_wdt.ko` in the image, but no `/dev/watchdog*` in the QEMU guest (no watchdog device was given to QEMU); `ubus call system watchdog`: `"status": "offline", "timeout": 30` | `kmod-softdog`, `kmod-i6300esb-wdt` |
| Module signing | `MODULE_SIG` | not set | — | — |
| Hardening (extra) | `STACKPROTECTOR` y (`_STRONG` not set), `FORTIFY_SOURCE` y, `HARDENED_USERCOPY` y, `STRICT_DEVMEM` not set | as listed | — | — |
| 32-bit ARM specifics | `ARM_LPAE` y, `VMSPLIT_3G` y (`PAGE_OFFSET=0xC0000000`), `HIGHMEM` not set, `ARCH_VIRT` y, `ARCH_MULTI_V7` y, `SMP` y, `NR_CPUS=4`, `VFP`/`NEON`/`KERNEL_MODE_NEON` y, `EFI` y, `EFI_STUB` y | as listed | the guest CPU flags include `lpae` and `neon` (`/proc/cpuinfo`) | — |
| Other | `BPF_SYSCALL` y, `XDP_SOCKETS` not set, `TCP_CONG_BBR` m, `IPV6` y, `PREEMPT_NONE` y, `HZ=100`, `DEVTMPFS` not set; virtio net/blk/pci/mmio/console built in; `SQUASHFS`, `EXT4_FS`, `F2FS_FS`, `OVERLAY_FS` y; `VFAT_FS` m | as listed | — | — |

- **What the 32-bit options mean for NMB, as far as read here.** `HIGHMEM` is off, so the kernel uses at most the
  low memory below `PAGE_OFFSET`; with 512 MB the guest saw `Memory: 503392K/524288K available` (all of it). Which
  CPUs an `ARM_LPAE=y` kernel excludes was not checked here (the Kconfig help is not in the SDK); QEMU's
  `cortex-a15` worked.

## 5. Modules and firmware files carried

- **Modules.** 88 `.ko` files in `lib/modules/6.12.112/` (`ls *.ko | wc -l`), plus `modules.builtin` (136 lines;
  virtio, PL011, bridge, 8021q, squashfs, ext4, overlay and 9p-virtio are built in). Guest `lsmod | wc -l` = 88.
  - Groups: NIC drivers (`e1000e`, `ena`, `atlantic`, `genet`, `stmmac`/`stmmac-platform`/`dwmac-generic`,
    `vmxnet3`), PHY and MDIO (`aquantia`, `bcm7xxx`, `broadcom`, `marvell`, `marvell10g`, `realtek`, `smsc`, `sfp`,
    `phylink`, `pcs_xpcs`, `mdio-*`), nftables and conntrack, PPP/PPPoE, `macsec`, I2C/GPIO/RTC
    (`i2c-mux-pca954x`, `gpio-pca953x`, `rtc-rx8025`), `sp805_wdt`, crypto helpers, `vfat`/NLS.
  - 37 installed kmod packages carry no `.ko` file on this target (`R:` entries of `lib/apk/db/installed`), among
    them the device packages `kmod-mvneta`, `kmod-mvpp2`, `kmod-fsl-dpaa1-net`, `kmod-fsl-dpaa2-net`,
    `kmod-fsl-enetc-net`, `kmod-fsl-fec`, `kmod-dwmac-imx`, `kmod-dwmac-rockchip`, `kmod-dwmac-sun8i`,
    `kmod-octeontx2-net` and `kmod-renesas-net-avb`. Whether their drivers are built in or simply not built for
    armv7 was not determined.
- **Firmware.** `lib/firmware` is **empty** (`find -type f | wc -l` = 0). No NIC, Wi-Fi or other firmware is in the
  image, and no firmware package is installed.

## 6. Init system, service manager, C library, package manager

- **Init:** procd (`/sbin/procd` is PID 1 in the guest's `ps w`) with ubus, ujail and seccomp (`procd-ujail`,
  `procd-seccomp`). Init scripts in `/etc/init.d`, linked in `/etc/rc.d`.
- **Enabled services** (guest loop `$s enabled`): boot, cron, dnsmasq, done, dropbear, firewall, gpio_switch, led, log,
  network, odhcpd, packet_steering, rpcd, sysctl, sysfixtime, sysntpd, system, ucitrack, uhttpd, umount,
  urandom_seed, urngd (22 of 22; the same set as x86-64).
- **Services known to procd** (guest `ubus call service list`): cron, dnsmasq, dropbear, firewall, gpio_switch, log,
  network, odhcpd, packet_steering, rpcd, sysntpd, system, ubus, ucitrack, uhttpd, urandom_seed, urngd.
- **C library: musl 1.2.5.** The guest's `/lib/libc.so` prints `musl libc (armhf)` / `Version 1.2.5`. Package
  `libc 1.2.5-r5`; the loader is `/lib/ld-musl-armhf.so.1` (hard-float EABI). Toolchain GCC 14.3.0 (kernel version
  string; section 15).
- **Package manager: apk-tools 3.0.5** (guest `apk --version`: `apk-tools 3.0.5, compiled for armv7.`). opkg is
  absent (guest `command -v opkg`). Package architecture `arm_cortex-a15_neon-vfpv4` (`/etc/apk/arch`,
  `profiles.json`). Feeds in `/etc/apk/repositories.d/distfeeds.list`: target, kmods, base, luci, packages,
  routing, telephony and video for 25.12.6.
- **Installed packages: 203** in the generic image (`awk '/^P:/…' lib/apk/db/installed`; the squashfs and ext4
  roots, the targz root and the initramfs give the same list). The base root (`rootfs.tar.gz`, `.cpio.gz`,
  `.manifest`) has 131. The list is in section 10.
- **Licences of the installed packages** (`L:` fields, `sort | uniq -c`): GPL-2.0-only 102, ISC 29, GPL-2.0 26,
  Apache-2.0 15 (LuCI), LGPL-2.1 7, MIT 4, GPL-3.0-with-GCC-exception 4 (libc, libgcc and friends),
  GPL-2.0-or-later 4, LGPL-2.1-or-later 3, BSD-4-Clause 2 (`ppp`, `ppp-mod-pppoe`), and one each of Zlib, LGPL-2.1+,
  GPL-3.0-or-later (`grub2-efi-arm`), GPL-2.0-or-later MPL-2.0 (`ca-bundle`), GPL-2.0+, GPL-2.0 BSD-3-Clause and
  BSD-3-Clause. Every package has a licence field (no firmware packages, unlike x86-64).

## 7. Tools of 6A criterion 3, with versions

Presence: the guest's `command -v` loop (the same result in the unpacked root with `find`). Versions: the guest's
own output where it exists, otherwise the installed package version.

| Tool | Present | Version |
| :--- | :--- | :--- |
| procd (init) | yes, `/sbin/procd` | `procd 2026.03.13~58eb263d-r1` (package) |
| systemd, OpenRC | no (`systemctl`, `rc-service` absent) | — |
| nftables `nft` | yes, `/usr/sbin/nft` | `nftables v1.1.6 (Commodore Bullmoose #7)`; package `nftables-json 1.1.6-r2` |
| dnsmasq | yes | `Dnsmasq version 2.93`; package `dnsmasq 2.93-r1` (not `dnsmasq-full`) |
| hostapd | **no** | (feed: section 9a) |
| wpa_supplicant | **no** | (feed: section 9a) |
| iw | **no** | (feed: section 9a) |
| pppd | yes | `pppd version 2.5.2`; `ppp 2.5.2-r4`, `ppp-mod-pppoe 2.5.2-r4` |
| WireGuard (`wg`, module) | **no** | (feed: section 9a) |
| OpenVPN | **no** | (feed: section 9a) |
| `ip` | yes, but the **BusyBox applet** (`/sbin/ip -> ../bin/busybox`) | BusyBox v1.37.0; `ip link add … type bridge` and `ip addr add` worked (section 9) |
| `tc`, `ethtool`, `iptables`, `ss` | no | — |
| DHCP client | `udhcpc` (BusyBox), `odhcp6c 2026.06.27~10a52220-r3` | — |
| DHCP/RA server | `odhcpd-ipv6only 2026.06.29~5d7be43f-r1` and dnsmasq | — |
| SSH | `dropbear` (with `dbclient`, `scp`, `ssh` links in `/usr/bin`) | `Dropbear v2025.89` |
| Web | `uhttpd 2026.08.24~373145f7-r1` with LuCI `26.288.56989~75a306d` | — |
| HTTP client | `uclient-fetch` (`wget` is a link to it) | `uclient-fetch 2026.06.07~daad21fa-r1` |
| BusyBox | yes | `busybox 1.37.0-r6` (banner: `BusyBox v1.37.0 (2026-10-17 21:45:38 UTC)`). Applets include `nc`, `sha256sum`, `flock`, `tar`, `hexdump`; no `base64`, `tftp` or `timeout` |

## 8. Python (6A criterion 8)

**No Python is in the image.** The checks:

- In the unpacked squashfs root, `find . -iname "python*" -o -name "*.py" -o -name "*.pyc"` found 0 files;
  `grep -rlI "^#!.*python" .` found 0, and `grep -rlI "^#!.*perl" .` found 0.
- In the guest, `find / -xdev …` and `find /rom -xdev …` both printed 0; `command -v python python3 perl` printed
  absent for all three.
- `python3` and `python3-base 3.13.18-r1` exist in the packages feed but are not installed (section 9a).
- The ImageBuilder and SDK contain Python scripts for the **build host**: `meson.py` and `b43-fwsquash.py` in the
  ImageBuilder's `staging_dir/host/bin`, and 126 `.py` files in the SDK (its `scripts/` folder and the kernel tree;
  `grep '\.py$'` on the archive listings). They are build tools and are not in the image.

## 9. Own network and firewall manager, and how each is switched off (6A criterion 5)

The guest runs **netifd** (`/sbin/netifd`, package `netifd 2026.02.26~cbb83a18-r2`) and **firewall4**
(`/sbin/fw4`, ucode, package `firewall4 2025.03.17~b6e51575-r2`). Ruleset: one table `table inet fw4`
(guest `nft list tables`), 145 lines (`nft list ruleset | wc -l`), the same as x86-64.

**Default configuration** (guest `uci show network`, `uci show firewall`; `/etc/board.d/99-default_network`):

- `eth0` is bridged into `br-lan` with the static built-in LAN default `192.168.1.1/24` and `ip6assign 60`.
- `eth1` (only if it exists) is `wan` with DHCP and `wan6` with DHCPv6. In the guest, `eth1` got `192.0.2.15/24`
  and the default route from QEMU's user-mode network.
- A ULA prefix and a DHCP DUID are generated on the first boot (values are per guest and not recorded).
- Firewall zones: `lan` input/output/forward ACCEPT; `wan` input REJECT, forward DROP, masquerading on; one
  forwarding `lan` → `wan`; defaults input and forward REJECT.
- Board scripts know one real armsr board by name (`traverse,ten64`); every other board, QEMU's `linux,dummy-virt`
  included, gets the generic `eth0`/`eth1` layout.

Coexistence tests in the guest (run 1; every probe line ended with exit 0):

| Test | Command | Seen |
| :--- | :--- | :--- |
| Own nft table across firewall restart and reload | `nft add table inet nmbtest` + chain; `/etc/init.d/firewall restart`; `nft list tables`; `fw4 reload`; `nft list tables` | `inet nmbtest` **kept** after both |
| Plain `fw4 stop` / `fw4 start` | `fw4 stop`; `nft list tables`; `fw4 start` | after stop only `inet nmbtest`; after start `inet nmbtest` and `inet fw4`: `fw4 stop` removes only its own table |
| Firewall service stop | `/etc/init.d/firewall stop`; `nft list tables`; `/etc/init.d/firewall start` | **empty after stop, NMB's table included**; after start only `inet fw4`. The same as x86-64: the init script's stop flushes the whole ruleset |
| Firewall disable | `/etc/init.d/firewall disable`; `… enabled; echo $?`; `ls /etc/rc.d \| grep -c firewall` | `enabled` gave 1 and 0 links were left, so it is off at boot; `enable` brought back 1 link |
| Own bridge and foreign address | `ip link add nmbbr0 type bridge`, `198.51.100.77/24` on it; `203.0.113.9/24` added on netifd's `br-lan` | both present |
| `network reload` + 10 s | `ip addr show …` | `nmbbr0` address present; `203.0.113.9` on `br-lan` **present** |
| `network restart` + 15 s | `ip addr show …` | `nmbbr0` address present; `203.0.113.9` on `br-lan` **gone** (netifd rebuilds `br-lan`) |
| netifd stop / start | `/etc/init.d/network stop`; `pgrep -l netifd`; `ip route`; `start`; 15 s; `pgrep`, `ip route` | after stop no netifd (pgrep rc 1) and only the `nmbbr0` route left; after start netifd ran again and `br-lan` and `nmbbr0` routes were present (the WAN DHCP lease was not yet back after 15 s) |
| netifd disable | `/etc/init.d/network disable`; `ls /etc/rc.d \| grep -c network` | 0 links, so off at boot; `enable` brought back 2 links (`S20network`, `K90network`) |

**What this means for criterion 5.** The same as on x86-64 (I-01): both managers can be switched off cleanly with
`disable` (plus `stop` for netifd); netifd leaves devices it does not manage alone and keeps foreign addresses on its
own devices across `reload` but not `restart`; `fw4` leaves foreign tables alone, but `/etc/init.d/firewall stop`
flushes the whole ruleset, so NMB must use `fw4 stop` or `disable` and a reboot while its own tables are loaded.

### 9a. Feed lookup (run 1)

`apk update` in the guest over the user-mode WAN gave rc 0 and `OK: 11132 distinct packages available`. Then, for
each name, `apk search $p | grep "^$p-[0-9]"`. **None of these packages is installed in the image.** The kmod
versions agree with the tree's own `kmods/…/index.json` (read on the PC with `python3 -I` and `json`).

| Package | Version in the 25.12.6 feeds for `arm_cortex-a15_neon-vfpv4` |
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
| `kmod-veth`, `kmod-bonding`, `kmod-usb-net-rndis` | `6.12.112-r1` |
| `python3` / `python3-base` | `3.13.18-r1` (in the feed only; never installed for NMB, criterion 8) |

These are the same versions as in the x86-64 feeds (I-01 section 9a).

## 10. Installed packages (203, `lib/apk/db/installed`; the same in every generic root)

`apk-mbedtls 3.0.5-r3`, `attendedsysupgrade-common 10`, `base-files 1718~981c978eff`, `blkid 2.41.5-r1`, `busybox 1.37.0-r6`, `ca-bundle 20260816-r1`, `cgi-io 2026.07.21~31cb3c89-r1`, `dnsmasq 2.93-r1`, `dropbear 2025.89-r3`, `e2fsprogs 1.47.3-r1`, `firewall4 2025.03.17~b6e51575-r2`, `fstools 2026.05.23~16718b6e-r1`, `fwtool 2025.10.03~04cd252e-r1`, `getrandom 2025.10.30~6f78fa49-r1`, `grub2-efi-arm 2.12-r1`, `jansson4 2.14.1-r1`, `jshn 2026.06.19~7dd12784-r1`, `jsonfilter 2026.03.16~b9034210-r1`, `kernel 6.12.112~30028b278c7723afe2edd04be365895c-r1`, `kmod-acpi-mdio 6.12.112-r1`, `kmod-amazon-ena 6.12.112-r1`, `kmod-atlantic 6.12.112-r1`, `kmod-bcmgenet 6.12.112-r1`, `kmod-crypto-aead 6.12.112-r1`, `kmod-crypto-crc32 6.12.112-r1`, `kmod-crypto-crc32c 6.12.112-r1`, `kmod-crypto-ctr 6.12.112-r1`, `kmod-crypto-gcm 6.12.112-r1`, `kmod-crypto-geniv 6.12.112-r1`, `kmod-crypto-gf128 6.12.112-r1`, `kmod-crypto-ghash 6.12.112-r1`, `kmod-crypto-hash 6.12.112-r1`, `kmod-crypto-hmac 6.12.112-r1`, `kmod-crypto-manager 6.12.112-r1`, `kmod-crypto-null 6.12.112-r1`, `kmod-crypto-rng 6.12.112-r1`, `kmod-crypto-seqiv 6.12.112-r1`, `kmod-crypto-sha3 6.12.112-r1`, `kmod-crypto-sha512 6.12.112-r1`, `kmod-dwmac-imx 6.12.112-r1`, `kmod-dwmac-rockchip 6.12.112-r1`, `kmod-dwmac-sun8i 6.12.112-r1`, `kmod-e1000e 6.12.112-r1`, `kmod-fixed-phy 6.12.112-r1`, `kmod-fs-vfat 6.12.112-r1`, `kmod-fsl-dpaa1-net 6.12.112-r1`, `kmod-fsl-dpaa2-net 6.12.112-r1`, `kmod-fsl-enetc-net 6.12.112-r1`, `kmod-fsl-fec 6.12.112-r1`, `kmod-fsl-mc-dpio 6.12.112-r1`, `kmod-fsl-pcs-lynx 6.12.112-r1`, `kmod-fsl-xgmac-mdio 6.12.112-r1`, `kmod-gpio-pca953x 6.12.112-r1`, `kmod-hwmon-core 6.12.112-r1`, `kmod-i2c-core 6.12.112-r1`, `kmod-i2c-mux 6.12.112-r1`, `kmod-i2c-mux-pca954x 6.12.112-r1`, `kmod-lib-crc-ccitt 6.12.112-r1`, `kmod-lib-crc-itu-t 6.12.112-r1`, `kmod-lib-crc16 6.12.112-r1`, `kmod-lib-crc32c 6.12.112-r1`, `kmod-libphy 6.12.112-r1`, `kmod-macsec 6.12.112-r1`, `kmod-marvell-mdio 6.12.112-r1`, `kmod-mdio-bcm-unimac 6.12.112-r1`, `kmod-mdio-bus-mux 6.12.112-r1`, `kmod-mdio-devres 6.12.112-r1`, `kmod-mdio-gpio 6.12.112-r1`, `kmod-mii 6.12.112-r1`, `kmod-mvneta 6.12.112-r1`, `kmod-mvpp2 6.12.112-r1`, `kmod-net-selftests 6.12.112-r1`, `kmod-nf-conntrack 6.12.112-r1`, `kmod-nf-conntrack6 6.12.112-r1`, `kmod-nf-flow 6.12.112-r1`, `kmod-nf-log 6.12.112-r1`, `kmod-nf-log6 6.12.112-r1`, `kmod-nf-nat 6.12.112-r1`, `kmod-nf-reject 6.12.112-r1`, `kmod-nf-reject6 6.12.112-r1`, `kmod-nfnetlink 6.12.112-r1`, `kmod-nft-core 6.12.112-r1`, `kmod-nft-fib 6.12.112-r1`, `kmod-nft-nat 6.12.112-r1`, `kmod-nft-offload 6.12.112-r1`, `kmod-nls-base 6.12.112-r1`, `kmod-nls-cp437 6.12.112-r1`, `kmod-nls-iso8859-1 6.12.112-r1`, `kmod-nls-utf8 6.12.112-r1`, `kmod-octeontx2-net 6.12.112-r1`, `kmod-of-mdio 6.12.112-r1`, `kmod-pcs-xpcs 6.12.112-r1`, `kmod-phy-aquantia 6.12.112-r1`, `kmod-phy-bcm7xxx 6.12.112-r1`, `kmod-phy-broadcom 6.12.112-r1`, `kmod-phy-marvell 6.12.112-r1`, `kmod-phy-marvell-10g 6.12.112-r1`, `kmod-phy-realtek 6.12.112-r1`, `kmod-phy-smsc 6.12.112-r1`, `kmod-phylib-broadcom 6.12.112-r1`, `kmod-phylink 6.12.112-r1`, `kmod-ppp 6.12.112-r1`, `kmod-pppoe 6.12.112-r1`, `kmod-pppox 6.12.112-r1`, `kmod-pps 6.12.112-r1`, `kmod-ptp 6.12.112-r1`, `kmod-regmap-core 6.12.112-r1`, `kmod-regmap-i2c 6.12.112-r1`, `kmod-renesas-net-avb 6.12.112-r1`, `kmod-rtc-rx8025 6.12.112-r1`, `kmod-sfp 6.12.112-r1`, `kmod-slhc 6.12.112-r1`, `kmod-stmmac-core 6.12.112-r1`, `kmod-vmxnet3 6.12.112-r1`, `kmod-wdt-sp805 6.12.112-r1`, `libblkid1 2.41.5-r1`, `libblobmsg-json20260213 2026.06.19~7dd12784-r1`, `libc 1.2.5-r5`, `libcomerr0 1.47.3-r1`, `libe2p2 1.47.3-r1`, `libext2fs2 1.47.3-r1`, `libf2fs6 1.16.0-r4`, `libgcc1 14.3.0-r5`, `libiwinfo-data 2026.05.26~66bdd1a0-r1`, `libiwinfo20230701 2026.05.26~66bdd1a0-r1`, `libjson-c5 0.18-r1`, `libjson-script20260213 2026.06.19~7dd12784-r1`, `liblucihttp-ucode 2023.03.15~9b5b683f-r1`, `liblucihttp0 2023.03.15~9b5b683f-r1`, `libmbedtls21 3.6.7-r1`, `libmnl0 1.0.5-r1`, `libnftnl11 1.3.1-r1`, `libnl-tiny1 2025.12.02~40493a65-r1`, `libpthread 1.2.5-r5`, `librt 1.2.5-r5`, `libsmartcols1 2.41.5-r1`, `libss2 1.47.3-r1`, `libubox20260213 2026.06.19~7dd12784-r1`, `libubus20251202 2026.06.28~24864e78-r2`, `libuci20250120 2025.12.02~66127cd7-r1`, `libuclient20201210 2026.06.07~daad21fa-r1`, `libucode20230711 2026.01.16~85922056-r2`, `libudebug 2026.01.16~875e1a7a`, `libustream-mbedtls20201210 2026.03.01~99f1c0db-r1`, `libuuid1 2.41.5-r1`, `logd 2025.10.30~6f78fa49-r1`, `luci 26.288.56989~75a306d`, `luci-app-attendedsysupgrade 26.288.56989~75a306d`, `luci-app-firewall 26.288.56989~75a306d`, `luci-app-package-manager 26.288.56989~75a306d`, `luci-base 26.288.56989~75a306d`, `luci-lib-uqr 26.288.56989~75a306d`, `luci-light 26.288.56989~75a306d`, `luci-mod-admin-full 26.288.56989~75a306d`, `luci-mod-network 26.288.56989~75a306d`, `luci-mod-status 26.288.56989~75a306d`, `luci-mod-system 26.288.56989~75a306d`, `luci-proto-ipv6 26.288.56989~75a306d`, `luci-proto-ppp 26.288.56989~75a306d`, `luci-ssl 26.288.56989~75a306d`, `luci-theme-bootstrap 26.288.56989~75a306d`, `mkf2fs 1.16.0-r4`, `mtd 27`, `netifd 2026.02.26~cbb83a18-r2`, `nftables-json 1.1.6-r2`, `odhcp6c 2026.06.27~10a52220-r3`, `odhcpd-ipv6only 2026.06.29~5d7be43f-r1`, `openwrt-keyring 2025.12.10~f0670054-r1`, `owut 2026.08.30~9b51eeec-r1`, `partx-utils 2.41.5-r1`, `ppp 2.5.2-r4`, `ppp-mod-pppoe 2.5.2-r4`, `procd 2026.03.13~58eb263d-r1`, `procd-seccomp 2026.03.13~58eb263d-r1`, `procd-ujail 2026.03.13~58eb263d-r1`, `px5g-mbedtls 11`, `rpcd 2026.10.04~d99f703e-r1`, `rpcd-mod-file 2026.10.04~d99f703e-r1`, `rpcd-mod-iwinfo 2026.10.04~d99f703e-r1`, `rpcd-mod-luci 20240305-r1`, `rpcd-mod-rpcsys 2026.10.04~d99f703e-r1`, `rpcd-mod-rrdns 20170710`, `rpcd-mod-ucode 2026.10.04~d99f703e-r1`, `ubox 2025.10.30~6f78fa49-r1`, `ubus 2026.06.28~24864e78-r2`, `ubusd 2026.06.28~24864e78-r2`, `uci 2025.12.02~66127cd7-r1`, `uclient-fetch 2026.06.07~daad21fa-r1`, `ucode 2026.01.16~85922056-r2`, `ucode-mod-fs 2026.01.16~85922056-r2`, `ucode-mod-html 1`, `ucode-mod-log 2026.01.16~85922056-r2`, `ucode-mod-math 2026.01.16~85922056-r2`, `ucode-mod-ubus 2026.01.16~85922056-r2`, `ucode-mod-uci 2026.01.16~85922056-r2`, `ucode-mod-uclient 2026.06.07~daad21fa-r1`, `ucode-mod-uloop 2026.01.16~85922056-r2`, `uhttpd 2026.08.24~373145f7-r1`, `uhttpd-mod-ubus 2026.08.24~373145f7-r1`, `urandom-seed 3`, `urngd 2025.10.03~f17e33d9-r1`, `usign 2025.10.03~c4c72b1b-r1`, `zlib 1.3.1-r1`

The 131 packages of the base root (`rootfs.tar.gz`, `.manifest`) are this list without these 72: `kmod-acpi-mdio`,
`kmod-amazon-ena`, `kmod-atlantic`, `kmod-bcmgenet`, `kmod-crypto-aead`, `kmod-crypto-crc32`, `kmod-crypto-ctr`,
`kmod-crypto-gcm`, `kmod-crypto-geniv`, `kmod-crypto-gf128`, `kmod-crypto-ghash`, `kmod-crypto-hmac`,
`kmod-crypto-manager`, `kmod-crypto-null`, `kmod-crypto-rng`, `kmod-crypto-seqiv`, `kmod-crypto-sha3`,
`kmod-crypto-sha512`, `kmod-dwmac-imx`, `kmod-dwmac-rockchip`, `kmod-dwmac-sun8i`, `kmod-e1000e`, `kmod-fixed-phy`,
`kmod-fsl-dpaa1-net`, `kmod-fsl-dpaa2-net`, `kmod-fsl-enetc-net`, `kmod-fsl-fec`, `kmod-fsl-mc-dpio`,
`kmod-fsl-pcs-lynx`, `kmod-fsl-xgmac-mdio`, `kmod-gpio-pca953x`, `kmod-hwmon-core`, `kmod-i2c-core`, `kmod-i2c-mux`,
`kmod-i2c-mux-pca954x`, `kmod-lib-crc-itu-t`, `kmod-lib-crc16`, `kmod-libphy`, `kmod-macsec`, `kmod-marvell-mdio`,
`kmod-mdio-bcm-unimac`, `kmod-mdio-bus-mux`, `kmod-mdio-devres`, `kmod-mdio-gpio`, `kmod-mii`, `kmod-mvneta`,
`kmod-mvpp2`, `kmod-net-selftests`, `kmod-octeontx2-net`, `kmod-of-mdio`, `kmod-pcs-xpcs`, `kmod-phy-aquantia`,
`kmod-phy-bcm7xxx`, `kmod-phy-broadcom`, `kmod-phy-marvell`, `kmod-phy-marvell-10g`, `kmod-phy-realtek`,
`kmod-phy-smsc`, `kmod-phylib-broadcom`, `kmod-phylink`, `kmod-pps`, `kmod-ptp`, `kmod-regmap-core`,
`kmod-regmap-i2c`, `kmod-renesas-net-avb`, `kmod-rtc-rx8025`, `kmod-sfp`, `kmod-stmmac-core`, `kmod-vmxnet3`,
`kmod-wdt-sp805`, `libsmartcols1`, `partx-utils` (`comm -13` of the two sorted lists).

## 11. Size on disk

- Download: 8,006,271 bytes (squashfs) and 9,135,764 bytes (ext4), as `.img.gz`.
- Disk image: 243,564,032 bytes (232.3 MiB) for each; ESP 128 MiB, root partition 104 MiB
  (`CONFIG_TARGET_KERNEL_PARTSIZE=128`, `CONFIG_TARGET_ROOTFS_PARTSIZE=104` in the ImageBuilder's `.config`).
- squashfs root: 3,836,034 bytes compressed. Unpacked, the root is 8,001,828 bytes (`du -sb`), the ext4 root the
  same files. Of that:
  - `lib/modules`: 2,315,672 bytes
  - `lib/firmware`: 0
  - `www` (LuCI): 1,337,317 bytes
  - `usr`: 5,464,822 bytes
- The base root (`rootfs.tar.gz`, unpacked with `--exclude=./dev/*`) is 9,863,114 bytes by `du -sb`.
- Guest `df -k` after `init complete` (squashfs image, run 1):
  - `/rom`: 3,840 kB, 100 % used
  - `/overlay` (f2fs on loop0): 100,672 kB, 39,344 kB used, 61,328 kB free. The 39 MB "used" right after
    formatting are f2fs's own reserved space, not files: at the end of the run it was 39,384 kB.
  - `/tmp`: tmpfs, 253,332 kB
  - The ESP is not mounted; `mdir` shows 129,417,216 bytes free on it.
- ext4 image: the 104 MiB root has 21,515 of 26,624 4-KiB blocks free (about 84 MiB), from `debugfs -R stats`.

## 12. QEMU boot test (squashfs combined-efi image; this is the 32-bit ARM test guest of F-06)

**How it was run.** Each boot ran on a fresh copy of the decompressed image (and of `u-boot.bin`) in its own
`boot<N>/` folder, under `flock ~/nmb-work/qemu.lock` for the whole run, one guest at a time. The serial console
was driven by a small Python script (`python3 -I tools/drive.py`, kept outside the data folders) through a Unix
socket; it sends each probe line in 32-byte pieces and waits for an end marker with the line's exit code.

**The working QEMU command line** (run 1; `<w>` is the boot folder):

```
flock ~/nmb-work/qemu.lock \
qemu-system-arm -M virt,highmem=off -cpu cortex-a15 -accel tcg,thread=multi,tb-size=128 -smp 2 -m 512 \
  -nographic -monitor none -serial unix:<w>/serial.sock,server=on,wait=on \
  -bios <w>/u-boot.bin \
  -drive file=<w>/disk.img,format=raw,if=none,id=d0 -device virtio-blk-pci,drive=d0 \
  -drive file=<w>/data.img,format=raw,if=none,id=d1,readonly=on -device virtio-blk-pci,drive=d1 \
  -netdev user,id=lan0,net=198.51.100.0/24,restrict=on -device virtio-net-pci,netdev=lan0,mac=02:00:00:00:03:01 \
  -netdev user,id=wan0,net=192.0.2.0/24,guestfwd=tcp:192.0.2.100:8000-cmd:cat\ <file> \
  -device virtio-net-pci,netdev=wan0,mac=02:00:00:00:03:02
```

- QEMU 8.2.2 (`qemu-system-arm 1:8.2.2+ds-0ubuntu1.18`); the firmware is the tree's own `u-boot.bin`
  (U-Boot 2025.04), no host EFI package needed. TCG with two vCPUs (`thread=multi`); KVM cannot run an ARM guest
  on this x86 host.
- `-cpu cortex-a15` provides LPAE (guest `/proc/cpuinfo` flag `lpae`), which the kernel is built for (`ARM_LPAE=y`).
  `highmem=off` was set in every run; a run without it was not tried.
- U-Boot printed `DRAM: 512 MiB`, found both NICs (`eth0: virtio-net#34, eth1: virtio-net#35`), warned
  `Loading Environment from Flash... *** Warning - bad CRC, using default environment` (empty flash, harmless) and
  booted through `efi_mgr` → `Booting: virtio 0` → GRUB → the kernel.
- In the guest the first NIC is `eth0` (in `br-lan`), the second `eth1` (`wan`). The first disk is `vda` (root), the
  data disk `vdb`.

**How a file is put into the guest.** Three ways worked (1 and 2 in run 1, 3 in run 3), all offline and without
root on the PC:

1. **A read-only data disk.** On the PC: `dd if=/dev/zero of=data.img bs=1M count=8`, `mformat -i data.img -v
   NMBDATA ::`, `mcopy -i data.img <file> ::/`. It is the second `-drive` above. In the guest:
   `modprobe vfat; mkdir -p /mnt/data; mount -t vfat -o ro /dev/vdb /mnt/data`. The test file (8,922 bytes, SHA-256
   `374ad8f98a2ca42be7ca5a0daea8638070e74be4b4cbe9f5b7ebd43bda4a9de6`) was read back with the same SHA-256 by the
   guest's `sha256sum` (exit 0). The kernel logs `vdb: p1 size 17136 extends beyond EOD, truncated`, because it
   reads mformat's boot sector as a partition table; mounting the whole `/dev/vdb` worked regardless.
2. **A QEMU `guestfwd` on the WAN network.** `guestfwd=tcp:192.0.2.100:8000-cmd:cat <file>` makes QEMU run
   `cat <file>` for each connection the guest opens to 192.0.2.100:8000. In the guest:
   `nc 192.0.2.100 8000 > /tmp/<file> & sleep 15; kill $!` (BusyBox `nc`). The file arrived complete (8,922 bytes,
   same SHA-256). This needs no disk and no network service on the PC; it also works on a `restrict=on` network.

3. **A 9p shared folder (run 3; the simplest for F-06).** The kernel has `NET_9P=y`, `NET_9P_VIRTIO=y` and
   `9P_FS=y` (`modules.builtin`: `9p.ko`, `9pnet.ko`, `9pnet_virtio.ko`), so no module is needed. QEMU option:
   `-virtfs local,path=<host folder>,mount_tag=nmbxfer,security_model=none,readonly=on`. In the guest:
   `mkdir -p /mnt/x; mount -t 9p -o trans=virtio,version=9p2000.L,ro nmbxfer /mnt/x`. `ls -l /mnt/x` showed the
   folder's two files, and the test file had the same SHA-256 on the share and after `cp` to `/tmp` (exit 0).
   With `security_model=none` the files appear with the PC user's numeric owner (1000). A writable share was not
   tried.

Not used: SSH into the guest (the WAN zone rejects input, and the LAN network is restricted).

| Run | Settings | Result |
| :--- | :--- | :--- |
| 1 | the command line above, squashfs image, U-Boot | Booted; all 30 probe lines exit 0; `poweroff` clean (`reboot: Power down`), driver exit 0, QEMU exit 0, `flock`+run exit 0. 06:36:36–06:51:02 UTC; host load not recorded (an arm64 guest of another package had just released the lock) |
| 2 | the same, fresh image copy, probes `logread`/`dmesg`/`free` only | Booted; console after 56.4 s; `init complete` was logged by 573 s of guest uptime (the wait loop's last check), `br-lan` up at 463 s; 4 probes exit 0; clean power-off, driver 0, QEMU 0. 06:53:37–07:03:44 UTC. Host load average at the start 1.26 (1 min) |
| 3 | **initramfs kernel, direct boot:** `-kernel <w>/kernel.bin -append "console=ttyAMA0,115200 earlycon"` in place of `-bios` and the root `-drive`, plus `-virtfs` (above); otherwise the same line | Booted; console after 48.4 s; `/` is tmpfs (14,164 kB used; after `init complete` `free` showed 21,972 kB used, 445,252 kB available), kernel `Memory: 493344K/524288K available`; `init complete` about 670 s after kernel start; 4 probes exit 0, 9p transfer OK; clean power-off, driver 0, QEMU 0. 07:04:45–07:16:57 UTC |
| 4 | as run 2, sampling `top` every 40 s during the first boot | Booted; console after 68.2 s; host load average 6.05 at the start (other packages were working). 9 `top -b -n 1` samples from 67 s to 485 s of guest uptime, then `init complete` by 622 s; 3 probes exit 0; clean power-off, driver 0, QEMU 0. 07:17:22–07:28:55 UTC |

Figures from run 1 (TCG on an x86 host, two vCPUs, 512 MB; **not** a performance figure for real hardware):

- **Boot time.**
  - Host clock from QEMU start: `Linux version` on the console after 12.8 s (U-Boot, its 2 s autoboot countdown and
    GRUB's 5 s timeout), `Please press Enter to activate this console` after 62.3 s.
  - Kernel timestamps: root mounted at 1.89 s, `Run /sbin/init` at 2.25 s, `init: - preinit -` at 8.30 s, overlay
    formatted and mounted by 35.76 s, `procd: - init -` at 48.63 s, module loading done at 83.90 s.
  - `procd: - init complete -` (`logread`) was logged at 06:43:50; the RTC set the clock to 06:36:51 at 1.49 s, so
    init completed about 420 s after kernel start. `br-lan` and `eth1` came up at about 361 s (kernel log).
  - **Where the time goes (run 4, `top -b -n 1` every 40 s).** From 67 s to about 345 s of uptime the first-boot
    scripts ran one after another (`S10boot`, then `S10system`, then `S12log`) while one kernel worker thread,
    `[kworker/0:1+events]`, took 59–78 % CPU in five of the six samples; what that worker did was not determined.
    At 395 s `dropbearkey -t rsa` was generating the RSA host key (48 % CPU); by 438 s both dropbear host keys
    existed, and by 485 s uhttpd had generated its EC key and self-signed certificate (`logread`: `Generating EC
    private key`). So most of the first boot is first-boot work under emulation; a second boot of the same disk was
    not measured.
- **Memory** (guest `free`, 506,668 kB total; kernel `Memory: 503392K/524288K available`):
  - after `init complete`: used 24,128 kB, free 456,368 kB, available 444,788 kB;
  - at the end of the run (uptime 826 s, after `apk update` and all tests): used 30,916 kB, available 435,176 kB;
    `/tmp` 1,748 kB used (the apk indexes), `/overlay` 39,384 kB used.
- **Storage after boot:** section 11.
- **Listening ports** (guest `netstat -tulnp`):
  - TCP 22: dropbear, all addresses (IPv4 and IPv6)
  - TCP 80 and 443: uhttpd, all addresses
  - TCP and UDP 53: dnsmasq, on loopback, the LAN address, the WAN address 192.0.2.15, and IPv6 link-local, site
    and ULA addresses
  - UDP 546: odhcp6c; UDP 547: odhcpd
- **Running services** (guest `ps w`): procd, ubusd, logd, rpcd, urngd, dropbear, netifd, odhcpd, udhcpc on eth1,
  odhcp6c on eth1, uhttpd, dnsmasq (inside `ujail`, user `dnsmasq`), ntpd (inside `ujail`, user `ntp`), askfirst
  with the login shell on the console. Kernel threads include `watchdogd` and the four f2fs threads.
- **Security observations** (the same as x86-64):
  - The banner says "There is no root password defined on this device!".
  - `/etc/config/dropbear` has `PasswordAuth 'on'` and `RootPasswordAuth 'on'`.
  - LuCI listens on 80 and 443 on all addresses; the default firewall rejects input from the wan zone.
  - The kernel logs `CPU0: Spectre v2: firmware did not set auxiliary control register IBE bit, system vulnerable`
    under QEMU's firmware.
- **Kernel messages under TCG:** `hrtimer: interrupt took 77335719 ns`. This comes from the emulation and host load,
  not from the image.

## 13. Fields not determined

- **usign signature of `sha256sums`:** not checked; no usign/signify tool on the PC and the signing key
  (`b5043e70f9a75cde`) is not in the tree or the image (section 1).
- **Boot time and memory on real 32-bit ARM hardware:** not determined; only TCG emulation on an x86 PC is possible
  here, and no armv7 board is part of step 1.
- **Ext4 image boot:** not run. It was taken apart only, and its file and package lists equal the squashfs image's.
- **Watchdog device:** not determined; the guest had no `/dev/watchdog*` (QEMU was given no watchdog device), so
  procd's watchdog stayed `offline`.
- **Which CPUs an `ARM_LPAE=y` kernel excludes:** not checked (no Kconfig help text in the SDK).
- **Why 37 kmod packages carry no module on armv7:** not determined (section 5).
- **Running the SDK's compiler:** not done in this package; compiler, binutils and C library versions were read from
  the files (section 15). A first cross build belongs to the build packages.
- **Boot time of a second (not first) boot:** not measured; every run started from a fresh image copy, so all
  figures include the first-boot work (overlay format, host keys, TLS certificate).
- **What the busy kernel worker (`kworker/0:1+events`) did during the first boot:** not determined (section 12).
- **A writable 9p share and the ext4 combined image under QEMU:** not tried.

## 14. Tools missing on the PC

`usign`/`signify`, `parted`, `expect`, `socat` and `extract-ikconfig` are missing, and there is no root, so no loop
mounts. Present and used: `gzip`, `xz`, `zstd`, `tar`, `cpio`, `sgdisk`, `fdisk`, `mtools`, `debugfs`,
`unsquashfs` (squashfs-tools 4.6.1), `binwalk` 2.3.4, `file`, `strings`, `gpg`, `qemu-system-arm` 8.2.2 and `flock`.
`cpio` and `bsdtar` are now installed (they were missing for I-01).

## 15. SDK, ImageBuilder and toolchain of the tree (5.3 point 4)

The archives were only listed (`zstd -dc … | tar -tvf -`, streamed; exit 0 each) and single files were extracted
by name; the large extractions were deleted afterwards. Listing sizes are the sum of the entries.

| Archive | Entries | Unpacked size (bytes) | What it is for NMB |
| :--- | ---: | ---: | :--- |
| SDK `openwrt-sdk-…_gcc-14.3.0_musl_eabi.Linux-x86_64` | 35,256 | 1,044,061,685 | Building a first `nmbd` `.apk` package for this target with OpenWrt's package Makefiles, and the C11/musl builds; carries the kernel tree's `.config` and headers |
| Toolchain `openwrt-toolchain-…_gcc-14.3.0_musl_eabi.Linux-x86_64` | 2,988 | 278,917,017 | The bare cross compiler and musl sysroot for C11/musl builds outside the OpenWrt build system |
| ImageBuilder `openwrt-imagebuilder-…Linux-x86_64` | 3,739 | 212,995,925 | Assembling an image with NMB's own package selection from the published packages, without rebuilding the base |
| `kernel-debug.tar.zst` | 1,355 | — | `debug/vmlinux` (53,101,528 bytes) and 1,352 unstripped modules: symbols for kernel crash analysis only |
| `llvm-bpf-21.1.6.Linux-x86_64` | 405 | 210,924,030 | clang/llc for BPF programs (SDK option `CONFIG_SDK_LLVM_BPF=y`); NMB does not plan BPF code |

- **Compiler: GCC 14.3.0** (`OpenWrt GCC 14.3.0 r33311-981c978eff`, `strings` of
  `bin/.arm-openwrt-linux-muslgnueabi-gcc-14.3.0.bin`). Target triplet `arm-openwrt-linux-muslgnueabi`
  (`info.mk`: `TARGET_CROSS=arm-openwrt-linux-muslgnueabi-`). The compiler binaries are x86-64 ELF
  (`file`) started through small wrapper scripts with the archive's own host loader and `runas.so`.
  Toolchain folder name: `toolchain-arm_cortex-a15+neon-vfpv4_gcc-14.3.0_musl_eabi`.
- **C library: musl 1.2.5** (`info.mk`: `LIBC_TYPE=musl`, `LIBC_VERSION=1.2.5`; `strings lib/libc.so` contains
  `1.2.5`; `file`: ELF 32-bit ARM EABI5 shared object). The same in the SDK and in the separate toolchain archive
  (both `info.mk` identical).
- **Binutils: 2.44** (`strings` of `arm-openwrt-linux-muslgnueabi/bin/.ld.bin`: `(GNU Binutils) 2.44`; the kernel
  version string says `GNU ld (GNU Binutils) 2.44`).
- **Debugger:** `gdb` 16.3 (`strings` of `.arm-openwrt-linux-muslgnueabi-gdb.bin`).
- **Feeds the SDK builds against** (`feeds.conf.default`): base `openwrt.git^64900e756efa667774214f9c6b17c113222e01f7`,
  packages `^093d5516…`, luci `^75a306df…`, routing `^48c80803…`, telephony `^2618106d…`, video `^094bf58d…` (the same
  commits as `feeds.buildinfo`).
- **ImageBuilder.** It carries only three packages itself (`base-files`, `kernel`, `libc`); everything else comes
  from the feeds in its `repositories` file (the same eight URLs as the image's `distfeeds.list`), so building an
  image needs network access or a local mirror. Its `.config`: squashfs, ext4, targz, cpio and initramfs outputs,
  `CONFIG_GRUB_EFI_IMAGES=y`, GRUB timeout 5, kernel partition 128 MB, root partition 104 MB. Signing key for apk:
  `keys/openwrt-25.12.pem` (the same name as the image's `/etc/apk/keys/openwrt-25.12.pem`).
- **The tree's package folders.** `packages/` holds the 88 target-specific packages (for example `grub2-efi-arm`,
  `uboot-envtools`, `iptables-nft`, `umbim`, `uqmi`, the ATM tools), `kmods/` the 1,121 kmod packages of this kernel.
- **Use for NMB (5.3 point 4):**
  - Go code for this target is cross-built without the SDK (`GOARCH=arm`, `GOARM=7`, no cgo); the SDK is needed
    for an OpenWrt `.apk` package of `nmbd`, for any C11/musl part, and for packages that link against OpenWrt's
    libraries.
  - The kernel `.config` above confirms that no own kernel is needed for bridge, VLAN, nftables, PPPoE, WireGuard
    (module), namespaces, seccomp and cgroup2; Landlock is not available, and Wi-Fi (`cfg80211`/`mac80211`) comes
    only as backports modules.
  - The base root (`rootfs.tar.gz`, 131 packages, musl armhf) is the quick user-mode test root for arm v7 binaries
    (5.3 point 3); the generic image is the full guest.
