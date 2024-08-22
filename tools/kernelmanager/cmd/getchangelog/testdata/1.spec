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
%global LKAver 1.0

# Define the buildid, if required.
%global buildid 123

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

Name: kernel
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

%changelog
* Tue Feb 22 2024 Mustafa Gezen - 1.0-123
- Test changelog
