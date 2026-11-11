# Image record: OpenWrt 25.12.6 x86/generic (32-bit x86; squashfs and ext4, combined and combined-efi)

Package I-04 (work plan section 5.4): the 32-bit x86 tree, as a cross-check of I-01 (OpenWrt 25.12.6 x86-64).
The method is work plan 5.1, and every field of 5.2 is below, under the same headings as the I-01 record so that
I-09 can compare them. Sections 15 to 17 are added for this package: the build trees (5.3 point 4), the
differences from I-01, and the work log. Each value names the command that produced it. A value that could not
be read says "not determined" and why. Nothing here comes from memory. Only lists, versions and checksums are
recorded. No file taken out of the image is in this repository.

- Date of the analysis: 2026-11-11.
- Where it ran: dev PC, Windows with WSL2 Ubuntu 24.04, about 1.8 GB RAM and 4 threads, no root. The QEMU guests
  ran on the same PC.
- Work folder: `~/nmb-work/I-04/` (outside every repository; one fresh sub-folder per unpack).
- Material, read only: `~/src/NMB-depythonization-firmwares/openwrt-target-trees/openwrt-25.12.6-x86-generic/`.

## 1. Files and checksums

The tree holds the whole published target folder. Every file below was checked against the tree's own
`sha256sums` (1,282 lines, SHA-256 `9d367bee4aa225d32c3c922f082436d6b965d1be5106e891687af49e90d63b08`).

| File | Size (bytes) | SHA-256 (publisher list) | Check |
| :--- | ---: | :--- | :--- |
| `openwrt-25.12.6-x86-generic-generic-squashfs-combined.img.gz` | 11,548,556 | `a133620d7afe07fedb0f83f43ad00ae7071d7bb80a2214f7ea4c7698cc79654e` | OK |
| `openwrt-25.12.6-x86-generic-generic-squashfs-combined-efi.img.gz` | 11,724,212 | `eac36da875b212ff9f989b1cc260838c3326a66609535a4643f4b06e574617f0` | OK |
| `openwrt-25.12.6-x86-generic-generic-ext4-combined.img.gz` | 12,725,259 | `30f7d36486f9b7cc879434112c8c6abb1f7501b3192eee85da914b21529cb906` | OK |
| `openwrt-25.12.6-x86-generic-generic-ext4-combined-efi.img.gz` | 12,899,053 | `1d7b733c42551b5c50c53c924937325aabdc84634c23d12e23d06b742e1701ff` | OK |
| `openwrt-25.12.6-x86-generic-generic-squashfs-rootfs.img.gz` | 5,511,970 | `c276be48e2dfdb4dcc25ff3a4308609c8e7bf9e5594a39ff52381da75252d9aa` | OK |
| `openwrt-25.12.6-x86-generic-generic-ext4-rootfs.img.gz` | 6,690,871 | `b6543f60715777ae8e9ddb81888d4f9547ce3f829c7bbcb8b89e4cc43785e070` | OK |
| `openwrt-25.12.6-x86-generic-generic-targz-rootfs.tar.gz` | 6,575,472 | `ca1d1dcbc84a16faee72c814e2abdd63a2d61e0bdf528222c9c26a2f5be42d9c` | OK |
| `openwrt-25.12.6-x86-generic-rootfs.tar.gz` | 4,688,633 | `89c853268b391279d38df966be90a0bf0986468063042822bedfe843a94ed120` | OK |
| `openwrt-25.12.6-x86-generic-generic-kernel.bin` | 5,628,416 | `e6059a4b9efccd50ee25ca760b2887997de2907b80cc36209c89e161086629fc` | OK |
| `openwrt-sdk-25.12.6-x86-generic_gcc-14.3.0_musl.Linux-x86_64.tar.zst` | 275,319,725 | `c04cec9e97f73cb13dab6135c397640478ea5a15694efe9bf8a5ea03fe9e0927` | OK |
| `openwrt-toolchain-25.12.6-x86-generic_gcc-14.3.0_musl.Linux-x86_64.tar.zst` | 75,770,711 | `602881789bc6ea29cbd44de14f229b075ea91f6a57341065358e7b7ff8832928` | OK |
| `openwrt-imagebuilder-25.12.6-x86-generic.Linux-x86_64.tar.zst` | 37,337,093 | `39594039b9ad3e0f14e9ddd8928b2831f4e73add2556ee6dbca7cccd52ca25cb` | OK |
| `kernel-debug.tar.zst` | 126,829,484 | `ada2acf5d7046348f6ab497b9bf2e107061ded3768ae876c137f926191f9b6bb` | OK |
| `llvm-bpf-21.1.6.Linux-x86_64.tar.zst` | 51,309,188 | `c290b7d9df8e1d4a9e466a3d5b7204e4a2ecc18e121e066ebc0e3881d3f91c1f` | OK |
| `openwrt-25.12.6-x86-generic.manifest` | 4,227 | `e4d60113c7c738282cc5fffdbc8a67ae2b4924cbcaecfba7213c11458d68d289` | OK |
| `profiles.json` | 3,053 | `e5f154de68a5c88c8f49c599bd7e733499e84a304845676dce459e2aa12d03b4` | OK |
| `openwrt-25.12.6-x86-generic.bom.cdx.json` | 18,186 | `89bf116bd187ddcdc247e02114d75c58e767c5f97a1acde2621b0102d4b769a0` | OK |
| `config.buildinfo` | 2,199 | `2fc104739546634e222a6430acb78581a7292f01f385f22573358349a003d105` | OK |
| `feeds.buildinfo` | 487 | `037f68494e8cf059f7cb49dbce7e0f73a3da4db1e73b2b8b7477ff2d2676f40b` | OK |
| `version.buildinfo` | 18 | `47cc6a2e45cd952f6b4d5edaf3ba2259a038a8f7f834b37edcdac78f97078c9c` | OK |

- **How it was checked.** `grep -E '\*(openwrt-25|openwrt-sdk|openwrt-image|openwrt-tool|llvm|kernel-debug|profiles|config.build|feeds.build|version.build)' sha256sums | sha256sum -c`
  printed 20 `OK` lines, one for each of the 20 top-level files in the list (all 20 are in the table above), and
  exit 0. The `bom.cdx.json` file matches the `*openwrt-25` pattern; sizes are from `stat -c %s`. 20 top-level,
  1,165 `kmods/…` and 97 `packages/…` lines make up the 1,282 lines of the list.
  - The 1,165 `kmods/…` lines and the 97 `packages/…` lines of the list also check `OK` (`sha256sum -c --quiet`,
    exit 0). Only the four signature files `packages.adb.asc` and `packages.adb.sig` (in `kmods/…/` and
    `packages/`) are not in the list, which is expected.
- **Signature of the list (gpg).** `gpg --verify sha256sums.asc sha256sums` in a fresh, isolated `GNUPGHOME`
  gave "Good signature from OpenWrt Build System (Nitrokey3)" (exit 0). The signature was made on 2026-10-18 at
  05:45:39 UTC with the signing subkey `92C5 61DE 55AE 6552 F3C7 36B8 2B01 5109 0606 D1D9`, primary key
  `8A8B C12F 46B8 36C0 F9CD B36F 1D53 D187 7742 E911`.
  - **Where the key came from.** This package fetched nothing from the network. The key is the copy that I-01
    fetched from keys.openpgp.org on 2026-11-11 and compared with OpenWrt's keyring repository
    (`~/nmb-work/I-01/pub/key.asc`, SHA-256 `f5a73197ea11a233f0783aba5080ff47ec803c28b90f773d783b7b18a24a6b46`).
    gpg says "not certified with a trusted signature", because no web of trust was set up.
  - The same key also gives "Good signature" on `packages/packages.adb.asc` and on
    `kmods/6.12.112-1-87a1956e9c0e98a41de66ca5d2374ff9/packages.adb.asc`.
- **Signature of the list (usign): not checked.** `sha256sums.sig` names the key `b5043e70f9a75cde` (bytes 3–10 of
  the signature). The only usign key in the image, `/etc/opkg/keys/8a11255d14aef6c8`, is the "ASU CA pubkey 2022",
  a different key. The release key is not in the material, and `usign` is not installed on the PC, so this check
  would need a key from the network.
- **Release identity.** `version.buildinfo` reads `r33311-981c978eff`, and the image's own `/etc/openwrt_release`
  gives `DISTRIB_REVISION='r33311-981c978eff'`, `DISTRIB_TARGET='x86/generic'`, `DISTRIB_ARCH='i386_pentium4'`.
  This is the same release as I-01.
- **Decompressed images.** `gzip -dc` gave exit 0 for all four combined images and printed no "trailing garbage"
  warning.
  - squashfs-combined: 126,353,408 bytes, `1f87bd4e626fc9d3738268f1de8dcca1a273179e29bd1562ca44b808db4494b4`
  - squashfs-combined-efi: 126,123,520 bytes, `061efadce5c2824d22506a453f84d4f52d393e11cfd2b0ae13c172e5f2b85e9f`
  - ext4-combined: 126,353,408 bytes, `66b16a1a646ed014d92339ad515950c3faf396019ef54c33f9d1f58dae31374e`
  - ext4-combined-efi: 126,123,520 bytes, `38983b2bca0ea89f4629e4aadfadcab79abf48eac104624fbfccf4cb7b7d904d`

## 2. Partition and file system layout, boot method

The layout was read with `sgdisk -p disk.img` (exit 0) and `fdisk -l` (exit 0). The tree has two layouts, and the
squashfs and ext4 images of each layout differ only in partition 2.

**combined-efi (GPT), as in I-01:**

| No. | Sectors (512 B) | Size | GPT type | Contents (`file`) |
| :--- | :--- | :--- | :--- | :--- |
| 128 | 34–511 | 239 KiB | BIOS boot (`EF02`) | GRUB core for legacy BIOS boot |
| 1 | 512–33279 | 16.0 MiB | Linux filesystem (`8300`), not the EFI System type | FAT, label `kernel` |
| 2 | 33280–246271 | 104.0 MiB | Linux filesystem | squashfs image: squashfs 4.0, xz, 5,410,898 bytes, 1,273 inodes. ext4 image: ext4, label `rootfs`, 26,624 blocks of 4 KiB, 20,317 free (`debugfs -R stats`) |

**combined (MBR, no EFI):**

| No. | Sectors (512 B) | Size | MBR type | Contents (`file`) |
| :--- | :--- | :--- | :--- | :--- |
| 1 (boot flag) | 512–33279 | 16.0 MiB | `83` | ext2-format file system, label `kernel` (no journal; the guest mounts it with the ext4 driver) |
| 2 | 33792–246783 | 104.0 MiB | `83` | the same squashfs or ext4 as in the combined-efi image (`cmp` of the two squashfs partitions: identical) |

- **Partition table (combined-efi).** A GPT behind a protective MBR. As in I-01 the backup GPT header is missing;
  sgdisk reports "invalid backup GPT header", and fdisk says "The backup GPT table is not on the end of the device".
  In the guest of run 1 (combined image, MBR) no GPT message appears. In run 3 (combined-efi) the kernel printed
  `Alternate GPT is invalid, using primary GPT.` and booted on (as in I-01).
- **Boot partition (combined-efi).** Read with `mdir -/ -i p1.img ::` and `mcopy -s` (both exit 0). It holds
  `/boot/vmlinuz` (5,628,416 bytes), `/boot/grub/{boot.img,core.img,grub.cfg}` and **`/efi/boot/bootia32.efi`**
  (438,272 bytes, `file`: "PE32 executable (EFI application) Intel 80386"). There is **no `bootx64.efi`**, so the
  image boots only from 32-bit UEFI firmware or from legacy BIOS.
- **Boot partition (combined).** Read with `debugfs -R "rdump / …" p1.img` (exit 0; ownership warnings only). It
  holds `/boot/vmlinuz` and `/boot/grub/{boot.img,core.img,grub.cfg}`.
  - `vmlinuz` is the same file in both layouts and equals the publisher's `…-generic-kernel.bin` (SHA-256
    `e6059a4b…629fc`, `sha256sum`).
- **Boot method.**
  - Legacy BIOS: SeaBIOS 1.16.3 loads GRUB 2.12 from the disk (run 1, console: `Booting from Hard Disk...`,
    `GNU GRUB version 2.12`). GRUB boots the default entry after the 5 s timeout. The combined image's `grub.cfg`
    sets `root='(hd0,msdos1)'` and has no `search` line, and run 1 printed no `error:` line (`grep -c 'error:'` = 0).
    The combined-efi `grub.cfg` has `search -l kernel -s root`, as in I-01.
  - Legacy BIOS on the GPT image (run 3): SeaBIOS printed `Booting from Hard Disk...` and `GRUB loading....`,
    loaded the GRUB core from the BIOS boot partition 128, and GRUB 2.12 booted `OpenWrt` after the timeout; no
    `error:` line (`grep -c 'error:'` = 0). The kernel found its root by `root=PARTUUID=<disk id>-02` and mounted
    partition 1 (FAT) on `/boot` (guest `mount`: `vfat`). So the combined-efi image also serves legacy-BIOS PCs.
  - UEFI: **not determined.** The image carries only an IA32 EFI loader, and the PC has only x64 OVMF
    (`/usr/share/OVMF/OVMF_CODE_4M.fd`, package `ovmf 2024.02-2ubuntu0.10`); no 32-bit OVMF is installed.
  - Kernel command line, from the guest's `cat /proc/cmdline`: `BOOT_IMAGE=/boot/vmlinuz root=PARTUUID=<disk id>-02
    rootwait console=tty1 console=ttyS0,38400n8 noinitrd`. **The serial console runs at 38400 baud**
    (`CONFIG_GRUB_BAUDRATE=38400` in the ImageBuilder `.config`), where I-01 uses 115200. There is no initramfs.
  - A second GRUB entry boots with `failsafe=true`.
- **Root file system at run time (squashfs, run 1).** From the guest's `mount` and `df -k`: the squashfs is mounted
  read-only on `/rom`; the rest of partition 2 is `/dev/loop0` (202,368 sectors), formatted ext4 on the first boot
  (`mount_root: overlay filesystem in /dev/loop0 has not been formatted yet`, then `switching to ext4 overlay`);
  `overlayfs:/overlay` is `/`; partition 1 is mounted read-write on `/boot` (ext4 driver, combined image); `/tmp` is
  tmpfs; cgroup2 is on `/sys/fs/cgroup`.
- **Root file system (ext4 image).** Not booted, only taken apart: `debugfs -R "rdump / out" p2.img` (exit 0). Its
  file list equals the squashfs root (`find -type f | sort`, `diff`, exit 0).
- **Unpacking the squashfs.** `unsquashfs -no-xattrs -d … p2.img` gave exit 2; the one failure was
  `could not create character device …/dev/console, because you're not superuser`, expected without root. It
  created 928 files, 121 directories and 223 symlinks.
- **Root archives.** `…-generic-targz-rootfs.tar.gz` has the same 928 files as the squashfs root
  (`tar -tzvf`, `diff`, exit 0). `…-x86-generic-rootfs.tar.gz` is the target's default root without the device
  packages: 753 files and 136 packages (`tar -xzOf … ./lib/apk/db/installed`). The 43 packages it lacks are the
  NIC, DRM/framebuffer and firmware packages of the `generic` profile.

## 3. How an update is written

- **sysupgrade.** Present: `/sbin/sysupgrade` (guest `command -v`). The x86 rule `/lib/upgrade/platform.sh` is
  **byte-identical to I-01's** (`diff`, exit 0): `platform_check_image` accepts the boot-sector magic `eb48`/`eb63`
  and compares partition tables, writing the full image if the layout changed; `platform_copy_config` copies the
  saved settings into partition 1 (vfat or ext4); `platform_do_bootloader_upgrade` reruns `grub-bios-setup`.
- **No A/B scheme.** One root partition, overwritten in place; settings travel in the backup file.
- **Other update helpers.** `owut` and `attendedsysupgrade-common` are installed (package list), as in I-01.
- **Packages at run time.** `apk` installs into the overlay (run 1: 33 packages added, section 9a).

## 4. Kernel and the configuration NMB needs

- **Kernel version: 6.12.112, 32-bit.** Sources:
  - `file vmlinuz`: "Linux kernel x86 boot executable bzImage, version 6.12.112 (builder@buildhost) #0 SMP Sat Oct
    17 21:45:38 2026, RO-rootFS".
  - `strings` of the decompressed payload (xz stream at offset 16532 found by `binwalk`, `tail -c +16533 vmlinuz |
    xz -dc --single-stream`, exit 0; `file`: "ELF 32-bit LSB executable, Intel 80386"): `Linux version 6.12.112
    (builder@buildhost) (i486-openwrt-linux-musl-gcc (OpenWrt GCC 14.3.0 r33311-981c978eff) 14.3.0, GNU ld (GNU
    Binutils) 2.44) #0 SMP Sat Oct 17 21:45:38 2026`.
  - Guest `uname -a`: `Linux OpenWrt 6.12.112 #0 SMP Sat Oct 17 21:45:38 2026 i686 GNU/Linux`; `ubus call system
    board` gives `"kernel": "6.12.112"`.
  - Installed package `kernel 6.12.112~87a1956e9c0e98a41de66ca5d2374ff9-r1`; `profiles.json` gives the same
    vermagic `87a1956e9c0e98a41de66ca5d2374ff9` (I-01: `b4125196…`).
- **Where the config came from: not in the image.** The decompressed payload has no `IKCFG_ST` marker
  (`LC_ALL=C grep -c -a IKCFG_ST` = 0), the guest has no `/proc/config.gz`, and the image carries no `.config`.
- **Sources used instead, all from the checked tree, all identical.**
  - The official `kmod-ikconfig-6.12.112-r1.apk` from the tree's `kmods/` folder (SHA-256 `3bbd6aee…4c137`, `OK`
    in the signed list). The ADB file (`ADBd`, raw deflate) was inflated with a 7-line Python script run with `-I`
    from `~/nmb-work/I-04/scripts/`, outside the data folders, and the gzip blob between `IKCFG_ST` and `IKCFG_ED`
    was decompressed (exit 0).
  - The SDK's and the ImageBuilder's `build_dir/target-i386_pentium4_musl/linux-x86_generic/linux-6.12.112/.config`
    (single files taken out of the archives with `zstd -dc | tar -xf - <path>`, exit 0).
  - All three have SHA-256 `7c324c8551642a4a8eee90f9e30b21ffdac4fde4bb1c781416ab3f38fa890a0f`, 8,375 lines,
    starting `# Linux/x86 6.12.112 Kernel Configuration` and
    `CONFIG_CC_VERSION_TEXT="i486-openwrt-linux-musl-gcc (OpenWrt GCC 14.3.0 r33311-981c978eff) 14.3.0"`.
  - The `configs.ko` in `kernel-debug.tar.zst` (23,492 bytes) has no `IKCFG_ST` marker (`grep -c -a` = 0), so it
    was not usable as a source.
- **Source of every option below: this config.** "In image" = the module is in `lib/modules/6.12.112/` (67 `.ko`
  files). "Package" = a kmod package is in the tree's `kmods/` folder (1,163 kmod packages). The last column
  compares with I-01's config (`c936e1f6…df3a`, the file I-01 extracted, read from `~/nmb-work/I-01/kmod/`).

| Feature (5.2) | Option(s) | Value | In image | Package | Same as I-01 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Bridge, VLAN filtering | `BRIDGE`, `BRIDGE_VLAN_FILTERING` | y, y | built in | — | yes |
| Bonding | `BONDING` | m | no | `kmod-bonding` | yes |
| Team | `NET_TEAM` | m | no | `kmod-team` | yes |
| VLAN 802.1Q | `VLAN_8021Q` | y | built in | — | yes |
| nftables | `NF_TABLES` m, `NF_TABLES_INET` y, `NF_TABLES_NETDEV` y, `NF_TABLES_BRIDGE` m, `NFT_NAT` m, `NFT_MASQ` m, `NFT_FIB_INET` m, `NF_FLOW_TABLE` m | as listed | yes (`nf_tables.ko`, `nft_*.ko`) | `kmod-nft-*` | yes |
| WireGuard | `WIREGUARD` | m | no | `kmod-wireguard` (installed and loaded in run 1: `ip link add … type wireguard` worked) | yes |
| PPP, PPPoE | `PPP` m, `PPPOE` m, `PPP_ASYNC` m | m | yes (`ppp_generic`, `pppoe`, `pppox`, `ppp_async`) | `kmod-ppp`, `kmod-pppoe` | yes |
| cfg80211, mac80211 | `CFG80211` | **not set in the kernel** | no | `kmod-cfg80211`, `kmod-mac80211` from backports `6.18.55` (`6.12.112.6.18.55-r2`) | yes |
| mac80211_hwsim | — | from backports | no | `kmod-mac80211-hwsim-6.12.112.6.18.55-r2` (installed in run 1; `radios=2` gave `phy0`, `phy1` in `iw dev`) | yes |
| Wi-Fi drivers (examples) | — | from backports / mt76 | no | `kmod-ath9k`, `kmod-ath10k(-ct)`, `kmod-iwlwifi`, `kmod-mt76*`, `kmod-mt7601u`, `kmod-rtw88-*`, `kmod-rtw89-*` | yes (same Wi-Fi package names) |
| USB host | `USB` y, `USB_XHCI_HCD` y, `USB_EHCI_HCD` y | built in | — | — | yes |
| USB network | `USB_USBNET` m, `USB_RTL8152` m, `USB_NET_AX88179_178A` m, `USB_NET_SMSC95XX` m, `USB_NET_AQC111` m | m | no | `kmod-usb-net-*` | yes |
| USB tethering (4.27) | `USB_NET_RNDIS_HOST`, `USB_NET_CDCETHER`, `USB_NET_CDC_NCM`, `USB_IPHETH`, `USB_NET_QMI_WWAN`, `USB_NET_CDC_MBIM`, `USB_NET_HUAWEI_CDC_NCM` | all m | no | `kmod-usb-net-rndis`, `-cdc-ether`, `-cdc-ncm`, `-ipheth`, `-qmi-wwan`, `-cdc-mbim` | yes |
| TUN, veth, macvlan, dummy, IFB | `TUN`, `VETH`, `MACVLAN`, `DUMMY`, `IFB` | all m | no | `kmod-tun`, `kmod-veth`, `kmod-macvlan`, `kmod-dummy` | yes |
| Policy routing | `IP_ADVANCED_ROUTER`, `IP_MULTIPLE_TABLES`, `IPV6_MULTIPLE_TABLES` | y | built in | — | yes |
| Traffic control | `NET_SCHED` y, `NET_SCH_FQ_CODEL` y, `NET_SCH_CAKE` m, `NET_SCH_HTB` m, `NET_SCH_NETEM` m, `NET_SCH_TBF` m | as listed | fq_codel built in | `kmod-sched-cake`, `kmod-netem` | yes |
| Namespaces | `NAMESPACES`, `NET_NS`, `PID_NS`, `USER_NS`, `UTS_NS`, `IPC_NS` | all y | built in | — | yes |
| seccomp | `SECCOMP`, `SECCOMP_FILTER` | y, y | built in | (`procd-seccomp` installed) | yes |
| **Landlock** | `SECURITY` | **not set**, so `SECURITY_LANDLOCK` is absent | — | not available without an own kernel | yes |
| cgroups | `CGROUPS` y, `MEMCG` y (`MEMCG_V1` not set), `CGROUP_PIDS` y, `CPUSETS` y, `CGROUP_BPF` y, `BLK_CGROUP` y; `CGROUP_FREEZER`, `CGROUP_DEVICE` not set | as listed | cgroup2 on `/sys/fs/cgroup` (guest `mount`) | — | yes |
| Watchdog | `WATCHDOG` y; `SOFT_WATCHDOG`, `I6300ESB_WDT`, `ITCO_WDT` m | as listed | no `/dev/watchdog*` in the guest; `ubus call system watchdog`: `"status": "offline", "timeout": 30` | `kmod-softdog`, `kmod-i6300esb-wdt`, `kmod-itco-wdt` | yes |
| Module signing | `MODULE_SIG` | not set | — | — | yes |
| **CPU and memory model** | `X86_32` y, `MPENTIUM4` y, `X86_GENERIC` y, `X86_MINIMUM_CPU_FAMILY=6`, `HIGHMEM64G` y with `X86_PAE` y, `VMSPLIT_3G` y, `PGTABLE_LEVELS=3` | as listed | — | — | **no** (I-01: `X86_64`, `PGTABLE_LEVELS=4`) |
| **CPU count** | `NR_CPUS` | **4** | — | — | **no** (I-01: 512) |
| **Other differences that touch NMB** | `IO_URING` **not set**, `ARCH_MMAP_RND_BITS=8` (I-01: 28), `VMAP_STACK` absent, `IOMMU_SUPPORT` not set, `BLK_DEV_NVME` **m** (I-01: y) | as listed | `kmod-nvme` is not installed | `kmod-nvme` | **no** |
| Other | `BPF_SYSCALL` y, `XDP_SOCKETS` not set, `TCP_CONG_BBR` m, `IPV6` y, `PREEMPT_NONE` y, `HZ=100`, `VIRTIO_NET` y, `VIRTIO_BLK` y | as listed | — | — | yes |

- **Whole-config comparison** (`CONFIG_` lines, `# … is not set` read as `n`, `LC_ALL=C sort -t= -k1,1`, then
  name and value split at the **first** `=` with `awk` and compared by name; files `~/nmb-work/I-04/cmp/a.cfg`
  (I-01) and `b.cfg` (I-04); `awk 'NR==FNR{i=index($0,"=");a[substr($0,1,i-1)]=substr($0,i+1);next}
  {i=index($0,"=");k=substr($0,1,i-1);v=substr($0,i+1);if(k in a){if(a[k]!=v)d++;s[k]=1}else ob++}
  END{for(k in a)if(!(k in s))oa++;print d,oa,ob}' a.cfg b.cfg` printed `95 323 192`):
  95 options have different values, 323 options exist only in I-01's config (199 of them `y`/`m`, almost all
  64-bit-only: IOMMU, VFIO, Xen PV, AVX/SSSE3 crypto, `MLXSW`, `X86_X2APIC`, `VMAP_STACK`) and 192 only in this one
  (56 `y`/`m`: `X86_32`, `HIGHMEM64G`, `X86_PAE`, ISA, Geode/SCx200, `UID16`, 586 crypto). None of the NMB
  options of 5.2 differs (the rows above marked "yes").
  - Correction after verification: a first count gave 96, because `join -t=` split the value of
    `CONFIG_CC_IMPLICIT_FALLTHROUGH="-Wimplicit-fallthrough=5"` at its inner `=` and reported equal values as
    different. Splitting at the first `=` gives 95; an independent comparison by the verifier also gave 95.
- **What follows for NMB on 32-bit x86.** A box without NVMe support in the image cannot boot from an NVMe disk
  until `kmod-nvme` is added in an image build; Go programs for `GOARCH=386` need SSE2 by default, which the
  `MPENTIUM4` kernel and `-march=pentium4` userland also require; `io_uring` is unavailable; one kernel uses at
  most 4 CPUs; ASLR has 8 bits of mmap randomisation instead of 28.

## 5. Modules and firmware files carried

- **Modules.** 67 `.ko` files in `lib/modules/6.12.112/` (`ls *.ko | wc -l`; guest `lsmod | wc -l` = 67), plus
  `modules.builtin` (278 lines). Names:
  `3c59x 8139too 8390 button-hotplug crc-ccitt drm_buddy drm_display_helper drm_exec drm_suballoc_helper e100 e1000
  fat forcedeth i2c-dev i915 libcrc32c mii natsemi ne2k-pci nf_conntrack nf_defrag_ipv4 nf_defrag_ipv6 nf_flow_table
  nf_flow_table_inet nf_log_syslog nf_nat nf_reject_ipv4 nf_reject_ipv6 nf_tables nfnetlink nft_chain_nat nft_ct
  nft_fib nft_fib_inet nft_fib_ipv4 nft_fib_ipv6 nft_flow_offload nft_hash nft_limit nft_log nft_masq nft_nat
  nft_numgen nft_quota nft_redir nft_reject nft_reject_inet nft_reject_ipv4 nft_reject_ipv6 nls_cp437 nls_iso8859-1
  nls_utf8 pcnet32 ppp_async ppp_generic pppoe pppox r8169 realtek sis900 slhc tg3 vfat via-rhine via-velocity video
  wmi`.
  - Groups: old and common PCI NIC drivers (`3c59x`, `8139too`, `e100`, `e1000`, `natsemi`, `ne2k-pci`, `pcnet32`,
    `sis900`, `via-rhine`, `via-velocity`, `forcedeth`, `r8169`, `tg3`), nftables and conntrack, PPP/PPPoE, `i915`
    with DRM helpers, and `vfat`/NLS. virtio-net and virtio-blk are built in.
  - I-01 has instead the newer server/SoC NIC drivers (`e1000e`, `igb`, `igc`, `ixgbe`, `bnx2`, `amd-xgbe`, `ena`,
    `stmmac`/`dwmac-intel`); see section 16.
- **Firmware.** 68 files under `lib/firmware`, 1,198,205 bytes (`find -type f | wc -l`, `du -sb`): `e100/` (3),
  `i915/` (33) and `rtl_nic/` (32). No Wi-Fi firmware is in the image.
  - The three firmware packages (`e100-firmware`, `i915-firmware-dmc`, `r8169-firmware`) have an empty licence
    field in the apk database (`grep -c '^L:$'` = 3).

## 6. Init system, service manager, C library, package manager

- **Init:** procd (`/sbin/procd` is PID 1 in the guest's `ps w`) with ubus, ujail and seccomp (`procd-ujail`,
  `procd-seccomp`). Init scripts in `/etc/init.d`, linked in `/etc/rc.d`; `etc/init.d` is identical to I-01's
  (`diff -r`, exit 0).
- **Enabled services** (guest loop `$s enabled`): boot, cron, dnsmasq, done, dropbear, firewall, gpio_switch, led,
  log, network, odhcpd, packet_steering, rpcd, sysctl, sysfixtime, sysntpd, system, ucitrack, uhttpd, umount,
  urandom_seed, urngd (22, the same as I-01).
- **Services known to procd** (guest `ubus call service list`): cron, dnsmasq, dropbear, firewall, gpio_switch, log,
  network, odhcpd, packet_steering, rpcd, sysntpd, system, ubus, ucitrack, uhttpd, urandom_seed, urngd.
- **C library: musl 1.2.5, i386.** The guest's `/lib/libc.so` prints `musl libc (i386)` / `Version 1.2.5`.
  Package `libc 1.2.5-r5`; `/lib/ld-musl-i386.so.1` is a link to `libc.so` (`ls -l`). `file usr/sbin/dnsmasq`:
  "ELF 32-bit LSB pie executable, Intel 80386 … interpreter /lib/ld-musl-i386.so.1".
- **Package manager: apk-tools 3.0.5** (guest `apk --version`: `apk-tools 3.0.5, compiled for x86.`); `apk` arch
  `i386_pentium4` (`/etc/apk/arch`). opkg is absent (guest `command -v opkg`).
  - Feeds in `/etc/apk/repositories.d/distfeeds.list`: target `x86/generic`, kmods
    `6.12.112-1-87a1956e9c0e98a41de66ca5d2374ff9`, and base, luci, packages, routing, telephony, video for
    `i386_pentium4`, release 25.12.6. apk key: `/etc/apk/keys/openwrt-25.12.pem`.
- **Installed packages: 179.** `awk '/^P:/…' lib/apk/db/installed` and guest `apk info | wc -l` both give 179.
  The squashfs and ext4 images give the same list. The list is in section 10.
  - The published `openwrt-25.12.6-x86-generic.manifest` lists the 136 packages of the plain target root
    (`…-x86-generic-rootfs.tar.gz`); the 43 device packages of the `generic` profile in `profiles.json` (NIC drivers,
    `kmod-fs-vfat`, `kmod-drm-i915` and their dependencies) come on top.
- **Licences of the installed packages** (`L:` fields, `sort | uniq -c`): GPL-2.0-only 73, ISC 29, GPL-2.0 27,
  Apache-2.0 15 (LuCI), LGPL-2.1 7, MIT 4, GPL-3.0-with-GCC-exception 4 (libc, libgcc, libpthread, librt),
  GPL-2.0-or-later 4, GPL-3.0-or-later 3 (GRUB), LGPL-2.1-or-later 2, BSD-4-Clause 2, and one each of Zlib,
  LGPL-2.1+, "GPL-2.0-or-later MPL-2.0", GPL-2.0+, "GPL-2.0 BSD-3-Clause", BSD-3-Clause; 3 empty (the firmware).

## 7. Tools of 6A criterion 3, with versions

The presence check is the guest's `command -v` loop (run 1). Versions come from the guest's `--version` output
where it exists, otherwise from the installed package. **Every userland version is the same as in I-01**: the
two package lists differ in version only for `kernel` (`join` of the two lists, section 16).

| Tool | Present | Version |
| :--- | :--- | :--- |
| procd (init) | yes, `/sbin/procd` | `procd 2026.03.13~58eb263d-r1` (package) |
| systemd, OpenRC | no | — |
| nftables `nft` | yes, `/usr/sbin/nft` | `nftables v1.1.6 (Commodore Bullmoose #7)`; package `nftables-json 1.1.6-r2` |
| dnsmasq | yes | `Dnsmasq version 2.93`; package `dnsmasq 2.93-r1` (not `dnsmasq-full`) |
| hostapd | **no** | (feeds: section 9a) |
| wpa_supplicant | **no** | (feeds: section 9a) |
| iw | **no** | (feeds: section 9a; installed in run 1 for the hwsim test) |
| pppd | yes | `pppd version 2.5.2`; `ppp 2.5.2-r4`, `ppp-mod-pppoe 2.5.2-r4` |
| WireGuard (`wg`, module) | **no** | (feeds: section 9a) |
| OpenVPN | **no** | (feeds: section 9a) |
| `ip` | yes, the **BusyBox applet** (`/sbin/ip -> ../bin/busybox`) | BusyBox v1.37.0; `ip link add … type bridge` and `type wireguard` worked |
| `tc`, `ethtool`, `iptables`, `ss` | no | — |
| DHCP client | `udhcpc` (BusyBox), `odhcp6c 2026.06.27~10a52220-r3` | — |
| DHCP/RA server | `odhcpd-ipv6only 2026.06.29~5d7be43f-r1` and dnsmasq | — |
| SSH | `dropbear` | `Dropbear v2025.89` |
| Web | `uhttpd 2026.08.24~373145f7-r1` with LuCI `26.288.56989~75a306d` | — |
| HTTP client | `uclient-fetch` | `uclient-fetch 2026.06.07~daad21fa-r1` |
| BusyBox | yes | `busybox 1.37.0-r6`; `strings bin/busybox`: `BusyBox v1.37.0 (2026-10-17 21:45:38 UTC)`. (The guest's bare `busybox` call printed "applet not found", so its banner is not used.) |
| `ucode` | yes, `/usr/bin/ucode` | `ucode 2026.01.16~85922056-r2` |
| `lua`, `perl`, `nproc` | no | — |

## 8. Python (6A criterion 8)

**No Python is in the image.** The checks:

- In the unpacked squashfs root, `find . -iname "python*" -o -name "*.py" -o -name "*.pyc"` found 0 files.
- `LC_ALL=C grep -rlI "^#!.*python" .` found 0 files.
- In the guest, `find / -xdev …` and `find /rom -xdev …` both found 0.
- In the guest, `command -v python python3 perl lua` printed absent for all four.
- `python3` and `python3-base` `3.13.18-r1` exist in the packages feed and are not installed (section 9a).
- The build trees carry Python for the build host only (section 15): the SDK and ImageBuilder have
  `staging_dir/host/bin/python3` as a link to the build machine's `/opt/venv/bin/python3.9` (dangling here), and the
  toolchain has GCC's libstdc++ pretty-printers under `share/gcc-14.3.0/python/`. None of it reaches an image.

## 9. Own network and firewall manager, and how each is switched off (6A criterion 5)

The guest runs **netifd** (`/sbin/netifd`, package `netifd 2026.02.26~cbb83a18-r2`) and **firewall4**
(`/sbin/fw4`, ucode, `firewall4 2025.03.17~b6e51575-r2`). The ruleset is one table, `table inet fw4`
(guest `nft list tables`), with 145 lines (`nft list ruleset | wc -l`), as in I-01.

**Default configuration** (guest `uci show network` and `uci show firewall`; `etc/config` and `etc/board.d` are
identical to I-01's, `diff -r`, exit 0):

- `eth0` is bridged into `br-lan`, which gets the static built-in LAN default address.
- `eth1` is `wan` with DHCP, and `wan6` with DHCPv6.
- A ULA prefix is generated on the first boot.
- Firewall defaults: input and forward `REJECT`; zone `lan` accepts; zone `wan` (`wan`, `wan6`) rejects input,
  drops forward, masquerades.

Coexistence tests in the guest of run 1 (the same probe lines as I-01 run 4; each probe ended with exit 0):

| Test | Command | Seen |
| :--- | :--- | :--- |
| Own nft table across firewall actions | `nft add table inet nmbtest` + chain; `fw4 stop`; `fw4 start`; `fw4 reload`; `/etc/init.d/firewall restart` | `inet nmbtest` **kept** after all four (`fw4 stop` left only `inet nmbtest`) |
| Firewall service stop | `/etc/init.d/firewall stop`; `nft list tables` | **empty**: NMB's table is deleted too (the service stop runs `fw4 flush`). `start` brought back `inet fw4` only |
| Firewall disable | `disable`; `enabled; echo $?`; `ls /etc/rc.d \| grep -c firewall` | `enabled` gave 1 and 0 links; `enable` gave 1 link again |
| netifd stop | `/etc/init.d/network stop`; `pgrep -l netifd` | no netifd process (`pg=1`); `start` brought it back |
| netifd disable | `disable`; `ls /etc/rc.d \| grep -c network` | 0 links; `enable` gave 2 links |
| Own bridge `nmbbr0` (198.51.100.77/24) | `network reload`, `restart`, `stop`, `start` | present after all of them; `ip route` still had `198.51.100.0/24 dev nmbbr0` |
| Own address 203.0.113.9/24 on `br-lan` | `network reload` + 10 s, then `restart` + 15 s | **kept** by `reload`, **removed** by `restart` |

**What this means for criterion 5:** the same as I-01. Both managers can be switched off with `disable` (plus
`stop` for netifd). The firewall service's `stop` flushes the whole nftables ruleset, so NMB must use `fw4 stop` or
`disable` and a reboot. netifd leaves foreign devices alone and drops foreign addresses on its own devices on
`restart`.

### 9a. Feed lookup and function tests (run 1)

In the guest, over user-mode networking: `apk update` gave rc 0 and `OK: 11142 distinct packages available`
(I-01: 11234). Then `apk search $p | grep "^$p-[0-9]"` per name. **None of these packages is installed in the
image.** All versions equal I-01's.

| Package | Version in the 25.12.6 feeds (`i386_pentium4`, `x86/generic`) |
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
| `python3` / `python3-base` | `3.13.18-r1` (in the feed only; never installed for NMB, criterion 8) |

Function tests (not run in I-01):

- `apk add kmod-mac80211-hwsim iw kmod-wireguard` installed 33 packages (`OK: 17.5 MiB in 212 packages`, rc 0).
  The kernel logged `Loading modules backported from Linux version v6.18.55-0-g725bd2f3c` and
  `wireguard: WireGuard 1.0.0 loaded`.
- `modprobe mac80211_hwsim radios=2` gave two radios, `phy0` and `phy1` (`iw list`).
- `ip link add nmbwg0 type wireguard` created the link (`mtu 1420`), and `ip link del` removed it. `lsmod` then
  showed `cfg80211`, `mac80211`, `mac80211_hwsim`, `wireguard`, `compat`, `udp_tunnel`, `ip6_udp_tunnel`,
  `libchacha20poly1305`, `libcurve25519_generic`.

## 10. Installed packages (179, `lib/apk/db/installed`; the same in the squashfs and ext4 images)

`apk-mbedtls 3.0.5-r3`, `attendedsysupgrade-common 10`, `base-files 1718~981c978eff`, `busybox 1.37.0-r6`, `ca-bundle 20260816-r1`, `cgi-io 2026.07.21~31cb3c89-r1`, `dnsmasq 2.93-r1`, `dropbear 2025.89-r3`, `e100-firmware 20260221-r1`, `e2fsprogs 1.47.3-r1`, `firewall4 2025.03.17~b6e51575-r2`, `fstools 2026.05.23~16718b6e-r1`, `fwtool 2025.10.03~04cd252e-r1`, `getrandom 2025.10.30~6f78fa49-r1`, `grub2 2.12-r1`, `grub2-bios-setup 2.12-r1`, `grub2-efi 2.12-r1`, `i915-firmware-dmc 20260221-r1`, `jansson4 2.14.1-r1`, `jshn 2026.06.19~7dd12784-r1`, `jsonfilter 2026.03.16~b9034210-r1`, `kernel 6.12.112~87a1956e9c0e98a41de66ca5d2374ff9-r1`, `kmod-3c59x 6.12.112-r1`, `kmod-8139too 6.12.112-r1`, `kmod-acpi-video 6.12.112-r1`, `kmod-backlight 6.12.112-r1`, `kmod-button-hotplug 6.12.112-r3`, `kmod-crypto-crc32c 6.12.112-r1`, `kmod-crypto-hash 6.12.112-r1`, `kmod-dma-buf 6.12.112-r1`, `kmod-drm 6.12.112-r1`, `kmod-drm-buddy 6.12.112-r1`, `kmod-drm-display-helper 6.12.112-r1`, `kmod-drm-exec 6.12.112-r1`, `kmod-drm-i915 6.12.112-r1`, `kmod-drm-kms-helper 6.12.112-r1`, `kmod-drm-suballoc-helper 6.12.112-r1`, `kmod-drm-ttm 6.12.112-r1`, `kmod-drm-ttm-helper 6.12.112-r1`, `kmod-e100 6.12.112-r1`, `kmod-e1000 6.12.112-r1`, `kmod-fb 6.12.112-r1`, `kmod-fb-cfb-copyarea 6.12.112-r1`, `kmod-fb-cfb-fillrect 6.12.112-r1`, `kmod-fb-cfb-imgblt 6.12.112-r1`, `kmod-fb-sys-fops 6.12.112-r1`, `kmod-fb-sys-ram 6.12.112-r1`, `kmod-forcedeth 6.12.112-r1`, `kmod-fs-vfat 6.12.112-r1`, `kmod-hwmon-core 6.12.112-r1`, `kmod-i2c-algo-bit 6.12.112-r1`, `kmod-i2c-core 6.12.112-r1`, `kmod-input-core 6.12.112-r1`, `kmod-lib-crc-ccitt 6.12.112-r1`, `kmod-lib-crc32c 6.12.112-r1`, `kmod-libphy 6.12.112-r1`, `kmod-mdio-devres 6.12.112-r1`, `kmod-mii 6.12.112-r1`, `kmod-natsemi 6.12.112-r1`, `kmod-ne2k-pci 6.12.112-r1`, `kmod-nf-conntrack 6.12.112-r1`, `kmod-nf-conntrack6 6.12.112-r1`, `kmod-nf-flow 6.12.112-r1`, `kmod-nf-log 6.12.112-r1`, `kmod-nf-log6 6.12.112-r1`, `kmod-nf-nat 6.12.112-r1`, `kmod-nf-reject 6.12.112-r1`, `kmod-nf-reject6 6.12.112-r1`, `kmod-nfnetlink 6.12.112-r1`, `kmod-nft-core 6.12.112-r1`, `kmod-nft-fib 6.12.112-r1`, `kmod-nft-nat 6.12.112-r1`, `kmod-nft-offload 6.12.112-r1`, `kmod-nls-base 6.12.112-r1`, `kmod-nls-cp437 6.12.112-r1`, `kmod-nls-iso8859-1 6.12.112-r1`, `kmod-nls-utf8 6.12.112-r1`, `kmod-pcnet32 6.12.112-r1`, `kmod-phy-realtek 6.12.112-r1`, `kmod-ppp 6.12.112-r1`, `kmod-pppoe 6.12.112-r1`, `kmod-pppox 6.12.112-r1`, `kmod-pps 6.12.112-r1`, `kmod-ptp 6.12.112-r1`, `kmod-r8169 6.12.112-r1`, `kmod-sis900 6.12.112-r1`, `kmod-slhc 6.12.112-r1`, `kmod-tg3 6.12.112-r1`, `kmod-via-rhine 6.12.112-r1`, `kmod-via-velocity 6.12.112-r1`, `libblkid1 2.41.5-r1`, `libblobmsg-json20260213 2026.06.19~7dd12784-r1`, `libc 1.2.5-r5`, `libcomerr0 1.47.3-r1`, `libe2p2 1.47.3-r1`, `libext2fs2 1.47.3-r1`, `libf2fs6 1.16.0-r4`, `libgcc1 14.3.0-r5`, `libiwinfo-data 2026.05.26~66bdd1a0-r1`, `libiwinfo20230701 2026.05.26~66bdd1a0-r1`, `libjson-c5 0.18-r1`, `libjson-script20260213 2026.06.19~7dd12784-r1`, `liblucihttp-ucode 2023.03.15~9b5b683f-r1`, `liblucihttp0 2023.03.15~9b5b683f-r1`, `libmbedtls21 3.6.7-r1`, `libmnl0 1.0.5-r1`, `libnftnl11 1.3.1-r1`, `libnl-tiny1 2025.12.02~40493a65-r1`, `libpthread 1.2.5-r5`, `librt 1.2.5-r5`, `libsmartcols1 2.41.5-r1`, `libss2 1.47.3-r1`, `libubox20260213 2026.06.19~7dd12784-r1`, `libubus20251202 2026.06.28~24864e78-r2`, `libuci20250120 2025.12.02~66127cd7-r1`, `libuclient20201210 2026.06.07~daad21fa-r1`, `libucode20230711 2026.01.16~85922056-r2`, `libudebug 2026.01.16~875e1a7a`, `libustream-mbedtls20201210 2026.03.01~99f1c0db-r1`, `libuuid1 2.41.5-r1`, `logd 2025.10.30~6f78fa49-r1`, `luci 26.288.56989~75a306d`, `luci-app-attendedsysupgrade 26.288.56989~75a306d`, `luci-app-firewall 26.288.56989~75a306d`, `luci-app-package-manager 26.288.56989~75a306d`, `luci-base 26.288.56989~75a306d`, `luci-lib-uqr 26.288.56989~75a306d`, `luci-light 26.288.56989~75a306d`, `luci-mod-admin-full 26.288.56989~75a306d`, `luci-mod-network 26.288.56989~75a306d`, `luci-mod-status 26.288.56989~75a306d`, `luci-mod-system 26.288.56989~75a306d`, `luci-proto-ipv6 26.288.56989~75a306d`, `luci-proto-ppp 26.288.56989~75a306d`, `luci-ssl 26.288.56989~75a306d`, `luci-theme-bootstrap 26.288.56989~75a306d`, `mkf2fs 1.16.0-r4`, `mtd 27`, `netifd 2026.02.26~cbb83a18-r2`, `nftables-json 1.1.6-r2`, `odhcp6c 2026.06.27~10a52220-r3`, `odhcpd-ipv6only 2026.06.29~5d7be43f-r1`, `openwrt-keyring 2025.12.10~f0670054-r1`, `owut 2026.08.30~9b51eeec-r1`, `partx-utils 2.41.5-r1`, `ppp 2.5.2-r4`, `ppp-mod-pppoe 2.5.2-r4`, `procd 2026.03.13~58eb263d-r1`, `procd-seccomp 2026.03.13~58eb263d-r1`, `procd-ujail 2026.03.13~58eb263d-r1`, `px5g-mbedtls 11`, `r8169-firmware 20260221-r1`, `rpcd 2026.10.04~d99f703e-r1`, `rpcd-mod-file 2026.10.04~d99f703e-r1`, `rpcd-mod-iwinfo 2026.10.04~d99f703e-r1`, `rpcd-mod-luci 20240305-r1`, `rpcd-mod-rpcsys 2026.10.04~d99f703e-r1`, `rpcd-mod-rrdns 20170710`, `rpcd-mod-ucode 2026.10.04~d99f703e-r1`, `ubox 2025.10.30~6f78fa49-r1`, `ubus 2026.06.28~24864e78-r2`, `ubusd 2026.06.28~24864e78-r2`, `uci 2025.12.02~66127cd7-r1`, `uclient-fetch 2026.06.07~daad21fa-r1`, `ucode 2026.01.16~85922056-r2`, `ucode-mod-fs 2026.01.16~85922056-r2`, `ucode-mod-html 1`, `ucode-mod-log 2026.01.16~85922056-r2`, `ucode-mod-math 2026.01.16~85922056-r2`, `ucode-mod-ubus 2026.01.16~85922056-r2`, `ucode-mod-uci 2026.01.16~85922056-r2`, `ucode-mod-uclient 2026.06.07~daad21fa-r1`, `ucode-mod-uloop 2026.01.16~85922056-r2`, `uhttpd 2026.08.24~373145f7-r1`, `uhttpd-mod-ubus 2026.08.24~373145f7-r1`, `urandom-seed 3`, `urngd 2025.10.03~f17e33d9-r1`, `usign 2025.10.03~c4c72b1b-r1`, `zlib 1.3.1-r1`

## 11. Size on disk

- Download: 11,548,556 bytes (squashfs-combined), 11,724,212 (squashfs-combined-efi), 12,725,259 (ext4-combined),
  12,899,053 (ext4-combined-efi), as `.img.gz`.
- Disk image: 126,353,408 bytes (combined) and 126,123,520 bytes (combined-efi).
- squashfs root: 5,410,898 bytes compressed (I-01: 6,096,242). Unpacked, the root is 16,519,194 bytes (`du -sb`;
  I-01: 20,384,863). Of that (`du -sb` per folder):
  - `lib/modules`: 4,536,449 bytes (I-01: 7,971,111)
  - `lib/firmware`: 1,198,205 bytes (I-01: 1,411,732)
  - `www` (LuCI): 1,337,317 bytes (the same as I-01)
  - `usr/lib`: 3,616,020 bytes
- Guest `df -k` after the first boot (run 1, squashfs-combined):
  - `/rom`: 5,376 kB, 100 % used
  - `/overlay` (ext4 on loop0): 89,195 kB, 62 kB used, 82,051 kB free
  - `/boot`: 16,084 kB, 5,680 kB used, 10,080 kB free
  - `/tmp`: tmpfs, 122,116 kB
- After installing 33 packages (end of run 1): `/overlay` 2,255 kB used, `/tmp` 1,784 kB used.
- ext4 image: the 104 MiB root has 20,317 of 26,624 4-KiB blocks free (about 79.4 MiB), from `debugfs -R stats`.

## 12. QEMU boot test (squashfs image)

**How it was run.** Each boot ran on a fresh copy of a decompressed image, under `flock -w … ~/nmb-work/qemu.lock`
for the whole run, one guest at a time (`~/nmb-work/I-04/scripts/run.sh`). The QEMU command line:

```
qemu-system-i386 -machine pc -accel <kvm|tcg,tb-size=128> -cpu <host|max> -smp 1 -m <256|512> -nographic -monitor none
  -serial unix:<work>/serial.sock,server=on,wait=on
  -drive id=d0,file=disk.img,format=raw,if=none -device virtio-blk-pci,drive=d0
  -netdev user,id=lan0,net=198.51.100.0/24,restrict=on -device virtio-net-pci,netdev=lan0,mac=02:00:00:00:04:01
  -netdev user,id=wan0,net=192.0.2.0/24             -device virtio-net-pci,netdev=wan0,mac=02:00:00:00:04:02
```

- QEMU 8.2.2 (Debian 1:8.2.2+ds-0ubuntu1.18), SeaBIOS 1.16.3 (package `seabios 1.16.3-2`), machine `pc`
  (i440FX + PIIX).
- The serial console was driven by a small Python script (`python3 -I`, kept outside the data folders) that waits
  for markers and sends input in 32-byte chunks.
- The guest saw `pae cx8 cmov sse2 nx` in `/proc/cpuinfo` (run 1).

| Run | Image, settings | Result |
| :--- | :--- | :--- |
| 1 | squashfs-combined (MBR, legacy BIOS), **KVM** (`-cpu host`), `-smp 1 -m 256` | Booted cleanly (`Booting paravirtualized kernel on KVM`, `Hypervisor detected: KVM`). All 34 probes gave exit 0. `poweroff` was clean: `reboot: Power down`, QEMU exit 0, driver exit 0, flock exit 0; 110 s in all (06:51:47–06:53:37 UTC) |
| 2 | squashfs-combined-efi through its BIOS boot partition, KVM, `-m 512` | **Did not start**: `Could not access KVM kernel module: Permission denied` (QEMU exit 1, driver exit 1). The user's ACL entry on `/dev/kvm` was there before run 1 (`getfacl` listed a named-user entry with `rw-` for the working user) and gone after it (`getfacl` shows no user entry; the node's change time is 06:53:36 UTC). Nothing in this package changed it; there is no sudo |
| 3 | squashfs-combined-efi through its BIOS boot partition, **TCG** (`-cpu max`, `-accel tcg,tb-size=128`), `-smp 1 -m 512`; fresh copy of the decompressed image (SHA-256 `061efadc…85e9f` checked before the copy) | Booted cleanly through SeaBIOS and GRUB 2.12 from the GPT image (`Booting paravirtualized kernel on bare hardware`; `ubus call system board`: `"system": "QEMU TCG CPU version 2.5+"`, `"rootfs_type": "squashfs"`). All 10 probes (`scripts/probes2.txt`) gave exit 0. `poweroff` was clean: `reboot: Power down`, QEMU exit 0, driver exit 0, flock+run exit 0; 360 s in all (07:41:01–07:47:01 UTC). The first probe polled `logread` before `logd` was up, so the console shows 22 `Failed to find log object: Not found` lines from that loop; the probe still ended with exit 0 |

Figures from run 3 (TCG, one vCPU, 512 MB; emulation on the dev PC, **not** a figure for real hardware):

- **Boot time.** Host clock (driver timestamps from QEMU start): kernel banner at 25.5 s, `Please press Enter to
  activate this console` at 62.9 s, shell prompt at 74.5 s. Kernel timestamps: `Run /sbin/init` at 7.10 s,
  `procd: - init -` at 36.53 s (logged as 07:42:43 UTC, so kernel start was about 07:42:06). `procd: - init
  complete -` was logged at 07:46:08 UTC, about 241 s after kernel start (1 s resolution); most of that is the
  modules and services starting under emulation (`kmodloader: done` at 57.4 s, `br-lan` forwarding at 185.3 s).
  The first boot includes formatting the overlay.
- **Memory.** Kernel line: `Memory: 501724K/523756K available (10324K kernel code, …, 0K highmem)`; `HighMem empty`
  (512 MB lies below the 32-bit low-memory limit, so PAE high memory is not used). Guest `free` after `init
  complete`: total 503,628 kB, used 23,464 kB, free 457,812 kB, available 442,428 kB; at the end (uptime 302 s)
  used 23,388 kB.
- **Storage** (guest `df -k`, combined-efi): `/rom` 5,376 kB (100 %), `/overlay` 89,195 kB with 62 kB used,
  `/boot` (FAT) 16,334 kB with 6,090 kB used. `/proc/partitions`: `vda1` 16,384, `vda2` 106,496, `vda128` 239,
  `loop0` 101,184 (1 KiB blocks).
- **CPU:** `QEMU TCG CPU version 2.5+`, flags `pae cx8 cmov sse2 nx`. `nproc` is not in the image (`-ash: nproc:
  not found`).
- **Listening ports** (`netstat -tulnp`): the same set as run 1 and I-01 (dropbear 22, uhttpd 80 and 443 on all
  addresses; dnsmasq 53 on loopback, LAN, WAN 192.0.2.15, IPv6 link-local and ULA; odhcp6c 546; odhcpd 547).
- **Other probes, the same as run 1:** `uname -m` `i686`; `/etc/openwrt_release` `r33311-981c978eff`,
  `x86/generic`, `i386_pentium4`; musl 1.2.5; `apk-tools 3.0.5`, 179 packages, no `opkg`; 67 modules loaded;
  no `/proc/config.gz`; no `/dev/watchdog*`, `ubus call system watchdog` `"status": "offline"`; cgroup2 on
  `/sys/fs/cgroup`.

Figures from run 1 (KVM, one vCPU, 256 MB; **not** a figure for real hardware, and not comparable with I-01's TCG
figures):

- **Boot time.**
  - Host clock (driver timestamps from QEMU start): kernel banner at 8.4 s, `Please press Enter to activate this
    console` at 17.8 s. This includes SeaBIOS and the 5 s GRUB timeout.
  - Kernel timestamps: `Run /sbin/init` at 3.36 s, `procd: - init -` at 9.43 s.
  - `procd: - init complete -` (from `logread`) at 06:52:13 UTC. The RTC line `setting system clock to
    2026-11-11T06:51:58 UTC` at 2.95 s puts kernel start at about 06:51:55, so init completed about 18 s after
    kernel start (1 s resolution). At uptime 19.5 s the first probe had already found it.
  - The first boot includes formatting the overlay.
- **Memory.** Kernel line: `Memory: 242436K/261612K available (10324K kernel code, …, 0K highmem)`; `HighMem empty`.
  Guest `free` right after `init complete`: total 244,236 kB, used 21,772 kB, free 200,084 kB, available
  185,736 kB. At the end of the run (uptime 94 s, after `apk update`, 33 installed packages and the hwsim and
  WireGuard modules loaded): used 24,240 kB, available 177,984 kB.
- **Storage after boot:** see section 11.
- **Listening ports** (guest `netstat -tulnp`), the same set as I-01:
  - TCP 22: dropbear, all addresses
  - TCP 80 and 443: uhttpd, all addresses
  - TCP and UDP 53: dnsmasq, on loopback, the LAN address, the WAN address 192.0.2.15, and IPv6 link-local and ULA
  - UDP 546: odhcp6c; UDP 547: odhcpd
  - Ephemeral UDP ports: dnsmasq and ntpd
- **Running services** (guest `ps w`): procd, ubusd, logd, rpcd, urngd, dropbear, netifd, odhcpd, udhcpc on eth1,
  odhcp6c on eth1, uhttpd, dnsmasq (inside `ujail`, user `dnsmasq`), ntpd (inside `ujail`, user `ntp`), and askfirst
  with the login shell on the console. A `[watchdogd]` kernel thread runs, but there is no `/dev/watchdog`.
- **Security observations** (as in I-01): `/etc/config/dropbear` has `PasswordAuth 'on'` and
  `RootPasswordAuth 'on'`; uhttpd listens on `0.0.0.0` and `[::]` for 80 and 443; the default firewall rejects
  input from the wan zone. No root password is set.

## 13. Fields not determined

- **UEFI boot:** not determined. The combined-efi images carry only `bootia32.efi`, and the PC has no 32-bit OVMF.
- **Boot time and memory on real 32-bit hardware:** not determined; no such machine was used.
- **Ext4 image boot:** not run. It was taken apart only, and its file list equals the squashfs image's.
- **usign signature of `sha256sums`:** not checked; the release usign key is not in the material and `usign` is not
  installed (section 1). The gpg signature was checked.
- **Licence field of the three firmware packages:** empty in the apk database.
- **Toolchain versions by running the tools:** not determined on purpose. The toolchain binaries are untrusted host
  executables, so their versions were read without running them (section 15).

## 14. Tools missing on the PC

`cpio`, `bsdtar`, `xorriso`, `usign`, a 32-bit (IA32) OVMF and `extract-ikconfig` are missing, and there is no sudo,
so no loop mounts. `/dev/kvm` was usable for run 1 only (section 12).

Present and used: `gzip`, `xz`, `zstd`, `tar`, `sgdisk`, `fdisk`, `mtools`, `debugfs`, `unsquashfs`
(squashfs-tools 4.6.1), `binwalk`, `file`, `strings`, `gpg`, `qemu-system-i386` 8.2.2, SeaBIOS 1.16.3 and `flock`.

## 15. Build trees: SDK, toolchain, ImageBuilder, kernel-debug, llvm-bpf (5.3 point 4)

Each archive was listed with `zstd -dc … | tar -tvf -` (exit 0/0 for all four), and only single files were taken
out. No compiler or other binary from the trees was run.

| Archive | Entries | Bytes unpacked (sum of `tar -tv` sizes) |
| :--- | ---: | ---: |
| SDK | 36,096 | 1,147,252,475 |
| toolchain | 3,196 | 307,898,841 |
| ImageBuilder | 955 | 188,229,919 |
| llvm-bpf 21.1.6 | 405 | 210,924,030 |
| kernel-debug | 1,378 (1,375 `.ko` with debug information and an unstripped 32-bit `vmlinux` of 77,697,076 bytes) | — |

- **Compiler: GCC 14.3.0, target `i486-openwrt-linux-musl`.** Sources: the toolchain's
  `toolchain-i386_pentium4_gcc-14.3.0_musl/info.mk` (`TARGET_CROSS=i486-openwrt-linux-musl-`, `GCC_VERSION=14.3.0`),
  identical to the SDK's `staging_dir/toolchain-i386_pentium4_gcc-14.3.0_musl/info.mk` (`cmp`, exit 0), and
  `strings` of `bin/i486-openwrt-linux-musl-gcc-14.3.0` (`OpenWrt GCC 14.3.0 r33311-981c978eff`; configured with
  `--target=i486-openwrt-linux-musl --enable-languages=c,c++ --enable-lto` and target flags
  `-Os -pipe -march=pentium4 … -fstack-protector -D_FORTIFY_SOURCE=1 -Wl,-z,now -Wl,-z,relro`).
  The compiler binary itself is "ELF 64-bit LSB executable, x86-64" (`file`): it runs on an x86-64 build host.
- **C library: musl 1.2.5** (`info.mk`: `LIBC_TYPE=musl`, `LIBC_VERSION=1.2.5`; `strings lib/libc.so` contains
  `1.2.5`). The same as the image.
- **binutils: 2.44** (`strings i486-openwrt-linux-musl/bin/ld.bfd`: `(GNU Binutils) 2.44`; also in the kernel's
  version string).
- **gdb:** `strings bin/i486-openwrt-linux-musl-gdb` contains `16.3`; this is a weak reading, not a version output.
- **Release identity:** the toolchain's `version.mk` reads `REVISION:=r33311-981c978eff`. The SDK's
  `feeds.conf.default` pins the same feed commits as `feeds.buildinfo` (packages `093d5516…`, luci `75a306df…`,
  routing `48c80803…`, telephony `2618106d…`, video `094bf58d…`) plus `base` at `64900e75…`.
- **ImageBuilder:** `.config` has `CONFIG_TARGET_KERNEL_PARTSIZE=16`, `CONFIG_TARGET_ROOTFS_PARTSIZE=104`,
  `CONFIG_GRUB_IMAGES=y`, `CONFIG_GRUB_EFI_IMAGES=y`, `CONFIG_GRUB_BAUDRATE=38400`, `CONFIG_GRUB_TIMEOUT="5"`,
  squashfs, ext4 and targz roots. It carries only 3 packages (`libc`, `kernel`, `base-files`); everything else comes
  from the feeds in its `repositories` file (the same eight feed URLs as the image's `distfeeds.list`), so an image
  build needs network access or a local mirror. `profiles.json` has one profile, `generic` ("Generic x86"), with the
  29 default packages and 15 device packages listed in it, and seven image names.
- **kernel-debug:** debug modules and `vmlinux` for reading crash addresses; its `configs.ko` has no embedded config.
- **llvm-bpf 21.1.6:** a clang/LLVM for building BPF programs on the build host; it contains no Python
  (`grep -ci python` of the listing = 0).
- **What they would be used for (5.3 point 4):**
  - The toolchain (or the SDK's copy of it) is the C11/musl cross compiler for a 32-bit x86 build of the later C code
    base, and the libc and headers to link a Go program with cgo, if one is ever needed for `386`.
  - The SDK builds an `nmbd` apk for `i386_pentium4` against these exact feeds and this kernel (vermagic
    `87a1956e…`).
  - The ImageBuilder assembles an NMB image for 32-bit x86 from the published packages plus that apk, without
    building the base.
  - The kernel `.config` (section 4) confirms that no own kernel is needed for NMB's options on 32-bit x86 either.
  - kernel-debug and llvm-bpf are not needed for the uses listed in 5.3.

## 16. Differences from I-01 (OpenWrt 25.12.6 x86-64)

Compared against the I-01 record and the files I-01 left in `~/nmb-work/I-01/` (its kernel config, its kmod
listing and its unpacked squashfs root).

| Area | I-01 x86-64 | I-04 x86/generic |
| :--- | :--- | :--- |
| Package arch / CPU baseline | `x86_64`, `X86_64`, `GENERIC_CPU` | `i386_pentium4`, `X86_32`, `MPENTIUM4`, `X86_GENERIC`, minimum CPU family 6, PAE required (`HIGHMEM64G`, `X86_PAE`) |
| Compiler prefix | `x86_64-openwrt-linux-musl-gcc` 14.3.0 | `i486-openwrt-linux-musl-gcc` 14.3.0, `-march=pentium4` |
| Kernel vermagic hash | `b4125196ecec75391093e1e2edeba490` | `87a1956e9c0e98a41de66ca5d2374ff9` |
| Kernel config | 8,537 lines | 8,375 lines; 95 values differ, 323 options only in I-01, 192 only here; **no NMB option of 5.2 differs** |
| CPUs, ASLR, io_uring, NVMe | `NR_CPUS=512`, mmap random bits 28, `IO_URING=y`, NVMe built in | `NR_CPUS=4`, 8 bits, `IO_URING` not set, NVMe as a module that is not installed |
| EFI loader | `bootx64.efi` (638,976 bytes) | `bootia32.efi` (438,272 bytes) |
| Serial console | 115200 baud | 38400 baud |
| Layouts in the tree | combined-efi only (in the set) | combined (MBR) and combined-efi (GPT) |
| Installed packages | 182 | 179; the same userland with the same versions. Only in I-01: `bnx2-firmware`, `kmod-amazon-ena`, `kmod-amd-xgbe`, `kmod-bnx2`, `kmod-dwmac-intel`, `kmod-e1000e`, `kmod-igb`, `kmod-igc`, `kmod-ixgbe`, `kmod-mdio`, `kmod-pcs-xpcs`, `kmod-phylink`, `kmod-stmmac-core`. Only here: `e100-firmware`, `kmod-3c59x`, `kmod-8139too`, `kmod-e100`, `kmod-natsemi`, `kmod-ne2k-pci`, `kmod-pcnet32`, `kmod-sis900`, `kmod-via-rhine`, `kmod-via-velocity` |
| Modules in the image | 71 `.ko` | 67 `.ko` |
| Firmware in the image | 70 files, 1,411,732 bytes (`bnx2` 5, `i915` 37, `rtl_nic` 28) | 68 files, 1,198,205 bytes (`e100` 3, `i915` 33, `rtl_nic` 32) |
| kmod packages published | 1,184 | 1,163; none is only here. Only in I-01: `kmod-amazon-ena`, `kmod-amd-xgbe`, `kmod-drm-ivpu`, `kmod-dwmac-intel`, `kmod-fs-efivarfs`, `kmod-i2c-mlxcpld`, `kmod-i2c-mux-mlxcpld`, `kmod-lib-objagg`, `kmod-lib-parman`, `kmod-mlxreg`, `kmod-mlxreg-lc`, `kmod-mlxreg-sn2201`, `kmod-mlxsw-*` (5), `kmod-pcs-xpcs`, `kmod-stmmac-core`, `kmod-vfio`, `kmod-vfio-pci`. Every NMB kmod (WireGuard, bonding, team, cfg80211, mac80211, hwsim, veth, tun, macvlan, dummy, cake, netem, USB tethering, softdog, i6300esb, iTCO) is in both |
| Root size | squashfs 6,096,242 bytes; unpacked 20,384,863 | squashfs 5,410,898; unpacked 16,519,194 |
| Feed size (`apk update`) | 11,234 distinct packages | 11,142 |
| Identical | — | `/etc/config`, `/etc/board.d`, `/etc/init.d`, `/lib/upgrade/platform.sh` (`diff -r`, exit 0); enabled services; listening ports; netifd and fw4 behaviour in every coexistence test; feed versions of every tool in 9a; Python absent |
| QEMU | TCG only (no KVM access then), OVMF | KVM in run 1 (MBR image) and TCG in run 3 (GPT image through its BIOS boot partition), SeaBIOS; run 1 figures are not comparable, run 3 is TCG on the same PC but with a different firmware path and CPU model |

**Cross-check result.** The 32-bit tree confirms I-01 on every point of 5.2 that does not depend on the CPU: the
same kernel options for NMB, the same userland and versions, the same network and firewall managers with the same
behaviour, and no Python. The differences are the CPU model and its limits (PAE and SSE2 needed, at most 4 CPUs,
no `io_uring`, weaker mmap ASLR), the older NIC set, the IA32-only EFI loader and the 38400-baud console.

## 17. Work log (what ran where, with exit codes)

All on the dev PC under WSL2, as the normal user, in `~/nmb-work/I-04/`:

- `sha256sum -c` on the tree's list: 20 top-level files, 1,165 kmods and 97 packages, exit 0; `gpg --verify` on
  `sha256sums.asc` and both `packages.adb.asc`: exit 0.
- `gzip -dc` of four combined images: exit 0 each. `sgdisk -p`, `fdisk -l`: exit 0. `dd` cuts by offset: exit 0.
- `mdir`/`mcopy` on the FAT boot partition: exit 0. `debugfs -R "rdump …"` on the ext2 boot partition and on the
  ext4 root: exit 0 (ownership warnings). `unsquashfs -no-xattrs`: exit 2 (the `/dev/console` node only).
- `binwalk`, `xz -dc --single-stream` of the kernel payload: exit 0.
- `python3 -I scripts/ikcfg.py` on `kmod-ikconfig`: exit 0; on kernel-debug's `configs.ko`: exit 1 (no marker).
- `zstd -dc | tar -tvf -` of SDK, toolchain, ImageBuilder, llvm-bpf, kernel-debug: exit 0. Single-file
  extractions: exit 0, except one toolchain `tar` call with exit 2 (`ld.bfd` is a hard link to a file not named in
  that call; repeated with the link target, exit 0).
- QEMU run 1 (KVM): flock+run exit 0, QEMU exit 0, driver exit 0. Run 2 (KVM): QEMU exit 1, driver exit 1
  (`/dev/kvm` permission denied). Run 3 (TCG, squashfs-combined-efi, 512 MB, under `flock -w 3600
  ~/nmb-work/qemu.lock` for the whole run, started after the lock was free): flock+run exit 0, QEMU exit 0,
  driver exit 0, all 10 probes exit 0.
- Kernel config count re-derived after verification (section 4): `awk` split at the first `=`, output
  `95 323 192`.
- Large extractions removed afterwards: the decompressed disk images, the guest disk copies and the unpacked
  toolchain binaries.
