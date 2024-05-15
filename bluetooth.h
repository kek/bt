// NOTE: You could use https://michael-f-bryan.github.io/rust-ffi-guide/cbindgen.html to generate
// this header automatically from your Rust code.  But for now, we'll just write it by hand.

void hello(char *name);
void whisper(char *message);
void bluetooth_upload(char *program);
char *make_string();
