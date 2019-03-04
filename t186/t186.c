#include <stdlib.h> // getenv
#include <stdio.h>  // puts
#include <openssl/x509.h>

int main(void)
{
    const char *dir;

    dir = getenv(X509_get_default_cert_dir_env());
    if (!dir)
        dir = X509_get_default_cert_dir();

    puts(dir);
}
