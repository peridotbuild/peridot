# All global changes to build and install should follow this line.

# Disable LTO in userspace packages.
%global _lto_cflags %{nil}

# The libexec directory is not used by the linker, so the shared object there
# should not be exported to RPM provides.
%global __provides_exclude_from ^%{_libexecdir}/kselftests

# Disable the find-provides.ksyms script.
%global __provided_ksyms_provides %{nil}

# All global wide changes should be above this line otherwise
# the %%install section will not see them.
%global __spec_install_pre %{___build_pre}

# Kernel has several large (hundreds of mbytes) rpms, they take ~5 mins
# to compress by single-threaded xz. Switch to threaded compression,
# and from level 2 to 3 to keep compressed sizes close to "w2" results.
#
# NB: if default compression in /usr/lib/rpm/redhat/macros ever changes,
# this one might need tweaking (e.g. if default changes to w3.xzdio,
# change below to w4T.xzdio):
%global _binary_payload w3T.xzdio

# Define the version of the Linux Kernel Archive tarball.
%global LKAver {{.Version}}

# Define the buildid, if required.
%global buildid {{.BuildID}}

# Determine the sublevel number and set pkg_version.
%define sublevel %(echo %{LKAver} | %{__awk} -F\. '{ print $3 }')
%if "%{sublevel}" == ""
%global pkg_version %{LKAver}.0
%else
%global pkg_version %{LKAver}
%endif

# Set pkg_release.
%global pkg_release 1%{?buildid}%{?dist}

# Architectures upon which we can sign the kernel
# for secure boot authentication.
%ifarch x86_64 || aarch64
%global signkernel 1
%else
%global signkernel 0
%endif

# Sign modules on all architectures that build modules.
%ifarch x86_64 || aarch64
%global signmodules 1
%else
%global signmodules 0
%endif

# Compress modules on all architectures that build modules.
%ifarch x86_64 || aarch64
%global zipmodules 1
%else
%global zipmodules 0
%endif

%if %{zipmodules}
%global zipsed -e 's/\.ko$/\.ko.xz/'
# For parallel xz processes. Replace with 1 to go back to single process.
%global zcpu `nproc --all`
%endif

# The following build options are enabled by default, but may become disabled
# by later architecture-specific checks. These can also be disabled by using
# --without <opt> in the rpmbuild command, or by forcing these values to 0.
#
# {{.KernelPackage}}
%define with_std          %{?_without_std:          0} %{?!_without_std:          1}
#
# {{.KernelPackage}}-headers
%define with_headers      %{?_without_headers:      0} %{?!_without_headers:      1}
#
# {{.KernelPackage}}-doc
%define with_doc          %{?_without_doc:          0} %{?!_without_doc:          1}
#
# perf
%define with_perf         %{?_without_perf:         0} %{?!_without_perf:         1}
#
# tools
%define with_tools        %{?_without_tools:        0} %{?!_without_tools:        1}
#
# bpf tool
%define with_bpftool      %{?_without_bpftool:      0} %{?!_without_bpftool:      1}
#
# control whether to install the vdso directories
%define with_vdso_install %{?_without_vdso_install: 0} %{?!_without_vdso_install: 1}
#
# Additional option for toracat-friendly, one-off, {{.KernelPackage}} building.
# Only build the base {{.KernelPackage}} (--with baseonly):
%define with_baseonly     %{?_with_baseonly:        1} %{?!_with_baseonly:        0}

%global KVERREL %{pkg_version}-%{pkg_release}.%{_target_cpu}

# If requested, only build base {{.KernelPackage}} package.
%if %{with_baseonly}
%define with_doc 0
%define with_perf 0
%define with_tools 0
%define with_bpftool 0
%define with_vdso_install 0
%endif

%ifarch noarch
%define with_std 0
%define with_headers 0
%define with_perf 0
%define with_tools 0
%define with_bpftool 0
%define with_vdso_install 0
%endif

%ifarch x86_64 || aarch64
%define with_doc 0
### as of {{.KernelPackage}}-6.5.4, no more perf and bpftool -ay
%define with_perf 0
%define with_bpftool 0
%endif

%ifarch x86_64
%define asmarch x86
%define bldarch x86_64
%define hdrarch x86_64
%define make_target bzImage
%define kernel_image arch/x86/boot/bzImage
%endif

%ifarch aarch64
%define asmarch arm64
%define bldarch arm64
%define hdrarch arm64
%define make_target Image.gz
%define kernel_image arch/arm64/boot/Image.gz
%endif

%if %{with_vdso_install}
%define use_vdso 1
%define _use_vdso 1
%else
%define _use_vdso 0
%endif

#
# Packages that need to be installed before the kernel is installed,
# as they will be used by the %%post scripts.
#
%define kernel_ml_prereq  coreutils, systemd >= 203-2, /usr/bin/kernel-install
%define initrd_prereq  dracut >= 027

Name: {{.KernelPackage}}
Summary: The Linux kernel. (The core of any Linux kernel based operating system.)
License: GPLv2 and Redistributable, no modification permitted.
URL: https://www.kernel.org/
Version: %{pkg_version}
Release: %{pkg_release}
ExclusiveArch: x86_64 aarch64 noarch
ExclusiveOS: Linux
Provides: kernel = %{version}-%{release}
Provides: installonlypkg(kernel)
Requires: %{name}-core-uname-r = %{KVERREL}
Requires: %{name}-modules-uname-r = %{KVERREL}

#
# List the packages required for the {{.KernelPackage}} build.
#
BuildRequires: bash, bc, binutils, bison, bzip2, coreutils, diffutils, dwarves, elfutils-devel
BuildRequires: findutils, flex, gawk, gcc, gcc-c++, gcc-plugin-devel, git-core, glibc-static
BuildRequires: gzip, hmaccalc, hostname, kernel-rpm-macros >= 185-9, kmod, m4, make, net-tools
BuildRequires: patch, perl-Carp, perl-devel, perl-generators, perl-interpreter, python3-devel
BuildRequires: redhat-rpm-config, tar, which, xz

%ifarch x86_64 || aarch64
BuildRequires: bpftool, openssl-devel
%endif

%if %{with_headers}
BuildRequires: rsync
%endif

%if %{with_doc}
BuildRequires: asciidoc, python3-sphinx, python3-sphinx_rtd_theme, xmlto
%endif

%if %{with_perf}
BuildRequires: asciidoc, audit-libs-devel, binutils-devel, bison, flex, java-devel
BuildRequires: libbabeltrace-devel, libbpf-devel, libtraceevent-devel, newt-devel
BuildRequires: numactl-devel, openssl-devel, perl(ExtUtils::Embed), xmlto
BuildRequires: xz-devel, zlib-devel
%ifarch aarch64
BuildRequires: opencsd-devel >= 1.0.0
%endif
%endif

%if %{with_tools}
BuildRequires: asciidoc, gettext, libcap-devel, libcap-ng-devel, libnl3-devel
BuildRequires: ncurses-devel, openssl-devel, pciutils-devel
%endif

%if %{with_bpftool}
BuildRequires: binutils-devel, python3-docutils, zlib-devel
%endif

%if %{signkernel} || %{signmodules}
BuildRequires: openssl
%if %{signkernel}
BuildRequires: nss-tools, pesign >= 0.10-4, system-sb-certs
%endif
%endif

BuildConflicts: rhbuildsys(DiskFree) < 500Mb

###
### Sources
###
Source0: https://www.kernel.org/pub/linux/kernel/v6.x/linux-%{LKAver}.tar.xz

Source2: config-x86_64
Source4: config-aarch64

Source20: mod-denylist.sh
Source21: mod-sign.sh
Source23: x509.genkey
Source26: mod-extra.list

Source34: filter-x86_64.sh
Source37: filter-aarch64.sh
Source40: filter-modules.sh

Source100: rockydup1.x509
Source101: rockykpatch1.x509

Source2000: cpupower.service
Source2001: cpupower.config
Source2002: kvm_stat.logrotate

Source3000: secureboot-sig-kernel-x86_64.cer
Source3001: secureboot-sig-kernel-aarch64.cer

%if %{signkernel}
%define secureboot_ca_0 %{_datadir}/pki/sb-certs/secureboot-ca-%{_arch}.cer

%ifarch x86_64
%define secureboot_key_0 %{SOURCE3000}
%endif
%ifarch aarch64
%define secureboot_key_0 %{SOURCE3001}
%endif

%define pesign_name_0 rockybootsigningsigkernelcert
%endif

%description
The %{name} meta package.

#
# This macro does requires, provides, conflicts, obsoletes for a {{.KernelPackage}} package.
#	%%kernel_ml_reqprovconf <subpackage>
# It uses any kernel_ml_<subpackage>_conflicts and kernel_ml_<subpackage>_obsoletes
# macros defined above.
#
%define kernel_ml_reqprovconf \
Provides: %{name} = %{pkg_version}-%{pkg_release}\
Provides: %{name}-%{_target_cpu} = %{pkg_version}-%{pkg_release}%{?1:+%{1}}\
Provides: %{name}-drm-nouveau = 16\
Provides: %{name}-uname-r = %{KVERREL}%{?1:+%{1}}\
Requires(pre): %{kernel_ml_prereq}\
Requires(pre): %{initrd_prereq}\
Requires(pre): ((linux-firmware >= 20150904-56.git6ebf5d57) if linux-firmware)\
Recommends: linux-firmware\
Requires(preun): systemd >= 200\
Conflicts: xfsprogs < 4.3.0-1\
Conflicts: xorg-x11-drv-vmmouse < 13.0.99\
%{expand:%%{?kernel_ml%{?1:_%{1}}_conflicts:Conflicts: %%{%{name}%{?1:_%{1}}_conflicts}}}\
%{expand:%%{?kernel_ml%{?1:_%{1}}_obsoletes:Obsoletes: %%{%{name}%{?1:_%{1}}_obsoletes}}}\
%{expand:%%{?kernel_ml%{?1:_%{1}}_provides:Provides: %%{%{name}%{?1:_%{1}}_provides}}}\
# We can't let RPM do the dependencies automatically because it'll then pick up\
# a correct but undesirable perl dependency from the module headers which\
# isn't required for the kernel proper to function.\
AutoReq: no\
AutoProv: yes\
%{nil}

%package headers
Summary: Header files for the Linux kernel, used by glibc.
Obsoletes: glibc-kernheaders < 3.0-46
Provides: glibc-kernheaders = 3.0-46
%description headers
The Linux kernel headers includes the C header files that specify
the interface between the Linux kernel and userspace libraries and
programs. The header files define structures and constants that are
needed for building most standard programs and are also needed for
rebuilding the glibc package.

%package doc
Summary: Various documentation bits found in the Linux kernel source.
Group: Documentation
%description doc
This package contains documentation files from the Linux kernel
source. Various bits of information about the Linux kernel and the
device drivers shipped with it are documented in these files.

You'll want to install this package if you need a reference to the
options that can be passed to Linux kernel modules at load time.

%if %{with_perf}
%package -n perf
Summary: Performance monitoring for the Linux kernel.
Requires: bzip2
License: GPLv2
%description -n perf
This package contains the perf tool, which enables performance
monitoring of the Linux kernel.

%package -n python3-perf
Summary: Python bindings for apps which will manipulate perf events.
%description -n python3-perf
This package contains a module that permits applications written
in the Python programming language to use the interface to
manipulate perf events.
%endif

%if %{with_tools}
%package -n %{name}-tools
Summary: Assortment of tools for the Linux kernel.
License: GPLv2
Obsoletes: kernel-tools < %{version}
Provides:  kernel-tools = %{version}-%{release}
Obsoletes: cpupowerutils < 1:009-0.6.p1
Provides:  cpupowerutils = 1:009-0.6.p1
Obsoletes: cpufreq-utils < 1:009-0.6.p1
Provides:  cpufreq-utils = 1:009-0.6.p1
Obsoletes: cpufrequtils < 1:009-0.6.p1
Provides:  cpufrequtils = 1:009-0.6.p1
Obsoletes: cpuspeed < 1:1.5-16
Requires: %{name}-tools-libs = %{version}-%{release}
%define __requires_exclude ^%{_bindir}/python
%description -n %{name}-tools
This package contains the tools/ directory from the Linux kernel
source and the supporting documentation.

%package -n %{name}-tools-libs
Summary: Libraries for the %{name}-tools.
License: GPLv2
Obsoletes: kernel-tools-libs < %{version}
Provides:  kernel-tools-libs = %{version}-%{release}
%description -n %{name}-tools-libs
This package contains the libraries built from the tools/ directory
of the Linux kernel source.

%package -n %{name}-tools-libs-devel
Summary: Development files for the %{name}-tools libraries.
License: GPLv2
Obsoletes: kernel-tools-libs-devel < %{version}
Provides:  kernel-tools-libs-devel = %{version}-%{release}
Obsoletes: cpupowerutils-devel < 1:009-0.6.p1
Provides:  cpupowerutils-devel = 1:009-0.6.p1
Provides: %{name}-tools-devel
Requires: %{name}-tools-libs = %{version}-%{release}
Requires: %{name}-tools = %{version}-%{release}
%description -n %{name}-tools-libs-devel
This package contains the development files for the tools/ directory
of the Linux kernel source.
%endif

%if %{with_bpftool}
%package -n bpftool
Summary: Inspection and simple manipulation of eBPF programs and maps.
License: GPLv2
%description -n bpftool
This package contains the bpftool, which allows inspection
and simple manipulation of eBPF programs and maps.
%endif

#
# This macro creates a {{.KernelPackage}}-<subpackage>-devel package.
#	%%kernel_ml_devel_package [-m] <subpackage> <pretty-name>
#
%define kernel_ml_devel_package(m) \
%package %{?1:%{1}-}devel\
Summary: Development package for building %{name} modules to match the %{?2:%{2} }%{name}.\
Provides: %{name}%{?1:-%{1}}-devel-%{_target_cpu} = %{version}-%{release}\
Provides: %{name}-devel-%{_target_cpu} = %{version}-%{release}%{?1:+%{1}}\
Provides: %{name}-devel-uname-r = %{KVERREL}%{?1:+%{1}}\
Provides: installonlypkg(kernel)\
Provides: installonlypkg({{.KernelPackage}})\
AutoReqProv: no\
Requires(pre): findutils\
Requires: findutils\
Requires: perl-interpreter\
Requires: openssl-devel\
Requires: elfutils-libelf-devel\
Requires: bison\
Requires: flex\
Requires: make\
Requires: gcc\
%if %{-m:1}%{!-m:0}\
Requires: %{name}-devel-uname-r = %{KVERREL}\
%endif\
%description %{?1:%{1}-}devel\
This package provides %{name} headers and makefiles sufficient to build modules\
against the %{?2:%{2} }%{name} package.\
%{nil}

#
# This macro creates an empty {{.KernelPackage}}-<subpackage>-devel-matched package that
# requires both the core and devel packages locked on the same version.
#	%%kernel_ml_devel_matched_package [-m] <subpackage> <pretty-name>
#
%define kernel_ml_devel_matched_package(m) \
%package %{?1:%{1}-}devel-matched\
Summary: Meta package to install matching core and devel packages for a given %{?2:%{2} }%{name}.\
Requires: %{name}%{?1:-%{1}}-devel = %{version}-%{release}\
Requires: %{name}%{?1:-%{1}}-core = %{version}-%{release}\
%description %{?1:%{1}-}devel-matched\
This meta package is used to install matching core and devel packages for a given %{?2:%{2} }%{name}.\
%{nil}

#
# This macro creates a {{.KernelPackage}}-<subpackage>-modules-extra package.
#	%%kernel_ml_modules_extra_package [-m] <subpackage> <pretty-name>
#
%define kernel_ml_modules_extra_package(m) \
%package %{?1:%{1}-}modules-extra\
Summary: Extra %{name} modules to match the %{?2:%{2} }%{name}.\
Provides: %{name}%{?1:-%{1}}-modules-extra-%{_target_cpu} = %{version}-%{release}\
Provides: %{name}%{?1:-%{1}}-modules-extra-%{_target_cpu} = %{version}-%{release}%{?1:+%{1}}\
Provides: %{name}%{?1:-%{1}}-modules-extra = %{version}-%{release}%{?1:+%{1}}\
Provides: installonlypkg(kernel-module)\
Provides: installonlypkg({{.KernelPackage}}-module)\
Provides: %{name}%{?1:-%{1}}-modules-extra-uname-r = %{KVERREL}%{?1:+%{1}}\
Requires: %{name}-uname-r = %{KVERREL}%{?1:+%{1}}\
Requires: %{name}%{?1:-%{1}}-modules-uname-r = %{KVERREL}%{?1:+%{1}}\
%if %{-m:1}%{!-m:0}\
Requires: %{name}-modules-extra-uname-r = %{KVERREL}\
%endif\
AutoReq: no\
AutoProv: yes\
%description %{?1:%{1}-}modules-extra\
This package provides less commonly used %{name} modules for the %{?2:%{2} }%{name} package.\
%{nil}

#
# This macro creates a {{.KernelPackage}}-<subpackage>-modules package.
#	%%kernel_ml_modules_package [-m] <subpackage> <pretty-name>
#
%define kernel_ml_modules_package(m) \
%package %{?1:%{1}-}modules\
Summary: %{name} modules to match the %{?2:%{2}-}core %{name}.\
Provides: %{name}%{?1:-%{1}}-modules-%{_target_cpu} = %{version}-%{release}\
Provides: %{name}-modules-%{_target_cpu} = %{version}-%{release}%{?1:+%{1}}\
Provides: %{name}-modules = %{version}-%{release}%{?1:+%{1}}\
Provides: installonlypkg(kernel-module)\
Provides: installonlypkg({{.KernelPackage}}-module)\
Provides: %{name}%{?1:-%{1}}-modules-uname-r = %{KVERREL}%{?1:+%{1}}\
Requires: %{name}-uname-r = %{KVERREL}%{?1:+%{1}}\
%if %{-m:1}%{!-m:0}\
Requires: %{name}-modules-uname-r = %{KVERREL}\
%endif\
AutoReq: no\
AutoProv: yes\
%description %{?1:%{1}-}modules\
This package provides commonly used %{name} modules for the %{?2:%{2}-}core %{name} package.\
%{nil}

#
# this macro creates a {{.KernelPackage}}-<subpackage> meta package.
#	%%kernel_ml_meta_package <subpackage>
#
%define kernel_ml_meta_package() \
%package %{1}\
Summary: %{name} meta-package for the %{1} ${name}.\
Requires: %{name}-%{1}-core-uname-r = %{KVERREL}+%{1}\
Requires: %{name}-%{1}-modules-uname-r = %{KVERREL}+%{1}\
Provides: installonlypkg(kernel)\
Provides: installonlypkg({{.KernelPackage}})\
%description %{1}\
The meta-package for the %{1} %{name}.\
%{nil}

#
# This macro creates a {{.KernelPackage}}-<subpackage> and its -devel.
#	%%define variant_summary The Linux {{.KernelPackage}} compiled for <configuration>
#	%%kernel_ml_variant_package [-n <pretty-name>] [-m] <subpackage>
#
%define kernel_ml_variant_package(n:m) \
%package %{?1:%{1}-}core\
Summary: %{variant_summary}.\
Provides: %{name}-%{?1:%{1}-}core-uname-r = %{KVERREL}%{?1:+%{1}}\
Provides: installonlypkg(kernel)\
Provides: installonlypkg({{.KernelPackage}})\
%if %{-m:1}%{!-m:0}\
Requires: %{name}-core-uname-r = %{KVERREL}\
%endif\
%{expand:%%kernel_ml_reqprovconf}\
%if %{?1:1} %{!?1:0} \
%{expand:%%kernel_ml_meta_package %{?1:%{1}}}\
%endif\
%{expand:%%kernel_ml_devel_package %{?1:%{1}} %{!?{-n}:%{1}}%{?{-n}:%{-n*}} %{-m:%{-m}}}\
%{expand:%%kernel_ml_devel_matched_package %{?1:%{1}} %{!?{-n}:%{1}}%{?{-n}:%{-n*}} %{-m:%{-m}}}\
%{expand:%%kernel_ml_modules_package %{?1:%{1}} %{!?{-n}:%{1}}%{?{-n}:%{-n*}} %{-m:%{-m}}}\
%{expand:%%kernel_ml_modules_extra_package %{?1:%{1}} %{!?{-n}:%{1}}%{?{-n}:%{-n*}} %{-m:%{-m}}}\
%{nil}

# And, finally, the main -core package.

%define variant_summary The Linux kernel.
%kernel_ml_variant_package
%description core
The %{name} package contains the Linux kernel (vmlinuz), the core of any
Linux kernel based operating system. The %{name} package handles the basic
functions of the operating system: memory allocation, process allocation,
device input and output, etc.

# Disable the building of the debug package(s).
%global debug_package %{nil}

# Disable the creation of build_id symbolic links.
%global _build_id_links none

# Set up our "big" %%{make} macro.
%global make %{__make} -s HOSTCFLAGS="%{?build_cflags}" HOSTLDFLAGS="%{?build_ldflags}"

%prep
%ifarch x86_64 || aarch64
%if %{with_baseonly}
%if !%{with_std}
echo "Cannot build --with baseonly as the standard build is currently disabled."
exit 1
%endif
%endif
%endif

%setup -q -n %{name}-%{version} -c
mv linux-%{LKAver} linux-%{KVERREL}

pushd linux-%{KVERREL} > /dev/null

# Purge the source tree of all unrequired dot-files.
find . -name '.*' -type f -delete

# Mangle all Python shebangs to be Python 3 explicitly.
# -i specifies the interpreter for the shebang
# -n prevents creating ~backup files
# -p preserves timestamps
# This fixes errors such as
# *** ERROR: ambiguous python shebang in /usr/bin/kvm_stat: #!/usr/bin/python. Change it to python3 (or python2) explicitly.
# Process all files in the Documentation, scripts and tools directories.
pathfix.py -i "%{__python3} %{py3_shbang_opts}" -n -p \
	tools/kvm/kvm_stat/kvm_stat \
	scripts/show_delta \
	scripts/jobserver-exec \
	scripts/diffconfig \
	scripts/clang-tools \
	scripts/bloat-o-meter \
	tools \
	scripts \
	Documentation \
	2>&1 | grep -Ev 'recursedown|no change'

mv COPYING COPYING-%{version}-%{release}

cp -a %{SOURCE2} .
cp -a %{SOURCE4} .

# Set the EXTRAVERSION string in the top level Makefile.
sed -i "s@^EXTRAVERSION.*@EXTRAVERSION = -%{release}.%{_target_cpu}@" Makefile

%ifarch x86_64 || aarch64
cp config-%{_target_cpu} .config
%{__make} -s ARCH=%{bldarch} listnewconfig | grep -E '^CONFIG_' > newoptions-%{_target_cpu}.txt || true
if [ -s newoptions-%{_target_cpu}.txt ]; then
  # We're just automatically adding new options, think of it as rolling.
  # If something breaks we can disable it explicitly.
  cat newoptions-%{_target_cpu}.txt >> .config
fi
rm -f newoptions-%{_target_cpu}.txt
%endif

# Add DUP and kpatch certificates to system trusted keys for Rocky.
%if %{signkernel} || %{signmodules}
openssl x509 -inform der -in %{SOURCE100} -out rockydup3.pem
openssl x509 -inform der -in %{SOURCE101} -out rockykpatch1.pem
cat rockydup3.pem rockykpatch1.pem > certs/rocky.pem
for i in config-*; do
	sed -i 's@CONFIG_SYSTEM_TRUSTED_KEYS="*"@CONFIG_SYSTEM_TRUSTED_KEYS="certs/rocky.pem"@' $i
done
%else
for i in config-*; do
        sed -i 's@CONFIG_SYSTEM_TRUSTED_KEYS="*"@CONFIG_SYSTEM_TRUSTED_KEYS=""@' $i
done
%endif

# Adjust the FIPS module name for Rocky9.
for i in config-*; do
	sed -i 's@CONFIG_CRYPTO_FIPS_NAME=.*@CONFIG_CRYPTO_FIPS_NAME="Rocky Linux 9 - Kernel Cryptographic API"@' $i
done

%{__make} -s distclean

popd > /dev/null

%build
pushd linux-%{KVERREL} > /dev/null

%ifarch x86_64 || aarch64
cp config-%{_target_cpu} .config

%{__make} -s ARCH=%{bldarch} oldconfig

%if %{signkernel} || %{signmodules}
cp %{SOURCE23} certs/
%endif

%if %{with_std}
%{make} %{?_smp_mflags} ARCH=%{bldarch} %{make_target}

%{make} %{?_smp_mflags} ARCH=%{bldarch} modules || exit 1

%ifarch aarch64
%{make} %{?_smp_mflags} ARCH=%{bldarch} dtbs
%endif

%if %{with_bpftool}
# Generate a vmlinux.h file.
bpftool btf dump file vmlinux format c > tools/bpf/bpftool/vmlinux.h
RPM_VMLINUX_H=vmlinux.h
%endif
%endif

%if %{with_perf}
%ifarch aarch64
%global perf_build_extra_opts CORESIGHT=1
%endif

%global perf_make \
	%{__make} -s -C tools/perf NO_PERF_READ_VDSO32=1 NO_PERF_READ_VDSOX32=1 WERROR=0 NO_LIBUNWIND=1 HAVE_CPLUS_DEMANGLE=1 NO_GTK2=1 NO_STRLCPY=1 NO_BIONIC=1 LIBTRACEEVENT_DYNAMIC=1 %{?perf_build_extra_opts} prefix=%{_prefix} PYTHON=%{__python3}

# Make sure that check-headers.sh is executable.
chmod +x tools/perf/check-headers.sh

%{perf_make} all
%endif

%if %{with_tools}
# Make sure that version-gen.sh is executable.
chmod +x tools/power/cpupower/utils/version-gen.sh

pushd tools/power/cpupower > /dev/null
%{__make} -s %{?_smp_mflags} CPUFREQ_BENCH=false DEBUG=false
popd > /dev/null

%ifarch x86_64
pushd tools/power/cpupower/debug/x86_64 > /dev/null
%{__make} -s %{?_smp_mflags} centrino-decode powernow-k8-decode
popd > /dev/null

pushd tools/power/x86/x86_energy_perf_policy > /dev/null
%{__make} -s %{?_smp_mflags}
popd > /dev/null

pushd tools/power/x86/turbostat > /dev/null
%{__make} -s %{?_smp_mflags}
popd > /dev/null

pushd tools/power/x86/intel-speed-select > /dev/null
%{__make} -s %{?_smp_mflags}
popd > /dev/null
%endif

pushd tools/thermal/tmon > /dev/null
%{__make} -s %{?_smp_mflags}
popd > /dev/null

pushd tools/iio > /dev/null
%{__make} -s %{?_smp_mflags}
popd > /dev/null

pushd tools/gpio > /dev/null
%{__make} -s %{?_smp_mflags}
popd > /dev/null

### BCAT
%if 0
pushd tools/vm > /dev/null
%{__make} -s %{?_smp_mflags} slabinfo page_owner_sort
popd > /dev/null
%endif
### BCAT
%endif

%if %{with_bpftool}
%global bpftool_make \
	%{__make} -s EXTRA_CFLAGS="${RPM_OPT_FLAGS}" EXTRA_LDFLAGS="%{__global_ldflags}" DESTDIR=$RPM_BUILD_ROOT VMLINUX_H="${RPM_VMLINUX_H}"

pushd tools/bpf/bpftool > /dev/null
%{bpftool_make}
popd > /dev/null
%endif
%endif

popd > /dev/null

%install
%define __modsign_install_post \
pushd linux-%{KVERREL} > /dev/null \
if [ "%{signmodules}" -eq "1" ]; then \
	if [ "%{with_std}" -ne "0" ]; then \
		%{SOURCE21} certs/signing_key.pem.sign certs/signing_key.x509.sign $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/ \
	fi \
fi \
if [ "%{zipmodules}" -eq "1" ]; then \
	find $RPM_BUILD_ROOT/lib/modules/ -name '*.ko' -type f | xargs --no-run-if-empty -P%{zcpu} xz \
fi \
%{nil}

#
# Ensure modules are signed *after* all invocations of
# strip have occured, which are in __os_install_post.
#
%define __spec_install_post \
	%{__arch_install_post}\
	%{__os_install_post}\
	%{__modsign_install_post}

pushd linux-%{KVERREL} > /dev/null

rm -fr $RPM_BUILD_ROOT

%ifarch x86_64 || aarch64
mkdir -p $RPM_BUILD_ROOT

%if %{with_std}
mkdir -p $RPM_BUILD_ROOT/boot
mkdir -p $RPM_BUILD_ROOT%{_libexecdir}
mkdir -p $RPM_BUILD_ROOT/lib/modules/%{KVERREL}
mkdir -p $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/systemtap

%ifarch aarch64
%{make} ARCH=%{bldarch} dtbs_install INSTALL_DTBS_PATH=$RPM_BUILD_ROOT/boot/dtb-%{KVERREL}
cp -r $RPM_BUILD_ROOT/boot/dtb-%{KVERREL} $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/dtb
find arch/%{bldarch}/boot/dts -name '*.dtb' -type f -delete
%endif

# Install the results within the RPM_BUILD_ROOT directory.
install -m 644 .config $RPM_BUILD_ROOT/boot/config-%{KVERREL}
install -m 644 .config $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/config
install -m 644 System.map $RPM_BUILD_ROOT/boot/System.map-%{KVERREL}
install -m 644 System.map $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/System.map

# We estimate the size of the initramfs because rpm needs to take this size
# into consideration when performing disk space calculations. (See bz #530778)
dd if=/dev/zero of=$RPM_BUILD_ROOT/boot/initramfs-%{KVERREL}.img bs=1M count=20

%if %{signkernel}
# Sign the kernel image if we're using EFI.
# aarch64 kernels are gziped EFI images.
%ifarch x86_64
SignImage=arch/x86/boot/bzImage
%endif

%ifarch aarch64
SignImage=arch/arm64/boot/Image
%endif

%pesign -s -i $SignImage -o vmlinuz.signed -a %{secureboot_ca_0} -c %{secureboot_key_0} -n %{pesign_name_0}

if [ ! -s vmlinuz.signed ]; then
	echo "pesigning failed"
	exit 1
fi

mv vmlinuz.signed $SignImage

%ifarch aarch64
gzip -f9 $SignImage
%endif
%endif

cp %{kernel_image} $RPM_BUILD_ROOT/boot/vmlinuz-%{KVERREL}
chmod 755 $RPM_BUILD_ROOT/boot/vmlinuz-%{KVERREL}
cp $RPM_BUILD_ROOT/boot/vmlinuz-%{KVERREL} $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/vmlinuz

sha512hmac $RPM_BUILD_ROOT/boot/vmlinuz-%{KVERREL} | sed -e "s,$RPM_BUILD_ROOT,," > $RPM_BUILD_ROOT/boot/.vmlinuz-%{KVERREL}.hmac
cp $RPM_BUILD_ROOT/boot/.vmlinuz-%{KVERREL}.hmac $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/.vmlinuz.hmac

# Override mod-fw because we don't want it to install any firmware.
# We'll get it from the linux-firmware package and we don't want conflicts.
%{make} %{?_smp_mflags} ARCH=%{bldarch} INSTALL_MOD_PATH=$RPM_BUILD_ROOT modules_install KERNELRELEASE=%{KVERREL} mod-fw=

# Add a noop %%defattr statement because rpm doesn't like empty file list files.
echo '%%defattr(-,-,-)' > ../%{name}-ldsoconf.list

%if %{with_vdso_install}
%{make} %{?_smp_mflags} ARCH=%{bldarch} INSTALL_MOD_PATH=$RPM_BUILD_ROOT vdso_install KERNELRELEASE=%{KVERREL}

if [ -s ldconfig-%{name}.conf ]; then
	install -D -m 444 ldconfig-%{name}.conf $RPM_BUILD_ROOT/etc/ld.so.conf.d/%{name}-%{KVERREL}.conf
	echo /etc/ld.so.conf.d/%{name}-%{KVERREL}.conf >> ../%{name}-ldsoconf.list
fi
%endif

#
# This looks scary but the end result is supposed to be:
#
# - all arch relevant include/ files.
# - all Makefile and Kconfig files.
# - all script/ files.
#
rm -f $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
rm -f $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/source
mkdir -p $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build

pushd $RPM_BUILD_ROOT/lib/modules/%{KVERREL} > /dev/null
ln -s build source
popd > /dev/null

mkdir -p $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/updates
mkdir -p $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/weak-updates

# CONFIG_KERNEL_HEADER_TEST generates some extra files during testing so just delete them.
find . -name *.h.s -delete

# First copy everything . . .
cp --parents `find  -type f -name "Makefile*" -o -name "Kconfig*"` $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build

if [ ! -e Module.symvers ]; then
	touch Module.symvers
fi

cp Module.symvers $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp System.map $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build

if [ -s Module.markers ]; then
	cp Module.markers $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
fi

gzip -c9 < Module.symvers > $RPM_BUILD_ROOT/boot/symvers-%{KVERREL}.gz
cp $RPM_BUILD_ROOT/boot/symvers-%{KVERREL}.gz $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/symvers.gz

# . . . then drop all but the needed Makefiles and Kconfig files.
rm -fr $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/scripts
rm -fr $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/include
cp .config $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a scripts $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
rm -fr $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/scripts/tracing
rm -f $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/scripts/spdxcheck.py

# Files for 'make scripts' to succeed with {{.KernelPackage}}-devel.
mkdir -p $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/security/selinux/include
cp -a --parents security/selinux/include/classmap.h $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents security/selinux/include/initial_sid_to_string.h $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
mkdir -p $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/tools/include/tools
cp -a --parents tools/include/tools/be_byteshift.h $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/include/tools/le_byteshift.h $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build

# Files for 'make prepare' to succeed with {{.KernelPackage}}-devel.
cp -a --parents tools/include/linux/compiler* $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/include/linux/types.h $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/build/Build.include $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp --parents tools/build/Build $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp --parents tools/build/fixdep.c $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp --parents tools/objtool/sync-check.sh $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/bpf/resolve_btfids $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build

cp --parents security/selinux/include/policycap_names.h $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp --parents security/selinux/include/policycap.h $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build

cp -a --parents tools/include/asm-generic $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/include/linux $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/include/uapi/asm $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/include/uapi/asm-generic $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/include/uapi/linux $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/include/vdso $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp --parents tools/scripts/utilities.mak $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/lib/subcmd $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp --parents tools/lib/*.c $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp --parents tools/objtool/*.[ch] $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp --parents tools/objtool/Build $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp --parents tools/objtool/include/objtool/*.h $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/lib/bpf $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp --parents tools/lib/bpf/Build $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build

if [ -f tools/objtool/objtool ]; then
	cp -a tools/objtool/objtool $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/tools/objtool/ || :
fi
if [ -f tools/objtool/fixdep ]; then
	cp -a tools/objtool/fixdep $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/tools/objtool/ || :
fi
if [ -d arch/%{bldarch}/scripts ]; then
	cp -a arch/%{bldarch}/scripts $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/arch/%{_arch} || :
fi
if [ -f arch/%{bldarch}/*lds ]; then
	cp -a arch/%{bldarch}/*lds $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/arch/%{_arch}/ || :
fi
if [ -f arch/%{asmarch}/kernel/module.lds ]; then
	cp -a --parents arch/%{asmarch}/kernel/module.lds $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
fi

find $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/scripts \( -iname "*.o" -o -iname "*.cmd" \) -exec rm -f {} +

if [ -d arch/%{asmarch}/include ]; then
	cp -a --parents arch/%{asmarch}/include $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
fi

%ifarch aarch64
# arch/arm64/include/asm/xen references arch/arm
cp -a --parents arch/arm/include/asm/xen $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
# arch/arm64/include/asm/opcodes.h references arch/arm
cp -a --parents arch/arm/include/asm/opcodes.h $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
%endif

cp -a include $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/include

%ifarch x86_64
# Files required for 'make prepare' to succeed with {{.KernelPackage}}-devel.
cp -a --parents arch/x86/entry/syscalls/syscall_32.tbl $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/entry/syscalls/syscall_64.tbl $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/tools/relocs_32.c $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/tools/relocs_64.c $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/tools/relocs.c $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/tools/relocs_common.c $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/tools/relocs.h $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/purgatory/purgatory.c $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/purgatory/stack.S $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/purgatory/setup-x86_64.S $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/purgatory/entry64.S $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/boot/string.h $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/boot/string.c $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents arch/x86/boot/ctype.h $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/

cp -a --parents scripts/syscalltbl.sh $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/
cp -a --parents scripts/syscallhdr.sh $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/

cp -a --parents tools/arch/x86/include/asm $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/arch/x86/include/uapi/asm $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/objtool/arch/x86/lib $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/arch/x86/lib/ $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/arch/x86/tools/gen-insn-attr-x86.awk $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
cp -a --parents tools/objtool/arch/x86/ $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build
%endif

# Clean up the intermediate tools files.
find $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/tools \( -iname "*.o" -o -iname "*.cmd" \) -exec rm -f {} +

# Make sure that the Makefile and the version.h file have a matching timestamp
# so that external modules can be built.
touch -r $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/Makefile \
	$RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build/include/generated/uapi/linux/version.h

find $RPM_BUILD_ROOT/lib/modules/%{KVERREL} -name "*.ko" -type f > modnames

# Mark the modules executable, so that strip-to-file can strip them.
xargs --no-run-if-empty chmod u+x < modnames

# Generate a list of modules for block and networking.
grep -F /drivers/ modnames | xargs --no-run-if-empty nm -upA | \
	sed -n 's,^.*/\([^/]*\.ko\):  *U \(.*\)$,\1 \2,p' > drivers.undef

collect_modules_list()
{
	sed -r -n -e "s/^([^ ]+) \\.?($2)\$/\\1/p" drivers.undef | \
		LC_ALL=C sort -u > $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/modules.$1

	if [ ! -z "$3" ]; then
		sed -r -e "/^($3)\$/d" -i $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/modules.$1
	fi
}

collect_modules_list networking \
    'register_netdev|ieee80211_register_hw|usbnet_probe|phy_driver_register|rt(l_|2x00)(pci|usb)_probe|register_netdevice'

collect_modules_list block \
    'ata_scsi_ioctl|scsi_add_host|scsi_add_host_with_dma|blk_alloc_queue|blk_init_queue|register_mtd_blktrans|scsi_esp_register|scsi_register_device_handler|blk_queue_physical_block_size' 'pktcdvd.ko|dm-mod.ko'

collect_modules_list drm \
    'drm_open|drm_init'

collect_modules_list modesetting \
    'drm_crtc_init'

# Detect any missing or incorrect license tags.
( find $RPM_BUILD_ROOT/lib/modules/%{KVERREL} -name '*.ko' -type f | xargs --no-run-if-empty /sbin/modinfo -l | \
	grep -E -v 'GPL( v2)?$|Dual BSD/GPL$|Dual MPL/GPL$|GPL and additional rights$' ) && exit 1

remove_depmod_files()
{
	# Remove all the files that will be auto generated by depmod at the kernel install time.
	pushd $RPM_BUILD_ROOT/lib/modules/%{KVERREL} > /dev/null
	rm -f modules.{alias,alias.bin,builtin.alias.bin,builtin.bin} \
		modules.{dep,dep.bin,devname,softdep,symbols,symbols.bin}
	popd > /dev/null
}

remove_depmod_files

# Identify modules in the {{.KernelPackage}}-modules-extras package
%{SOURCE20} $RPM_BUILD_ROOT lib/modules/%{KVERREL} %{SOURCE26}

#
# Generate the {{.KernelPackage}}-core and {{.KernelPackage}}-modules file lists.
#

# Make a copy of the System.map file for depmod to use.
cp System.map $RPM_BUILD_ROOT/

pushd $RPM_BUILD_ROOT > /dev/null

# Create a backup of the full module tree so it can be
# restored after the filtering has been completed.
mkdir restore
cp -r lib/modules/%{KVERREL}/* restore/

# Don't include anything going into {{.KernelPackage}}-modules-extra in the file lists.
xargs rm -fr < mod-extra.list

# Find all the module files and filter them out into the core and modules lists.
# This actually removes anything going into {{.KernelPackage}}-modules from the directory.
find lib/modules/%{KVERREL}/kernel -name *.ko -type f | sort -n > modules.list
cp $RPM_SOURCE_DIR/filter-*.sh .
./filter-modules.sh modules.list %{_target_cpu}
rm -f filter-*.sh

### BCAT
%if 0
# Run depmod on the resulting module tree to make sure that the tree isn't broken.
depmod -b . -aeF ./System.map %{KVERREL} &> depmod.out
if [ -s depmod.out ]; then
	echo "Depmod failure"
	cat depmod.out
	exit 1
else
	rm -f depmod.out
fi

remove_depmod_files
%endif
### BCAT

# Go back and find all of the various directories in the tree.
# We use this for the directory lists in {{.KernelPackage}}-core.
find lib/modules/%{KVERREL}/kernel -mindepth 1 -type d | sort -n > module-dirs.list

# Cleanup.
rm -f System.map
cp -r restore/* lib/modules/%{KVERREL}/
rm -fr restore

popd > /dev/null

# Make sure that the files lists start with absolute paths or rpmbuild fails.
# Also add in the directory entries.
sed -e 's/^lib*/\/lib/' %{?zipsed} $RPM_BUILD_ROOT/k-d.list > ../%{name}-modules.list
sed -e 's/^lib*/%dir \/lib/' %{?zipsed} $RPM_BUILD_ROOT/module-dirs.list > ../%{name}-core.list
sed -e 's/^lib*/\/lib/' %{?zipsed} $RPM_BUILD_ROOT/modules.list >> ../%{name}-core.list
sed -e 's/^lib*/\/lib/' %{?zipsed} $RPM_BUILD_ROOT/mod-extra.list >> ../%{name}-modules-extra.list

# Cleanup.
rm -f $RPM_BUILD_ROOT/k-d.list
rm -f $RPM_BUILD_ROOT/module-dirs.list
rm -f $RPM_BUILD_ROOT/modules.list
rm -f $RPM_BUILD_ROOT/mod-extra.list

# Move the development files out of the /lib/modules/ file system.
mkdir -p $RPM_BUILD_ROOT/usr/src/kernels
mv $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build $RPM_BUILD_ROOT/usr/src/kernels/%{KVERREL}

# This is going to create a broken link during the build but we don't use
# it after this point. We need the link to actually point to something
# for when the {{.KernelPackage}}-devel package is installed.
ln -sf /usr/src/kernels/%{KVERREL} $RPM_BUILD_ROOT/lib/modules/%{KVERREL}/build

# Move the generated vmlinux.h file into the {{.KernelPackage}}-devel directory structure.
if [ -f tools/bpf/bpftool/vmlinux.h ]; then
	mv tools/bpf/bpftool/vmlinux.h $RPM_BUILD_ROOT/usr/src/kernels/%{KVERREL}/
fi

# Purge the {{.KernelPackage}}-devel tree of leftover junk.
find $RPM_BUILD_ROOT/usr/src/kernels -name ".*.cmd" -type f -delete

# Red Hat UEFI Secure Boot CA certificate, which can be used to authenticate the kernel.
mkdir -p $RPM_BUILD_ROOT%{_datadir}/doc/%{name}-keys/%{KVERREL}
%if %{signkernel}
install -m 0644 %{secureboot_ca_0} $RPM_BUILD_ROOT%{_datadir}/doc/%{name}-keys/%{KVERREL}/kernel-signing-ca.cer
%endif

%if %{signmodules}
# Save the signing keys so that we can sign the modules in __modsign_install_post.
cp certs/signing_key.pem certs/signing_key.pem.sign
cp certs/signing_key.x509 certs/signing_key.x509.sign
%endif
%endif

# We have to do the headers install before the tools install because the
# {{.KernelPackage}} headers_install will remove any header files in /usr/include that
# it doesn't install itself.

%if %{with_headers}
# Install {{.KernelPackage}} headers
%{__make} -s ARCH=%{hdrarch} INSTALL_HDR_PATH=$RPM_BUILD_ROOT/usr headers_install

find $RPM_BUILD_ROOT/usr/include \
  \( -name .install -o -name .check -o \
     -name ..install.cmd -o -name ..check.cmd \) -delete
%endif

%if %{with_perf}
# perf tool binary and supporting scripts/binaries
%{perf_make} DESTDIR=$RPM_BUILD_ROOT lib=%{_lib} install-bin
# Remove the 'trace' symlink.
rm -f $RPM_BUILD_ROOT%{_bindir}/trace

# For both of the below, yes, this should be using a macro but right now
# it's hard coded and we don't actually want it anyway.
# Remove examples.
rm -fr $RPM_BUILD_ROOT/usr/lib/perf/examples
rm -fr $RPM_BUILD_ROOT/usr/lib/perf/include

# python-perf extension
%{perf_make} DESTDIR=$RPM_BUILD_ROOT install-python_ext

# perf man pages (note: implicit rpm magic compresses them later)
mkdir -p $RPM_BUILD_ROOT%{_mandir}/man1
%{perf_make} DESTDIR=$RPM_BUILD_ROOT install-man

# Remove any tracevent files, eg. its plugins still gets built and installed,
# even if we build against system's libtracevent during perf build (by setting
# LIBTRACEEVENT_DYNAMIC=1 above in perf_make macro). Those files should already
# ship with libtraceevent package.
rm -fr $RPM_BUILD_ROOT%{_libdir}/traceevent
%endif

%if %{with_tools}
%{__make} -s -C tools/power/cpupower DESTDIR=$RPM_BUILD_ROOT libdir=%{_libdir} mandir=%{_mandir} CPUFREQ_BENCH=false install

rm -f $RPM_BUILD_ROOT%{_libdir}/*.{a,la}

%find_lang cpupower
mv cpupower.lang ../

%ifarch x86_64
pushd tools/power/cpupower/debug/x86_64 > /dev/null
install -m755 centrino-decode $RPM_BUILD_ROOT%{_bindir}/centrino-decode
install -m755 powernow-k8-decode $RPM_BUILD_ROOT%{_bindir}/powernow-k8-decode
popd > /dev/null
%endif

chmod 0755 $RPM_BUILD_ROOT%{_libdir}/libcpupower.so*
mkdir -p $RPM_BUILD_ROOT%{_unitdir} $RPM_BUILD_ROOT%{_sysconfdir}/sysconfig
install -m644 %{SOURCE2000} $RPM_BUILD_ROOT%{_unitdir}/cpupower.service
install -m644 %{SOURCE2001} $RPM_BUILD_ROOT%{_sysconfdir}/sysconfig/cpupower

%ifarch x86_64
mkdir -p $RPM_BUILD_ROOT%{_mandir}/man8
pushd tools/power/x86/x86_energy_perf_policy > /dev/null
%{__make} -s %{?_smp_mflags} DESTDIR=$RPM_BUILD_ROOT install
popd > /dev/null

pushd tools/power/x86/turbostat > /dev/null
%{__make} -s %{?_smp_mflags} DESTDIR=$RPM_BUILD_ROOT install
popd > /dev/null

pushd tools/power/x86/intel-speed-select > /dev/null
%{__make} -s %{?_smp_mflags} DESTDIR=$RPM_BUILD_ROOT install
popd > /dev/null
%endif

pushd tools/thermal/tmon > /dev/null
%{__make} -s %{?_smp_mflags} INSTALL_ROOT=$RPM_BUILD_ROOT install
popd > /dev/null

pushd tools/iio > /dev/null
%{__make} -s %{?_smp_mflags} DESTDIR=$RPM_BUILD_ROOT install
popd > /dev/null

pushd tools/gpio > /dev/null
%{__make} -s %{?_smp_mflags} DESTDIR=$RPM_BUILD_ROOT install
popd > /dev/null

install -m644 -D %{SOURCE2002} $RPM_BUILD_ROOT%{_sysconfdir}/logrotate.d/kvm_stat

pushd tools/kvm/kvm_stat > /dev/null
%{__make} -s INSTALL_ROOT=$RPM_BUILD_ROOT install-tools
%{__make} -s INSTALL_ROOT=$RPM_BUILD_ROOT install-man
install -m644 -D kvm_stat.service $RPM_BUILD_ROOT%{_unitdir}/kvm_stat.service
popd > /dev/null

### BCAT
%if 0
pushd tools/vm > /dev/null
install -m755 slabinfo $RPM_BUILD_ROOT%{_bindir}/slabinfo
install -m755 page_owner_sort $RPM_BUILD_ROOT%{_bindir}/page_owner_sort
popd > /dev/null
%endif
### BCAT
%endif

%if %{with_bpftool}
pushd tools/bpf/bpftool > /dev/null
%{bpftool_make} prefix=%{_prefix} bash_compdir=%{_sysconfdir}/bash_completion.d/ mandir=%{_mandir} install doc-install
popd > /dev/null
%endif
%endif

%ifarch noarch
mkdir -p $RPM_BUILD_ROOT

%if %{with_doc}
# Sometimes non-world-readable files sneak into the kernel source tree.
chmod -R a=rX Documentation
find Documentation -type d | xargs --no-run-if-empty chmod u+w

DocDir=$RPM_BUILD_ROOT%{_datadir}/doc/%{name}-doc-%{version}-%{release}

# Copy the source over.
mkdir -p $DocDir
tar -h -f - --exclude=man --exclude='.*' -c Documentation | tar xf - -C $DocDir
%endif
%endif

popd > /dev/null

###
### Scripts.
###
%if %{with_tools}
%post -n %{name}-tools-libs
/sbin/ldconfig

%postun -n %{name}-tools-libs
/sbin/ldconfig
%endif

#
# This macro defines a %%post script for a {{.KernelPackage}}*-devel package.
#	%%kernel_ml_devel_post [<subpackage>]
# Note we don't run hardlink if ostree is in use, as ostree is
# a far more sophisticated hardlink implementation.
# https://github.com/projectatomic/rpm-ostree/commit/58a79056a889be8814aa51f507b2c7a4dccee526
#
%define kernel_ml_devel_post() \
%{expand:%%post %{?1:%{1}-}devel}\
if [ -f /etc/sysconfig/kernel ]\
then\
    . /etc/sysconfig/kernel || exit $?\
fi\
if [ "$HARDLINK" != "no" -a -x /usr/bin/hardlink -a ! -e /run/ostree-booted ] \
then\
    (cd /usr/src/kernels/%{KVERREL}%{?1:+%{1}} &&\
        /usr/bin/find . -type f | while read f; do\
          hardlink -c /usr/src/kernels/*%{?dist}.*/$f $f > /dev/null\
        done)\
fi\
%{nil}

#
# This macro defines a %%post script for a {{.KernelPackage}}*-modules-extra package.
# It also defines a %%postun script that does the same thing.
#	%%kernel_ml_modules_extra_post [<subpackage>]
#
%define kernel_ml_modules_extra_post() \
%{expand:%%post %{?1:%{1}-}modules-extra}\
/sbin/depmod -a %{KVERREL}%{?1:+%{1}}\
%{nil}\
%{expand:%%postun %{?1:%{1}-}modules-extra}\
/sbin/depmod -a %{KVERREL}%{?1:+%{1}}\
%{nil}

#
# This macro defines a %%post script for a {{.KernelPackage}}*-modules package.
# It also defines a %%postun script that does the same thing.
#	%%kernel_ml_modules_post [<subpackage>]
#
%define kernel_ml_modules_post() \
%{expand:%%post %{?1:%{1}-}modules}\
/sbin/depmod -a %{KVERREL}%{?1:+%{1}}\
%{nil}\
%{expand:%%postun %{?1:%{1}-}modules}\
/sbin/depmod -a %{KVERREL}%{?1:+%{1}}\
%{nil}

# This macro defines a %%posttrans script for a {{.KernelPackage}} package.
#	%%kernel_ml_variant_posttrans [<subpackage>]
# More text can follow to go at the end of this variant's %%post.
#
%define kernel_ml_variant_posttrans() \
%{expand:%%posttrans %{?1:%{1}-}core}\
if [ -x %{_sbindir}/weak-modules ]\
then\
    %{_sbindir}/weak-modules --add-kernel %{KVERREL}%{?1:+%{1}} || exit $?\
fi\
/bin/kernel-install add %{KVERREL}%{?1:+%{1}} /lib/modules/%{KVERREL}%{?1:+%{1}}/vmlinuz || exit $?\
%{nil}

#
# This macro defines a %%post script for a {{.KernelPackage}} package and its devel package.
#	%%kernel_ml_variant_post [-v <subpackage>] [-r <replace>]
# More text can follow to go at the end of this variant's %%post.
#
%define kernel_ml_variant_post(v:r:) \
%{expand:%%kernel_ml_devel_post %{?-v*}}\
%{expand:%%kernel_ml_modules_post %{?-v*}}\
%{expand:%%kernel_ml_modules_extra_post %{?-v*}}\
%{expand:%%kernel_ml_variant_posttrans %{?-v*}}\
%{expand:%%post %{?-v*:%{-v*}-}core}\
%{-r:\
if [ `uname -i` == "x86_64" ] &&\
    [ -f /etc/sysconfig/kernel ]; then\
    /bin/sed -r -i -e 's/^DEFAULTKERNEL=%{-r*}$/DEFAULTKERNEL=%{name}%{?-v:-%{-v*}}/' /etc/sysconfig/kernel || exit $?\
fi}\
%{nil}

#
# This macro defines a %%preun script for a {{.KernelPackage}} package.
#	%%kernel_ml_variant_preun <subpackage>
#
%define kernel_ml_variant_preun() \
%{expand:%%preun %{?1:%{1}-}core}\
/bin/kernel-install remove %{KVERREL}%{?1:+%{1}} /lib/modules/%{KVERREL}%{?1:+%{1}}/vmlinuz || exit $?\
if [ -x %{_sbindir}/weak-modules ]\
then\
    %{_sbindir}/weak-modules --remove-kernel %{KVERREL}%{?1:+%{1}} || exit $?\
fi\
%{nil}

%kernel_ml_variant_preun
%kernel_ml_variant_post -r kernel-smp

if [ -x /sbin/ldconfig ]
then
    /sbin/ldconfig -X || exit $?
fi

###
### File lists.
###
%if %{with_headers}
%files headers
/usr/include/*
%endif

%if %{with_doc}
%files doc
%defattr(-,root,root)
%{_datadir}/doc/%{name}-doc-%{version}-%{release}/Documentation/*
%dir %{_datadir}/doc/%{name}-doc-%{version}-%{release}/Documentation
%dir %{_datadir}/doc/%{name}-doc-%{version}-%{release}
%endif

%if %{with_perf}
%files -n perf
%{_bindir}/perf
%{_libdir}/libperf-jvmti.so
%dir %{_libexecdir}/perf-core
%{_libexecdir}/perf-core/*
%{_datadir}/perf-core/*
%{_mandir}/man[1-8]/perf*
%{_sysconfdir}/bash_completion.d/perf
%doc linux-%{KVERREL}/tools/perf/Documentation/examples.txt
%{_docdir}/perf-tip/tips.txt

%files -n python3-perf
%{python3_sitearch}/*
%endif

%if %{with_tools}
%files -n %{name}-tools -f cpupower.lang
%{_bindir}/cpupower
%{_datadir}/bash-completion/completions/cpupower
%ifarch x86_64
%{_bindir}/centrino-decode
%{_bindir}/powernow-k8-decode
%endif
%{_unitdir}/cpupower.service
%{_mandir}/man[1-8]/cpupower*
%config(noreplace) %{_sysconfdir}/sysconfig/cpupower
%ifarch x86_64
%{_bindir}/x86_energy_perf_policy
%{_mandir}/man8/x86_energy_perf_policy*
%{_bindir}/turbostat
%{_mandir}/man8/turbostat*
%{_bindir}/intel-speed-select
%endif
%{_bindir}/tmon
%{_bindir}/iio_event_monitor
%{_bindir}/iio_generic_buffer
%{_bindir}/lsiio
%{_bindir}/lsgpio
%{_bindir}/gpio-hammer
%{_bindir}/gpio-event-mon
%{_bindir}/gpio-watch
%{_mandir}/man1/kvm_stat*
%{_bindir}/kvm_stat
%{_unitdir}/kvm_stat.service
%config(noreplace) %{_sysconfdir}/logrotate.d/kvm_stat
### BCAT
%if 0
%{_bindir}/page_owner_sort
%{_bindir}/slabinfo
%endif
### BCAT

%files -n %{name}-tools-libs
{{if (eq (compareVersion .Version "6.2") 1)}}
%{_libdir}/libcpupower.so.1
{{else}}
%{_libdir}/libcpupower.so.0
{{end}}
%{_libdir}/libcpupower.so.0.0.1

%files -n %{name}-tools-libs-devel
%{_libdir}/libcpupower.so
%endif

%if %{with_bpftool}
%files -n bpftool
%{_sbindir}/bpftool
%{_sysconfdir}/bash_completion.d/bpftool
%{_mandir}/man8/bpftool-cgroup.8.gz
%{_mandir}/man8/bpftool-gen.8.gz
%{_mandir}/man8/bpftool-iter.8.gz
%{_mandir}/man8/bpftool-link.8.gz
%{_mandir}/man8/bpftool-map.8.gz
%{_mandir}/man8/bpftool-prog.8.gz
%{_mandir}/man8/bpftool-perf.8.gz
%{_mandir}/man8/bpftool.8.gz
%{_mandir}/man8/bpftool-net.8.gz
%{_mandir}/man8/bpftool-feature.8.gz
%{_mandir}/man8/bpftool-btf.8.gz
%{_mandir}/man8/bpftool-struct_ops.8.gz
%endif

# Empty meta-package.
%ifarch x86_64 || aarch64
%files
%endif

#
# This macro defines the %%files sections for a {{.KernelPackage}} package
# and its devel package.
#	%%kernel_ml_variant_files [-k vmlinux] <use_vdso> <condition> <subpackage>
#
%define kernel_ml_variant_files(k:) \
%if %{2}\
%{expand:%%files -f %{name}-%{?3:%{3}-}core.list %{?1:-f %{name}-%{?3:%{3}-}ldsoconf.list} %{?3:%{3}-}core}\
%{!?_licensedir:%global license %%doc}\
%license linux-%{KVERREL}/COPYING-%{version}-%{release}\
/lib/modules/%{KVERREL}%{?3:+%{3}}/%{?-k:%{-k*}}%{!?-k:vmlinuz}\
%ghost /boot/%{?-k:%{-k*}}%{!?-k:vmlinuz}-%{KVERREL}%{?3:+%{3}}\
/lib/modules/%{KVERREL}%{?3:+%{3}}/.vmlinuz.hmac \
%ghost /boot/.vmlinuz-%{KVERREL}%{?3:+%{3}}.hmac \
%ifarch aarch64\
/lib/modules/%{KVERREL}%{?3:+%{3}}/dtb \
%ghost /boot/dtb-%{KVERREL}%{?3:+%{3}} \
%endif\
%attr(0600, root, root) /lib/modules/%{KVERREL}%{?3:+%{3}}/System.map\
%ghost %attr(0600, root, root) /boot/System.map-%{KVERREL}%{?3:+%{3}}\
/lib/modules/%{KVERREL}%{?3:+%{3}}/symvers.gz\
/lib/modules/%{KVERREL}%{?3:+%{3}}/config\
%ghost %attr(0600, root, root) /boot/symvers-%{KVERREL}%{?3:+%{3}}.gz\
%ghost %attr(0600, root, root) /boot/initramfs-%{KVERREL}%{?3:+%{3}}.img\
%ghost %attr(0644, root, root) /boot/config-%{KVERREL}%{?3:+%{3}}\
%dir /lib/modules\
%dir /lib/modules/%{KVERREL}%{?3:+%{3}}\
%dir /lib/modules/%{KVERREL}%{?3:+%{3}}/kernel\
/lib/modules/%{KVERREL}%{?3:+%{3}}/build\
/lib/modules/%{KVERREL}%{?3:+%{3}}/source\
/lib/modules/%{KVERREL}%{?3:+%{3}}/updates\
/lib/modules/%{KVERREL}%{?3:+%{3}}/weak-updates\
/lib/modules/%{KVERREL}%{?3:+%{3}}/systemtap\
%{_datadir}/doc/%{name}-keys/%{KVERREL}%{?3:+%{3}}\
%if %{1}\
/lib/modules/%{KVERREL}%{?3:+%{3}}/vdso\
%endif\
/lib/modules/%{KVERREL}%{?3:+%{3}}/modules.*\
%{expand:%%files -f %{name}-%{?3:%{3}-}modules.list %{?3:%{3}-}modules}\
%{expand:%%files %{?3:%{3}-}devel}\
%defverify(not mtime)\
/usr/src/kernels/%{KVERREL}%{?3:+%{3}}\
%{expand:%%files %{?3:%{3}-}devel-matched}\
%{expand:%%files -f %{name}-%{?3:%{3}-}modules-extra.list %{?3:%{3}-}modules-extra}\
%config(noreplace) /etc/modprobe.d/*-blacklist.conf\
%if %{?3:1} %{!?3:0}\
%{expand:%%files %{3}}\
%endif\
%endif\
%{nil}

%kernel_ml_variant_files %{_use_vdso} %{with_std}

%changelog
{{range $val := .Changelog}}* {{if $val.Subject}}{{$val.Subject}}{{else}}{{$val.Date}} {{$val.Name}} - {{$val.Version}}-{{$val.BuildID}}{{end}}
{{range $text := $val.Messages}}- {{$text}}
{{end}}
{{end}}
