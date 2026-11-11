# Image record: Fedora IoT 44 (aarch64 raw disk image and x86_64 ostree installer ISO)

Package I-08 (work plan section 5.4). The method is work plan 5.1, and every field of 5.2 is below.
Each value names the command that produced it. A value that could not be read says "not determined" and why.
Nothing here comes from memory. Only lists, versions and checksums are recorded. No file taken out of the image
is in this repository.

- Date of the analysis: 2026-11-11.
- Where it ran: dev PC, Windows with WSL2 Ubuntu 24.04, about 1.8 GB RAM and 4 threads, no root. The QEMU guests
  (aarch64 under TCG, x86_64 installer under KVM) ran on the same PC, one at a time; none reached a booted system
  (section 12).
- Work folder: `~/nmb-work/I-08/` (outside every repository; one fresh sub-folder per unpack).
- Material, read only: `~/src/NMB-depythonization-firmwares/platforms/raspberry-pi-4-arm64/` and `.../x86-64/`.
- **One compose, two forms.** Both files are Fedora IoT compose `44.20260427.0`. The values are tagged by source:
  - **[raw]** read from the aarch64 disk image: its partition table, the three file systems, the ostree repository
    and the deployment on the root partition. This image *is* an installed system (it is written to the SD card as it
    is), so for aarch64 no install step was needed.
  - **[ISO]** read from the x86_64 installer ISO: its boot files and the ostree repository it carries (Anaconda
    deploys that repository; the ISO's kickstart is one `ostreesetup` line, section 2).
  - **[commit]** read from an ostree commit with `ostree` 2024.5 (the PC's): `ostree show`, `ls -R`, `cat`. Both
    commits carry the rpm package list in their metadata (`rpmostree.rpmdb.pkglist`) and the rpm database
    (`/usr/share/rpm/rpmdb.sqlite`). The commit is the installed `/usr` byte for byte, so these values hold for an
    installed system of either architecture.
  - **[guest]** read from the aarch64 image booting under QEMU (section 12a). No run reached `multi-user.target`, so only kernel-log values carry this tag.

## 1. Files and checksums

| File | Size (bytes) | SHA-256 (publisher list) | Check |
| :--- | ---: | :--- | :--- |
| `Fedora-IoT-raw-44-20260427.0.aarch64.raw.xz` | 885,964,916 | `e563c67b576d0dfdbcf5c2d814bbadaca5bc363a046a12391d9e0b0b87cc9303` | OK |
| `Fedora-IoT-ostree-44-20260427.0.x86_64.iso` | 2,397,593,600 | `7e11fbee94fa1533a0fae66bf291fb7e7af398bfa249d19218a06ad888e9c0da` | OK |

- **How it was checked.** The set's README names both files but lists no sums for them. The publisher's lists were
  fetched on 2026-11-11 at 13:00 UTC from
  `https://dl.fedoraproject.org/pub/alt/iot/44/IoT/aarch64/images/Fedora-IoT-44-aarch64-20260427.0-CHECKSUM` and
  `.../x86_64/iso/Fedora-IoT-44-x86_64-20260427.0-CHECKSUM`. Both are clear-signed.
  - List SHA-256: aarch64 images `d62b20f644ca6ae1c46d42c6513b2e5842da21c990df3e7fb394e4f1728ccea5` (3 files: the raw
    image and two bootc OCI archives), x86_64 iso `24be240f24d531c177b644320e007fc4ba0394dffa4f78bcfc7eaef8992fdc21`
    (2 files: this ISO and the provisioner ISO). The sizes in the lists are the sizes above.
  - `sha256sum` of the two material files (exit 0) gave the two values above; compared with the lists: equal.
- **Signature of the lists.** `gpg --verify <list>` in an isolated `GNUPGHOME`: "Good signature from Fedora (44)
  <…>" (e-mail part of the key identity left out) for both (exit 0), made 2026-04-27 20:03:45 and 20:04:26 CEST.
  - Key `36F6 12DC F27F 7D1A 48A8 35E4 DBFC F71C 6D9F 90A6` (rsa4096, 2025-01-14), from
    `https://fedoraproject.org/fedora.gpg` (fetched the same minute, SHA-256
    `3b8a3660cca4851a67f67f56819e37e2f6b37ac8e905cb571b25d2127d1b7618`, keys for Fedora 43–46). The same fingerprint
    is printed on `https://fedoraproject.org/security/` (fetched 13:01 UTC). gpg says "not certified with a trusted
    signature", because no web of trust was set up.
- **Signature of the ostree commits.** Both commits carry a detached signature (`ostree show`: "Found 1 signature").
  Verified with `ostree show --gpg-verify-remote=nmbverify <commit>` after `ostree remote add --gpg-import=fedora.gpg
  nmbverify file:///nonexistent` in the work copies of the repositories: "Good signature from Fedora
  <…>" (e-mail part left out), key ID `DBFCF71C6D9F90A6` (the key above), for both (exit 0).
- **Inside the ISO.** `ostree fsck` on the ISO's repository (archive mode; every object is checked): exit 0, 26,042
  objects, no error.
- **Inside the raw image.** `ostree fsck` on the dumped repository fails on the first content object ("Corrupted file
  object"). That is expected and not a fault of the image: in a `bare` repository the object checksum covers owner,
  group and extended attributes (SELinux labels), and `debugfs rdump` without root keeps neither. The commit and its
  signature were checked instead (above); the content was not re-hashed.
- **Release identity.** Commit metadata `version` = `44.20260427.0` for both; `/usr/lib/os-release`:
  `PRETTY_NAME="Fedora Linux 44.20260427.0 (IoT Edition)"`, `VARIANT_ID=iot`, `SUPPORT_END=2027-05-19`. ISO
  `.discinfo`: `1777308640.054404`, `Fedora 44`, `x86_64`.

| | aarch64 [raw] | x86_64 [ISO] |
| :--- | :--- | :--- |
| ostree ref | `fedora-iot:fedora/stable/aarch64/iot` | `fedora/stable/x86_64/iot` |
| Commit | `9927129bba1220bb14b87702588f391a5227e900be900757308133c150ffadf3` | `35a2e036cdcf8f3067effe5a7a7415993481e9beaaca7eed7eabf53381852192` |
| Content checksum | `4596ceac2c61a99f1866a3ef351ec4b40f8c91cbd188a02d4c19e4b315041835` | `c208413acab7ebfe37937163a5c7b802d7da4117fc583128fd16b1e9c9e40dee` |
| Commit date | 2026-04-27 16:35:47 UTC | 2026-04-27 16:34:26 UTC |
| Parent (not in the repository) | `aacecd91…0d76` | `5141b654…4f38` |

- The decompressed raw image (`xz -dc | dd conv=sparse`, both exit 0; `xz -l`: one stream, 6,144 blocks of 1 MiB,
  CRC64) is 6,442,450,944 bytes, SHA-256 `af99a4ec92dd98d36824341896434b0368734383162ab0738d0d1b9c9b60ebed`.

## 2. Partition and file system layout, boot method

**[raw] aarch64.** Read with `fdisk -l` (exit 0) and `file` on each partition cut out with `dd skip= count=`. The table
is an **MBR** (`Disklabel type: dos`, identifier `0xc1748067`); `sgdisk -p` finds no GPT ("Found invalid GPT and
valid MBR") and only shows the MBR converted in memory.

| No. | Sectors (512 B) | Size | MBR type | Contents (`file`, `debugfs -R stats`, `mdir`) |
| :--- | :--- | :--- | :--- | :--- |
| 1 (boot flag) | 34816–1060863 | 501 MiB | `06` FAT16 | FAT16, label `ESP`, serial `6C32-40D6`; 417 files, 33,748,164 bytes, 488,775,680 bytes free |
| 2 | 1060864–5255167 | 2 GiB | `83` | ext4, label `boot`, UUID `3372c521-…`; 524,288 blocks of 4 KiB, 450,496 free (302,252,032 bytes used) |
| 3 | 5255168–12582911 | 3.5 GiB | `83` | ext4, label `root`, UUID `b2492255-…`; 915,968 blocks, 403,105 free (2,100,686,848 bytes used) |

- Both ext4 file systems: features `has_journal ext_attr … orphan_file … metadata_csum`, last mounted
  2026-04-27 18:45 (the compose). The image is meant to be grown on the card at first boot; the root partition ends
  at the end of the 6 GiB image.
- **ESP** (`mdir -/ -i p1.img ::`, `mtype`): `bootcode.bin`, `start*.elf`/`fixup*.dat` (Pi firmware, with
  `start4*.elf`/`fixup4*.dat` for the Pi 4), 10 DTBs (`bcm2710-*`, `bcm2711-rpi-4-b.dtb`, `bcm2711-rpi-400.dtb`,
  `bcm2711-rpi-cm4.dtb`, `-cm4s.dtb`), `overlays/` (368 entries), `rpi-u-boot.bin` (928,352 bytes), `config.txt`, and
  `EFI/BOOT/BOOTAA64.EFI` (shim), `EFI/BOOT/fbaa64.efi`, `EFI/fedora/{shimaa64.efi,grubaa64.efi,mmaa64.efi,grub.cfg}`.
  - `config.txt` (non-comment lines): `arm_64bit=1`, **`kernel=rpi-u-boot.bin`**, `enable_uart=1`, `dtparam=audio=on`,
    `camera_auto_detect=0`, `display_auto_detect=1`, `disable_fw_kms_setup=1`, `hdmi_ignore_cec_init=1`; under `[pi4]`
    `dtoverlay=upstream-pi4`, `dtoverlay=vc4-kms-v3d-pi4,cma-256`, `max_framebuffers=2`.
  - `EFI/fedora/grub.cfg` searches for `/grub2/grub.cfg` and loads it from the `boot` partition.
- **`/boot`** (`debugfs -R "ls -l"`): `loader -> loader.1`, `loader.1/entries/ostree-1.conf` (one BLS entry),
  `grub2/grub.cfg` (with `blscfg`), `grub2/grubenv`, `ignition.firstboot`, `efi/` (empty mount point) and
  `ostree/fedora-iot-bf0859…e2fe/` with `vmlinuz-6.19.10-300.fc44.aarch64` (16,650,752 bytes),
  `initramfs-6.19.10-300.fc44.aarch64.img` (89,891,371 bytes) and `dtb/`.
  - BLS entry `options`: `modprobe.blacklist=vc4 rw coreos.no_persist_ip ignition.platform.id=metal
    $ignition_firstboot root=LABEL=root ostree=/ostree/boot.1/fedora-iot/bf0859…e2fe/0`; `fdtdir` = the `dtb/` folder.
  - `ignition.firstboot` contains `set ignition_network_kcmdline='systemd.firstboot=off systemd.condition-first-boot=true'`;
    `grub.cfg` adds `ignition.firstboot` to the kernel line while that file exists.
- **Root partition**: `/ostree/repo` (`mode=bare`, `[sysroot] bootloader=none`, `readonly=true`), `/ostree/deploy/fedora-iot/`
  with one deployment `9927129b….0` (its `.origin`: `refspec=fedora-iot:fedora/stable/aarch64/iot`) and the shared
  `var/` (`home`, `roothome`, `opt`, `srv`, `usrlocal`, `mnt`, `spool`, `lib`, `log`, `.ostree-selabeled`), and
  `/ostree/boot.1 -> boot.1.1`.
  - The deployment's `/etc/fstab`: `/` `ext4 defaults,ro`, `/boot` `ext4 defaults`, `/boot/efi` `vfat
    umask=0077,shortname=winnt`.
  - `/usr/lib/ostree/prepare-root.conf` [commit]: `[sysroot] readonly = true`, `[composefs] enabled = yes`, written by
    the compose (`treefile.json` `postprocess`).
- **Boot chain on the Pi 4** (read from the files, not booted on a Pi): Pi GPU firmware (`start4.elf`) reads
  `config.txt` and starts **U-Boot** (`rpi-u-boot.bin`, package `uboot-images-armv8 2026.04-2.fc44`) as the "kernel";
  U-Boot's UEFI starts shim (`BOOTAA64.EFI`, `shim-aa64 16.1-5`), shim starts GRUB (`grub2-efi-aa64 2.12-56.fc44`), GRUB
  reads the BLS entry on `/boot` and loads the kernel, initramfs and the device tree from `fdtdir`. The same image
  boots on QEMU `virt` through UEFI firmware (section 12).

**[ISO] x86_64.** `fdisk -l` (exit 0): GPT; p1 sectors 64–4641175 (2.2 GiB, "Microsoft basic data", the ISO 9660
file system), p2 4641176–4682135 (20 MiB, EFI System), p3 4682136–4682735 (300 KiB). `file`: `ISO 9660 … 'Fedora-44-IoT-x86_64'
(bootable)`. Listed with `7z l` (26,368 files, 278 folders) and read with `bsdtar` (Rock Ridge names).

- Boot paths: UEFI `EFI/BOOT/BOOTX64.EFI` (shim) → `grubx64.efi`, with `BOOTIA32.EFI`/`grubia32.efi` too; BIOS through
  `images/eltorito.img` and `boot/grub2/i386-pc/`. Both `grub.cfg` offer "Install Fedora 44" with
  `linux /images/pxeboot/vmlinuz inst.stage2=hd:LABEL=Fedora-44-IoT-x86_64 inst.ks=hd:LABEL=Fedora-44-IoT-x86_64:/osbuild.ks`.
- `images/pxeboot/vmlinuz` 18,479,464 bytes (`file`: bzImage `6.19.10-300.fc44.x86_64 … #1 SMP PREEMPT_DYNAMIC Wed Mar
  25 18:23:49 UTC 2026`), `images/pxeboot/initrd.img` 262,850,180 bytes, `images/install.img` 1,159,270,400 bytes (squashfs
  4.0, xz, 45,926 inodes: the Anaconda installer), `images/efiboot.img` 20 MiB.
- `osbuild.ks` (134 bytes) is a single line: `ostreesetup --osname=fedora-iot --url=file:///run/install/repo/ostree/repo
  --ref=fedora/stable/x86_64/iot --remote=fedora-iot --nogpg`. Disk layout, users and network are not given, so
  Anaconda asks for them; the installed x86_64 layout therefore depends on the install and was not determined.
- `ostree/repo` (`mode=archive-z2`, one ref `fedora/stable/x86_64/iot`): 22,535 `.filez`, 3,448 `.dirtree`, 58
  `.dirmeta`, one `.commit` and one `.commitmeta`; 896 MB on disk after extraction.
- **Pitfall:** `7z x` read the Joliet names, which are cut to 64 characters, so the objects lost their checksum
  names and `ostree` found no commit. `bsdtar -xf` reads the Rock Ridge names and gave a working repository.

## 3. How an update is written

- **rpm-ostree / ostree, A/B by deployment.** [commit] `rpm-ostree 2026.1-2.fc44`, `ostree 2025.7-2.fc44`,
  `bootc 1.12.1-1.fc44`, `bootupd 0.2.32-3.fc44`. An update pulls a new signed commit from the remote
  (`/etc/ostree/remotes.d/fedora-iot.conf`: `url=https://ostree.fedoraproject.org/iot`, `gpg-verify=true`,
  `gpgkeypath=/etc/pki/rpm-gpg/`, `contenturl=mirrorlist=https://ostree.fedoraproject.org/iot/mirrorlist`) and writes it
  as a **second deployment** next to the running one, with its own BLS entry; the next boot uses it and the old one
  stays as the rollback entry. `/usr` is read-only (composefs, `prepare-root.conf`); `/etc` is merged three-way; `/var`
  is shared. Package layering (`rpm-ostree install`) creates a new deployment the same way.
- **Automatic updates: off by default.** `rpm-ostreed-automatic.timer` and `bootc-fetch-apply-updates.timer` exist but
  are not enabled (not under `/usr/etc/systemd/system/*.wants/`, `ls -R`); `/etc/rpm-ostreed.conf` has only `[Daemon]`
  (no `AutomaticUpdatePolicy`). Enabled timers: `dnf-makecache.timer`, `fstrim.timer`, `fwupd-refresh.timer`,
  `rpm-ostree-countme.timer`.
- **Health check and rollback: greenboot.** [commit] `greenboot 0.16.2-0.fc44` and
  `greenboot-default-health-checks`; `/usr/libexec/greenboot/greenboot` is an ELF binary (`od -c`: `177 E L F`), the
  checks are shell scripts: `required.d/01_repository_dns_check.sh`, `required.d/02_watchdog.sh`,
  `wanted.d/01_update_platforms_check.sh` (`/usr/lib/greenboot/check/`). `/etc/greenboot/greenboot.conf`:
  `GREENBOOT_MAX_BOOT_ATTEMPTS=3`, `DISABLED_HEALTHCHECKS=()`.
  - `greenboot-healthcheck.service` (`ExecStart=/usr/libexec/greenboot/greenboot health-check`, `RequiredBy=boot-complete.target`,
    drop-in `After=network-online.target`) is enabled; `greenboot-set-rollback-trigger.service` runs when `/etc` or
    `/var` need an update (`ConditionNeedsUpdate`).
  - The counter lives in GRUB: `/boot/grub2/grub.cfg` decrements `boot_counter` while `boot_success=0` and boots entry
  1 (the previous deployment) when it reaches 0, then `save_env`. The image's `grubenv` holds no variables
  (`debugfs cat`: only the `#` padding).
- **Security support** (criterion 2): `SUPPORT_END=2027-05-19` in `os-release`. The release schedule
  (`https://fedorapeople.org/groups/schedule/f-44/f-44-key-tasks.html`, fetched 2026-11-11) gives final release
  2026-04-28 and "Fedora Linux 44 End of Life" 2027-06-02 ("a changeable date"); the lifecycle page
  (`docs.fedoraproject.org/en-US/releases/lifecycle/`) says a release every six months, each maintained "for
  approximately 13 months". No subscription. Staying supported means a rebase to the next release every year
  (`rpm-ostree rebase`), not tested here.

## 4. Kernel and the configuration NMB needs

- **Kernel version: 6.19.10** (`kernel 6.19.10-300.fc44`, both architectures; commit metadata `ostree.linux` =
  `6.19.10-300.fc44.aarch64` / `.x86_64`). The kernel is in the commit at `/usr/lib/modules/<ver>/vmlinuz`.
  - SHA-256 of the x86_64 kernel in the commit `4b37e4e542a62c580c751787848be6c99e6f908f6712c8c6da85516b8d541de2`
    (18,479,464 bytes), the same file as the ISO's `images/pxeboot/vmlinuz` (`sha256sum`). aarch64
    `e55a5f9a9177267f58d1e957f58e6b294a4552cfe7cdf72174e686edd6b474da` (16,650,752 bytes), the same file as
    `/boot/ostree/…/vmlinuz-6.19.10-300.fc44.aarch64` on the raw image.
  - [guest] kernel log of the guest runs (12a): `Linux version 6.19.10-300.fc44.aarch64 (mockbuild@…) (gcc (GCC) 16.0.1 20260321 (Red Hat 16.0.1-0), GNU ld version 2.46-1.fc44) #1 SMP PREEMPT_DYNAMIC`; `uname` itself was not reached.
- **Where the config came from:** `/usr/lib/modules/6.19.10-300.fc44.<arch>/config` in each commit, taken out with
  `ostree cat <commit> <path>` (exit 0).
  - x86_64: 12,299 lines, `# Linux/x86_64 6.19.10-300.fc44.x86_64 Kernel Configuration`, SHA-256
    `1c0d2e478cdc33747fb3bdd9e332677c517b5791c8cbed0454d22e6e430042d5`.
  - aarch64: 13,862 lines, `# Linux/arm64 6.19.10-300.fc44.aarch64 Kernel Configuration`, SHA-256
    `044be830f7df5cb19c1b38ea42383e1a770b3722efb2b375cc86ac2efcfb2530`.
  - `CONFIG_IKCONFIG` is not set: there is no `/proc/config.gz`; the config file next to the modules is the source on an
    installed system too.
- **Source of every option below: those two files** (a `grep -E "^CONFIG_X=|^# CONFIG_X is not set"` loop). "Module"
  says whether the `.ko.xz` is in the commit (`ostree ls -R`, regular files under `/usr/lib/modules/<ver>/kernel/`).
  "Guest" was to be the `modprobe <name>; echo $?` loop in the booted aarch64 image; the guest did not get that far (12a), so only what its kernel log shows is filled in.

| Feature (5.2) | Option(s) | x86_64 | aarch64 | Module in commit | Guest (aarch64) |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Bridge, VLAN filtering | `BRIDGE`, `BRIDGE_VLAN_FILTERING`, `BRIDGE_NETFILTER` | m, y, m | m, y, m | `bridge` both | not reached (12a) |
| Bonding, team | `BONDING`, `NET_TEAM` | m, m | m, m | `bonding`, `team` both | not reached (12a) |
| VLAN 802.1Q | `VLAN_8021Q` | m | m | `8021q` both | not reached (12a) |
| nftables | `NF_TABLES` m, `NF_TABLES_INET` y, `NF_TABLES_NETDEV` y, `NF_TABLES_BRIDGE` m, `NFT_NAT` m, `NFT_MASQ` m, `NFT_CT` m, `NFT_FIB_INET` m, `NF_FLOW_TABLE` m, `NFT_FLOW_OFFLOAD` m, `NFT_COMPAT` m | as listed | as listed | `nf_tables`, `nft_nat`, `nft_masq`, `nft_fib_inet` both | not reached (12a) |
| iptables (legacy) | `IP_NF_IPTABLES`, `IP_NF_IPTABLES_LEGACY` m, `NETFILTER_XTABLES` y | as listed | as listed | — | — |
| WireGuard | `WIREGUARD` | m | m | `wireguard` both | not reached (12a) |
| PPP, PPPoE | `PPP`, `PPPOE`, `PPP_ASYNC`, `PPP_SYNC_TTY`, `PPP_DEFLATE`, `PPP_MPPE` | all m | all m | `ppp_generic`, `pppoe`, `pppox` both | not reached (12a) |
| L2TP | `L2TP`, `PPPOL2TP` | m, m | m, m | **not in the commit**: `l2tp_ppp`, `l2tp_netlink`, `l2tp_eth` are in `kernel-modules-extra` (below) | not reached (12a) |
| cfg80211, mac80211 | `CFG80211`, `MAC80211` (`CFG80211_WEXT` not set) | m, m | m, m | both | not reached (12a) |
| **mac80211_hwsim** | `MAC80211_HWSIM` | m | m | **not in the commit**: in `kernel-modules-internal` (below) | not reached (12a) |
| Wi-Fi drivers (examples) | `ATH9K`, `ATH9K_HTC`, `ATH10K_PCI`, `ATH11K`, `ATH12K`, `IWLWIFI`, `MT7601U`, `MT76x2U`, `MT7921U`, `MT7925U`, `RTL8XXXU`, `RTW88_8822BU`, `RTW89`, `RT2800USB`, `BRCMFMAC` (`_SDIO` y, `_USB` y) | all m | all m | 173 (x86_64) / 176 (aarch64) `.ko.xz` under `drivers/net/wireless/` | — |
| USB host | `USB`, `USB_XHCI_HCD`, `USB_EHCI_HCD`, `USB_XHCI_PCI` y; `USB_DWC2`; `USB_DWC3` m | y (DWC2 not set) | y (DWC2 m) | — | — |
| USB network | `USB_USBNET`, `USB_RTL8152`, `USB_NET_AX88179_178A`, `USB_NET_AX8817X`, `USB_NET_SMSC95XX`, `USB_LAN78XX`, `USB_NET_AQC111` | all m | all m | `usbnet`, `r8152`, `ax88179_178a`, `smsc95xx`, `lan78xx` both | not reached (12a) |
| USB tethering (4.27) | `USB_NET_RNDIS_HOST`, `USB_NET_CDCETHER`, `USB_NET_CDC_NCM`, `USB_IPHETH`, `USB_NET_QMI_WWAN`, `USB_NET_CDC_MBIM`, `USB_NET_HUAWEI_CDC_NCM`, `USB_NET_CDC_EEM` | all m | all m | `rndis_host`, `cdc_ether`, `cdc_ncm`, `ipheth`, `qmi_wwan`, `cdc_mbim`, `huawei_cdc_ncm` both | not reached (12a) |
| TUN, veth, macvlan, ipvlan, dummy, IFB | `TUN`, `VETH`, `MACVLAN`, `IPVLAN`, `DUMMY`, `IFB` | all m | all m | all present both | not reached (12a) |
| Policy routing | `IP_ADVANCED_ROUTER`, `IP_MULTIPLE_TABLES`, `IPV6_MULTIPLE_TABLES` | y | y | built in | — |
| Traffic control | `NET_SCHED` y, `NET_SCH_FQ_CODEL` y, `NET_SCH_CAKE`, `NET_SCH_HTB`, `NET_SCH_NETEM`, `NET_SCH_TBF`, `NET_SCH_FQ`, `NET_SCH_INGRESS` m; `NET_SCH_DEFAULT` not set | as listed | as listed | `sch_cake`, `sch_netem` both | not reached (12a) |
| Namespaces | `NAMESPACES`, `NET_NS`, `PID_NS`, `USER_NS`, `UTS_NS`, `IPC_NS`, `TIME_NS` | all y | all y | built in | — |
| seccomp | `SECCOMP`, `SECCOMP_FILTER` | y, y | y, y | built in | systemd banner `+SECCOMP` |
| **Landlock** | `SECURITY` y, `SECURITY_LANDLOCK` y, `LSM="lockdown,yama,integrity,selinux,bpf,landlock,ipe"` | y | y | built in | `landlock: Up and running.` (kernel log) |
| Other LSMs | `SECURITY_SELINUX` y (`DEFAULT_SECURITY_SELINUX` y, `SECURITY_SELINUX_BOOTPARAM` y), `SECURITY_YAMA` y, `SECURITY_LOCKDOWN_LSM` y, `LOCK_DOWN_KERNEL_FORCE_NONE` y, `SECURITY_IPE` y, `IMA` y; **`SECURITY_APPARMOR` not set** | as listed | as listed | — | `enforcing=1` after the policy load (kernel log, run 2) |
| cgroups | `CGROUPS`, `MEMCG`, `MEMCG_V1`, `CGROUP_PIDS`, `CPUSETS`, `CGROUP_BPF`, `BLK_CGROUP`, `CGROUP_FREEZER`, `CGROUP_DEVICE` | all y | all y | — | not reached (12a) |
| Watchdog | `WATCHDOG` y, `WATCHDOG_HANDLE_BOOT_ENABLED` y; `SOFT_WATCHDOG` m, `I6300ESB_WDT` m; `ITCO_WDT` m (x86_64 only); `BCM2835_WDT` m, `ARM_SMC_WATCHDOG` m (aarch64 only) | as listed | as listed | `softdog`, `i6300esb` both; `iTCO_wdt` x86_64; `bcm2835_wdt` aarch64 | not reached (12a) |
| Module signing | `MODULE_SIG` y, `MODULE_SIG_FORCE` not set, `MODULE_COMPRESS_XZ` y | as listed | as listed | `.ko.xz` | — |
| Raspberry Pi 4 platform | `ARCH_BCM2835` y, `ARCH_BCM` y, `PCIE_BRCMSTB` y, `RASPBERRYPI_FIRMWARE` y, `RASPBERRYPI_POWER` y, `RESET_RASPBERRYPI` y, `BCMGENET` m, `MMC_SDHCI_IPROC` m, `DRM_VC4` m, `HW_RANDOM_BCM2835` m, `SERIAL_8250_BCM2835AUX` y, `SENSORS_RASPBERRYPI_HWMON` m, `BCM2711_THERMAL` m | absent (`BCMGENET` m) | as listed | `genet`, `bcm2835_wdt`, `bcm2711_thermal`; DTBs `broadcom/bcm2711-rpi-4-b`, `-rpi-400`, `-rpi-cm4-io` | not on a Pi |
| Other | `BPF_SYSCALL` y, `XDP_SOCKETS` y, `TCP_CONG_BBR` m, `MPTCP` y, `TLS` m, `IPV6` y, `PREEMPT_DYNAMIC` y (x86_64 `PREEMPT_LAZY` y, aarch64 `PREEMPT_VOLUNTARY` y), `HZ=1000`, `NR_CPUS` 8192 / 4096, `OVERLAY_FS` m, `SQUASHFS` m, `EROFS_FS` m, `DM_VERITY` m, `FS_VERITY` y, `EFI_STUB` y, `ZRAM` m, `FW_CFG_SYSFS` m | as listed | as listed | — | — |

- **Built but not shipped.** `modules.order` in the commit lists 4,960 modules (x86_64) and 5,715 (aarch64); the
  commit carries 4,692 and 5,475 (`comm` of the two lists). The 268 missing x86_64 modules are exactly the two Fedora
  sub-packages the compose leaves out: `kernel-modules-extra` (112 `.ko.xz`) and `kernel-modules-internal` (156), both
  `6.19.10-300.fc44.x86_64`, fetched from `https://kojipkgs.fedoraproject.org/packages/kernel/6.19.10/300.fc44/x86_64/`
  (SHA-256 `9728950c…8069` and `00161277…cbad9`) and listed with `bsdtar -tf` (their RPM signatures were not checked).
  - `kernel-modules-internal` holds `internal/drivers/net/wireless/virtual/mac80211_hwsim.ko.xz` (and `wwan_hwsim`, the
    KUnit tests); `kernel-modules-extra` holds the four `l2tp_*` modules (with
    `/etc/modprobe.d/l2tp_*-blacklist.conf`), `slip`, `can` USB drivers, ham radio, `rds`, the rarer TCP congestion
    modules and `pktgen`.
  - For NMB: **`mac80211_hwsim` is not in the base**; the Wi-Fi tests that need it (roadmap 6A) need
    `kernel-modules-internal` layered (`rpm-ostree install`) on the test image only. L2TP is not used by NMB.

## 5. Modules and firmware files carried

- **Modules [commit]** (regular `.ko.xz` files under `/usr/lib/modules/<ver>/kernel/`, `ostree ls -R`): x86_64 4,692
  files, 175,979,452 bytes; aarch64 5,475 files, 101,083,428 bytes. `modules.builtin`: 348 lines (x86_64), 619
  (aarch64). aarch64 carries 1,334 `.dtb` files under `/usr/lib/modules/<ver>/dtb/`.
- [guest] the loaded-module list was not reached (12a).
- **Firmware [commit]** (`/usr/lib/firmware`, regular files): x86_64 1,370 files, 206,344,017 bytes; aarch64 1,218
  files, 189,565,777 bytes; all `.xz`-compressed except the regulatory database (`regulatory.db` 6,308 bytes with
  `regulatory.db.p7s`, from `wireless-regdb 2026.02.04`). Firmware packages: `linux-firmware 20260309-1.fc44` with
  `atheros-`, `brcmfmac-`, `realtek-`, `iwlwifi-mld-`, `iwlwifi-mvm-firmware`; x86_64 also `microcode_ctl`.
  - Largest folders (aarch64, files): `intel` 233, `qca` 94, `ath11k` 80, `ti` 67, `brcm` 56, `rtl_nic` 40, `ath10k` 40,
    `rtlwifi` 38, `rtl_bt` 33, `ar3k` 32, `mrvl` 28, `ath6k` 28, `cypress` 21, `rtw89` 17 (`awk` over the listing).
    `intel/iwlwifi` 169 files on both, `rtw88` 11, `ath9k_htc` 2.
  - **No MediaTek firmware**: no file under `/usr/lib/firmware/mediatek/` and no `mt7601u.bin` on either
    architecture, although the `mt7601u`, `mt76x2u`, `mt7921u` and `mt7925u` drivers are carried. A MediaTek USB
    adapter needs the separate firmware package layered in.
  - **Pi 4 Wi-Fi** (aarch64): `brcm/brcmfmac43455-sdio.bin.xz -> ../cypress/cyfmac43455-sdio.bin.xz` (363,984 bytes),
    `brcmfmac43455-sdio.clm_blob.xz -> ../cypress/cyfmac43455-sdio.clm_blob.xz`, and
    `brcmfmac43455-sdio.raspberrypi,4-model-b.txt.xz` (936 bytes; `,4-compute-module` links to it).
  - **Pi boot firmware** is on the ESP (section 2) and in the commit under `/usr/lib/ostree-boot/efi/` (406 files,
    24,953,936 bytes), from `bcm2711-firmware`, `bcm2835-firmware`, `bcm283x-firmware`, `bcm2712-firmware` (all
    `20260408-1.6bc3936.fc44`) and `bcm283x-overlays`; `bootupd` updates the ESP from there.

## 6. Init system, service manager, C library, package manager

- **Init: systemd 259** (`systemd 259.5-1.fc44`, package list; boot banner in section 12). Default target
  `multi-user.target` (`treefile.json` `default-target`).
- **Units enabled by default [commit]** (`ostree ls -R` of `/usr/etc/systemd/system/`; the aarch64 deployment's
  `/etc/systemd/system/` has the same 76 entries, `diff` of the two lists):
  - `multi-user.target.wants`: **`NetworkManager`**, **`firewalld`**, **`ModemManager`**, **`sshd`**, `chronyd`,
    `fdo-client-linuxapp`, `greenboot-healthcheck`, `greenboot-success.target`, `bootloader-update`,
    `coreos-ignition-write-issues`, `fedora-iot-config-remote-fix`, `authselect-apply-changes`,
    `lvm-devices-import.{path,service}`, `systemd-homed`, `systemd-oomd`, `remote-fs.target`,
    `remote-{cryptsetup,integritysetup,veritysetup}.target`.
  - `sysinit.target.wants`: `systemd-resolved`, `systemd-network-generator`, `systemd-sysext`, `systemd-confext`,
    `systemd-pstore`, `selinux-autorelabel-mark`, `lvm2-monitor`, `lvm2-lvmpolld.socket`, `fips-crypto-policy-overlay`,
    `systemd-boot-clear-sysfail`; `basic.target.wants`: `ignition-firstboot-complete`, `rpmdb-rebuild`;
    `network-online.target.wants`: `NetworkManager-wait-online`; `local-fs.target.wants`: `ostree-remount`.
  - Timers: section 3. Sockets: `dbus`, `dm-event`, `systemd-journald-audit`, `systemd-oomd`, `systemd-resolved-monitor`,
    `systemd-resolved-varlink`, `systemd-userdbd`.
  - Present but **not** enabled: `nftables.service`, `iwd.service`, `wpa_supplicant.service`, `rpm-ostreed-automatic.timer`,
    `bootc-fetch-apply-updates.timer`. `systemd-networkd` is not installed (no `systemd-networkd` package).
- **C library: glibc 2.43** (`glibc 2.43-2.fc44`, both architectures). [guest] `ldd --version` was not reached (12a).
- **Package manager:** rpm-ostree on top of `rpm 6.0.1-2.fc44` (`rpm-ostree 2026.1-2.fc44`); `dnf5 5.4.1.0` is present
  too. Kernel modules are signed and xz-compressed (section 4).
- **Package counts, and how they were obtained.**
  - **383 packages (x86_64), 393 (aarch64).** Source 1: the commit metadata `rpmostree.rpmdb.pkglist` (`ostree show
    --print-metadata-key=rpmostree.rpmdb.pkglist <commit>`), a list of (name, epoch, version, release, arch), covered by
    the commit signature. Source 2: the rpm database in the commit, `/usr/share/rpm/rpmdb.sqlite` (`ostree cat`;
    `sqlite3 -readonly … "select count(*) from Packages"` = 383 for x86_64), whose headers were read with a 30-line
    reader of the rpm header format (`python3 -I`, own script, for NAME, VERSION, RELEASE, ARCH, SIZE, LICENSE). The
    two sources give the same 383 name-version-release values (`diff`, empty). `rpm -qa` in the guest was not reached (12a).
  - The compose input `/usr/share/rpm-ostree/treefile.json` [commit] asks for 104 packages (x86_64) with
    `"recommends": false` and `"documentation": false`; the other 279 are dependencies.
  - Only aarch64: `arm-image-installer`, the five `bcm*` packages, `btrfs-progs`, `grub2-efi-aa64`, `lzo`,
    `optee_client`, `parted`, `shim-aa64`, `uboot-images-armv8`. Only x86_64: `grub2-efi-x64`, `microcode_ctl`,
    `shim-x64`. Every package both have is the same version (`join` of the two lists).
  - Sum of the rpm `SIZE` tags: 1,267,493,276 bytes (x86_64), 1,859,305,250 (aarch64).
- **Licences (criterion 7)** (LICENSE tag, same reader): 263 of the 383 x86_64 packages name a GPL variant. The firmware
  packages are `LicenseRef-Callaway-Redistributable-no-modification-permitted` (`atheros-`, `brcmfmac-`, `realtek-`,
  `iwlwifi-*-firmware`; `linux-firmware` adds GPL and MIT parts); the Pi packages `bcm2711/2712/2835/283x-firmware` are
  `LicenseRef-Fedora-Firmware`, `uboot-images-armv8` GPL-2.0-or-later with BSD/LGPL parts, `shim-*` BSD-3-Clause,
  `grub2-efi-*` GPL-3.0-or-later. Redistribution of the firmware unmodified is what an NMB image would do; that this is
  enough was not judged here.
- **Accounts [raw]:** `root` locked (`!` in `/etc/shadow`, only the lock state was read), no other login user; SSH
  keys and users come from Ignition (or FDO, `fdo-client-linuxapp`) at first boot.

## 7. Tools of 6A criterion 3, with versions

Presence: path in the commit listing (`ostree ls -R`, `/usr/bin`, `/usr/sbin` (a link to `bin`), `/usr/libexec`).
Versions: package version from the signed package list. The tools' own output in the booted aarch64 guest was not reached (12a).

| Tool | In the base | Package (both arches) | Guest output (aarch64) |
| :--- | :--- | :--- | :--- |
| systemd (init) | yes, `/usr/lib/systemd/systemd` | `systemd 259.5-1.fc44` | boot banner `systemd 259.5-1.fc44 running in system mode` (12a) |
| OpenRC, procd | no | — | — |
| nftables `nft` | **yes** | `nftables 1:1.1.6-2.fc44` (+ `nftables-services`, `python3-nftables`) | not reached (12a) |
| dnsmasq | **no** | — | — |
| hostapd | **no** | — | — |
| wpa_supplicant | **yes** (`wpa_supplicant`, `wpa_cli`) | `wpa_supplicant 1:2.11-9.fc44` | not reached (12a) |
| iwd | yes (`/usr/libexec/iwd`, `iwctl`) | `iwd 3.12-1.fc44` | not reached (12a) |
| iw | **yes** | `iw 6.17-2.fc44` | not reached (12a) |
| pppd | **no** | — (no `ppp`, no `NetworkManager-ppp`) | — |
| WireGuard `wg`, `wg-quick` | **no** (module only) | — | — |
| OpenVPN | **no** | — | — |
| `ip`, `ss`, `bridge` | **yes** | `iproute 6.17.0-2.fc44` | not reached (12a) |
| `tc` | **no** (`iproute-tc` not installed) | — | — |
| ethtool | **no** | — | — |
| iptables | **no** binary (`iptables-libs 1.8.11` only) | — | — |
| DHCP client | NetworkManager's internal client; no `dhclient`, `dhcpcd`, `udhcpc` | — | — |
| NetworkManager, `nmcli` | **yes** | `NetworkManager 1:1.56.0-1.fc44` (+ `-wifi`, `-wwan`) | not reached (12a) |
| firewalld, `firewall-cmd` | **yes** | `firewalld 2.4.0-2.fc44` | not reached (12a) |
| ModemManager, `mmcli`, `qmicli`, `mbimcli` | yes | `ModemManager 1.24.2-3.fc44`, `libqmi-utils 1.36.0`, `libmbim-utils 1.32.0` | — |
| systemd-networkd | no (package not installed); `resolvectl` yes (`systemd-resolved`) | — | — |
| podman, skopeo | yes | `podman 5:5.8.1-1.fc44`, `skopeo 1:1.22.0-2.fc44` | not reached (12a) |
| curl, ssh, sudo | yes | `curl 8.18.0-4.fc44`, `openssh 10.2p1-7.fc44`, `sudo 1.9.17-7.p2.fc44` | not reached (12a) |
| traceroute, ping, tracepath | yes | `traceroute 3:2.1.6-4.fc44`, `iputils 20250605-2.fc44` | — |
| tcpdump, mtr, iperf3, sqlite3, wget, busybox, perl | no | — | — |

- So of criterion 3 the base brings systemd, nftables, wpa_supplicant and iw. **dnsmasq, hostapd, pppd, WireGuard
  tools and OpenVPN are not in the base**; they would be layered with `rpm-ostree install` or built into an own
  compose (neither tried here).
- **What the Fedora 44 repositories hold for the missing ones** (Fedora's package metadata service,
  `curl https://mdapi.fedoraproject.org/f44/pkg/<name>` and `.../f44-updates/pkg/<name>`, fetched 2026-11-11 15:2x
  UTC, HTTP 200 unless said; read with `python3 -I`): `dnsmasq 2.92rel2-9.fc44`, `hostapd 2.11-4.fc44` (release only;
  updates 400), `ppp 2.5.1-7.fc44` (release only), `wireguard-tools 1.0.20260223-1.fc44`, `openvpn 2.7.9-1.fc44`,
  `iproute-tc 6.17.0-2.fc44` (release only), `ethtool 2:7.2-1.fc44`. None of them names a Python package or
  `/usr/bin/python3` among its direct requirements (the metadata's `requires`; not resolved transitively).
- The same service lists `kernel-modules-internal 7.3.3-200.fc44` (the sub-package with `mac80211_hwsim`, section 4)
  and `mt7xxx-firmware` / `mediatek-firmware 20261019-1.fc44` (MediaTek firmware, section 5); `kernel-core` is
  `7.3.3-200.fc44` (repository `updates`, queried 15:37 UTC). So the F44 kernel has moved from the image's 6.19.10 to
  7.3.3; the stock image (compose of 2026-04-27) is half a year behind and updates itself only when `rpm-ostree
  upgrade` is run (automatic updates are off, section 3). The kernel options of section 4 were read for 6.19.10 only.

## 8. Python (6A criterion 8)

- **The base brings a Python interpreter.** `python3 3.14.3-2.fc44` and `python3-libs` are in both package lists;
  `/usr/bin/python3 -> python3.14` (11,544 bytes) in both commits. Python packages in the base (13 named `python*`): `python3`,
  `python3-libs`, `python-pip-wheel`, `python3-audit`, `python3-dbus`, `python3-distro`, `python3-firewall`,
  `python3-gobject-base`, `python3-libselinux`, `python3-libsemanage`, `python3-nftables`, `python3-policycoreutils`,
  `python3-setools`; and `policycoreutils-python-utils` (Python scripts).
- `/usr/lib64/python3*` and `/usr/lib/python3*`: 2,923 regular files, 59,547,577 bytes (x86_64; aarch64 61,420,419);
  2,758 `.py`/`.pyc` files in the whole commit (both).
- **Why it is there** (rpm database, `Requirename` joined with `Name`): `firewalld` requires `/usr/bin/python3` and
  `python3-firewall`; `policycoreutils-python-utils` requires `/usr/bin/python3` and `python3-policycoreutils`;
  `setools-console` requires `/usr/bin/python3` and `python3-setools`. No installed package requires
  `policycoreutils-python-utils` or `setools-console`; `setools-console` is asked for in `treefile.json`,
  `policycoreutils-python-utils` is not (how it entered the compose was not determined). The compose's own
  post-processing runs `semanage permissive --add bootupd_t` (`treefile.json`), a Python tool.
- [guest] the scan for Python scripts in the booted system was not reached (12a); the counts above are from the commit.
- **Can it be left out:** only by not using the stock commit. Removing `firewalld`, `policycoreutils-python-utils` and
  `setools-console` (`rpm-ostree override remove …` on a device, or an own compose without them) is what it would
  take; whether anything else then still pulls `python3` was not tested. As shipped, Fedora IoT 44 **fails
  criterion 8**.

## 9. Own network and firewall manager, and how each is switched off (6A criterion 5)

What the base brings [commit]: **NetworkManager enabled**, **firewalld enabled** (nftables back end), ModemManager
enabled, `systemd-resolved` enabled, `nftables.service` disabled, `iwd` and `wpa_supplicant` units disabled (`wpa_supplicant.service` is
`Type=dbus`, `BusName=fi.w1.wpa_supplicant1`, with a D-Bus activation file in `/usr/share/dbus-1/system-services/`), no
systemd-networkd, no netplan.

- **firewalld configuration** (`/usr/etc/firewalld/firewalld.conf`, non-comment lines): `DefaultZone=public`,
  `FirewallBackend=nftables`, `FlushAllOnReload=yes`, `CleanupOnExit=yes`, `NftablesTableOwner=yes`,
  `ReloadPolicy=INPUT:DROP,FORWARD:DROP,OUTPUT:DROP`, `IPv6_rpfilter=strict`, `StrictForwardPorts=no`,
  `NftablesFlowtable=off`. `/etc/firewalld/zones/` is empty (zones come from `/usr/lib/firewalld/zones/`).
- **NetworkManager configuration:** `/etc/NetworkManager/NetworkManager.conf` has empty `[main]` and `[logging]`;
  `/usr/lib/NetworkManager/conf.d/` holds only `22-wifi-mac-addr.conf`; no connection profiles in the image.
- **nftables.service** (`nftables-services`): `ExecStart=/sbin/nft -f /etc/sysconfig/nftables.conf`,
  `ExecReload=/sbin/nft 'flush ruleset; include …'`, `ExecStop=/sbin/nft flush ruleset`. The shipped
  `/etc/sysconfig/nftables.conf` holds only comments (its `include "/etc/nftables/main.nft"` is commented out), and
  the sample `main.nft` has `flush ruleset` on line 15 (`grep -n`). It is disabled by default.

**How each is switched on, read from the image** (`ostree ls -R`, the deployment's `/etc` dumped with `debugfs`):

| Component | How it starts | Switching it off (proposal; not tested live) |
| :--- | :--- | :--- |
| NetworkManager | `multi-user.target.wants/NetworkManager.service`, `NetworkManager-wait-online` under `network-online.target`, `dbus-org.freedesktop.nm-dispatcher.service` alias | `systemctl mask NetworkManager NetworkManager-wait-online NetworkManager-dispatcher`, or leave the packages out of an own compose. No D-Bus activation file for it under `/usr/share/dbus-1/system-services/` |
| firewalld | `multi-user.target.wants/firewalld.service`, alias `dbus-org.fedoraproject.FirewallD1.service`; `netavark-firewalld-reload.service` (podman) is tied to it | `systemctl mask firewalld`, or leave it out (it is also what pulls Python in, section 8). No D-Bus activation file |
| ModemManager | `multi-user.target.wants/ModemManager.service`, alias `dbus-org.freedesktop.ModemManager1.service`, D-Bus activation file `org.freedesktop.ModemManager1.service` | `systemctl mask ModemManager` (with an activation file, a D-Bus call can start it while it is only disabled), or leave it out |
| wpa_supplicant, iwd | not enabled; D-Bus activation files `fi.w1.wpa_supplicant1.service`, `net.connman.iwd.service`, `net.connman.ead.service` | started on demand over D-Bus (by NetworkManager in the stock setup); `nmbd` starts `wpa_supplicant` itself, `iwd` masked or left out |
| nftables.service | not enabled | `systemctl mask nftables` (reload and stop flush every table, above) |
| systemd-resolved | `sysinit.target.wants/systemd-resolved.service` | not judged here; belongs to NMB's DNS (dnsmasq) design |

- Masking on a device writes links into `/etc`; in an own compose the units are simply not shipped.
- **Live behaviour not determined:** whether firewalld (`FlushAllOnReload=yes`, `NftablesTableOwner=yes`) leaves an
  own `inet` table of NMB alone across a reload, restart and stop, and whether NetworkManager takes over a bridge or
  dummy interface that NMB creates. The probe had these tests (`scripts/nmb-probe.sh`), but no guest run reached them
  (section 12a).

## 10. Installed packages

**[commit] x86_64, 383 packages** (`rpmostree.rpmdb.pkglist`; `name epoch:version-release.arch`, epoch shown only when
not 0):

`ModemManager 1.24.2-3.fc44.x86_64`, `ModemManager-glib 1.24.2-3.fc44.x86_64`, `NetworkManager 1:1.56.0-1.fc44.x86_64`, `NetworkManager-libnm 1:1.56.0-1.fc44.x86_64`, `NetworkManager-wifi 1:1.56.0-1.fc44.x86_64`, `NetworkManager-wwan 1:1.56.0-1.fc44.x86_64`, `aardvark-dns 2:1.17.0-3.fc44.x86_64`, `alternatives 1.33-5.fc44.x86_64`, `atheros-firmware 20260309-1.fc44.noarch`, `attr 2.5.2-8.fc44.x86_64`, `audit-libs 4.1.4-1.fc44.x86_64`, `authselect 1.7.1-1.fc44.x86_64`, `authselect-libs 1.7.1-1.fc44.x86_64`, `bash 5.3.9-3.fc44.x86_64`, `bash-completion 1:2.17-2.fc44.noarch`, `bootc 1.12.1-1.fc44.x86_64`, `bootupd 0.2.32-3.fc44.x86_64`, `brcmfmac-firmware 20260309-1.fc44.noarch`, `bubblewrap 0.11.0-4.fc44.x86_64`, `bzip2-libs 1.0.8-23.fc44.x86_64`, `ca-certificates 2025.2.80_v9.0.304-6.fc44.noarch`, `catatonit 0.2.1-5.fc44.x86_64`, `checkpolicy 3.10-1.fc44.x86_64`, `chrony 4.8-5.fc44.x86_64`, `clevis 21-14.fc44.x86_64`, `clevis-dracut 21-14.fc44.x86_64`, `clevis-luks 21-14.fc44.x86_64`, `clevis-pin-tpm2 0.5.4-1.fc44.x86_64`, `clevis-systemd 21-14.fc44.x86_64`, `composefs 1.0.8-5.fc44.x86_64`, `composefs-libs 1.0.8-5.fc44.x86_64`, `conmon 2:2.2.1-2.fc44.x86_64`, `container-selinux 4:2.247.0-1.fc44.noarch`, `containernetworking-plugins 1.9.1-1.fc44.x86_64`, `containers-common 5:0.67.0-1.fc44.noarch`, `containers-common-extra 5:0.67.0-1.fc44.noarch`, `coreutils 9.10-3.fc44.x86_64`, `coreutils-common 9.10-3.fc44.x86_64`, `cpio 2.15-9.fc44.x86_64`, `cracklib 2.9.11-10.fc44.x86_64`, `cracklib-dicts 2.9.11-10.fc44.x86_64`, `criu 4.2-13.fc44.x86_64`, `criu-libs 4.2-13.fc44.x86_64`, `crun 1.27-1.fc44.x86_64`, `crypto-policies 20251128-3.git19878fe.fc44.noarch`, `cryptsetup 2.8.4-1.fc44.x86_64`, `cryptsetup-libs 2.8.4-1.fc44.x86_64`, `curl 8.18.0-4.fc44.x86_64`, `cyrus-sasl-lib 2.1.28-35.fc44.x86_64`, `dbus 1:1.16.2-1.fc44.x86_64`, `dbus-broker 37-8.fc44.x86_64`, `dbus-common 1:1.16.2-1.fc44.noarch`, `dbus-libs 1:1.16.2-1.fc44.x86_64`, `device-mapper 1.02.212-2.fc44.x86_64`, `device-mapper-event 1.02.212-2.fc44.x86_64`, `device-mapper-event-libs 1.02.212-2.fc44.x86_64`, `device-mapper-libs 1.02.212-2.fc44.x86_64`, `device-mapper-persistent-data 1.3.1-3.fc44.x86_64`, `diffutils 3.12-5.fc44.x86_64`, `dnf5 5.4.1.0-1.fc44.x86_64`, `dnf5-plugins 5.4.1.0-1.fc44.x86_64`, `dosfstools 4.2-18.fc44.x86_64`, `dracut 108-6.fc44.x86_64`, `dracut-config-generic 108-6.fc44.x86_64`, `dracut-network 108-6.fc44.x86_64`, `duktape 2.7.0-11.fc44.x86_64`, `e2fsprogs 1.47.3-4.fc44.x86_64`, `e2fsprogs-libs 1.47.3-4.fc44.x86_64`, `efi-filesystem 6-6.fc44.noarch`, `efibootmgr 18-11.fc44.x86_64`, `efivar-libs 39-12.fc44.x86_64`, `elfutils-default-yama-scope 0.194-5.fc44.noarch`, `elfutils-libelf 0.194-5.fc44.x86_64`, `elfutils-libs 0.194-5.fc44.x86_64`, `expat 2.7.3-2.fc44.x86_64`, `fdo-client 0.5.5-8.fc44.x86_64`, `fdo-owner-cli 0.5.5-8.fc44.x86_64`, `fedora-gpg-keys 44-1.noarch`, `fedora-iot-config 1-5.fc44.noarch`, `fedora-release-common 44-17.noarch`, `fedora-release-identity-iot 44-17.noarch`, `fedora-release-iot 44-17.noarch`, `fedora-repos 44-1.noarch`, `file 5.46-9.fc44.x86_64`, `file-libs 5.46-9.fc44.x86_64`, `filesystem 3.18-52.fc44.x86_64`, `findutils 1:4.10.0-7.fc44.x86_64`, `firewalld 2.4.0-2.fc44.noarch`, `firewalld-filesystem 2.4.0-2.fc44.noarch`, `fmt 11.2.0-4.fc44.x86_64`, `fuse-common 3.18.2-1.fc44.x86_64`, `fuse3 3.18.2-1.fc44.x86_64`, `fuse3-libs 3.18.2-1.fc44.x86_64`, `fwupd 2.1.1-1.fc44.x86_64`, `fwupd-efi 1.8-1.fc44.x86_64`, `fwupd-plugin-modem-manager 2.1.1-1.fc44.x86_64`, `fwupd-plugin-uefi-capsule-data 2.1.1-1.fc44.x86_64`, `gawk 5.3.2-3.fc44.x86_64`, `gdbm 1:1.23-11.fc44.x86_64`, `gdbm-libs 1:1.23-11.fc44.x86_64`, `gdisk 1.0.10-5.fc44.x86_64`, `gettext-envsubst 0.26-4.fc44.x86_64`, `gettext-libs 0.26-4.fc44.x86_64`, `gettext-runtime 0.26-4.fc44.x86_64`, `glib2 2.88.0-1.fc44.x86_64`, `glibc 2.43-2.fc44.x86_64`, `glibc-common 2.43-2.fc44.x86_64`, `glibc-minimal-langpack 2.43-2.fc44.x86_64`, `gmp 1:6.3.0-5.fc44.x86_64`, `gnulib-l10n 20241231-2.fc44.noarch`, `gnupg2 2.4.9-5.fc44.x86_64`, `gnupg2-dirmngr 2.4.9-5.fc44.x86_64`, `gnupg2-gpg-agent 2.4.9-5.fc44.x86_64`, `gnupg2-gpgconf 2.4.9-5.fc44.x86_64`, `gnupg2-keyboxd 2.4.9-5.fc44.x86_64`, `gnupg2-verify 2.4.9-5.fc44.x86_64`, `gnutls 3.8.12-1.fc44.x86_64`, `gpgme 2.0.1-3.fc44.x86_64`, `greenboot 0.16.2-0.fc44.x86_64`, `greenboot-default-health-checks 0.16.2-0.fc44.x86_64`, `grep 3.12-3.fc44.x86_64`, `grub2-common 1:2.12-56.fc44.noarch`, `grub2-efi-x64 1:2.12-56.fc44.x86_64`, `grub2-tools 1:2.12-56.fc44.x86_64`, `grub2-tools-minimal 1:2.12-56.fc44.x86_64`, `gzip 1.14-2.fc44.x86_64`, `hostname 3.25-4.fc44.x86_64`, `hwdata 0.405-1.fc44.noarch`, `ignition 2.26.0-2.fc44.x86_64`, `ignition-edge 2.26.0-2.fc44.x86_64`, `ima-evm-utils 1.6.2-8.fc44.x86_64`, `ima-evm-utils-libs 1.6.2-8.fc44.x86_64`, `inih 62-2.fc44.x86_64`, `iproute 6.17.0-2.fc44.x86_64`, `iptables-libs 1.8.11-13.fc44.x86_64`, `iputils 20250605-2.fc44.x86_64`, `iw 6.17-2.fc44.x86_64`, `iwd 3.12-1.fc44.x86_64`, `iwlwifi-mld-firmware 20260309-1.fc44.noarch`, `iwlwifi-mvm-firmware 20260309-1.fc44.noarch`, `jansson 2.14-4.fc44.x86_64`, `jose 14-6.fc44.x86_64`, `jq 1.8.1-2.fc44.x86_64`, `json-c 0.18-8.fc44.x86_64`, `json-glib 1.10.8-5.fc44.x86_64`, `kbd 2.9.0-3.fc44.x86_64`, `kbd-legacy 2.9.0-3.fc44.noarch`, `kbd-misc 2.9.0-3.fc44.noarch`, `kernel 6.19.10-300.fc44.x86_64`, `kernel-core 6.19.10-300.fc44.x86_64`, `kernel-modules 6.19.10-300.fc44.x86_64`, `kernel-modules-core 6.19.10-300.fc44.x86_64`, `keyutils 1.6.3-7.fc44.x86_64`, `keyutils-libs 1.6.3-7.fc44.x86_64`, `kmod 34.2-4.fc44.x86_64`, `kmod-libs 34.2-4.fc44.x86_64`, `krb5-libs 1.22.2-2.fc44.x86_64`, `less 691-2.fc44.x86_64`, `libacl 2.3.2-6.fc44.x86_64`, `libaio 0.3.111-23.fc44.x86_64`, `libarchive 3.8.4-2.fc44.x86_64`, `libassuan 2.5.7-5.fc44.x86_64`, `libattr 2.5.2-8.fc44.x86_64`, `libblkid 2.41.3-12.fc44.x86_64`, `libbpf 2:1.6.3-1.fc44.x86_64`, `libbrotli 1.2.0-3.fc44.x86_64`, `libbsd 0.12.2-7.fc44.x86_64`, `libcap 2.77-2.fc44.x86_64`, `libcap-ng 0.9.2-1.fc44.x86_64`, `libcbor 0.13.0-2.fc44.x86_64`, `libcom_err 1.47.3-4.fc44.x86_64`, `libcurl 8.18.0-4.fc44.x86_64`, `libdnf5 5.4.1.0-1.fc44.x86_64`, `libdnf5-cli 5.4.1.0-1.fc44.x86_64`, `libdrm 2.4.131-1.fc44.x86_64`, `libeconf 0.7.9-3.fc44.x86_64`, `libedit 3.1-58.20251016cvs.fc44.x86_64`, `libell 0.83-1.fc44.x86_64`, `libevent 2.1.12-17.fc44.x86_64`, `libfdisk 2.41.3-12.fc44.x86_64`, `libffi 3.5.2-2.fc44.x86_64`, `libfido2 1.16.0-5.fc44.x86_64`, `libfsverity 1.6-4.fc44.x86_64`, `libgcc 16.0.1-0.10.fc44.x86_64`, `libgcrypt 1.12.1-1.fc44.x86_64`, `libgomp 16.0.1-0.10.fc44.x86_64`, `libgpg-error 1.58-2.fc44.x86_64`, `libgudev 238-9.fc44.x86_64`, `libidn2 2.3.8-3.fc44.x86_64`, `libjcat 0.2.5-2.fc44.x86_64`, `libjose 14-6.fc44.x86_64`, `libkcapi 1.5.0-7.fc44.x86_64`, `libkcapi-hasher 1.5.0-7.fc44.x86_64`, `libkcapi-hmaccalc 1.5.0-7.fc44.x86_64`, `libksba 1.6.7-5.fc44.x86_64`, `liblastlog2 2.41.3-12.fc44.x86_64`, `libluksmeta 10-2.fc44.x86_64`, `libmbim 1.32.0-3.fc44.x86_64`, `libmbim-utils 1.32.0-3.fc44.x86_64`, `libmd 1.1.0-9.fc44.x86_64`, `libmnl 1.0.5-9.fc44.x86_64`, `libmodulemd 2.15.2-7.fc44.x86_64`, `libmount 2.41.3-12.fc44.x86_64`, `libndp 1.9-5.fc44.x86_64`, `libnet 1.3-7.fc44.x86_64`, `libnetfilter_conntrack 1.1.1-1.fc44.x86_64`, `libnfnetlink 1.0.1-32.fc44.x86_64`, `libnftnl 1.3.1-2.fc44.x86_64`, `libnghttp2 1.68.0-3.fc44.x86_64`, `libnghttp3 1.15.0-1.fc44.x86_64`, `libnl3 3.12.0-3.fc44.x86_64`, `libnsl2 2.0.1-5.fc44.x86_64`, `libnvme 1.16.1-2.fc44.x86_64`, `libpciaccess 0.16-17.fc44.x86_64`, `libpkgconf 2.5.1-1.fc44.x86_64`, `libpq 18.0-4.fc44.x86_64`, `libpsl 0.21.5-7.fc44.x86_64`, `libpwquality 1.4.5-15.fc44.x86_64`, `libqb 2.0.9-2.fc44.x86_64`, `libqmi 1.36.0-3.fc44.x86_64`, `libqmi-utils 1.36.0-3.fc44.x86_64`, `libqrtr-glib 1.2.2-9.fc44.x86_64`, `librepo 1.20.0-5.fc44.x86_64`, `libseccomp 2.6.0-3.fc44.x86_64`, `libsecret 0.21.7-10.fc44.x86_64`, `libselinux 3.10-1.fc44.x86_64`, `libselinux-utils 3.10-1.fc44.x86_64`, `libsemanage 3.10-1.fc44.x86_64`, `libsepol 3.10-1.fc44.x86_64`, `libslirp 4.9.1-3.fc44.x86_64`, `libsmartcols 2.41.3-12.fc44.x86_64`, `libsolv 0.7.36-2.fc44.x86_64`, `libss 1.47.3-4.fc44.x86_64`, `libssh 0.12.0-1.fc44.x86_64`, `libssh-config 0.12.0-1.fc44.noarch`, `libsss_idmap 2.12.0-4.fc44.x86_64`, `libsss_nss_idmap 2.12.0-4.fc44.x86_64`, `libsss_sudo 2.12.0-4.fc44.x86_64`, `libstdc++ 16.0.1-0.10.fc44.x86_64`, `libtasn1 4.21.0-1.fc44.x86_64`, `libtextstyle 0.26-4.fc44.x86_64`, `libtirpc 1.3.7-2.fc44.x86_64`, `libtool-ltdl 2.5.4-10.fc44.x86_64`, `libunistring 1.1-11.fc44.x86_64`, `libusb1 1.0.29-5.fc44.x86_64`, `libutempter 1.2.1-20.fc44.x86_64`, `libuuid 2.41.3-12.fc44.x86_64`, `libverto 0.3.2-12.fc44.x86_64`, `libxcrypt 4.5.2-3.fc44.x86_64`, `libxml2 2.12.10-6.fc44.x86_64`, `libxmlb 0.3.25-1.fc44.x86_64`, `libyaml 0.2.5-18.fc44.x86_64`, `libzstd 1.5.7-5.fc44.x86_64`, `linux-firmware 20260309-1.fc44.noarch`, `linux-firmware-whence 20260309-1.fc44.noarch`, `lua-libs 5.4.8-5.fc44.x86_64`, `luksmeta 10-2.fc44.x86_64`, `lvm2 2.03.38-2.fc44.x86_64`, `lvm2-libs 2.03.38-2.fc44.x86_64`, `lz4-libs 1.10.0-4.fc44.x86_64`, `microcode_ctl 2:2.1-74.fc44.x86_64`, `mokutil 2:0.7.2-3.fc44.x86_64`, `mpdecimal 4.0.1-3.fc44.x86_64`, `mpfr 4.2.2-3.fc44.x86_64`, `ncurses-base 6.6-1.fc44.noarch`, `ncurses-libs 6.6-1.fc44.x86_64`, `netavark 2:1.17.2-1.fc44.x86_64`, `nettle 3.10.1-3.fc44.x86_64`, `nftables 1:1.1.6-2.fc44.x86_64`, `nftables-services 1:1.1.6-2.fc44.noarch`, `ngtcp2 1.21.0-1.fc44.x86_64`, `ngtcp2-crypto-ossl 1.21.0-1.fc44.x86_64`, `npth 1.8-4.fc44.x86_64`, `nss-altfiles 2.23.0-9.fc44.x86_64`, `oniguruma 6.9.10-4.fc44.x86_64`, `openldap 2.6.13-1.fc44.x86_64`, `openssh 10.2p1-7.fc44.x86_64`, `openssh-clients 10.2p1-7.fc44.x86_64`, `openssh-server 10.2p1-7.fc44.x86_64`, `openssl 1:3.5.5-1.fc44.x86_64`, `openssl-libs 1:3.5.5-1.fc44.x86_64`, `os-prober 1.81-11.fc44.x86_64`, `ostree 2025.7-2.fc44.x86_64`, `ostree-libs 2025.7-2.fc44.x86_64`, `p11-kit 0.26.2-1.fc44.x86_64`, `p11-kit-trust 0.26.2-1.fc44.x86_64`, `pam 1.7.2-1.fc44.x86_64`, `pam-libs 1.7.2-1.fc44.x86_64`, `passim-libs 0.1.10-3.fc44.x86_64`, `passt 0^20260120.g386b5f5-1.fc44.x86_64`, `passt-selinux 0^20260120.g386b5f5-1.fc44.noarch`, `pcre2 10.47-1.fc44.1.x86_64`, `pcre2-syntax 10.47-1.fc44.1.noarch`, `pinentry 1.3.2-2.fc44.x86_64`, `pkgconf 2.5.1-1.fc44.x86_64`, `pkgconf-m4 2.5.1-1.fc44.noarch`, `pkgconf-pkg-config 2.5.1-1.fc44.x86_64`, `podman 5:5.8.1-1.fc44.x86_64`, `podman-sequoia 0.3.2-1.fc44.x86_64`, `policycoreutils 3.10-1.fc44.x86_64`, `policycoreutils-python-utils 3.10-1.fc44.noarch`, `polkit 127-2.fc44.2.x86_64`, `polkit-libs 127-2.fc44.2.x86_64`, `popt 1.19-10.fc44.x86_64`, `procps-ng 4.0.6-1.fc44.x86_64`, `protobuf 3.19.6-20.fc44.x86_64`, `protobuf-c 1.5.2-2.fc44.x86_64`, `psmisc 23.7-7.fc44.x86_64`, `publicsuffix-list-dafsa 20260116-1.fc44.noarch`, `python-pip-wheel 26.0.1-2.fc44.noarch`, `python3 3.14.3-2.fc44.x86_64`, `python3-audit 4.1.4-1.fc44.x86_64`, `python3-dbus 1.4.0-9.fc44.x86_64`, `python3-distro 1.9.0-11.fc44.noarch`, `python3-firewall 2.4.0-2.fc44.noarch`, `python3-gobject-base 3.56.2-1.fc44.x86_64`, `python3-libs 3.14.3-2.fc44.x86_64`, `python3-libselinux 3.10-1.fc44.x86_64`, `python3-libsemanage 3.10-1.fc44.x86_64`, `python3-nftables 1:1.1.6-2.fc44.noarch`, `python3-policycoreutils 3.10-1.fc44.noarch`, `python3-setools 4.6.0-5.fc44.x86_64`, `readline 8.3-4.fc44.x86_64`, `realtek-firmware 20260309-1.fc44.noarch`, `rootfiles 9.0-6.fc44.noarch`, `rpm 6.0.1-2.fc44.x86_64`, `rpm-build-libs 6.0.1-2.fc44.x86_64`, `rpm-libs 6.0.1-2.fc44.x86_64`, `rpm-ostree 2026.1-2.fc44.x86_64`, `rpm-ostree-libs 2026.1-2.fc44.x86_64`, `rpm-plugin-ima 6.0.1-2.fc44.x86_64`, `rpm-plugin-selinux 6.0.1-2.fc44.x86_64`, `rpm-sequoia 1.10.1-1.fc44.x86_64`, `rpm-sign-libs 6.0.1-2.fc44.x86_64`, `screen 5.0.1-6.fc44.x86_64`, `sdbus-cpp 2.2.1-2.fc44.x86_64`, `sed 4.9-7.fc44.x86_64`, `selinux-policy 43.3-1.fc44.noarch`, `selinux-policy-targeted 43.3-1.fc44.noarch`, `setools-console 4.6.0-5.fc44.x86_64`, `setup 2.15.0-28.fc44.noarch`, `shadow-utils 2:4.19.0-6.fc44.x86_64`, `shadow-utils-subid 2:4.19.0-6.fc44.x86_64`, `shared-mime-info 2.4-3.fc44.x86_64`, `shim-x64 16.1-5.x86_64`, `skopeo 1:1.22.0-2.fc44.x86_64`, `slirp4netns 1.3.1-4.fc44.x86_64`, `sqlite-libs 3.51.2-1.fc44.x86_64`, `sssd-client 2.12.0-4.fc44.x86_64`, `sudo 1.9.17-7.p2.fc44.x86_64`, `systemd 259.5-1.fc44.x86_64`, `systemd-libs 259.5-1.fc44.x86_64`, `systemd-pam 259.5-1.fc44.x86_64`, `systemd-resolved 259.5-1.fc44.x86_64`, `systemd-shared 259.5-1.fc44.x86_64`, `systemd-sysusers 259.5-1.fc44.x86_64`, `systemd-udev 259.5-1.fc44.x86_64`, `tar 2:1.35-8.fc44.x86_64`, `tmux 3.6a-1.fc44.x86_64`, `tpm2-pkcs11 1.9.1-7.fc44.x86_64`, `tpm2-tools 5.7-5.fc44.x86_64`, `tpm2-tss 4.1.3-9.fc44.x86_64`, `tpm2-tss-fapi 4.1.3-9.fc44.x86_64`, `traceroute 3:2.1.6-4.fc44.x86_64`, `tzdata 2026a-1.fc44.noarch`, `usbguard 1.1.3-7.fc44.x86_64`, `usbguard-selinux 1.1.3-7.fc44.noarch`, `userspace-rcu 0.15.6-1.fc44.x86_64`, `util-linux 2.41.3-12.fc44.x86_64`, `util-linux-core 2.41.3-12.fc44.x86_64`, `vim-data 2:9.2.240-1.fc44.noarch`, `vim-minimal 2:9.2.240-1.fc44.x86_64`, `wireless-regdb 2026.02.04-1.fc44.noarch`, `wpa_supplicant 1:2.11-9.fc44.x86_64`, `xfsprogs 6.18.0-2.fc44.x86_64`, `xz 1:5.8.2-2.fc44.x86_64`, `xz-libs 1:5.8.2-2.fc44.x86_64`, `yajl 2.1.0-40.fc44.x86_64`, `zchunk-libs 1.5.1-4.fc44.x86_64`, `zlib-ng-compat 2.3.3-3.fc44.x86_64`, `zram-generator 1.2.1-5.fc44.x86_64`, `zram-generator-defaults 1.2.1-5.fc44.noarch`, `zstd 1.5.7-5.fc44.x86_64`

**[raw] aarch64, 393 packages:** the x86_64 list with `grub2-efi-x64`, `microcode_ctl`, `shim-x64` replaced by
`arm-image-installer 5.3-3.fc44`, `bcm2711-firmware`, `bcm2712-firmware`, `bcm2835-firmware`, `bcm283x-firmware`,
`bcm283x-overlays` (all `20260408-1.6bc3936.fc44`), `btrfs-progs 6.19.1-1.fc44`, `grub2-efi-aa64 1:2.12-56.fc44`,
`lzo 2.10-16.fc44`, `optee_client 4.9.0-1.fc44`, `parted 3.6-14.fc44`, `shim-aa64 16.1-5`,
`uboot-images-armv8 1:2026.04-2.fc44`; every shared package has the same version, with the arch `aarch64`.

## 11. Size on disk

| | Bytes | Source |
| :--- | ---: | :--- |
| aarch64 `raw.xz` (download) | 885,964,916 | `stat` |
| aarch64 raw image (card size) | 6,442,450,944; 2,227,609,600 allocated as a sparse file | `xz -l`, `ls -s` |
| aarch64 ESP used | 33,748,164 in 417 files (of 501 MiB) | `mdir -/` |
| aarch64 `/boot` used | 302,252,032 (of 2 GiB); kernel 16,650,752 + initramfs 89,891,371 + DTBs | `debugfs stats` |
| aarch64 root partition used | 2,100,686,848 (of 3.5 GiB; 1,651,118,080 free) | `debugfs stats` |
| aarch64 commit (the OS tree) | 1,862,832,177 in 24,759 files, 4,602 links, 5,414 folders | `ostree ls -R`, sum of regular files |
| aarch64 ostree repository dump | 1.8 GiB | `du -sh` |
| x86_64 ISO | 2,397,593,600 | `stat` |
| x86_64 ostree repository in the ISO | 861,029,959 in the ISO listing, 896 MiB extracted | `7z x` summary, `du -sh` |
| x86_64 commit (the OS tree) | 1,382,684,853 in 21,998 files, 4,578 links, 5,172 folders | `ostree ls -R` |

- Of the x86_64 tree: kernel modules 175,979,452, firmware 206,344,017, Python 59,547,577 (`/usr/lib*/python3*`).
  `documentation: false` in the compose already leaves out docs; `install-langs` keeps 40 locales.
- [guest] storage after boot was not reached (12a).

## 12. QEMU boot tests

### 12a. aarch64 raw image (the installed system), TCG

**How it was run.** A sparse copy of the decompressed raw image (`cp --sparse=always`). In a copy of partition 2 cut
out by offset, the one BLS entry was replaced with `debugfs -w` (`rm`, `write`, owner `0:0`, label
`system_u:object_r:boot_t:s0`; `e2fsck -fn` rc 0) and written back with `dd seek=1060864 conv=notrunc` (`cmp`
equal). The only change to the entry: `ignition.platform.id=metal` → `qemu`, and `console=ttyAMA0,115200n8
systemd.default_timeout_start_sec=900 systemd.default_device_timeout_sec=900` added (slow TCG). An Ignition 3.4.0
config passed through `fw_cfg` (`opt/com.coreos/config`) sets the host name `nmb-i08` and writes
`/etc/nmb-probe.sh` with an enabled `nmb-probe.service` (`After=multi-user.target`) that waits for `systemctl
is-system-running --wait`, prints the figures to a virtio serial port and powers off. Probe, host name and kernel
arguments are test additions, not part of the base; sections 1–11 were read before them. Each run held `flock
~/nmb-work/qemu.lock` for the whole boot:

```
qemu-system-aarch64 -machine virt -accel tcg,thread=multi,tb-size=<256|64> -cpu cortex-a57 -smp 3 -m 512
  -nographic -monitor none -serial file:serial-a64-<run>.log -no-reboot
  -drive if=pflash,format=raw,readonly=on,file=/usr/share/AAVMF/AAVMF_CODE.fd -drive if=pflash,format=raw,file=vars-a64.fd
  -drive file=disk-a64.raw,format=raw,if=none,id=d0 -device virtio-blk-pci,drive=d0
  -netdev user,id=n0,net=192.0.2.0/24 -device virtio-net-pci,netdev=n0,mac=02:00:00:00:08:01 -device virtio-rng-pci
  -device virtio-serial-pci -chardev file,id=pr,path=probe-a64-<run>.log -device virtserialport,chardev=pr,name=nmb.probe
  -fw_cfg name=opt/com.coreos/config,file=probe.ign
```

- QEMU 8.2.2 on the PC, UEFI firmware `qemu-efi-aarch64 2024.02-2ubuntu0.10` (AAVMF). arm64 under **TCG** (no KVM for
  arm64 on an x86 PC). The path was UEFI → `EFI/BOOT/BOOTAA64.EFI` (shim) → GRUB → BLS entry (kernel line
  `BOOT_IMAGE=(hd0,msdos2)/ostree/fedora-iot-bf0859…/vmlinuz-6.19.10-300.fc44.aarch64`); the Pi path through
  `rpi-u-boot.bin` was not exercised.

| Run | Settings | Result |
| :--- | :--- | :--- |
| 1 | tb-size 256, QEMU timeout 4200 s | 13:13:40–13:18:55 UTC: stopped by signal 15 from outside after 5 min (an earlier attempt of this package was ended); still in the initrd. Data not used |
| 2 | as 1 | 13:19:45–about 13:43 UTC: stopped by signal 15 when that attempt's session ended. Reached the real root: `Switching root` at 1245 s, SELinux `enforcing=1`, `Welcome to Fedora Linux 44.20260427.0 (IoT Edition)!` at 1357 s, host name `nmb-i08` from Ignition. Probe not reached |
| 3 | fresh disk, as 1, timeout 5400 s | 13:55:38–14:42:30 UTC, QEMU rc 137: the PC's kernel log says `Out of memory: Killed process … (qemu-system-aar)`; another job on the PC was compiling and the PC's swap was full. Ignition ran all stages (`files` wrote the probe at 1774 s) and the initrd finished (`initrd-cleanup` at 2219 s). Probe not reached |
| 4 | fresh disk, tb-size 64, timeout 7200 s | 15:36:57–16:11:35 UTC: slower still (the PC's load average 10–16 from other jobs, swap full): `Run /init` at 98.1 s, no Ignition stage by 1592 s guest time. Stopped by this package with `kill -TERM` (QEMU rc 0, flock rc 0): at that rate the boot would not have finished in the timeout, and the guest held 0.5 GB of a PC without free memory. Probe not reached |

Figures taken from the serial logs of runs 2 and 3 (TCG with three vCPUs on a loaded PC, **not** a figure for a
Pi 4):

- Kernel `Memory: 271992K/524288K available (… 181632K reserved, 65536K cma-reserved)`; `Freeing initrd memory:
  87784K`; `landlock: Up and running.`; `SELinux:  Initializing.`; `Run /init as init process` at 57.9 s (run 3) /
  69.4 s (run 2).
- First boot runs Ignition in the initrd (`fetch-offline`, `kargs`, `disks`, `mount`, `files`, `umount`), each stage
  slow under TCG (run 3: `fetch-offline` at 1029 s, `files` at 1766 s). After the switch: `Successfully loaded
  SELinux policy in 17.553018s`, `Relabeled /dev/, /dev/shm/, /run/ in 12.663392s` (run 2).
- systemd banner: `systemd 259.5-1.fc44 running in system mode (+PAM +AUDIT +SELINUX -APPARMOR +IMA +IPE +SMACK
  +SECCOMP … +TPM2 … +BPF_FRAMEWORK +BTF …)`; `bpf-restrict-fs: LSM BPF program attached`.

**Not reached in any run:** boot time to `multi-user.target`, free memory and storage after boot, listening ports, running services, the live `modprobe` loop, the tools' own version output and the coexistence tests of section 9. The probe that would have read them is in the work folder (`~/nmb-work/I-08/scripts/nmb-probe.sh`). As far as it came, the image ran in 512 MB without an out-of-memory event inside the guest (runs 2–4); the one OOM kill (run 3) was the PC's.

### 12b. x86_64 installer ISO, KVM

The ISO is an installer; an installed x86_64 system needs an Anaconda run. `test -w /dev/kvm` succeeded, so the
installer ran with **`-accel kvm -cpu host`**, two vCPUs, **512 MB**, under the lock. Kickstart on a separate 8 MiB
FAT image labelled `OEMDRV` (`mkfs.vfat`, `mcopy`): the ISO's own `ostreesetup` line plus `text`, `zerombr`,
`clearpart --all`, `autopart`, `rootpw --lock`, `poweroff` and the same probe in `%post`. The ISO was attached
read-only as an IDE CD; the target was an empty 8 GiB qcow2.

| Run | Settings | Result |
| :--- | :--- | :--- |
| 1 | OVMF (`ovmf 2024.02-2ubuntu0.10`), `-kernel`/`-initrd` from the ISO, `inst.ks=hd:LABEL=OEMDRV:/ks.cfg inst.text` | 14:44:28 UTC start: OVMF booted the CD's own GRUB instead of the given kernel; its default entry failed with `error: ../../grub-core/loader/i386/efi/linux.c:grub_cmd_initrd:260:you need to load the kernel` and the menu came back. Stopped by the 2400 s QEMU timeout (rc 124). The PC's clock jumped about 40 min during this run (the WSL2 VM stalled under memory pressure), so the end time 15:27:07 UTC is not a duration |
| 2 | as 1 but SeaBIOS (no OVMF), so `-kernel` is used | 15:30:44 UTC: kernel `Memory: 186232K/523748K available`, `Trying to unpack rootfs image as initramfs...`, **`Initramfs unpacking failed: write error`**, `Freeing initrd memory: 256692K`, then `Kernel panic - not syncing: VFS: Unable to mount root fs on unknown-block(0,0)` at 9.5 s. Guest stopped with `kill` at 15:33:28 UTC |

- Why: the installer initrd is 262,850,180 bytes xz (`xz -l`: 250.7 MiB compressed, 349.8 MiB uncompressed;
  `xz -dc | wc -c` = 366,817,792). Compressed and unpacked copies together do not fit in 512 MB. Largest parts of
  the initrd (`bsdtar -tvf`, sizes summed per folder): `usr/lib/firmware` 147,822,737 bytes, `usr/lib/modules`
  66,104,789, `usr/lib64/python3.14` 11,466,133 (Anaconda is Python).
- **Result: the x86_64 installer does not fit in a 512 MB guest; no x86_64 system was installed or booted.** The
  installed x86_64 values in this record are those of the commit the installer deploys ([commit], [ISO]).

## 13. Fields not determined

- **x86_64 installed system** (partition layout chosen by Anaconda, `/etc` after install, boot time, memory, ports,
  services): not determined. The installer panics in 512 MB (12b), the limit of this work; a larger guest was not
  run.
- **aarch64 boot figures** (boot time, free memory and storage after boot, listening ports, running services, live
  `modprobe` checks, coexistence tests): not determined. No run reached `multi-user.target` (12a): two ended from outside, one by the PC's OOM killer, one was stopped because the loaded PC made TCG too slow. Run 2 reached the real root after 1245 s (about 21 min) under TCG; the first-boot user space was not measured. A run on an idle PC, or on the Pi 4 with the owner's go-ahead, would fill them.
- **The Pi 4 boot itself** (U-Boot → shim → GRUB on the real board, Wi-Fi, `genet`, the root partition grown on a
  card): not determined; the Pi 4 image is taken apart only (skill rule), and QEMU `virt` uses UEFI firmware, not the
  Pi's.
- **Content hashes of the raw image's ostree objects:** not re-checked (`ostree fsck` needs owners and SELinux
  labels that `debugfs rdump` without root drops, section 1); the commit and its signature were checked.
- **Signatures of the two koji kernel sub-packages** (section 4): not checked; they were only listed.
- **How `policycoreutils-python-utils` entered the compose:** not determined (not in `treefile.json`, required by no
  installed package).
- **Whether removing `firewalld`, `policycoreutils-python-utils` and `setools-console` frees the image of Python**,
  and whether `rpm-ostree install` of the missing tools works offline or online on the device: not tried.
- **Kernel options of the current F44 kernel 7.3.3** (section 7): not read; section 4 is for the image's 6.19.10.
- **Reproducibility of an own compose (criterion 4):** not tried (`rpm-ostree compose tree` with the commit's
  `treefile.json`, or `bootc`/osbuild).
- **Ookla and RTR binaries for glibc (criterion 6):** not part of the image; left to the comparison I-09.
- **Licence sufficiency for redistributing the firmware packages:** the LICENSE tags are listed (section 6); not
  judged.

## 14. Tools on the PC

- Used: `xz` 5.4.5 (`xz -dc | dd conv=sparse`), `fdisk`, `sgdisk`, `file`, `dd`, `mtools` 4.0.43, `debugfs`/`e2fsck`
  1.47.0, `7z` 23.01, `bsdtar` 3.7.2 (libarchive 3.7.2), `ostree` 2024.5, `sqlite3` 3.45.1, `gpg` 2.4.4, `curl`,
  `python3 -I` (the rpm header reader and JSON reading, scripts outside the material), `qemu-system-aarch64` and
  `qemu-system-x86_64` 8.2.2, AAVMF and OVMF 2024.02, `mkfs.vfat`, `qemu-img`, `flock`.
- **No root was used.** The PC has sudo since 2026-11-11 (work status log), but every step here ran as the user:
  partitions cut by offset, ext4 read with `debugfs`, no loop mount.
- `/dev/kvm` was writable for the user (`crw-rw-rw-+ root kvm`, `test -w /dev/kvm` rc 0), so the x86_64 guest used
  KVM; arm64 has no KVM on this PC.
- **Pitfall** (section 2): `7z x` of the ISO gives Joliet names cut to 64 characters, which breaks the ostree
  repository; `bsdtar` reads the Rock Ridge names.
- **Memory on the PC was the limit.** The WSL2 VM has 1,841 MiB and 1 GiB swap; with other jobs running on the PC,
  available memory went down to about 80 MiB, swap was full, and the host OOM killer ended run 3 of the aarch64
  guest.

## 15. Fedora IoT as the NMB base: what the native build adds and what it can drop

The Python version's `setup.sh` (read in `~/src/NetworkMeasureBox`, lines 218–237 and 249–255) handles Fedora with
`dnf install -y --setopt=install_weak_deps=False python3 iproute nftables ethtool pciutils usbutils polkit dnsmasq
ppp curl ca-certificates procps-ng libcap wireguard-tools traceroute iputils mtr iperf3 sqlite gnupg2 gcc make
glibc-devel openssl-devel libcurl-devel pkgconf-pkg-config`, then optional `hostapd iw wireless-regdb openvpn
python3-qrcode qrencode iptables-nft NetworkManager-ppp`, and opens its ports in firewalld when firewalld runs. On
Fedora IoT `/usr` is read-only and packages are layered with `rpm-ostree install` (section 3); whether that
`dnf install` line works on this base was not tried.

On top of the stock commit the native build needs:

- the criterion-3 tools not in the base (section 7): `dnsmasq`, `hostapd`, `ppp`, `wireguard-tools`, `openvpn`;
  `nftables`, `iproute`, `iw`, `wpa_supplicant` are already in it; `iproute-tc` and `ethtool` if `nmbd` uses them;
- firmware the base leaves out, at least `mt7xxx-firmware`/`mediatek-firmware` for MediaTek USB adapters (section
  5), because a read-only image cannot fetch it later (roadmap 6A);
- `kernel-modules-internal` on test images only (for `mac80211_hwsim`);
- its own systemd units for `nmbd`, and NetworkManager, firewalld, ModemManager, `iwd` and `nftables.service` masked
  or removed (section 9);
- an own compose instead of the stock commit if criterion 8 is to be met (section 8), since layering cannot take
  Python out of the base commit (only `rpm-ostree override remove`, untested).

It can drop, compared with the stock commit (each a proposal for an own compose, nothing tested): `firewalld` and
`python3-firewall`; `policycoreutils-python-utils`, `setools-console` and the other `python3-*` packages with
`python3`/`python3-libs`; `NetworkManager` (+ `-wifi`, `-wwan`, `-libnm`); `ModemManager` with `libqmi-utils` and
`libmbim-utils` (if `nmbd` talks to modems itself); `iwd` (`nmbd` uses `wpa_supplicant`); the container stack
(`podman`, `skopeo`, `crun`, `conmon`, `netavark`, `aardvark-dns`, `passt`, `slirp4netns`, `criu`,
`containers-common`, `containernetworking-plugins`); `fdo-client`, `clevis`, `tpm2-tools`, `usbguard`, `fwupd`; `dnf5`;
`screen`, `tmux`; `lvm2` and `xfsprogs` if the layout has no LVM/XFS. The Python version's compilers and headers are
not in the base and not needed by the native build. The A/B update (`rpm-ostree`, `greenboot`, GRUB counter) stays.

## 16. 6A criteria for this candidate

| Criterion | Finding | Where |
| :--- | :--- | :--- |
| 1. Stock kernel, Pi 4 and x86-64, network features | Met by the configs: one Fedora kernel 6.19.10 for both (now 7.3.3 in updates), with bridge, bonding, VLAN, nftables, WireGuard, PPP/PPPoE, cfg80211/mac80211, every tethering driver of 4.27, namespaces, seccomp, Landlock (active LSM), cgroups, watchdog, and the Pi 4 platform (`ARCH_BCM2835`, `PCIE_BRCMSTB`, `BCMGENET`, DTBs, Pi firmware on the ESP). **`mac80211_hwsim` is built but not shipped** (`kernel-modules-internal`). Live: Landlock and SELinux (enforcing) came up in the aarch64 guest; modules not loaded live (12a). | 4, 5 |
| 2. Security updates for years, free | **No: about 13 months per release** (`SUPPORT_END=2027-05-19`; schedule EOL 2027-06-02); free, but a yearly rebase to the next release is needed | 3 |
| 3. Fits nmbd | systemd 259.5, nftables 1.1.6, wpa_supplicant 2.11, iw 6.17, iproute 6.17.0 in the base; **dnsmasq, hostapd, pppd, wireguard-tools, OpenVPN missing** but in the F44 repositories (2.92, 2.11, 2.5.1, 1.0.20260223, 2.7.9), none requiring Python directly | 7 |
| 4. Read-only A/B image, reproducible | **Built in:** read-only `/usr` (composefs), ostree deployments as A/B, signed commits, greenboot health checks with a GRUB boot counter (3 attempts) and rollback; own composes from a treefile are possible (the compose's `treefile.json` is in the commit), reproducibility not tried | 2, 3 |
| 5. Does not fight NMB | **NetworkManager, firewalld and ModemManager are enabled by default**; `nftables.service` is off; started with the shipped `/etc/sysconfig/nftables.conf` (comments only) it loads nothing, but reload and stop flush the whole ruleset, so it must stay masked. The live tests (own nftables table across a firewalld reload, restart and stop; a bridge and a dummy interface across a NetworkManager restart) were not reached (12a): not determined. | 9 |
| 6. Size, memory, glibc | aarch64 card image 6 GiB (root 2.1 GB used, `/boot` 302 MB, ESP 34 MB); OS tree 1.86 GB (aarch64) / 1.38 GB (x86_64); glibc 2.43. Memory after boot not determined (12a); before user space the guest kernel reports 271,992 kB available of 524,288 kB. The x86_64 installer does not fit in 512 MB. Ookla/RTR not checked here | 11, 12 |
| 7. Licence | Free distribution; firmware packages under redistributable-no-modification and Fedora firmware licences, Pi boot files `LicenseRef-Fedora-Firmware` | 6 |
| 8. No Python in the base | **Fails:** `python3 3.14.3` with 13 `python*` packages, pulled in by `firewalld`, `policycoreutils-python-utils` and `setools-console`; the compose itself runs `semanage`. Only an own compose without them could meet it (untested) | 8 |
