extern crate libc;

use std::ffi::CStr;

extern "C" {
    fn SendGoMessage(_: *const u8) -> ();
}

#[no_mangle]
pub extern "C" fn bluetooth_upload(message: *const libc::c_char) {
    let message_cstr = unsafe { CStr::from_ptr(message) };
    let message = message_cstr.to_str().unwrap();
    println!("({})", message);
}

#[no_mangle]
pub extern "C" fn make_string() -> *const u8 {
    "Whatevery\0".as_ptr()
}

#[no_mangle]
pub extern "C" fn send_me_a_message() {
    let s = "Hey there\0".as_ptr();
    unsafe { SendGoMessage(s) }
}

#[no_mangle]
pub extern "C" fn bluetooth_scan() {
    let s = "Hey there\0".as_ptr();
    unsafe { SendGoMessage(s) }
}

#[cfg(test)]
pub mod test {
    use std::ffi::CString;

    use super::*;

    #[test]
    fn simulated_main_function() {
        bluetooth_upload(CString::new("undefined").unwrap().into_raw());
    }
}
