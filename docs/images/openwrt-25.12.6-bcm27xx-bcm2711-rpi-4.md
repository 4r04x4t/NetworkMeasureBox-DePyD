# Image record: OpenWrt 25.12.6 Raspberry Pi 4 (bcm27xx/bcm2711, squashfs and ext4, factory)

Package I-05 (work plan section 5.4). The method is work plan 5.1, and every field of 5.2 is below.
Each value names the command that produced it. A value that could not be read says "not determined" and why.
Nothing here comes from memory. Only lists, versions and checksums are recorded. No file taken out of the image
is in this repository.

**This image was taken apart only. It was not booted anywhere and not written to any card or device.** The Pi 4
image cannot be booted in QEMU on this PC by the package's rule, and a boot on the owner's Pi 4 needs the backup
procedure of work plan section 7 point 5; that is a later step. Every field that needs a running system says
"not determined: needs a boot on the Pi 4, later step".

- Date of the analysis: 2026-11-11.
- Where it ran: dev PC, Windows with WSL2 Ubuntu 24.04, about 1.8 GB RAM and 4 threads, no root. No QEMU guest and
  no qemu-user run: no binary from the image was executed.
- Work folder: `~/nmb-work/I-05/` (outside every repository; one fresh sub-folder per unpack).
- Material, read only: `~/src/NMB-depythonization-firmwares/platforms/raspberry-pi-4-arm64/`, only the two
  `openwrt-25.12.6-bcm27xx-bcm2711-rpi-4-*-factory.img.gz` files.

## 1. Files and checksums

| File | Size (bytes) | SHA-256 (publisher list) | Check |
| :--- | ---: | :--- | :--- |
| `openwrt-25.12.6-bcm27xx-bcm2711-rpi-4-squashfs-factory.img.gz` | 15,432,229 | `0ddb84a616fdd5afa406297c79d3ff7717718aa7a8d942ed788cedd84a689eca` | OK |
| `openwrt-25.12.6-bcm27xx-bcm2711-rpi-4-ext4-factory.img.gz` | 17,006,660 | `90f6cf7a2fd7e3bf2eab060c7460612be7b43c0c3e075fad01ac80a70a41a917` | OK |

- **How it was checked.** The set's README lists the files but not their OpenWrt sums, and the bcm27xx/bcm2711 tree
  is not in the set (`openwrt-target-trees/` holds only `x86/generic`, `armsr/armv8` and `armsr/armv7`). So the
  publisher's list was fetched on 2026-11-11 at 07:12 UTC from
  `https://downloads.openwrt.org/releases/25.12.6/targets/bcm27xx/bcm2711/sha256sums`, together with
  `sha256sums.asc` (gpg) and `sha256sums.sig` (usign). The list has 1,296 lines and SHA-256
  `c15bd0bd4a93103e1d502017eace5a1362e845b423fccc55a6329185712a7aa3`.
  - In the material folder, `grep -E "rpi-4-(squashfs|ext4)-factory.img.gz" sha256sums | sha256sum -c` printed both
    files `OK` (exit 0). File sizes from `stat -c %s`; they equal the sizes in the publisher's `profiles.json`.
- **gpg signature of the list.** Verified in a fresh, isolated `GNUPGHOME` with `gpg --verify sha256sums.asc sha256sums`:
  "Good signature from OpenWrt Build System (Nitrokey3)" (exit 0).
  - Signing subkey `92C5 61DE 55AE 6552 F3C7 36B8 2B01 5109 0606 D1D9`, primary key
    `8A8B C12F 46B8 36C0 F9CD B36F 1D53 D187 7742 E911`; signature made 2026-10-18 (`gpg --list-packets`).
  - The key was taken from OpenWrt's keyring repository (`github.com/openwrt/keyring`, `gpg/0x1D53D1877742E911.asc`)
    and also fetched from keys.openpgp.org; `gpg --show-keys --with-subkey-fingerprints` gave the same primary and
    subkey fingerprints for both. gpg says "not certified with a trusted signature", because no web of trust was set up.
- **usign signature of the list.** `usign` is not installed on the PC, so the Ed25519 signature was checked with
  OpenSSL 3.0.13 instead: the key number in `sha256sums.sig` is `b5043e70f9a75cde`; the public key of that number was
  taken from the keyring repository (`usign/b5043e70f9a75cde`, comment "Public usign key for unattended snapshot
  builds", file SHA-256 `d7ac10f9…228f`); its 32 key bytes were wrapped in an Ed25519 SubjectPublicKeyInfo and
  `openssl pkeyutl -verify -pubin -rawin -in sha256sums -sigfile sig.bin` printed "Signature Verified Successfully"
  (exit 0). Control: the same check on a copy of the list with one byte appended printed "Signature Verification
  Failure" (exit 1).
- **Release identity.** `r33311-981c978eff`, from the publisher's `version.buildinfo` (in the signed list, `sha256sum -c`
  OK) and from the image's own `/etc/openwrt_release` (`DISTRIB_TARGET='bcm27xx/bcm2711'`,
  `DISTRIB_ARCH='aarch64_cortex-a72'`) and `/etc/openwrt_version`.
- **Decompressed images.** `gzip -dc` gave exit 0 for both and printed no "trailing garbage" warning. The factory
  images carry no fwtool metadata at the end (`tail -c 2048 | strings` empty); the publisher's sysupgrade images
  (not in the set) are 314 bytes larger in `profiles.json`.
  - squashfs: 82,342,912 bytes, `bdee113d8a7bd7c1e38796c8cadb5084bf7a66a0b9d062f64049b57faee86448`
  - ext4: 184,549,376 bytes, `53eeefd3746436f37db106824b467e4d6f6a4cde70915842a56465d34ad33c07`

## 2. Partition and file system layout, boot method

The layout was read with `fdisk -l disk.img` (exit 0). `sgdisk -p` reports "Found invalid GPT and valid MBR": there
is no GPT. It is an MBR (DOS) table, disk identifier `0xd9cce260`, the same in both images.

| No. | Sectors (512 B) | Size | MBR type | Contents (`file`) |
| :--- | :--- | :--- | :--- | :--- |
| 1 (boot flag) | 8192–139263 | 64 MiB | `0c` W95 FAT32 (LBA) | FAT16, label `boot`, serial `6AD3-EC82`. **Identical in both images** (`cmp`, exit 0) |
| 2 | 147456–360447 | 104 MiB | `83` Linux | squashfs image: squashfs 4.0, xz, 5,796,528 bytes, 1,332 inodes, block 256 KiB. ext4 image: ext4, label `rootfs`, 26,624 blocks of 4 KiB, 20,018 blocks free, 6,656 inodes (5,315 free), features `has_journal ext_attr resize_inode filetype extent sparse_super large_file uninit_bg` (`debugfs -R stats`) |

- **Squashfs file is shorter than its table.** The squashfs image file ends at sector 160,825, so only the first
  13,370 sectors of partition 2 (the squashfs itself) are in the file; the table still declares 104 MiB. The ext4
  image file ends exactly at sector 360,447.
- **Partitions were cut out by offset** with `dd bs=512 skip=8192 count=131072` and `dd bs=512 skip=147456`
  (all exit 0).
- **Boot method.** No UEFI, no GRUB, no U-Boot in the image. The partition holds the Raspberry Pi VideoCore
  firmware (`start4*.elf`, `fixup4*.dat`), `config.txt`, `cmdline.txt`, three device trees, 353 overlays and
  `kernel8.img`, a raw arm64 kernel `Image` (`file`: "Linux kernel ARM64 boot executable Image, little-endian, 4K
  pages"). The Pi 4's first-stage bootloader lives in its SPI EEPROM, which is not part of the image (its version:
  not determined, needs the Pi 4).
- **Kernel command line** (`cmdline.txt`, 96 bytes): `console=tty1 console=serial0,115200
  root=PARTUUID=d9cce260-02 rootfstype=squashfs,ext4 rootwait`. No initramfs file is on the boot partition.
  - The device tree `bcm2711-rpi-4-b.dtb` carries its own `/chosen/bootargs` (`dtc -I dtb -O dts`): `coherent_pool=1M
    8250.nr_uarts=1 snd_bcm2835.enable_headphones=0 cgroup_disable=memory numa_policy=interleave`. How the firmware
    merges this with `cmdline.txt`, and so the real `/proc/cmdline`: not determined, needs a boot on the Pi 4, later
    step. If `cgroup_disable=memory` is in effect, the memory cgroup controller is off at run time although the
    kernel has `MEMCG=y` (section 4).
- **Root file system at run time.** Read from the scripts only:
  - `/lib/preinit/79_move_config` mounts partition 1 read-write as vfat on `/boot` during preinit (it loads
    `nls_cp437`, `nls_iso8859-1`, `fat`, `vfat` first) and picks up a saved sysupgrade backup from there.
  - `/lib/preinit/81_set_root_part` rewrites `root=` in `/boot/cmdline.txt` from `/boot/partuuid.txt` after an upgrade.
  - The squashfs image's overlay handling is done by `fstools 2026.05.23~16718b6e-r1`, the same version as in I-01
    (where the rest of partition 2 became an ext4 overlay on a loop device on the first boot). On the Pi 4:
    not determined, needs a boot on the Pi 4, later step.
  - `/etc/fstab` contains only its comment header.
- **Root file system (ext4 image).** The root is the ext4 partition itself, with no overlay.
  - Unpacked with `debugfs -R "rdump / out" p2.img` (exit 0; "Operation not permitted while changing ownership"
    warnings, because there is no root).
  - Its file list is identical to the squashfs root (`find -type f | sort`, 959 files each, then `diff`, exit 0), and
    so is its package list (`diff`, exit 0).
- **Unpacking the squashfs.** `unsquashfs -no-xattrs -d … p2.img` gave exit 2. The one failure was
  `could not create character device …/dev/console, because you're not superuser`, which is expected without root.
  It created 959 files, 128 directories and 244 symlinks.

### 2a. Boot partition contents

Listed with `mdir -/ -i p1.img ::` (exit 0) and copied out with `mcopy -s -n -i p1.img ::/ boot/` (exit 0).
`mdir`: 373 files, 20,268,541 bytes, 46,303,232 bytes free of 64 MiB, cluster 4 sectors.

| File | Size (bytes) | SHA-256 |
| :--- | ---: | :--- |
| `kernel8.img` | 12,886,024 | `4e3fcbbfd14cd404996b553d6c87133d44d3dc4d0fbb4cc6d19003d1c3bb11f9` |
| `start4.elf` | 2,263,968 | `2cc344f6171797317ae0b6197191803b45e2f3cda091f0da31c0a3f0bfc47593` |
| `start4cd.elf` | 814,140 | `3b395c8863690d016ab97fa34ca804bc8ca5b1883af6f308779e6be529db348d` |
| `start4x.elf` | 3,011,592 | `afa1df2e4fad7f451094fc5d92af189a3246b2850c4ee389cef4bdef680e1210` |
| `fixup4.dat` | 5,456 | `ca84b2fae239fc9041ac1bdc8648c1e33b9c6e67c4bd073b651b96dc2726c553` |
| `fixup4cd.dat` | 3,230 | `5c3fbb36333731e503c5a2a4b57edff680d837441d0f098cb5fdb1b0120ef740` |
| `fixup4x.dat` | 8,449 | `9e289996528566451095d6f0bebf48c19cff74c75de8f1836ede731e994be42b` |
| `bcm2711-rpi-4-b.dtb` | 56,211 | `4bc9dd6182025add23a750b18ff748e46d4e939434530e8057c6bfa1ce6f1c16` |
| `bcm2711-rpi-400.dtb` | 56,215 | `3ce06c773943f5bb7b92ebc0a8fa0031980bf9af2511612e986ea58f07feecf5` |
| `bcm2711-rpi-cm4.dtb` | 56,732 | `8d31c84f86b092f44fe4e676b688d206e88a712f8af55e76969cfef2019ff944` |
| `config.txt` | 572 | `7b1e5e909946c77429cf6f355a83c46eb9c6411d821bb7d1f21a2c560d1fd865` |
| `distroconfig.txt` | 678 | `e8a9ef9714f2fe59c91b8cafe36a57c18a73422d1ff9abec7a211b2e2028cd88` |
| `cmdline.txt` | 96 | `39c0c5c434a27dabeff20b495fd78d0949ea5aed4b091d8cbd0a8c6322885d17` |
| `partuuid.txt` | 9 | `f3f8bc0be8a418dcdb360250a7d6ef66a42cf006dc73e87b293bf986351aa310` (content `d9cce260`) |
| `COPYING.linux` | 496 | `fb5a425bd3b3cd6071a3a9aff9909a859e7c1158d54d32e07658398cd67eb6a0` |
| `LICENCE.broadcom` | 1,594 | `c7283ff51f863d93a275c66e3b4cb08021a5dd4d8c1e7acc47d872fbe52d3d6b` |
| `overlays/` (353 `.dtbo` + `README`) | 1,103,079 | SHA-256 of the sorted `sha256sum *.dtbo` output: `274ef126fac1bab4dbab0706fa6faf6bb23684ba04f7c2909f4fa0368b662a89`; `README` `8ada23f1…12e1` |

- **VideoCore firmware build** (`strings start4*.elf | grep VC_BUILD_ID`): all three are branch `bcm2711_2`, built
  Apr 30 2025, version `5560078dcc8591a00f57b9068d13e5544aeef3aa (clean)`, platform `raspberrypi_linux`; variants
  `start`, `start_cd`, `start_x`. They belong to package `bcm27xx-gpu-fw 2025.04.30-r1` (the package lists no file
  in the root; its files are on the boot partition). Licence: `LICENCE.broadcom` (Broadcom / Raspberry Pi
  redistribution licence for binaries); the apk `L:` field of `bcm27xx-gpu-fw` is empty.
- **`config.txt`**: comments, `include distroconfig.txt`, then an empty `[all]` section for own settings. No `kernel=`,
  `arm_64bit=`, `enable_uart=`, `gpu_mem=`, `dtoverlay=` or `dtparam=` lines of its own.
- **`distroconfig.txt`**: `[pi0w]`, `[pi3]` and `[pi4]` sections each with `dtoverlay=disable-bt` (the comment says
  this gives the PL011 UART `ttyAMA0` back to GPIO 14/15 instead of the mini UART), and under `[pi4]` also
  `arm_boost=1`.
- **Device trees** (`fdtget -t s <dtb> / model` and `compatible`): `bcm2711-rpi-4-b.dtb` "Raspberry Pi 4 Model B"
  (`raspberrypi,4-model-b`, `brcm,bcm2711`); `bcm2711-rpi-400.dtb` "Raspberry Pi 400" (`raspberrypi,400`);
  `bcm2711-rpi-cm4.dtb` "Raspberry Pi Compute Module 4" (`raspberrypi,4-compute-module`). These are the three
  `supported_devices` of the `rpi-4` profile in the publisher's `profiles.json`.
- **Overlays.** 353 `.dtbo` files; full name list in section 15. Present among them, relevant for NMB:
  `disable-bt`, `disable-wifi`, `miniuart-bt`, `wifimac`, `sdio`, `dwc2`, `dwc-otg`, `i2c-rtc`, `i2c-rtc-gpio`,
  `rpi-poe`, `rpi-poe-plus`, `act-led`, `gpio-led`, `gpio-shutdown`, `gpio-poweroff`, `ramoops-pi4`,
  `pcie-32bit-dma`, `vl805`, `upstream-pi4`, `vc4-kms-v3d-pi4`.

## 3. How an update is written

- **sysupgrade.** Present: `/sbin/sysupgrade` (identical to I-01's, `cmp`). The Pi rule is `/lib/upgrade/platform.sh`
  (read in the unpacked root; it says "copied from x86's platform.sh" and differs from I-01's file):
  - `REQUIRE_IMAGE_METADATA=1`: sysupgrade wants the fwtool metadata, which the `*-sysupgrade.img.gz` files carry and
    these factory files do not (section 1).
  - `platform_check_image` compares the image's MBR partition map with the card's; if it differs it writes the full
    image ("Partition layout has changed. Full image will be written.").
  - `platform_do_upgrade` writes the whole card with `dd bs=2M` when the layout changed or partitions are not kept,
    otherwise each partition of the image to the matching card partition, then copies the 4-byte disk identifier
    (offset 440).
  - `platform_copy_config` mounts partition 1 on `/boot`, restores `boot/cmdline.txt` and `boot/config.txt` from the
    backup, fixes `root=` (`bcm27xx_set_root_part`) and leaves the settings archive on `/boot` for preinit.
- **No A/B scheme.** One root partition; an update overwrites it, and the boot partition with it. Settings are carried
  in the backup archive. A failed boot has no automatic way back.
- **Other update helpers.** `owut 2026.08.30~9b51eeec-r1` and `attendedsysupgrade-common 10` are installed (package
  list): online image-rebuild helpers.
- **Packages at run time.** `apk` installs into the overlay (squashfs) or into the ext4 root.
- **VideoCore firmware and device trees** are updated only as part of the boot partition in a full image
  (`bcm27xx-gpu-fw` has no root files). The SPI EEPROM bootloader is not handled by this image: no `rpi-eeprom` package
  is installed (package list).

## 4. Kernel and the configuration NMB needs

- **Kernel version: 6.12.112, arm64.** Sources:
  - `strings kernel8.img`: `Linux version 6.12.112 (builder@buildhost) (aarch64-openwrt-linux-musl-gcc (OpenWrt GCC
    14.3.0 r33311-981c978eff) 14.3.0, GNU ld (GNU Binutils) 2.44) #0 SMP Sat Oct 17 21:45:38 2026`
  - the installed package `kernel 6.12.112~46807ee337a11df7be7d383dc0872b8e-r1`, and `linux_kernel` in
    `profiles.json` (`version 6.12.112`, `release 1`, `vermagic 46807ee337a11df7be7d383dc0872b8e`)
  - `modinfo -F vermagic` of the image's modules: `6.12.112 SMP mod_unload aarch64`
  - Uncompressed `Image`; header image size field `0xcb0000` (`xxd -s 0x10`).
- **Where the config came from: not in the image.** `kernel8.img` contains no `IKCFG_ST` marker (`grep -c -a IKCFG_ST`
  printed 0), the image carries no `.config`, and `/proc/config.gz` could not be read without a boot.
- **Source used instead: the official `kmod-ikconfig` for this exact kernel build** (`CONFIG_IKCONFIG=m`).
  - Address: `https://downloads.openwrt.org/releases/25.12.6/targets/bcm27xx/bcm2711/kmods/6.12.112-1-46807ee337a11df7be7d383dc0872b8e/kmod-ikconfig-6.12.112-r1.apk`,
    fetched 2026-11-11, 50,776 bytes, checked against the signed `sha256sums` (`sha256sum -c`, OK, exit 0).
  - The kmods folder hash `46807ee3…2b8e` equals the hash in the installed `kernel` package's version.
  - The apk (ADB, magic `ADBd`, raw deflate) was inflated with a short script run with `python3 -I`, kept outside
    the data folder; the gzip blob between `IKCFG_ST` and `IKCFG_ED` was decompressed (exit 0).
  - Result: 8,366 lines, `# Linux/arm64 6.12.112 Kernel Configuration`,
    `CONFIG_CC_VERSION_TEXT="aarch64-openwrt-linux-musl-gcc (OpenWrt GCC 14.3.0 r33311-981c978eff) 14.3.0"`;
    1,399 options `=y`, 1,278 `=m`. SHA-256 of the extracted config:
    `ccc2335f7f05c9e7bb8354caca87ec3b658bf5611764a1c9a353e6c10c897cb5`.
  - **Source of every option below: this config** (`grep -E "^CONFIG_X=|^# CONFIG_X is not set"`). "In image" says
    whether the module is in `lib/modules/6.12.112/` or in `modules.builtin` (192 lines). "Package" says whether a
    kmod is published in the kmods folder (file names in the signed `sha256sums`; 1,188 kmod packages).

| Feature (5.2) | Option(s) | Value | In image | Package |
| :--- | :--- | :--- | :--- | :--- |
| Bridge, VLAN filtering | `BRIDGE`, `BRIDGE_VLAN_FILTERING` | y, y | built in (`bridge.ko` in `modules.builtin`) | — |
| Bonding | `BONDING` | m | no | `kmod-bonding-6.12.112-r1` |
| Team | `NET_TEAM` | m | no | `kmod-team` and its modes |
| VLAN 802.1Q | `VLAN_8021Q` | y | built in (`8021q.ko`) | — |
| nftables | `NF_TABLES` m, `NF_TABLES_INET` y, `NF_TABLES_NETDEV` y, `NF_TABLES_BRIDGE` m, `NFT_NAT` m, `NFT_MASQ` m, `NFT_FIB_INET` m, `NF_FLOW_TABLE` m, `NFT_COMPAT` m, `NFT_SOCKET` m, `NFT_TPROXY` m | as listed | yes (29 `nf_*`/`nft_*` modules, `nf_tables` included) | `kmod-nft-*` (e.g. `kmod-nft-bridge-6.12.112-r1`) |
| WireGuard | `WIREGUARD` | m | no | `kmod-wireguard-6.12.112-r1` |
| PPP, PPPoE | `PPP` m, `PPPOE` m, `PPP_ASYNC` m, `L2TP` m | m | yes (`ppp_generic`, `pppoe`, `pppox`, `ppp_async`, `slhc`) | `kmod-ppp`, `kmod-pppoe` |
| cfg80211, mac80211 | `CFG80211` | **not set in the kernel** | `cfg80211.ko` and `compat.ko` from backports (`kmod-cfg80211 6.12.112.6.18.55-r2` installed); **no `mac80211`** | `kmod-mac80211-6.12.112.6.18.55-r2` |
| mac80211_hwsim | — | from backports | no | `kmod-mac80211-hwsim-6.12.112.6.18.55-r2` |
| Pi Wi-Fi (`brcmfmac`) | — | from backports | yes, see 5a | `kmod-brcmfmac 6.12.112.6.18.55-r2` |
| Pi Ethernet (GENET) | `BCMGENET`, `BROADCOM_PHY`, `MDIO_BCM_UNIMAC` (`mdio-bcm-unimac.ko`) | y | built in, see 5b | — |
| USB host | `USB` y, `PCIE_BRCMSTB` y, `USB_XHCI_HCD` y, `USB_XHCI_PCI` y, `USB_DWCOTG` y, `USB_DWC2` m, `USB_STORAGE` y, `USB_UAS` y | as listed | built in (`pcie-brcmstb`, `xhci-pci`, `dwc_otg`) | — |
| USB network | `USB_USBNET` m, `USB_LAN78XX` m, `USB_RTL8152` m, `USB_NET_AX88179_178A` m, `USB_NET_AX8817X` m, `USB_NET_SMSC95XX` m, `USB_NET_AQC111` m, `USB_NET_CDC_EEM` m | m | `usbnet`, `lan78xx`, `r8152` (and `r8169` for PCIe) | `kmod-usb-net-asix-ax88179`, `-smsc95xx`, `-aqc111` |
| USB tethering (4.27) | `USB_NET_RNDIS_HOST`, `USB_NET_CDCETHER`, `USB_NET_CDC_NCM`, `USB_IPHETH`, `USB_NET_QMI_WWAN`, `USB_NET_CDC_MBIM`, `USB_NET_HUAWEI_CDC_NCM`; `USB_WDM` m, `USB_ACM` m, `USB_SERIAL_OPTION` m | all m | no | `kmod-usb-net-rndis`, `-cdc-ether`, `-cdc-ncm`, `-ipheth`, `-qmi-wwan`, `-cdc-mbim`, `-huawei-cdc-ncm` (all `6.12.112-r1`) |
| TUN, veth, macvlan, dummy, IFB, ipvlan, VXLAN | `TUN`, `VETH`, `MACVLAN`, `DUMMY`, `IFB`, `IPVLAN`, `VXLAN` | all m | no | `kmod-tun`, `kmod-veth`, `kmod-macvlan`, `kmod-dummy`, `kmod-ifb` |
| Policy routing | `IP_ADVANCED_ROUTER`, `IP_MULTIPLE_TABLES`, `IPV6_MULTIPLE_TABLES` | y | built in | — |
| Traffic control | `NET_SCHED` y, `NET_SCH_FQ_CODEL` y, `NET_SCH_CAKE` m, `NET_SCH_HTB` m, `NET_SCH_NETEM` m, `NET_SCH_TBF` m, `NET_SCH_INGRESS` m, `NET_CLS_BPF` m, `NET_ACT_MIRRED` m | as listed | fq_codel built in | `kmod-sched-cake`, `kmod-netem` |
| Namespaces | `NAMESPACES`, `NET_NS`, `PID_NS`, `USER_NS`, `UTS_NS`, `IPC_NS` | all y | built in | — |
| seccomp | `SECCOMP`, `SECCOMP_FILTER` | y, y | built in | (`procd-seccomp` installed) |
| **Landlock** | `SECURITY` | **not set**, so `SECURITY_LANDLOCK` is absent | — | not available without an own kernel |
| cgroups | `CGROUPS` y, `MEMCG` y (`MEMCG_V1` not set), `CGROUP_PIDS` y, `CPUSETS` y, `CGROUP_BPF` y, `BLK_CGROUP` y; `CGROUP_FREEZER`, `CGROUP_DEVICE` not set | as listed | the DT `bootargs` has `cgroup_disable=memory` (section 2); what is mounted: not determined, needs a boot | — |
| Watchdog | `WATCHDOG` y, `BCM2835_WDT` y, `SOFT_WATCHDOG` m | as listed | `bcm2835_wdt` built in; DT node `watchdog@7e100000`. Whether procd opens `/dev/watchdog`: not determined, needs a boot | `kmod-softdog-6.12.112-r1` |
| Hardware RNG | `HW_RANDOM_IPROC_RNG200` y (`HW_RANDOM_BCM2835` not set) | y | built in (`iproc-rng200`); `urngd` is **not** installed (it is in I-01) | — |
| Bluetooth | `BT` m, `BT_HCIUART` m | m | no; `distroconfig.txt` disables BT with `disable-bt` | `kmod-bluetooth-6.12.112-r1` |
| Module signing | `MODULE_SIG` | not set | — | — |
| Other | `BPF_SYSCALL` y, `XDP_SOCKETS` not set, `TCP_CONG_BBR` m, `DEFAULT_TCP_CONG="cubic"`, `IPV6` y, `PREEMPT_NONE` y, `HZ=100`, `NR_CPUS=4`, `ARM64_4K_PAGES` y, `ARM64_VA_BITS=39`, `EFI` not set, `VIRTUALIZATION` not set, `ARM_RASPBERRYPI_CPUFREQ` y with `CPU_FREQ_DEFAULT_GOV_ONDEMAND` y, `BCM2711_THERMAL` y, `SQUASHFS` y, `OVERLAY_FS` y, `EXT4_FS` y, `F2FS_FS` y, `VFAT_FS` m | as listed | — | — |

## 5. Modules and firmware files carried

- **Modules.** 77 `.ko` files in `lib/modules/6.12.112/` (`ls *.ko | wc -l`; 3,624,882 bytes with
  `modules.builtin` and `modules.builtin.modinfo`), all `vermagic 6.12.112 SMP mod_unload aarch64`. List:
  `brcmfmac`, `brcmfmac-bca`, `brcmfmac-cyw`, `brcmfmac-wcc`, `brcmutil`, `cfg80211`, `compat`, `crc-ccitt`, `evdev`,
  `fat`, `vfat`, `hid`, `hid-generic`, `usbhid`, `hwmon`, `i2c-bcm2835`, `i2c-brcmstb`, `i2c-dev`, `lan78xx`, `r8152`,
  `r8169`, `realtek`, `microchip`, `usbnet`, `libcrc32c`, `nf_conntrack`, `nf_defrag_ipv4`, `nf_defrag_ipv6`,
  `nf_flow_table`, `nf_flow_table_inet`, `nf_log_syslog`, `nf_nat`, `nf_reject_ipv4`, `nf_reject_ipv6`, `nf_tables`,
  `nfnetlink`, `nft_chain_nat`, `nft_ct`, `nft_fib`, `nft_fib_inet`, `nft_fib_ipv4`, `nft_fib_ipv6`,
  `nft_flow_offload`, `nft_hash`, `nft_limit`, `nft_log`, `nft_masq`, `nft_nat`, `nft_numgen`, `nft_quota`,
  `nft_redir`, `nft_reject`, `nft_reject_inet`, `nft_reject_ipv4`, `nft_reject_ipv6`, `nls_cp437`, `nls_iso8859-1`,
  `nls_utf8`, `ppp_async`, `ppp_generic`, `pppoe`, `pppox`, `slhc`, `snd`, `snd-bcm2835`, `snd-compress`,
  `snd-hwdep`, `snd-mixer-oss`, `snd-pcm`, `snd-pcm-dmaengine`, `snd-pcm-oss`, `snd-rawmidi`, `snd-seq-device`,
  `snd-timer`, `soundcore`, `spi-bcm2835`, `spi-bcm2835aux`.
  - Load order: `/etc/modules-boot.d` (11 entries) and `/etc/modules.d` (49 entries, e.g. `brcmutil`, `brcmfmac`,
    `usb-net-lan78xx`, `usb-net-rtl8152`, `ppp`, `pppoe`, `nft-*`).
  - Built in (`modules.builtin`, 192 lines), Pi-relevant: `genet`, `mdio-bcm-unimac`, `broadcom`, `bcm7xxx`,
    `bcm-phy-lib`, `pcie-brcmstb`, `xhci-hcd`, `xhci-pci`, `xhci-plat-hcd`, `dwc_otg`, `sdhci-iproc`, `bcm2835-mmc`,
    `mmc_block`, `bcm2835_wdt`, `iproc-rng200`, `bcm2711_thermal`, `bcm2835-mailbox`, `vchiq`, `vc_mem`,
    `usb-storage`, `uas`, `bridge`, `8021q`.
- **Firmware.** 49 files and 14 symlinks under `lib/firmware`, 1,544,121 bytes (`find -type f`, `find -type l`,
  `du -sb`):
  - `brcm/` and `cypress/`: Wi-Fi, see 5a.
  - `rtl_nic/` (39 files): `rtl8105e-1`, `rtl8106e-1/-2`, `rtl8107e-1/-2`, `rtl8125a-3`, `rtl8125b-1/-2`,
    `rtl8125bp-2`, `rtl8125cp-1`, `rtl8125d-1/-2`, `rtl8125k-1`, `rtl8126a-2/-3`, `rtl8127a-1`, `rtl8153a-2/-3/-4`,
    `rtl8153b-2`, `rtl8153c-1`, `rtl8156a-2`, `rtl8156b-2`, `rtl8168d-1/-2`, `rtl8168e-1/-2/-3`, `rtl8168f-1/-2`,
    `rtl8168fp-3`, `rtl8168g-1/-2/-3`, `rtl8168h-1/-2`, `rtl8402-1`, `rtl8411-1/-2` (`.fw`), from `r8152-firmware`
    and `r8169-firmware` (both `20260221-r1`).
  - `regulatory.db` (6,332 bytes, `c682b8e0f2f61385929586409f092638ba84dd795503d220be15b5c3d16dd75d`) from
    `wireless-regdb 2026.09.03-r1`.
  - Six firmware packages have an empty licence field in the apk database: `bcm27xx-gpu-fw`, `brcmfmac-firmware-usb`,
    `brcmfmac-nvram-43455-sdio`, `cypress-firmware-43455-sdio`, `r8152-firmware`, `r8169-firmware`.

### 5a. Pi-specific Wi-Fi (CYW43455 over SDIO, `brcmfmac`)

- **Device tree** (`bcm2711-rpi-4-b.dtb` decompiled with `dtc`): `mmcnr@7e300000` (`brcm,bcm2835-mmc`), status
  `okay`, bus width 4, with child `wifi@1` `compatible = "brcm,bcm4329-fmac"`; the `__symbols__` alias is
  `brcmf = "/soc/mmcnr@7e300000/wifi@1"`.
- **Driver.** `brcmfmac.ko` with `brcmutil.ko`, plus the vendor parts `brcmfmac-cyw`, `brcmfmac-bca`, `brcmfmac-wcc`,
  from `kmod-brcmfmac 6.12.112.6.18.55-r2` / `kmod-brcmutil 6.12.112.6.18.55-r2` (backports 6.18.55).
  `modinfo -F depends brcmfmac.ko`: `brcmutil,cfg80211,compat`; licence "Dual BSD/GPL". It is a FullMAC driver on
  cfg80211 only: `mac80211` is not installed (package list). `strings brcmfmac.ko` shows the name pattern
  `brcm/brcmfmac43455-sdio` with `.txt` and `.clm_blob`.
- **Firmware and NVRAM files** (`find … -print0`, `stat`, `sha256sum`):

| Path under `lib/firmware` | Kind | Size (bytes) | SHA-256 / target | Package |
| :--- | :--- | ---: | :--- | :--- |
| `cypress/cyfmac43455-sdio.bin` | firmware | 612,273 | `0c618567dfbad2adbece68cf339611aa92f7305f95d7c9a4b655e152ffc4939b` | `cypress-firmware-43455-sdio 5.10.9.2022.0909-r1` |
| `cypress/cyfmac43455-sdio.clm_blob` | regulatory (CLM) blob | 4,733 | `e0bd46f3ae0b9ec988967dae860a490041395379fd6aab3e11bdfefe1f827bfc` | same |
| `brcm/brcmfmac43455-sdio.bin` | link | — | `../cypress/cyfmac43455-sdio.bin` | same |
| `brcm/brcmfmac43455-sdio.clm_blob` | link | — | `../cypress/cyfmac43455-sdio.clm_blob` | same |
| `brcm/brcmfmac43455-sdio.raspberrypi,4-model-b.txt` | NVRAM | 1,883 | `edb6f4e4fb19e18940004124feb4ffe160d72fc607243a07a4480338a28b2748` | `brcmfmac-nvram-43455-sdio 20260221-r1` |
| `brcm/brcmfmac43455-sdio.raspberrypi,4-compute-module.txt`, `…5-model-b.txt`, `…5-compute-module.txt` | links | — | `brcmfmac43455-sdio.raspberrypi,4-model-b.txt` | same |
| `brcm/brcmfmac43455-sdio.Raspberry Pi Foundation-Raspberry Pi 4 Model B.txt` (and `…Pi 5 Model B`, `…Compute Module 4`, `…Compute Module 5`) | links | — | `brcmfmac43455-sdio.raspberrypi,4-model-b.txt` | same |
| `brcm/brcmfmac43455-sdio.raspberrypi,3-model-b-plus.txt` | NVRAM | 1,884 | `ae2f9e43c609766036cd7d82d709e62135fc266cbe8f40fb10f782114d585658` | same (`3-model-a-plus` links to it) |
| `brcm/brcmfmac43455-sdio.AW-CM256SM.txt` | NVRAM | 1,475 | `f50362207b7536521cadaa59bdc294a9ad946b149fbb6e332ec2a245c1f34fbf` | same (four Pine64/Beagle links to it) |
| `brcm/brcmfmac43455-sdio.acepc-t8.txt` | NVRAM | 1,723 | `b87e0dc904dbc660a1c467bf0e3ca1cf8babe7e9ec0559041758808646e4f3e3` | same |
| `brcm/brcmfmac43455-sdio.MINIX-NEO Z83-4.txt` | NVRAM | 2,510 | `030aa76f3554075df9a576ea46b7d0a00b46dc8070c9e969b7d813cd0dac3981` | same |
| `brcm/brcmfmac43143.bin` | USB Wi-Fi firmware | 397,312 | `a0ace23094652c3aa3bcd480906b1a769d0169afcc3b5f074b187a30d8617e03` | `brcmfmac-firmware-usb 20260221-r1` |
| `brcm/brcmfmac43236b.bin` | USB Wi-Fi firmware | 348,160 | `2138a2edca76a1f808250fd89381b30529938e04b2a9befa138afda71a9c96e4` | same |

- **Firmware version** (`strings cyfmac43455-sdio.bin`): `43455c0-roml/43455_sdio-pno-aoe-pktfilter-pktctx-wfds-mfp-
  dfsradar-wowlpf-noclminc-clm_min-obss-obssdump-swdiv-gtkoe-roamprof-txbf-ve-extsae-dpp-sr-okc-bpd Version: 7.45.250
  (70e9766 CY) … Date: Fri 2022-07-29 … FWID 01-e53e306b`. `brcmfmac43143.bin`: `Version: 6.10.198.66`, 2014-04-03.
  `brcmfmac43236b.bin`: no version string found by `strings`.
- **Pi 4 NVRAM file**: header "NVRAM config file for the BCM43455 WiFi/BT chip as found on the Raspberry Pi 4 Model B",
  `SPDX-License-Identifier: GPL-2.0+`, `(C) Copyright 2018 Raspberry Pi (Trading) Ltd.`; `devid=0x43ab`,
  `boardtype=0x6e4`, `boardrev=0x1304`, `aa2g=1`, `aa5g=1`, `txchain=1`, `rxchain=1`. It also contains a fixed
  `macaddr=` line (the value is not recorded here). Which address the interface really gets: not determined, needs a
  boot on the Pi 4, later step.
- **Which NVRAM name the driver picks** (board-compatible name or the "Raspberry Pi Foundation-…" name): not
  determined, needs a boot on the Pi 4, later step. Both names lead to the same file.
- **User space for Wi-Fi.** `wpad-basic-mbedtls 2025.08.26~ca266cc2-r6` (`hostapd` and `wpa_supplicant` are links to
  `/usr/sbin/wpad`), `hostapd-common`, `iw 6.17-r1`, `iwinfo`, `wifi-scripts 1.0-r1`, `ucode-mod-nl80211`,
  `wireless-regdb`. The init script `wpad` is enabled (`S19wpad`).
- **Default wireless configuration** (read in `/lib/wifi/mac80211.uc`, not run): for each detected radio it writes a
  `wifi-device` of type `mac80211` and an AP `wifi-iface` with SSID "OpenWrt", `encryption 'none'` and
  `disabled='1'` unless `board.json` carries `wlan.defaults`. `/etc/board.d/02_network` sets no wlan defaults for the
  Pi 4 (only `ucidef_set_interface_lan "eth0"`), so the AP would be generated **disabled**, with an empty country.
  Not seen on a running system: not determined, needs a boot on the Pi 4, later step.
- **Bluetooth** on the same chip: switched off by `dtoverlay=disable-bt` in `distroconfig.txt`; no BT firmware
  (`BCM4345C0.hcd`) and no BT modules are in the image.

### 5b. Pi-specific Ethernet (GENET)

- **Device tree:** `ethernet@7d580000`, `compatible = "brcm,bcm2711-genet-v5"`, status `okay`, `phy-mode =
  "rgmii-rxid"`, with `mdio@e14` (`brcm,genet-mdio-v5`) and `ethernet-phy@1`; alias `ethernet0`. The DT also offers a
  parameter `eee` that adds `genet.eee=N` to the bootargs.
- **Driver:** built into the kernel (`BCMGENET=y`; `genet.ko`, `mdio-bcm-unimac.ko`, `broadcom.ko`, `bcm7xxx.ko`,
  `bcm-phy-lib.ko` in `modules.builtin`). It needs no firmware file. `kmod-phy-microchip` and `kmod-phy-realtek` are
  installed for USB/PCIe adapters.
- **Name and role:** `/etc/board.d/02_network` and `/lib/preinit/05_set_preinit_iface_brcm2708` make `eth0` the LAN
  (and the failsafe interface) for `raspberrypi,4-model-b`, `raspberrypi,400` and `raspberrypi,4-compute-module`.
  There is **no WAN** in the default configuration of this board (no `ucidef_set_interface_wan`).
- **Link speed, the PHY seen and the MAC address:** not determined, needs a boot on the Pi 4, later step.
- **USB 3 controller:** the VL805 is on PCIe (`pcie@7d500000`, `usb@0,0` child); `pcie-brcmstb` and `xhci-pci` are
  built in. The USB-C port uses `dwc_otg` (built in) or `dwc2` (overlay).

## 6. Init system, service manager, C library, package manager

- **Init:** procd (`/sbin/procd`, 97,697 bytes; `/sbin/init` 16,649 bytes) with ubus, ujail and seccomp
  (`procd 2026.03.13~58eb263d-r1`, `procd-ujail`, `procd-seccomp`). `/etc/inittab`: `::sysinit:/etc/init.d/rcS S boot`,
  `::askconsole:/usr/libexec/login.sh` and `tty1::askfirst`. Init scripts are in `/etc/init.d` (22) and linked in
  `/etc/rc.d`.
- **Enabled services** (`ls etc/rc.d`, `S` links): sysfixtime, boot, system, sysctl, log, rpcd, dnsmasq, dropbear,
  firewall, wpad, network, packet_steering, odhcpd, cron, uhttpd, ucitrack, gpio_switch, done, led, sysntpd,
  urandom_seed (21 `S` links; `umount` has only a `K` link). Against I-01: `wpad` is added, `urngd` is absent.
- **Services running after boot:** not determined, needs a boot on the Pi 4, later step.
- **C library: musl 1.2.5.** `lib/libc.so` (`file`: "ELF 64-bit LSB shared object, ARM aarch64 … static-pie linked");
  `strings lib/libc.so` contains `1.2.5`; package `libc 1.2.5-r5`; `lib/ld-musl-aarch64.so.1` links to `libc.so`.
  Toolchain GCC 14.3.0 and GNU ld 2.44, from the kernel's version string.
- **Package manager: apk.** `/usr/bin/apk`, package `apk-mbedtls 3.0.5-r3` (its `--version` was not run, no binary
  from the image is executed here). opkg is absent. Feeds in `/etc/apk/repositories.d/distfeeds.list`;
  `customfeeds.list` is present.
- **Installed packages: 177.** The squashfs and ext4 images give the same list (`awk '/^P:/…/^V:/…'
  lib/apk/db/installed`, then `diff`, exit 0). The publisher's target manifest
  (`openwrt-25.12.6-bcm27xx-bcm2711.manifest`, 142 lines, `sha256sum -c` OK) is a strict subset; the 35 more are the
  `rpi-4` profile's device packages (`profiles.json`: `cypress-firmware-43455-sdio`, `brcmfmac-nvram-43455-sdio`,
  `kmod-brcmfmac`, `wpad-basic-mbedtls`, `kmod-i2c-bcm2835`, `kmod-spi-bcm2835`, `kmod-spi-bcm2835-aux`,
  `kmod-i2c-brcmstb`, `kmod-usb-net-lan78xx`, `kmod-usb-net-rtl8152`, `kmod-r8169`) and their dependencies. The list
  is in section 10.
- **Licences of the installed packages** (`L:` fields, `sort | uniq -c`): GPL-2.0-only 64, ISC 33, GPL-2.0 28,
  Apache-2.0 15, LGPL-2.1 7, MIT 4, GPL-3.0-with-GCC-exception 4, GPL-2.0-or-later 4, BSD-3-Clause 4,
  LGPL-2.1-or-later 2, BSD-4-Clause 2, Zlib, LGPL-2.1+, "GPL-2.0-or-later MPL-2.0", GPL-2.0+ (1 each), and 6 empty
  (the firmware packages of section 5).

## 7. Tools of 6A criterion 3, with versions

The presence check is a loop over `bin sbin usr/bin usr/sbin` in the unpacked root (`[ -e ] || [ -L ]`, links read
with `readlink`). Versions come from the installed package and, where it has one, a version string found by
`strings` in the binary. **No binary was run**, so `--version` outputs are not available.

| Tool | Present | Version |
| :--- | :--- | :--- |
| procd (init) | yes, `/sbin/procd` | package `procd 2026.03.13~58eb263d-r1` |
| systemd, OpenRC | no | — |
| nftables `nft` | yes, `/usr/sbin/nft` | strings `1.1.6`, `Commodore Bullmoose #7`; package `nftables-json 1.1.6-r2` |
| dnsmasq | yes | strings `2.93`; package `dnsmasq 2.93-r1` (not `dnsmasq-full`) |
| hostapd | **yes**, `/usr/sbin/hostapd -> wpad` | `wpad-basic-mbedtls 2025.08.26~ca266cc2-r6`; strings `wpa_supplicant v2.12-devel` |
| wpa_supplicant | **yes**, `/usr/sbin/wpa_supplicant -> wpad` | same |
| iw | **yes**, `/usr/sbin/iw` | strings `6.17`; package `iw 6.17-r1` |
| pppd | yes, `/usr/sbin/pppd` | plugin path `/usr/lib/pppd/2.5.2`; `ppp 2.5.2-r4`, `ppp-mod-pppoe 2.5.2-r4` |
| WireGuard (`wg`, module) | **no** | (feeds: section 9a) |
| OpenVPN | **no** | (feeds: section 9a) |
| `ip` | yes, but the **BusyBox applet** (`/sbin/ip -> ../bin/busybox`) | BusyBox v1.37.0 |
| `tc`, `ethtool`, `iptables`, `ss` | no | — |
| `brctl` | BusyBox applet | — |
| DHCP client | `udhcpc` (BusyBox), `odhcp6c 2026.06.27~10a52220-r3` | — |
| DHCP/RA server | `odhcpd-ipv6only 2026.06.29~5d7be43f-r1` and dnsmasq | — |
| SSH | `dropbear` | strings `2025.89`; package `dropbear 2025.89-r3` |
| Web | `uhttpd 2026.08.24~373145f7-r1` with LuCI `26.288.56989~75a306d` | — |
| HTTP client | `uclient-fetch` (`/usr/bin/wget -> /bin/uclient-fetch`) | `uclient-fetch 2026.06.07~daad21fa-r1` |
| NTP | `ntpd` (BusyBox applet, service `sysntpd`) | — |
| BusyBox | yes | `busybox 1.37.0-r6`; strings `BusyBox v1.37.0 (2026-10-17 21:45:38 UTC)` |
| Pi tools | `vcgencmd`, `vcmailbox`, `vclog`, `dtoverlay`, `dtparam`, `dtmerge`, `pinctrl`, `raspinfo`, `eepdump`, `eepmake`, `eepflash.sh`, `kdtc`, `otpset`, `overlaycheck`, `ovmerge` | package `bcm27xx-utils 2025.03.14-r1`; see section 8 |

## 8. Python (6A criterion 8)

**No Python interpreter is in the image, but one Python script is.** The checks, in both unpacked roots:

- `find . -iname "python*" -o -name "*.py" -o -name "*.pyc"` found 0 files (squashfs and ext4).
- `grep -rlI "^#!.*python" .` found **1 file**: `/usr/bin/otpset` (5,934 bytes, `#!/usr/bin/env python3`).
- `grep -rlI "^#!.*perl" .` found **3 files**: `/usr/bin/overlaycheck`, `/usr/bin/ovmerge`, `/usr/bin/kdtc`.
- All four belong to `bcm27xx-utils 2025.03.14-r1` (its `F:`/`R:` entries in `lib/apk/db/installed`).
- `python`, `python3` and `perl` are absent from `bin sbin usr/bin usr/sbin`, so these four scripts cannot run on the
  image as built.
- The base's own tools are shell, ucode and C (`fw4` is ucode; netifd and procd are C).
- `python3` / `python3-base 3.13.18-r1` are in the packages feed only (section 9a), not installed.
- **For criterion 8:** an NMB image on this base would have to leave out `bcm27xx-utils` or at least these scripts;
  the release check for a Python interpreter alone would not find `otpset`.

## 9. Own network and firewall manager, and how each is switched off (6A criterion 5)

The image carries **netifd** (`/sbin/netifd`, `netifd 2026.02.26~cbb83a18-r2`) and **firewall4** (`/sbin/fw4`, ucode,
`firewall4 2025.03.17~b6e51575-r2`), both enabled (`S20network`, `S19firewall`).

- **Same scripts as I-01.** `cmp` against the I-01 x86-64 root: `sbin/fw4`, `etc/init.d/firewall`,
  `etc/init.d/network`, `usr/share/firewall4/main.uc`, `bin/config_generate`, `sbin/sysupgrade` and
  `lib/upgrade/common.sh` are byte-identical. `etc/init.d/firewall` has `stop_service() { fw4 flush }`, and `fw4 flush`
  deletes every table (`nft list tables | while read …; nft delete table`, `sbin/fw4` line 58 on), so I-01's caution
  holds as written: the firewall service's `stop` would delete NMB's tables too; plain `fw4 stop` deletes only
  `inet fw4`.
- **How they are switched off:** `/etc/init.d/firewall disable` and `/etc/init.d/network disable` remove the `rc.d`
  links (rc.common, as in I-01). On the Pi 4 itself the coexistence tests of I-01 (own nft table across a firewall
  restart, own bridge across netifd reload/restart, own address on `br-lan`): not determined, needs a boot on the
  Pi 4, later step. netifd and fw4 are architecture-independent C/ucode of the same versions, so I-01's results are
  the expectation, not a measurement on this image.
- **Default configuration** (read from `/etc/board.d/02_network`, `/etc/board.d/99-default_network` and
  `bin/config_generate`, not run): `eth0` in `br-lan` with the built-in static LAN default address, dnsmasq serving
  DHCP and DNS on the LAN, a ULA prefix generated on the first boot (`12_network-generate-ula`). **No WAN interface**
  for the Pi 4. Wi-Fi AP generated disabled (5a).
- **Other defaults relevant for security** (read in `/etc`): `root` has an empty password field in `/etc/shadow`;
  `/etc/config/dropbear` has `PasswordAuth 'on'`, `RootPasswordAuth 'on'`, port 22; `/etc/config/uhttpd` listens on
  `0.0.0.0:80`, `[::]:80`, `0.0.0.0:443`, `[::]:443`; `/etc/sysctl.d/10-default.conf` sets `net.ipv4.ip_forward=1` and
  `net.ipv6.conf.default.forwarding=1`.

### 9a. Feed lookup (no boot; read on the PC)

The feeds could not be read with `apk search` in a guest, as in I-01. Instead the publisher's `index.json` of the
feeds `base`, `packages` and `routing` for `aarch64_cortex-a72` and of the target `packages` folder were fetched on
2026-11-11 at about 07:19 UTC and read with a short `python3 -I` script kept outside the data folder (742, 4,565, 56
and 89 packages). `index.json` is **not signed** (the signed index is `packages.adb`, also fetched, not parsed); kmod
versions come from the signed target `sha256sums` instead. **None of these packages is installed in the image**
unless marked.

| Package | Version in the 25.12.6 feeds |
| :--- | :--- |
| `hostapd-mbedtls` / `wpad-mbedtls` / `wpad-basic-mbedtls` / `wpa-supplicant-mbedtls` | `2025.08.26~ca266cc2-r7` (the image has `wpad-basic-mbedtls …-r6` installed) |
| `iw` / `iw-full` | `6.17-r1` (`iw` installed) |
| `wireguard-tools` | `1.0.20260223-r1` |
| `kmod-wireguard` | `6.12.112-r1` (signed list) |
| `openvpn-mbedtls` | `2.7.8-r1` |
| `ip-full` / `tc-full` (iproute2) | `6.18.0-r2` |
| `ethtool` | `6.15-r1` |
| `dnsmasq-full` | `2.93-r2` |
| `iperf3` | `3.22-r1` |
| `kmod-mac80211`, `kmod-mac80211-hwsim` | `6.12.112.6.18.55-r2` (signed list) |
| `kmod-veth`, `kmod-bonding`, `kmod-usb-net-rndis`, `kmod-softdog`, `kmod-bluetooth`, `kmod-ikconfig` | `6.12.112-r1` (signed list) |
| `bluez-daemon` | `5.83-r1` |
| `python3` / `python3-base` | `3.13.18-r1` (in the feed only; never installed for NMB, criterion 8) |

## 10. Installed packages (177, `lib/apk/db/installed`; the same in both images)

`apk-mbedtls 3.0.5-r3`, `attendedsysupgrade-common 10`, `base-files 1718~981c978eff`, `bcm27xx-gpu-fw 2025.04.30-r1`, `bcm27xx-utils 2025.03.14-r1`, `brcmfmac-firmware-usb 20260221-r1`, `brcmfmac-nvram-43455-sdio 20260221-r1`, `busybox 1.37.0-r6`, `ca-bundle 20260816-r1`, `cgi-io 2026.07.21~31cb3c89-r1`, `cypress-firmware-43455-sdio 5.10.9.2022.0909-r1`, `dnsmasq 2.93-r1`, `dropbear 2025.89-r3`, `e2fsprogs 1.47.3-r1`, `firewall4 2025.03.17~b6e51575-r2`, `fstools 2026.05.23~16718b6e-r1`, `fwtool 2025.10.03~04cd252e-r1`, `getrandom 2025.10.30~6f78fa49-r1`, `hostapd-common 2025.08.26~ca266cc2-r6`, `iw 6.17-r1`, `jansson4 2.14.1-r1`, `jshn 2026.06.19~7dd12784-r1`, `jsonfilter 2026.03.16~b9034210-r1`, `kernel 6.12.112~46807ee337a11df7be7d383dc0872b8e-r1`, `kmod-brcmfmac 6.12.112.6.18.55-r2`, `kmod-brcmutil 6.12.112.6.18.55-r2`, `kmod-cfg80211 6.12.112.6.18.55-r2`, `kmod-crypto-crc32c 6.12.112-r1`, `kmod-crypto-hash 6.12.112-r1`, `kmod-crypto-sha256 6.12.112-r1`, `kmod-fixed-phy 6.12.112-r1`, `kmod-fs-vfat 6.12.112-r1`, `kmod-hid 6.12.112-r1`, `kmod-hid-generic 6.12.112-r1`, `kmod-hwmon-core 6.12.112-r1`, `kmod-i2c-bcm2835 6.12.112-r1`, `kmod-i2c-brcmstb 6.12.112-r1`, `kmod-i2c-core 6.12.112-r1`, `kmod-input-core 6.12.112-r1`, `kmod-input-evdev 6.12.112-r1`, `kmod-lib-crc-ccitt 6.12.112-r1`, `kmod-lib-crc32c 6.12.112-r1`, `kmod-libphy 6.12.112-r1`, `kmod-mdio-devres 6.12.112-r1`, `kmod-mii 6.12.112-r1`, `kmod-mmc 6.12.112-r1`, `kmod-nf-conntrack 6.12.112-r1`, `kmod-nf-conntrack6 6.12.112-r1`, `kmod-nf-flow 6.12.112-r1`, `kmod-nf-log 6.12.112-r1`, `kmod-nf-log6 6.12.112-r1`, `kmod-nf-nat 6.12.112-r1`, `kmod-nf-reject 6.12.112-r1`, `kmod-nf-reject6 6.12.112-r1`, `kmod-nfnetlink 6.12.112-r1`, `kmod-nft-core 6.12.112-r1`, `kmod-nft-fib 6.12.112-r1`, `kmod-nft-nat 6.12.112-r1`, `kmod-nft-offload 6.12.112-r1`, `kmod-nls-base 6.12.112-r1`, `kmod-nls-cp437 6.12.112-r1`, `kmod-nls-iso8859-1 6.12.112-r1`, `kmod-nls-utf8 6.12.112-r1`, `kmod-of-mdio 6.12.112-r1`, `kmod-phy-microchip 6.12.112-r1`, `kmod-phy-realtek 6.12.112-r1`, `kmod-ppp 6.12.112-r1`, `kmod-pppoe 6.12.112-r1`, `kmod-pppox 6.12.112-r1`, `kmod-r8169 6.12.112-r1`, `kmod-slhc 6.12.112-r1`, `kmod-sound-arm-bcm2835 6.12.112-r1`, `kmod-sound-core 6.12.112-r1`, `kmod-spi-bcm2835 6.12.112-r1`, `kmod-spi-bcm2835-aux 6.12.112-r1`, `kmod-usb-common 6.12.112-r1`, `kmod-usb-core 6.12.112-r1`, `kmod-usb-hid 6.12.112-r1`, `kmod-usb-net 6.12.112-r1`, `kmod-usb-net-lan78xx 6.12.112-r1`, `kmod-usb-net-rtl8152 6.12.112-r1`, `libblkid1 2.41.5-r1`, `libblobmsg-json20260213 2026.06.19~7dd12784-r1`, `libc 1.2.5-r5`, `libcomerr0 1.47.3-r1`, `libe2p2 1.47.3-r1`, `libext2fs2 1.47.3-r1`, `libf2fs6 1.16.0-r4`, `libfdt 1.7.2-r1`, `libgcc1 14.3.0-r5`, `libiwinfo-data 2026.05.26~66bdd1a0-r1`, `libiwinfo20230701 2026.05.26~66bdd1a0-r1`, `libjson-c5 0.18-r1`, `libjson-script20260213 2026.06.19~7dd12784-r1`, `liblucihttp-ucode 2023.03.15~9b5b683f-r1`, `liblucihttp0 2023.03.15~9b5b683f-r1`, `libmbedtls21 3.6.7-r1`, `libmnl0 1.0.5-r1`, `libnftnl11 1.3.1-r1`, `libnl-tiny1 2025.12.02~40493a65-r1`, `libpthread 1.2.5-r5`, `librt 1.2.5-r5`, `libsmartcols1 2.41.5-r1`, `libss2 1.47.3-r1`, `libubox20260213 2026.06.19~7dd12784-r1`, `libubus20251202 2026.06.28~24864e78-r2`, `libuci20250120 2025.12.02~66127cd7-r1`, `libuclient20201210 2026.06.07~daad21fa-r1`, `libucode20230711 2026.01.16~85922056-r2`, `libudebug 2026.01.16~875e1a7a`, `libustream-mbedtls20201210 2026.03.01~99f1c0db-r1`, `libuuid1 2.41.5-r1`, `logd 2025.10.30~6f78fa49-r1`, `luci 26.288.56989~75a306d`, `luci-app-attendedsysupgrade 26.288.56989~75a306d`, `luci-app-firewall 26.288.56989~75a306d`, `luci-app-package-manager 26.288.56989~75a306d`, `luci-base 26.288.56989~75a306d`, `luci-lib-uqr 26.288.56989~75a306d`, `luci-light 26.288.56989~75a306d`, `luci-mod-admin-full 26.288.56989~75a306d`, `luci-mod-network 26.288.56989~75a306d`, `luci-mod-status 26.288.56989~75a306d`, `luci-mod-system 26.288.56989~75a306d`, `luci-proto-ipv6 26.288.56989~75a306d`, `luci-proto-ppp 26.288.56989~75a306d`, `luci-ssl 26.288.56989~75a306d`, `luci-theme-bootstrap 26.288.56989~75a306d`, `mkf2fs 1.16.0-r4`, `mtd 27`, `netifd 2026.02.26~cbb83a18-r2`, `nftables-json 1.1.6-r2`, `odhcp6c 2026.06.27~10a52220-r3`, `odhcpd-ipv6only 2026.06.29~5d7be43f-r1`, `openwrt-keyring 2025.12.10~f0670054-r1`, `owut 2026.08.30~9b51eeec-r1`, `partx-utils 2.41.5-r1`, `ppp 2.5.2-r4`, `ppp-mod-pppoe 2.5.2-r4`, `procd 2026.03.13~58eb263d-r1`, `procd-seccomp 2026.03.13~58eb263d-r1`, `procd-ujail 2026.03.13~58eb263d-r1`, `px5g-mbedtls 11`, `r8152-firmware 20260221-r1`, `r8169-firmware 20260221-r1`, `rpcd 2026.10.04~d99f703e-r1`, `rpcd-mod-file 2026.10.04~d99f703e-r1`, `rpcd-mod-iwinfo 2026.10.04~d99f703e-r1`, `rpcd-mod-luci 20240305-r1`, `rpcd-mod-rpcsys 2026.10.04~d99f703e-r1`, `rpcd-mod-rrdns 20170710`, `rpcd-mod-ucode 2026.10.04~d99f703e-r1`, `ubox 2025.10.30~6f78fa49-r1`, `ubus 2026.06.28~24864e78-r2`, `ubusd 2026.06.28~24864e78-r2`, `uci 2025.12.02~66127cd7-r1`, `uclient-fetch 2026.06.07~daad21fa-r1`, `ucode 2026.01.16~85922056-r2`, `ucode-mod-digest 2026.01.16~85922056-r2`, `ucode-mod-fs 2026.01.16~85922056-r2`, `ucode-mod-html 1`, `ucode-mod-log 2026.01.16~85922056-r2`, `ucode-mod-math 2026.01.16~85922056-r2`, `ucode-mod-nl80211 2026.01.16~85922056-r2`, `ucode-mod-rtnl 2026.01.16~85922056-r2`, `ucode-mod-ubus 2026.01.16~85922056-r2`, `ucode-mod-uci 2026.01.16~85922056-r2`, `ucode-mod-uclient 2026.06.07~daad21fa-r1`, `ucode-mod-uloop 2026.01.16~85922056-r2`, `uhttpd 2026.08.24~373145f7-r1`, `uhttpd-mod-ubus 2026.08.24~373145f7-r1`, `urandom-seed 3`, `usign 2025.10.03~c4c72b1b-r1`, `wifi-scripts 1.0-r1`, `wireless-regdb 2026.09.03-r1`, `wpad-basic-mbedtls 2025.08.26~ca266cc2-r6`, `zlib 1.3.1-r1`

## 11. Size on disk

- Download: 15,432,229 bytes (squashfs) and 17,006,660 bytes (ext4), as `.img.gz`.
- Disk image: 82,342,912 bytes (78.5 MiB, squashfs) and 184,549,376 bytes (176 MiB, ext4). Both declare a card layout
  ending at sector 360,447 (176 MiB); the rest of a card is not used by the layout.
- Boot partition: 64 MiB, 20,268,541 bytes in files (kernel 12,886,024; VideoCore firmware 6,106,835 in the six
  `start4*`/`fixup4*` files; overlays 1,103,079), 46,303,232 bytes free (`mdir`).
- squashfs root: 5,796,528 bytes compressed. Unpacked, the root is 17,645,429 bytes (`du -sb`), the same for the
  ext4 root. Of that:
  - `lib/modules`: 3,624,882 bytes
  - `lib/firmware`: 1,544,121 bytes
  - `usr/lib`: 3,998,904 bytes
  - `www` (LuCI): 1,337,317 bytes
- ext4 image: the 104 MiB root has 20,018 of 26,624 4-KiB blocks free (81,993,728 bytes, about 78.2 MiB), from
  `debugfs -R stats`.
- Overlay size and free space after the first boot (squashfs image): not determined, needs a boot on the Pi 4, later
  step.

## 12. Boot test

**Not run, by the package's rule.** The Pi 4 image is not booted in QEMU on this PC and is not written to a card here.
The boot on the owner's Pi 4 is a later step under work plan section 7 point 5 (full backup of settings, database and
the card's image first, a tried way back, the time recorded as the box's own action, the Python version put back
afterwards).

Fields 5.2 takes from a boot, all **not determined: needs a boot on the Pi 4, later step**:

- boot time (firmware to `procd: - init complete -`)
- free memory after boot (`free`, kernel `Memory:` line) on the Pi 4's RAM size
- free storage after boot (`df -k`: `/rom`, `/overlay`, `/boot`, `/tmp`)
- listening ports (`netstat -tulnp`); expected from the configuration only: dropbear 22, uhttpd 80/443, dnsmasq 53,
  odhcpd 547
- running services (`ps w`, `ubus call service list`)
- `/proc/cmdline` (merge of `cmdline.txt` and the DT bootargs), mounted cgroup controllers, `/dev/watchdog*` and
  `ubus call system watchdog`
- the netifd and fw4 coexistence tests of I-01 on this hardware
- Wi-Fi: the NVRAM file chosen, firmware load, `iw list` (bands, interface modes, AP/station/monitor support,
  concurrent interfaces), regulatory domain
- Ethernet: GENET link speed, PHY, MAC address, EEE state
- SPI EEPROM bootloader version and boot order

## 13. Fields not determined

- **Everything that needs a running system:** see section 12. Reason: this package only takes the image apart.
- **Tool versions from the binaries themselves** (`nft --version`, `apk --version` and so on): not determined; no
  binary from the image was run (no boot, no qemu-user). Package versions and `strings` were used instead.
- **Which kernel file and which `start4*.elf` the firmware picks**: not determined; `config.txt` names neither.
- **`/proc/cmdline`**: not determined; how the firmware merges the DT `bootargs` with `cmdline.txt` is not visible in
  the files.
- **`brcmfmac43236b.bin` version**: no version string found by `strings`.
- **Licence field of six firmware packages**: empty in the apk database.
- **SDK, toolchain, ImageBuilder and kernel-debug of bcm27xx/bcm2711** (5.3 point 4): not in the set, not downloaded
  and not analysed. The signed publisher list names them:
  `openwrt-sdk-25.12.6-bcm27xx-bcm2711_gcc-14.3.0_musl.Linux-x86_64.tar.zst` (`d5faf36e…8df8e`),
  `openwrt-toolchain-25.12.6-bcm27xx-bcm2711_gcc-14.3.0_musl.Linux-x86_64.tar.zst` (`b1bf188a…b579`),
  `openwrt-imagebuilder-25.12.6-bcm27xx-bcm2711.Linux-x86_64.tar.zst` (`78191c19…4d2a`), `kernel-debug.tar.zst`
  (`766423ec…dc78`) and `llvm-bpf-21.1.6.Linux-x86_64.tar.zst` (`3f07164c…f574`). From the image alone: GCC 14.3.0,
  GNU ld (binutils) 2.44, musl 1.2.5, target `aarch64_cortex-a72`. Use: the ImageBuilder would assemble an NMB Pi 4
  image from these packages without rebuilding; the native `nmbd` for arm64 can be built with the `armsr/armv8` SDK
  of I-02 if its compiler and musl match (I-09 to compare).
- **Feed versions from the signed `packages.adb`**: not parsed; the unsigned `index.json` was used (9a).

## 14. Tools missing on the PC

`usign`, `cpio`, `bsdtar`, `xorriso`, `parted`, `expect`, `socat`, `lz4` and `extract-ikconfig` are missing, and there
is no sudo, so no loop mounts. `usign`'s check was done with OpenSSL (section 1).

Present and used: `gzip`, `fdisk`, `sgdisk`, `dd`, `mtools` (`mdir`, `mcopy`), `debugfs`, `unsquashfs`
(squashfs-tools), `dtc`/`fdtget`, `modinfo`, `file`, `strings`, `xxd`, `gpg`, `openssl` 3.0.13, `curl`, `sha256sum`
and `python3 -I` (only for the ADB inflate and the `index.json` lookup, scripts outside the data folders).

## 15. Overlays on the boot partition (353 names, `ls overlays/*.dtbo`)

`act-led`, `adafruit-st7735r`, `adafruit18`, `adau1977-adc`, `adau7002-simple`, `ads1015`, `ads1115`, `ads7846`, `adv7282m`, `adv728x-m`, `akkordion-iqdacplus`, `allo-boss-dac-pcm512x-audio`, `allo-boss2-dac-audio`, `allo-digione`, `allo-katana-dac-audio`, `allo-piano-dac-pcm512x-audio`, `allo-piano-dac-plus-pcm512x-audio`, `anyspi`, `apds9960`, `applepi-dac`, `arducam-64mp`, `arducam-pivariety`, `at86rf233`, `audioinjector-addons`, `audioinjector-bare-i2s`, `audioinjector-isolated-soundcard`, `audioinjector-ultra`, `audioinjector-wm8731-audio`, `audiosense-pi`, `audremap-pi5`, `audremap`, `balena-fin`, `bcm2712d0`, `camera-mux-2port`, `camera-mux-4port`, `cap1106`, `chipcap2`, `chipdip-dac`, `cirrus-wm5102`, `cm-swap-i2c0`, `cma`, `crystalfontz-cfa050_pi_m`, `cutiepi-panel`, `dacberry400`, `dht11`, `dionaudio-kiwi`, `dionaudio-loco-v2`, `dionaudio-loco`, `disable-bt-pi5`, `disable-bt`, `disable-emmc2`, `disable-wifi-pi5`, `disable-wifi`, `dpi18`, `dpi18cpadhi`, `dpi24`, `draws`, `dwc-otg`, `dwc2`, `edt-ft5406`, `enc28j60-spi2`, `enc28j60`, `exc3000`, `ezsound-6x8iso`, `fbtft`, `fe-pi-audio`, `fsm-demo`, `gc9a01`, `ghost-amp`, `goodix`, `googlevoicehat-soundcard`, `gpio-charger`, `gpio-fan`, `gpio-hog`, `gpio-ir-tx`, `gpio-ir`, `gpio-key`, `gpio-led`, `gpio-no-bank0-irq`, `gpio-no-irq`, `gpio-poweroff`, `gpio-shutdown`, `hd44780-i2c-lcd`, `hd44780-lcd`, `hdmi-backlight-hwhack-gpio`, `hifiberry-adc`, `hifiberry-adc8x`, `hifiberry-amp`, `hifiberry-amp100`, `hifiberry-amp3`, `hifiberry-amp4pro`, `hifiberry-dac`, `hifiberry-dac8x`, `hifiberry-dacplus-pro`, `hifiberry-dacplus-std`, `hifiberry-dacplus`, `hifiberry-dacplusadc`, `hifiberry-dacplusadcpro`, `hifiberry-dacplusdsp`, `hifiberry-dacplushd`, `hifiberry-digi-pro`, `hifiberry-digi`, `highperi`, `hy28a`, `hy28b-2017`, `hy28b`, `i-sabre-q2m`, `i2c-bcm2708`, `i2c-fan`, `i2c-gpio`, `i2c-mux`, `i2c-pwm-pca9685a`, `i2c-rtc-gpio`, `i2c-rtc`, `i2c-sensor`, `i2c0-pi5`, `i2c0`, `i2c1-pi5`, `i2c1`, `i2c2-pi5`, `i2c3-pi5`, `i2c3`, `i2c4`, `i2c5`, `i2c6`, `i2s-dac`, `i2s-gpio28-31`, `i2s-master-dac`, `ilitek251x`, `imx219`, `imx258`, `imx290`, `imx296`, `imx327`, `imx335`, `imx378`, `imx415`, `imx462`, `imx477`, `imx500-pi5`, `imx500`, `imx519`, `imx708`, `interludeaudio-analog`, `interludeaudio-digital`, `iqaudio-codec`, `iqaudio-dac`, `iqaudio-dacplus`, `iqaudio-digi-wm8804-audio`, `iqs550`, `irs1125`, `jedec-spi-nor`, `justboom-both`, `justboom-dac`, `justboom-digi`, `ltc294x`, `max98357a`, `maxtherm`, `mbed-dac`, `mcp23017`, `mcp23s17`, `mcp2515-can0`, `mcp2515-can1`, `mcp2515`, `mcp251xfd`, `mcp3008`, `mcp3202`, `mcp342x`, `media-center`, `merus-amp`, `midi-uart0-pi5`, `midi-uart0`, `midi-uart1-pi5`, `midi-uart1`, `midi-uart2-pi5`, `midi-uart2`, `midi-uart3-pi5`, `midi-uart3`, `midi-uart4-pi5`, `midi-uart4`, `midi-uart5`, `minipitft13`, `miniuart-bt`, `mipi-dbi-spi`, `mlx90640`, `mmc`, `mz61581`, `ov2311`, `ov5647`, `ov64a40`, `ov7251`, `ov9281`, `papirus`, `pca953x`, `pcf857x`, `pcie-32bit-dma-pi5`, `pcie-32bit-dma`, `pciex1-compat-pi5`, `pibell`, `pifacedigital`, `pifi-40`, `pifi-dac-hd`, `pifi-dac-zero`, `pifi-mini-210`, `piglow`, `pimidi`, `pineboards-hat-ai`, `pineboards-hatdrive-poe-plus`, `piscreen`, `piscreen2r`, `pisound-micro`, `pisound-pi5`, `pisound`, `pitft22`, `pitft28-capacitive`, `pitft28-resistive`, `pitft35-resistive`, `pps-gpio`, `proto-codec`, `pwm-2chan`, `pwm-gpio-fan`, `pwm-gpio`, `pwm-ir-tx`, `pwm-pio`, `pwm`, `pwm1`, `qca7000-uart0`, `qca7000`, `ramoops-pi4`, `ramoops`, `rootmaster`, `rotary-encoder`, `rpi-backlight`, `rpi-codeczero`, `rpi-dacplus`, `rpi-dacpro`, `rpi-digiampplus`, `rpi-ft5406`, `rpi-fw-uart`, `rpi-poe-plus`, `rpi-poe`, `rpi-sense-v2`, `rpi-sense`, `rpi-tv`, `rra-digidac1-wm8741-audio`, `sainsmart18`, `sc16is750-i2c`, `sc16is750-spi0`, `sc16is752-i2c`, `sc16is752-spi0`, `sc16is752-spi1`, `sdhost`, `sdio-pi5`, `sdio`, `seeed-can-fd-hat-v1`, `seeed-can-fd-hat-v2`, `sh1106-spi`, `si446x-spi0`, `smi-dev`, `smi-nand`, `smi`, `spi-gpio35-39`, `spi-gpio40-45`, `spi-rtc`, `spi0-0cs`, `spi0-1cs-inverted`, `spi0-1cs`, `spi0-2cs`, `spi1-1cs`, `spi1-2cs`, `spi1-3cs`, `spi2-1cs-pi5`, `spi2-1cs`, `spi2-2cs-pi5`, `spi2-2cs`, `spi2-3cs`, `spi3-1cs-pi5`, `spi3-1cs`, `spi3-2cs-pi5`, `spi3-2cs`, `spi4-1cs`, `spi4-2cs`, `spi5-1cs-pi5`, `spi5-1cs`, `spi5-2cs-pi5`, `spi5-2cs`, `spi6-1cs`, `spi6-2cs`, `ssd1306-spi`, `ssd1306`, `ssd1327-spi`, `ssd1331-spi`, `ssd1351-spi`, `sunfounder-pipower3`, `sunfounder-pironman5`, `superaudioboard`, `sx150x`, `tc358743-audio`, `tc358743-pi5`, `tc358743`, `tinylcd35`, `tpm-slb9670`, `tpm-slb9673`, `uart0-pi5`, `uart0`, `uart1-pi5`, `uart1`, `uart2-pi5`, `uart2`, `uart3-pi5`, `uart3`, `uart4-pi5`, `uart4`, `uart5`, `udrc`, `ugreen-dabboard`, `upstream-pi4`, `upstream`, `vc4-fkms-v3d-pi4`, `vc4-fkms-v3d`, `vc4-kms-dpi-generic`, `vc4-kms-dpi-hyperpixel2r`, `vc4-kms-dpi-hyperpixel4`, `vc4-kms-dpi-hyperpixel4sq`, `vc4-kms-dpi-panel`, `vc4-kms-dsi-7inch`, `vc4-kms-dsi-generic`, `vc4-kms-dsi-ili9881-5inch`, `vc4-kms-dsi-ili9881-7inch`, `vc4-kms-dsi-lt070me05000-v2`, `vc4-kms-dsi-lt070me05000`, `vc4-kms-dsi-waveshare-800x480`, `vc4-kms-dsi-waveshare-panel-v2`, `vc4-kms-dsi-waveshare-panel`, `vc4-kms-kippah-7inch`, `vc4-kms-v3d-pi4`, `vc4-kms-v3d-pi5`, `vc4-kms-v3d`, `vc4-kms-vga666`, `vga666`, `vl805`, `w1-gpio-pi5`, `w1-gpio-pullup-pi5`, `w1-gpio-pullup`, `w1-gpio`, `w5500`, `watterott-display`, `waveshare-can-fd-hat-mode-a`, `waveshare-can-fd-hat-mode-b`, `wifimac`, `wittypi`, `wm8960-soundcard`, `ws2812-pio`
