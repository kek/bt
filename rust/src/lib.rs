extern crate libc;

use std::ffi::{c_char, CStr, CString};

extern {
    fn addTogether(_: u8, _: u8) -> u8;
}

#[no_mangle]
pub extern "C" fn bluetooth_upload(message: *const libc::c_char) {
    let message_cstr = unsafe { CStr::from_ptr(message) };
    let message = message_cstr.to_str().unwrap();
    println!("({})", message);
}

#[no_mangle]
pub extern "C" fn make_c_string() -> *const c_char {
    let cs = CString::new("Whatevery").expect("Failed to make C string");
    cs.as_ptr()
}

#[no_mangle]
pub extern "C" fn make_string() -> *const u8 {
    println!("{}", unsafe { addTogether(1,2) });
    "Whatevery\0".as_ptr()
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
