package util

import (
	"os"
	"os/user"
	"runtime"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

const (
	_display_country_NDX = iota
	_display_language_NDX
	_display_script_NDX
	_display_variant_NDX
	_file_encoding_NDX
	_file_separator_NDX
	_format_country_NDX
	_format_language_NDX
	_format_script_NDX
	_format_variant_NDX
	_ftp_nonProxyHosts_NDX
	_ftp_proxyHost_NDX
	_ftp_proxyPort_NDX
	_http_nonProxyHosts_NDX
	_http_proxyHost_NDX
	_http_proxyPort_NDX
	_https_proxyHost_NDX
	_https_proxyPort_NDX
	_java_io_tmpdir_NDX
	_line_separator_NDX
	_os_arch_NDX
	_os_name_NDX
	_os_version_NDX
	_path_separator_NDX
	_socksNonProxyHosts_NDX
	_socksProxyHost_NDX
	_socksProxyPort_NDX
	_sun_arch_abi_NDX
	_sun_arch_data_model_NDX
	_sun_cpu_endian_NDX
	_sun_cpu_isalist_NDX
	_sun_io_unicode_encoding_NDX
	_sun_jnu_encoding_NDX
	_sun_os_patch_level_NDX
	_sun_stderr_encoding_NDX
	_sun_stdout_encoding_NDX
	_user_dir_NDX
	_user_home_NDX
	_user_name_NDX
	FIXED_LENGTH
)

func init() {
	sysPropsRaw(vmProperties, "vmProperties", "()[Ljava/lang/String;")
	sysPropsRaw(platformProperties, "platformProperties", "()[Ljava/lang/String;")
}

func sysPropsRaw(method native.Method, name, desc string) {
	native.Register("jdk/internal/util/SystemProps$Raw", name, desc, method)
}

// private static native String[] vmProperties();
func vmProperties(frame *rtda.Frame) {
	// TODO
	rt := frame.GetRuntime()
	jStrs := make([]*heap.Object, 2)
	jStrs[0] = rt.JSFromGoStr("java.home")
	jStrs[1] = rt.JSFromGoStr(frame.Thread.VMOptions.AbsJavaHome)
	frame.PushRef(frame.GetRuntime().NewStringArray(jStrs))
}

// private static native String[] platformProperties();
func platformProperties(frame *rtda.Frame) {
	rt := frame.GetRuntime()
	jStrs := make([]*heap.Object, FIXED_LENGTH)
	jStrs[_display_country_NDX] = rt.JSFromGoStr("")
	jStrs[_display_language_NDX] = rt.JSFromGoStr("")
	jStrs[_display_script_NDX] = rt.JSFromGoStr("")
	jStrs[_display_variant_NDX] = rt.JSFromGoStr("")
	jStrs[_file_encoding_NDX] = rt.JSFromGoStr("UTF-8")
	jStrs[_file_separator_NDX] = rt.JSFromGoStr(string(os.PathSeparator))
	jStrs[_format_country_NDX] = rt.JSFromGoStr("")
	jStrs[_format_language_NDX] = rt.JSFromGoStr("")
	jStrs[_format_script_NDX] = rt.JSFromGoStr("")
	jStrs[_format_variant_NDX] = rt.JSFromGoStr("")
	jStrs[_ftp_nonProxyHosts_NDX] = rt.JSFromGoStr("")
	jStrs[_ftp_proxyHost_NDX] = rt.JSFromGoStr("")
	jStrs[_ftp_proxyPort_NDX] = rt.JSFromGoStr("")
	jStrs[_http_nonProxyHosts_NDX] = rt.JSFromGoStr("")
	jStrs[_http_proxyHost_NDX] = rt.JSFromGoStr("")
	jStrs[_http_proxyPort_NDX] = rt.JSFromGoStr("")
	jStrs[_https_proxyHost_NDX] = rt.JSFromGoStr("")
	jStrs[_https_proxyPort_NDX] = rt.JSFromGoStr("")
	jStrs[_java_io_tmpdir_NDX] = rt.JSFromGoStr(os.TempDir())
	jStrs[_line_separator_NDX] = rt.JSFromGoStr("\n") // TODO
	jStrs[_os_arch_NDX] = rt.JSFromGoStr(runtime.GOARCH)
	jStrs[_os_name_NDX] = rt.JSFromGoStr(runtime.GOOS)
	jStrs[_os_version_NDX] = rt.JSFromGoStr("")
	jStrs[_path_separator_NDX] = rt.JSFromGoStr(string(os.PathListSeparator))
	jStrs[_socksNonProxyHosts_NDX] = rt.JSFromGoStr("")
	jStrs[_socksProxyHost_NDX] = rt.JSFromGoStr("")
	jStrs[_socksProxyPort_NDX] = rt.JSFromGoStr("")
	jStrs[_sun_arch_abi_NDX] = rt.JSFromGoStr("")
	jStrs[_sun_arch_data_model_NDX] = rt.JSFromGoStr("")
	jStrs[_sun_cpu_endian_NDX] = rt.JSFromGoStr("")
	jStrs[_sun_cpu_isalist_NDX] = rt.JSFromGoStr("")
	jStrs[_sun_io_unicode_encoding_NDX] = rt.JSFromGoStr("")
	jStrs[_sun_jnu_encoding_NDX] = rt.JSFromGoStr("")
	jStrs[_sun_os_patch_level_NDX] = rt.JSFromGoStr("")
	jStrs[_sun_stderr_encoding_NDX] = rt.JSFromGoStr("UTF-8")
	jStrs[_sun_stdout_encoding_NDX] = rt.JSFromGoStr("UTF-8")

	var userDir, userHome, userName string
	if dir, err := os.Getwd(); err != nil {
		userDir = dir
	}
	if user, err := user.Current(); err == nil {
		userHome = user.HomeDir
		userName = user.Name
	}
	jStrs[_user_dir_NDX] = rt.JSFromGoStr(userDir)
	jStrs[_user_home_NDX] = rt.JSFromGoStr(userHome)
	jStrs[_user_name_NDX] = rt.JSFromGoStr(userName)

	frame.PushRef(frame.GetRuntime().NewStringArray(jStrs))
}

/**
 * System properties. The following properties are guaranteed to be defined:
 * <dl>
 * <dt>java.version         <dd>Java version number
 * <dt>java.version.date    <dd>Java version date
 * <dt>java.vendor          <dd>Java vendor specific string
 * <dt>java.vendor.url      <dd>Java vendor URL
 * <dt>java.vendor.version  <dd>Java vendor version
 * <dt>java.home            <dd>Java installation directory
 * <dt>java.class.version   <dd>Java class version number
 * <dt>java.class.path      <dd>Java classpath
 * <dt>os.name              <dd>Operating System Name
 * <dt>os.arch              <dd>Operating System Architecture
 * <dt>os.version           <dd>Operating System Version
 * <dt>file.separator       <dd>File separator ("/" on Unix)
 * <dt>path.separator       <dd>Path separator (":" on Unix)
 * <dt>line.separator       <dd>Line separator ("\n" on Unix)
 * <dt>user.name            <dd>User account name
 * <dt>user.home            <dd>User home directory
 * <dt>user.dir             <dd>User's current working directory
 * </dl>
 */
