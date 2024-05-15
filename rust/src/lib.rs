use std::ffi::CStr;

#[no_mangle]
pub extern "C" fn bluetooth_upload(message: *const libc::c_char) {
    let message_cstr = unsafe { CStr::from_ptr(message) };
    let message = message_cstr.to_str().unwrap();
    println!("({})", message);
}
#[no_mangle]
pub extern "C" fn make_string() -> *const u8 {
    "Whatevery".as_ptr()
}

#[cfg(test)]
pub mod test {
    use std::ffi::CString;

    use super::*;

    // This is meant to do the same stuff as the main function in the .go files
    #[test]
    fn simulated_main_function () {
        bluetooth_upload(CString::new("undefined").unwrap().into_raw());
        // let actual = make_string().to_string().into_bytes();
        // let expected = "Hey".as_bytes();
        // assert_eq!(actual, expected);
    }
}
