#include <mylib/macros.h>
#include <mylib/defines.h>

int main() {
  scanll(h);

   debug(h);

   ll h_ = 0;
   rep(i, 60) {
      h_ += (ll)1 << i;

      if (h_ > h) {
         pl(i + 1);
         return 0;
      }
   }
}
